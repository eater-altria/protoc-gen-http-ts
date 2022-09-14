package main

import (
	"flag"
	//"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	//"fmt"
	//"github.com/ditashi/jsbeautifier-go/jsbeautifier"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	//"strings"
)

//import vector "container/vector"

type ProtoMessage struct {
	Prefix   string
	messages []protoreflect.FullName
	//sevicer protoreflect.ServiceDescriptors
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
	var needImportInterfacesMap = mapset.NewSet[string]()
	var needImportInterfaces = make([]string, 10)
	var allMethods = mapset.NewSet[string]()
	var fileName = strings.Replace(protoFile.Desc.Path(), ".proto", "", 1)
	var generateFileName = fileName + ".http.ts"
	var path = protogen.GoImportPath(protoFile.Desc.Path())
	var t = plugin.NewGeneratedFile(generateFileName, path)
	for _, service := range protoFile.Services {
		var methods = service.Methods
		for _, method := range methods {
			needImportInterfacesMap.Add(string(method.Input.Desc.Name()))
			needImportInterfacesMap.Add(string(method.Output.Desc.Name()))
			allMethods.Add(string(method.Desc.Name()))
		}
	}
	needImportInterfacesMap.Each(func(a string) bool {
		needImportInterfaces = append(needImportInterfaces, string(a))
		return false
	})
	protoMessage.GenerateImportSourceCode(needImportInterfaces, t, fileName)
	return nil
}

func (protoMessage *ProtoMessage) GenerateImportSourceCode(
	needImportInterfaces []string,
	t *protogen.GeneratedFile,
	fileName string,
) {
	t.P("import {")
	for _, interfaceName := range needImportInterfaces {
		if len(interfaceName) > 0 {
			t.P("  " + interfaceName + ",")
		}
	}
	slices := strings.Split(fileName, "/")
	var sliceslen = len(slices)

	t.P("}from \"./" + slices[sliceslen-1] + ".ts\"")
}

func (protoMessage *ProtoMessage) GenerateServiceClass(
	serviceName string,
	serviceSlice []string,

) {

}
