package services

import (
	"errors"
	"go-echo/entity"
	"go-echo/repository"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	GetAll() ([]*entity.Biodata, error)
	FindByID(ID int) (*entity.Biodata, error)
	Create(biodataRequest *entity.Request) error
	Update(ID int, biodataUpdate *entity.Update) (*entity.Biodata, error)
	Delete(ID int) error
}

type service struct {
	repository repository.Repository
	validator  *validator.Validate
}

func NewService(repository repository.Repository) *service {
	return &service{
		repository: repository,
		validator:  validator.New(),
	}
}

func (s *service) GetAll() ([]*entity.Biodata, error) {
	biodatas, err := s.repository.GetAll()
	return biodatas, err
}

func (s *service) FindByID(ID int) (*entity.Biodata, error) {
	biodata, err := s.repository.FindByID(ID)
	return biodata, err
}

func (s *service) Create(biodataRequest *entity.Request) error {
	err := s.validator.Struct(biodataRequest)
	if err != nil {
		return err
	}
	biodata := &entity.Biodata{
		NAME:    biodataRequest.NAME,
		AGE:     biodataRequest.AGE,
		ADDRESS: biodataRequest.ADDRESS,
	}

	err = s.repository.Create(biodata)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Update(ID int, biodataUpdate *entity.Update) (*entity.Biodata, error) {
	biodata, err := s.repository.FindByID(ID)
	if err != nil {
		return nil, err
	}

	if biodata == nil {
		return nil, errors.New("Biodata not found")
	}

	err = biodataUpdate.ValidateUpdate()
	if err != nil {
		return nil, err
	}

	biodata.NAME = biodataUpdate.NAME
	biodata.AGE = biodataUpdate.AGE
	biodata.ADDRESS = biodataUpdate.ADDRESS

	updatedBiodata, err := s.repository.Update(biodata)
	if err != nil {
		return nil, err
	}
	return updatedBiodata, nil
}

func (s *service) Delete(ID int) error {
	biodata, err := s.repository.FindByID(ID)
	if err != nil {
		return err
	}

	if biodata == nil {
		return errors.New("Biodata not found")
	}

	err = s.repository.Delete(biodata)
	if err != nil {
		return err
	}

	return nil
}
