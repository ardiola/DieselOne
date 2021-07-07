package main

import "fmt"

func main() {
	soal1(10)
	fmt.Println()
	soal2()
	fmt.Println()
	angkaA := 17
	angkaB := 8
	soal3(angkaA, angkaB)
}

func soal1(angka int) {

	var angkaModify = (angka * 2) + 1

	for i := 1; i <= angkaModify; i++ {
		// fmt.Printf("X ")
		for j := 1; j <= angkaModify; j++ {
			if (i == 1 || j == 1) || (i == angkaModify || j == angkaModify) || (i == angka+1 || j == angka+1) || (j == (angkaModify-i+1) || i == j) {
				fmt.Print("X ")
			} else {
				fmt.Print("  ")
			}

		}
		fmt.Println()
	}

}

func soal2() {
	fmt.Println("soal 2")
}

func soal3(a int, b int) {

	jumlah := a + b

	for i := 1; i <= jumlah; i++ {

		if i%3 == 0 {
			fmt.Printf("b")
		} else {
			fmt.Printf("a")
		}

	}
}
