package main

import (
	"bufio"
	"clierp/internal/database"
	"context"
	"database/sql"
	"fmt"
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

func NewReceipt(scanner *bufio.Scanner, stateStruct *State, docType string) {
	currentReceipt := database.CreateReceiptParams{
		DocType: docType,
	}
	if docType == docTypeEntrace {
		purchases, err := stateStruct.db.GetAllPurchases(context.Background())
		if len(purchases) == 0 {
			fmt.Println("Приемна бележка се създава на база на договор за покупка. В момента няма активни договори. Моля, създайте нов договор.")
			return
		}
		if err != nil {
			fmt.Println("Грешка при търсене на договори.")
			return
		}
		fmt.Println("Изберете номер договoр за покупка. Активни договори към момента са:")
		printPurchases(stateStruct)
		for {
			purchase, err := stateStruct.db.GetPurchaseById(context.Background(), int32(scanInt(scanner)))
			if err != nil {
				fmt.Println("Невалиден номер на договор опитайте пак.")
				continue
			}
			currentReceipt.PurchaseID = sql.NullInt32{
				Valid: true,
				Int32: purchase.ID,
			}
			break
		}
	}
	fmt.Println("Въведете номер на камион.")
	scanner.Scan()
	currentReceipt.TruckReg = scanner.Text()
	fmt.Println("Въведете номер на ремарке.")
	scanner.Scan()
	currentReceipt.TrailerReg = scanner.Text()
	fmt.Println("Въведете вид зърно.")
	currentReceipt.GrainType = scanGrainType(scanner)
	fmt.Println("Въведете количество бруто.")
	currentReceipt.Gross = int32(scanInt(scanner))
	fmt.Println("Въведете количество тара.")
	currentReceipt.Tare = int32(scanInt(scanner))
	currentReceipt.Net = currentReceipt.Gross - currentReceipt.Tare
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

func NewPurchase(scanner *bufio.Scanner, stateStruct *State) {
	var purchase database.CreatePurchaseParams
	fmt.Println("Въведете доставчик.")
	scanner.Scan()
	purchase.Suplier = scanner.Text()
	fmt.Println("Въведете вид зърно.")
	purchase.GrainType = scanGrainType(scanner)
	fmt.Println("Въведете количество.")
	purchase.Quantity = int32(scanInt(scanner))
	fmt.Println("Въведете цена.")
	purchase.Price = int32(scanInt(scanner))
	if err := stateStruct.db.CreatePurchase(context.Background(), purchase); err != nil {
		fmt.Println("Неуспешно създаване на документа.")
	}
	fmt.Println("Документът е създаден успешно.")
	fmt.Println(refLineSeparator)
}

func scanGrainType(scanner *bufio.Scanner) string {
	for {
		scanner.Scan()
		text := scanner.Text()
		if _, exist := grainTypes[text]; exist {
			return text
		}
		fmt.Println("Неизвестен тип зърно. Позволени са следните типове:")
		for grain := range grainTypes {
			fmt.Println(grain)
		}
	}
}

func scanInt(scanner *bufio.Scanner) int {
	for {
		scanner.Scan()
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Невалидно число: %v", err)
			continue
		}
		return num
	}
}
