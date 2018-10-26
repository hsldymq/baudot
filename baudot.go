package baudot

type Charset byte

const (
	Letters  Charset = 0
	Figures  Charset = 1
	Cyrillic Charset = 2
)

type Codec interface {
	Encode(string) ([]byte, error)
	Decode([]byte) (string, error)
}

func NewITA1() {

}

func NewITA2() {

}

func NewUSTTY() {

}

func NewMTK2() {

}
