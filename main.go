package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-tty"
)

type Player struct {
	Name   string
	Symbol string
	Score  int
}

func PrintPlayerInfo(player1, player2, currentPlayer Player) {
	if player1 == currentPlayer {
		fmt.Println("\033[34m╔═════════════╗\033[0m     ╔═════════════╗")
		fmt.Printf("\033[34m║ %s%s ║\033[0m     ║ %s%s ║\n", player1.Symbol, player1.Name, player2.Symbol, player2.Name)
		fmt.Printf("\033[34m║      %d      ║\033[0m     ║      %d      ║\n", player1.Score, player2.Score)
		fmt.Println("\033[34m╚═════════════╝\033[0m     ╚═════════════╝")
	} else {
		fmt.Println("╔═════════════╗     \033[34m╔═════════════╗\033[0m")
		fmt.Printf("║ %s%s ║     \033[34m║ %s%s ║\033[0m\n", player1.Symbol, player1.Name, player2.Symbol, player2.Name)
		fmt.Printf("║      %d      ║     \033[34m║      %d      ║\033[0m\n", player1.Score, player2.Score)
		fmt.Println("╚═════════════╝     \033[34m╚═════════════╝\033[0m")
	}
}

func PrintBoard(gameBoard [6][7]string, player1, player2, currentPlayer Player) {
	PrintPlayerInfo(player1, player2, currentPlayer)

	fmt.Println("\n\n  1    2    3    4    5    6    7")
	fmt.Println("╔════╦════╦════╦════╦════╦════╦════╗")
	for i := range gameBoard {
		for j := range gameBoard[i] {
			if gameBoard[i][j] == "" {
				gameBoard[i][j] = "   "
			}
			fmt.Printf("║ %v", gameBoard[i][j])
			if j == 6 {
				fmt.Print("║")
			}
		}
		if i != 5 {
			fmt.Println("\n╠════╬════╬════╬════╬════╬════╬════╣")
		}
	}
	fmt.Println("\n╚════╩════╩════╩════╩════╩════╩════╝")
}

func ClearScreen() {
	fmt.Print("\033[H")
	fmt.Print("\033[0J")
	fmt.Print("\033[?25l")
}

func SwitchPlayer(currentPlayer, player1, player2 *Player) *Player {
	if currentPlayer == player1 {
		return player2
	}
	return player1
}

func IsValidMove(move int, board [6][7]string) bool {
	return move >= 0 && move < 7 && board[0][move] == ""
}

func DropPiece(move int, currentPlayer *Player, board *[6][7]string) {
	for i := 5; i >= 0; i-- {
		if board[i][move] == "" {
			board[i][move] = currentPlayer.Symbol
			break
		}
	}
}

func CheckWin(board [6][7]string) bool {
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			if board[i][j] != "" {
				// Check horizontally
				if j <= 3 && board[i][j] == board[i][j+1] && board[i][j] == board[i][j+2] && board[i][j] == board[i][j+3] {
					return true
				}
				// Check vertically
				if i <= 2 && board[i][j] == board[i+1][j] && board[i][j] == board[i+2][j] && board[i][j] == board[i+3][j] {
					return true
				}
				// Check diagonally (positive slope)
				if i <= 2 && j <= 3 && board[i][j] == board[i+1][j+1] && board[i][j] == board[i+2][j+2] && board[i][j] == board[i+3][j+3] {
					return true
				}
				// Check diagonally (negative slope)
				if i <= 2 && j >= 3 && board[i][j] == board[i+1][j-1] && board[i][j] == board[i+2][j-2] && board[i][j] == board[i+3][j-3] {
					return true
				}
			}
		}
	}
	return false
}

func Checkdrow(board [6][7]string) bool{
	for i := 0; i < 7; i++ {
		if board[0][i] == ""{
			return false
		}
	}
	return true
}

func PrintResult(winnerName string, tty tty.TTY) {
	if winnerName == "draw" {
		fmt.Println("\033[9;3H\033[48;5;241m                                ")
		fmt.Print("\033[10;3H              draw!             ")
		fmt.Println("\033[11;3H                                ")
		fmt.Println("\033[12;3H       Press 'q' to exit        ")
		fmt.Println("\033[13;3H   or press enter to continue   ")
		fmt.Println("\033[14;3H                                \033[0m ║")
	}else {
		fmt.Println("\033[9;3H\033[48;5;241m                                ")
		fmt.Printf("\033[10;3H         %s wins!         ", winnerName)
		fmt.Println("\033[11;3H                                ")
		fmt.Println("\033[12;3H       Press 'q' to exit        ")
		fmt.Println("\033[13;3H   or press enter to continue   ")
		fmt.Println("\033[14;3H                                \033[0m ║")
	}

	for {
		char, _ := tty.ReadRune()
		if char == 'q' {
			Quit()
		} else if char == 13 {
			return
		}
	}
}

func Quit() {
	fmt.Print("\033[H")
	fmt.Print("\033[0J")
	fmt.Print("\033[?25h")
	os.Exit(0)
}


func GetUserInput(tty tty.TTY) int {
	for {
		char, _ := tty.ReadRune()
		if char >= '0' && char <= '9' {
			return int(char - '0')
		}else if char == 'q' {
			Quit()
		}
	}
}

func main() {
	gameBoard := [6][7]string{}
	player1 := Player{Name: "Player 1", Symbol: "🔴 "}
	player2 := Player{Name: "Player 2", Symbol: "🟡 "}
	currentPlayer := &player1
	tty, _ := tty.Open()
	defer tty.Close()

	for {
		ClearScreen()
		PrintBoard(gameBoard, player1, player2, *currentPlayer)

		move := GetUserInput(*tty)

		move-- 
		if IsValidMove(move, gameBoard) {
			DropPiece(move, currentPlayer, &gameBoard)
			if Checkdrow(gameBoard){
				ClearScreen()
				PrintBoard(gameBoard, player1, player2, *currentPlayer)
				PrintResult("draw", *tty)
				gameBoard = [6][7]string{}
			}
			if CheckWin(gameBoard) {
				ClearScreen()
				PrintBoard(gameBoard, player1, player2, *currentPlayer)
				currentPlayer.Score++
				PrintResult(currentPlayer.Name, *tty)
				gameBoard = [6][7]string{}
			}
			currentPlayer = SwitchPlayer(currentPlayer, &player1, &player2)
		}
	}
}
