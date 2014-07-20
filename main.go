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

func main() {
	session, err := connect()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("locadora").C("actors")
	err = c.Insert(&Actor{"Lucas", 29})
	if err != nil {
		fmt.Println("Erro ao inserir")
	}
	fmt.Println("Inserido com sucesso")

	result := Actor{}
	err = c.Find(bson.M{"name": "Lucas"}).One(&result)
	if err != nil {
		fmt.Println("Erro ao requisitar actor")
	}
	fmt.Println("Name: ", result.Name)

	err = c.Insert(&Actor{"Joao", 50})
	if err != nil {
		fmt.Println("Erro ao inserir")
	}
	fmt.Println("Inserido com sucesso")

	err = c.Remove(bson.M{"name": "Joao"})
	if err != nil {
		fmt.Println("Erro ao remover")
	}

	result2 := Actor{}
	err = c.Find(bson.M{"name": "Joao"}).One(&result2)
	if err != nil {
		fmt.Println("Erro ao requisitar actor")
	}

	err = insertActor("John", 39, c)
	if err != nil {
		fmt.Println("Erro ao inserir novo ator")
	}
}

func insertActor(name string, age int, c *mgo.Collection) error {
	actor := Actor{name, age}
	err := c.Insert(actor)
	if err != nil {
		return err
	}
	return nil
}

func connect() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}
