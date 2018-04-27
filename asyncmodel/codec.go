package asyncmodel

import "bytes"

type Encode interface {
	Encode() *bytes.Buffer
}

type Decode interface {
	Decode(buf *bytes.Buffer)
}
