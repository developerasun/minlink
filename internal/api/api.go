package api

import (
	"log"
	"net/http"

	"github.com/developerasun/minlink/pkg"
	"github.com/gin-gonic/gin"
)

// Health godoc
// @Summary Show the health status
// @Description Get server health status
// @Tags api
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /api/health [get]
func Health(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, HealthResponse{
		Message: "ok",
	})
}

// Crawl godoc
// @Summary visit tradingview and extract daily dollar index
// @Description Get server health status
// @Tags api
// @Produce json
// @Success 200 {object} CrawlResponse
// @Router /api/crawl [get]
func Crawl(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	dxy, rcErr := pkg.RunCrawler()
	if rcErr != nil {
		log.Println(`Crawl: failed to run a crawler`)
		ctx.Error(rcErr)
	}

	ctx.JSON(http.StatusOK, CrawlResponse{
		Dxy: dxy,
	})
}

// RenderMainPage godoc
// @Summary show main page, returning html
// @Description show main page, returning html
// @Tags view
// @Router / [get]
func RenderMainPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
