package repository

import (
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	FullName  string             `json:"fullName" bson:"full_name,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Phone     string             `json:"phone" bson:"phone,omitempty"`
	AccountId string             `json:"accountId" bson:"account_id,omitempty"`
	RoleCode  string             `json:"roleCode" bson:"role_code,omitempty"`
	UserName  string             `json:"userName" bson:"user_name,omitempty"`
	HashPass  string             `json:"hashPass" bson:"hash_pass,omitempty"`
	IsActive  *bool              `json:"isActive" bson:"is_active,omitempty"`

	NewUser  *bool  `json:"newUser" bson:"new_user,omitempty"`
	Password string `json:"password"`

	CreatedTime *time.Time `json:"createdTime" bson:"created_time,omitempty"`
	UpdatedTime *time.Time `json:"updatedTime" bson:"last_updated_time,omitempty"`
}

func (User) CollectionName() string {
	return "user"
}

type Claims struct {
	UserName string `json:"userName"`
	RoleCode string `json:"roleCode"`
	IsActive *bool  `json:"isActive"`
	jwt.StandardClaims
}

type Session struct {
	Id       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	UserName string             `json:"userName" bson:"user_name,omitempty"`
	Token    string             `json:"token" bson:"token,omitempty"`
	IsLogout *bool              `json:"isLogout" bson:"is_logout,omitempty"`
	Expired  *time.Time         `json:"expired" bson:"expired,omitempty"`
}

func (Session) CollectionName() string {
	return "session"
}

type GetUserRequest struct {
	Username string `json:"userName"`
	IsActive *bool  `json:"isActive"`
	NewUser  *bool  `json:"newUser"`
	RoleCode string `json:"roleCode"`
}

type GetSessionRequest struct {
	Token string `json:"token"`
}
