package dtos

import "github.com/kraikub/account-management-api/api/v1/internal/entities"

type UserDTO struct {
	AppOwned        int    `json:"appOwned"`
	AppQuota        int    `json:"appQuota"`
	SigninSignature string `json:"signinSignature"`
	Uid             string `json:"uid"`
	ShouldUpdate    bool   `json:"shouldUpdate"`
	ProfileImageUrl string `json:"profileImageUrl"`
	UniversityEmail string `json:"universityEmail"`
	PersonalEmail   string `json:"personalEmail"`
}

func (UserDTO) FromEntity(e entities.UserEntity) UserDTO {
	return UserDTO{
		AppOwned:        e.AppOwned,
		AppQuota:        e.AppQuota,
		SigninSignature: e.SigninSignature,
		Uid:             e.Uid,
		ShouldUpdate:    e.ShouldUpdate,
		ProfileImageUrl: e.ProfileImageUrl,
		UniversityEmail: e.UniversityEmail,
		PersonalEmail:   e.PersonalEmail,
	}
}

func (u UserDTO) ToUpdateableEntity() entities.UpdateableUser {
	return entities.UpdateableUser{
		ProfileImageUrl: u.ProfileImageUrl,
		UniversityEmail: u.UniversityEmail,
		PersonalEmail:   u.PersonalEmail,
	}
}
