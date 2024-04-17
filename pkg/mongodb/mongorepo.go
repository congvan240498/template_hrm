package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrContextNotFoundKeyRegion = errors.New("mongo multi conn: context not found key region")
	ErrNotFoundRegion           = errors.New("mongo multi conn: mapping collections not found region")
)

type ModelInterface interface {
	CollectionName() string
}

type Repository[T ModelInterface] struct {
	*mongo.Collection
}

func NewRepository[T ModelInterface](dbStorage *DatabaseStorage) *Repository[T] {
	var t T

	return &Repository[T]{
		Collection: dbStorage.db.Collection(t.CollectionName()),
	}
}

func (r *Repository[T]) FindOneDoc(ctx context.Context, filter interface{}, options ...*options.FindOneOptions) (*T, error) {
	var m T
	err := r.Collection.FindOne(ctx, filter, options...).Decode(&m)
	if err != nil {
		return nil, err
	}

	return &m, err
}

func (r *Repository[T]) FindDocs(ctx context.Context, filter interface{}, options ...*options.FindOptions) ([]*T, error) {
	cs, err := r.Collection.Find(ctx, filter, options...)
	if err != nil {
		return nil, err
	}
	ms := make([]*T, 0)
	err = cs.All(ctx, &ms)
	if err != nil {
		return nil, err
	}

	return ms, nil
}
