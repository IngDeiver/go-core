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


func (u UserRepository) List(filter any, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto) (*dtos.PagedResponse[userDomain.User], error) {
	return u.base.List(filter, pagination, sort)
}

func (u UserRepository) Add(user userDomain.User) (userDomain.User, error) {
    return u.base.Add(user)
}

func (u UserRepository) Get(ID string) (userDomain.User, error) {
	return u.base.Get(ID)
}
func (u UserRepository) Remove(ID string) (userDomain.User, error) {
	return u.base.Remove(ID)
}

/* are extended from this same domain if I add new functions to the base repository domain
type UserRepository userRepositoryDomain.UserBaseRepositoryDomain*/
