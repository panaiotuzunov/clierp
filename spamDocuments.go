package main

import (
	"bufio"
	"clierp/internal/database"
	"context"
	"fmt"
	"time"
)

func SpamNewPurchase(scanner *bufio.Scanner, stateStruct *State) {
	fmt.Println("Колко документа да бъдат създадени?")
	reps := scanInt(scanner)
	for i := range reps {
		time.Sleep(time.Second)
		if err := stateStruct.db.CreatePurchase(context.Background(), database.CreatePurchaseParams{
			Suplier:   "Доставчик",
			Price:     "300",
			Quantity:  "100",
			GrainType: "пшеница",
		}); err != nil {
			fmt.Printf("Error creating document %d - %v\n", i, err)
			continue
		}
		fmt.Printf("Документ %d е създаден успешно.\n", i+1)
	}
}
