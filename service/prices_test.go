package service

import (
	"github.com/jesper-nord/pcc-planner/tibber"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateCheapestPrices(t *testing.T) {
	prices := []tibber.Price{
		{
			Total:    0.6054,
			StartsAt: "2023-12-27T00:00:00.000+01:00",
		},
		{
			Total:    0.5502,
			StartsAt: "2023-12-27T01:00:00.000+01:00",
		},
		{
			Total:    0.5444,
			StartsAt: "2023-12-27T02:00:00.000+01:00",
		},
		{
			Total:    0.5486,
			StartsAt: "2023-12-27T03:00:00.000+01:00",
		},
		{
			Total:    0.5702,
			StartsAt: "2023-12-27T04:00:00.000+01:00",
		},
		{
			Total:    0.5878,
			StartsAt: "2023-12-27T05:00:00.000+01:00",
		},
		{
			Total:    0.6248,
			StartsAt: "2023-12-27T06:00:00.000+01:00",
		},
		{
			Total:    0.878,
			StartsAt: "2023-12-27T07:00:00.000+01:00",
		},
		{
			Total:    1.1921,
			StartsAt: "2023-12-27T08:00:00.000+01:00",
		},
		{
			Total:    1.1939,
			StartsAt: "2023-12-27T09:00:00.000+01:00",
		},
		{
			Total:    1.1325,
			StartsAt: "2023-12-27T10:00:00.000+01:00",
		},
		{
			Total:    1.1208,
			StartsAt: "2023-12-27T11:00:00.000+01:00",
		},
		{
			Total:    1.1136,
			StartsAt: "2023-12-27T12:00:00.000+01:00",
		},
		{
			Total:    1.0961,
			StartsAt: "2023-12-27T13:00:00.000+01:00",
		},
		{
			Total:    1.1308,
			StartsAt: "2023-12-27T14:00:00.000+01:00",
		},
		{
			Total:    1.1909,
			StartsAt: "2023-12-27T15:00:00.000+01:00",
		},
		{
			Total:    1.2203,
			StartsAt: "2023-12-27T16:00:00.000+01:00",
		},
		{
			Total:    1.2279,
			StartsAt: "2023-12-27T17:00:00.000+01:00",
		},
		{
			Total:    1.1457,
			StartsAt: "2023-12-27T18:00:00.000+01:00",
		},
		{
			Total:    1.0149,
			StartsAt: "2023-12-27T19:00:00.000+01:00",
		},
		{
			Total:    0.8756,
			StartsAt: "2023-12-27T20:00:00.000+01:00",
		},
		{
			Total:    0.6302,
			StartsAt: "2023-12-27T21:00:00.000+01:00",
		},
		{
			Total:    0.5808,
			StartsAt: "2023-12-27T22:00:00.000+01:00",
		},
		{
			Total:    0.555,
			StartsAt: "2023-12-27T23:00:00.000+01:00",
		},
	}

	activeHours := 12
	forcedHours := []int{19, 20}

	expected := map[int]bool{
		0:  true,
		1:  true,
		2:  true,
		3:  true,
		4:  true,
		5:  true,
		6:  true,
		7:  false,
		8:  false,
		9:  false,
		10: false,
		11: false,
		12: false,
		13: false,
		14: false,
		15: false,
		16: false,
		17: false,
		18: false,
		19: true,
		20: true,
		21: true,
		22: true,
		23: true,
	}

	result := CalculateCheapestPrices(activeHours, forcedHours, prices)
	assert.Equal(t, expected, result)
}
