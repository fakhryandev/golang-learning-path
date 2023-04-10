package main

import "fmt"

func main() {
	var i int = 21
	var j = true
	k := 123.456

	ya := 'Я'

	fmt.Printf("Menampilkan nilai i : %v\n", i)
	fmt.Printf("Menampilkan tipe data dari variable i : %T\n", i)
	fmt.Printf("Menampilkan tanda : %% \n")
	fmt.Printf("Menampilkan nilai j : %v\n", j)
	fmt.Printf("Menampilkan nilai kebalikan j : %v\n", !j)
	fmt.Printf("Menampilkan unicode russia : %s\n", "\u042F (ya)")
	fmt.Printf("Menampilkan nilai base 10 : %d \n", 21)
	fmt.Printf("Menampilkan nilai base 8 : %o \n", 21)
	fmt.Printf("Menampilkan nilai base 16 : %x \n", 15)
	fmt.Printf("Menampilkan nilai base 16 : %X \n", 15)
	fmt.Printf("Menampilkan unicode karakter Я : %U\n", ya)
	fmt.Printf("Menampilkan nilai float : %f \n", k)
	fmt.Printf("Menampilkan nilai scientific : %E \n", k)
}
