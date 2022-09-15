package main

import (
	"flag"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
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

func main() {
	var g = ProtoMessage{}

	var flags flag.FlagSet

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

	// find all rpc methods‘s input and output type interface
	var needImportInterfacesMap = mapset.NewSet[string]()
	var needImportInterfaces = make([]string, 20)

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
			needImportInterfacesMap.Add(string(method.Input.Desc.Name()))
			needImportInterfacesMap.Add(string(method.Output.Desc.Name()))
		}
		services = append(services, protoService)
	}
	needImportInterfacesMap.Each(func(a string) bool {
		needImportInterfaces = append(needImportInterfaces, string(a))
		return false
	})
	protoMessage.GenerateImportSourceCode(needImportInterfaces, fileName, t)
	protoMessage.GenerateGeneralServiceClass(t)
	for _, service := range services {
		if len(service.serviceName) > 0 {
			protoMessage.GenerateServiceClass(service.serviceName, service.methods, t)
		}
	}
	return nil
}

func (protoMessage *ProtoMessage) GenerateImportSourceCode(
	needImportInterfaces []string,
	fileName string,
	t *protogen.GeneratedFile,
) {
	t.P("import {")
	for _, interfaceName := range needImportInterfaces {
		if len(interfaceName) > 0 {
			t.P("  " + interfaceName + ",")
		}
	}
	slices := strings.Split(fileName, "/")
	var sliceslen = len(slices)

	t.P("}from \"./" + slices[sliceslen-1] + "\"")
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
