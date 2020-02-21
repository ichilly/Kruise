package utils

func GetAllKeys(m map[string]string) []string {
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func GetAllValues(m map[string]string) []string {
	values := []string{}
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

func GetFirstKey(m map[string]string) string {
	for key := range m {
		return key
	}
	return ""
}

func GetFirstValue(m map[string]string) string {
	for _, value := range m {
		return value
	}
	return ""
}
