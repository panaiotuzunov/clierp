package main

import (
	"bufio"
	"clierp/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

func SpamNewPurchase(scanner *bufio.Scanner, stateStruct *State) {
	fmt.Println("Колко документа да бъдат създадени?")
	reps := scanInt(scanner)
	for i := range reps {
		time.Sleep(time.Second)
		if err := stateStruct.db.CreatePurchase(context.Background(), database.CreatePurchaseParams{
			Suplier:   "Доставчик",
			Price:     decimal.NewFromInt(300),
			Quantity:  decimal.NewFromInt(100),
			GrainType: "пшеница",
		}); err != nil {
			fmt.Printf("Error creating document %d - %v\n", i, err)
			continue
		}
		fmt.Printf("Документ %d е създаден успешно.\n", i+1)
	}
}
