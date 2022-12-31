package handler

import (
	"database/sql"

	"github.com/be3/go_vue_todo/server/internal/gateway"
	"github.com/be3/go_vue_todo/server/internal/middleware"
	"github.com/be3/go_vue_todo/server/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) *gin.Engine {
	router.Use(middleware.Cors())                                  // CORSの設定
	router.Use(middleware.SetSessionCookie("secret", "mysession")) // クッキーに認証キーを作成

	group := router.Group("/api")
	{
		var db *sql.DB
		userRepository := gateway.NewUserRepository(db)

		signupHandler := SignupHandler{
			Usecase: usecase.NewSignupUsecase(userRepository),
		}
		signinHandler := SigninHandler{
			Usecase: usecase.NewSigninUsecase(userRepository),
		}

		group.POST("/signup", signupHandler.Func)
		group.POST("/signin", signinHandler.Func)

		// ミドルウェアによる認証時のみ利用可能なハンドラ
		authGroup := group.Group("/auth", middleware.Authenticate())
		{
			taskRepository := gateway.NewTaskRepository(db)

			getTaskHandler := GetTaskHandler{
				Usecase: usecase.NewGetTaskUsecase(taskRepository),
			}
			getUserTasksHandler := GetUserTasksHandler{
				getUserTasks: usecase.NewGetUserTasksUsecase(taskRepository),
				getAuthUser:  usecase.NewGetAuthUserUsecase(userRepository),
			}
			postTaskHandler := PostTaskHandler{
				getAuthUser: usecase.NewGetAuthUserUsecase(userRepository),
				createTask:  usecase.NewCreateTaskUsecase(taskRepository),
			}
			putTaskHandler := PutTaskHandler{
				Usecase: usecase.NewUpdateTaskUsecase(taskRepository),
			}
			deleteTaskHandler := DeleteTaskHandler{
				Usecase: usecase.NewDeleteTaskUsecase(taskRepository),
			}
			signoutHandler := SignoutHandler{
				Usecase: usecase.NewSignoutUsecase(),
			}

			authGroup.GET("/tasks/:id", getTaskHandler.Func)
			authGroup.GET("/tasks", getUserTasksHandler.Func)
			authGroup.POST("/tasks", postTaskHandler.Func)
			authGroup.GET("/signout", signoutHandler.Func)
			authGroup.PUT("/tasks/:id", putTaskHandler.Func)
			authGroup.DELETE("/tasks/:id", deleteTaskHandler.Func)
		}
	}

	return router
}
