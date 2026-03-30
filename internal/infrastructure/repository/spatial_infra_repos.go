package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/tigerbig/spatial-data-plateform/internal/domain/collection"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SpatialRepo struct {
	db *mongo.Database
}

func NewSpatialRepo(db *mongo.Database) *SpatialRepo {
	return &SpatialRepo{
		db: db,
	}
}

func (repo *SpatialRepo) Create(ctx context.Context, feature *collection.Features) (*collection.Features, error) {
	coll := repo.db.Collection("features")

	response, err := coll.InsertOne(ctx, feature)
	if err != nil {
		return nil, err
	}

	feature.ID = response.InsertedID.(bson.ObjectID)

	return feature, nil
}

func bsonDtoM(d bson.D) bson.M {
	m := bson.M{}
	for _, elem := range d {
		m[elem.Key] = elem.Value
	}

	return m
}

func (repo *SpatialRepo) FindAll(ctx context.Context) ([]collection.Features, error) {
	coll := repo.db.Collection("features")

	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var raw []bson.M
	errRaw := cursor.All(ctx, &raw)
	if errRaw != nil {
		return nil, errRaw
	}

	result := make([]collection.Features, 0)
	for _, doc := range raw {
		geoRaw := doc["geometry"]
		var geo bson.M
		log.Printf("RAW GEO: %+v", geo)
		var props map[string]interface{}
		switch value := geoRaw.(type) {
		case bson.M:
			geo = value
		case bson.D:
			geo = bsonDtoM(value)
		default:
			return nil, fmt.Errorf("invalid geometry type: %T", value)
		}

		switch valueProps := doc["properties"].(type) {
		case bson.M:
			props = valueProps
		case bson.D:
			props = bsonDtoM(valueProps)
		case nil:
			props = map[string]interface{}{}
		default:
			return nil, fmt.Errorf("invalid properties type: %T", valueProps)
		}
		geometry, errGeometry := decodeGeometry(geo)
		if errGeometry != nil {
			return nil, errGeometry
		}

		idRaw, ok := doc["_id"]
		if !ok {
			return nil, fmt.Errorf("_id not found")
		}

		id, ok := idRaw.(bson.ObjectID)
		if !ok {
			return nil, fmt.Errorf("_id is not ObjectID, got %T", idRaw)
		}

		typeStr, _ := doc["type"].(string)

		feature := collection.Features{
			ID:         id,
			Type:       typeStr,
			Geometry:   geometry,
			Properties: props,
		}

		log.Printf("Feature Type is : %T", feature)
		log.Println("Feature value is :", feature)
		result = append(result, feature)
	}

	return result, nil
}

func decodeGeometry(raw bson.M) (interface{}, error) {
	types, ok := raw["type"].(string)
	if !ok {
		return nil, fmt.Errorf("geometry type invalid")
	}
	bytes, _ := bson.Marshal(raw)

	switch types {
	case "Point":
		var point collection.Point
		bson.Unmarshal(bytes, &point)
		return point, nil

	case "LineString":
		var lineString collection.LineString
		bson.Unmarshal(bytes, &lineString)
		return lineString, nil

	case "Polygon":
		var polygon collection.Polygon
		bson.Unmarshal(bytes, &polygon)
		return polygon, nil

	default:
		return raw, nil
	}
}

func (repo *SpatialRepo) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	coll := repo.db.Collection("features")

	fmt.Println("DELETE ID:", id.Hex())

	var findID bson.M
	err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&findID)
	log.Printf("Find ID type : %T", findID)
	log.Printf("Find ID : %v", findID)

	if err != nil {
		fmt.Println("NOT FOUND BEFORE DELETE")
	} else {
		fmt.Println("FOUND:", findID)
	}

	// test string
	var testStr bson.M
	errStr := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&testStr)

	if errStr == nil {
		fmt.Println("FOUND AS STRING:", testStr)
	}

	filter := bson.M{
		"_id": id,
	}

	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Delete not success : %v", err)
		return err
	}

	fmt.Println("DB NAME:", repo.db.Name())
	fmt.Println("COLLECTION:", coll.Name())

	if result.DeletedCount == 0 {
		return fmt.Errorf("document not found")
	}

	log.Println("Infrastructure delete not success", result)
	return nil
}
