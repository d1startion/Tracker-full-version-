package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("недостаточно данных для расчёта")
	}
	distKm := Distance(steps, height)
	calories := distKm * weight * walkingCaloriesCoefficient
	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("недостаточно данных для расчёта")
	}
	distKm := Distance(steps, height)
	calories := distKm * weight
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}
	distKm := Distance(steps, height)
	speedKmh := distKm / duration.Hours()
	return speedKmh
}

func Distance(steps int, height float64) float64 {
	stepSize := height * stepLengthCoefficient
	distMetr := float64(steps) * stepSize
	distKm := distMetr / mInKm
	return distKm
}
