package rest

type OrderOptions struct {
	ClientOrderID         string              `json:"client_order_id"`
	ProductID             string              `json:"product_id"`
	OrderConfiguration    *OrderConfiguration `json:"order_configuration,omitempty"`
	SelfTradePreventionID string              `json:"self_trade_prevention_id,omitempty"`
	Leverage              string              `json:"leverage,omitempty"`
	MarginType            string              `json:"margin_type,omitempty"`
	RetailPortfolioID     string              `json:"retail_portfolio_id,omitempty"`
}

type OrderConfiguration struct {
	MarketMarketIOC       MarketMarketIOC       `json:"market_market_ioc,omitempty"`
	SORLimitIOC           SORLimitIOC           `json:"sor_limit_ioc,omitempty"`
	LimitLimitGTCT        LimitLimitGTC         `json:"limit_limit_gtc,omitempty"`
	StopLimitStopLimitGTC StopLimitStopLimitGTC `json:"stop_limit_stop_limit_gtc,omitempty"`
	StopLimitStopLimitGTD StopLimitStopLimitGTD `json:"stop_limit_stop_limit_gtd,omitempty"`
	TriggerBracketGTC     TriggerBracketGTC     `json:"trigger_bracket_gtc,omitempty"`
	TriggerBracketGTD     TriggerBracketGTD     `json:"trigger_bracket_gtd,omitempty"`
}
type MarketMarketIOC struct {
	QuoteSize string `json:"quote_size,omitempty"`
	BaseSize  string `json:"base_size,omitempty"`
}
type SORLimitIOC struct {
	BaseSize   string `json:"base_size,omitempty"`
	LimitPrice string `json:"limit_price,omitempty"`
}
type LimitLimitGTC struct {
	BaseSize   string `json:"base_size,omitempty"`
	LimitPrice string `json:"limit_price,omitempty"`
	PostOnly   bool   `json:"post_only,omitempty"`
}
type LimitLimitGTD struct {
	BaseSize   string `json:"base_size,omitempty"`
	LimitPrice string `json:"limit_price,omitempty"`
	PostOnly   bool   `json:"post_only,omitempty"`
	EndTime    string `json:"end_time,omitempty"`
}
type StopLimitStopLimitGTC struct {
	BaseSize      string `json:"base_size,omitempty"`
	LimitPrice    string `json:"limit_price,omitempty"`
	StopPrice     string `json:"stop_price,omitempty"`
	StopDirection string `json:"stop_direction,omitempty"`
}
type StopLimitStopLimitGTD struct {
	BaseSize      string `json:"base_size,omitempty"`
	LimitPrice    string `json:"limit_price,omitempty"`
	StopPrice     string `json:"stop_price,omitempty"`
	EndTime       string `json:"end_time,omitempty"`
	StopDirection string `json:"stop_direction,omitempty"`
}
type TriggerBracketGTC struct {
	BaseSize         string `json:"base_size,omitempty"`
	LimitPrice       string `json:"limit_price,omitempty"`
	StopTriggerPrice string `json:"stop_trigger_price,omitempty"`
}
type TriggerBracketGTD struct {
	BaseSize         string `json:"base_size,omitempty"`
	LimitPrice       string `json:"limit_price,omitempty"`
	StopTriggerPrice string `json:"stop_trigger_price,omitempty"`
	EndTime          string `json:"end_time,omitempty"`
}

type OrderResponse struct {
	Success bool `json:"success"`
}
