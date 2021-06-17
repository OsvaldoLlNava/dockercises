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
	people := ReadFromXml("./app/resources/info/people.xml")

	// fmt.Println(content)

	db := client.Database("dockercise1")
	coleccion := db.Collection("Personas")

	DockerciseDatabase.InsertarTodo(ctx, coleccion, people.Personas)

	fmt.Println("Proceso Finalizado")

}

func ReadFromXml(path string) DockerciseModels.People {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var people DockerciseModels.People

	err = xml.Unmarshal(content, &people)
	if err != nil {
		panic(err)
	}

	return people
}
