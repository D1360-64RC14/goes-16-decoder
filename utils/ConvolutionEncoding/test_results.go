package convolutionencoding

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
var testPolys = []byte{ 0b111, 0b011, 0b101 }

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
	{ 0, 0b0110010, []byte{ 0, 1 }, true  }, // 8
	{ 1, 0b1011001, []byte{ 0, 0 }, false },
}

type Result struct {
	Bit         byte
	Buffer      byte
	Result    []byte
	ShoudFail   bool
}
