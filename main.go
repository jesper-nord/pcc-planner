package main

import (
	"flag"
	"fmt"
	"github.com/jesper-nord/pcc-planner/file"
	"github.com/jesper-nord/pcc-planner/service"
	"github.com/jesper-nord/pcc-planner/tibber"
	"os"
	"strconv"
	"strings"
)

var tokenFlag = flag.String("token", "", "Tibber API token")
var activeHoursFlag = flag.Int("hours", 0, "Number of active hours per period (value between 0-24)")
var forcedHoursFlag = flag.String("forced-hours", "", "Comma separated list of forced hours (e.g. 5,6,7)")
var outputDirFlag = flag.String("output-dir", "/tmp/pccplanner/out", "Output file directory")
var notifyFlag = flag.Bool("notify", false, "Send Tibber notification with plan")
var helpFlag = flag.Bool("help", false, "show available commands")

func main() {
	if len(os.Args) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()
	if *helpFlag {
		flag.PrintDefaults()
		os.Exit(1)
	}

	prices, err := service.FetchTomorrowPrices(*tokenFlag)
	if err != nil {
		panic(err)
	}

	var forcedHours []int
	if *forcedHoursFlag != "" {
		split := strings.Split(*forcedHoursFlag, ",")
		forcedHours = make([]int, len(split))
		for i, s := range split {
			forcedHours[i], _ = strconv.Atoi(s)
		}
	}
	result := service.CalculateCheapestPrices(*activeHoursFlag, forcedHours, prices)

	err = file.WriteToOutput(result.HourResult, *outputDirFlag)
	if err != nil {
		panic(err)
	}

	if *notifyFlag {
		var msg []string
		for _, hourResult := range result.HourResult {
			msg = append(msg, fmt.Sprintf("%d: %t", hourResult.Hour, hourResult.Enabled))
		}
		_, err = tibber.SendNotification(*tokenFlag, fmt.Sprintf("PCC plan %s", result.Date.Format("2006-01-02")), strings.Join(msg, ", "))
		if err != nil {
			panic(err)
		}
	}
}
