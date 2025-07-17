package services

import (
	"errors"
	"fmt"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	AddMember(member models.Member) int
	ListMembers() []models.Member
}

type Library struct {
	Books         map[int]models.Book
	Members       map[int]models.Member
	BorrowedBooks map[int]int // bookID -> memberID
	nextBookID    int
	nextMemberID  int
}

func NewLibrary() *Library {
	return &Library{
		Books:         make(map[int]models.Book),
		Members:       make(map[int]models.Member),
		BorrowedBooks: make(map[int]int),
		nextBookID:    1,
		nextMemberID:  1,
	}
}

func (l *Library) AddBook(book models.Book) {
	book.ID = l.nextBookID
	l.nextBookID++
	l.Books[book.ID] = book
	fmt.Println("Book added successfully.")
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
	delete(l.BorrowedBooks, bookID)
	fmt.Println("Book removed.")
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	if _, ok := l.Members[memberID]; !ok {
		return errors.New("member does not exist")
	}
	if _, ok := l.Books[bookID]; !ok {
		return errors.New("book does not exist")
	}
	if _, borrowed := l.BorrowedBooks[bookID]; borrowed {
		return errors.New("book is already borrowed")
	}
	l.BorrowedBooks[bookID] = memberID
	fmt.Println("Book borrowed successfully.")
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	if borrower, ok := l.BorrowedBooks[bookID]; !ok || borrower != memberID {
		return errors.New("this book was not borrowed by this member")
	}
	delete(l.BorrowedBooks, bookID)
	fmt.Println("Book returned successfully.")
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	available := []models.Book{}
	for id, book := range l.Books {
		if _, borrowed := l.BorrowedBooks[id]; !borrowed {
			available = append(available, book)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	borrowed := []models.Book{}
	for bookID, borrowerID := range l.BorrowedBooks {
		if borrowerID == memberID {
			if book, ok := l.Books[bookID]; ok {
				borrowed = append(borrowed, book)
			}
		}
	}
	return borrowed
}

func (l *Library) AddMember(member models.Member) int {
	member.ID = l.nextMemberID
	l.nextMemberID++
	l.Members[member.ID] = member
	fmt.Printf("Member added successfully. ID: %d\n", member.ID)
	return member.ID
}

func (l *Library) ListMembers() []models.Member {
	members := []models.Member{}
	for _, member := range l.Members {
		members = append(members, member)
	}
	return members
}
