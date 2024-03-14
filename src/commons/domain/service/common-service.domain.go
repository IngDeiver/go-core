package commonServiceDomain

type ServiceDomain[T any] interface {
	List() ([]T, error)
	Add(document T) (T, error)
	Get(ID string) (T, error)
	Remove(ID string) (T, error)
}
