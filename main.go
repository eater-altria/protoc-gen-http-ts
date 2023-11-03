package main

import (
	"strings"

	"github.com/eater-altria/protoc-gen-http-ts/util"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"
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
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
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
	var importInfo map[string][]util.InterfaceName
	importInfo = make(map[string][]util.InterfaceName)

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
			var input util.InterfaceName
			input.RealName = string(method.Input.Desc.Name())
			input.FullName = string(method.Input.Desc.FullName())
			// get the imported interfaces from inputImportPath
			var inputInterfaces = importInfo[inputImportPath]
			// if inputInterfaces don't have current interface, push it
			if !util.IsContainInterfaceName(inputInterfaces, input.RealName) {
				inputInterfaces = append(inputInterfaces, input)
			}
			// update interfaces
			importInfo[inputImportPath] = inputInterfaces

			// do the same thing to output message.
			var outputImportPath = method.Output.Location.SourceFile
			var output util.InterfaceName
			output.RealName = string(method.Output.Desc.Name())
			output.FullName = string(method.Output.Desc.FullName())
			var outputInterfaces = importInfo[outputImportPath]
			if !util.IsContainInterfaceName(outputInterfaces, output.RealName) {
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
	needImportInterfaces map[string][]util.InterfaceName,
	sourcePath string,
) []string {
	var code []string
	code = append(code, "/* eslint-disable */")

	for path, interfaces := range needImportInterfaces {
		code = append(code, "import {")
		for _, eachInterface := range interfaces {
			if len(eachInterface.RealName) > 0 {
				code = append(code, "  "+eachInterface.RealName+" as "+util.ConvertToUnderscore(eachInterface.FullName)+",")
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
		"export type IFormatFn = <TResp>(res: Record<string, any>, config: Record<string, any>) => TResp;",
		"",
		"export class GeneralClass {",
		"  generalRequestMethod: GeneralRequest;",
		"  formatFn: IFormatFn;",
		"  constructor(generalRequestMethod: any, formatFn: IFormatFn) {",
		"    this.generalRequestMethod = generalRequestMethod as GeneralRequest;",
		"    this.formatFn = formatFn as IFormatFn;",
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

		// 获取message filed 的类型对象
		filedTypeConfig := util.GetFiledTypeConfig(method.Output.Fields, "")

		// 后台需要 帕斯卡命名规则
		transformdName, err := util.TransformNameStyle(string(name), util.PascalCase)

		// 后端请求 url 命名规则：package + serviceName + transformdName
		// 举例： /lixin.dos.boss_bff.v1.BossBffService/CreateRole
		requestMethodName := serviceName + "/" + transformdName

		if err != nil {
			return make([]string, 0), err
		}

		var input = util.ConvertToUnderscore(string(method.Input.Desc.FullName()))
		var output = util.ConvertToUnderscore(string(method.Output.Desc.FullName()))

		// protoc 里面的注释
		commentStr := util.GenerateComment(method.Comments, false, "")
		typeVarName, filedTypeConfigStr := util.GenerateFormatObject(filedTypeConfig, transformdName)

		// 增加注释
		code = append(code, "  "+commentStr)
		code = append(code, "  "+transformdName+"(payload: "+input+", options?: any): Promise<"+output+"> {")
		code = append(code, "  "+filedTypeConfigStr)
		code = append(code, "    return new Promise((resolve, reject) => {")
		code = append(code, "      this.generalRequestMethod<"+input+", "+output+">('"+requestMethodName+"', payload, "+"options).then((res) => {")
		code = append(code, "        resolve(this.formatFn(res,"+typeVarName+"));")
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
