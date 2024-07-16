package userRepository

import (
	"errors"

	authDto "github.com/ingdeiver/go-core/src/auth/domain/dto"
	"github.com/ingdeiver/go-core/src/commons/domain/dtos"
	"github.com/ingdeiver/go-core/src/commons/infrastructure/helpers"
	mongoBaseRepository "github.com/ingdeiver/go-core/src/commons/infrastructure/mongo/repository"
	userDomain "github.com/ingdeiver/go-core/src/users/domain"
	userDtos "github.com/ingdeiver/go-core/src/users/domain/dto"
	"go.mongodb.org/mongo-driver/bson"
)

// composition from base repository domain and implements  BaseRepositoryDomain
type UserRepository struct {
	base *mongoBaseRepository.MongoBaseRepository[userDomain.User]
	//add another compositions here
}

func New() *UserRepository {
	BaseRepo := mongoBaseRepository.New[userDomain.User]()
	return &UserRepository{base: &BaseRepo}
}

func (u UserRepository) FindAll(filter any, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto, customPipeline bson.A) (*dtos.PagedResponse[userDomain.User], error) {
	userCustomPipeline := bson.A{
		bson.D{{Key: "$project", Value: bson.M{"password": 0}}},
	}

	userCustomPipeline = append(userCustomPipeline, customPipeline...)
	return u.base.FindAll(filter, pagination, sort, userCustomPipeline)
}

func (u UserRepository) FindAllWithoutPagination(filter any, customPipeline bson.A) ([]userDomain.User, error) {
	userCustomPipeline := bson.A{
		bson.D{{Key: "$project", Value: bson.M{"password": 0}}},
	}

	userCustomPipeline = append(userCustomPipeline, customPipeline...)
	return u.base.FindAllWithoutPagination(filter, userCustomPipeline)
}

func convertToUserDomain(user interface{}) (userDomain.User, error) {
	switch userInfo := user.(type) {
	case userDtos.CreateUserDto:
		return userDomain.User{
			FirstName: userInfo.FirstName,
			LastName:  userInfo.LastName,
			Password:  userInfo.Password,
			Email:     userInfo.Email,
			Role:      userInfo.Role,
		}, nil
	case authDto.RegisterDto:
		return userDomain.User{
			FirstName: userInfo.FirstName,
			LastName:  userInfo.LastName,
			Password:  userInfo.Password,
			Email:     userInfo.Email,
			Role:      userInfo.Role,
		}, nil
	default:
		return userDomain.User{}, errors.New("user conversion failed")
	}
}

func (u UserRepository) Create(user any) (userDomain.User, error) {
	userInfo, err := convertToUserDomain(user)
	if err !=nil {
		return userInfo, err
	}
	hash, err := helpers.CreateHash(userInfo.Password)
	if err != nil {
		return userDomain.User{}, err
	}
	userInfo.Password = hash

	// Set default role
	if userInfo.Role == "" {
		userInfo.Role = userDomain.UserRole
	}

	return u.base.Create(userInfo)
}

func (u UserRepository) UpdateOne(filter interface{}, document any) (*userDomain.User, error) {
	return u.base.UpdateOne(filter, document)
}

func (u UserRepository) UpdateById(ID string, document any) (*userDomain.User, error) {
	return u.base.UpdateById(ID, document)
}

func (u UserRepository) FindById(ID string) (userDomain.User, error) {
	return u.base.FindById(ID)
}

func (u UserRepository) RemoveById(ID string) (userDomain.User, error) {
	return u.base.RemoveById(ID)
}

func (u UserRepository) FindOne(filter interface{}) (*userDomain.User, error) {
	return u.base.FindOne(filter)
}

func (u UserRepository) RemoveOne(filter interface{}) (*userDomain.User, error) {
	return u.base.RemoveOne(filter)
}

/* are extended from this same domain if I add new functions to the base repository domain
type UserRepository userRepositoryDomain.UserBaseRepositoryDomain*/
