package convolutionencoding

import (
	"bytes"
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
	for i := 0; i < len(polinomials); i++ {
		polinomials[i] = bitFill(k) & polinomials[i]
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

	andSummed := bytes.Map(func(poly rune) rune {
		and := byte(poly) & ce.buffer
		sum := sumByte(and)
		return rune(sum)
	}, ce.polinomials)
	
	return andSummed
}

// bitFill cria um byte preenchido com uma quantidade de bits 1.
// Por exemplo:
//
// bitFill(2) = 0b11
//
// bitFill(4) = 0b1111
//
// bitFill(5) = 0b11111
func bitFill(length uint8) (result byte) {
	result = 0

	for i := uint8(0); i < length; i++ {
		result = result | 1 << i
	}

	return
}

// readBitPosition lê um bit em uma posição específica do byte.
// Por exemplo:
//
// readBitPosition(0b101, 1) = 0 // 0b1[0]1
//
// readBitPosition(0b10011101, 3) = 1 // 0b1001[1]101
func readBitPosition(data byte, position uint8) byte {
	return 0b1 & (data >> position)
}

// sumByte efetua uma soma entre todos os bits de um byte.
// Por exemplo:
//
// sumByte(0b101001) = 1 + 0 + 1 + 0 + 0 + 1 = 1
func sumByte(data byte) (result byte) {
	for i := uint8(0); i < 8; i++ {
		result = result + readBitPosition(data, i)
	}
	return 0b1 & result
}