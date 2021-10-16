package convolutionencoding

func ConvolutionStep(buffer *byte, nextBit byte, polinomials []byte, k int) []byte {
	nextBit &= 1

	*buffer = *buffer >> 1
	*buffer = *buffer | (nextBit << (k - 1))

	limitPolinomialsByK(&polinomials, k)
	return andSumPolynomials(buffer, &polinomials)
}

func limitPolinomialsByK(polynomials *[]byte, limit int) {
	for index := range *polynomials {
		(*polynomials)[index] &= byteFill(limit)
	}
}

// can be made by LUT
func byteFill(qnt int) (result byte) {
	for i := 0; i < qnt; i++ {
		result |= 1 << i
	}
	return
}

func andSumPolynomials(buffer *byte, polinomials *[]byte) (result []byte) {
	result = make([]byte, len(*polinomials))

	for index, poly := range *polinomials {
		result[index] = poly & *buffer
		result[index] = xorByte(result[index])
	}
	return
}

// can be made by LUT
func xorByte(data byte) (result byte) {
	for i := 0; i < 8; i++ {
		result ^= 1 & (data >> i)
	}
	return
}