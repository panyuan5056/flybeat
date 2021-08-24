package topology

type Input struct {
	Messages chan map[string]interface{}
	Stop     bool
	Config   map[string]interface{}
}

func (i *Input) OnEventBar() (map[string]interface{}, bool) {
	event, ok := <-i.Messages
	return event, ok
}

func (i *Input) Len() int {
	return len(i.Messages)
}
