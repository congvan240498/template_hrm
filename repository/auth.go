package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"hrm/pkg/mongodb"
)

type AuthRepository struct {
	*mongodb.Repository[Session]
}

func NewAuthRepository(dbStorage *mongodb.DatabaseStorage) *AuthRepository {
	return &AuthRepository{
		Repository: mongodb.NewRepository[Session](dbStorage),
	}
}

func (r *AuthRepository) CreateSession(ctx context.Context, req *Session) error {
	_, err := r.InsertOne(ctx, req)
	return err
}

func (r *AuthRepository) Logout(ctx context.Context, token string) error {
	filter := bson.M{"token": token}
	_, err := r.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"is_logout": true}})
	return err
}
