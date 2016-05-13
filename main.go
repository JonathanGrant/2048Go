package main

import "fmt"
import "math/rand"

type Grid struct {
    squares [4][4]int
}

type xyLoc struct {
    x int
    y int
}

func getFreeSquares(grid [4][4] int) []xyLoc {
    var freeSquares []xyLoc
    //loop through the grid and add the freeSquares to the slice
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            if grid[i][j] == 0 {
                freeSquares = append(freeSquares, xyLoc{i, j})
            }
        }
    }
    return freeSquares
}

func addTwoOrFourToGrid(grid [4][4]int) [4][4]int {
    freeSquares := getFreeSquares(grid)
    if len(freeSquares) > 0 {
        square, num := freeSquares[rand.Intn(len(freeSquares))], rand.Intn(2)
        grid[square.x][square.y] = (num + 1) * 2
    }
    return grid
}

func createAndReturnNewGridWithInitialSquares() Grid {
    rand.Seed(44)
    squares := [4][4]int{}
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            //Create a square and add it to squares
            squares[i][j] = 0 //num starts at 0, will be set later
        }
    }
    //Now add the starting numbers twice
    squares = addTwoOrFourToGrid(squares)
    squares = addTwoOrFourToGrid(squares)
    myGrid := Grid{squares}
    return myGrid
}

func printGrid(grid [4][4]int) {
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            fmt.Print(grid[i][j])
        }
        fmt.Println()
    }
}

func main() {
    fmt.Println("Ola!")
    myGrid := createAndReturnNewGridWithInitialSquares()
    printGrid(myGrid.squares)
    //Okay now that I have a grid, I need to be able to take in user input
}