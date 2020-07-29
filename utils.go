package glav

import "bytes"

func writeUintToBuffer(w *bytes.Buffer, v uint64) {
	var buf []byte
	for {
		c := uint8(v & 0x7f)
		v >>= 7
		if v != 0 {
			c |= 0x80
		}
		buf = append(buf, c)
		if c&0x80 == 0 {
			break
		}
	}
	w.Write(buf)
}

func writeStringToBuffer(buffer *bytes.Buffer, val string) {
	writeUintToBuffer(buffer, uint64(len(val)))
	buffer.Write([]byte(val))
}
