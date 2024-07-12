package util

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s %v", msg, err)
	}
}

func CloseChannel(ch chan error) {
	if _, ok := <-ch; ok {
		close(ch)
	}
}

func DefaultErrorHandler(ctx *fiber.Ctx, err error) error {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	code := http.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	return ctx.Status(code).JSON(fiber.Map{
		"success": false,
		"error":   e.Message,
		"meta":    GetMetaData(ctx),
	})
}

func GetMetaData(ctx *fiber.Ctx) interface{} {
	return fiber.Map{
		"timestamp": time.Now(),
		"path":      ctx.Path(),
		"method":    ctx.Method(),
	}
}
func Response(ctx *fiber.Ctx, code int, data any) error {
	return ctx.Status(code).JSON(fiber.Map{
		"success": true,
		"data":    data,
		"meta":    GetMetaData(ctx),
	})
}

func Ok(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusOK, data)
}

func Created(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusCreated, data)
}
