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
			reference(&receipts)
		case "2":
			newDoc(&docEnumerator, &receipts)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Невалиден избор.")
		}
	}
}

func reference(receipts *[]entranceReceipt) {
	fmt.Printf("%-12v %-12v %-12v %-12v %-12v %-12v %-12v\n", "НОМЕР", "ДАТА", "КАМИОН", "РЕМАРКЕ", "БРУТО", "ТАРА", "НЕТО")
	for _, e := range *receipts {
		fmt.Printf("%-12v %-12v %-12v %-12v %-12v %-12v %-12v\n", e.id, e.date.Format("00/00/0000"), e.truck, e.trailer, e.gross, e.tare, e.net)
	}
}

func newDoc(id *int, receipts *[]entranceReceipt) {
	*receipts = append(*receipts, NewEntranceReceipt(id))
}
