package userServices

import (
	"database/sql"
	"fmt"
	"go-postgres/middleware"
	"go-postgres/models"
	"log"
)

func searchUser(email string) (models.User, error) {
	db := middleware.CreateConnection()
	defer db.Close()

	var user models.User

	sqlStatement := `select * from users where email = $1`
	row := db.QueryRow(sqlStatement, email)

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
