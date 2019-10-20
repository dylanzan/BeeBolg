package utils

import "strings"

func HandleCategoryAddStr(str string) (res []string) {

	res = strings.Split(str, " ")

	return res
}
