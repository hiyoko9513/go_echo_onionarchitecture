package interactor

import (
	drep "hiyoko-echo/domain/repository"
	"hiyoko-echo/infrastructure/database"
	prep "hiyoko-echo/infrastructure/persistence/repository"
	"hiyoko-echo/presenter/http/handler"
	"hiyoko-echo/usecase"
)

type Interactor interface {
	NewTableRepository() drep.TableRepository
	NewUserRepository() drep.UserRepository
	NewUserUseCase() usecase.UserUseCase
	NewUserHandler() handler.UserHandler
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	Conn *database.EntClient
}

func NewInteractor(conn *database.EntClient) Interactor {
	return &interactor{conn}
}

type appHandler struct {
	handler.UserHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.UserHandler = i.NewUserHandler()
	return appHandler
}

func (i *interactor) NewTableRepository() drep.TableRepository {
	return prep.NewTableRepository(i.Conn)
}

func (i *interactor) NewUserRepository() drep.UserRepository {
	return prep.NewUserRepository(i.Conn)
}

func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(i.NewUserRepository())
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUseCase())
}
