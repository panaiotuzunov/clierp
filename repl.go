package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func startRepl() {
	docEnumerator := 1
	receipts := []*entranceReceipt{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Изберете опцията от менюто. (само код)")
	fmt.Println("1. Справки")
	fmt.Println("2. Нов документ")
	for {
		scanner.Scan()
		selection := scanner.Text()
		switch selection {
		case "1":
			reference(receipts)
		case "2":
			receipts = append(receipts, newDoc(&docEnumerator))
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Невалиден избор.")
		}
	}
}

func reference(receipts []*entranceReceipt) {
	fmt.Printf("%-12v %-12v %-12v %-12v %-12v %-12v %-12v\n", "НОМЕР", "ДАТА", "КАМИОН", "РЕМАРКЕ", "БРУТО", "ТАРА", "НЕТО")
	for _, e := range receipts {
		fmt.Printf("%-12v %-12v %-12v %-12v %-12v %-12v %-12v\n", e.id, e.date.Format("dd-mm-yyyy"), e.truck, e.trailer, e.gross, e.tare, e.net)
	}
}

func newDoc(id *int) *entranceReceipt {
	scanner := bufio.NewScanner(os.Stdin)
	result := &entranceReceipt{
		id: *id,
	}
	result.date = time.Now()
	fmt.Println("Въведете номер на камион.")
	scanner.Scan()
	result.truck = scanner.Text()
	fmt.Println("Въведете номер на ремарке.")
	scanner.Scan()
	result.trailer = scanner.Text()
	fmt.Println("Въведете количество бруто.")
	for {
		scanner.Scan()
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Невалидно число: %v", err)
			continue
		}
		result.gross = num
		break
	}
	for {
		scanner.Scan()
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Невалидно число: %v", err)
			continue
		}
		result.tare = num
		break
	}
	result.net = result.gross - result.tare
	*id++
	return result
}
