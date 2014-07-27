package main

import (
	"fmt"

	"github.com/lucasweiblen/gomovieapi/db"
)

func main() {
	context, err := db.GetContext()
	if err != nil {
		panic(err)
	}
	defer context.SessionWrapper.Session.Close()

	err = db.InsertActor("Paura", 50, context)
	if err != nil {
		fmt.Println("Error inserting new actor")
		return
	}

	actor, err := db.GetActor("Paura", context)
	if err != nil {
		fmt.Println("Error getting actor")
		return
	}
	fmt.Printf("Actor has name %s\n", actor.Name)

	err = db.DeleteActor("Paura", context)
	if err != nil {
		fmt.Println("Error deleting actor")
		return
	}
	fmt.Println("Deleted successfully")

	err = db.InsertActor("Joao", 36, context)
	if err != nil {
		fmt.Println("Error inserting new actor")
	}

	err = db.UpdateActor("Joao", "Lucas", context)
	if err != nil {
		fmt.Println("Error updating actor")
	}
}
