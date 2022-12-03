package application

type UseCase[I, O any] interface {
	Execute(*I) (*O, error)
}
