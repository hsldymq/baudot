/*
 * ITA2(International Telegraph Alphabet No.2) as known as Baubot-Murray code which is a modication of Baudot code(ITA1)
 */

package baudot

// Encode string into byte array represent the sequence of Baudot-Murray code(ITA2)
func (c *ita1) Encode(msg string) ([]byte, error) {
	return encode(msg, c.ignErr, versionITA1)
}

// Decode Baudot-Murray code(ITA2) to string
func (c *ita1) Decode(codes []byte) (string, error) {
	return decode(codes, c.ignErr, versionITA1)
}

// EncodeChar encodes a character into Baudot-Murray code(ITA2)
func (c *ita1) EncodeChar(char rune, currentCharset Charset) (byte, bool, error) {
	code, shiftedCharset, err := encodeChar(char, currentCharset, versionITA1)

	return code, shiftedCharset != currentCharset, err
}

// DecodeChar decodes a Baudot-Murray code(ITA2) to rune
func (c *ita1) DecodeChar(code byte, currentCharset Charset) (rune, bool, error) {
	char, shiftedCharset, err := decodeChar(code, currentCharset, versionITA1)

	return char, currentCharset != shiftedCharset, err
}
