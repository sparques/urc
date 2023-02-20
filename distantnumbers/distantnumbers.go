package distantnumbers

// distantnumbers marshals/unmarshals a slice of uint16s

// Marshal converts the uint16 slice data into a byte slice
// Normally a marshal function allocates its return value, but since
// this is meant for embedded systems (and ala tinygo) we have the
// return buffer passed in, pre-allocated
func Marshal(data []uint16, dst []byte) {
	if len(dst) < len(data)*2 {
		panic("dst buffer smaller than input data")
	}
	for i := 0; i < len(data); i += 2 {
		dst[i], dst[i+1] = byte(data[i]>>1&0xFF), byte(data[i]&0xFF)
	}
}

func Unmarshal(dst []uint16, buf []byte) {
	max := len(dst)
	if len(buf)/2 < max {
		max = len(buf) / 2
	}

	for i := 0; i < len(dst); i++ {
		dst[i] = uint16(buf[i*2])<<1 | uint16(buf[i*2+1])
	}
}
