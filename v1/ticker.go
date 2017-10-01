package bitfinex

import (
	"strings"
)

type TickerService struct {
	client *Client
}

type Tick struct {
	Mid       string
	Bid       string
	Ask       string
	LastPrice string `json:"last_price"`
	Low       string
	High      string
	Volume    string
}

// Get(pair) - return last Tick for specified pair
func (s *TickerService) Get(pair string) (Tick, error) {
	pair = strings.ToUpper(pair)
	req, err := s.client.newRequest("GET", "pubticker/"+pair, nil)

	if err != nil {
		return Tick{}, err
	}

	var v = &Tick{}
	_, err = s.client.do(req, v)

	if err != nil {
		return Tick{}, err
	}

	return *v, nil
}
