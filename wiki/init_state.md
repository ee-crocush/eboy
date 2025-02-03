# Контракт на получение начального состояния игры

## REST ручки

### Получение начального состояния игры

**Ендпоинт**: GET /state

Request: Frontend -> Backend
```json
{
  "user_id": 456
}
```

Response: Backend -> Frontend
```json
{
  "session_id": 1,
  "user_id": 123,
  "area_id": 456,
  "neutrals": [
    {
      "id": 101,
      "name": "Forest",
      "coordinates": {"q": 1, "r": 0, "s": -1},
      "capacity": 1000,
      "current_amount": 500
    }
  ],
  "buildings": [
    {
      "id": 201,
      "name": "Town Hall",
      "coordinates": {"q": 0, "r": -1, "s": 1},
      "level": 3,
      "hp": 4000
    }
  ],
  "heroes": [
    {
      "id": 301,
      "name": "Knight",
      "coordinates": {"q": -1, "r": 1, "s": 0},
      "hp": 150,
      "experience": 50,
      "level": 1
    }
  ],
  "timestamp": "2023-10-01T12:00:00Z"
}
```


## WebSocket контракты

### Запрос текущего состояния мира

**Канал**: get_world_state

Request: Frontend -> Backend
```json
{
  "type": "get_world_state",
  "data": {
    "user_id": 123
  }
}
```

Response: Backend -> Frontend
```json
{
  "type": "world_state",
  "data": {
    "session_id": 1,
    "user_id": 123,
    "area_id": 456,
    "actions": [
      {
        "id": 789,
        "type": "move",
        "status": "in_progress"
      }
    ],
    "timestamp": "2023-10-01T12:00:00Z"
  }
}
```