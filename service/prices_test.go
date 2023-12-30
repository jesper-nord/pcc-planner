package service

import (
	"github.com/jesper-nord/pcc-planner/tibber"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCalculateCheapestPrices(t *testing.T) {
	prices := []tibber.Price{
		{
			Total:    0.6054,
			StartsAt: time.Date(2023, 12, 27, 0, 0, 0, 0, time.Local),
		},
		{
			Total:    0.5502,
			StartsAt: time.Date(2023, 12, 27, 1, 0, 0, 0, time.Local),
		},
		{
			Total:    0.5444,
			StartsAt: time.Date(2023, 12, 27, 2, 0, 0, 0, time.Local),
		},
		{
			Total:    0.5486,
			StartsAt: time.Date(2023, 12, 27, 3, 0, 0, 0, time.Local),
		},
		{
			Total:    0.5702,
			StartsAt: time.Date(2023, 12, 27, 4, 0, 0, 0, time.Local),
		},
		{
			Total:    0.5878,
			StartsAt: time.Date(2023, 12, 27, 5, 0, 0, 0, time.Local),
		},
		{
			Total:    0.6248,
			StartsAt: time.Date(2023, 12, 27, 6, 0, 0, 0, time.Local),
		},
		{
			Total:    0.878,
			StartsAt: time.Date(2023, 12, 27, 7, 0, 0, 0, time.Local),
		},
		{
			Total:    1.1921,
			StartsAt: time.Date(2023, 12, 27, 8, 0, 0, 0, time.Local),
		},
		{
			Total:    1.1939,
			StartsAt: time.Date(2023, 12, 27, 9, 0, 0, 0, time.Local),
		},
		{
			Total:    1.1325,
			StartsAt: time.Date(2023, 12, 27, 10, 0, 0, 0, time.Local),
		},
		{
			Total:    1.1208,
			StartsAt: time.Date(2023, 12, 27, 11, 0, 0, 0, time.Local),
		},
		{
			Total:    1.1136,
			StartsAt: time.Date(2023, 12, 27, 12, 0, 0, 0, time.Local),
		},
		{
			Total:    1.0961,
			StartsAt: time.Date(2023, 12, 27, 13, 0, 0, 0, time.Local),
		},
		{
			Total:    1.1308,
			StartsAt: time.Date(2023, 12, 27, 14, 0, 0, 0, time.Local),
		},
		{
			Total:    1.1909,
			StartsAt: time.Date(2023, 12, 27, 15, 0, 0, 0, time.Local),
		},
		{
			Total:    1.2203,
			StartsAt: time.Date(2023, 12, 27, 16, 0, 0, 0, time.Local),
		},
		{
			Total:    1.2279,
			StartsAt: time.Date(2023, 12, 27, 17, 0, 0, 0, time.Local),
		},
		{
			Total:    1.1457,
			StartsAt: time.Date(2023, 12, 27, 18, 0, 0, 0, time.Local),
		},
		{
			Total:    1.0149,
			StartsAt: time.Date(2023, 12, 27, 19, 0, 0, 0, time.Local),
		},
		{
			Total:    0.8756,
			StartsAt: time.Date(2023, 12, 27, 20, 0, 0, 0, time.Local),
		},
		{
			Total:    0.6302,
			StartsAt: time.Date(2023, 12, 27, 21, 0, 0, 0, time.Local),
		},
		{
			Total:    0.5808,
			StartsAt: time.Date(2023, 12, 27, 22, 0, 0, 0, time.Local),
		},
		{
			Total:    0.555,
			StartsAt: time.Date(2023, 12, 27, 23, 0, 0, 0, time.Local),
		},
	}

	activeHours := 16
	forcedHours := []int{0, 1, 2, 3, 4, 5, 6, 7, 22, 23}

	expected := []HourResult{
		{Hour: 0, Enabled: true},
		{Hour: 1, Enabled: true},
		{Hour: 2, Enabled: true},
		{Hour: 3, Enabled: true},
		{Hour: 4, Enabled: true},
		{Hour: 5, Enabled: true},
		{Hour: 6, Enabled: true},
		{Hour: 7, Enabled: true},
		{Hour: 8, Enabled: false},
		{Hour: 9, Enabled: false},
		{Hour: 10, Enabled: false},
		{Hour: 11, Enabled: true},
		{Hour: 12, Enabled: true},
		{Hour: 13, Enabled: true},
		{Hour: 14, Enabled: false},
		{Hour: 15, Enabled: false},
		{Hour: 16, Enabled: false},
		{Hour: 17, Enabled: false},
		{Hour: 18, Enabled: false},
		{Hour: 19, Enabled: true},
		{Hour: 20, Enabled: true},
		{Hour: 21, Enabled: true},
		{Hour: 22, Enabled: true},
		{Hour: 23, Enabled: true},
	}

	result := CalculateCheapestPrices(activeHours, forcedHours, prices)
	assert.Equal(t, expected, result.HourResult)
}
