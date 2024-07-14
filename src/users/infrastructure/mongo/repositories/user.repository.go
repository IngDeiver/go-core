package userRepository

import (
	"errors"

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

func  New() *UserRepository {
	BaseRepo :=  mongoBaseRepository.New[userDomain.User]()
	return  &UserRepository{base: &BaseRepo}
}


func (u UserRepository) FindAll(filter any, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto, customPipeline bson.A) (*dtos.PagedResponse[userDomain.User], error) {
	userCustomPipeline := bson.A{
		bson.D{{Key: "$project",Value:  bson.M{"password": 0}}},
	}

	userCustomPipeline = append(userCustomPipeline,customPipeline...)
	return u.base.FindAll(filter, pagination, sort, userCustomPipeline)
}

func (u UserRepository) Create(user any) (userDomain.User, error) {
	userInfo, ok := user.(userDtos.CreateUserDto)
	if !ok {
		return userDomain.User{}, errors.New("user convertion fail")
	}

	hash, err :=helpers.CreateHash(userInfo.Password )
	if err != nil {
		return userDomain.User{}, err
	}
	userInfo.Password = hash 
    return u.base.Create(userInfo)
}

func (u  UserRepository) UpdateById(ID string, document any) (*userDomain.User, error){
	return u.base.UpdateById(ID, document)
}

func (u UserRepository) FindById(ID string) (userDomain.User, error) {
	return u.base.FindById(ID)
}
func (u UserRepository) RemoveById(ID string) (userDomain.User, error) {
	return u.base.RemoveById(ID)
}

/* are extended from this same domain if I add new functions to the base repository domain
type UserRepository userRepositoryDomain.UserBaseRepositoryDomain*/
