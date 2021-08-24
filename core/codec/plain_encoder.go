package codec

type PlainEncoder struct{}

func (e *PlainEncoder) Encode(v interface{}) ([]byte, error) {
	return v.([]byte), nil
}
