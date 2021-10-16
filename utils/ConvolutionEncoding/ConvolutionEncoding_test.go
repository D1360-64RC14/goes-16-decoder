package convolutionencoding

import (
	"bytes"
	"testing"
)

func TestNewConvolutionEncoding_test(t *testing.T) {
	testConv, err := NewConvolutionEncoding(3, testPolys)
	if err != nil {
		t.Fatal(err)
	}

	for index, data := range testResults {
		result := testConv.SendBit(data.Bit)

		if buffer_p := testConv.GetBuffer(); buffer_p != data.Buffer {
			t.Logf("Loop %d | Buffer: produzido (%b); esperado (%b).\n", index, buffer_p, data.Buffer)
			t.Fail()
		}
		if !bytes.Equal(result, data.Result) {
			t.Logf("Loop %d | Resultado: produzido %v; esperado %v.\n", index, result, data.Result)
			t.Fail()
		}
	}
}

func TestNewConvolutionEncoding_goes(t *testing.T) {
	goesConv, err := NewConvolutionEncoding(7, goesPolys)
	if err != nil {
		t.Fatal(err)
	}

	for index, data := range goesResults {
		result := goesConv.SendBit(data.Bit)

		if buffer_p := goesConv.GetBuffer(); buffer_p != data.Buffer {
			t.Logf("Loop %d | Buffer: produzido (%b); esperado (%b).\n", index, buffer_p, data.Buffer)
			t.Fail()
		}
		if !bytes.Equal(result, data.Result) {
			t.Logf("Loop %d | Resultado: produzido %v; esperado %v.\n", index, result, data.Result)
			t.Fail()
		}
	}
}

func TestNewConvolutionEncoding_goesErr(t *testing.T) {
	goesConvErr, err := NewConvolutionEncoding(7, goesPolys)
	if err != nil {
		t.Fatal(err)
	}

	for index, data := range goesResultsErr {
		result := goesConvErr.SendBit(data.Bit)

		if buffer_p := goesConvErr.GetBuffer(); buffer_p != data.Buffer && !data.ShoudFail {
			t.Logf("Loop %d | Buffer: produzido (%b); esperado (%b).\n", index, buffer_p, data.Buffer)
			t.Fail()
		}
		if !bytes.Equal(result, data.Result) && !data.ShoudFail {
			t.Logf("Loop %d | Resultado: produzido %v; esperado %v.\n", index, result, data.Result)
			t.Fail()
		}
	}
}

func BenchmarkNewConvolutionEncoding_test(b *testing.B) {
	testConv, err := NewConvolutionEncoding(3, testPolys)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		testConv.SendBit(byte(i))
	}
}

func BenchmarkNewConvolutionEncoding_goes(b *testing.B) {
	goesConv, err := NewConvolutionEncoding(3, goesPolys)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		goesConv.SendBit(byte(i))
	}
}