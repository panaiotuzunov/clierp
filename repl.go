package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	docEnumerator := 1
	receipts := []entranceReceipt{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Изберете опцията от менюто. (само код)")
		fmt.Println("1. Справки")
		fmt.Println("2. Нов документ")
		scanner.Scan()
		selection := scanner.Text()
		switch selection {
		case "1":
			referenceRepl(&receipts)
		case "2":
			newDocRepl(&docEnumerator, &receipts, scanner)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Невалиден избор.")
		}
	}
}

func referenceRepl(receipts *[]entranceReceipt) {
	fmt.Printf("%-12v %-12v %-12v %-12v %-12v %-12v %-12v\n", "НОМЕР", "ДАТА", "КАМИОН", "РЕМАРКЕ", "БРУТО", "ТАРА", "НЕТО")
	for _, e := range *receipts {
		fmt.Printf("%-12v %-12v %-12v %-12v %-12v %-12v %-12v\n", e.id, e.date.Format("02/01/2006"), e.truck, e.trailer, e.gross, e.tare, e.net)
	}
}

func newDocRepl(id *int, receipts *[]entranceReceipt, scanner *bufio.Scanner) {
outer:
	for {
		fmt.Println("Изберете тип документ. За връщане назад изберете '0'")
		fmt.Println("1. Приемна бележка")
		fmt.Println("2. Пропуск за извозване")
		scanner.Scan()
		selection := scanner.Text()
		switch selection {
		case "1":
			*receipts = append(*receipts, NewEntranceReceipt(id))
		case "2":
			fmt.Println("Документа все още не съществува.")
		case "0":
			break outer
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Невалиден избор.")
		}
	}
}
