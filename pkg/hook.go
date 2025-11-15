package pkg

import (
	"fmt"

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
		setStatusColor(data.Dai.Usd), data.Dai.Usd,
		setStatusColor(data.PaypalUsd.Usd), data.PaypalUsd.Usd,
		setStatusColor(data.Tether.Usd), data.Tether.Usd,
		setStatusColor(data.UsdCoin.Usd), data.UsdCoin.Usd,
		setStatusColor(data.Usds.Usd), data.Usds.Usd,
	)

	return html
}
