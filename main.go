package main

import (
	"flag"
	"github.com/jesper-nord/pcc-planner/file"
	"github.com/jesper-nord/pcc-planner/service"
	"os"
)

var tokenFlag = flag.String("token", "", "Tibber API token")
var hoursFlag = flag.Int("hours", 0, "Number of active hours per period (value between 0-24)")
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

	result := service.CalculateCheapestPrices(*hoursFlag, prices)

	err = file.WriteToOutput(result, *outputDirFlag)
	if err != nil {
		panic(err)
	}
}
