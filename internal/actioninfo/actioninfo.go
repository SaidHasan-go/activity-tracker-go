package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

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
