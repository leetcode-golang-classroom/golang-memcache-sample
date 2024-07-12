package photo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/leetcode-golang-classroom/golang-memcache-sample/internal/types"
)

type Handler struct {
	jsonServerURL string
	cache         types.MemCache
	client        *http.Client
	sync.RWMutex
}

func NewHandler(jsonServerURL string, cache types.MemCache) *Handler {
	return &Handler{
		jsonServerURL: jsonServerURL,
		cache:         cache,
		client:        &http.Client{},
	}
}
func (handler *Handler) GetPhoto(ctx *fiber.Ctx) error {
	handler.Lock()
	defer handler.Unlock()
	// GET id from param
	id := ctx.Params("id")
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id not provide")
	}
	// make request with jsonServerURL
	timeout, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()
	// format url
	requestURL := fmt.Sprintf("%s/%s", handler.jsonServerURL, id)
	req, err := http.NewRequestWithContext(timeout, http.MethodGet, requestURL, nil)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	res, err := handler.client.Do(req)
	if err != nil {
		if errors.Is(err, http.ErrHandlerTimeout) {
			return fiber.NewError(fiber.ErrRequestTimeout.Code, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer res.Body.Close()
	var photo types.Photo
	err = json.NewDecoder(res.Body).Decode(&photo)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Errorf("marshal failed %w", err).Error())
	}
	resultByte, err := json.Marshal(photo)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Errorf("marshal failed %w", err).Error())
	}
	err = handler.cache.Set(id, resultByte)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(fiber.Map{
		"Data":     photo,
		"IsCached": false,
	})
}
