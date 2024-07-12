package photo

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/leetcode-golang-classroom/golang-memcache-sample/internal/types"
)

func (handler *Handler) VerifyCache(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id not provide")
	}
	val, err := handler.cache.Get(id)
	if err != nil {
		return ctx.Next()
	}
	var photo types.Photo
	err = json.Unmarshal(val.Value, &photo)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to unmarshal")
	}
	return ctx.JSON(fiber.Map{
		"Data":     photo,
		"IsCached": true,
	})
}
