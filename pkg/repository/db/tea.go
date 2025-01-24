package db

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Tea struct {
	InternalId *bson.ObjectID         `bson:"_id"`
	Created    time.Time              `bson:"created"`
	Updated    time.Time              `bson:"updated"`
	Data       map[string]interface{} `bson:"data"`
}
func (tea *Tea) Id() string {
	if tea.InternalId == nil {
		return ""
	}
	return tea.InternalId.Hex()
}

type TeaQuery struct {
	query Query
}

func (q *TeaQuery) CreateCollection() error {
	return q.query.createCollection()
}

func (q *TeaQuery) DropCollection() error {
	return q.query.dropCollection()
}

func (q *TeaQuery) FindAll(filter bson.M, res *[]Tea, sort bson.M) error {
	return q.query.findAll(filter, res, sort)
}

func (q *TeaQuery) Find(filter bson.M, res *Tea) error {
	return q.query.find(filter, res)
}

func (q *TeaQuery) Create(data []byte) (string, error) {
	var datamap map[string]interface{}
	if err := json.Unmarshal(data, &datamap); err != nil {
		return "", err
	}

	now := time.Now()
	type Record struct {
		Data    map[string]interface{} `bson:"data"`
		Created time.Time              `bson:"created"`
		Updated time.Time              `bson:"updated"`
	}
	tea := Record{
		Data:    datamap,
		Created: now,
		Updated: now,
	}
	return q.query.create(tea)
}

func (q *TeaQuery) Update(teaId string, data []byte) (string, error) {
	var datamap map[string]interface{}
	if err := json.Unmarshal(data, &datamap); err != nil {
		return "", err
	}

	now := time.Now()
	type Record struct {
		Data    map[string]interface{} `bson:"data"`
		Updated time.Time              `bson:"updated"`
	}
	tea := Record{
		Data:    datamap,
		Updated: now,
	}
	return q.query.update(teaId, tea)
}

func (q *TeaQuery) Delete(teaId string) error {
	id, err := bson.ObjectIDFromHex(teaId)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": id,
	}

	return q.query.delete(filter)
}
