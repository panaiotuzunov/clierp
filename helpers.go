package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"strconv"
)

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

func scanFloat64(scanner *bufio.Scanner) float64 {
	for {
		scanner.Scan()
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Printf("Невалидно число: %v\n", err)
			continue
		}
		return num
	}
}

func float64toStrForDB(num float64) string {
	return strconv.FormatFloat(num, 'f', 3, 64)
}

func nullIntToStr(n sql.NullInt32) any {
	if !n.Valid {
		return ""
	}
	return n.Int32
}
