package handler

import (
	"TopUpWEb/sdk"
	"TopUpWEb/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GamesHandler struct {
	gameService service.GameService
}

func NewGamesHandler(gamesService service.GameService) *GamesHandler {
	return &GamesHandler{gamesService}
}

type gameResult struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama"`
	Link string `json:"link"`
}

func (h *GamesHandler) GetAllGames(c *gin.Context) {
	games, err := h.gameService.FindAll()

	if err != nil {
		sdk.Fail(c, 500, "Failed to load games")
		return
	}
	var results []gameResult
	for _, game := range games {
		result := gameResult{
			ID:   game.ID,
			Nama: game.Nama,
			Link: game.Gambar,
		}
		results = append(results, result)
	}

	sdk.Success(c, http.StatusOK, "Games loaded	", results)
}

type gameByIdResult struct {
	ID        uint           `json:"id"`
	Nama      string         `json:"nama"`
	Developer string         `json:"developer"`
	Link      string         `json:"link"`
	TopUps    []listTopUpRes `json:"topups"`
}

type listTopUpRes struct {
	ID    uint   `json:"id"`
	Price string `json:"price"`
	Type  string `json:"type"`
}

func (h *GamesHandler) GetGamebyID(c *gin.Context) {
	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)

	game, err := h.gameService.FindbyId(uint(id))

	if err != nil {
		sdk.FailOrError(c, http.StatusNotFound, "Data not found", err)
		return
	}

	listTopUps, err := h.gameService.GetListTopUpsByGameID(uint(id))
	if err != nil {
		sdk.FailOrError(c, http.StatusInternalServerError, "Failed to retrieve ListTopUps", err)
		return
	}
	results := gameByIdResult{
		ID:        game.ID,
		Nama:      game.Nama,
		Developer: game.Developer,
		Link:      game.Gambar,
	}
	var topupResults []listTopUpRes
	for _, topup := range listTopUps {
		topupResults = append(topupResults, listTopUpRes{
			ID:    topup.ID,
			Price: topup.Harga,
			Type:  topup.JenisPaket,
		})
	}
	results.TopUps = topupResults
	sdk.Success(c, 200, "Data Found", results)
}
