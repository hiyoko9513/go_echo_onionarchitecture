package interactor

import (
	drep "hiyoko-echo/internal/domain/repository"
	"hiyoko-echo/internal/infrastructure/database"
	prep "hiyoko-echo/internal/infrastructure/persistence/repository"
	"hiyoko-echo/internal/presenter/http/handler"
	"hiyoko-echo/internal/usecase"
)

type Interactor interface {
	NewTableRepository() drep.TableRepository
	NewUserRepository() drep.UserRepository
	NewUserUseCase() usecase.UserUseCase
	NewUserHandler() handler.UserHandler
	NewAppHandler() handler.AppHandler
}

type interactor struct {
	conn *database.EntClient
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
	return prep.NewTableRepository(i.conn)
}

func (i *interactor) NewUserRepository() drep.UserRepository {
	return prep.NewUserRepository(i.conn)
}

func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(i.NewUserRepository())
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUseCase())
}
