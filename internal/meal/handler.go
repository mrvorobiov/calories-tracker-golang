package meal

import (
	"github.com/gofiber/fiber/v3"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h Handler) InitRoutes(app *fiber.App) {
	app.Get("/meals", h.GetAll)
	app.Post("/meals", h.Create)
	app.Delete("/meals/:id", h.Delete)
	app.Patch("/meals:id", h.Update)
}

func (h Handler) GetAll(ctx fiber.Ctx) error {
	meals, err := h.service.GetAll()
	if err != nil {
		return err
	}
	return ctx.JSON(meals)
}

func (h Handler) Create(ctx fiber.Ctx) error {
	var meal Meal
	if err := ctx.Bind().Body(&meal); err != nil {
		return err
	}
	if err := h.service.Track(meal); err != nil {
		return err
	}
	return ctx.JSON(meal)
}

func (h Handler) Delete(ctx fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}
	if err := h.service.Delete(uint(id)); err != nil {
		return err
	}
	return ctx.JSON(true)
}

func (h Handler) Update(ctx fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return err
	}
	var body Meal
	if err := ctx.Bind().Body(&body); err != nil {
		return err
	}
	body.Id = uint(id)
	if err := h.service.Update(body); err != nil {
		return err
	}
	return ctx.JSON(body)
}
