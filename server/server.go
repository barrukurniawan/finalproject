package server

import (
	"mygram/server/services"
	"net/http"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type Router struct {
	control *services.HandlersController
}

func UserRouther(control *services.HandlersController) *Router {
	return &Router{control: control}
}

func (r *Router) Start(port string) {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome Final Project My Gram"})
	})

	/* List API */

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Users
	router.POST("/users/register", r.control.Register_User)
	router.POST("/users/login", r.control.Login_User)
	router.PUT("/users", r.control.PUT_User)
	router.DELETE("/users", r.control.Delete_User)
	// Photos
	router.POST("/photos", r.control.Post_Photos)
	router.GET("/photos", r.control.Get_Photos)
	router.PUT("/photos/:photoId", r.control.Put_Photos)
	router.DELETE("/photos/:photoId", r.control.Delete_Photos)
	// Comments
	router.POST("/comments", r.control.Comments_Post)
	router.GET("/comments", r.control.Comment_Get)
	router.PUT("/comments/:commentId", r.control.Comment_Put)
	router.DELETE("/comments/:commentId", r.control.Comment_Delete)
	// Social Medias
	router.POST("/socialmedias", r.control.SocialMedias_Post)
	router.GET("/socialmedias", r.control.SocialMedias_Get)
	router.PUT("/socialmedias/:socialMediaId", r.control.SocialMedias_Put)
	router.DELETE("/socialmedias/:socialMediaId", r.control.SocialMedias_Delete)

	router.Run(port)
}
