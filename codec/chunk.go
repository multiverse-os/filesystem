package codec

type Chunk struct {
	contents []byte
	digest   []byte
}

func NewChunk(contents, digest []byte) *Chunk {
	return &Chunk{contents: contents, digest: digest}
}
