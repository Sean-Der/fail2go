package fail2go

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
