package main

import (
	"bufio"
	"context"
	"database/sql"
	"errors"
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
		scanner.Scan()
		selection := scanner.Text()
		switch selection {
		case "1":
			printEntranceAndExitReciеpts(stateStruct)
		case "0":
			break outer
		case "exit":
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
		scanner.Scan()
		selection := scanner.Text()
		switch selection {
		case "1":
			NewEntranceReceipt(stateStruct)
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

func printEntranceAndExitReciеpts(stateStruct *State) {
	receipts, err := stateStruct.db.GetAllReceipts(context.Background())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("Няма намерени документи.")
			return
		}
		fmt.Println("Грешка при търсене на документи.")
		return
	}
	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Printf("%-12v %-12v %-12v %-12v %-12v %-12v %-12v\n", "НОМЕР", "ДАТА", "КАМИОН", "РЕМАРКЕ", "БРУТО", "ТАРА", "НЕТО")
	for _, r := range receipts {
		fmt.Printf("%-12v %-12v %-12v %-12v %-12v %-12v %-12v\n",
			r.ID,
			r.CreatedAt.Format("02/01/2006"),
			r.TruckReg, r.TrailerReg, r.Gross, r.Tare, r.Net)
	}
	fmt.Println("----------------------------------------------------------------------------------------")
}
