package file

import (
	"fmt"
	"github.com/jesper-nord/pcc-planner/service"
	"log"
	"os"
	"path"
	"time"
)

func WriteToOutput(hours []service.HourResult, outputDir string) error {
	fileName := fmt.Sprintf("pccplan.%s.tsv", time.Now().Add(time.Hour*24).Format("2006-01-02"))

	var output string
	for _, hourResult := range hours {
		output += fmt.Sprintf("%d\t%t\n", hourResult.Hour, hourResult.Enabled)
	}

	file := path.Join(outputDir, fileName)
	err := os.WriteFile(file, []byte(output), 0666)
	if err != nil {
		return fmt.Errorf("unable to write to output file: %w", err)
	}
	log.Printf("wrote output to %s", file)
	return nil
}
