package convolutionencoding

import (
	"bytes"
	"testing"
)

func TestConvolutionStep_test(t *testing.T) {
	buffer := byte(0b000)

	for index, data := range testResults {
		result := ConvolutionStep(&buffer, data.Bit, testPolys, 3)

		if data.Buffer != buffer {
			t.Logf("Loop %d | Buffer: produzido (%b); esperado (%b).\n", index, buffer, data.Buffer)
			t.Fail()
		}
		if !bytes.Equal(data.Result, result) {
			t.Logf("Loop %d | Resultado: produzido (%b); esperado (%b).\n", index, result, data.Result)
			t.Fail()
		}
	}
}

func TestConvolutionStep_goes(t *testing.T) {
	buffer := byte(0b000)

	for index, data := range goesResults {
		result := ConvolutionStep(&buffer, data.Bit, goesPolys, 7)

		if data.Buffer != buffer {
			t.Logf("Loop %d | Buffer: produzido (%b); esperado (%b).\n", index, buffer, data.Buffer)
			t.Fail()
		}
		if !bytes.Equal(data.Result, result) {
			t.Logf("Loop %d | Resultado: produzido (%b); esperado (%b).\n", index, result, data.Result)
			t.Fail()
		}
	}
}

func TestConvolutionStep_goesErr(t *testing.T) {
	buffer := byte(0b000)

	for index, data := range goesResultsErr {
		result := ConvolutionStep(&buffer, data.Bit, goesPolys, 7)

		if data.Buffer != buffer && !data.ShoudFail {
			t.Logf("Loop %d | Buffer: produzido (%b); esperado (%b).\n", index, buffer, data.Buffer)
			t.Fail()
		}
		if data.ShoudFail {
			buffer = data.Buffer
		}
		if !bytes.Equal(data.Result, result) && !data.ShoudFail {
			t.Logf("Loop %d | Resultado: produzido (%b); esperado (%b).\n", index, result, data.Result)
			t.Fail()
		}
	}
}