package helper


func XorShift(r uint64) uint64 {
	r ^= r << 13
	r ^= r >> 7
	r ^= r << 17
	return r
}
