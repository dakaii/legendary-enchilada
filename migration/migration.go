package migration

import (
	"context"
	"fmt"
	"legendaryenchilada/model"
	"legendaryenchilada/storage"
	"log"
	"math/rand"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// defining this migration function for demonstration purposes
func DataMigration() {
	db := storage.InitDatabase()
	for i := 0; i < 500; i++ {
		// giving a unique name to each scooter.
		u, err := uuid.NewV4()
		if err != nil {
			log.Fatal(err)
		}
		lat := randomFloat(1.48, 1.21)
		lon := randomFloat(104.1, 103.5)
		point := model.Point{
			Title: "purple-scooter-" + u.String(),
			Location: model.Location{
				GeoJSONType: "Point",
				Coordinates: []float64{lon, lat},
			}}
		addPoint(db, point)
	}
}
func randomFloat(max float64, min float64) float64 {
	return min + rand.Float64()*(max-min)
}
func addPoint(db storage.Storage, point model.Point) error {
	coll := db.Mongo.Collection(storage.Scooter)
	point.ID = primitive.NewObjectID()
	insertResult, err := coll.InsertOne(context.Background(), point)
	if err != nil {
		fmt.Printf("Could not insert new Point. Id: %s\n", point.ID)
		return err
	}
	fmt.Printf("Inserted new Point. ID: %s\n", insertResult.InsertedID)
	return nil
}
