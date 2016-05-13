package main

import "fmt"
import "math/rand"

type Grid struct {
    squares [4][4]int
}

func createAndReturnNewGridWithInitialSquares() Grid {
    i1, j1, num1, i2, j2, num2 := rand.Intn(4), rand.Intn(4), rand.Intn(2), rand.Intn(4), rand.Intn(4), rand.Intn(2)
    fmt.Println(i1, j1, num1, i2, j2, num2)
    squares := [4][4]int{}
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            //Create a square and add it to squares
            if i1 == i {
                if j1 == j {
                    squares[i][j] = (num1 + 1) * 2
                }
            } else if i2 == i {
                if j2 == j {
                    squares[i][j] = (num2 + 1) * 2
                }
            } else {
                squares[i][j] = 0 //num starts at 0, will be set later
            }
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