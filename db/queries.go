package db

import "go.mongodb.org/mongo-driver/mongo"

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("ClusterMD-test").Collection(collectionName)
	return collection
}
