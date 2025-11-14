package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"

	// "math/rand"
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

	for {
		select {
		case <-ticker.C:

			if isProduction {
				log.Println("requesting real coingekco api in production")

				request, _ := http.NewRequest(http.MethodGet, constant.ENDPOINT_COINS, nil)
				client := http.Client{}
				response, _ := client.Do(request)

				raw, _ := io.ReadAll(response.Body)
				response.Body.Close() // @dev prevent FD leak

				var data CoinGeckoApiResponse
				if err := json.Unmarshal(raw, &data); err != nil {
					log.Println("RenderStats: " + err.Error())
				}

				// @dev queue stream data to response buffer
				ctx.SSEvent(constant.SSE_STATS, SseStatsResponse{
					Data: data,
				})
			} else {
				log.Println("immitating dummy coingekco api in development")

				data := CoinGeckoApiResponse{
					Dai:       Currency{Usd: rand.Float32()},
					PaypalUsd: Currency{Usd: rand.Float32()},
					Tether:    Currency{Usd: rand.Float32()},
					UsdCoin:   Currency{Usd: rand.Float32()},
					Usds:      Currency{Usd: rand.Float32()},
				}

				alert := func(value float32) string {
					if value <= constant.ALERT_DEPEGGING_MIN || value >= constant.ALERT_DEPEGGING_MAX {
						return `
						  <div class="w-4 flex justify-center">
								<div class="w-4 h-4 rounded-full bg-red-500"></div>
							</div>
						`
					}
					return `
						<div class="w-4 flex justify-center">
							<div class="w-4 h-4 rounded-full bg-green-500"></div>
						</div>
					`
				}

				_html := fmt.Sprintf(`
				<div class="flex justify-center">
					<ul class="flex flex-col p-6 text-3xl gap-2">
						<li class="flex items-center gap-2">
							%s
							<div>DAI: %f</div>
						</li>
						<li class="flex items-center gap-2">
							%s
							<div>PYUSD: %f</div>
						</li>
						<li class="flex items-center gap-2">
							%s
							<div>USDT: %f</div>
						</li>
						<li class="flex items-center gap-2">
							%s
							<div>USDC: %f</div>
						</li>
						<li class="flex items-center gap-2">
							%s
							<div>USDS: %f</div>
						</li>
					</ul>
				</div>
				`,
					alert(data.Dai.Usd), data.Dai.Usd,
					alert(data.PaypalUsd.Usd), data.PaypalUsd.Usd,
					alert(data.Tether.Usd), data.Tether.Usd,
					alert(data.UsdCoin.Usd), data.UsdCoin.Usd,
					alert(data.Usds.Usd), data.Usds.Usd,
				)

				ctx.SSEvent(constant.SSE_STATS, _html)
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
