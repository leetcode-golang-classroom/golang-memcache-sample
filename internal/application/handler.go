package application

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leetcode-golang-classroom/golang-memcache-sample/internal/util"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) HealthCheck(ctx *fiber.Ctx) error {
	return util.Ok(ctx, fiber.Map{
		"message": "ok",
	})
}
