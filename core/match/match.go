package match

var HandleMatchs map[string]func(data string, config map[string]interface{}) map[string]interface{} = make(map[string]func(data string, config map[string]interface{}) map[string]interface{})

func Register(category string, fn func(data string, config map[string]interface{}) map[string]interface{}) {
	HandleMatchs[category] = fn
}
