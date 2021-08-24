package topology

type Output struct {
	Messages chan map[string]interface{}
	Stop     bool
	Config   map[string]interface{}
}

func (o *Output) Push(message map[string]interface{}) {
	o.Messages <- message
}
