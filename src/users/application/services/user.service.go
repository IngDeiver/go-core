package userService

import (
	baseService "github.com/ingdeiver/go-core/src/commons/application/services/base"
	"github.com/ingdeiver/go-core/src/commons/domain/dtos"
	errorsDomain "github.com/ingdeiver/go-core/src/commons/domain/errors"
	userDomain "github.com/ingdeiver/go-core/src/users/domain"
	userDtos "github.com/ingdeiver/go-core/src/users/domain/dto"
	userRepo "github.com/ingdeiver/go-core/src/users/infrastructure/mongo/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// composition from base service domain adb implements BaseServiceDomain
type UserService struct {
	base *baseService.BaseService[userDomain.User]
	userRepo *userRepo.UserRepository
	//add another compositions here
}

func New(repository *userRepo.UserRepository) *UserService {
	BaseService := baseService.New[userDomain.User](repository)
	return &UserService{base: &BaseService, userRepo: repository}
}

func (s *UserService) FindAll(filter userDtos.UserFilterDto, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto) (*dtos.PagedResponse[userDomain.User], error) {
	return s.base.FindAll(filter, pagination, sort)
}

func (s *UserService) FindAllWithoutPagination(filter any) ([]userDomain.User, error) {
	return s.base.FindAllWithoutPagination(filter)
}

func (s *UserService) Create(data userDtos.CreateUserDto) (userDomain.User, error) {

	existUser, err := s.base.FindOne(bson.M{"email": data.Email})
	if err != nil && err != errorsDomain.ErrNotFoundError {
		return userDomain.User{}, err
	}

	if existUser != nil {
		return *existUser, errorsDomain.ErrUserAlreadyExistsError
	}
	return s.base.Create(data)
}

func (s *UserService) UpdateOne(filter interface{}, document userDtos.UpdateUserDto) (*userDomain.User, error) {
	return s.base.UpdateOne(filter, document)
}

func (s *UserService) UpdatePasswordById(ID string, password string) (error) {
	return s.userRepo.UpdatePasswordById(ID, password)
}

func (s *UserService) UpdateById(ID string, document userDtos.UpdateUserDto) (*userDomain.User, error) {
	if document.Email != "" {
		objectId, err := primitive.ObjectIDFromHex(ID)
		if err != nil {
			return nil, err
		}
		existUser, err := s.base.FindOne(bson.M{"email": document.Email, "_id": bson.M{"$ne": objectId}})
		if err != nil && err != errorsDomain.ErrNotFoundError {
			return nil, err
		}

		if existUser != nil {
			return existUser, errorsDomain.ErrUserAlreadyExistsError
		}
	}
	return s.base.UpdateById(ID, document)
}

func (s *UserService) FindById(ID string) (userDomain.User, error) {
	return s.base.FindById(ID)
}

func (s *UserService) FindOne(filter interface{}) (*userDomain.User, error) {
	return s.base.FindOne(filter)
}

func (s *UserService) RemoveOne(filter interface{}) (*userDomain.User, error) {
	return s.base.RemoveOne(filter)
}

func (s *UserService) RemoveById(ID string) (userDomain.User, error) {
	return s.base.RemoveById(ID)
}
