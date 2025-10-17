package main

import (
	"bufio"
	"clierp/internal/database"
	"context"
	"fmt"
	"os"
	"strconv"
)

func NewEntranceReceipt(stateStruct *State) {
	scanner := bufio.NewScanner(os.Stdin)
	currentReceipt := database.CreateReceiptParams{}
	fmt.Println("Въведете номер на камион.")
	scanner.Scan()
	currentReceipt.TruckReg = scanner.Text()
	fmt.Println("Въведете номер на ремарке.")
	scanner.Scan()
	currentReceipt.TrailerReg = scanner.Text()
	fmt.Println("Въведете количество бруто.")
	for {
		scanner.Scan()
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Невалидно число: %v", err)
			continue
		}
		currentReceipt.Gross = int32(num)
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
		currentReceipt.Tare = int32(num)
		currentReceipt.Net = currentReceipt.Gross - currentReceipt.Tare
		break
	}
	if err := stateStruct.db.CreateReceipt(context.Background(), currentReceipt); err != nil {
		fmt.Println("Неуспешно създаване на документа.")
	}
	fmt.Println("Документът е създаден успешно.")
	fmt.Println("----------------------------------------------------------------------------------------")
}
