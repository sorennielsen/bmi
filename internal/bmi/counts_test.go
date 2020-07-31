package bmi

import "testing"

func TestNew(t *testing.T) {
	// Given
	// The Void

	// When
	c := NewCounts()

	// Then
	if c.Total != 0 {
		t.Errorf("Total not correct. Want %v, got %v", 0, c.Total)
	}

	if len(c.Calculations) != 0 {
		t.Errorf("Number of calculations not correct. Want %v, got %v", 0, len(c.Calculations))
	}

	if c.Errors != 0 {
		t.Errorf("Number of errors not correct. Want %v, got %v", 0, c.Errors)
	}

	if c.Average != 0 {
		t.Errorf("Average not correct. Want %v, got %v", 0, c.Average)
	}
}

func TestErrors(t *testing.T) {
	// Given
	c := NewCounts()

	// When
	c.registerError()

	// Then
	if c.Errors != 1 {
		t.Errorf("Number of errors not correct. Want %v, got %v", 0, c.Errors)
	}
}

func TestCalculationForUnderweight(t *testing.T) {
	// Given
	c := NewCounts()

	// When
	c.register(15.0) // UNDERWEIGHT

	// Then
	if c.Errors != 0 {
		t.Errorf("Number of errors not correct. Want %v, got %v", 0, c.Errors)
	}

	count, ok := c.Calculations["UNDERWEIGHT"]
	if !ok {
		t.Errorf("Did not register any UNDERWEIGHT calculations. Want %v", 1)
		return
	}

	if count != 1 {
		t.Errorf("Number of UNDERWEIGHT not corect. Want %v, got %v", 1, count)
	}
}

func TestCalculationForNormal(t *testing.T) {
	// Given
	c := NewCounts()

	// When
	c.register(22.0) // NORMAL

	// Then
	if c.Errors != 0 {
		t.Errorf("Number of errors not correct. Want %v, got %v", 0, c.Errors)
	}

	count, ok := c.Calculations["NORMAL"]
	if !ok {
		t.Errorf("Did not register any NORMAL calculations. Want %v", 1)
		return
	}

	if count != 1 {
		t.Errorf("Number of NORMAL not corect. Want %v, got %v", 1, count)
	}
}

func TestCalculationForOverweight(t *testing.T) {
	// Given
	c := NewCounts()

	// When
	c.register(26.0) // OVERWEIGHT

	// Then
	if c.Errors != 0 {
		t.Errorf("Number of errors not correct. Want %v, got %v", 0, c.Errors)
	}

	count, ok := c.Calculations["OVERWEIGHT"]
	if !ok {
		t.Errorf("Did not register any OVERWEIGHT calculations. Want %v", 1)
		return
	}

	if count != 1 {
		t.Errorf("Number of OVERWEIGHT not corect. Want %v, got %v", 1, count)
	}
}

func TestCalculationForObese(t *testing.T) {
	// Given
	c := NewCounts()

	// When
	c.register(32.0) // OBESE

	// Then
	if c.Errors != 0 {
		t.Errorf("Number of errors not correct. Want %v, got %v", 0, c.Errors)
	}

	count, ok := c.Calculations["OBESE"]
	if !ok {
		t.Errorf("Did not register any OBESE calculations. Want %v", 1)
		return
	}

	if count != 1 {
		t.Errorf("Number of OVERWEIGHT not corect. Want %v, got %v", 1, count)
	}
}
