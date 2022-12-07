package handler

import (
	"traveland/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}
func NewHandler(service *service.Service)*Handler{
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRountes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		guide := api.Group("/guide")
		{
			guide.POST("/add-guide", h.addGuide)
			guide.DELETE("/delete-guide/:id", h.delteGuide)
			guide.PUT("/update-guide/:id", h.updateGuide)
			guide.GET("/get-guide/:id", h.getGuideByID)

			guide.GET("/get-all-guide/", h.getAllGuide)
		}
		place := api.Group("/place")
		{
			place.POST("/add-place/", h.addPlace)
			place.DELETE("/delete-place/:id", h.deltePlace)
			place.PUT("/update-place/:id", h.updatePlace)
			place.GET("/get-place/:id", h.getPlaceByID)

			place.GET("/get-all-place/", h.getAllPlace)
			place.GET("/get-place-by-type/:type-id", h.getPlaceByType)
		}
	}
	return router
}