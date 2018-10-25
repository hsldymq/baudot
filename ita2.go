package baudot

var letters = map[byte]rune{
	1:  'E',
	2:  '\n',
	3:  'A',
	4:  ' ',
	5:  'S',
	6:  'I',
	7:  'U',
	8:  '\r',
	9:  'D',
	10: 'R',
	11: 'J',
	12: 'N',
	13: 'F',
	14: 'C',
	15: 'K',
	16: 'T',
	17: 'Z',
	18: 'L',
	19: 'W',
	20: 'H',
	21: 'Y',
	22: 'P',
	23: 'Q',
	24: 'O',
	25: 'B',
	26: 'G',
	28: 'M',
	29: 'X',
	30: 'V',
}

var figures = map[byte]rune{
	1:  '3',
	2:  '\n',
	3:  '-',
	4:  ' ',
	5:  '\'',
	6:  '8',
	7:  '7',
	8:  '\r',
	9:  '\u0005',
	10: '4',
	11: '\u0007',
	12: ',',
	13: '!',
	14: ':',
	15: '(',
	16: '5',
	17: '+',
	18: ')',
	19: '2',
	20: '£',
	21: '6',
	22: '0',
	23: '1',
	24: '9',
	25: '?',
	26: '&',
	28: '.',
	29: '/',
	30: ';',
}

// the character set switch between letters and figures
var charset *map[byte]rune

var charmap = map[rune][2]int8{
	'E':      {1, -1},
	'\n':     {2, 2},
	'A':      {3, -1},
	' ':      {4, 4},
	'S':      {5, -1},
	'I':      {6, -1},
	'U':      {7, -1},
	'\r':     {8, 8},
	'D':      {9, -1},
	'R':      {10, -1},
	'J':      {11, -1},
	'N':      {12, -1},
	'F':      {13, -1},
	'C':      {14, -1},
	'K':      {15, -1},
	'T':      {16, -1},
	'Z':      {17, -1},
	'L':      {18, -1},
	'W':      {19, -1},
	'H':      {20, -1},
	'Y':      {21, -1},
	'P':      {22, -1},
	'Q':      {23, -1},
	'O':      {24, -1},
	'B':      {25, -1},
	'G':      {26, -1},
	'M':      {28, -1},
	'X':      {29, -1},
	'V':      {30, -1},
	'3':      {-1, 1},
	'-':      {-1, 3},
	'\'':     {-1, 5},
	'8':      {-1, 6},
	'7':      {-1, 7},
	'\u0005': {-1, 9},
	'4':      {-1, 10},
	'\u0007': {-1, 11},
	',':      {-1, 12},
	'!':      {-1, 13},
	':':      {-1, 14},
	'(':      {-1, 15},
	'5':      {-1, 16},
	'+':      {-1, 17},
	')':      {-1, 18},
	'2':      {-1, 19},
	'£':      {-1, 20},
	'6':      {-1, 21},
	'0':      {-1, 22},
	'1':      {-1, 23},
	'9':      {-1, 24},
	'?':      {-1, 25},
	'&':      {-1, 26},
	'.':      {-1, 28},
	'/':      {-1, 29},
	';':      {-1, 30},
}

// Encode string into byte array represent the sequence of Baudot-Murray code(ITA2)
// The sequence always starts with a null Control followed by a LS(Shift to Letters) Control
func Encode(msg string) ([]byte, error) {
	return []byte{}, nil
}

// Decode Baudot-Murray code(ITA2) to string
func Decode(codes []byte) (string, error) {
	return "", nil
}

// EncodeChar encodes a character into Baudot-Murray code(ITA2)
func EncodeChar(c rune) (byte, error) {
	return 1, nil
}

// DecodeChar decodes a Baudot-Murray code(ITA2) to rune
func DecodeChar(byte) (rune, error) {
	return '1', nil
}
