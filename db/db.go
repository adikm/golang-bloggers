package db

import (
	"context"
	"github.com/adikm/golang-bloggers/app/feed"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func connect() (*mongo.Client, context.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return mongoClient, ctx
}

func InsertEntries(entries *[]feed.Entry) { // TODO make it a function for storing newsletter issues
	mongoClient, ctx := connect()
	collection := mongoClient.Database("bloggers").Collection("bloggers")
	issue := NewsletterIssue{
		Entries: *entries,
		Date:    time.Now(),
		Number:  0,
	}
	_, err := collection.InsertOne(ctx, issue)
	if err != nil {
		log.Println("Can't store")
	}
}

type NewsletterIssue struct {
	Entries []feed.Entry
	Date    time.Time
	Number  int
}
