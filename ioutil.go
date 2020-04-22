package gox

import "io"

// WriteAll writes all data into w
func WriteAll(w io.Writer, data []byte) error {
	n, err := w.Write(data)
	for n < len(data) && err == nil {
		data = data[n:]
		n, err = w.Write(data)
	}
	return err
}
