package userRepository

import (
	mongoBaseRepository "github.com/ingdeiver/go-core/src/commons/infrastructure/mongo/repository"
	userDomain "github.com/ingdeiver/go-core/src/users/domain"
)

// composition from base repository domain and implements  BaseRepositoryDomain
type UserRepository struct {
    Base *mongoBaseRepository.MongoBaseRepository[userDomain.User]
	//add another compositions here
}

func  New() UserRepository {
	BaseRepo :=  mongoBaseRepository.New[userDomain.User]()
	return  UserRepository{Base: &BaseRepo}
}

func (u UserRepository) List() ([]userDomain.User, error) {
    return u.Base.List()
}

func (u UserRepository) Add(user userDomain.User) (userDomain.User, error) {
    return u.Base.Add(user)
}

func (u UserRepository) Get(ID string) (userDomain.User, error) {
	return u.Base.Get(ID)
}
func (u UserRepository) Remove(ID string) (userDomain.User, error) {
	return u.Base.Remove(ID)
}

/* are extended from this same domain if I add new functions to the base repository domain
type UserRepository userRepositoryDomain.UserBaseRepositoryDomain*/
