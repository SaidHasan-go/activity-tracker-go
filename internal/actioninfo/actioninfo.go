package actioninfo

import (
	"fmt"
	"log"
)

// DataParser defines an interface for parsing input records
// and producing a formatted activity summary.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info processes a dataset using the provided DataParser.
// It parses each record, prints formatted information to stdout,
// and logs errors without stopping the processing.
func Info(dataset []string, dp DataParser) {
	for _, record := range dataset {
		err := dp.Parse(record)
		if err != nil {
			log.Printf("error parsing record \"%s\": %v", record, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("error getting activity information: %v", err)
			continue
		}

		fmt.Println(info)
	}
}
