package main

import "fmt"
import "./engine"

func main() {
	var bord = engine.NewBoard()
	fmt.Println(bord.GetBoardAsString())
}
