package main

import (
	"fmt"
	"math/rand"
	"minesweeper/gameStats"
	"strconv"
	"time"
)

/* ------------------------------------------------------ Overview & Running ---------------------------------------------------------------------- */

// This is a refactor of my Java Minesweeper project into GO, this is to solidify my learning of basic Go concepts/syntax

// To run do: go run .

/* ------------------------------------------------------- Global Variables ----------------------------------------------------------------------- */

const (

	heading = 			
	"\n  __  __   _                   _____                                                 \n" +
	" |  \\/  | (_)                 / ____|                                               \n" +
	" | \\  / |  _   _ __     ___  | (___   __      __   ___    ___   _ __     ___   _ __ \n" +
	" | |\\/| | | | | '_ \\   / _ \\  \\___ \\  \\ \\ /\\ / /  / _ \\  / _ \\ | '_ \\   / _ \\ | '__|\n" +
	" | |  | | | | | | | | |  __/  ____) |  \\ V  V /  |  __/ |  __/ | |_) | |  __/ | |   \n" +
	" |_|  |_| |_| |_| |_|  \\___| |_____/    \\_/\\_/    \\___|  \\___| | .__/   \\___| |_|   \n" +
	"                                                               | |                  \n" +
	"                                                               |_|                  \n"  

	rules = 
	"\nThe rules of this simplified minesweeper clone are: \n" +
	"on the grid, there are 100 cells, a cell could be empty \n"+
	"or it could contain a bomb. In the console, type in an X \n" +
	"(i.e. 0 to 9) coordinate and then a Y coordinate (also 0 to 9). \n" + "If that location contains a bomb it will be game over. \n "

	version = "\nVersion: Go \nMade By Stanley \n"

	question = "?"
	bomb = "\u001B[37m" + "?" + "\u001B[0m" // can change ? to X to test it works and to see where bombs are
	rBomb = "\u001B[31m" + "X" + "\u001B[0m"
)

var (
	gameOver bool = false
	nBombs int = 15
	clearCells int = 99 - nBombs
	count int = 0
	gameGrid = [10][10]string{}
	xInput int
	yInput int
	correctInput bool = false
	startInput string
	bombDisplay int = 0
)

/* ------------------------------------------------------------ Functions ------------------------------------------------------------------------- */

// General

func Wait(sec time.Duration){
	time.Sleep(sec * time.Second)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Game Functions

func GameStart(){
	for (!correctInput){

		fmt.Printf("\nPlease enter S or s to start playing \n")
		fmt.Scan(&startInput)

		if (startInput == "S" || startInput == "s"){
			correctInput = true
		}

	}
}

func GameEnd(){
	fmt.Println("Game will exit")
	Wait(1)
	fmt.Println("Game exited")
}

func GameRestart(){
	Wait(1)
	fmt.Println("Game Restarting....")
	Wait(1)
	fmt.Println("Updating stats....")
	Wait(1)
	fmt.Println("Game Restarted")
	Wait(1)
	Intro()
	InitGame()
}

// GameGrid Functions

func FillGrid(){
	for i:= 0; i < len(gameGrid); i++{
		for j:= 0; j < len(gameGrid[i]); j++{
			gameGrid[i][j] = question
		}
	}
}

func PlaceBombs(){
	
	var bombCount int = 0

	for(bombCount <= nBombs){

		rR := rand.Intn(10)
		rC := rand.Intn(10)
		if(gameGrid[rR][rC] == question){
			gameGrid[rR][rC] = bomb
			// fmt.Printf("\nThere is a bomb at X:%v Y:%v \n", rR, rC)
			// un-comment above line to see where bombs are if you are testing it works
			bombCount++
		}
	
	}

}

func DisplayGrid(){
	fmt.Println("")
	for _, row := range gameGrid{
		for _, cell := range row{
			fmt.Printf("| %v |", cell)
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func AdjacentBombs(xInput int, yInput int){

	if (gameGrid[xInput][yInput] == question){
		
      for x := max(0, xInput-1); x <= min(9, xInput+1); x++ {
            for y := max(0, yInput-1); y <= min(9, yInput+1); y++ {
                if gameGrid[x][y] == bomb && (x != xInput || y != yInput) {
                    bombDisplay++
                }
            }
        }
		}
	
}

func CheckInput(){

	for(!gameOver && count < clearCells){

		fmt.Printf("Please enter a X coordinate \n")
		xInput := 0
		fmt.Scan(&xInput)
		fmt.Printf("\nPlease enter a Y coordinate \n")
		yInput := 0
		fmt.Scan(&yInput)

		if (xInput >= 0 && xInput <= 9  && yInput >= 0 && yInput <= 9 && gameGrid[xInput][yInput] == question){

			count++
			AdjacentBombs(xInput, yInput)
			if(bombDisplay > 0){
				gameGrid[xInput][yInput] = "\u001B[38;2;255;165;0m" + strconv.FormatInt(int64(bombDisplay), 10) + "\u001B[0m"
				fmt.Printf("\n That cell is clear! There are %v cells remaining \n", clearCells - count)
				DisplayGrid()
			} else {
			gameGrid[xInput][yInput] = " "
			fmt.Printf("\n That cell is clear! There are %v cells remaining \n", clearCells - count)
			DisplayGrid()
			}

		} else if (xInput >= 10 || xInput < 0  || yInput >= 10 || yInput < 0){

			fmt.Printf("\nInvalid input \n \n")

		}  else if (gameGrid[xInput][yInput] != question){

			gameGrid[xInput][yInput] = rBomb
			DisplayGrid()

			Wait(1)
			fmt.Printf("\nBOOM! \n")
			fmt.Printf("\n========================================= GAME OVER ========================================= \n")
			gameStats.IncrementLosses()

			GameRestart()

		}
	}

	if (count == clearCells){

		Wait(1)
		fmt.Println("Congrats, you won!")
		gameStats.IncrementWins()

		GameRestart()
	}

}

// GameState

func Intro(){

	fmt.Print(heading)
	Wait(1)
	fmt.Print(version)
	Wait(1)
	fmt.Printf("\nYour Wins: %v & Losses: %v \n", gameStats.Wins, gameStats.Losses)
	Wait(1)
	fmt.Print(rules)
	Wait(1)

}

func InitGame(){

	GameStart()
	Wait(1)
	FillGrid()
	PlaceBombs()
	DisplayGrid()
	Wait(1)
	CheckInput()
	
}

// Game Runner

func main(){

Intro()
InitGame()

}