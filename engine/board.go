package engine

type Board struct {
	new_area map[int]map[int]Field
}

type Field struct {
	is_filled bool
	isWhite   bool
	figurine  Figurine
}

func NewBoard() Board {
	p := new(Board)
	p.initNewBoard()
	return *p
}

func (b Board) initNewBoard(){
	b.new_area = make(map[int]map[int]Field)
	for x := 1; x <= 8; x++{
		if(b.new_area[x] == nil){
			b.new_area[x] = make(map[int]Field, 8)
		}
		var x_channel <-chan Figurine
		if(x == 1){
			x_channel = getInitialFigureRow(true)
			// Spawn White figures
		} else if (x == 8) {
			// Spawn Black figures
			x_channel = getInitialFigureRow(false)
		}else if(x == 2){
			// Spawn White pawn
			x_channel = getPawnRow(true)
		}else if(x == 7){
			// Spawn Black pawn
			x_channel = getPawnRow(false)
		}else{
			continue
		}
		for y := 1; y <= 8; y++{
			if (x_channel != nil){
				b.new_area[x][y] = Field{true, false, (<-x_channel)}
			}

		}
		//b.new_area[x]
	}
}

func getInitialFigureRow(isWhite bool) <-chan Figurine {
	out := make(chan Figurine, 8)
	out <- Figurine{isWhite, Knight}
	out <- Figurine{isWhite, Bishop}
	out <- Figurine{isWhite, Rook}
	out <- Figurine{isWhite, Queen}
	out <- Figurine{isWhite, King}
	out <- Figurine{isWhite, Rook}
	out <- Figurine{isWhite, Bishop}
	out <- Figurine{isWhite, Knight}

	return out
}

func getPawnRow(isWhite bool) <-chan Figurine{
	out := make(chan Figurine, 8)
	for x:=0; x > 8; x++{
		out <- Figurine{isWhite, Pawn}
	}

	return out
}

func (b Board)GetBoardAsString() string{
	boardString := ""
	for x:=1; x<=8;x++{
		for y:=1; y<=8;y++{
			//
		}
	}
	return boardString
}
