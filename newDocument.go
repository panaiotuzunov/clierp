package main

import (
	"bufio"
	"clierp/internal/database"
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const receiptTypeExit = "Пропуск за извозване"
const receiptTypeEntrace = "Приемна бележка"

var refLineSeparator = strings.Repeat("-", 120)
var grainTypes = map[string]struct{}{
	"пшеница":    {},
	"ечемик":     {},
	"царевица":   {},
	"слънчоглед": {},
	"рапица":     {},
}

func NewReceipt(scanner *bufio.Scanner, stateStruct *State, receiptType string) {
	currentReceipt := database.CreateReceiptParams{
		DocType: receiptType,
	}
	if receiptType == receiptTypeEntrace {
		purchases, err := stateStruct.db.GetAllPurchases(context.Background())
		if len(purchases) == 0 {
			fmt.Println("Приемна бележка се създава на база на договор за покупка. В момента няма активни договори. Моля, създайте нов договор.")
			return
		}
		if err != nil {
			fmt.Printf("Грешка при търсене на договори - %v\n", err)
			return
		}
		fmt.Println("Изберете номер na договoр за покупка. Активни договори към момента са:")
		printPurchases(stateStruct)
		for {
			purchase, err := stateStruct.db.GetPurchaseById(context.Background(), int32(scanInt(scanner)))
			if err != nil {
				fmt.Printf("Невалиден номер на договор опитайте пак - %v\n", err)
				continue
			}
			currentReceipt.PurchaseID = sql.NullInt32{
				Valid: true,
				Int32: purchase.ID,
			}
			currentReceipt.SaleID = sql.NullInt32{
				Valid: false,
			}
			currentReceipt.GrainType = purchase.GrainType
			break
		}
	} else { // receiptType == receiptTypeExit
		sales, err := stateStruct.db.GetAllSales(context.Background())
		if len(sales) == 0 {
			fmt.Println("Пропуск за извозване се създава на база на договор за продажба. В момента няма активни договори. Моля, създайте нов договор.")
			return
		}
		if err != nil {
			fmt.Printf("Грешка при търсене на договори - %v\n", err)
			return
		}
		fmt.Println("Изберете номер na договoр за продажба. Активни договори към момента са:")
		printAllSales(stateStruct)
		for {
			sale, err := stateStruct.db.GetSaleById(context.Background(), int32(scanInt(scanner)))
			if err != nil {
				fmt.Printf("Невалиден номер на договор опитайте пак - %v\n", err)
				continue
			}
			currentReceipt.SaleID = sql.NullInt32{
				Valid: true,
				Int32: sale.ID,
			}
			currentReceipt.PurchaseID = sql.NullInt32{
				Valid: false,
			}
			currentReceipt.GrainType = sale.GrainType
			break
		}
	}
	fmt.Println("Въведете номер на камион.")
	scanner.Scan()
	currentReceipt.TruckReg = scanner.Text()
	fmt.Println("Въведете номер на ремарке.")
	scanner.Scan()
	currentReceipt.TrailerReg = scanner.Text()
	fmt.Println("Въведете количество бруто.")
	currentReceipt.Gross = int32(scanInt(scanner))
	fmt.Println("Въведете количество тара.")
	currentReceipt.Tare = int32(scanInt(scanner))
	currentReceipt.Net = currentReceipt.Gross - currentReceipt.Tare
	if receiptType == receiptTypeExit {
		currentReceipt.Gross *= -1
		currentReceipt.Tare *= -1
		currentReceipt.Net *= -1
	}
	if err := stateStruct.db.CreateReceipt(context.Background(), currentReceipt); err != nil {
		fmt.Printf("Неуспешно създаване на документа - %v\n", err)
		return
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
		fmt.Printf("Неуспешно създаване на документа - %v\n", err)
		return
	}
	fmt.Println("Документът е създаден успешно.")
	fmt.Println(refLineSeparator)
}

func NewSale(scanner *bufio.Scanner, stateStruct *State) {
	var sale database.CreateSaleParams
	fmt.Println("Въведете клиент.")
	scanner.Scan()
	sale.Client = scanner.Text()
	fmt.Println("Въведете вид зърно.")
	sale.GrainType = scanGrainType(scanner)
	fmt.Println("Въведете количество.")
	sale.Quantity = int32(scanInt(scanner))
	fmt.Println("Въведете цена.")
	sale.Price = int32(scanInt(scanner))
	if err := stateStruct.db.CreateSale(context.Background(), sale); err != nil {
		fmt.Printf("Неуспешно създаване на документа - %v\n", err)
		return
	}
	fmt.Println("Документът е създаден успешно.")
	fmt.Println(refLineSeparator)
}

func NewTransport(scanner *bufio.Scanner, stateStruct *State) {
	var transport database.CreateTransportParams
	// select purchase
	purchases, err := stateStruct.db.GetAllPurchases(context.Background())
	if len(purchases) == 0 {
		fmt.Println("В момента няма активни договори за покупка. Моля, създайте нов договор.")
		return
	}
	if err != nil {
		fmt.Printf("Грешка при търсене на договори - %v\n", err)
		return
	}
	fmt.Println("Изберете номер на договoр за покупка. Активни договори към момента са:")
	printPurchases(stateStruct)
	for {
		purchase, err := stateStruct.db.GetPurchaseById(context.Background(), int32(scanInt(scanner)))
		if err != nil {
			fmt.Printf("Невалиден номер на договор опитайте пак - %v\n", err)
			continue
		}
		transport.PurchaseID = sql.NullInt32{
			Valid: true,
			Int32: purchase.ID,
		}
		transport.GrainType = purchase.GrainType
		break
	}
	// select sale
	sales, err := stateStruct.db.GetSalesByGrainType(context.Background(), transport.GrainType)
	if len(sales) == 0 {
		fmt.Println("В момента няма активни договори за продажба. Моля, създайте нов договор.")
		return
	}
	if err != nil {
		fmt.Printf("Грешка при търсене на договори - %v\n", err)
		return
	}
	fmt.Println("Изберете номер нa договoр за продажба. Активни договори към момента са:")
	printSalesByGrainType(stateStruct, transport.GrainType)
	for {
		sale, err := stateStruct.db.GetSaleByIdandGrainType(context.Background(), database.GetSaleByIdandGrainTypeParams{
			ID:        int32(scanInt(scanner)),
			GrainType: transport.GrainType,
		})
		if err != nil {
			fmt.Printf("Невалиден номер на договор опитайте пак - %v\n", err)
			continue
		}
		transport.SaleID = sql.NullInt32{
			Valid: true,
			Int32: sale.ID,
		}
		break
	}
	fmt.Println("Въведете номер на камион.")
	scanner.Scan()
	transport.TruckReg = scanner.Text()
	fmt.Println("Въведете номер на ремарке.")
	scanner.Scan()
	transport.TrailerReg = scanner.Text()
	fmt.Println("Въведете количество нето.")
	transport.Net = int32(scanInt(scanner))
	if err := stateStruct.db.CreateTransport(context.Background(), transport); err != nil {
		fmt.Printf("Неуспешно създаване на документа - %v\n", err)
		return
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
			fmt.Printf("Невалидно число: %v\n", err)
			continue
		}
		return num
	}
}

func SpamNewPurchase(scanner *bufio.Scanner, stateStruct *State) {
	fmt.Println("Колко документа да бъдат създадени?")
	reps := scanInt(scanner)
	for i := range reps {
		time.Sleep(time.Second)
		if err := stateStruct.db.CreatePurchase(context.Background(), database.CreatePurchaseParams{
			Suplier:   "Доставчик",
			Price:     300,
			Quantity:  100,
			GrainType: "пшеница",
		}); err != nil {
			fmt.Printf("Error creating document %d - %v\n", i, err)
			continue
		}
		fmt.Printf("Документ %d е създаден успешно.\n", i+1)
	}
}
