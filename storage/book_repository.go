package storage

import (
	"database/sql"
	"fmt"

	"my-book-app/internal/book"
)

type PostgresBookRepository struct {
	db *sql.DB
}

func NewPostgresBookRepository(db *sql.DB) *PostgresBookRepository {
    return &PostgresBookRepository{
        db: db,
    }
}

func (r *PostgresBookRepository) Create(book *book.Book) error {
	_, err := r.db.Exec("INSERT INTO books (id, title, author, genre, year) VALUES ($1, $2, $3, $4, $5)",
		book.ID, book.Title, book.Author, book.Genre, book.Year)
	return err
}

func (r *PostgresBookRepository) Read(id int) (*book.Book, error) {
	row := r.db.QueryRow("SELECT id, title, author, genre, year FROM books WHERE id = $1", id)
	book := &book.Book{}
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre, &book.Year)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("book not found")
	} else if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *PostgresBookRepository) Update(book *book.Book) error {
	_, err := r.db.Exec("UPDATE books SET title = $1, author = $2, genre = $3, year = $4 WHERE id = $5",
		book.Title, book.Author, book.Genre, book.Year, book.ID)
	return err
}

func (r *PostgresBookRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id = $1", id)
	return err
}
