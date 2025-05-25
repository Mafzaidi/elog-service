package usecase

import (
	"fmt"

	"github.com/mafzaidi/elog/internal/account"
	"github.com/mafzaidi/elog/internal/entities"
	"github.com/mafzaidi/elog/internal/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceUC struct {
	repo        service.Repository
	accountRepo account.Repository
}

func NewServiceUseCase(repo service.Repository, accountRepo account.Repository) service.UseCase {
	return &ServiceUC{
		repo:        repo,
		accountRepo: accountRepo,
	}
}

func (u *ServiceUC) ServicesHasAccount(userID primitive.ObjectID, isActive *bool) ([]entities.MasterData, error) {
	accountFilter := bson.M{
		"userID": userID,
	}
	if isActive != nil {
		accountFilter["isActive"] = *isActive
	}

	accounts, err := u.accountRepo.FindManyByFilter(accountFilter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch accounts: %w", err)
	}

	serviceIDMap := make(map[primitive.ObjectID]bool)
	for _, acc := range accounts {
		if !acc.Service.ID.IsZero() {
			serviceIDMap[acc.Service.ID] = true
		}
	}

	if len(serviceIDMap) == 0 {
		return []entities.MasterData{}, nil
	}

	var serviceIDs []primitive.ObjectID
	for id := range serviceIDMap {
		serviceIDs = append(serviceIDs, id)
	}

	serviceFilter := bson.M{
		"_id": bson.M{"$in": serviceIDs},
	}
	services, err := u.repo.FindManyByFilter(serviceFilter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch services by ID: %w", err)
	}

	return services, nil

}
