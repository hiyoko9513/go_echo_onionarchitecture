package interactor

import (
	"hiyoko-echo/internal/application/usecase"
	"hiyoko-echo/internal/domain/services"
	"hiyoko-echo/internal/infrastructure/database"
	prep "hiyoko-echo/internal/infrastructure/persistence/repository"
	"hiyoko-echo/internal/presentation/http/app/handler"
)

type Interactor interface {
	NewTableRepository() services.TableRepository
	NewUserRepository() services.UserRepository
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

func (i *interactor) NewTableRepository() services.TableRepository {
	return prep.NewTableRepository(i.conn)
}

func (i *interactor) NewUserRepository() services.UserRepository {
	return prep.NewUserRepository(i.conn)
}

func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(i.NewUserRepository())
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUseCase())
}
