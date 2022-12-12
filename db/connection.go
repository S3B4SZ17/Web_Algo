package db

import (
	"context"

	"github.com/S3B4SZ17/Web_Algo/management"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
var (
	dbConn string
	//Client instance
	DB *mongo.Client
)

/*
InitializeOAuthGoogle Function
*/
func IntializeDB() {
	dbConn = viper.GetString("mongoDB.driver") + "://" + viper.GetString("mongoDB.username") + ":" + viper.GetString("mongoDB.password") + "@" + viper.GetString("mongoDB.host") + viper.GetString("mongoDB.options")
	DB = ConnectMongodb()
}

func ConnectMongodb() *mongo.Client {
	management.Log.Info("MongoDB initialization ..." + dbConn)
	// intializeConnString()
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbConn))
	if err != nil {
		management.Log.Panic(err.Error())
	}
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		management.Log.Panic(err.Error())
	// 	}
	// 	management.Log.Info("Client disconnected ...")
	// }()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		management.Log.Panic(err.Error())
	}
	management.Log.Info("Successfully connected and pinged.")
	return client
}

func CloseClientDB() {
	if DB == nil {
		return
	}

	err := DB.Disconnect(context.TODO())
	if err != nil {
		management.Log.Panic(err.Error())
	}

	management.Log.Info("Connection to MongoDB closed.")
}
