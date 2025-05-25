package usecase

import (
	"errors"

	"github.com/mafzaidi/elog/internal/account"
	"github.com/mafzaidi/elog/internal/entities"
	"github.com/mafzaidi/elog/internal/service"
	"github.com/mafzaidi/elog/internal/user"
	"github.com/mafzaidi/elog/pkg/authorizer/masterkey"
	"github.com/mafzaidi/elog/pkg/authorizer/pwd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountUC struct {
	repo    account.Repository
	svcRepo service.Repository
	usrRepo user.Repository
}

func NewAccountUseCase(
	repo account.Repository,
	svcRepo service.Repository,
	usrRepo user.Repository) account.UseCase {
	return &AccountUC{
		repo:    repo,
		svcRepo: svcRepo,
		usrRepo: usrRepo,
	}
}

func (u *AccountUC) Store(pl *account.CreateParams) error {

	svcFilter := bson.M{}

	svcFilter["isActive"] = true
	svcFilter["group"] = "SERVICE"

	if pl.Service != "" {
		svcFilter["key"] = pl.Service
	}

	svc, err := u.svcRepo.FindByFilter(svcFilter)
	if err != nil {
		return errors.New("failed to get service")
	}

	filter := bson.M{
		"$and": []bson.M{
			{"service._id": svc.ID},
			{"service.code": svc.Attributes.Code},
			{"service.key": svc.Key},
			{"service.name": svc.Attributes.Name},
			{"username": pl.Username},
		},
	}

	user, err := u.usrRepo.FindByID(pl.UserID)
	if err != nil || !pwd.CheckHash(user.Password, pl.PasswordApp) {
		return errors.New("user or password is not valid")
	}

	masterKey, err := masterkey.Decrypt(user.MasterKeyEnc, pl.PasswordApp, user.Salt)
	if err != nil {
		return errors.New("failed to decrypt master key")
	}

	encPwd, err := masterkey.Encrypt(masterKey.MasterKey, pl.Password)
	if err != nil {
		return errors.New("failed to encrypt password")
	}

	newAcc := &entities.Account{
		UserID: pl.UserID,
		Service: struct {
			ID   primitive.ObjectID `json:"id,omitempty"`
			Code string             `json:"code"`
			Key  string             `json:"key"`
			Name string             `json:"name"`
		}{
			ID:   svc.ID,
			Code: svc.Attributes.Code,
			Key:  svc.Key,
			Name: svc.Attributes.Name,
		},
		Username:          pl.Username,
		PasswordEncrypted: encPwd.EncodedCipher,
		Salt:              encPwd.EncodedSalt,
		Host:              pl.Host,
		Notes:             pl.Notes,
		IsActive:          *pl.IsActive,
	}

	return u.repo.Upsert(filter, newAcc)
}

func (u *AccountUC) UserAccounts(userID primitive.ObjectID) ([]entities.Account, error) {
	return nil, nil
}
