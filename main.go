package main

import (
	"fmt"
	"os"
)

func main() {
	var correctUsername string = "barmarket"
	var correctPassword string = "allien"
	fmt.Println("=============================")
	fmt.Println("Login")
	fmt.Println("=============================")
	var input string

	for input != correctUsername {
		fmt.Print("Enter Username: ")
		fmt.Scanln(&input)

		if input != correctUsername {
			fmt.Println("Username not found")
		}
	}

	for input != correctPassword {
		fmt.Print("Enter Password: ")
		fmt.Scanln(&input)

		if input != correctPassword {
			fmt.Println("Password not found")
		}

	}

	fmt.Println("Login Success")
	fmt.Println("=============================")
	fmt.Println("Welcome to Bar Market")
	for {
		var input int
		fmt.Println("Menu,masakan apa yang ingin anda pesan?")
		fmt.Println("1. Nasi Goreng")
		fmt.Println("2. Mie Goreng")
		fmt.Println("3. Sate Ayam")
		fmt.Println("4. Keluar")

		fmt.Print("Masukan pilihan anda: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			fmt.Println("Menu makanan yang anda pesan adalah Nasi Goreng")
		case 2:
			fmt.Println("Anda memesan Mie Goreng")
		case 3:
			fmt.Println("Anda memesan Sate Ayam")
		case 4:
			os.Exit(0)

		}
	}
}
