package main

import (
	"flag"
	"github.com/jesper-nord/pcc-planner/file"
	"github.com/jesper-nord/pcc-planner/service"
	"os"
	"strconv"
	"strings"
)

var tokenFlag = flag.String("token", "", "Tibber API token")
var activeHoursFlag = flag.Int("hours", 0, "Number of active hours per period (value between 0-24)")
var forcedHoursFlag = flag.String("forcedhours", "", "Comma separated list of forced hours (e.g. 5,6,7)")
var outputDirFlag = flag.String("outputdir", "/tmp/pccplanner/out", "Output file directory")
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

	err = file.WriteToOutput(result, *outputDirFlag)
	if err != nil {
		panic(err)
	}
}
