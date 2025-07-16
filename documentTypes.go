package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type entranceReceipt struct {
	id      int
	date    time.Time
	truck   string
	trailer string
	gross   int
	tare    int
	net     int
}

func NewEntranceReceipt(id *int) entranceReceipt {
	scanner := bufio.NewScanner(os.Stdin)
	result := entranceReceipt{
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
	fmt.Println("Въведете количество тара.")
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
	fmt.Printf("Документ с номер %v е създаден успешно.\n", result.id)
	return result
}
