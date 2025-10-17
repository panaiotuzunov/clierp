package main

import (
	"bufio"
	"clierp/internal/database"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const docTypeExit = "Пропуск за извозване"
const docTypeEntrace = "Приемна бележка"

var refLineSeparator = strings.Repeat("-", 120)
var grainTypes = map[string]struct{}{
	"пшеница":    {},
	"ечемик":     {},
	"царевица":   {},
	"слънчоглед": {},
	"рапица":     {},
}

func NewReceipt(stateStruct *State, docType string) {
	scanner := bufio.NewScanner(os.Stdin)
	currentReceipt := database.CreateReceiptParams{
		DocType: docType,
	}
	fmt.Println("Въведете номер на камион.")
	scanner.Scan()
	currentReceipt.TruckReg = scanner.Text()
	fmt.Println("Въведете номер на ремарке.")
	scanner.Scan()
	currentReceipt.TrailerReg = scanner.Text()
	fmt.Println("Въведете вид зърно.")
	for {
		scanner.Scan()
		text := scanner.Text()
		if _, exist := grainTypes[text]; exist {
			currentReceipt.GrainType = text
			break
		}
		fmt.Println("Неизвестен тип зърно. Позволени са следните типове:")
		for grain := range grainTypes {
			fmt.Println(grain)
		}
	}
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
	if docType == docTypeExit {
		currentReceipt.Gross *= -1
		currentReceipt.Tare *= -1
		currentReceipt.Net *= -1
	}
	if err := stateStruct.db.CreateReceipt(context.Background(), currentReceipt); err != nil {
		fmt.Println("Неуспешно създаване на документа.")
	}
	fmt.Println("Документът е създаден успешно.")
	fmt.Println(refLineSeparator)
}
