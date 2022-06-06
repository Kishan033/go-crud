package userdb

import (
	"database/sql"
	"fmt"
	"go-postgres/models"
	"log"
	"time"
)

type repo struct {
	client *sql.DB
}

func NewRepo(client *sql.DB) *repo {
	return &repo{
		client: client,
	}
}

func (repo *repo) Add(user models.User) int64 {
	sqlStatement := `INSERT INTO users (name, location, age, email, password, createdat, updatedat) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING userid`
	var id int64

	err := repo.client.QueryRow(sqlStatement, user.Name, user.Location, user.Age, user.Email, user.Password, time.Now(), time.Now()).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}

func (repo *repo) Edit(id int64, user models.User) int64 {
	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`

	// execute the sql statement
	res, err := repo.client.Exec(sqlStatement, id, user.Name, user.Location, user.Age)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func (repo *repo) Get(id int64) (u models.User, e error) {
	var user models.User

	sqlStatement := `SELECT * FROM users WHERE userid=$1`

	row := repo.client.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return user, nil
}

func (repo *repo) List() (users []models.User, err error) {
	var userList []models.User

	// create the select sql query
	sqlStatement := `SELECT * FROM users`

	rows, err := repo.client.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		userList = append(userList, user)

	}

	return userList, err
}

func (repo *repo) Delete(id int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM users WHERE userid=$1`

	// execute the sql statement
	res, err := repo.client.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
