package db

import (
	"context"
	"log"

	"github.com/olusamimaths/go-book-application/src/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBHandler struct {
	MongoClient mongo.Client
	database *mongo.Database
}

func NewDBHandler(connectString string, dbname string) (DBHandler, error) {
	dbHandler := DBHandler{}
	clientOptions := options.Client().ApplyURI(connectString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}

	err =  client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return dbHandler, err
	}

	dbHandler.MongoClient = *client
	dbHandler.database = client.Database(dbname)
	return dbHandler, nil
}

func (dbHandler DBHandler) FindAllBooks() ([]*domain.Book, error) {
	var results []*domain.Book
	collection := dbHandler.database.Collection("books")

	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem domain.Book
		err = cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	return results, nil
}

func (dbHandler DBHandler) SaveBook(book domain.Book) error {
	collection := dbHandler.database.Collection("books")

	_, err := collection.InsertOne(context.TODO(), book)
	if err != nil {
		return err
	}
	return nil
}

func (dbHandler DBHandler) SaveAuthor(author domain.Author) error {
	collection := dbHandler.database.Collection("authors")

	_, err := collection.InsertOne(context.TODO(), author)
	if err != nil {
		return err
	}
	return nil
}