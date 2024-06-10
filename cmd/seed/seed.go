package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to db
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// insert to books
	sqls := []string{
		"INSERT INTO books (title, author, description) VALUES ('The Great Gatsby', 'F. Scott Fitzgerald', 'The Great Gatsby is a 1925 novel by American writer F. Scott Fitzgerald. Set in the Jazz Age on Long Island, near New York City, the novel depicts first-person narrator Nick Carraway''s interactions with mysterious millionaire Jay Gatsby and Gatsby''s obsession to reunite with his former lover, Daisy Buchanan.');",
		"INSERT INTO books (title, author, description) VALUES ('To Kill a Mockingbird', 'Harper Lee', 'To Kill a Mockingbird is a novel by Harper Lee published in 1960. Instantly successful, widely read in high schools and middle schools in the United States, it has become a classic of modern American literature, winning the Pulitzer Prize.');",
		"INSERT INTO books (title, author, description) VALUES ('1984', 'George Orwell', 'Nineteen Eighty-Four: A Novel, often referred to as 1984, is a dystopian social science fiction novel by English novelist George Orwell. It was published on 8 June 1949 by Secker & Warburg as Orwell''s ninth and final book completed in his lifetime.');",
		"INSERT INTO books (title, author, description) VALUES ('The Catcher in the Rye', 'J.D. Salinger', 'The Catcher in the Rye is a novel by J. D. Salinger, partially published in serial form in 1945–1946 and as a novel in 1951. It was originally intended for adults, but is often read by adolescents for its themes of angst and alienation, and as a critique on superficiality in society.');",
		"INSERT INTO books (title, author, description) VALUES ('Lord of the Flies', 'William Golding', 'Lord of the Flies is a 1954 novel by Nobel Prize-winning British author William Golding. The book focuses on a group of British');",
	}

	for _, s := range sqls {
		_, err := db.Exec(s)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("table books seeded")

	// insert to users
	sqls = []string{
		"INSERT INTO users (email, name, password) VALUES ('ann@test.com', 'Ann', 'password');",
		"INSERT INTO users (email, name, password) VALUES ('bob@test.com', 'Bob', 'password');",
		"INSERT INTO users (email, name, password) VALUES ('cat@test.com', 'Cat', 'password');",
		"INSERT INTO users (email, name, password) VALUES ('dan@test.com', 'Dan', 'password');",
		"INSERT INTO users (email, name, password) VALUES ('eli@test.com', 'Eli', 'password');",
	}

	for _, s := range sqls {
		_, err := db.Exec(s)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("table users seeded")

	// insert to orders
	sqls = []string{
		"INSERT INTO orders (user_id) VALUES (1);",
		"INSERT INTO orders (user_id) VALUES (2);",
		"INSERT INTO orders (user_id) VALUES (3);",
		"INSERT INTO orders (user_id) VALUES (4);",
		"INSERT INTO orders (user_id) VALUES (5);",
	}

	for _, s := range sqls {
		_, err := db.Exec(s)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("table orders seeded")

	// insert to order_books
	sqls = []string{
		"INSERT INTO order_books (order_id, book_id, quantity) VALUES (1, 1, 1);",
		"INSERT INTO order_books (order_id, book_id, quantity) VALUES (1, 2, 2);",
		"INSERT INTO order_books (order_id, book_id, quantity) VALUES (2, 3, 1);",
		"INSERT INTO order_books (order_id, book_id, quantity) VALUES (2, 4, 1);",
		"INSERT INTO order_books (order_id, book_id, quantity) VALUES (3, 5, 1);",
	}
	for _, s := range sqls {
		_, err := db.Exec(s)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("table order_books seeded")
}
