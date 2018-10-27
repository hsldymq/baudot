package baudot

import (
	"fmt"
	"testing"
)

func TestEncodeChar(t *testing.T) {
	tt := []struct {
		caseName    string
		char        rune
		charset     Charset
		expectCode  byte
		expectShift bool
		shouldFail  bool
		failedText  string
	}{
		{
			caseName:    "test regular char",
			char:        'R',
			charset:     Letters,
			expectCode:  10,
			expectShift: false,
			shouldFail:  false,
			failedText:  "code for 'R' should be 10, got %v",
		},
		{
			caseName:    "test invalid char",
			char:        '^',
			charset:     Letters,
			expectCode:  0,
			expectShift: false,
			shouldFail:  true,
			failedText:  "encode code for char '^' should return error, got %v",
		},
		{
			caseName:    "test shift to letters charset",
			char:        'A',
			charset:     Figures,
			expectCode:  3,
			expectShift: true,
			shouldFail:  false,
			failedText:  "value of shifted Should Be true, got %v",
		},
		{
			caseName:    "test shift to figures charset",
			char:        '6',
			charset:     Letters,
			expectCode:  21,
			expectShift: true,
			shouldFail:  false,
			failedText:  "value of shifted Should Be true, got %v",
		},
	}

	c := NewITA2(false)
	for _, tc := range tt {
		t.Run(tc.caseName, func(t *testing.T) {
			code, shifted, err := c.EncodeChar(tc.char, tc.charset)
			if err != nil {
				if !tc.shouldFail {
					t.Errorf(tc.failedText, fmt.Sprintf("%v, %v, %v", code, shifted, err))
				}
			} else {
				if tc.expectCode != code || tc.expectShift != shifted {
					t.Errorf(tc.failedText, fmt.Sprintf("%v, %v, %v", code, shifted, err))
				}
			}
		})
	}
}

func TestDecodeChar(t *testing.T) {
	tt := []struct {
		caseName    string
		code        byte
		charset     Charset
		expectChar  rune
		expectShift bool
		shouldFail  bool
		failedText  string
	}{
		{
			caseName:    "test letter code",
			code:        10,
			charset:     Letters,
			expectChar:  'R',
			expectShift: false,
			shouldFail:  false,
			failedText:  "expect 'R', got %v",
		},
		{
			caseName:    "test figure code",
			code:        10,
			charset:     Figures,
			expectChar:  '4',
			expectShift: false,
			shouldFail:  false,
			failedText:  "expect '4', got %v",
		},
		{
			caseName:    "test invalid code",
			code:        50,
			charset:     Letters,
			expectChar:  '\u0000',
			expectShift: false,
			shouldFail:  true,
			failedText:  "expect an error, got %v",
		},
		{
			caseName:    "test shift to letters",
			code:        31,
			charset:     Figures,
			expectChar:  '\u0000',
			expectShift: true,
			shouldFail:  false,
			failedText:  "expect LS control, got %v",
		},
		{
			caseName:    "test shift to figures",
			code:        27,
			charset:     Letters,
			expectChar:  '\u0000',
			expectShift: true,
			shouldFail:  false,
			failedText:  "expect FS control, got %v",
		},
	}

	c := NewITA2(false)
	for _, tc := range tt {
		t.Run(tc.caseName, func(t *testing.T) {
			char, shifted, err := c.DecodeChar(tc.code, tc.charset)
			if err != nil {
				if !tc.shouldFail {
					t.Errorf(tc.failedText, fmt.Sprintf("%v, %v, %v", char, shifted, err))
				}
			} else {
				if tc.expectChar != char || tc.expectShift != shifted {
					t.Errorf(tc.failedText, fmt.Sprintf("%v, %v, %v", char, shifted, err))
				}
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tt := []struct {
		caseName   string
		msg        string
		ignErr     bool
		expect     []byte
		shouldFail bool
		failedText string
	}{
		{
			caseName:   "test letters and figures",
			msg:        "X&Y",
			ignErr:     false,
			expect:     []byte{0, 31, 29, 27, 26, 31, 21},
			shouldFail: false,
			failedText: fmt.Sprintf("expect %v, got %%v", []byte{0, 31, 29, 27, 26, 31, 21}),
		},
		{
			caseName:   "test invalid msg",
			msg:        "1 + 1 ~ 2",
			ignErr:     false,
			expect:     nil,
			shouldFail: true,
			failedText: "expect an error, got %v",
		},
		{
			caseName:   "test invalid msg, ignore error",
			msg:        "1 + 1 ~ 2",
			ignErr:     true,
			expect:     []byte{0, 31, 27, 23, 4, 17, 4, 23, 4, 4, 19},
			shouldFail: true,
			failedText: fmt.Sprintf("expect %v, got %%v", []byte{0, 31, 27, 23, 4, 17, 4, 23, 4, 4, 19}),
		},
	}

	c := NewITA2(false)
	for _, tc := range tt {
		t.Run(tc.caseName, func(t *testing.T) {
			c.ignErr = tc.ignErr
			codes, err := c.Encode(tc.msg)
			if err != nil {
				if !tc.shouldFail {
					t.Errorf(tc.failedText, fmt.Sprintf("%v, %v", codes, err))
				}
			} else {
				if len(tc.expect) != len(codes) {
					t.Errorf(tc.failedText, fmt.Sprintf("%v, %v", codes, err))
				} else {
					for index, value := range tc.expect {
						if value != codes[index] {
							t.Errorf(tc.failedText, fmt.Sprintf("%v, %v", codes, err))
						}
					}
				}
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tt := []struct {
		caseName   string
		codes      []byte
		ignErr     bool
		expect     string
		shouldFail bool
		failedText string
	}{
		{
			caseName:   "test decoding valid code",
			codes:      []byte{0, 31, 29, 27, 26, 31, 21},
			ignErr:     false,
			expect:     "X&Y",
			shouldFail: false,
			failedText: "expect 'X&Y', got %v",
		},
		{
			caseName:   "test decoding valid code contains null",
			codes:      []byte{0, 27, 1, 31, 22, 0, 0, 24, 27, 26, 0, 31, 10, 27, 19, 31, 9, 27, 19},
			ignErr:     false,
			expect:     "3PO&R2D2",
			shouldFail: false,
			failedText: "expect '3PO&R2D2', got %v",
		},
		{
			caseName:   "test valid code",
			codes:      []byte{0, 27, 1, 31, 100, 12},
			ignErr:     false,
			expect:     "",
			shouldFail: true,
			failedText: "expect an error, got %v",
		},
		{
			caseName:   "test valid code, ignore error",
			codes:      []byte{0, 27, 1, 31, 100, 12},
			ignErr:     true,
			expect:     "3N",
			shouldFail: true,
			failedText: "expect an error, got %v",
		},
	}

	c := NewITA2(false)
	for _, tc := range tt {
		t.Run(tc.caseName, func(t *testing.T) {
			c.ignErr = tc.ignErr
			str, err := c.Decode(tc.codes)
			if err != nil {
				if !tc.shouldFail {
					t.Errorf(tc.failedText, fmt.Sprintf("%v, %v", str, err))
				}
			} else {
				if tc.expect != str {
					t.Errorf(tc.failedText, fmt.Sprintf("%v, %v", str, err))
				}
			}
		})
	}
}
