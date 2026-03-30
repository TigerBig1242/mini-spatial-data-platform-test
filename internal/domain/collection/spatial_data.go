package collection

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type GeometryType string

const (
	PointType GeometryType = "Point"
	LinesStringType GeometryType = "LineString"
	PolygonType GeometryType = "Polygon"
	MultiPolygonType GeometryType = "MultiPolygon"
)

type Point struct {
	Type        GeometryType `bson:"type" json:"type"`
	Coordinates [2]float64   `bson:"coordinates" json:"coordinates"`
}

type LineString struct {
	Type        GeometryType `bson:"type" json:"type"`
	Coordinates [][]float64  `bson:"coordinates" json:"coordinates"`
}

type Polygon struct {
	Type        GeometryType   `bson:"type" json:"type"`
	Coordinates [][][]float64  `bson:"coordinates" json:"coordinates"`
}

type Features struct {
	ID         bson.ObjectID     	  `bson:"_id,omitempty" json:"id"`
	Type       string                 `bson:"type" json:"type"`
	Geometry   interface{}            `bson:"geometry" json:"geometry"`
	Properties map[string]interface{} `bson:"properties" json:"properties"`
}