package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	ID              primitive.ObjectID `bson:"_id"`
	AppOwned        int                `bson:"appOwned"`
	AppQuota        int                `bson:"appQuota"`
	SigninSignature string             `bson:"signinSignature"`
	Uid             string             `bson:"uid"`
	ShouldUpdate    bool               `bson:"shouldUpdate"`
	ProfileImageUrl string             `bson:"profileImageUrl"`
	UniversityEmail string             `bson:"universityEmail"`
	PersonalEmail   string             `bson:"personalEmail"`
	CreatedAt       time.Time          `bson:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt"`
}

type UpdateableUser struct {
	ProfileImageUrl string             `bson:"profileImageUrl"`
	UniversityEmail string             `bson:"universityEmail"`
	PersonalEmail   string             `bson:"personalEmail"`
}

func (UserEntity) CollectionName() string {
	return "users"
}

func (e *UserEntity) MarshalBSON() ([]byte, error) {
	if e.ID.IsZero() {
		e.ID = primitive.NewObjectID()
	}

	now := time.Now().UTC()
	if e.CreatedAt.IsZero() {
		e.CreatedAt = now
	}
	e.UpdatedAt = now

	type ue UserEntity
	return bson.Marshal((*ue)(e))
}
