package controllers

import (
	"bufio"
	"fmt"
	"libary-management/models"
	"libary-management/services"
	"libary-management/types"
	"os"
	"strconv"
	"strings"
)

func Start(library services.LibraryManager) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Welcome to the Golang Library Management System 📚")
		fmt.Println("Here is a list of actions you can perform")
		fmt.Println("1. Add a book")
		fmt.Println("2. Remove a book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return Borrowed book")
		fmt.Println("5. List all books")
		fmt.Println("6. List all borrowed books")
		fmt.Println("7. Exit Application")
		fmt.Println("Start typing below ...")
		fmt.Println()

		scanner.Scan()
		userChoice := scanner.Text()

		switch userChoice {
		case "1":
			addBook(library)
		case "2":
			removeBook(library, scanner)
		case "3":
			borrowBook(library, scanner)
		case "4":
			returnBook(library, scanner)
		case "5":
			listBooks(library)
		case "6":
			listBorrowedBooks(library, scanner)
		case "7":
			return
		default:
			println("input is invalid")
		}
	}

}

var BookId = 1

func addBook(library services.LibraryManager) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter book Title: ")
	title, _ := reader.ReadString('\n')

	fmt.Println("Enter book Author: ")
	author, _ := reader.ReadString('\n')

	book := models.Book{
		Id:     BookId,
		Title:  strings.TrimSpace(title),
		Author: strings.TrimSpace(author),
		Status: types.Available,
	}

	BookId++

	library.AddBook(book)
	fmt.Println("Book added successfully")
}

func removeBook(library services.LibraryManager, scanner *bufio.Scanner) {
	fmt.Println("Enter book Id: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	library.RemoveBook(id)
	fmt.Println("Book removed successfully")
}

func borrowBook(library services.LibraryManager, scanner *bufio.Scanner) {
	fmt.Println("Enter book Id: ")
	scanner.Scan()
	bookId, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter member Id: ")
	scanner.Scan()
	memberId, _ := strconv.Atoi(scanner.Text())

	err := library.BorrowBook(bookId, memberId)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Book with id %d borrowed successfully\n", bookId)
	}
}

func returnBook(library services.LibraryManager, scanner *bufio.Scanner) {
	fmt.Println("Enter book Id: ")
	scanner.Scan()
	bookId, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter member Id: ")
	scanner.Scan()
	memberId, _ := strconv.Atoi(scanner.Text())

	err := library.ReturnBook(bookId, memberId)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("Book with id %d returned successfully\n", bookId)
	}
}

func listBooks(library services.LibraryManager) {
	books := library.ListAvailableBooks()

	if len(books) == 0 {
		fmt.Println("No books found")
	}

	for i, book := range books {
		fmt.Printf("%d:\nBook ID: %d\nBook Title: %s\nBook Author: %s\nBook Status: %v\n", i+1, book.Id, book.Title, book.Author, book.Status)
	}
}

func listBorrowedBooks(library services.LibraryManager, scanner *bufio.Scanner) {
	fmt.Println("Enter Member Id: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	books := library.ListBorrowedBooks(id)

	if len(books) == 0 {
		fmt.Println("No Borrowed books found")
	}

	for i, book := range books {
		fmt.Printf("%d:\nBook ID: %d\nBook Title: %s\nBook Author: %s\nBook Status: %v\n", i+1, book.Id, book.Title, book.Author, book.Status)
	}
}
