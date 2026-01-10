package main

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

func printEntranceAndExitReciepts(stateStruct *State) {
	receipts, err := stateStruct.db.GetAllReceipts(context.Background())
	if err != nil {
		fmt.Printf("Грешка при търсене на документи - %v\n", err)
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
	fmt.Fprintln(w, "ТИП\tНОМЕР\tДАТА\tЗЪРНО\tКАМИОН\tРЕМАРКЕ\tБРУТО\tТАРА\tНЕТО\tДОСТАВЧИК\tПОКУПКА №\tКЛИЕНТ\tПРОДАЖБА №")
	for _, r := range receipts {
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%s\t%d\t%d\t%d\t%s\t%v\t%s\t%v\n",
			r.DocType,
			r.ID,
			r.CreatedAt.Format("02/01/2006"),
			r.GrainType,
			r.TruckReg,
			r.TrailerReg,
			r.Gross, r.Tare,
			r.Net,
			r.Suplier.String,
			nullIntToStr(r.PurchaseID),
			r.Client.String,
			nullIntToStr(r.SaleID))
	}
	w.Flush()
	fmt.Println(refLineSeparator)
}

func printInventory(stateStruct *State) {
	inventory, err := stateStruct.db.GetCurrentInventoryByType(context.Background())
	if err != nil {
		fmt.Printf("Грешка при калкулиране на наличност - %v", err)
		return
	}
	if len(inventory) == 0 {
		fmt.Println(refLineSeparator)
		fmt.Println("Няма текуща наличност.")
		fmt.Println(refLineSeparator)
		return
	}
	fmt.Println(refLineSeparator)
	fmt.Println("Текущата наличност по култури е:")
	for _, item := range inventory {
		fmt.Printf("%s - %d т.\n", item.GrainType, item.Sum)
	}
	fmt.Println(refLineSeparator)
}

func printPurchases(stateStruct *State) {
	purchases, err := stateStruct.db.GetAllPurchases(context.Background())
	if err != nil {
		fmt.Printf("Грешка при търсене на документи - %v\n", err)
		return
	}
	if len(purchases) == 0 {
		fmt.Println(refLineSeparator)
		fmt.Println("Няма намерени документи.")
		fmt.Println(refLineSeparator)
		return
	}
	fmt.Println(refLineSeparator)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "НОМЕР\tДАТА\tДОСТАВЧИК\tЗЪРНО\tЦЕНА\tКОЛИЧЕСТВО\tЕКСПЕДИРАНО\tОСТАТЪК")
	for _, p := range purchases {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%d\t%d\t%d\t%d\n",
			p.ID,
			p.CreatedAt.Format("02/01/2006"),
			p.Suplier,
			p.GrainType,
			p.Price,
			p.Quantity,
			p.Expedited,
			p.Quantity-p.Expedited)
	}
	w.Flush()
	fmt.Println(refLineSeparator)
}

func printSales(stateStruct *State) {
	sales, err := stateStruct.db.GetAllSales(context.Background())
	if err != nil {
		fmt.Printf("Грешка при търсене на документи - %v\n", err)
		return
	}
	if len(sales) == 0 {
		fmt.Println(refLineSeparator)
		fmt.Println("Няма намерени документи.")
		fmt.Println(refLineSeparator)
		return
	}
	fmt.Println(refLineSeparator)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "НОМЕР\tДАТА\tКЛИЕНТ\tЗЪРНО\tЦЕНА\tКОЛИЧЕСТВО\tЕКСПЕДИРАНО\tОСТАТЪК")
	for _, s := range sales {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%d\t%d\t%d\t%d\n",
			s.ID,
			s.CreatedAt.Format("02/01/2006"),
			s.Client,
			s.GrainType,
			s.Price,
			s.Quantity,
			int32(math.Abs(float64(s.Expedited))),
			s.Quantity-int32(math.Abs(float64(s.Expedited))))
	}
	w.Flush()
	fmt.Println(refLineSeparator)
}

func printTransports(stateStruct *State) {
	transports, err := stateStruct.db.GetAllTransports(context.Background())
	if err != nil {
		fmt.Printf("Грешка при търсене на документи - %v\n", err)
		return
	}
	if len(transports) == 0 {
		fmt.Println(refLineSeparator)
		fmt.Println("Няма намерени документи.")
		fmt.Println(refLineSeparator)
		return
	}
	fmt.Println(refLineSeparator)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "НОМЕР\tДАТА\tКАМИОН\tРЕМАРКЕ\tКОЛИЧЕСТВО\tЗЪРНО\tДОСТАВЧИК\tПОКУПКА №\tКЛИЕНТ\tПРОДАЖБА №")
	for _, t := range transports {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%d\t%s\t%s\t%d\t%s\t%d\n",
			t.ID,
			t.CreatedAt.Format("02/01/2006"),
			t.TruckReg,
			t.TrailerReg,
			t.Net,
			t.GrainType,
			t.Suplier.String,
			t.PurchaseID.Int32,
			t.Client.String,
			t.SaleID.Int32)
	}
	w.Flush()
	fmt.Println(refLineSeparator)
}

func nullIntToStr(n sql.NullInt32) any {
	if !n.Valid {
		return ""
	}
	return n.Int32
}
