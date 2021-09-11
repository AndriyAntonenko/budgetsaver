package service

import (
	"github.com/AndriyAntonenko/budgetSaver/pkg/domain"
	"github.com/AndriyAntonenko/budgetSaver/pkg/repository"
)

type GroupService struct {
	groupRepo repository.Group
}

func NewGroupService(groupRepo repository.Group) *GroupService {
	return &GroupService{groupRepo}
}

func (gs *GroupService) CreateGroup(userId string, payload *domain.CreateGroupPayload) (*domain.Group, error) {
	return gs.groupRepo.CreateGroup(userId, payload)
}
