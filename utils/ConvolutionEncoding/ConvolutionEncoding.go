package convolutionencoding

import (
	"errors"
)

type ConvolutionEncoding struct {
	buffer        byte
	k             uint8
	polinomials []byte
}

// NewConvolutionEncoding8 cria uma instância de convolução de até 8 bits.
//
// - `k` diz respeito à quantidade de bits do buffer.
//
// - `polinomials` diz respeito aos polinômios geradores.
//
// A razão de codificação (r) é dada pela quantidade de polinômios.
func NewConvolutionEncoding(k uint8, polinomials []byte) (*ConvolutionEncoding, error) {
	if k > 8 {
		return &ConvolutionEncoding{}, errors.New("'k' deve ser menor ou igual a 8")
	}

	// Filtra todos os polinômios para terem
	// a quantidade correta de bits (k).
	for index := range polinomials {
		polinomials[index] = bitFill(k) & polinomials[index]
	}

	return &ConvolutionEncoding{
		buffer:      byte(0),
		k:           k,
		polinomials: polinomials,
	}, nil
}

func (ce *ConvolutionEncoding) SetBuffer(newBuffer byte) {
	newBuffer = bitFill(ce.k) & newBuffer
	ce.buffer = newBuffer
}
func (ce ConvolutionEncoding) GetBuffer() byte {
	return ce.buffer
}

func (ce *ConvolutionEncoding) SendBit(bit byte) []byte {
	ce.buffer = ce.buffer >> 1
	ce.buffer = ce.buffer | (0b1 & bit) << (ce.k - 1)

	andSum := make([]byte, len(ce.polinomials))
	for index, poly := range ce.polinomials {
		andSum[index] = sumByte(poly & ce.buffer)
	}
	
	return andSum
}

// bitFill cria um byte preenchido com uma quantidade de bits 1.
// Por exemplo:
//
// bitFill(2) = 0b11
//
// bitFill(4) = 0b1111
//
// bitFill(5) = 0b11111
func bitFill(length uint8) byte {
	// result = 0

	// for i := uint8(0); i < length; i++ {
	// 	result = result | 1 << i
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
	return LUT[length]
}

// readBitPosition lê um bit em uma posição específica do byte.
// Por exemplo:
//
// readBitPosition(0b101, 1) = 0 // 0b1[0]1
//
// readBitPosition(0b10011101, 3) = 1 // 0b1001[1]101
// func readBitPosition(data byte, position uint8) byte {
// 	return 0b1 & (data >> position)
// }

// sumByte efetua uma soma entre todos os bits de um byte.
// Por exemplo:
//
// sumByte(0b101001) = 1 + 0 + 1 + 0 + 0 + 1 = 1
func sumByte(data byte) byte {
	// for i := uint8(0); i < 8; i++ {
	// 	result = result + readBitPosition(data, i)
	// }
	// return 0b1 & result
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