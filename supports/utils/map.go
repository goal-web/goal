package utils

func GetMapKeys(data map[string]interface{}) (keys []string) {
	for key, _ := range data {
		keys = append(keys, key)
	}
	return
}
