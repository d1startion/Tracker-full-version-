package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку вида "3456,Walking,3h00m"
func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("invalid data format, expected 'steps,type,duration'")
	}

	// шаги
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid steps value: %w", err)
	}
	if steps <= 0 {
		return errors.New("steps must be greater than zero")
	}
	t.Steps = steps

	// тип тренировки
	t.TrainingType = parts[1]

	// продолжительность
	dur, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("invalid duration value: %w", err)
	}
	if dur <= 0 {
		return errors.New("duration must be greater than zero")
	}
	t.Duration = dur

	return nil
}

// ActionInfo возвращает информацию о тренировке в виде строки
func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("failed to calculate walking calories: %w", err)
		}
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("failed to calculate running calories: %w", err)
		}
	default:
		return "", fmt.Errorf("unknown training type: %s", t.TrainingType)
	}

	str := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		dist,
		speed,
		calories,
	)

	return str, nil
}
