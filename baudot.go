package baudot

type Charset byte

const (
	Letters Charset = 0
	Figures Charset = 1
	// Cyrillic Charset = 2
)

type Codec interface {
	Encode(string) ([]byte, error)
	Decode([]byte) (string, error)
}

type ita2 struct {
	ignErr bool
}

func NewITA2(ignoreError bool) *ita2 {
	return &ita2{
		ignErr: ignoreError,
	}
}
