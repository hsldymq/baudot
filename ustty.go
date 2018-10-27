/*
 * US as known as Baubot-Murray code which is a modication of Baudot code(ITA1)
 */

package baudot

// Encode string into byte array represent the sequence of US TTY(variant of ITA2)
func (c *ustty) Encode(msg string) ([]byte, error) {
	return encode(msg, c.ignErr, versionUSTTY)
}

// Decode Baudot-Murray code(ITA2) to string
func (c *ustty) Decode(codes []byte) (string, error) {
	return decode(codes, c.ignErr, versionUSTTY)
}

// EncodeChar encodes a character into Baudot-Murray code(ITA2)
func (c *ustty) EncodeChar(char rune, currentCharset Charset) (byte, bool, error) {
	code, shiftedCharset, err := encodeChar(char, currentCharset, versionUSTTY)

	return code, shiftedCharset != currentCharset, err
}

// DecodeChar decodes a Baudot-Murray code(ITA2) to rune
func (c *ustty) DecodeChar(code byte, currentCharset Charset) (rune, bool, error) {
	char, shiftedCharset, err := decodeChar(code, currentCharset, versionUSTTY)

	return char, currentCharset != shiftedCharset, err
}
