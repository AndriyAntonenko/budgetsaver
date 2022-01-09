package service

import (
	dto "github.com/AndriyAntonenko/budgetSaver/pkg/dtos"
	"github.com/AndriyAntonenko/budgetSaver/pkg/repository"
)

type TxCategoryService struct {
	txCategoryRepo   *repository.TxCategory
	financeGroupRepo *repository.FinanceGroup
}

func NewTxCategoryService(txCategoryRepo *repository.TxCategory, financeGroupRepo *repository.FinanceGroup) *TxCategoryService {
	return &TxCategoryService{
		txCategoryRepo:   txCategoryRepo,
		financeGroupRepo: financeGroupRepo,
	}
}

func (s *TxCategoryService) CreateTxCategory(userId string, payload *dto.CreateTxCategoryDto) (*dto.TxCategoryDto, *ServiceError) {
	if payload.FinanceGroup != nil {
		userRole, err := (*s.financeGroupRepo).GetUserRoleInFinanceGroup(*payload.FinanceGroup, userId)
		if userRole == "" {
			return nil, NewServiceError(UnknownFinanceGroupMemberError, "user not in this group")
		}

		if err != nil {
			return nil, NewServiceError(UnexpectedError, err.Error())
		}
	}

	txCategoryRepo, err := (*s.txCategoryRepo).AddTxCategory(repository.CreateTxCategoryRecord{
		Name:         payload.Name,
		Creator:      &userId,
		FinanceGroup: payload.FinanceGroup,
	})

	if err != nil {
		return nil, NewServiceError(UnexpectedError, err.Error())
	}

	return &dto.TxCategoryDto{
		Id:           txCategoryRepo.Id,
		Name:         txCategoryRepo.Name,
		Creator:      &txCategoryRepo.Creator.String,
		FinanceGroup: &txCategoryRepo.FinanceGroup.String,
	}, nil
}
