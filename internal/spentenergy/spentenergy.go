package spentenergy

import (
	"errors"
	"fmt"
	"time"
)

// Base constants used for distance and calorie calculations.
const (
	mInKm                      = 1000 // number of meters in one kilometer
	minInH                     = 60   // number of minutes in one hour
	stepLengthCoefficient      = 0.45 // coefficient for step length based on height
	walkingCaloriesCoefficient = 0.5  // coefficient for walking calorie calculation
)

var (
	ErrSteps    = errors.New("Incorrect steps parameter")
	ErrDuration = errors.New("Incorrect duration parameter")
	ErrCalor    = errors.New("Invalid calorie parameter")
	ErrWeight   = errors.New("Incorrect weight parameter")
	ErrHeight   = errors.New("Incorrect height parameter")
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, ErrSteps
	}
	if duration <= 0 {
		return 0, ErrDuration
	}
	if weight <= 0 {
		return 0, ErrWeight
	}
	if height <= 0 {
		return 0, ErrHeight
	}

	meanSpeed := MeanSpeed(steps, height, duration)

	calories := (weight * meanSpeed * duration.Minutes()) / minInH

	if calories <= 0 {
		return 0, fmt.Errorf("Negative amount of calories %v", ErrCalor)
	}

	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, ErrSteps
	}
	if duration <= 0 {
		return 0, ErrDuration
	}
	if weight <= 0 {
		return 0, ErrWeight
	}
	if height <= 0 {
		return 0, ErrHeight
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	return (weight * meanSpeed * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	distanc := Distance(steps, height)
	return distanc / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient
	stepsFloat := float64(steps) * stepLength
	return stepsFloat / mInKm
}
