package service

import (
	"errors"
	"github.com/jesper-nord/pcc-planner/tibber"
	"log"
	"sort"
	"time"
)

func FetchTomorrowPrices(token string) ([]tibber.Price, error) {
	response, err := tibber.GetPrices(token)
	if err != nil {
		return nil, err
	}

	for _, home := range response.Data.Viewer.Homes {
		if home.Subscription != nil {
			return home.Subscription.PriceInfo.Tomorrow, nil
		}
	}
	return nil, errors.New("unable to find home")
}

func CalculateCheapestPrices(targetHours int, prices []tibber.Price) map[int]bool {
	sort.Slice(prices, func(i, j int) bool {
		startsAt1, _ := time.Parse(time.RFC3339, prices[i].StartsAt)
		startsAt2, _ := time.Parse(time.RFC3339, prices[j].StartsAt)
		return startsAt1.Before(startsAt2)
	})

	result := make(map[int]bool)

	var totalDayPrice float64
	for _, price := range prices {
		startsAt, _ := time.Parse(time.RFC3339, price.StartsAt)
		result[startsAt.Hour()] = false
		totalDayPrice += price.Total
	}

	sort.Slice(prices, func(i, j int) bool {
		return prices[i].Total < prices[j].Total
	})

	var total float64
	for i := 0; i < targetHours; i++ {
		price := prices[i]
		startsAt, _ := time.Parse(time.RFC3339, price.StartsAt)
		result[startsAt.Hour()] = true
		total += price.Total
	}

	log.Printf("result: %v", result)
	return result
}
