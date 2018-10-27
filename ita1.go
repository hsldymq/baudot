/*
 * ITA1(International Telegraph Alphabet No.1) is Original Baudot Code.
 * This is The UK Version.
 */

package baudot

// Encode string into byte array represent the sequence of Baudot code
func (c *ita1) Encode(msg string) ([]byte, error) {
	return encode(msg, c.ignErr, versionITA1)
}

// Decode Baudot code to string
func (c *ita1) Decode(codes []byte) (string, error) {
	return decode(codes, c.ignErr, versionITA1)
}

// EncodeChar encodes a character into Baudot code
func (c *ita1) EncodeChar(char rune, currentCharset Charset) (byte, bool, error) {
	code, shiftedCharset, err := encodeChar(char, currentCharset, versionITA1)

	return code, shiftedCharset != currentCharset, err
}

// DecodeChar decodes a Baudot code to rune
func (c *ita1) DecodeChar(code byte, currentCharset Charset) (rune, bool, error) {
	char, shiftedCharset, err := decodeChar(code, currentCharset, versionITA1)

	return char, currentCharset != shiftedCharset, err
}
