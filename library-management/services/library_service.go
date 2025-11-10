package services

import (
	"fmt"
	"libary-management/models"
	"libary-management/types"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookId int)
	BorrowBook(bookId int, memberId int) error
	ReturnBook(bookId int, memberId int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberId int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func GolangLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
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
	book, err := library.GetBook(bookId)
	if err != nil {
		return err
	}

	if book.Status == types.Available {
		member, err := library.GetMember(memberId)
		if err != nil {
			return err
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
