package convolutionencoding

// ConvolutionStep funciona com: dado um buffer, próximo bit, polinômios e quantidade de bits (k),
// executa as operações e retorna os devidos valores.
func ConvolutionStep(buffer *byte, nextBit byte, polinomials []byte, k int) []byte {
	nextBit &= 1

	*buffer = *buffer >> 1
	*buffer = *buffer | (nextBit << (k - 1))

	limitPolinomialsByK(polinomials, k)
	return andSumPolynomials(*buffer, polinomials)
}

// limitPolinomialsByK limita o tamanho dos polinômio de acordo com a quantidade de bits (k).
func limitPolinomialsByK(polynomials []byte, limit int) {
	for index := range polynomials {
		polynomials[index] &= byteFill(limit)
	}
}

// byteFill preenche um byte com bits 1 de acordo com a quantidade informada.
// Feito utilizando LUT.
func byteFill(qnt int) byte {
	// for i := 0; i < qnt; i++ {
	// 	result |= 1 << i
	// }
	// return
	LUT := [9]byte{
		0b00000000,
		0b00000001,
		0b00000011,
		0b00000111,
		0b00001111,
		0b00011111,
		0b00111111,
		0b01111111,
		0b11111111,
	}
	return LUT[qnt]
}

// andSumPolynomials executa as operações retornando o resultado final.
func andSumPolynomials(buffer byte, polinomials []byte) (result []byte) {
	result = make([]byte, len(polinomials))

	for index, poly := range polinomials {
		result[index] = poly & buffer
		result[index] = xorByte(result[index])
	}
	return
}

// xorByte soma todos os bits de um byte.
// Feito utilizando LUT.
func xorByte(data byte) byte {
	// for i := 0; i < 8; i++ {
	// 	result ^= 1 & (data >> i)
	// }
	// return
	LUT := [255]byte{
		0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1,
		0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0,
		0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1,
		0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1,
		0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0,
		1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0,
		0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0,
		1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0,
		0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1,
		0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1,
		1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0,
		1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1,
		0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1,
	}
	return LUT[data]
}