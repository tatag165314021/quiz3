package repository

import (
	"database/sql"
	"quiz3/structs"
)

func GetAllCategory(db *sql.DB) (results []structs.Category, err error) {
	sql := "SELECT * FROM category"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = structs.Category{}
		err = rows.Scan(&category.ID, &category.Name, &category.Created_at, &category.Updated_at)
		if err != nil {
			panic(err)
		}
		results = append(results, category)
	}
	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO category (id, name, created_at, updated_at) VALUES ($1,$2,$3,$4)"
	errs := db.QueryRow(sql, category.ID, category.Name, category.Created_at, category.Updated_at)
	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE category SET name = $1, created_at = $2, updated_at = $3 WHERE id = $4"
	errs := db.QueryRow(sql, category.Name, category.Created_at, category.Updated_at, category.ID)
	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM category WHERE id = $1"
	errs := db.QueryRow(sql, category.ID)

	return errs.Err()
}
