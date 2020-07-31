package bmi

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	BMITooLow  = errors.New("BMI too low!")
	BMITooHigh = errors.New("BMI too high!")
)

func Calculate(height, weight string) (bmi float64, desc string, err error) {
	bmi, desc, err = calculate(height, weight)
	if err != nil {
		counts.registerError()
	}
	counts.register(bmi)
	return
}

func CalculateWithoutStats(height, weight string) (bmi float64, desc string, err error) {
	return calculate(height, weight)
}

// calc calculates the BMI and a description
func calculate(height, weight string) (float64, string, error) {
	h, err := strconv.ParseUint(height, 10, 64)
	if err != nil {
		return 0, "", errors.New("Cannot convert first argument to a number")
	}
	w, err := strconv.ParseUint(weight, 10, 64)
	if err != nil {
		return 0, "", errors.New("Cannot convert second argument to a number")
	}
	bmi := float64(w) / ((float64(h) / 100.0) * (float64(h) / 100.0))
	if bmi < 10 {
		return 0, "", BMITooLow
	}
	if bmi > 50 {
		return 0, "", BMITooHigh
	}
	bmiLevel := DescribeLevel(bmi)
	desc := fmt.Sprintf("BMI for %vcm and %vkg is %2.1f (%v)", h, w, bmi, bmiLevel)
	return bmi, desc, nil
}

// Describe describes a calculated BMI with one of the following categories: UNDERWEIGHT, NORMAL, OVERWEIGHT, OBESE.
func DescribeLevel(bmi float64) (desc string) {
	switch {
	case bmi < 18.5:
		desc = "UNDERWEIGHT"
	case bmi < 25:
		desc = "NORMAL"
	case bmi < 30:
		desc = "OVERWEIGHT"
	default:
		desc = "OBESE"
	}
	return
}
