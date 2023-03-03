package service

import (
	logger "github.com/sirupsen/logrus"
)

type Service interface {
}

type service struct {
	repo Repository
}

func NewCommandService(repository Repository) (c Service) {
	if repository == nil {
		logger.Fatal("Repository can not be nil!")
	}
	c = &service{repo: repository}
	return
}

type Repository interface {
}
