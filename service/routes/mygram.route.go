package routes

import (
	"final-project/connection"
	"final-project/middleware"
	"final-project/service/mygram/handler"
	"final-project/service/mygram/repository"
	"final-project/service/mygram/usecase"

	"github.com/gin-gonic/gin"
)

type myGramRoutes struct {
	Handler handler.MyGramHandler
	Router  *gin.RouterGroup
}

func InitMygramRoute(
	router *gin.RouterGroup, gormDb *connection.GormDB,
) *myGramRoutes {
	handler := handler.InitMyGramHandler(
		usecase.InitMyGramUsecase(
			repository.InitRequestOrderRepository(
				gormDb,
			),
			gormDb,
		),
	)
	return &myGramRoutes{
		Handler: *handler,
		Router:  router,
	}
}

func (r *myGramRoutes) Routes() {
	router := r.Router.Group("/user")
	router.POST("/register", r.Handler.Register)
	router.POST("/login", r.Handler.Login)

	photoRouter := r.Router.Group("/photo")
	photoRouter.Use(middleware.Authentication())
	photoRouter.POST("/createPhoto", r.Handler.PostPhoto)
	photoRouter.PUT("/updatePhoto/:id", r.Handler.UpdatePhoto)
	photoRouter.DELETE("/deletePhoto/:id", r.Handler.DeletePhoto)
	photoRouter.GET("/getAll", r.Handler.GetAllPhoto)
	photoRouter.GET("/get/:id", r.Handler.GetOnePhoto)

	sosmedRouter := r.Router.Group("/sosmed")
	sosmedRouter.Use(middleware.Authentication())
	sosmedRouter.POST("/createSocialMedia", r.Handler.PostSosmed)
	sosmedRouter.PUT("/updateSocialMedia/:id", r.Handler.UpdateSosmed)
	sosmedRouter.DELETE("/deleteSocialMedia/:id", r.Handler.DeleteSosmed)
	sosmedRouter.GET("/getAll", r.Handler.GetAllSosmed)
	sosmedRouter.GET("/get/:id", r.Handler.GetOneSosmed)

	commentRouter := r.Router.Group("/comment")
	commentRouter.Use(middleware.Authentication())
	commentRouter.POST("/createComment", r.Handler.PostComment)
	commentRouter.PUT("/updateComment/:id", r.Handler.UpdateComment)
	commentRouter.DELETE("/deleteComment/:id", r.Handler.DeleteComment)
	commentRouter.GET("/getAll", r.Handler.GetAllComment)
	commentRouter.GET("/get/:id", r.Handler.GetOneComment)
}
