package db

import (
	"errors"

	"github.com/lucasweiblen/gomovieapi/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Context struct {
	SessionWrapper SessionWrapper
	Colls          map[string]*mgo.Collection
}

type SessionWrapper struct {
	Session *mgo.Session
}

type QueryParams struct {
	Kind, Attr, Value string
}

func (sw *SessionWrapper) getCollections() map[string]*mgo.Collection {
	m := make(map[string]*mgo.Collection)
	m["actors"] = sw.Session.DB("locadora").C("actors")
	m["movies"] = sw.Session.DB("locadora").C("movies")
	return m
}

func GetActor(name string, context *Context) (*models.Actor, error) {
	actor := models.Actor{}
	err := context.Colls["actors"].Find(bson.M{"name": name}).One(&actor)
	if err != nil {
		return nil, err
	}
	return &actor, nil
}

func DeleteActor(name string, context *Context) error {
	err := context.Colls["actors"].Remove(bson.M{"name": name})
	if err != nil {
		return err
	}
	return nil
}

func Get(q *QueryParams, cnt *Context) (interface{}, error) {
	var err error
	//`var result models.Actor
	//ar result interface{}
	if q.Kind == "actor" {
		result := models.Actor{}
		err = cnt.Colls["actors"].Find(bson.M{q.Attr: q.Value}).One(&result)
		return result, nil
	} else {
		result := models.Movie{}
		err = cnt.Colls["movies"].Find(bson.M{q.Attr: q.Value}).One(&result)
		return result, nil
	}
	if err != nil {
		return nil, err
	}
	return "", nil
}

func Insert(doc interface{}, cnt *Context) error {
	var err error
	if kind := checkType(doc); kind == "actor" {
		err = cnt.Colls["actors"].Insert(doc)
	} else if kind == "movie" {
		err = cnt.Colls["movies"].Insert(doc)
	} else {
		return errors.New("no collection")
	}
	if err != nil {
		return err
	}
	return nil
}

func Update(oldDoc interface{}, newDoc interface{}, cnt *Context) error {
	var err error
	if kind := checkType(newDoc); kind == "actor" {
		err = cnt.Colls["actors"].Update(oldDoc, newDoc)
	} else if kind == "movie" {
		err = cnt.Colls["movies"].Update(oldDoc, newDoc)
	} else {
		return errors.New("no collection")
	}
	if err != nil {
		return err
	}
	return nil
}

func checkType(obj interface{}) string {
	if _, ok := obj.(models.Actor); ok {
		return "actor"
	}
	if _, ok := obj.(models.Movie); ok {
		return "movie"
	}
	return ""
}

func UpdateActor(name string, newName string, context *Context) error {
	actor := models.Actor{newName, 90}
	err := context.Colls["actors"].Update(bson.M{"name": name}, actor)
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
