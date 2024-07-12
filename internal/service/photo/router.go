package photo

import "github.com/gofiber/fiber/v2"

type Router struct {
	photoHdr *Handler
}

func NewRouter(photoHdr *Handler) *Router {
	return &Router{
		photoHdr: photoHdr,
	}
}

func (route *Router) SetupRoutes(router fiber.Router) {
	router.Get("/photo/:id", route.photoHdr.VerifyCache, route.photoHdr.GetPhoto)
}
