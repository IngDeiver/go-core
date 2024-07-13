package userRepository

import (
	"github.com/ingdeiver/go-core/src/commons/domain/dtos"
	mongoBaseRepository "github.com/ingdeiver/go-core/src/commons/infrastructure/mongo/repository"
	userDomain "github.com/ingdeiver/go-core/src/users/domain"
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


func (u UserRepository) FindAll(filter any, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto) (*dtos.PagedResponse[userDomain.User], error) {
	return u.base.FindAll(filter, pagination, sort)
}

func (u UserRepository) Create(user userDomain.User) (userDomain.User, error) {
    return u.base.Create(user)
}

func (u UserRepository) FindById(ID string) (userDomain.User, error) {
	return u.base.FindById(ID)
}
func (u UserRepository) RemoveById(ID string) (userDomain.User, error) {
	return u.base.RemoveById(ID)
}

/* are extended from this same domain if I add new functions to the base repository domain
type UserRepository userRepositoryDomain.UserBaseRepositoryDomain*/
