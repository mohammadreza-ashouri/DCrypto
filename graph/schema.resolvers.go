package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/mohammadreza-ashouri/oracle-pricefeed-lab/graph/generated"
	"github.com/mohammadreza-ashouri/oracle-pricefeed-lab/graph/model"
)

type Response struct {
	timestamp             string `json:"timestamp"`
	low                   string `json:"low"`
	high                  string `json:"high"`
	last                  string `json:"last"`
	volume                string `json:"volume"`
	volume30d             string `json:"volume30d"`
	priceChange           string `json:"priceChange"`
	priceChangePercentage string `json:"priceChangePercentage"`
	pair                  string `json:"pair"`
}

func (r *mutationResolver) CreateToken(ctx context.Context, input model.NewToken) (*model.Token, error) {
	var token model.Token
	var localtimestamp string
	localtimestamp = strconv.FormatInt(time.Now().UnixNano(), 10)

	token.Symbol = input.Symbol
	token.Usdprice = input.Usdprice
	token.Europrice = input.Usdprice
	token.Nokprice = input.Nokprice
	token.Timestamp = localtimestamp

	return &token, nil
}

func (r *mutationResolver) CreatePrice(ctx context.Context, input model.NewPrice) (*model.Price, error) {
	var price model.Price
	price.Timestamp = input.Timestamp
	price.Low = input.Low
	price.High = input.High
	price.Last = input.Last
	price.Volume = input.Volume
	price.Volume30d = input.Volume30d
	price.PriceChange = input.PriceChange
	price.PriceChangePercentage = input.PriceChangePercentage
	price.Pair = input.Pair

	return &price, nil
}

func (r *queryResolver) Tokens(ctx context.Context) ([]*model.Token, error) {
	var tokens []*model.Token
	var localtimestamp string
	localtimestamp = strconv.FormatInt(time.Now().UnixNano(), 10)

	dummyToken := model.Token{

		Symbol:    "ETH",
		Usdprice:  "4000",
		Europrice: "3344",
		Nokprice:  "33540",
		Timestamp: localtimestamp,
	}
	tokens = append(tokens, &dummyToken)
	return tokens, nil
}

func (r *queryResolver) Prices(ctx context.Context) ([]*model.Price, error) {

	var prices []*model.Price

	//---------get prices from a REST API

	response, err := http.Get("https://cex.io/api/ticker/ETH/USD")

	responseData, err := ioutil.ReadAll(response.Body)
	fmt.Println(responseData)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.last)
	fmt.Println(responseData)
	//--------------------

	dummyPrice := model.Price{

		Timestamp:             responseObject.timestamp,
		Low:                   responseObject.low,
		High:                  responseObject.high,
		Last:                  responseObject.last,
		Volume:                responseObject.volume,
		Volume30d:             responseObject.volume30d,
		PriceChange:           responseObject.priceChange,
		PriceChangePercentage: responseObject.priceChangePercentage,
		Pair:                  responseObject.pair,
	}
	prices = append(prices, &dummyPrice)
	return prices, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
