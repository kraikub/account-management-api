package repositories

import (
	"context"

	"github.com/kraikub/account-management-api/api/v1/internal/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Database
}

type UserRepository interface {
	FindUserWithUserId(ctx context.Context, id string) (entities.UserEntity, error)
	UpdateUserWithUserId(ctx context.Context, id string, updateE entities.UpdateableUser) error
}

func CreateUserRepository(db *mongo.Database) userRepository {
	return userRepository{
		db: db,
	}
}

func (u userRepository) collection() *mongo.Collection {
	return u.db.Collection(entities.UserEntity{}.CollectionName())
}

func (u userRepository) FindUserWithUserId(ctx context.Context, id string) (entities.UserEntity, error) {
	var user entities.UserEntity
	collection := u.collection()
	idFilter := bson.M{
		"uid": id,
	}
	err := collection.FindOne(ctx, idFilter).Decode(&user)
	return user, err
}

func (u userRepository) UpdateUserWithUserId(ctx context.Context, id string, updateE entities.UpdateableUser) error {
	collection := u.collection()
	idFilter := bson.M{
		"uid": id,
	}

	update := bson.M{}
	if updateE.UniversityEmail != "" {
		update["universityEmail"] = updateE.UniversityEmail
	}
	if updateE.PersonalEmail != "" {
		update["personalEmail"] = updateE.PersonalEmail
	}
	if updateE.ProfileImageUrl != "" {
		update["profileImageUrl"] = updateE.ProfileImageUrl
	}

	updateDocument := bson.M{
		"$set": update,
	}
	_, err := collection.UpdateOne(ctx, idFilter, updateDocument)
	return err
}
