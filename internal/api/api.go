package api

import (
	"log"
	"os"
	"strconv"

	"net/http"
	"time"

	"github.com/developerasun/minlink/internal/constant"
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

// CrawlDollarIndex godoc
// @Summary visit tradingview and extract daily dollar index
// @Description Get server health status
// @Tags api
// @Produce json
// @Success 200 {object} CrawlResponse
// @Router /api/dxy [get]
func CrawlDollarIndex(ctx *gin.Context) {
	dxy, rcErr := pkg.CrawlDollarIndex()

	if rcErr != nil {
		log.Println(`CrawlDollarIndex: failed to run a crawler`)
		ctx.Error(rcErr)
	}

	ctx.JSON(http.StatusOK, CrawlResponse{
		Data: dxy,
	})
}

// CrawlDaiPrice godoc
// @Summary visit metamask and extract dai token price at the moment
// @Description Get dai coin price
// @Tags api
// @Produce json
// @Success 200 {object} CrawlResponse
// @Router /api/dai [get]
func CrawlDaiPrice(ctx *gin.Context) {
	dai, rcErr := pkg.CrawlDaiPrice()

	if rcErr != nil {
		log.Println(`CrawlDaiPrice: failed to run a crawler`)
		ctx.Error(rcErr)
	}

	ctx.JSON(http.StatusOK, CrawlResponse{
		Data: dai,
	})
}

// Health godoc
// @Summary Send target index stats as streams
// @Description use server side event to dynamically render target data
// @Tags api
// @Produce json
// @Success 200 {object} SseStatsResponse
// @Router /api/sse_stats [get]
func RenderStats(ctx *gin.Context) {
	isProduction := os.Getenv("PRODUCTION") == "true"
	_interval := os.Getenv("SSE_INTERVAL")

	interval, aErr := strconv.Atoi(_interval)
	if aErr != nil {
		log.Println("RenderStats: " + aErr.Error())
	}
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	defer ticker.Stop()

	// handle first event right away
	data := pkg.GetDataSources()
	html := pkg.FinalizePeggingStats(data)
	log.Println("handle first event: ", data)
	ctx.SSEvent(constant.SSE_STATS, html)
	ctx.Writer.Flush()

	// stream next events when client stays on page
	for {
		select {
		case <-ticker.C:

			var data pkg.CoinGeckoApiType
			if isProduction {
				log.Println("requesting real coingekco api in production")

				data := pkg.GetDataSources()
				html := pkg.FinalizePeggingStats(data)

				// @dev queue stream data to response buffer
				ctx.SSEvent(constant.SSE_STATS, html)
			} else {
				log.Println("immitating dummy coingekco api in development")

				data = pkg.GetDummyDataSources()
				html := pkg.FinalizePeggingStats(data)
				ctx.SSEvent(constant.SSE_STATS, html)
			}

			// @dev deliver the buffer to client
			ctx.Writer.Flush()

		case <-time.After(time.Minute * 1):
			log.Println("timed out")
			return

		case <-ctx.Request.Context().Done():
			log.Println("client disconnected")
			return
		}
	}
}

// RenderMainPage godoc
// @Summary show main page, returning html
// @Description show main page, returning html
// @Tags view
// @Router / [get]
func RenderMainPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
