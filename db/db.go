package db

import (
	"github.com/lucasweiblen/gomovieapi/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

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

func GetActor(name string, context *Context) (*models.Actor, error) {
	actor := models.Actor{}
	err := context.Collections["actors"].Find(bson.M{"name": name}).One(&actor)
	if err != nil {
		return nil, err
	}
	return &actor, nil
}

func DeleteActor(name string, context *Context) error {
	err := context.Collections["actors"].Remove(bson.M{"name": name})
	if err != nil {
		return err
	}
	return nil
}

func InsertActor(name string, age int, context *Context) error {
	actor := models.Actor{name, age}
	err := context.Collections["actors"].Insert(actor)
	if err != nil {
		return err
	}
	return nil
}

func UpdateActor(name string, newName string, context *Context) error {
	actor := models.Actor{newName, 90}
	err := context.Collections["actors"].Update(bson.M{"name": name}, actor)
	if err != nil {
		return err
	}
	return nil
}

func GetContext() (*Context, error) {
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
