package mongoconnect

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateData(name string, alter string, typeKey int) {
	collection := MongConnect()
	defer collection.Database().Client().Disconnect(context.TODO())

	filter := bson.D{{Key: "name", Value: name}}
	var update bson.D

	if typeKey == 1 {
		update = bson.D{{"$set", bson.D{{"name", alter}}}}
	} else if typeKey == 2 {
		update = bson.D{{"$set", bson.D{{"email", alter}}}}
	} else {
		fmt.Println("Tipo de dado invalido.")
		return
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Documentos combinados: %d, Documentos modificados: %d\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
