package main

import (
	"fmt"

	"github.com/lucasweiblen/gomovieapi/db"
	"github.com/lucasweiblen/gomovieapi/models"
)

func main() {
	context, err := db.GetContext()
	if err != nil {
		panic(err)
	}
	defer context.SessionWrapper.Session.Close()

	newActor := models.Actor{"Lennon", 29}
	err = db.Insert(newActor, context)
	if err != nil {
		fmt.Printf("Error inserting %v", newActor)
		return
	}
	fmt.Println("insertiu normalmente")

	newMovie := models.Movie{"Home Alone", 1998, []*models.Actor{&newActor}}

	err = db.Insert(newMovie, context)
	if err != nil {
		fmt.Printf("Error inserting %v", newMovie)
		return
	}
	fmt.Println("insertiu normalmente")

	query := db.QueryParams{"actor", "name", "Lennon"}
	result, err := db.Get(&query, context)
	if err != nil {
		fmt.Println("Erro")
		return
	}
	fmt.Println(result)
	if _, ok := result.(models.Actor); ok {
		fmt.Println("Eh um ator")
		err = db.Delete(result, context)
		if err != nil {
			fmt.Println("Error deleting")
			return
		}
		fmt.Println("Deletado com sucesso")
	}
}
