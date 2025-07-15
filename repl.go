package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Изберете опцията от менюто. (само код)")
	fmt.Println("1. Справки")
	fmt.Println("2. Нов документ")
	for {
		scanner.Scan()
		selection := scanner.Text()
		switch selection {
		case "1":
			reference()
		case "2":
			newDoc()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Невалиден избор. ")
		}
	}
}

func reference() {
	fmt.Printf("%-12v %-12v %-12v %-12v %-12v %-12v %-12v\n", "НОМЕР", "ДАТА", "КАМИОН", "РЕМАРКЕ", "БРУТО", "ТАРА", "НЕТО")
	e := newDoc()
	fmt.Printf("%-12v %-12v %-12v %-12v %-12v %-12v %-12v\n", e.num, e.date.Format("dd-mm-yyyy"), e.truck, e.trailer, e.gross, e.tare, e.net)
}

func newDoc() *entranceReceipt {
	result := &entranceReceipt{
		num:     1,
		date:    time.Now(),
		truck:   "СВ7319ТМ",
		trailer: "СН5494АХ",
		gross:   40000,
		tare:    15000,
	}
	result.net = result.gross - result.tare
	return result
}
