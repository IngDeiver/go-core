package commonMongoRepository

// implements BaseRepositoryDomain
type  MongoBaseRepository[T any] struct {
	
}

func  New[T any]()  MongoBaseRepository[T] {
	return  MongoBaseRepository[T]{}
}

func (s  MongoBaseRepository[T]) List() ([]T, error) {
	return []T{}, nil
}

func (s  MongoBaseRepository[T]) Add(user T) (T, error) {
	var result T
	return result, nil
}
func (s  MongoBaseRepository[T]) Get(ID string) (T, error) {
	var result T
	return result, nil
}
func (s  MongoBaseRepository[T]) Remove(ID string) (T, error) {
	var result T
	return result, nil
}
