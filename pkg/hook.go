package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/developerasun/minlink/internal/constant"
)

type Currency struct {
	Usd float32 `json:"usd"`
}

type CoinGeckoApiType struct {
	Dai       Currency `json:"dai"`
	PaypalUsd Currency `json:"paypal-usd"`
	Tether    Currency `json:"tether"`
	UsdCoin   Currency `json:"usd-coin"`
	Usds      Currency `json:"usds"`
}

func GetDummyDataSources() CoinGeckoApiType {
	return CoinGeckoApiType{
		Dai:       Currency{Usd: rand.Float32()},
		PaypalUsd: Currency{Usd: rand.Float32()},
		Tether:    Currency{Usd: rand.Float32()},
		UsdCoin:   Currency{Usd: rand.Float32()},
		Usds:      Currency{Usd: rand.Float32()},
	}
}

func GetDataSources() CoinGeckoApiType {
	var data CoinGeckoApiType
	request, _ := http.NewRequest(http.MethodGet, constant.ENDPOINT_COINS, nil)
	client := http.Client{}
	response, _ := client.Do(request)

	raw, _ := io.ReadAll(response.Body)
	response.Body.Close() // @dev prevent FD leak

	if err := json.Unmarshal(raw, &data); err != nil {
		log.Println("GetDataSources: " + err.Error())
	}

	return data
}

func setStatusColor(value float32) string {
	var color string = "green"

	if value <= constant.ALERT_DEPEGGING_MIN || value >= constant.ALERT_DEPEGGING_MAX {
		color = "red"
	}

	html := fmt.Sprintf(`
		<div class="w-4 flex justify-center">
			<div class="w-4 h-4 rounded-full bg-%s-500"></div>
		</div>
	`, color)

	return html
}

func FinalizePeggingStats(data CoinGeckoApiType) string {
	html := fmt.Sprintf(`
		<div class="flex flex-col justify-center items-center">
			<div class="mx-auto text-center">Time: %s</div>
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
		// @dev in the format of: 2025-11-16 17:43:09.257622246
		(strings.Split(time.Now().Local().String(), "+")[0]),
		setStatusColor(data.Dai.Usd), data.Dai.Usd,
		setStatusColor(data.PaypalUsd.Usd), data.PaypalUsd.Usd,
		setStatusColor(data.Tether.Usd), data.Tether.Usd,
		setStatusColor(data.UsdCoin.Usd), data.UsdCoin.Usd,
		setStatusColor(data.Usds.Usd), data.Usds.Usd,
	)

	return html
}
