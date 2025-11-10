package main

import (
	"libary-management/controllers"
	"libary-management/models"
	"libary-management/services"
)

func main() {
	library := services.GolangLibrary()

	//Add a sample member
	library.Members[1] = models.Member{
		Id:            1,
		Name:          "bolu golang",
		BorrowedBooks: []models.Book{},
	}

	//start the program
	controllers.Start(library)
}
