package helper

// MyReader emits an infinite stream of the ASCII character 'A'
type MyReader struct{}

func (MyReader) Read(b []byte) (n int, err error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}
