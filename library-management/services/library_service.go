package services

import (
	"fmt"
	"libary-management/models"
	"libary-management/types"
	"sync"
	"time"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookId int)
	BorrowBook(bookId int, memberId int) error
	ReturnBook(bookId int, memberId int) error
	ReserveBook(bookId int, memberId int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberId int) []models.Book
}

type Library struct {
	Books        map[int]models.Book
	Members      map[int]models.Member
	Reservations map[int]models.Reservation
	Mutex        sync.Mutex
}

func GolangLibrary() *Library {
	return &Library{
		Books:        make(map[int]models.Book),
		Members:      make(map[int]models.Member),
		Reservations: make(map[int]models.Reservation),
	}
}

func (library *Library) GetBook(bookId int) (models.Book, error) {
	book, bookExists := library.Books[bookId]
	if !bookExists {
		return models.Book{}, fmt.Errorf("a book with id %d does not exist", bookId)
	}
	return book, nil
}

func (library *Library) GetMember(memberId int) (models.Member, error) {
	member, memberExists := library.Members[memberId]
	if !memberExists {
		return models.Member{}, fmt.Errorf("a member with id %d does not exist", memberId)
	}
	return member, nil
}

func (library *Library) RemoveBorrowedBook(bookId int, member models.Member) models.Member {
	for i, book := range member.BorrowedBooks {
		if book.Id == bookId {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			break
		}
	}
	return member
}

func (library *Library) AddBook(book models.Book) {
	library.Books[book.Id] = book
}

func (library *Library) RemoveBook(bookId int) {
	delete(library.Books, bookId)
}

func (library *Library) BorrowBook(bookId int, memberId int) error {
	book, bookExists := library.GetBook(bookId)
	if bookExists != nil {
		return bookExists
	}

	if book.Status == types.Available {
		member, memberExists := library.GetMember(memberId)
		if memberExists != nil {
			return memberExists
		}

		// add book to member borrowed book list
		member.BorrowedBooks = append(member.BorrowedBooks, book)
		library.Members[memberId] = member

		// change book status to borrowed
		book.Status = types.Borrowed
		library.Books[bookId] = book
	}

	return nil
}

func (library *Library) ReturnBook(bookId int, memberId int) error {
	book, err := library.GetBook(bookId)
	if err != nil {
		return err
	}

	if book.Status == types.Borrowed {
		member, err := library.GetMember(memberId)
		if err != nil {
			return err
		}

		// remove book from member borrowed book list
		member = library.RemoveBorrowedBook(bookId, member)
		library.Members[memberId] = member

		// change book status to available
		book.Status = types.Available
		library.Books[bookId] = book
	}
	return nil
}

func (library *Library) ListAvailableBooks() []models.Book {
	var books []models.Book
	for _, book := range library.Books {
		books = append(books, book)
	}
	return books
}

func (library *Library) ListBorrowedBooks(memberId int) []models.Book {
	var books []models.Book

	member, err := library.GetMember(memberId)
	if err != nil {
		return []models.Book{}
	}

	for _, book := range member.BorrowedBooks {
		books = append(books, book)
	}
	return books
}

func IsReservationExpired(bookingTime time.Time) bool {
	return time.Now().After(bookingTime.Add(5 * time.Second))
}

func (library *Library) GetReservation(bookId int) (models.Reservation, error) {
	reservation, reservationExists := library.Reservations[bookId]
	if !reservationExists {
		return models.Reservation{}, fmt.Errorf("a reservation with book id %d does not exist", bookId)
	}
	return reservation, nil
}

func (library *Library) ReserveBook(bookId int, memberId int) error {
	library.Mutex.Lock()
	book, getBookErr := library.GetBook(bookId)
	if getBookErr != nil {
		library.Mutex.Unlock()
		return getBookErr
	}

	if book.Status == types.Borrowed {
		library.Mutex.Unlock()
		return fmt.Errorf("a book with id %d is already borrowed", bookId)
	}

	bookReservation, reservationExists := library.Reservations[bookId]

	if reservationExists && !IsReservationExpired(bookReservation.BookingTime) {
		library.Mutex.Unlock()
		return fmt.Errorf("a book with id %d is already reserved", bookId)
	}

	if reservationExists && IsReservationExpired(bookReservation.BookingTime) {
		library.RemoveReservation(bookId)
	}

	reservation := models.Reservation{
		BookId:      bookId,
		MemberId:    memberId,
		BookingTime: time.Now(),
	}

	library.Reservations[bookId] = reservation

	timer := time.AfterFunc(5*time.Second, func() {
		library.Mutex.Lock()
		defer library.Mutex.Unlock()

		currentReservation, currReservationExists := library.Reservations[bookId]

		if currReservationExists &&
			currentReservation.MemberId == memberId &&
			currentReservation.BookingTime.Equal(reservation.BookingTime) {

			currBook, err := library.GetBook(bookId)
			if err == nil && currBook.Status != types.Borrowed {
				library.RemoveReservation(bookId)
				fmt.Printf("Auto-cancelled reservation for book %d (member %d)", bookId, memberId)
			}
		}
	})

	library.Mutex.Unlock()

	go func(bookId int, memberId int, timer *time.Timer) {
		err := library.BorrowBook(bookId, memberId)
		if err != nil {
			fmt.Printf("failed to borrow book with book id %d with member id %d: %v", bookId, memberId, err)

			library.RemoveReservation(bookId)
		}
		timer.Stop()
		fmt.Printf("book with id %d borrowed async successfully \n", bookId)
	}(bookId, memberId, timer)
	return nil
}

func (library *Library) RemoveReservation(bookId int) {
	library.Mutex.Lock()
	defer library.Mutex.Unlock()
	delete(library.Reservations, bookId)
}
