package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Migrate(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collections := []string{
		"point",
		"linstring",
		"polygon",
		"feature",
	}

	for _, collectionName := range collections {
		err := db.CreateCollection(ctx, collectionName)
		if err != nil {
			if !isCollectionExistsError(err) {
				return fmt.Errorf("failed to create collection %s :", collectionName)
			}
			log.Printf("Collection '%s' already exists, skipping...\n", collectionName)
			continue
		}
		log.Printf("Collection '%s' created\n", collectionName)
	}

	fmt.Println("DB", db.Name())

	fmt.Println("Migrated Complete")

	return nil
}

func isCollectionExistsError(err error) bool {
	return strings.Contains(err.Error(), "already exists")
}
