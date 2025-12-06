package cipher

type Header struct {
	header []byte
}

type Payload struct {
	rawData []byte
}

type Frame struct {
	*Header
	*Payload
}
