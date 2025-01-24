package srvteapod

import (
	"github.com/enuesaa/teatime/pkg/repository/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) List() ([]string, error) {
	query := srv.repos.DB.Teapods()
	list := []string{}

	teapods := []db.Teapod{}
	filter := bson.M{}
	sort := bson.M{
		"created": 1,
	}
	if err := query.FindAll(filter, &teapods, sort); err != nil {
		return list, err
	}

	for _, teapod := range teapods {
		list = append(list, teapod.Name)
	}
	return list, nil
}
