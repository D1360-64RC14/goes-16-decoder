package convolutionencoding

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

//  Valores
//     índice 0   = bit a ser inserido
//     índice 1   = buffer esperado
//     índice 2.. = valores esperados de retorno
var testResults = []Result{
	{ 1, 0b100, []byte{ 1, 0, 1 }, false },
	{ 1, 0b110, []byte{ 0, 1, 1 }, false },
	{ 0, 0b011, []byte{ 0, 0, 1 }, false },
	{ 1, 0b101, []byte{ 0, 1, 0 }, false },
	{ 0, 0b010, []byte{ 1, 1, 0 }, false },
	{ 0, 0b001, []byte{ 1, 1, 1 }, false },
	{ 1, 0b100, []byte{ 1, 0, 1 }, false },
	{ 1, 0b110, []byte{ 0, 1, 1 }, false },
	{ 0, 0b011, []byte{ 0, 0, 1 }, false },
	{ 1, 0b101, []byte{ 0, 1, 0 }, false },
}
// Polinômios de teste
var testPolys = []byte{ 0b111, 0b011, 0b101 }

// goes 11, 12, 13 e 14
var goesResults = []Result{
	{ 1, 0b1000000, []byte{ 1, 1 }, false },
	{ 1, 0b1100000, []byte{ 0, 1 }, false },
	{ 0, 0b0110000, []byte{ 0, 1 }, false },
	{ 1, 0b1011000, []byte{ 1, 1 }, false },
	{ 0, 0b0101100, []byte{ 0, 1 }, false },
	{ 0, 0b0010110, []byte{ 1, 0 }, false },
	{ 1, 0b1001011, []byte{ 1, 0 }, false },
	{ 1, 0b1100101, []byte{ 1, 0 }, false },
	{ 0, 0b0110010, []byte{ 0, 0 }, false },
	{ 1, 0b1011001, []byte{ 0, 0 }, false },
}
var goesPolys = []byte{ 0b1111001, 0b1011011 }

var goesResultsErr = []Result{
	{ 1, 0b1000000, []byte{ 1, 1 }, false },
	{ 1, 0b1100000, []byte{ 0, 1 }, false },
	{ 0, 0b0110000, []byte{ 0, 1 }, false },
	{ 1, 0b1011000, []byte{ 1, 1 }, false },
	{ 0, 0b0101100, []byte{ 0, 0 }, true  }, // 4
	{ 0, 0b0010110, []byte{ 1, 0 }, false },
	{ 1, 0b1001011, []byte{ 1, 0 }, false },
	{ 1, 0b1100101, []byte{ 1, 0 }, false },
	{ 0, 0b0111010, []byte{ 0, 0 }, true  }, // 8
	{ 1, 0b1011001, []byte{ 0, 0 }, false },
}

type Result struct {
	Bit         byte
	Buffer      byte
	Result    []byte
	ShoudFail   bool
}

func TestNewConvolutionEncoding8(t *testing.T) {
	testConv, err := NewConvolutionEncoding8(3, testPolys)
	if err != nil {
		t.Fatal(err)
	}
	goesConv, err := NewConvolutionEncoding8(7, goesPolys)
	if err != nil {
		t.Fatal(err)
	}
	goesConvErr, err := NewConvolutionEncoding8(7, goesPolys)
	if err != nil {
		t.Fatal(err)
	}

	if testConv.GetBuffer() != 0b0 || goesConv.GetBuffer() != 0b0 || goesConvErr.GetBuffer() != 0b0 {
		t.Logf("Buffers iniciaram em um estado não zero.\ntestConv: %b\bgoesConv: %b\n",
			testConv.GetBuffer(),
			goesConv.GetBuffer(),
		)
		t.Fail()
	}

	// [ 1, 1, 0, 1, 0, 0, 1, 1, 0, 1 ]

	looper("testConv", testConv, testResults, t)
	looper("goesConv", goesConv, goesResults, t)
	looper("goesConvErr", goesConvErr, goesResultsErr, t) // Isso deve tar erro nos índices 4 e 8
}

// looper executa testes com as informações dadas.
func looper(name string, convolutioner *ConvolutionEncoding, results []Result, t *testing.T) {
	t.Log("Executando testes de " + name)

	for loop_index, result := range results {
		result_got := convolutioner.SendBit(result.Bit)
		buffer_got := convolutioner.GetBuffer()

		if buffer_got != result.Buffer {
			if result.ShoudFail {
				t.Logf(
					"\n-- Falha esperada em Loop %d --\n" +
					"  Buffer esperado: %b\n"+
					"  Buffer recebido: %b\n",
					loop_index,
					result.Buffer,
					buffer_got,
				)
			} else {
				t.Logf(
					"\n-- Falha não esperada em Loop %d --\n" +
					"  Buffer esperado: %b\n"+
					"  Buffer recebido: %b\n",
					loop_index,
					result.Buffer,
					buffer_got,
				)
				t.Fail()
			}
		}
		if !bytes.Equal(result_got, result.Result) {
			exp_result := make([]string, len(result_got))
			got_result := make([]string, len(result_got))

			for i := 0; i < len(result_got); i++ {
				exp_result[i] = fmt.Sprintf("%b", result.Result[i])
				got_result[i] = fmt.Sprintf("%b", result_got[i])
			}
			
			if result.ShoudFail {
				t.Logf(
					"\n-- Falha esperada em Loop %d --\n" +
					"  Resultado esperado: %s\n"+
					"  Resultado recebido: %s\n",
					loop_index,
					strings.Join(exp_result, ", "),
					strings.Join(got_result, ", "),
				)
			} else {
				t.Logf(
					"\n-- Falha não esperada em Loop %d --\n" +
					"  Resultado esperado: %s\n"+
					"  Resultado recebido: %s\n",
					loop_index,
					strings.Join(exp_result, ", "),
					strings.Join(got_result, ", "),
				)
				t.Fail()
			}
		}
	}
}
