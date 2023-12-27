package service

import (
	"errors"
	"github.com/jesper-nord/pcc-planner/tibber"
	"log"
	"slices"
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

func CalculateCheapestPrices(activeHours int, forcedHours []int, prices []tibber.Price) map[int]bool {
	result := make(map[int]bool)

	for _, price := range prices {
		startsAt, _ := time.Parse(time.RFC3339, price.StartsAt)
		hour := startsAt.Hour()
		result[hour] = slices.Contains(forcedHours, hour)
	}

	sort.Slice(prices, func(i, j int) bool {
		return prices[i].Total < prices[j].Total
	})

	for i := 0; i < activeHours-len(forcedHours); i++ {
		price := prices[i]
		startsAt, _ := time.Parse(time.RFC3339, price.StartsAt)
		hour := startsAt.Hour()
		if slices.Contains(forcedHours, hour) {
			// hour already included
			continue
		}
		result[hour] = true
	}

	log.Printf("result: %v", result)
	return result
}
