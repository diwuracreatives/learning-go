package main

import (
	"fmt"
	"libary-management/models"
	"libary-management/services"
	"libary-management/services/concurrency"
	"libary-management/types"
	"sync"
	"time"
)

func main() {
	library := services.GolangLibrary()

	// seed multiple books and members
	book1 := models.Book{Id: 1, Title: "Adventure of Code of Conduct", Author: "ray lee", Status: types.Borrowed}
	book2 := models.Book{Id: 2, Title: "The Go Routine Guide", Author: "Max Concurrency", Status: types.Borrowed}

	library.Books[1] = models.Book{Id: 1, Title: "Adventure of Code of Conduct", Author: "ray lee", Status: types.Borrowed}
	library.Books[2] = models.Book{Id: 2, Title: "The Go Routine Guide", Author: "Max Concurrency", Status: types.Borrowed}
	library.Books[3] = models.Book{Id: 3, Title: "Mutex Mastery", Author: "Lana Sync", Status: types.Available}
	library.Books[4] = models.Book{Id: 4, Title: "Channel Theory", Author: "Igor Flow", Status: types.Available}
	library.Books[5] = models.Book{Id: 5, Title: "The Art of Defer", Author: "Ali Grace", Status: types.Available}

	library.Members[1] = models.Member{Id: 1, Name: "Noah Thomas", BorrowedBooks: []models.Book{book1, book2}}
	library.Members[2] = models.Member{Id: 2, Name: "Lola Ade", BorrowedBooks: []models.Book{}}
	library.Members[3] = models.Member{Id: 3, Name: "Kenneth Olaoluwa", BorrowedBooks: []models.Book{}}
	library.Members[4] = models.Member{Id: 4, Name: "Thomas Aina", BorrowedBooks: []models.Book{}}
	library.Members[5] = models.Member{Id: 5, Name: "Gabriel Mary", BorrowedBooks: []models.Book{}}

	const (
		NumWorkers    = 3
		TotalRequests = 5
	)

	requests := make(chan concurrency.ReservationRequest, TotalRequests)

	var wg sync.WaitGroup

	for i := 1; i <= NumWorkers; i++ {
		go concurrency.Worker(i, library, requests, &wg)
	}

	fmt.Println("sending reservation requests...")

	for i := 1; i <= TotalRequests; i++ {
		bookID := i
		memberID := i

		wg.Add(1)
		requests <- concurrency.ReservationRequest{BookID: bookID, MemberID: memberID}
	}

	close(requests)

	wg.Wait()
	fmt.Println("initial reservations processed by workers.")

	fmt.Println("still Processing reservation requests")
	time.Sleep(10 * time.Second)
}

//start the program
//controllers.Start(library)
