/*
 * Chopsticks.go created on Sun Jun 26 2022 by King Red Sanchez
 * Copyright (c) 2022
 */

package main

import (
	"fmt"
	st "strings"
	t "time"
)

var (
	p              = fmt.Print
	s              = fmt.Scan
	boolWinner     = false
	intMoveCounter = 0
	strPlayer1Name = ""
	arrPlayer1Hand = [2]int{1, 1}
	strPlayer2Name = ""
	arrPlayer2Hand = [2]int{1, 1}
	strMoveType    = ""
	strFrom        = ""
	strTo          = ""
	intLeft        = 0
	intRight       = 0
	intTotal       = 0
	intHandTotal   = 0
)

func main() {
	gameInitialize()
	for !boolWinner {
		showHands()
		if intMoveCounter%2 == 0 {
			p("\nMove #", intMoveCounter+1, " is ", strPlayer1Name, "'s turn\n")
			moveMaker()
			if loseChecker(arrPlayer1Hand) {
				boolWinner = loseChecker(arrPlayer1Hand)
			} else if loseChecker(arrPlayer2Hand) {
				boolWinner = loseChecker(arrPlayer2Hand)
			}
		} else {
			p("\nMove #", intMoveCounter+1, " is ", strPlayer2Name, "'s turn\n")
			moveMaker()
			if loseChecker(arrPlayer1Hand) {
				boolWinner = loseChecker(arrPlayer1Hand)
			} else if loseChecker(arrPlayer2Hand) {
				boolWinner = loseChecker(arrPlayer2Hand)
			}
		}
		intMoveCounter++
	}
	showHands()
	if loseChecker(arrPlayer1Hand) {
		p(strPlayer2Name, " wins! Congratulations!")
	} else if loseChecker(arrPlayer2Hand) {
		p(strPlayer1Name, " wins! Congratulations!")
	}
}

func moveMaker() {
	p("Move: ")
	s(&strMoveType)
	strMoveType = st.ToLower(strMoveType[:1])
	switch strMoveType {
	case "a":
		s(&strFrom, &strTo)
		strFrom = st.ToLower(strFrom)
		strTo = st.ToLower(strTo)
		if strFrom != "l" && strFrom != "r" || strTo != "l" && strTo != "r" {
			p("Hand input error occurred. Terminating Game.\n")
			boolWinner = true
		}
		if intMoveCounter%2 == 0 {
			// Player1 attacks player2
			if strFrom == "l" && strTo == "l" {
				if arrPlayer1Hand[0] == 0 || arrPlayer2Hand[0] == 0 {
					p("Transfer error occurred. Terminating Game.\n")
					boolWinner = true
				}
				intTotal = arrPlayer1Hand[0] + arrPlayer2Hand[0]
				if intTotal >= 5 {
					intTotal %= 5
				}
				arrPlayer2Hand[0] = intTotal
			} else if strFrom == "l" && strTo == "r" {
				if arrPlayer1Hand[0] == 0 || arrPlayer2Hand[1] == 0 {
					p("Transfer error occurred. Terminating Game.\n")
					boolWinner = true
				}
				intTotal = arrPlayer1Hand[0] + arrPlayer2Hand[1]
				if intTotal >= 5 {
					intTotal %= 5
				}
				arrPlayer2Hand[1] = intTotal
			} else if strFrom == "r" && strTo == "l" {
				if arrPlayer1Hand[1] == 0 || arrPlayer2Hand[0] == 0 {
					p("Transfer error occurred. Terminating Game.\n")
					boolWinner = true
				}
				intTotal = arrPlayer1Hand[1] + arrPlayer2Hand[0]
				if intTotal >= 5 {
					intTotal %= 5
				}
				arrPlayer2Hand[0] = intTotal
			} else {
				if arrPlayer1Hand[1] == 0 || arrPlayer2Hand[1] == 0 {
					p("Transfer error occurred. Terminating Game.\n")
					boolWinner = true
				}
				intTotal = arrPlayer1Hand[1] + arrPlayer2Hand[1]
				if intTotal >= 5 {
					intTotal %= 5
				}
				arrPlayer2Hand[1] = intTotal
			}
		} else {
			// Player2 attacks player1
			if strFrom == "l" && strTo == "l" {
				if arrPlayer2Hand[0] == 0 || arrPlayer1Hand[0] == 0 {
					p("Transfer error occurred. Terminating Game.\n")
					boolWinner = true
				}
				intTotal = arrPlayer2Hand[0] + arrPlayer1Hand[0]
				if intTotal >= 5 {
					intTotal %= 5
				}
				arrPlayer1Hand[0] = intTotal
			} else if strFrom == "l" && strTo == "r" {
				if arrPlayer2Hand[0] == 0 || arrPlayer1Hand[1] == 0 {
					p("Transfer error occurred. Terminating Game.\n")
					boolWinner = true
				}
				intTotal = arrPlayer2Hand[0] + arrPlayer1Hand[1]
				if intTotal >= 5 {
					intTotal %= 5
				}
				arrPlayer1Hand[1] = intTotal
			} else if strFrom == "r" && strTo == "l" {
				if arrPlayer2Hand[1] == 0 || arrPlayer1Hand[0] == 0 {
					p("Transfer error occurred. Terminating Game.\n")
					boolWinner = true
				}
				intTotal = arrPlayer2Hand[1] + arrPlayer1Hand[0]
				if intTotal >= 5 {
					intTotal %= 5
				}
				arrPlayer1Hand[0] = intTotal
			} else {
				if arrPlayer2Hand[1] == 0 || arrPlayer1Hand[1] == 0 {
					p("Transfer error occurred. Terminating Game.\n")
					boolWinner = true
				}
				intTotal = arrPlayer2Hand[1] + arrPlayer1Hand[1]
				if intTotal >= 5 {
					intTotal %= 5
				}
				arrPlayer1Hand[1] = intTotal
			}
		}
		p("Move is: ", strMoveType, "\tFrom: ", strFrom, "\tTo: ", strTo, "\n\n")
	case "d":
		s(&intLeft, &intRight)
		if intLeft < 0 || intLeft > 5 || intRight < 0 || intRight > 5 {
			p("Move range error occurred. Terminating Game.\n")
			boolWinner = true
		}
		intTotal = intLeft + intRight
		if intMoveCounter%2 == 0 {
			// Player1 splits
			intHandTotal = arrPlayer1Hand[0] + arrPlayer1Hand[1]
			if intTotal != intHandTotal {
				p("Divide range error occurred. Terminating Game.\n")
				boolWinner = true
			}
			if arrPlayer1Hand[0] == intLeft && arrPlayer1Hand[1] == intRight {
				p("Repeated hand error occurred. Terminating Game.\n")
				boolWinner = true
			}
			if arrPlayer1Hand[0] == intRight && arrPlayer1Hand[1] == intLeft {
				p("Switched hand error occurred. Terminating Game.\n")
				boolWinner = true
			}
			arrPlayer1Hand[0] = intLeft
			arrPlayer1Hand[1] = intRight
			if arrPlayer1Hand[0] > 5 || arrPlayer1Hand[1] > 5 {
				p("Invalid divide error occurred. Terminating Game.\n")
				boolWinner = true
			}
			if arrPlayer1Hand[0] == 5 {
				arrPlayer1Hand[0] = 0
			} else if arrPlayer1Hand[1] == 5 {
				arrPlayer1Hand[1] = 0
			}

		} else {
			// Player2 splits
			intHandTotal = arrPlayer2Hand[0] + arrPlayer2Hand[1]
			if intTotal != intHandTotal {
				p("Divide range error occurred. Terminating Game.\n")
				boolWinner = true
			}
			if arrPlayer2Hand[0] == intLeft && arrPlayer2Hand[1] == intRight {
				p("Repeated hand error occurred. Terminating Game.\n")
				boolWinner = true
			}
			if arrPlayer2Hand[0] == intRight && arrPlayer2Hand[1] == intLeft {
				p("Switched hand error occurred. Terminating Game.\n")
				boolWinner = true
			}
			arrPlayer2Hand[0] = intLeft
			arrPlayer2Hand[1] = intRight
			if arrPlayer2Hand[0] > 5 || arrPlayer2Hand[1] > 5 {
				p("Invalid divide error occurred. Terminating Game.\n")
				boolWinner = true
			}
			if arrPlayer2Hand[0] == 5 {
				arrPlayer2Hand[0] = 0
			} else if arrPlayer2Hand[1] == 5 {
				arrPlayer2Hand[1] = 0
			}
		}
		p("Move is: ", strMoveType, "\tLeft: ", intLeft, "\tRight: ", intRight, "\n\n")
	default:
		p("Move type input error occurred. Terminating Game.\n")
		boolWinner = true
	}

}

func loseChecker(hand [2]int) bool {
	if (hand[0] == 0) && (hand[1] == 0) {
		return true
	} else {
		return false
	}
}

func showHands() {
	p(strPlayer1Name, "\nLeft: ", arrPlayer1Hand[0], "\tRight: ", arrPlayer1Hand[1], "\n")
	p(strPlayer2Name, "\nLeft: ", arrPlayer2Hand[0], "\tRight: ", arrPlayer2Hand[1], "\n")
}

func gameInitialize() {
	p("Welcome to Chopsticks!\nThis game is made with the fickle console I/O of Go Programming Language, so enter only the right data types and singe word strings.\n")
	t.Sleep(1 * t.Second)
	p("Player 1, enter your name now: ")
	s(&strPlayer1Name)
	p("Player 2, enter your name now: ")
	s(&strPlayer2Name)
	p("\nPlayers ", strPlayer1Name, " and ", strPlayer2Name, "\nthe game is commencing\n\n")
	t.Sleep(2 * t.Second)
}
