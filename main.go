package main

import "fmt"

type Square struct {
    num int //Holds the number. 0 is empty
}

type Grid struct {
    squares [4][4]Square
}

func createAndReturnNewGridWithInitialSquares() {
    squares := [4][4]Square{}
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            //Create a square and add it to squares
            squares[i][j] = Square{0} //num starts at 0, will be set later
        }
    }
    myGrid := Grid{squares}
    return myGrid
}

func main() {
    fmt.Println("Ola!")
    myGrid := createAndReturnNewGridWithInitialSquares()
    fmt.Println(myGrid)
}