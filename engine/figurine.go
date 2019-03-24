package engine

import "strconv"


const(
	Empty string = "E"
	Pawn string = "P"
	Knight string = "N"
	Bishop string = "B"
	Rook string  ="R"
	Queen string = "Q"
	King string = "K"
)

type Figurine struct {
	//IsEmpty bool
	FigureWhite bool
	FigureType string
}

func (f Figurine) ToString()  string {
	//if (f.FigureType == nil || f.FigureSide == nil){
	//	return "empty row"
	//}
	return strconv.FormatBool(f.FigureWhite) + f.FigureType
}
/*
func (f Figurine) SpawnEmpty() Figurine {
	f.IsEmpty = true
	f.FigureWhite = true
	f.FigureType = Empty
	return f
}
*/
