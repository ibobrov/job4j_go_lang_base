package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"job4j.ru/go-lang-base/internal/domain"
)

type CreateItemRequest struct {
	Name string `json:"name"`
}

type CreateItemResponse struct {
	Item domain.Item `json:"newItem"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"name is required"`
}

// CreateItem godoc
//
// @Summary Создать задачу
// @Description Создает новую задачу
// @Tags Items
// @Accept json
// @Produce json
// @Param request body CreateItemRequest true "Данные задачи"
// @Success 201 {object} CreateItemResponse
// @Failure 400 {object} ErrorResponse "Некорректный запрос"
// @Router /item [post]
func (s *Server) CreateItem(c *fiber.Ctx) error {
	var req CreateItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if err := req.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	item, err := s.Repository.Create(c.Context(), domain.Item{
		ID:   uuid.New().String(),
		Name: req.Name,
	})
	if err != nil {
		log.Errorw("s.Repository.Create", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusCreated).JSON(CreateItemResponse{Item: item})
}

func (request CreateItemRequest) Validate() error {
	if len(request.Name) < 3 || len(request.Name) > 100 {
		return errors.New("name must be between 3 and 100 characters")
	}
	return nil
}
