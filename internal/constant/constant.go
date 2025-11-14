package constant

const (
	ROUTE_ROOT                             = "/"
	ROUTE_API                              = "/api"
	FilePermUserReadWrite                  = 0600
	FilePermUserReadWriteGroupRead         = 0644
	FilePermExecutable                     = 0755
	ETH_DECIMAL                    float64 = 1e18
	USDT_DECIMAL                   float64 = 1e6
	ETH_USDT_ADDRESS                       = "0xdAC17F958D2ee523a2206206994597C13D831ec7" // eth mainnet

	// sse event name list
	SSE_STATS = "sse_stats"

	// coins endpoint
	ENDPOINT_USDT  = ""
	ENDPOINT_USDC  = ""
	ENDPOINT_USDS  = "https://api.coingecko.com/api/v3/simple/price?ids=usds&vs_currencies=usd"
	ENDPOINT_DAI   = "https://api.coingecko.com/api/v3/simple/price?ids=dai&vs_currencies=usd"
	ENDPOINT_PYUSD = ""
)
