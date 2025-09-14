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

func (t *Training) Parse(datastring string) (err error) {
	parsStr := strings.Split(datastring, ",")
	if len(parsStr) != 3 {
		return errors.New("неверный формат данных")
	}
	t.Steps, err = strconv.Atoi(parsStr[0])
	if err != nil {
		return errors.New("неверный формат данных")
	}
	t.TrainingType = parsStr[1]
	t.Duration, err = time.ParseDuration(parsStr[2])
	if err != nil {
		return errors.New("неверный формат данных")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	var errr error
	switch t.TrainingType {
	case "Ходьба":
		calories, errr = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if errr != nil {
			return "", errr
		}
	case "Бег":
		calories, errr = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if errr != nil {
			return "", errr
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %2.f ч.\nДистанция: %2.f км.\nСкорость: %2.f км/ч\nСожгли калорий: %2.f", t.TrainingType, t.Duration.Hours(), dist, speed, calories)
	return str, nil
}
