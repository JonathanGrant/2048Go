package main

import (
    "fmt"
    "math/rand"
    "bufio"
    "os"
    "strings"
)

type Grid struct {
    squares [4][4]int
}

type xyLoc struct {
    y int
    x int
}

func getTakenSquares(grid [4][4] int) []xyLoc {
    var takenSquares []xyLoc
    //loop through the grid and add the takenSquares to the slice
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            if grid[i][j] != 0 {
                takenSquares = append(takenSquares, xyLoc{i, j})
            }
        }
    }
    return takenSquares
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
        grid[square.y][square.x] = (num + 1) * 2
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

func tiltGrid(grid [4][4]int, direction int) [4][4]int {
    takenSquares := getTakenSquares(grid)
    collisionHappened := false
    if direction == 1 {
        //Up
        for index := range takenSquares {
            square := takenSquares[index]
            for square.y != 0 {
                //check if the space above it is free
                if grid[square.y-1][square.x] == 0 {
                    //swap them!
                    grid[square.y-1][square.x] = grid[square.y][square.x]
                    grid[square.y][square.x] = 0
                    square = xyLoc{square.y-1, square.x}
                } else {
                    //Check for a collision
                    if grid[square.y-1][square.x] == grid[square.y][square.x] {
                        grid[square.y-1][square.x] = grid[square.y-1][square.x] * 2
                        grid[square.y][square.x] = 0
                        collisionHappened = true
                        break
                    } else {
                        break
                    }
                }
            }
        }
    } else if direction == 2 {
        //Down
        for index := range takenSquares {
            square := takenSquares[index]
            for square.y != 3 {
                //check if the space above it is free
                if grid[square.y+1][square.x] == 0 {
                    //swap them!
                    grid[square.y+1][square.x] = grid[square.y][square.x]
                    grid[square.y][square.x] = 0
                    square = xyLoc{square.y+1, square.x}
                } else {
                    //Check for a collision
                    if grid[square.y+1][square.x] == grid[square.y][square.x] {
                        grid[square.y+1][square.x] = grid[square.y+1][square.x] * 2
                        grid[square.y][square.x] = 0
                        collisionHappened = true
                        break
                    } else {
                        break
                    }
                }
            }
        }
    } else if direction == 3 {
        //Left
        for index := range takenSquares {
            square := takenSquares[index]
            for square.x != 0 {
                //check if the space above it is free
                if grid[square.y][square.x-1] == 0 {
                    //swap them!
                    grid[square.y][square.x-1] = grid[square.y][square.x]
                    grid[square.y][square.x] = 0
                    square = xyLoc{square.y, square.x-1}
                } else {
                    //Check for a collision
                    if grid[square.y][square.x-1] == grid[square.y][square.x] {
                        grid[square.y][square.x-1] = grid[square.y][square.x-1] * 2
                        grid[square.y][square.x] = 0
                        collisionHappened = true
                        break
                    } else {
                        break
                    }
                }
            }
        }
    } else if direction == 4 {
        //Right
        for index := range takenSquares {
            square := takenSquares[index]
            for square.x != 3 {
                //check if the space above it is free
                if grid[square.y][square.x+1] == 0 {
                    //swap them!
                    grid[square.y][square.x+1] = grid[square.y][square.x]
                    grid[square.y][square.x] = 0
                    square = xyLoc{square.y, square.x+1}
                } else {
                    //Check for a collision
                    if grid[square.y][square.x+1] == grid[square.y][square.x] {
                        grid[square.y][square.x+1] = grid[square.y][square.x+1] * 2
                        grid[square.y][square.x] = 0
                        collisionHappened = true
                        break
                    } else {
                        break
                    }
                }
            }
        }
    }
    return grid
    if collisionHappened {
        return grid//tiltGrid(grid, direction)
    } else {
        return grid
    }
}

func movePlayed(grid [4][4]int, direction int) [4][4]int {
    grid = tiltGrid(grid, direction)
    grid = addTwoOrFourToGrid(grid)
    return grid
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Ola!")
    myGrid := createAndReturnNewGridWithInitialSquares()
    printGrid(myGrid.squares)
    //Okay now that I have a grid, I need to be able to take in user input
    fmt.Println("Make your move!")
    for {
        text, _ := reader.ReadString('\n')
        if strings.Compare(text,"up\n") == 0 {
            myGrid.squares = movePlayed(myGrid.squares, 1)
            printGrid(myGrid.squares)
        } else if strings.Compare(text,"down\n") == 0 {
            myGrid.squares = movePlayed(myGrid.squares, 2)
            printGrid(myGrid.squares)
        } else if strings.Compare(text,"left\n") == 0 {
            myGrid.squares = movePlayed(myGrid.squares, 3)
            printGrid(myGrid.squares)
        } else if strings.Compare(text,"right\n") == 0 {
            myGrid.squares = movePlayed(myGrid.squares, 4)
            printGrid(myGrid.squares)
        }
    }
}