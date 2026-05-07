package use

import (
	"fmt"
	"strconv"
	"time"
)

/*
Problem 2 · Library Management System
Design a library system where members can borrow and return books, with fine calculation for late returns.
Requirements

Library has a catalog of books — each book has title, author, ISBN, and availability status
Members can search books by title or author
A member can borrow a book — generates a borrow record with due date (14 days from borrow date)
A member can return a book — if returned late, fine is ₹5 per day overdue
A member can only borrow 3 books at a time
Library tracks all currently borrowed books
*/

// library
type Library struct {
	name    string
	books   []Book
	borrows map[string][]BorrowedRecord
}

func NewLibrary(name string) *Library {
	return &Library{name: name, borrows: make(map[string][]BorrowedRecord)}
}

func (l *Library) AddBook(book Book) {
	l.books = append(l.books, book)
}

func (l *Library) AddBorrowedBook(memid int, record *BorrowedRecord) {
	//insert into the slice with the memid as key
	str := strconv.Itoa(memid)
	if v := l.borrows[str]; v != nil {
		l.borrows[str] = append(l.borrows[str], *record)
	} else {
		l.borrows[str] = append(make([]BorrowedRecord, 0), *record)

	}
}

func (l *Library) GetName() string {
	return l.name
}

func (l *Library) GetBooks() []Book {
	return l.books
}

// fine calculator
func (l *Library) FineCalculator(b *BorrowedRecord) float64 {
	if time.Now().After(b.dueDate) { // only enters if TODAY is past the due date
		days := int(time.Since(b.dueDate).Hours() / 24)
		return float64(days * 5)
	}
	return 0.0
}

// borrow book
func (l *Library) Borrow(book *Book, member *Member) error {
	//check
	if member.borrowedBooks >= 3 {
		return fmt.Errorf("return the borrowed books first....")
	} else {
		r := NewBorrowedRecord(*member, *book, time.Now().Add(14*24*time.Hour)) //borrowed book
		l.AddBorrowedBook(member.id, r)
		book.availability = false
		member.borrowedBooks += 1
	}
	return nil
}

// Return book
func (l *Library) Return(book *Book, member *Member) float64 {
	//remove the book from map(borrowed records), reduce the borrowed books from member, calculate fine??
	//1. removed borrowed book
	//3.calculate fine
	//i have list of borrowed books, use that to calculate fine??
	str := strconv.Itoa(member.id)
	total := 0.0
	for i, v := range l.borrows[str] {
		if v.GetBook().isbn == book.isbn {
			total += l.FineCalculator(&v)
			l.borrows[str] = append(l.borrows[str][:i], l.borrows[str][i+1:]...)
			break
		}
	}
	//2. reduce borrowed books
	member.borrowedBooks -= 1
	//--- change book availability status
	book.availability = true
	return total
}

func (l *Library) SearchByTitle(title string) []Book {
	res := make([]Book, 0)
	for _, v := range l.books {
		if v.GetTitle() == title {
			res = append(res, v)
		}
	}
	return res
}

func (l *Library) SearchByAuthor(author string) []Book {
	res := make([]Book, 0)
	for _, v := range l.books {
		if v.GetAuthor() == author {
			res = append(res, v)
		}
	}
	return res
}

// books
type Book struct {
	title        string
	author       string
	isbn         string
	availability bool
}

func NewBook(title string, author string, isbn string, availability bool) *Book {
	return &Book{
		title:        title,
		author:       author,
		isbn:         isbn,
		availability: availability,
	}
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetIsbn() string {
	return b.isbn
}

func (b *Book) GetAvalilability() bool {
	return b.availability
}

// member
type Member struct {
	id            int
	name          string
	borrowedBooks int
}

func NewMember(id int, name string, borrowedBooks int) *Member {
	return &Member{id: id, name: name, borrowedBooks: borrowedBooks}
}

func (m *Member) GetId() int            { return m.id }
func (m *Member) GetName() string       { return m.name }
func (m *Member) GetBorrowedBooks() int { return m.borrowedBooks }

// borrowed record
type BorrowedRecord struct {
	member  Member
	book    Book
	dueDate time.Time
}

func NewBorrowedRecord(member Member, book Book, dueDate time.Time) *BorrowedRecord {
	return &BorrowedRecord{member: member, book: book, dueDate: dueDate}
}

func (b *BorrowedRecord) GetMember() *Member {
	return &b.member
}

func (b *BorrowedRecord) GetBook() *Book {
	return &b.book
}

func (b *BorrowedRecord) GetDueDate() time.Time {
	return b.dueDate
}
