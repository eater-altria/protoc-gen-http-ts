package util

import (
	"regexp"
	"strings"
	"syscall"
	"unicode"
)

// IsContainInt Check if an element exists in an array
func IsContainInt(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
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
	PascalCase NameStyle = 0
	CamelCase  NameStyle = 1
	SnakeCase  NameStyle = 2
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
