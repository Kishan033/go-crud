package userdb

import (
	"context"
	"database/sql"
	"fmt"
	"go-postgres/models"
	dflog "log"
	"time"

	"github.com/go-kit/kit/log"
)

type repo struct {
	client *sql.DB
	logger log.Logger
}

func NewRepo(client *sql.DB, logger log.Logger) *repo {
	return &repo{
		client: client,
		logger: log.With(logger, "user-rep", "postgres"),
	}
}

func (repo *repo) Add(ctx context.Context, user models.User) (int64, error) {
	sqlStatement := `INSERT INTO users (name, location, age, email, password, createdat, updatedat) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING userid`
	var id int64

	err := repo.client.QueryRow(sqlStatement, user.Name, user.Location, user.Age, user.Email, user.Password, time.Now(), time.Now()).Scan(&id)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Inserted a single record %v", id)

	return id, nil
}

func (repo *repo) Edit(ctx context.Context, id int64, user models.User) (int64, error) {
	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`

	res, err := repo.client.Exec(sqlStatement, id, user.Name, user.Location, user.Age)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (repo *repo) Get(ctx context.Context, id int64) (models.User, error) {
	var user models.User

	sqlStatement := `SELECT * FROM users WHERE userid=$1`

	row := repo.client.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, err
	case nil:
		return user, nil
	default:
		return user, err
	}
}

func (repo *repo) List(ctx context.Context) (users []models.User, err error) {
	var userList []models.User

	// create the select sql query
	sqlStatement := `SELECT * FROM users`

	rows, err := repo.client.Query(sqlStatement)

	if err != nil {
		dflog.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			dflog.Fatalf("Unable to scan the row. %v", err)
		}
		userList = append(userList, user)
	}

	return userList, err
}

func (repo *repo) Delete(ctx context.Context, id int64) (int64, error) {
	sqlStatement := `DELETE FROM users WHERE userid=$1`

	res, err := repo.client.Exec(sqlStatement, id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
