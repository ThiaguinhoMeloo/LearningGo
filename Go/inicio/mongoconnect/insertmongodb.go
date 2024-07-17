package mongoconnect

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertMongoDb() {
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(context.TODO())

	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	collection := MongConnect()

	// collection := client.Database("PrimeiroBanco").Collection("PrimeiraCollection")

	users := []interface{}{
		bson.D{
			{Key: "name", Value: "Thiago Melo"},
			{Key: "email", Value: "thiagomelo329@hotmail.com"},
			{Key: "created_at", Value: time.Date(2024, 7, 16, 12, 34, 56, 0, time.UTC)},
		},
		bson.D{
			{Key: "name", Value: "Evelyn Pereira"},
			{Key: "email", Value: "e168486@dac.unicamp.br"},
			{Key: "created_at", Value: time.Date(2024, 7, 16, 12, 34, 56, 0, time.UTC)},
		},
		bson.D{
			{Key: "name", Value: "Leticia Correia de Melo"},
			{Key: "email", Value: "leticiacdemeloo@gmail.com"},
			{Key: "created_at", Value: time.Date(2024, 7, 16, 12, 34, 56, 0, time.UTC)},
		},
	}

	//Inserir Documentos
	insert, err := collection.InsertMany(context.TODO(), users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Documento inseridos com os Ids: \n", insert.InsertedIDs)
}
