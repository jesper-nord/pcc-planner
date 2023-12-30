package service

import (
	"errors"
	"github.com/jesper-nord/pcc-planner/tibber"
	"github.com/samber/lo"
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
			prices := home.Subscription.PriceInfo.Tomorrow
			if len(prices) == 0 {
				return nil, errors.New("tomorrow's prices not yet in")
			}
			return prices, nil
		}
	}
	return nil, errors.New("unable to find home")
}

func CalculateCheapestPrices(activeHours int, forcedHours []int, prices []tibber.Price) PriceResult {
	sort.Slice(prices, func(i, j int) bool {
		return prices[i].Total < prices[j].Total
	})

	var result []HourResult
	var date time.Time

	for _, price := range prices {
		date = price.StartsAt
		hour := date.Hour()
		if slices.Contains(forcedHours, hour) {
			result = append(result, HourResult{Hour: hour, Enabled: true})
			continue
		}
		countEnabled := lo.CountBy(result, func(item HourResult) bool {
			return item.Enabled
		})
		result = append(result, HourResult{Hour: hour, Enabled: countEnabled < activeHours})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Hour < result[j].Hour
	})

	log.Printf("result: %v", result)
	return PriceResult{
		HourResult: result,
		Date:       date,
	}
}

type PriceResult struct {
	Date       time.Time
	HourResult []HourResult
}

type HourResult struct {
	Hour    int
	Enabled bool
}
