package main

import (
	"flag"
	"fmt"
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

type ProtoService struct {
	serviceName string
	methods     []*protogen.Method
}

type ImportedFile struct {
	path protoreflect.SourcePath
	name []protoreflect.Name
}

func main() {
	var g = ProtoMessage{}

	var flags flag.FlagSet
	fmt.Println("xxx")

	flags.StringVar(&g.Prefix, "prefix", "/", "API path prefix")

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(g.Generate)
}

func (protoMessage *ProtoMessage) Generate(plugin *protogen.Plugin) error {
	var protoFiles = plugin.Files
	for _, file := range protoFiles {
		protoMessage.GenerateOneFile(file, plugin)
	}

	return nil
}

func (protoMessage *ProtoMessage) GenerateOneFile(protoFile *protogen.File, plugin *protogen.Plugin) error {
	if len(protoFile.Services) == 0 {
		return nil
	}

	var importInfo map[string][]string
	importInfo = make(map[string][]string)

	// fileName is just like ./test_protos/test. Remove prefix of .proto
	var fileName = strings.Replace(protoFile.Desc.Path(), ".proto", "", 1)
	// add .http.ts prefix
	var generateFileName = fileName + ".http.ts"

	var services = make([]ProtoService, 10)

	var path = protogen.GoImportPath(protoFile.Desc.Path())
	var t = plugin.NewGeneratedFile(generateFileName, path)
	for _, service := range protoFile.Services {
		var protoService ProtoService
		protoService.serviceName = string(service.Desc.Name())
		var methods = service.Methods
		for _, method := range methods {
			protoService.methods = append(protoService.methods, method)
			var inputImportPath = string(method.Input.Location.SourceFile)
			var input = string(method.Input.Desc.Name())
			var inputInterfaces = importInfo[inputImportPath]
			if !IsContainInt(inputInterfaces, input) {
				inputInterfaces = append(inputInterfaces, input)
			}
			importInfo[inputImportPath] = inputInterfaces
			var outputImportPath = string(method.Output.Location.SourceFile)
			var output = string(method.Output.Desc.Name())
			var outputInterfaes = importInfo[outputImportPath]
			if !IsContainInt(outputInterfaes, output) {
				outputInterfaes = append(outputInterfaes, output)
			}
			importInfo[outputImportPath] = outputInterfaes

		}
		services = append(services, protoService)
	}
	protoMessage.GenerateImportSourceCodeV2(importInfo, fileName, t)
	protoMessage.GenerateGeneralServiceClass(t)
	for _, service := range services {
		if len(service.serviceName) > 0 {
			protoMessage.GenerateServiceClass(service.serviceName, service.methods, t)
		}
	}
	return nil
}

func (protoMessage *ProtoMessage) GenerateImportSourceCodeV2(
	needImportInterfaces map[string][]string,
	sourcePath string,
	t *protogen.GeneratedFile,
) {
	for path, interfaces := range needImportInterfaces {
		// t.P("----path----")
		// t.P(path)
		// t.P("----interfaces----")
		// t.P(interfaces)
		t.P("import {")
		for _, interfae := range interfaces {
			if len(interfae) > 0 {
				t.P("  " + interfae + ",")
			}
		}
		var relativePath = getRelativePath(sourcePath, path)
		t.P("}from \"" + strings.Replace(relativePath, ".proto", "", 1) + "\"")
	}
	t.P("")
}
func (protoMessage *ProtoMessage) GenerateGeneralServiceClass(
	t *protogen.GeneratedFile,
) {
	t.P("export type GeneralRequest = <TReq, TResp>(TReq, cmd: string, options?: any) => Promise<TResp>")
	t.P("")
	t.P("export class GeneralClass {")
	t.P("  GeneralRequestMethod: GeneralRequest;")
	t.P("  constructor(GeneralRequestMethod: GeneralRequest) {")
	t.P("    this.GeneralRequestMethod = GeneralRequestMethod;")
	t.P("  }")
	t.P("}")
	t.P("")
}

func (protoMessage *ProtoMessage) GenerateServiceClass(
	serviceName string,
	methods []*protogen.Method,
	t *protogen.GeneratedFile,
) {
	t.P("export class " + serviceName + " extends GeneralClass {")
	t.P("  constructor(GeneralRequestMethod: GeneralRequest) {")
	t.P("    super(GeneralRequestMethod)")
	t.P("  }")
	for _, method := range methods {
		var name = method.Desc.Name()
		var input = method.Input.Desc.Name()
		var output = method.Output.Desc.Name()
		t.P("  " + name + "(payload: " + input + ", options?: any): Promise<" + output + "> {")
		t.P("    return new Promise((resolve, reject) => {")
		t.P("      this.GeneralRequestMethod<" + input + "," + output + ">(payload, '" + name + "', options).then(res => {")
		t.P("        resolve(res)")
		t.P("      }).catch(error => {")
		t.P("        reject(error)")
		t.P("      })")
		t.P("    })")
		t.P("  }")
	}
	t.P("}")
	t.P("")
}

func IsContainInt(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func ReverseSlice(slice []string) []string {
	var sliceReversed []string
	var sliceLen = len(slice)
	for i := sliceLen - 1; i >= 0; i-- {
		sliceReversed = append(sliceReversed, slice[i])
	}
	return sliceReversed
}

func getRelativePath(pathA string, pathB string) string {
	var pathASlice = strings.Split(pathA, "/")
	var pathBSlice = strings.Split(pathB, "/")
	pathASlice = ReverseSlice(pathASlice)
	var res = ""
	for i, _ := range pathASlice {
		if i == 0 {
			res = res + "./"
		} else {
			res = res + "../"
		}
	}

	for i, v := range pathBSlice {
		if i != len(pathBSlice)-1 {
			res = res + v + "/"
		} else {
			res = res + v
		}
	}

	return res
}
