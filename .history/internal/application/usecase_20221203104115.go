package application

type UseCase[I, O comparable] interface {
	Execute(*I) (*O, error)
}
