package daysteps

import (
	"errors"
	"fmt"
	"math"
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
		return errors.New("пустая строка данных")
	}

	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат данных")
	}

	stepStr := parts[0]
	durStr := parts[1]

	// шаги: строгое целое положительное, без пробелов
	if strings.TrimSpace(stepStr) != stepStr {
		return errors.New("неверный формат шагов")
	}
	steps, err := strconv.Atoi(stepStr)
	if err != nil || steps <= 0 {
		return errors.New("неверный формат шагов")
	}
	ds.Steps = steps

	// продолжительность: поддержка дробных значений (1.5h, 30.5m), но > 0
	if strings.Contains(durStr, " ") {
		return errors.New("неверный формат продолжительности")
	}
	if !strings.ContainsAny(durStr, "hm") {
		return errors.New("неверный формат продолжительности")
	}

	var total float64
	rest := durStr

	for rest != "" {
		i := strings.IndexAny(rest, "hm")
		if i == -1 {
			return errors.New("неверный формат продолжительности")
		}
		numStr := rest[:i]
		unit := rest[i : i+1]
		rest = rest[i+1:]

		val, err := strconv.ParseFloat(numStr, 64)
		if err != nil || val < 0 {
			return errors.New("неверный формат продолжительности")
		}

		switch unit {
		case "h":
			total += val * float64(time.Hour)
		case "m":
			total += val * float64(time.Minute)
		default:
			return errors.New("неверная единица измерения")
		}
	}

	if total <= 0 {
		return errors.New("некорректная продолжительность")
	}

	ds.Duration = time.Duration(math.Round(total))
	return nil
}

// ActionInfo возвращает строку с шагами, дистанцией и калориями
// ActionInfo возвращает строку с шагами, дистанцией и калориями
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Weight <= 0 || ds.Height <= 0 {
		return "", errors.New("недостаточно данных для расчёта")
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
