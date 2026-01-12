package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/MaximK0valev/activity-tracker-go/internal/personaldata"
	"github.com/MaximK0valev/activity-tracker-go/internal/spentenergy"
)

var (
	// ErrInvalidFormat indicates that the training input has an invalid format.
	ErrinvalidFormat = errors.New("Invalid Format")

	// ErrUnknownType indicates that the training type is not supported.
	ErrUnkownType = errors.New("неизвестный тип тренировки")
)

// Training represents a single training session,
// including its type, duration, steps, and personal parameters.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse parses a training record in the format
// "steps,training_type,duration" and populates the Training fields.
func (t *Training) Parse(datastring string) (err error) {
	slice := strings.Split(datastring, ",")
	if len(slice) != 3 {
		return fmt.Errorf("Incorrect amount of data %v", ErrinvalidFormat)
	}
	t.Steps, err = strconv.Atoi(slice[0])
	if err != nil {
		return fmt.Errorf("Incorrect number of steps %v", err)
	}
	if t.Steps <= 0 {
		return fmt.Errorf("Negative number of steps")
	}
	t.TrainingType = slice[1]

	t.Duration, err = time.ParseDuration(slice[2])
	if err != nil {
		return fmt.Errorf("Invalid walk duration format %v", err)
	}
	if t.Duration <= 0 {
		return fmt.Errorf("negative duration of the walk %v", ErrinvalidFormat)
	}
	return nil
}

// ActionInfo returns a formatted summary of the training session,
// including distance, average speed, and burned calories.
func (t Training) ActionInfo() (string, error) {
	distanc := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	calorie := 0.0
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calorie, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("Calorie calculation error")
		}
	case "Бег":
		calorie, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("Calorie calculation error")
		}
	default:
		return "", ErrUnkownType
	}
	str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distanc, meanSpeed, calorie)
	return str, nil
}
