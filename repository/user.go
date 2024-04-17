package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"hrm/pkg/mongodb"
)

type UserRepository struct {
	*mongodb.Repository[User]
}

func NewUserRepository(dbStorage *mongodb.DatabaseStorage) *UserRepository {
	return &UserRepository{
		Repository: mongodb.NewRepository[User](dbStorage),
	}
}

func (r *UserRepository) GetUser(ctx context.Context, req *GetUserRequest) ([]*User, error) {
	filter := bson.M{}
	if req != nil {
		if req.Username != "" {
			filter["user_name"] = req.Username
		}
	}

	return r.FindDocs(ctx, filter)
}

func (r *UserRepository) GetUserByUserName(ctx context.Context, username string) (*User, error) {
	filter := bson.M{"user_name": username}
	return r.FindOneDoc(ctx, filter)
}

func (r *UserRepository) CreateUser(ctx context.Context, req *User) error {
	_, err := r.InsertOne(ctx, req)
	return err
}

func (r *UserRepository) UpdateUser(ctx context.Context, req *User) error {
	filter := bson.M{"user_name": req.UserName}
	update := bson.M{"$set": bson.M{
		"password":  req.Password,
		"full_name": req.FullName,
		"phone":     req.Phone,
		"email":     req.Email,
	}}
	_, err := r.UpdateOne(ctx, filter, update)
	return err
}
