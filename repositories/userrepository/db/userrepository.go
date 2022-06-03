package userdb

import (
	"context"
	"database/sql"
	"fmt"
	"go-postgres/models"
	"log"
	"time"

	"github.com/lib/pq"
)

type repo struct {
	client *sql.DB
}

func NewRepo(client *sql.DB) *repo {
	return &repo{
		client: client,
	}
}

func (repo *repo) Add(ctx context.Context, user models.User) int64 {
	sqlStatement := `INSERT INTO users (name, location, age, email, password, createdat, updatedat) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING userid`
	var id int64

	err := repo.client.QueryRow(sqlStatement, user.Name, user.Location, user.Age, user.Email, user.Password, time.Now(), time.Now()).Scan(&id)

	pqErr := err.(*pq.Error)
	log.Println(pqErr.Code)

	if pqErr.Code == "23505" {
		log.Println("Email already exist")
	}
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id
}

// func (repo *repo) Add(ctx context.Context, user models.User) (err error) {

// 	brand.CreatedAt = time.Now()
// 	brand.UpdatedAt = time.Now()
// 	_, err = repo.client.Database(repo.dbname).Collection(collection).InsertOne(ctx, brand)
// 	if err != nil {
// 		if mongo.IsDuplicateKeyError(err) {
// 			return brandrepository.ErrBrandAlreadyExists
// 		}
// 		return err
// 	}

// 	return nil
// }

func (repo *repo) Get(ctx context.Context, id int64) (u models.User, e error) {
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

// func (repo *repo) List(ctx context.Context) (brands []models.Brand, err error) {
// 	cur, err := repo.client.Database(repo.dbname).Collection(collection).Find(ctx, bson.D{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = cur.All(ctx, &brands)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return brands, nil
// }

// func (repo *repo) Update(ctx context.Context, id primitive.ObjectID, name string, logoUrl string, priority int) (err error) {
// 	res, err := repo.client.Database(repo.dbname).Collection(collection).UpdateByID(ctx, id,
// 		bson.M{
// 			"$set": bson.M{
// 				"name":       name,
// 				"logo_url":   logoUrl,
// 				"priority":   priority,
// 				"updated_at": time.Now(),
// 			},
// 		},
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	if res.MatchedCount == 0 {
// 		return brandrepository.ErrBrandNotFound
// 	}

// 	return nil
// }
