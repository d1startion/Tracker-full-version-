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

// Parse разбирает строку вида "3456,Ходьба,3h00m"
func (t *Training) Parse(datastring string) (err error) {
	parsStr := strings.Split(datastring, ",")
	if len(parsStr) != 3 {
		return errors.New("неверный формат данных")
	}

	// шаги
	t.Steps, err = strconv.Atoi(parsStr[0])
	if err != nil || t.Steps <= 0 {
		return errors.New("неверный формат данных")
	}

	// тип тренировки
	t.TrainingType = parsStr[1]

	// продолжительность
	t.Duration, err = time.ParseDuration(parsStr[2])
	if err != nil || t.Duration <= 0 {
		return errors.New("неверный формат данных")
	}

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
			return "", err
		}
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
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
