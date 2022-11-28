package main

import (
	"protoc-gen-http/util"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//import vector "container/vector"

type ProtoMessage struct {
	nameCase util.NameStyle
	messages []protoreflect.FullName
	//service protector.ServiceDescriptors
}

// ProtoService rpc Service struct
type ProtoService struct {
	serviceName string             // rpc service name
	methods     []*protogen.Method // rpc method list
}

// ImportedFile dependency info
// for example: import {a, b} from './search_service'
// ./search_service is path, [a, b] is names
type ImportedFile struct {
	path protoreflect.SourcePath //imported file's path
	name []protoreflect.Name     // imported ts interfaces name
}

func main() {
	var g = ProtoMessage{}
	protogen.Options{
		ParamFunc: getCompileOption(&g),
	}.Run(g.Generate)
}

func getCompileOption(g *ProtoMessage) func(key string, value string) error {
	var setFunc = func(key string, value string) error {
		// set option values
		if key == "nameCase" {
			g.nameCase = util.TransStringToNameStyle(value)
		}
		return nil
	}
	return setFunc
}

// Generate main func
func (protoMessage *ProtoMessage) Generate(plugin *protogen.Plugin) error {
	if protoMessage.nameCase == util.UNKNOWN {
		return nil
	}
	// get all files need be compiled
	var protoFiles = plugin.Files
	// compile one by one
	for _, file := range protoFiles {
		err := protoMessage.GenerateOneFile(file, plugin)
		if err != nil {
			return err
		}
	}

	return nil
}

// GenerateOneFile compile one proto file
func (protoMessage *ProtoMessage) GenerateOneFile(protoFile *protogen.File, plugin *protogen.Plugin) error {
	// if the proto file don't have service
	// skip the proto.
	var code []string
	if len(protoFile.Services) == 0 {
		return nil
	}

	// save ts import info，
	// key is path, value is interfaces which are imported from the path
	var importInfo map[string][]string
	importInfo = make(map[string][]string)

	// fileName is just like ./test_protos/test. Remove prefix of .proto
	var fileName = strings.Replace(protoFile.Desc.Path(), ".proto", "", 1)
	// add .http.ts prefix
	var generateFileName = fileName + ".http.ts"

	// collect all services from the proto file
	var services []ProtoService

	// generated file's path
	var path = protogen.GoImportPath(protoFile.Desc.Path())

	// t.P() can write message to generated file
	var t = plugin.NewGeneratedFile(generateFileName, path)

	for _, service := range protoFile.Services {
		var protoService ProtoService
		protoService.serviceName = string(service.Desc.Name())
		var methods = service.Methods
		for _, method := range methods {
			// add method to corresponding service struct
			protoService.methods = append(protoService.methods, method)

			// inputImportPath is the path where the method input message from
			var inputImportPath = method.Input.Location.SourceFile
			// input is the method input message's name
			var input = string(method.Input.Desc.Name())
			// get the imported interfaces from inputImportPath
			var inputInterfaces = importInfo[inputImportPath]
			// if inputInterfaces don't have current interface, push it
			if !util.IsContainInt(inputInterfaces, input) {
				inputInterfaces = append(inputInterfaces, input)
			}
			// update interfaces
			importInfo[inputImportPath] = inputInterfaces

			// do the same thing to output message.
			var outputImportPath = method.Output.Location.SourceFile
			var output = string(method.Output.Desc.Name())
			var outputInterfaces = importInfo[outputImportPath]
			if !util.IsContainInt(outputInterfaces, output) {
				outputInterfaces = append(outputInterfaces, output)
			}
			importInfo[outputImportPath] = outputInterfaces

		}

		// collect services info.
		services = append(services, protoService)
	}

	// generate codes
	code = append(code, protoMessage.GenerateImportSourceCode(importInfo, fileName)...)
	code = append(code, protoMessage.GenerateGeneralServiceClass()...)
	for _, service := range services {
		if len(service.serviceName) > 0 {
			var serviceCode, err = protoMessage.GenerateServiceClass(service.serviceName, service.methods)
			if err == nil {
				code = append(code, serviceCode...)
			} else {
				return err
			}
		}
	}
	t.P(strings.Join(code[:], "\n"))
	return nil
}

// GenerateImportSourceCode generate code such as `import {xxx} from "xxx"`
/**
* needImportInterfaces: all interfaces need to be imported
sourcePath: generate code's path
*/
func (protoMessage *ProtoMessage) GenerateImportSourceCode(
	needImportInterfaces map[string][]string,
	sourcePath string,
) []string {
	var code []string
	code = append(code, "/* eslint-disable */")

	for path, interfaces := range needImportInterfaces {
		code = append(code, "import {")
		for _, interfae := range interfaces {
			if len(interfae) > 0 {
				code = append(code, "  "+interfae+",")
			}
		}
		// relativePath is based on sourcePath and the interface's path,
		var relativePath = util.GetRelativePath(sourcePath, path)
		code = append(code, "} from '"+strings.Replace(relativePath, ".proto", "", 1)+"';")
	}
	code = append(code, "")
	return code
}

// GenerateGeneralServiceClass generate GeneralServiceClass
// the class is extended by all other service.
func (protoMessage *ProtoMessage) GenerateGeneralServiceClass() []string {
	var code = []string{
		"export type GeneralRequest = <TReq, TResp>(cmd: string, payload: TReq, options?: any) => Promise<TResp>;",
		"",
		"export class GeneralClass {",
		"  generalRequestMethod: GeneralRequest;",
		"  constructor(generalRequestMethod: any) {",
		"    this.generalRequestMethod = generalRequestMethod as GeneralRequest;",
		"  };",
		"};",
		"",
	}
	return code
}

// GenerateServiceClass generate one rpc Service
func (protoMessage *ProtoMessage) GenerateServiceClass(
	serviceName string,
	methods []*protogen.Method,
) ([]string, error) {
	var code = []string{
		"export class " + serviceName + " extends GeneralClass {",
	}
	for _, method := range methods {
		var name = method.Desc.Name()
		transformdName, err := util.TransformNameStyle(string(name), protoMessage.nameCase)
		if err != nil {
			return make([]string, 0), err
		}
		var input = string(method.Input.Desc.Name())
		var output = string(method.Output.Desc.Name())
		code = append(code, "  "+transformdName+"(payload: "+input+", options?: any): Promise<"+output+"> {")
		code = append(code, "    return new Promise((resolve, reject) => {")
		code = append(code, "      this.generalRequestMethod<"+input+", "+output+">('"+transformdName+"', payload, "+"options).then((res) => {")
		code = append(code, "        resolve(res);")
		code = append(code, "      })")
		code = append(code, "        .catch((error) => {")
		code = append(code, "          reject(error);")
		code = append(code, "        });")
		code = append(code, "    });")
		code = append(code, "  };")
	}
	code = append(code, "};")
	code = append(code, "")

	return code, nil
}
