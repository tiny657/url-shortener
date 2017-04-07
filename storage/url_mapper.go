package storage

var pair map[string]string

func init() {
	pair = make(map[string]string)
}

func GetUrl(key string) string {
	value, _ := pair[key]
	return value
}

func SetUrl(key string, value string) {
	pair[key] = value
}
