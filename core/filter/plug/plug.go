package plug

var HandleRender map[string]func(data map[string]interface{}, filed string, render []string) (interface{}, bool) = make(map[string]func(data map[string]interface{}, filed string, render []string) (interface{}, bool))

func Register(category string, fn func(data map[string]interface{}, filed string, render []string) (interface{}, bool)) {
	HandleRender[category] = fn
}
