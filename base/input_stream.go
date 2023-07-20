package base

// InputStream is a helper type for managing input streams from inside
// the Data event.
type InputStream struct {
	b []byte
}

// Begin accepts a new packet and returns a working sequence of
// unprocessed bytes.
func (is *InputStream) Begin(packet []byte) (data []byte) {
	data = packet
	if len(data) > 0 {
		is.b = append(is.b, data...)
		//data = is.b
	}
	return is.b
}

// End shifts the stream to match the unprocessed data.
func (is *InputStream) End(data []byte) {
	if len(data) > 0 {
		//if len(data) != len(is.b) {
		is.b = append(is.b[:0], data...)
		//}
	} else if len(is.b) > 0 {
		is.b = is.b[:0]
	}
}
