package http

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"time"
)

// WorldStateRequest - запрос с фронта на начальное состояние игрового мира
type WorldStateRequest struct {
	Data struct {
		UserId int32 `json:"user_id"`
	} `json:"data"` // отправляемые данные с фронта
}

type InitWorldStateResponse struct {
	SessionId int32     `json:"session_id"`
	UserId    int32     `json:"user_id"`
	AreaId    int32     `json:"area_id"`
	Neutrals  []byte    `json:"neutrals"`
	Buildings []byte    `json:"buildings"`
	Heroes    []byte    `json:"heroes"`
	Units     []byte    `json:"units"`
	Enemies   []byte    `json:"enemies"`
	Timestamp time.Time `json:"timestamp"`
}

func GetInitWorldState(ctx fiber.Ctx) error {
	// Получаем тело запроса в виде сырых данных
	body := ctx.Body()

	// Парсим JSON из тела запроса
	var req WorldStateRequest
	if err := json.Unmarshal(body, &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request format",
		})
	}

	// Получение user_id из запроса
	userId := req.Data.UserId

	// Формирование мок-ответа
	response := InitWorldStateResponse{
		SessionId: 1,
		UserId:    userId,
		AreaId:    456,
		Neutrals: []byte(`[
			{
				"id": 101,
				"name": "Forest",
				"coordinates": {"q": 1, "r": 0, "s": -1},
				"capacity": 1000.00
			}
		]`),
		Buildings: []byte(`[
			{
		      "id": 201,
		      "name": "Town Hall",
		      "coordinates": {"q": 0, "r": -1, "s": 1},
		      "level": 3,
		      "hp": 4000
		    }
		]`),
		Heroes: []byte(`[
			{
		      "id": 301,
		      "name": "Knight",
		      "coordinates": {"q": -1, "r": 1, "s": 0},
		      "hp": 150,
		      "experience": 50,
		      "level": 1
		    }
		]`),
		// и остальные сущности
		Timestamp: time.Now(),
	}

	// Возврат ответа
	return ctx.Status(fiber.StatusOK).JSON(response)
}
