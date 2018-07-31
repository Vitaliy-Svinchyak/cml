package Config

var maxX, maxY int

func SetMaxXAndMaxY(x, y int) {
	maxX = x
	maxY = y
}

func GetTerminalSize() (int, int) {
	return maxX, maxY
}
