package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl(stateStruct *State) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Изберете опцията от менюто. (само код)")
		fmt.Println("1. Справки")
		fmt.Println("2. Нов документ")
		scanner.Scan()
		selection := scanner.Text()
		switch selection {
		case "1":
			referenceRepl(scanner, stateStruct)
		case "2":
			newDocRepl(scanner, stateStruct)
		case "exit":
			os.Exit(0)
		case "quit":
			os.Exit(0)
		default:
			fmt.Println("Невалиден избор.")
		}
	}
}

func referenceRepl(scanner *bufio.Scanner, stateStruct *State) {
outer:
	for {
		fmt.Println("Изберете справка от каталога. За връщане назад изберете '0'")
		fmt.Println("1. Кантарна книга")
		fmt.Println("2. Наличност")
		scanner.Scan()
		selection := scanner.Text()
		switch selection {
		case "1":
			printEntranceAndExitReciepts(stateStruct)
		case "2":
			printInventory(stateStruct)
		case "0":
			break outer
		case "exit":
			os.Exit(0)
		case "quit":
			os.Exit(0)
		default:
			fmt.Println("Невалиден избор.")
		}
	}
}

func newDocRepl(scanner *bufio.Scanner, stateStruct *State) {
outer:
	for {
		fmt.Println("Изберете тип документ. За връщане назад изберете '0'")
		fmt.Println("1. Приемна бележка")
		fmt.Println("2. Пропуск за извозване")
		scanner.Scan()
		selection := scanner.Text()
		switch selection {
		case "1":
			NewReceipt(stateStruct, docTypeEntrace)
		case "2":
			NewReceipt(stateStruct, docTypeExit)
		case "0":
			break outer
		case "exit":
			os.Exit(0)
		case "quit":
			os.Exit(0)
		default:
			fmt.Println("Невалиден избор.")
		}
	}
}
