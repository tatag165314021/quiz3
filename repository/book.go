package repository

import (
	"database/sql"
	"quiz3/structs"
)

func GetAllBooks(db *sql.DB) (results []structs.Books, err error) {
	sql := "SELECT * FROM books"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var books = structs.Books{}
		err = rows.Scan(&books.ID, &books.Title, &books.Description, &books.ImageURL, &books.ReleaseYear, &books.Price, &books.TotalPage, &books.Thickness, &books.Created_at, &books.Updated_at, &books.CategoryID)
		if err != nil {
			panic(err)
		}
		results = append(results, books)
	}
	return
}

func InsertBooks(db *sql.DB, books structs.Books) (err error) {
	sql := "INSERT INTO books ( title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id) VALUES ($1,$2,$3,$4, $5, $6, $7,$8,$9, $10)"
	errs := db.QueryRow(sql, books.Title, books.Description, books.ImageURL, books.ReleaseYear, books.Price, books.TotalPage, books.Thickness, books.Created_at, books.Updated_at, books.CategoryID)
	return errs.Err()
}

func UpdateBooks(db *sql.DB, books structs.Books) (err error) {
	sql := "UPDATE books SET title=$1, description=$2, image_url=$3, release_year=$4, price=$5, total_page=$6, thickness=$7, created_at=$8, updated_at=$9, category_id=$10 WHERE id = $11"
	errs := db.QueryRow(sql, books.Title, books.Description, books.ImageURL, books.ReleaseYear, books.Price, books.TotalPage, books.Thickness, books.Created_at, books.Updated_at, books.CategoryID, books.ID)
	return errs.Err()
}

func DeleteBooks(db *sql.DB, books structs.Books) (err error) {
	sql := "DELETE FROM books WHERE id = $1"
	errs := db.QueryRow(sql, books.ID)

	return errs.Err()
}

func GetBooksCategory(db *sql.DB, books structs.Books) (results []structs.Books, err error) {
	sql := "SELECT * FROM books where category_id = $1"
	rows, err := db.Query(sql, books.CategoryID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var books = structs.Books{}
		err = rows.Scan(&books.ID, &books.Title, &books.Description, &books.ImageURL, &books.ReleaseYear, &books.Price, &books.TotalPage, &books.Thickness, &books.Created_at, &books.Updated_at, &books.CategoryID)
		if err != nil {
			panic(err)
		}
		results = append(results, books)
	}
	return
}
