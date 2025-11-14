package constant

const (
	ROUTE_ROOT = "/"
	ROUTE_API  = "/api"

	// depegging alert
	ALERT_DEPEGGING_MIN float32 = 0.990
	ALERT_DEPEGGING_MAX float32 = 1.010

	// eth mainnet
	ETH_DECIMAL   float64 = 1e18
	USDT_DECIMAL  float64 = 1e6
	USDC_DECIMAL  float64 = 1e6
	USDS_DECIMAL  float64 = 1e18
	DAI_DECIMAL   float64 = 1e18
	PYUSD_DECIMAL float64 = 1e6

	USDT_CONTRACT  = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	USDC_CONTRACT  = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	USDS_CONTRACT  = "0xDC035D45D973e3EC169D2276dDab16f1E407384F"
	DAI_CONTRACT   = "0x6B175474E89094C44Da98b954EedeAC495271d0F"
	PYUSD_CONTRACT = "0x6c3ea9036406852006290770bedfcaba0e23a0e8"

	// sse event name list
	SSE_STATS = "sse_stats"

	// coins endpoint
	ENDPOINT_COINS = "https://api.coingecko.com/api/v3/simple/price?ids=tether,usd-coin,usds,dai,paypal-usd&vs_currencies=usd"
)
