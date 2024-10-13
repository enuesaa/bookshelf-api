package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBRepositoryInterface interface {
	Connect() error
	Disconnect() error
	WithTransaction(fn func() error) error
	CreateCollection(name string, schema bson.M) error
	Create(name string, document interface{}) (string, error)
	FindAll(name string, filter bson.M, res interface{}) error
	Find(name string, filter bson.M, res interface{}) error
	Delete(name string, filter bson.M) error
}
type DBRepository struct {
	client *mongo.Client
	db     *mongo.Database
	sc     context.Context // do not use this directly. instead use repo.ctx()
}

func (repo *DBRepository) Connect() error {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	repo.client = client
	repo.db = client.Database("app")
	return nil
}

func (repo *DBRepository) Disconnect() error {
	if repo.client != nil {
		return repo.client.Disconnect(context.Background())
	}
	return nil
}

func (repo *DBRepository) ctx() context.Context {
	if repo.sc != nil {
		return repo.sc
	}
	return context.Background()
}

func (repo *DBRepository) WithTransaction(fn func() error) error {
	ctx := context.Background()

	session, err := repo.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func (sc context.Context) error {
		repo.sc = sc

		if err := session.StartTransaction(); err != nil {
			return err
		}
		if err := fn(); err != nil {
			session.AbortTransaction(sc)
			return err
		}
		session.CommitTransaction(sc)
		return nil
	})
	repo.sc = nil

	return err
}

func (repo *DBRepository) CreateCollection(name string, schema bson.M) error {
	validator := bson.M{
		"$jsonSchema": schema,
	}
	err := repo.db.CreateCollection(repo.ctx(), name,
		options.CreateCollection().SetValidator(validator),
		options.CreateCollection().SetValidationLevel("strict"),
		options.CreateCollection().SetValidationAction("error"),
	)
	return err
}

func (repo *DBRepository) FindAll(name string, filter bson.M, res interface{}) error {
	collection := repo.db.Collection(name)

	cur, err := collection.Find(repo.ctx(), filter)
	if err != nil {
		return err
	}
	return cur.All(repo.ctx(), res)
}

func (repo *DBRepository) Find(name string, filter bson.M, res interface{}) error {
	collection := repo.db.Collection(name)

	return collection.FindOne(repo.ctx(), filter).Decode(res)
}

func (repo *DBRepository) Create(name string, document interface{}) (string, error) {
	collection := repo.db.Collection(name)
	res, err := collection.InsertOne(repo.ctx(), document)
	if err != nil {
		return "", err
	}
	id := res.InsertedID

	return fmt.Sprintf("%s", id), nil
}

func (repo *DBRepository) Delete(name string, filter bson.M) error {
	collection := repo.db.Collection(name)
	if _, err := collection.DeleteOne(repo.ctx(), filter); err != nil {
		return err
	}
	return nil
}
