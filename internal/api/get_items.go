package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type ItemRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetItemsResponse struct {
	Items []ItemRequest `json:"items"`
}

// GetItems godoc
//
// @Summary Получить список задач
// @Description Возвращает все задачи
// @Tags Items
// @Produce json
// @Success 200 {object} GetItemsResponse
// @Router /items [get]
func (s *Server) GetItems(c *fiber.Ctx) error {
	items, err := s.Repository.List(c.Context())
	if err != nil {
		log.Errorw("s.Repository.List", "failed to get items from repository", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	res := make([]ItemRequest, 0, len(items))
	for _, item := range items {
		res = append(res, ItemRequest{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	return c.Status(fiber.StatusOK).JSON(GetItemsResponse{Items: res})
}
