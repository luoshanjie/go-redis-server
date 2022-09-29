package domain

//go:generate mockery --name=BaseService
type BaseService interface {
	Run() error
}
