package interactor

import (
	"hiyoko-echo/internal/application/usecase"
	"hiyoko-echo/internal/domain/service"
	"hiyoko-echo/internal/infrastructure/database"
	prep "hiyoko-echo/internal/infrastructure/persistence/repository"
	"hiyoko-echo/internal/presentation/http/handler"
)

type Interactor interface {
	NewTableRepository() service.TableRepository
	NewUserRepository() service.UserRepository
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

func (i *interactor) NewTableRepository() service.TableRepository {
	return prep.NewTableRepository(i.conn)
}

func (i *interactor) NewUserRepository() service.UserRepository {
	return prep.NewUserRepository(i.conn)
}

func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(i.NewUserRepository())
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUseCase())
}
