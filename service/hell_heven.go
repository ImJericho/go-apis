package service

import (
	"apis/dao"
	"apis/model"
)

type Service struct {
	dao dao.MongoDao
}

func NewService(dao dao.MongoDao) Service {
	return Service{
		dao: dao,
	}
}

func (s Service) GetAllUsers() ([]model.MainDB, error) {
	return s.dao.GetAllUsers()
}

func (s Service) GetUserByname(name string) (model.MainDB, error) {
	return s.dao.GetUserByName(name)
}

func (s Service) CreateUser(user model.MainDB) error {
	_, err := s.dao.CreateUser(user)
	return err
}

func (s Service) DeleteUser(name string) error {
	_, err := s.dao.DeleteUser(name)
	return err
}

func (s Service) IncHell(name string) error {
	user, err := s.dao.GetUserByName(name)
	if err != nil {
		return err
	}
	user.HellPoints += 50
	_, err = s.dao.UpdateUser(user.Name, user)
	return err
}

func (s Service) IncHeaven(name string) error {
	user, err := s.dao.GetUserByName(name)
	if err != nil {
		return err
	}
	user.HevenPoints += 50
	_, err = s.dao.UpdateUser(user.Name, user)
	return err
}

func (s Service) DecHell(name string) error {
	user, err := s.dao.GetUserByName(name)
	if err != nil {
		return err
	}
	user.HellPoints -= 50
	_, err = s.dao.UpdateUser(user.Name, user)
	return err
}

func (s Service) DecHeaven(name string) error {
	user, err := s.dao.GetUserByName(name)
	if err != nil {
		return err
	}
	user.HevenPoints -= 50
	_, err = s.dao.UpdateUser(user.Name, user)
	return err
}
