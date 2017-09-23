package lib

import "time"

type APIStatsPricesItem struct {
	ItemID           string    `json:"item_id"`
	City             string    `json:"city"`
	SellPriceMin     int       `json:"sell_price_min"`
	SellPriceMinDate time.Time `json:"sell_price_min_date"`
	SellAmountMin    int       `json:"sell_amount_min"`
	SellPriceMax     int       `json:"sell_price_max"`
	SellPriceMaxDate time.Time `json:"sell_price_max_date"`
	SellAmountMax    int       `json:"sell_amount_max"`
	BuyPriceMin      int       `json:"buy_price_min"`
	BuyPriceMinDate  time.Time `json:"buy_price_min_date"`
	BuyAmountMin     int       `json:"buy_amount_min"`
	BuyPriceMax      int       `json:"buy_price_max"`
	BuyPriceMaxDate  time.Time `json:"buy_price_max_date"`
	BuyAmountMax     int       `json:"buy_amount_max"`
}

type APIStatsChartsResponse struct {
	Location string                         `json:"location"`
	Data     APIStatsChartsLocationResponse `json:"data"`
}

type APIStatsChartsLocationResponse struct {
	Timestamps []int64   `json:"timestamps"`
	PricesMin  []int     `json:"prices_min"`
	PricesMax  []int     `json:"prices_max"`
	PricesAvg  []float64 `json:"prices_avg"`
}
