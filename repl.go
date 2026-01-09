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
		fmt.Println("3. Покупки")
		fmt.Println("4. Продажби")
		scanner.Scan()
		selection := scanner.Text()
		switch selection {
		case "1":
			printEntranceAndExitReciepts(stateStruct)
		case "2":
			printInventory(stateStruct)
		case "3":
			printPurchases(stateStruct)
		case "4":
			printSales(stateStruct)
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
		fmt.Println("3. Договор за покупка")
		fmt.Println("4. Договор за продажба")
		scanner.Scan()
		selection := scanner.Text()
		switch selection {
		case "1":
			NewReceipt(scanner, stateStruct, receiptTypeEntrace)
		case "2":
			NewReceipt(scanner, stateStruct, receiptTypeExit)
		case "3":
			NewPurchase(scanner, stateStruct)
		case "4":
			NewSale(scanner, stateStruct)
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
