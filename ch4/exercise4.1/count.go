package count

// Different returns the number of bits that are different in two SHA256 hashes.
func Different(h1, h2 [32]byte) int {
	acc := 0
	for i, b := range h1 {
		acc += compare(b, h2[i])
	}
	return acc
}

func differentSlice(h1, h2 []byte) int {
	acc := 0
	for i, b := range h1 {
		acc += compare(b, h2[i])
	}
	return acc
}

func compare(b1, b2 byte) int {
	b := b1 ^ b2
	count := 0
	for b != 0 {
		b &= b - 1
		count++
	}
	return count
}
