package codec

type Encoder interface {
	Encode(interface{}) ([]byte, error)
}

func NewEncoder(t string) Encoder {
	switch t {
	case "json":
		return &JsonEncoder{}
	case "plain":
		return &PlainEncoder{}
	case "grok":
		return GrokEncoder{}
	}
	return nil

}
