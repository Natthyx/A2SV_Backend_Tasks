package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func RunLibrarySystem(library services.LibraryManager) {
	for {
		fmt.Println("\n--- Library Menu ---")
		fmt.Println("1. Add book")
		fmt.Println("2. Remove book")
		fmt.Println("3. Borrow book")
		fmt.Println("4. Return book")
		fmt.Println("5. List available books")
		fmt.Println("6. List borrowed books by member")
		fmt.Println("7. Add new member")
		fmt.Println("8. List members")
		fmt.Println("9. Exit")

		fmt.Print("Enter choice: ")
		choice := readInput()

		switch choice {
		case "1":
			fmt.Print("Enter title: ")
			title := readInput()
			fmt.Print("Enter author: ")
			author := readInput()
			book := models.Book{Title: title, Author: author}
			library.AddBook(book)

		case "2":
			fmt.Print("Enter book ID: ")
			bookID, _ := strconv.Atoi(readInput())
			library.RemoveBook(bookID)

		case "3":
			fmt.Print("Enter book ID: ")
			bookID, _ := strconv.Atoi(readInput())
			fmt.Print("Enter member ID: ")
			memberID, _ := strconv.Atoi(readInput())
			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case "4":
			fmt.Print("Enter book ID: ")
			bookID, _ := strconv.Atoi(readInput())
			fmt.Print("Enter member ID: ")
			memberID, _ := strconv.Atoi(readInput())
			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case "5":
			books := library.ListAvailableBooks()
			fmt.Println("Available Books:")
			for _, book := range books {
				fmt.Printf("ID: %d | %s by %s\n", book.ID, book.Title, book.Author)
			}

		case "6":
			fmt.Print("Enter member ID: ")
			memberID, _ := strconv.Atoi(readInput())
			books := library.ListBorrowedBooks(memberID)
			fmt.Println("Borrowed Books:")
			for _, book := range books {
				fmt.Printf("ID: %d | %s by %s\n", book.ID, book.Title, book.Author)
			}

		case "7":
			fmt.Print("Enter member name: ")
			name := readInput()
			member := models.Member{Name: name}
			library.AddMember(member)

		case "8":
			members := library.ListMembers()
			fmt.Println("Registered Members:")
			for _, m := range members {
				fmt.Printf("ID: %d | Name: %s\n", m.ID, m.Name)
			}

		case "9":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
