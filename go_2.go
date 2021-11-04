package main

import (
	"fmt"
	"math/rand"
	"time"
)

var NumberOfServices = 15

type Book struct {
	name string
	author string
}

type Reader struct {
	name string
}

type Librarian struct { }

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (reader Reader) takeBook(booksChan chan Book, readerSemaphore chan int,
	librarianSemaphore chan bool) {
	for {
		if <-readerSemaphore == -1 {
			return
		}
		firstBook := <-booksChan
		secondBook := <-booksChan
		fmt.Println("Reader "+reader.name+" got books:",
			firstBook, ", ", secondBook)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Reader " + reader.name + " read books.")
		fmt.Println()
		librarianSemaphore <- true
	}
}

func (Librarian) giveBooks(partsChan chan Book, readerSemaphore chan int,
	librarianSemaphore chan bool, endChan chan bool) {
	for i := 0; i < NumberOfServices; i++ {
		if i != 0 {
			<-librarianSemaphore
		}
		var firstBook = Book{name: randSeq(10), author: randSeq(8)}
		var secondBook = Book{name: randSeq(10), author: randSeq(8)}
		fmt.Println("Librarian sent ", firstBook, " and ",
			secondBook)
		partsChan <- firstBook
		partsChan <- secondBook
		readerSemaphore <- 1
	}
	<-librarianSemaphore
	readerSemaphore <- -1
	readerSemaphore <- -1
	readerSemaphore <- -1
	endChan <- true
}

func main() {
	reader1 := Reader{name: "Alex"}
	reader2 := Reader{name: "Dasha"}
	reader3 := Reader{name: "Andrew"}

	booksChan := make(chan Book, 2)
	librarianSemaphore := make(chan bool, 1)
	readerSemaphore := make(chan int, 1)
	endChan := make(chan bool, 1)

	go Librarian{}.giveBooks(booksChan, readerSemaphore, librarianSemaphore, endChan)
	go reader1.takeBook(booksChan, readerSemaphore, librarianSemaphore)
	go reader2.takeBook(booksChan, readerSemaphore, librarianSemaphore)
	go reader3.takeBook(booksChan, readerSemaphore, librarianSemaphore)
	<-endChan
}
