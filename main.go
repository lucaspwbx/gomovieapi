package main

import (
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Movie struct {
	Title  string
	Year   int
	Actors []*Actor
}

type Actor struct {
	Name string
	Age  int
}

type Context struct {
	SessionWrapper SessionWrapper
	Collections    map[string]*mgo.Collection
}

type SessionWrapper struct {
	Session *mgo.Session
}

func (sw *SessionWrapper) getCollections() map[string]*mgo.Collection {
	m := make(map[string]*mgo.Collection)
	m["actors"] = sw.Session.DB("locadora").C("actors")
	m["movies"] = sw.Session.DB("locadora").C("movies")
	return m
}

func main() {
	context, err := getContext()
	if err != nil {
		panic(err)
	}
	defer context.SessionWrapper.Session.Close()

	err = insertActor("Paura", 50, context)
	if err != nil {
		fmt.Println("Error inserting new actor")
		return
	}

	actor, err := getActor("Paura", context)
	if err != nil {
		fmt.Println("Error getting actor")
		return
	}
	fmt.Printf("Actor has name %s\n", actor.Name)

	err = deleteActor("Paura", context)
	if err != nil {
		fmt.Println("Error deleting actor")
		return
	}
	fmt.Println("Deleted successfully")

	err = insertActor("Joao", 36, context)
	if err != nil {
		fmt.Println("Error inserting new actor")
	}

	err = updateActor("Joao", "Lucas", context)
	if err != nil {
		fmt.Println("Error updating actor")
	}
}

func getActor(name string, context *Context) (*Actor, error) {
	actor := Actor{}
	err := context.Collections["actors"].Find(bson.M{"name": name}).One(&actor)
	if err != nil {
		return nil, err
	}
	return &actor, nil
}

func deleteActor(name string, context *Context) error {
	err := context.Collections["actors"].Remove(bson.M{"name": name})
	if err != nil {
		return err
	}
	return nil
}

func insertActor(name string, age int, context *Context) error {
	actor := Actor{name, age}
	err := context.Collections["actors"].Insert(actor)
	if err != nil {
		return err
	}
	return nil
}

func updateActor(name string, newName string, context *Context) error {
	actor := Actor{newName, 90}
	err := context.Collections["actors"].Update(bson.M{"name": name}, actor)
	if err != nil {
		return err
	}
	return nil
}

func getContext() (*Context, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	sw := SessionWrapper{session}
	collection := sw.getCollections()

	context := Context{sw, collection}
	return &context, nil
}
