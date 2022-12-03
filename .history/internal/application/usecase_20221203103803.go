package domain

type UseCase[I, O comparable] interface {
	Execute(*I) (*O, error)
}
