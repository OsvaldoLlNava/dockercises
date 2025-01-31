package DockerciseDatabase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/OsvaldoLlNava/dockercises1/structured/app/pkg/DockerciseModels"
)

func Connect() (context.Context, *mongo.Client, context.CancelFunc) {
	uri := "mongodb://mongo_compose:27017/"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	credential := options.Credential{
		Username: "quesitoUser",
		Password: "SecretCheese1234",
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(credential))
	if err != nil {
		fmt.Println(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	return ctx, client, cancel
}

func Disconnect(ctx context.Context, client *mongo.Client, cancel context.CancelFunc) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
	cancel()
}

func ObtenerPersonas() []DockerciseModels.Person {

	ctx, client, cancel := Connect()
	defer Disconnect(ctx, client, cancel)

	db := client.Database("dockercise1")
	collection := db.Collection("Personas")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	var p []DockerciseModels.Person

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &p)
	if err != nil {
		panic(err)
	}

	return p

}

func ObtenerUnaPersona(id string) DockerciseModels.Person {

	ctx, client, cancel := Connect()
	defer Disconnect(ctx, client, cancel)

	var p DockerciseModels.Person

	numero, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}

	db := client.Database("dockercise1")
	collection := db.Collection("Personas")
	err = collection.FindOne(ctx, bson.D{{Key: "id", Value: numero}}).Decode(&p)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return p
		} else {
			panic(err)
		}
	}

	return p

}

func InsertarTodo(ctx context.Context, coleccion *mongo.Collection, p []DockerciseModels.Person) {

	datos := make([]interface{}, len(p))

	for i, v := range p {
		datos[i] = bson.D{
			{Key: "id", Value: v.Id},
			{Key: "first_name", Value: v.First_Name},
			{Key: "last_name", Value: v.Last_Name},
			{Key: "company", Value: v.Company},
			{Key: "email", Value: v.Email},
			{Key: "ip_address", Value: v.Ip_Address},
			{Key: "phone_number", Value: v.Phone_Number},
		}
	}

	// fmt.Println(datos)

	_, err := coleccion.InsertMany(
		ctx, datos,
	)
	if err != nil {
		panic(err)
	}
}
