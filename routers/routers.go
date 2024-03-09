package routers

import (
	"quiz3/controllers"
	"quiz3/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/bangun-datar/segitiga-sama-sisi", controllers.GetSegitiga)
	router.GET("/bangun-datar/persegi", controllers.GetPersegi)
	router.GET("/bangun-datar/persegi-panjang", controllers.GetPersegiPanjang)
	router.GET("/bangun-datar/lingkaran", controllers.GetLingkaran)

	router.GET("/categories", controllers.GetAllCategory)
	router.POST("/categories", middleware.BasicAuth(), controllers.InsertCategory)
	router.PUT("/categories/:id", middleware.BasicAuth(), controllers.UpdateCategory)
	router.DELETE("/categories/:id", middleware.BasicAuth(), controllers.DeleteCategory)
	router.GET("/categories/:id/books", controllers.GetBooksCategory)

	router.GET("/books", controllers.GetAllBooks)
	router.POST("/books", middleware.BasicAuth(), controllers.InsertBooks)
	router.PUT("/books/:id", middleware.BasicAuth(), controllers.UpdateBooks)
	router.DELETE("/books/:id", middleware.BasicAuth(), controllers.DeleteBooks)

	return router
}
