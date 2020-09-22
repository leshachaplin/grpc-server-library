package repository

import (
	"context"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	//"gopkg.in/mgo.v2"
)

type Service struct {
	Id   bson.ObjectId `bson:"_id"`
	Type string        `bson:"type"`
	Url  string        `bson:"url"`
}

type RepositoryOfServiceState struct {
	db *mgo.Database
}

func NewRepositoryOfServiceState(database *mgo.Database) *RepositoryOfServiceState {
	return &RepositoryOfServiceState{
		db: database,
	}
}

func (r *RepositoryOfServiceState) GetServices(ctx context.Context, typeOfService string) (map[string]string, error) {
	var err error
	query := bson.M{}
	services := []Service{}
	result := make(map[string]string, 0)

	err = r.db.C("services").Find(query).All(&services)
	if err != nil {
		return nil, err
	}
	for _, v := range services {
		result[v.Type] = v.Url
	}
	return result, err
}

func (r *RepositoryOfServiceState) IfExistService(ctx context.Context, key string) (bool, error) {
	var err error
	query := bson.M{
		"type": bson.M{
			"$eq": key,
		},
	}
	service := Service{}

	err = r.db.C("services").Find(query).All(&service)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RepositoryOfServiceState) AddService(ctx context.Context, typeOfService, url string) error {
	var err error
	c := &Service{
		Id:   bson.NewObjectId(),
		Type: typeOfService,
		Url:  url,
	}
	err = r.db.C("services").Insert(&c)

	return err
}

func (r *RepositoryOfServiceState) DeleteService(ctx context.Context, typeOfService, url string) error {
	var err error
	_, err = r.db.C("services").RemoveAll(bson.M{"key": typeOfService, "value": url})
	return err
}
