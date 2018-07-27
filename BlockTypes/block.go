package BlockTypes

type BlockProperties struct {
	Width       string
	Height      string
	Row         int
	Col         int
	Border      int
	BorderColor string
	Children    []BlockProperties
	Parent      *BlockProperties
}

type BaseBlock interface {
}
