package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

// парсинг
func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	info := strings.Split(data, ",")
	if len(info) != 2 {
		return 0, time.Duration(0), fmt.Errorf("некорректные данные parsePackage")
	}
	st, err := strconv.Atoi(info[0])
	if err != nil {
		return 0, time.Duration(0), err
	}
	if st <= 0 {
		return 0, time.Duration(0), fmt.Errorf("количество шагов должно быть больше 0")
	}
	t, err := time.ParseDuration(info[1])
	if err != nil {
		return 0, time.Duration(0), err
	}
	if t <= 0 {
		return 0, time.Duration(0), fmt.Errorf("время тренировки должно быть больше 0")
	}
	return st, t, nil
}

// DayActionInfo вычисляет дистанцию в километрах
//
//	и количество потраченных калорий
func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	if weight <= 0 || height <= 0 {
		log.Println("рост и вес должен быть больше 0")
		return ""
	}
	step, times, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	dist := float64(step) * stepLength / float64(mInKm)
	c, err := spentcalories.WalkingSpentCalories(step, weight, height, times) //WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error)
	if err != nil {
		log.Println(err)
		return ""
	}
	str := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", step, dist, c)
	return str
}
