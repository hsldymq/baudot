package baudot

import "fmt"

type Charset byte

type version byte

const (
	Letters Charset = 0
	Figures Charset = 1
	// Cyrillic Charset = 2
)

const (
	NULL byte = 0
	// ITA2 Shift to Figures
	FS byte = 27
	// ITA2 Shift to Letters
	LS byte = 31
	// ITA1 Shift to Letters
	LS_ITA1 byte = 1
	// ITA1 Shift to Figures
	FS_ITA1 byte = 2
)

const (
	versionITA1  version = 0
	versionITA2  version = 1
	versionUSTTY version = 2
)

type Codec interface {
	Encode(string) ([]byte, error)
	Decode([]byte) (string, error)
}

type ita1 struct {
	ignErr bool
}

type ita2 struct {
	ignErr bool
}

type ustty struct {
	ignErr bool
}

func NewITA1(ignoreError bool) *ita1 {
	return &ita1{
		ignErr: ignoreError,
	}
}

func NewITA2(ignoreError bool) *ita2 {
	return &ita2{
		ignErr: ignoreError,
	}
}

func NewUSTTY(ignoreError bool) *ustty {
	return &ustty{
		ignErr: ignoreError,
	}
}

// The sequence always starts with a null Control followed by a LS(Shift to Letters) Control
func encode(msg string, ignoreError bool, ver version) ([]byte, error) {
	var (
		shifters       [2]byte
		shiftersITA2   = [2]byte{LS, FS}
		shiftersITA1   = [2]byte{LS_ITA1, FS_ITA1}
		currentCharset = Letters
		codes          = []byte{NULL, LS}
	)

	for _, char := range msg {
		if ver == versionITA1 {
			shifters = shiftersITA1
		} else if ver == versionITA2 || ver == versionUSTTY {
			shifters = shiftersITA2
		} else if ignoreError {
			break
		} else {
			return nil, fmt.Errorf("")
		}

		code, shiftedCharset, err := encodeChar(char, currentCharset, ver)

		if err != nil {
			if ignoreError {
				continue
			} else {
				return nil, err
			}
		}

		if currentCharset != shiftedCharset {
			currentCharset = shiftedCharset
			codes = append(codes, shifters[currentCharset])
		}

		codes = append(codes, code)
	}

	return codes, nil
}

func decode(codes []byte, ignoreError bool, ver version) (string, error) {
	var str []rune
	currentCharset := Letters

	for _, eachCode := range codes {
		ch, shiftedCharset, err := decodeChar(eachCode, currentCharset, ver)

		if err != nil {
			if ignoreError {
				continue
			} else {
				return "", err
			}
		}

		if currentCharset != shiftedCharset {
			currentCharset = shiftedCharset
			continue
		}

		if ch == '\u0000' {
			continue
		}

		str = append(str, ch)
	}

	return string(str), nil
}

func encodeChar(char rune, currentCharset Charset, ver version) (byte, Charset, error) {
	var (
		shiftedCharset = currentCharset
		charValues     [2]int8
		ok             bool
	)

	if ver == versionITA1 {
		charValues, ok = charmapITA1[char]
	} else if ver == versionITA2 {
		charValues, ok = charmapITA2[char]
	} else if ver == versionUSTTY {
		charValues, ok = charmapUSTTY[char]
	} else {
		return '\u0000', currentCharset, fmt.Errorf("Unsupported version: %d", ver)
	}

	if !ok {
		// always return error, not affect by ignErr field
		return 0, currentCharset, fmt.Errorf("Invalid Char: %c", char)
	}

	code := charValues[currentCharset]
	if code == -1 {
		shiftedCharset = Charset(currentCharset ^ 1)
		code = charValues[shiftedCharset]
	}

	return byte(code), shiftedCharset, nil
}

func decodeChar(code byte, currentCharset Charset, ver version) (rune, Charset, error) {
	var charset map[byte]rune

	if ver == versionITA1 {
		if code == LS_ITA1 {
			return '\u0000', Letters, nil
		} else if code == FS_ITA1 {
			return '\u0000', Figures, nil
		}

		if currentCharset == Letters {
			charset = lettersITA1
		} else {
			charset = figuresITA1
		}
	} else if ver == versionITA2 || ver == versionUSTTY {
		if code == LS {
			return '\u0000', Letters, nil
		} else if code == FS {
			return '\u0000', Figures, nil
		}

		if currentCharset == Letters {
			charset = lettersITA2
		} else if ver == versionITA2 {
			charset = figuresITA2
		} else {
			charset = figuresUSTTY
		}
	} else {
		return '\u0000', currentCharset, fmt.Errorf("Unsupported version: %d", ver)
	}

	char, ok := charset[code]
	if !ok {
		// always return error, not affect by ignErr field
		return '\u0000', currentCharset, fmt.Errorf("Invalid Code: %d", code)
	}

	return char, currentCharset, nil
}
