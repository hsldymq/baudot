/*
 * ITA2(International Telegraph Alphabet No.2) as known as Baubot-Murray code which is a modication of Baudot code(ITA1)
 */

package baudot

import (
	"fmt"
)

var charmapITA2 = map[rune][2]int8{
	'\u0000': {0, 0},
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
func (c *ita2) Encode(msg string) ([]byte, error) {
	currentCharset := Letters
	shifters := [2]byte{LS, FS}

	codes := []byte{NULL, LS}
	for _, char := range msg {
		code, shifted, err := c.EncodeChar(char, currentCharset)

		if err != nil {
			if c.ignErr {
				continue
			} else {
				return nil, err
			}
		}

		if shifted {
			currentCharset ^= 1
			codes = append(codes, shifters[currentCharset])
		}

		codes = append(codes, code)
	}

	return codes, nil
}

// Decode Baudot-Murray code(ITA2) to string
func (c *ita2) Decode(codes []byte) (string, error) {
	var str []rune
	currentCharset := Letters

	for _, eachCode := range codes {
		ch, shifted, err := c.DecodeChar(eachCode, currentCharset)

		if err != nil {
			if c.ignErr {
				continue
			} else {
				return "", err
			}
		}

		if shifted {
			currentCharset ^= 1
			continue
		}

		if ch == '\u0000' {
			continue
		}

		str = append(str, ch)
	}

	return string(str), nil
}

// EncodeChar encodes a character into Baudot-Murray code(ITA2)
func (c *ita2) EncodeChar(char rune, currentCharset Charset) (byte, bool, error) {
	shifted := false

	charValues, ok := charmapITA2[char]
	if !ok {
		// always return error, not affect by ignErr field
		return 0, false, fmt.Errorf("Invalid Char: %c", char)
	}

	code := charValues[currentCharset]
	if code == -1 {
		shifted = true
		code = charValues[currentCharset^1]
	}

	return byte(code), shifted, nil
}

// DecodeChar decodes a Baudot-Murray code(ITA2) to rune
func (c *ita2) DecodeChar(code byte, currentCharset Charset) (rune, bool, error) {
	var charset map[byte]rune

	if code == LS {
		return '\u0000', currentCharset != Letters, nil
	} else if code == FS {
		return '\u0000', currentCharset != Figures, nil
	}

	if currentCharset == Letters {
		charset = lettersITA2
	} else {
		charset = figuresITA2
	}

	char, ok := charset[code]
	if !ok {
		// always return error, not affect by ignErr field
		return '\u0000', false, fmt.Errorf("Invalid Code: %d", code)
	}

	return char, false, nil
}
