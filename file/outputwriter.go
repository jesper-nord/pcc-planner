package file

import (
	"fmt"
	"os"
	"path"
	"time"
)

func WriteToOutput(result map[int]bool, outputDir string) error {
	fileName := fmt.Sprintf("pcc-res-%s", time.Now().Add(time.Hour*24).Format("2006-01-02"))

	var output string
	for hour, active := range result {
		output += fmt.Sprintf("%d\t%t\n", hour, active)
	}

	file := path.Join(outputDir, fileName)
	err := os.WriteFile(file, []byte(output), 0666)
	if err != nil {
		return fmt.Errorf("unable to write to output file: %w", err)
	}
	return nil
}
