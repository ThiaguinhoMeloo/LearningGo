package mongoconnect

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongConnect() {
	//Conectar ao MongoDb
	clienteOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clienteOptions)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	//Verificar a conexão com o MongoDb
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	//Escolher a coleção no MongoDb
	collection := client.Database("PrimeiroBanco").Collection("PrimeiraCollection")

	//Consultar dados no MongoDb
	filter := bson.D{{}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	//Iterar sobre os resultados do MongoDb
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("MongoDb - ", result)
	}

	//Verificar por erros durante a iteração no MongoDb
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
}