package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку вида "678,1h30m" в шаги и длительность
func (ds *DaySteps) Parse(datastring string) error {
	if datastring == "" {
		return errors.New("empty data string")
	}

	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("invalid data format, expected 'steps,duration'")
	}

	stepStr := parts[0]
	durStr := parts[1]

	// шаги
	if strings.TrimSpace(stepStr) != stepStr {
		return errors.New("steps string contains extra spaces")
	}
	steps, err := strconv.Atoi(stepStr)
	if err != nil {
		return fmt.Errorf("invalid steps value: %w", err)
	}
	if steps <= 0 {
		return errors.New("steps must be greater than zero")
	}
	ds.Steps = steps

	// длительность
	if strings.Contains(durStr, " ") {
		return errors.New("duration contains spaces")
	}

	dur, err := time.ParseDuration(durStr)
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}
	if dur <= 0 {
		return errors.New("duration must be greater than zero")
	}

	ds.Duration = dur
	return nil
}

// ActionInfo возвращает строку с шагами, дистанцией и калориями
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", errors.New("steps must be greater than zero")
	}
	if ds.Duration <= 0 {
		return "", errors.New("duration must be greater than zero")
	}
	if ds.Weight <= 0 {
		return "", errors.New("weight must be greater than zero")
	}
	if ds.Height <= 0 {
		return "", errors.New("height must be greater than zero")
	}

	// длина шага (в метрах)
	stepLength := ds.Height * 0.45
	// дистанция (в км)
	distance := float64(ds.Steps) * stepLength / 1000.0
	// калории (0.5 * вес * дистанция)
	calories := distance * ds.Weight * 0.5

	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}
