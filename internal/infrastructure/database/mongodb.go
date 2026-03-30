package database

import (
	"context"
	"fmt"
	"time"

	"github.com/tigerbig/spatial-data-plateform/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectDatabase(config *config.Config) (*mongo.Client, error) {

	// Connect to database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s uri=%s sslmode=disable",
		config.DB_Host, config.DB_Port, config.DB_User, config.DB_Password,
		config.Uri)
	fmt.Println("DSN Connection string: ", dsn)

	client, err := mongo.Connect(options.Client().ApplyURI(config.Uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	errPing := client.Ping(ctx, nil)
	if errPing != nil {
		panic(errPing)
	}

	fmt.Println("Connected to MongoDB")
	fmt.Println("MIGRATE RUN")

	return client, nil
}
