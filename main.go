package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connection = "mongodb+srv://mukunth:mukunth@mycluster.jptcn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const newvalue = "value"
const oldvalue = "reviewvalue"

type datafunction struct {
	ID   primitive.ObjectID `bson: "_id, omitempty"`
	name string             `bson: "name,omitempty"`
	age  string             `bson: age, omitempty`
}

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
	database1 := data1.Database(newvalue)
	data2 := database1.Collection(oldvalue)

	//dat,3, err := data2.InsertOne(ctx, bson.D{
	//{"name", "arya"},
	//{"age", "21"},
	//})
	//if err != nil {

	//log.Fatal(err)
	//}
	//fmt.Println("Here comes the value", data3.InsertedID)

	//data4, err := data2.InsertMany(ctx, []interface{}{
	//bson.D{
	//{"name", "kumar"},
	//{"age", "23"},
	//},
	//bson.D{

	//{"name", "vaibhav"},
	//{"age", "40"},
	//},
	//})
	//if err != nil {
	//log.Fatal(err)
	//}
	//fmt.Println("here we added the more documents", data4.InsertedIDs)
	//data3, err := data2.Find(ctx, bson.D{})
	//if err != nil {
	//log.Fatal(err)

	//}
	//for data3.Next(ctx) {
	//var file []bson.M
	//if err = data3.All(ctx, &file); err != nil {

	//log.Fatal(err)
	//}

	//fmt.Println(file)

	//}
	//var data3 bson.M
	//if err = data2.FindOne(ctx, bson.M{}).Decode(&data3); err != nil {

	//log.Fatal(err)
	//}
	//fmt.Println("here comes the data", data3)
	//data3, err := data2.Find(ctx, bson.M{"name": "arya"})
	//if err != nil {
	//log.Fatal(err)
	//}

	//var dataone []bson.M

	//if err = data3.All(ctx, &dataone); err != nil {
	//log.Fatal(err)
	//}
	//fmt.Print(dataone)
	//data4, _ := primitive.ObjectIDFromHex("62125d2836cb6651a3758090")
	//if err != nil{
	//log.Fatal(err)
	//}
	//data3, err := data2.UpdateOne(
	//ctx,
	//bson.M{"_id": data4},
	//bson.D{
	//{"$set", bson.D{{"name", "kumarboy"}}},
	//},
	//)

	//if err != nil {
	//log.Fatal(err)
	//}
	//fmt.Println("datamodifed count", data3.ModifiedCount)

	//data6, err := data2.UpdateMany(
	//ctx,
	//bson.M{"age": "21"},
	//bson.D{
	//{"$set", bson.D{{"name", "periyasamy"}}},
	//},
	//)
	//if err != nil {

	//log.Fatal(err)
	//}
	//fmt.Println(data6.ModifiedCount)
	//data6,err := data2.ReplaceOne(
	//ctx,
	//)
	//data7, err := data2.ReplaceOne(

	//ctx,
	//bson.M{"name": "arya"},
	//bson.M{
	//"name": "aryahome",
	//"age":  "21",
	//},
	//)
	//fmt.Println(data7.ModifiedCount)
	datafield, err := data2.DeleteOne(ctx, bson.M{"name": "kumar"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(datafield.DeletedCount)
}
