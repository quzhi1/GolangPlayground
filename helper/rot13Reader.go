package helper

import (
	"io"
	"os"
	"strings"
)

// Rot13Reader reads a string but emit the rot13 encoding
type Rot13Reader struct {
	R io.Reader
}

func (rr *Rot13Reader) Read(b []byte) (n int, err error) {
	ORIGINAL := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ROT13 := "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
	n, err = rr.R.Read(b)
	for i := 0; i < len(b); i++ {
		index := strings.Index(ORIGINAL, string(b[i]))
		if index > -1 {
			b[i] = ROT13[index]
		}
	}
	return
}

// DisplayRot13 display decoded string
func DisplayRot13() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{R: s}
	io.Copy(os.Stdout, &r)
}
