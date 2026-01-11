package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/SaidHasan-go/activity-tracker-go/internal/personaldata"
	"github.com/SaidHasan-go/activity-tracker-go/internal/spentenergy"
)

var (
	ErrinvalidFormat = errors.New("Invalid Format")
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 2 {
		return fmt.Errorf("Incorrect amount of data %v", ErrinvalidFormat)
	}
	ds.Steps, err = strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("Incorrect number of steps %v", err)
	}
	if ds.Steps <= 0 {
		return fmt.Errorf("Negative number of steps")
	}
	ds.Duration, err = time.ParseDuration(slice[1])
	if err != nil {
		return fmt.Errorf("Invalid walk duration format %v", err)
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("negative duration of the walk %v", ErrinvalidFormat)
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distanc := spentenergy.Distance(ds.Steps, ds.Height)
	calorie, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Height, ds.Weight, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("Calorie calculation error %v", err)
	}

	str := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distanc, calorie)
	return str, nil
}
