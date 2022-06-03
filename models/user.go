package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// User schema of the user table
type User struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty" validate:"required"`
	Location  string    `json:"location,omitempty"`
	Age       int64     `json:"age,omitempty"`
	Email     string    `json:"email,omitempty" validate:"required" sql:"email"`
	Password  string    `json:"password,omitempty" validate:"required" sql:"password"`
	Username  string    `json:"username,omitempty" sql:"username"`
	TokenHash string    `json:"tokenhash,omitempty" sql:"tokenhash"`
	CreatedAt time.Time `json:"createdat,omitempty" sql:"createdat"`
	UpdatedAt time.Time `json:"updatedat,omitempty" sql:"updatedat"`
}

type DBStore struct {
	Store *sql.DB
}

var dbStore DBStore

func InnitializeUserDBService(conn *sql.DB) {
	dbStore.Store = conn
}

func (user *User) InsertUser() int64 {
	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO users (name, location, age, email, password, createdat, updatedat) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING userid`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := dbStore.Store.QueryRow(sqlStatement, user.Name, user.Location, user.Age, user.Email, user.Password, time.Now(), time.Now()).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

func (u User) GetUser(id int64) (User, error) {
	var user User

	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE userid=$1`

	// execute the sql statement
	row := dbStore.Store.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
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

	// return empty user on error
	return user, err
}

func (u User) GetAllUsers() ([]User, error) {
	var users []User

	// create the select sql query
	sqlStatement := `SELECT * FROM users`

	// execute the sql statement
	rows, err := dbStore.Store.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var user User

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		users = append(users, user)

	}

	// return empty user on error
	return users, err
}

func (user *User) UpdateUser(id int64) int64 {
	// create the update sql query
	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`

	// execute the sql statement
	res, err := dbStore.Store.Exec(sqlStatement, id, user.Name, user.Location, user.Age)

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

func (user User) DeleteUser(id int64) int64 {
	// create the delete sql query
	sqlStatement := `DELETE FROM users WHERE userid=$1`

	// execute the sql statement
	res, err := dbStore.Store.Exec(sqlStatement, id)

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

func (u User) GetUserByEmail(email string) (User, error) {
	var user User

	sqlStatement := `select * from users where email = $1`
	row := dbStore.Store.QueryRow(sqlStatement, email)

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

	// return empty user on error
	return user, err
}
