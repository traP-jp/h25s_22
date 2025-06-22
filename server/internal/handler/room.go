package handler

import (
	"backend/internal/repository"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"googlemaps.github.io/maps"
)

type (
	TimeOption struct {
		StartTime time.Time `json:"start_time"` // 開始時間
		EndTime   time.Time `json:"end_time"`   // 終了時間
	}
	CreateRoomRequest struct {
		Name          string       `json:"name"`            // ルーム名
		TimeOptions   []TimeOption `json:"time_options"`    // 時間候補
		PlaceOptions  []string     `json:"place_options"`   // 場所候補（GooglePlaceIDの配列）
		CenterPlaceID string       `json:"center_place_id"` // 中心位置(GooglePlaceID)
		Radius        int          `json:"radius"`          // 半径
		PlaceMax      int          `json:"place_max"`       // 最大値
	}

	ReturnTimeOption struct {
		ID        uuid.UUID `json:"id"`         // 時間候補のID
		StartTime time.Time `json:"start_time"` // 開始時間
		EndTime   time.Time `json:"end_time"`   // 終了時間
	}
	ReturnPlaceOption struct {
		ID             uuid.UUID   `json:"id"`              // 場所候補のID
		GooglePlaceID  string      `json:"google_place_id"` // Google Place ID
		Name           string      `json:"name"`
		Location       maps.LatLng `json:"location"`
		PhotoReference string      `json:"photoReference"`
		PriceLevel     int         `json:"priceLevel"`
		Rating         float32     `json:"rating"`
		Address        string      `json:"address"`
	}

	CreateRoomResponse struct {
		RoomID       uuid.UUID           `json:"room_id"`       // 作成されたルームのID
		TimeOptions  []ReturnTimeOption  `json:"time_options"`  // 作成された時間候補
		PlaceOptions []ReturnPlaceOption `json:"place_options"` // 作成された場所候補
	}

	GetRoomResponse struct {
		RoomID       uuid.UUID           `json:"room_id"`       // ルームのID
		Name		 string              `json:"name"`          // ルーム名
		CenterPoint  string              `json:"center_point"`  // 中心位置のGoogle Place ID
		Radius       int                 `json:"radius"`        // 半径
		PlaceMax     int                 `json:"place_max"`     // 最大値
		TimeOptions  []ReturnTimeOption  `json:"time_options"`  // 時間候補
		PlaceOptions []ReturnPlaceOption `json:"place_options"` // 場所候補
	}
)

func (h *Handler) GetClient(c echo.Context) (*maps.Client, error) {
	api_key, ok := os.LookupEnv("GOOGLE_API_KEY")
	if !ok {
		return nil, c.String(http.StatusInternalServerError, "Not Found APIKey")
	}
	client, err := maps.NewClient(maps.WithAPIKey(api_key))
	if err != nil {
		return nil, c.String(http.StatusInternalServerError, "Failed to create Google Maps client")
	}
	return client, nil
}

func (h *Handler) GetPlaceDetail(c echo.Context, client maps.Client, googlePlaceID string) (GetSearch, error) {
	r := &maps.PlaceDetailsRequest{
		PlaceID: googlePlaceID,
	}
	result, err := client.PlaceDetails(context.Background(), r)
	if err != nil {
		return GetSearch{}, err
	}
	return GetSearch{
		Name:           result.Name,
		PlaceID:        result.PlaceID,
		Location:       result.Geometry.Location,
		PhotoReference: result.Photos[0].PhotoReference,
		PriceLevel:     result.PriceLevel,
		Rating:         result.Rating,
		Address:        result.Vicinity,
	}, nil
}

// Room作成リクエスト
func (h *Handler) CreateRoom(c echo.Context) error {
	var req CreateRoomRequest
	if err := c.Bind(&req); err != nil { //リクエストをバインド
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	roomID, err := h.repo.CreateRoom(c.Request().Context(), repository.CreateRoomParams{ // ルームを作成
		Name:        req.Name,
		PlaceMax:    req.PlaceMax,
		CenterPoint: req.CenterPlaceID,
		Radius:      req.Radius,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create room"})
	}
	for _, timeOption := range req.TimeOptions { // 時間候補の検証
		if timeOption.StartTime.IsZero() || timeOption.EndTime.IsZero() {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid time option"})
		}
		if timeOption.StartTime.After(timeOption.EndTime) {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Start time must be before end time"})
		}
	}
	timeOptionsWithID := make([]repository.CreateTimeOptionParams, len(req.TimeOptions)) // TimeOptionの作成用スライス
	placeOptionsWithID := make([]repository.CreatePlaceParams, len(req.PlaceOptions))    // PlaceOptionの作成用スライス
	for i, timeOption := range req.TimeOptions {
		timeOptionsWithID[i] = repository.CreateTimeOptionParams{
			RoomID:    roomID,
			StartTime: timeOption.StartTime,
			EndTime:   timeOption.EndTime,
		}
	}
	for i, placeOption := range req.PlaceOptions {
		placeOptionsWithID[i] = repository.CreatePlaceParams{
			RoomID:        roomID,
			GooglePlaceID: placeOption,
		}
	}

	for _, timeOptionWithID := range timeOptionsWithID {
		if _, err := h.repo.CreateTimeOption(c.Request().Context(), timeOptionWithID); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create time option"})
		}
	}
	for _, placeOptionWithID := range placeOptionsWithID {
		if _, err := h.repo.CreatePlace(c.Request().Context(), placeOptionWithID); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create place option"})
		}
	}

	resTimeOptions := make([]ReturnTimeOption, len(timeOptionsWithID))    // レスポンス用のTimeOptionスライス
	resPlaceOptions := make([]ReturnPlaceOption, len(placeOptionsWithID)) // レスポンス用のPlaceOptionスライス

	roomTimeOptions, err := h.repo.GetTimeOptionsByRoomID(c.Request().Context(), roomID) // ルームIDに紐づくtimeOptionsを取得
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve time options"})
	}
	roomPlaces, err := h.repo.GetPlacesByRoomID(c.Request().Context(), roomID) // ルームIDに紐づくplacesを取得
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve place options"})
	}

	for i, timeOption := range roomTimeOptions { // ルームIDに紐づくtimeOptionsをレスポンス用に変換
		resTimeOptions[i] = ReturnTimeOption{
			ID:        timeOption.ID,
			StartTime: timeOption.StartTime,
			EndTime:   timeOption.EndTime,
		}
	}

	// placesの詳細情報をGoogle Maps APIから取得する
	googleClient, err := h.GetClient(c) // Google Mapsクライアントの取得
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create Google Maps client"})
	}

	for i, placeOption := range roomPlaces {
		detail, err := h.GetPlaceDetail(c, *googleClient, placeOption.GooglePlaceID) // Google Placeの詳細情報を取得
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve place details"})
		}
		resPlaceOptions[i] = ReturnPlaceOption{ // レスポンス用に変換
			ID:             placeOption.ID,
			GooglePlaceID:  placeOption.GooglePlaceID,
			Name:           detail.Name,
			Location:       detail.Location,
			PhotoReference: detail.PhotoReference,
			PriceLevel:     detail.PriceLevel,
			Rating:         detail.Rating,
			Address:        detail.Address,
		}
	}

	res := CreateRoomResponse{ // レスポンスの構築
		RoomID:       roomID,
		TimeOptions:  resTimeOptions,
		PlaceOptions: resPlaceOptions,
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *Handler) GetRoom(c echo.Context) error {
	roomIDStr := c.Param("room_id") // パラメータからroom_idを取得
	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid room_id")
	}
	room, err := h.repo.GetRoom(c.Request().Context(), roomID) // ルーム情報を取得
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "room not found")
	}
	roomTimeOptions, err := h.repo.GetTimeOptionsByRoomID(c.Request().Context(), roomID) // ルームIDに紐づくtimeOptionsを取得
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to retrieve time options")
	}
	roomPlaces, err := h.repo.GetPlacesByRoomID(c.Request().Context(), roomID) // ルームIDに紐づくplacesを取得
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to retrieve place options")
	}

	resTimeOptions := make([]ReturnTimeOption, len(roomTimeOptions)) // レスポンス用のTimeOptionスライス
	resPlaceOptions := make([]ReturnPlaceOption, len(roomPlaces))    // レスポンス用のPlaceOptionスライス
	for i, timeOption := range roomTimeOptions {                     // ルームIDに紐づくtimeOptionsをレスポンス用に変換
		resTimeOptions[i] = ReturnTimeOption{
			ID:        timeOption.ID,
			StartTime: timeOption.StartTime,
			EndTime:   timeOption.EndTime,
		}
	}
	// placesの詳細情報をGoogle Maps APIから取得する
	googleClient, err := h.GetClient(c) // Google Mapsクライアントの取得
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create Google Maps client")
	}

	for i, placeOption := range roomPlaces {
		detail, err := h.GetPlaceDetail(c, *googleClient, placeOption.GooglePlaceID) // Google Placeの詳細情報を取得
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to retrieve place details")
		}
		resPlaceOptions[i] = ReturnPlaceOption{ // レスポンス用に変換
			ID:             placeOption.ID,
			GooglePlaceID:  placeOption.GooglePlaceID,
			Name:           detail.Name,
			Location:       detail.Location,
			PhotoReference: detail.PhotoReference,
			PriceLevel:     detail.PriceLevel,
			Rating:         detail.Rating,
			Address:        detail.Address,
		}
	}
	res := GetRoomResponse{ // レスポンスの構築
		RoomID:       room.ID,
		Name:         room.Name,
		CenterPoint:  room.CenterPoint,
		Radius:       room.Radius,
		PlaceMax:     room.PlaceMax,
		TimeOptions:  resTimeOptions,
		PlaceOptions: resPlaceOptions,
	}
	return c.JSON(http.StatusOK, res)
}
