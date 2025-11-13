package concurrency

import (
	"fmt"
	"libary-management/services"
	"sync"
)

type ReservationRequest struct {
	BookID   int
	MemberID int
}

func Worker(id int, library *services.Library, requests <-chan ReservationRequest, wg *sync.WaitGroup) {
	fmt.Printf("worker %d started.\n", id)

	for request := range requests {
		fmt.Printf("worker %d processing Book %d for Member %d \n", id, request.BookID, request.MemberID)

		err := library.ReserveBook(request.BookID, request.MemberID)

		if err != nil {
			fmt.Printf("worker %d failed to reserve Book %d for Member %d: %v \n", id, request.BookID, request.MemberID, err)
		}

		wg.Done()
	}
	fmt.Printf("worker %d finished \n", id)
}
