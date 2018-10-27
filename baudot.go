package baudot

type Charset byte

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
)

type Codec interface {
	Encode(string) ([]byte, error)
	Decode([]byte) (string, error)
}

type ita2 struct {
	ignErr bool
}

type ustty struct {
	ignErr bool
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
