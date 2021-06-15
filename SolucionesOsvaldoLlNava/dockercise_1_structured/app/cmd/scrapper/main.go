package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/OsvaldoLlNava/dockercises1/structured/app/pkg/DockerciseDatabase"
	"github.com/OsvaldoLlNava/dockercises1/structured/app/pkg/DockerciseModels"
)

func main() {

	ctx, client, cancel := DockerciseDatabase.Connect()
	defer DockerciseDatabase.Disconnect(ctx, client, cancel)

	// leer el archivo
	content, err := ioutil.ReadFile("./app/resources/info/people.xml")
	if err != nil {
		panic(err)
	}
	// fmt.Println(content)

	var people DockerciseModels.People

	err = xml.Unmarshal(content, &people)
	if err != nil {
		panic(err)
	}
	db := client.Database("dockercise1")
	coleccion := db.Collection("Personas")

	// for _, v := range people.Personas {
	// 	DockerciseDatabase.InsertarDocumento(ctx, coleccion, v)
	// }

	DockerciseDatabase.InsertarTodo(ctx, coleccion, people.Personas)

	fmt.Println("Proceso Finalizado")

}
