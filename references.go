package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
)

func printEntranceAndExitReciepts(stateStruct *State) {
	receipts, err := stateStruct.db.GetAllReceipts(context.Background())
	if err != nil {
		fmt.Println("Грешка при търсене на документи.")
		return
	}
	if len(receipts) == 0 {
		fmt.Println(refLineSeparator)
		fmt.Println("Няма намерени документи.")
		fmt.Println(refLineSeparator)
		return
	}
	fmt.Println(refLineSeparator)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ТИП\tНОМЕР\tДАТА\tЗЪРНО\tКАМИОН\tРЕМАРКЕ\tБРУТО\tТАРА\tНЕТО")
	for _, r := range receipts {
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%s\t%d\t%d\t%d\n",
			r.DocType, r.ID, r.CreatedAt.Format("02/01/2006"),
			r.GrainType, r.TruckReg, r.TrailerReg, r.Gross, r.Tare, r.Net)
	}
	w.Flush()
	fmt.Println(refLineSeparator)
}

func printInventory(stateStruct *State) {
	inventory, err := stateStruct.db.GetCurrentInventoryByType(context.Background())
	if err != nil {
		fmt.Println("Грешка при калкулиране на наличност.")
	}
	if len(inventory) == 0 {
		fmt.Println("Няма текуща наличност.")
		return
	}
	fmt.Println(refLineSeparator)
	fmt.Println("Текущата наличност по култури е:")
	for _, item := range inventory {
		fmt.Printf("%s - %d т.\n", item.GrainType, item.Sum)
	}
	fmt.Println(refLineSeparator)
}
