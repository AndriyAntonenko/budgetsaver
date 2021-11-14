package service

import (
	"fmt"

	dto "github.com/AndriyAntonenko/budgetSaver/pkg/dtos"
	"github.com/AndriyAntonenko/budgetSaver/pkg/repository"
)

type FinanceGroupService struct {
	repo repository.FinanceGroup
}

func NewFinanceGroupService(repo repository.FinanceGroup) *FinanceGroupService {
	return &FinanceGroupService{repo}
}

func (s *FinanceGroupService) CreateFinanceGroup(userId string, payload dto.CreateFinanceGroupPayload) (*dto.FinanceGroup, error) {
	record, err := s.repo.CreateFinanceGroup(userId, repository.CreateFinanceGroupRecord{
		Name:        payload.Name,
		Description: payload.Description,
	})

	if err != nil {
		return nil, fmt.Errorf("database error: %s", err.Error())
	}

	return &dto.FinanceGroup{
		Id:           record.Id,
		Name:         record.Name,
		Description:  record.Description,
		MembersCount: 1,
	}, nil
}

func (s *FinanceGroupService) GetUsersFinanceGroups(userId string) ([]dto.FinanceGroup, error) {
	records, err := s.repo.GetUsersFinanceGroups(userId)
	if err != nil {
		return nil, fmt.Errorf("database error: %s", err.Error())
	}

	groups := make([]dto.FinanceGroup, 0)
	for _, rec := range records {
		groups = append(groups, dto.FinanceGroup{
			Id:           rec.Id,
			Name:         rec.Name,
			Description:  rec.Description,
			MembersCount: rec.MemberCount,
		})
	}

	return groups, nil
}
