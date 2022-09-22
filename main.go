package main

import (
	"flag"
	"protoc-gen-http/util"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//import vector "container/vector"

type ProtoMessage struct {
	Prefix   string
	messages []protoreflect.FullName
	//sevicer protoreflect.ServiceDescriptors
}

// rpc Service struct
type ProtoService struct {
	serviceName string             // rpc service name
	methods     []*protogen.Method // rpc method list
}

// dependency info
// for example: import {a, b} from './search_service'
// ./serch_service is path, [a, b] is name
type ImportedFile struct {
	path protoreflect.SourcePath //imported file's path
	name []protoreflect.Name     // imported ts interfaces name
}

func main() {
	var g = ProtoMessage{}

	var flags flag.FlagSet

	flags.StringVar(&g.Prefix, "prefix", "/", "API path prefix")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(g.Generate)
}

// Generate main func
func (protoMessage *ProtoMessage) Generate(plugin *protogen.Plugin) error {
	// get all files need be compiled
	var protoFiles = plugin.Files
	// compile one by one
	for _, file := range protoFiles {
		protoMessage.GenerateOneFile(file, plugin)
	}

	return nil
}

// compile one proto file
func (protoMessage *ProtoMessage) GenerateOneFile(protoFile *protogen.File, plugin *protogen.Plugin) error {
	// if the proto file don't have service
	// skip the proto.
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
			var inputImportPath = string(method.Input.Location.SourceFile)
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
			var outputImportPath = string(method.Output.Location.SourceFile)
			var output = string(method.Output.Desc.Name())
			var outputInterfaes = importInfo[outputImportPath]
			if !util.IsContainInt(outputInterfaes, output) {
				outputInterfaes = append(outputInterfaes, output)
			}
			importInfo[outputImportPath] = outputInterfaes

		}

		// collect services info.
		services = append(services, protoService)
	}

	// generate codes
	protoMessage.GenerateImportSourceCode(importInfo, fileName, t)
	protoMessage.GenerateGeneralServiceClass(t)
	for _, service := range services {
		if len(service.serviceName) > 0 {
			protoMessage.GenerateServiceClass(service.serviceName, service.methods, t)
		}
	}
	return nil
}

// generate code such as `import {xxx} from "xxx"`
/**
* needImportInterfaces: all interfaces need to be imported
sourcePath: generate code's path
*/
func (protoMessage *ProtoMessage) GenerateImportSourceCode(
	needImportInterfaces map[string][]string,
	sourcePath string,
	t *protogen.GeneratedFile,
) {
	t.P("/* eslint-disable */")
	for path, interfaces := range needImportInterfaces {
		t.P("import {")
		for _, interfae := range interfaces {
			if len(interfae) > 0 {
				t.P("  " + interfae + ",")
			}
		}
		// relativePath is based on sourcePath and the interface's path,
		var relativePath = util.GetRelativePath(sourcePath, path)
		t.P("} from '" + strings.Replace(relativePath, ".proto", "", 1) + "';")
	}
	t.P("")
}

// generate GeneralServiceClass
// the class is extended by all other service.
func (protoMessage *ProtoMessage) GenerateGeneralServiceClass(
	t *protogen.GeneratedFile,
) {
	t.P("export type GeneralRequest = <TReq, TResp>(cmd: string, payload: TReq, options?: any) => Promise<TResp>;")
	t.P("")
	t.P("export class GeneralClass {")
	t.P("  generalRequestMethod: GeneralRequest;")
	t.P("  constructor(generalRequestMethod: GeneralRequest) {")
	t.P("    this.generalRequestMethod = generalRequestMethod;")
	t.P("  };")
	t.P("};")
	t.P("")
}

// generate one rpc Service
func (protoMessage *ProtoMessage) GenerateServiceClass(
	serviceName string,
	methods []*protogen.Method,
	t *protogen.GeneratedFile,
) {
	t.P("export class " + serviceName + " extends GeneralClass {")
	for _, method := range methods {
		var name = method.Desc.Name()
		var input = method.Input.Desc.Name()
		var output = method.Output.Desc.Name()
		t.P("  " + name + "(payload: " + input + ", options?: any): Promise<" + output + "> {")
		t.P("    return new Promise((resolve, reject) => {")
		t.P("      this.generalRequestMethod<" + input + ", " + output + ">('" + name + "', payload, " + "options).then((res) => {")
		t.P("        resolve(res);")
		t.P("      })")
		t.P("        .catch((error) => {")
		t.P("          reject(error);")
		t.P("        });")
		t.P("    });")
		t.P("  };")
	}
	t.P("};")
	t.P("")
}
