package meal

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
	"time"
)

type Kind int

const (
	BREAKFAST Kind = iota
	LUNCH
	DINNER
	SLACK
)

const (
	ErrCaloriesAreNotPositiveInteger = "calories have to be positive integer"
)

type Meal struct {
	Id            uint      `json:"int" gorm:"primaryKey"`
	Kind          Kind      `json:"kind"`
	Title         string    `json:"title"`
	Calories      int       `json:"calories"`
	Fats          *int      `json:"fats"`
	Proteins      *int      `json:"proteins"`
	Carbohydrates *int      `json:"carbohydrates"`
	LoggedOn      time.Time `json:"loggedOn" gorm:"default:CURRENT_TIMESTAMP"`
}

func InitRouter(db *gorm.DB, app *fiber.App) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)
	handler.InitRoutes(app)
}
