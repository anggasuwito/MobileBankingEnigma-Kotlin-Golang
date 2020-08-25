package master

import (
	"banking/master/controller"
	"banking/master/repositories"
	"banking/master/usecase"
	"database/sql"

	"github.com/gorilla/mux"
)

func InitAll(R *mux.Router, DB *sql.DB) {
	TrancationRepo := repositories.InitTransactionRepoImpl(DB)
	TrancationUseCase := usecase.InitTransactionUseCase(TrancationRepo)
	controller.TransactionController(R, TrancationUseCase)
	UserRepo := repositories.InitUserRepoImpl(DB)
	UserUseCase := usecase.InitUserUseCase(UserRepo)
	controller.UserController(R, UserUseCase)
}
