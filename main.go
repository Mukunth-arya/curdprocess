package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connection = "mongodb+srv://mukunth:mukunth@mycluster.jptcn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbname = "value"
const colname = "reviewvalue"

func main() {

	data1, err := mongo.NewClient(options.Client().ApplyURI(connection))
	if err != nil {

		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = data1.Connect(ctx)
	if err != nil {

		log.Fatal(err)
	}

	dataconnect := data1.Database(dbname)
	data2 := dataconnect.Collection(colname)
	println("hey there the database is created", dataconnect)
	//var datavalue bson.M
	//if err = data2.FindOne(ctx, bson.M{}).Decode(&datavalue); err != nil {

	//log.Fatal(err)
	//}
	//fmt.Println(datavalue)
	//var data []bson.M
	//if err = data5.All(ctx, &data); err != nil {
	//log.Fatal(err)
	//}
	//for _, data2 := range data {
	//fmt.Println(data2)
	//}
	//defer data5.Close(ctx)
	//for data5.Next(ctx) {
	//var data3 bson.M
	//if err = data5.Decode(&data3); err != nil {

	//log.Fatal(err)
	//}
	//fmt.Println(data3)

	//}
	datakey, err := data2.Find(ctx, bson.M{"name": "mukunth"})
	if err != nil {
		log.Fatal(err)
	}
	var keyvalue []bson.M
	if err = datakey.All(ctx, &keyvalue); err != nil {
		log.Fatal(err)
	}
	for _, value := range keyvalue {
		fmt.Println(value)
	}
}
