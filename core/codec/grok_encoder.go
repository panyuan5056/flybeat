package codec

import (
	"github.com/vjeantet/grok"
)

type GrokEncoder struct{}

func (e *GrokEncoder) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
