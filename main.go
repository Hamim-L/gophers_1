package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	var duit int
	harga := 50000

	fmt.Println("Harga RTX:", harga)

	for {
		fmt.Print("Duit lu berapa? ")
		fmt.Scanln(&duit)

		if duit < harga {
			color.Red("Duit lu kurang cok, miskin bat awokawok")
			fmt.Println("Kurang segini:", harga-duit)
			fmt.Println()
			continue
		}

		color.Green("Duit lu cukup nih")
		fmt.Println("Sisa:", duit-harga)
		break
	}
}