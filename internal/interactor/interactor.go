package interactor

import (
	"hiyoko-echo/internal/application/usecase"
	"hiyoko-echo/internal/domain/services"
	"hiyoko-echo/internal/infrastructure/database"
	"hiyoko-echo/internal/infrastructure/persistence/repository"
	"hiyoko-echo/internal/presentation/http/app/handler"
	"hiyoko-echo/internal/presentation/http/app/oapi"
)

type Interactor interface {
	NewTableRepository() services.TableRepository
	NewUserRepository() services.UserRepository
	NewUserUseCase() usecase.UserUseCase
	NewUserHandler() handler.UserHandler
	NewAppHandler() oapi.ServerInterface
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

func (i *interactor) NewAppHandler() oapi.ServerInterface {
	appHandler := &appHandler{}
	appHandler.UserHandler = i.NewUserHandler()
	return appHandler
}

func (i *interactor) NewTableRepository() services.TableRepository {
	return repository.NewTableRepository(i.conn)
}

func (i *interactor) NewUserRepository() services.UserRepository {
	return repository.NewUserRepository(i.conn)
}

func (i *interactor) NewUserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(i.NewUserRepository())
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserUseCase())
}
