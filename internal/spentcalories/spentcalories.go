package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

// парсинг
func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	info := strings.Split(data, ",")
	if len(info) != 3 {
		return 0, "", time.Duration(0), fmt.Errorf("некорректные данные parseTraining")
	}
	s, err := strconv.Atoi(info[0])
	if err != nil {
		return 0, "", time.Duration(0), err
	}
	if s <= 0 {
		return 0, "", time.Duration(0), fmt.Errorf("некорректные количество шагов")
	}
	t, err := time.ParseDuration(info[2])
	if err != nil {
		return 0, "", time.Duration(0), err
	}
	if t <= 0 {
		return 0, "", time.Duration(0), fmt.Errorf("некорректное время")
	}
	return s, info[1], t, nil
}

// определяет пройденное расстояние
func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	l := stepLengthCoefficient * height
	d := l * float64(steps)
	return d / mInKm
}

// определение средней скорости
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	d := distance(steps, height)
	return d / float64(duration.Hours())
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, exercise, times, err := parseTraining(data)
	if err != nil {
		return "", err
	}
	dst := distance(steps, height)
	avg := meanSpeed(steps, height, times)

	switch {
	case exercise == "Бег":
		s, err := RunningSpentCalories(steps, weight, height, times)
		if err != nil {
			return "", err
		}
		str := fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", float64(times.Hours()), dst, avg, s)
		return str, nil
	case exercise == "Ходьба":
		s, err := WalkingSpentCalories(steps, weight, height, times)
		if err != nil {
			return "", err
		}
		str := fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", float64(times.Hours()), dst, avg, s)
		return str, nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

// RunningSpentCalories расчитывает потраченные калории
// при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 { ////////////////////
		return 0, fmt.Errorf("некорректные данные RunningSpentCalories")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	calories := (duration.Minutes() * weight * avgSpeed) / minInH
	return calories, nil
}

// WalkingSpentCalories расчитывает потраченные калории
// при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 {
		return 0, fmt.Errorf("некорректные данные WalkingSpentCalories")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	calories := (duration.Minutes() * weight * avgSpeed) / minInH
	return calories * walkingCaloriesCoefficient, nil
}
