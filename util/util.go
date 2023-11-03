package util

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"regexp"
	"strings"
	"syscall"
	"unicode"

	"google.golang.org/protobuf/compiler/protogen"
)

type InterfaceName struct {
	RealName string
	FullName string
}

// IsContainInt Check if an element exists in an array
func IsContainInt(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func ConvertToUnderscore(str string) string {
	// replace '.' with '_'
	str = strings.ReplaceAll(str, ".", "_")

	return str
}

func IsContainInterfaceName(items []InterfaceName, item string) bool {
	for _, eachItem := range items {
		if eachItem.RealName == item {
			return true
		}
	}
	return false
}

// ReverseSlice reverse an array
func ReverseSlice(slice []string) []string {
	var sliceReversed []string
	var sliceLen = len(slice)
	for i := sliceLen - 1; i >= 0; i-- {
		sliceReversed = append(sliceReversed, slice[i])
	}
	return sliceReversed
}

// GetRelativePath Calculate relative path of two files
func GetRelativePath(pathA string, pathB string) string {
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

type NameStyle int32

const (
	CamelCase  NameStyle = 0
	PascalCase NameStyle = 1
	SnakeCase  NameStyle = 2
	UNKNOWN    NameStyle = 3
)

func (nameStyle NameStyle) String() string {
	switch nameStyle {
	case PascalCase:
		return "pascal"
	case CamelCase:
		return "camel"
	case SnakeCase:
		return "snake"
	default:
		return "UNKNOWN"
	}
}

func TransStringToNameStyle(name string) NameStyle {
	if PascalCase.String() == name {
		return PascalCase
	} else if CamelCase.String() == name {
		return CamelCase
	} else if SnakeCase.String() == name {
		return SnakeCase
	}
	return UNKNOWN
}

func transPascalToCamel(oldName string) string {
	var camelName = ""
	for index, s := range oldName {
		if index == 0 {
			camelName = camelName + string(unicode.ToLower(s))
		} else {
			camelName = camelName + string(s)
		}
	}
	return camelName
}

func transPascalToSnake(oldName string) string {
	var snakeName = ""
	for index, s := range oldName {
		if index == 0 {
			snakeName = snakeName + string(unicode.ToLower(s))
		} else {
			if unicode.ToLower(s) != s {
				snakeName = snakeName + "_" + string(unicode.ToLower(s))
			} else {
				snakeName = snakeName + string(s)
			}
		}
	}
	return snakeName
}

func transCamelToPascal(oldName string) string {
	var pascalName = ""
	for index, s := range oldName {
		if index == 0 {
			pascalName = pascalName + string(unicode.ToUpper(s))
		} else {
			pascalName = pascalName + string(s)
		}
	}
	return pascalName
}
func transCamelToSnake(oldName string) string {
	var snakeName = ""
	for _, s := range oldName {
		if s == unicode.ToLower(s) {
			snakeName = snakeName + string(unicode.ToLower(s))
		} else {
			snakeName = snakeName + "_" + string(unicode.ToLower(s))
		}
	}
	return snakeName
}

func transSnakeToPascal(oldName string) string {
	var flag = false
	var pascalName = ""
	for index, s := range oldName {
		if index == 0 {
			pascalName = pascalName + string(unicode.ToUpper(s))
		} else {
			if string(s) == "_" {
				flag = true
			} else {
				if flag {
					pascalName = pascalName + string(unicode.ToUpper(s))
					flag = false
				} else {
					pascalName = pascalName + string(s)
				}
			}
		}
	}
	return pascalName
}

func transSnakeToCamel(oldName string) string {
	var flag = false
	var pascalName = ""
	for index, s := range oldName {
		if index == 0 {
			pascalName = pascalName + string(s)
		} else {
			if string(s) == "_" {
				flag = true
			} else {
				if flag {
					pascalName = pascalName + string(unicode.ToUpper(s))
					flag = false
				} else {
					pascalName = pascalName + string(s)
				}
			}
		}
	}
	return pascalName
}

func TransformNameStyle(oldName string, targetStyle NameStyle) (string, error) {
	var isMatchPascal, _ = regexp.MatchString("^[A-Z][a-z]+", oldName)
	var isMatchCamel, _ = regexp.MatchString("^[a-z]+", oldName)
	var isMatchSnake, _ = regexp.MatchString("[a-z]+_[a-z]+", oldName)

	if isMatchSnake {
		if targetStyle.String() == "camel" {
			return transSnakeToCamel(oldName), nil
		} else if targetStyle.String() == "snake" {
			return oldName, nil
		} else if targetStyle.String() == "pascal" {
			return transSnakeToPascal(oldName), nil
		} else {
			return "", syscall.Errno(1)
		}
	} else if isMatchPascal {
		if targetStyle.String() == "camel" {
			return transPascalToCamel(oldName), nil
		} else if targetStyle.String() == "snake" {
			return transPascalToSnake(oldName), nil
		} else if targetStyle.String() == "pascal" {
			return oldName, nil
		} else {
			return "", syscall.Errno(1)
		}
	} else if isMatchCamel {
		if targetStyle.String() == "camel" {
			return oldName, nil
		} else if targetStyle.String() == "snake" {
			return transCamelToSnake(oldName), nil
		} else if targetStyle.String() == "pascal" {
			return transCamelToPascal(oldName), nil
		} else {
			return "", syscall.Errno(1)
		}
	}

	return "", syscall.Errno(1)
}

// GenerateComment 生成 typescript 注释
func GenerateComment(methodCommentSet protogen.CommentSet, deprecated bool, prefix string) string {
	var lines []string

	// 顶部注释
	leadingStr := methodCommentSet.Leading.String()

	// 尾注释
	trailingStr := methodCommentSet.Trailing.String()

	// 有注释才操作
	if len(leadingStr) > 0 || len(trailingStr) > 0 {
		// 去掉 *
		content := strings.Replace(strings.TrimSpace(leadingStr+trailingStr), "*", "", -1)
		// 去掉 斜杠 //
		content = strings.Replace(content, "/", "", -1)

		// 前缀
		if prefix != "" {
			content = prefix + content
		}

		lines = strings.Split(content, "\n")
		for i := range lines {
			lines[i] = strings.TrimSpace(strings.TrimLeft(lines[i], " "))
			// 注意需要2个换行符，才会在IDE里面显示换行
			lines[i] = strings.TrimRight(lines[i], "\n")
		}
	}

	// 如果废弃函数，增加 jsDoc 语义化注释
	if deprecated {
		if len(lines) > 0 {
			lines = append(lines, "")
		}
		lines = append(lines, "@deprecated")
	}

	var commentStr string
	if len(lines) == 1 {
		commentStr = string((fmt.Sprintf("/** %s */", lines[0])))
	} else if len(lines) > 1 {
		// 这里为了跟输出代码对齐
		commentStr = string(fmt.Sprintf("/**\n"+"   "+"* %s \n   */", strings.Join(lines, "\n   * \n   * ")))
	}

	return commentStr
}

// GetFiledTypeConfig 递归遍历获取 message 所有字段的类型
func traversalFieldType(fields []*protogen.Field, prefix string, outputMap map[string]interface{}) {
	for _, field := range fields {
		filedKey := prefix

		// 转换成驼峰命名规范
		camelGoName := transPascalToCamel(field.GoName)

		if len(prefix) == 0 {
			filedKey = filedKey + camelGoName
		} else {
			filedKey = filedKey + "." + camelGoName
		}

		if field.Desc.Kind() == protoreflect.MessageKind {
			traversalFieldType(field.Message.Fields, filedKey, outputMap)
		} else {
			// 做一下约束，否则太多字段生成。
			if field.Desc.Kind() == protoreflect.Int64Kind {
				outputMap[filedKey] = field.Desc.Kind().String()
			}
		}
	}
}

// GetFiledTypeConfig 获取 message 所有字段的类型
func GetFiledTypeConfig(fields []*protogen.Field, prefix string) map[string]interface{} {
	outputMap := map[string]interface{}{}

	traversalFieldType(fields, prefix, outputMap)

	return outputMap
}

// GenerateFormatObject 生成格式化字符串
func GenerateFormatObject(config map[string]interface{}, varName string) (string, string) {
	newVarName := fmt.Sprintf("%sFiled", varName)

	var code = []string{
		fmt.Sprintf("  const %s = {", newVarName),
	}

	// 遍历字段定义配置
	for key, value := range config {
		keyValueStr := fmt.Sprintf("\"%s\": \"%v\",", key, value)
		code = append(code, keyValueStr)
	}

	code = append(code, "};")

	codeStr := strings.Join(code, "\n      ")

	return newVarName, codeStr
}
