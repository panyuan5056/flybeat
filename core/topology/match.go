package topology

type Match struct {
	Config map[string]interface{}
	Handle func(data string, config map[string]interface{}) map[string]interface{}
}

func (m *Match) Encoder(message string) map[string]interface{} {
	return m.Handle(message, m.Config)
}
