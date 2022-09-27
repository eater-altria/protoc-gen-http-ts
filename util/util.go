package util

import "strings"

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
