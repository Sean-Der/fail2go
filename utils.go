package fail2go

import (
	ogórek "github.com/kisielk/og-rek"
)

func interfaceSliceToStringSlice(interfaceSlice []interface{}) (stringSlice []string) {
	for _, v := range interfaceSlice {
		stringSlice = append(stringSlice, v.(string))
	}
	return
}

func stringInSliceIndex(input string, list []string) int {
	for index, value := range list {
		if value == input {
			return index
		}
	}
	return -1
}

func callSliceToStringSlice(callSlice []interface{}) (stringSlice []string) {
	for _, v := range callSlice {
		stringSlice = append(stringSlice, v.(ogórek.Call).Args[0].(string))
	}
	return
}
