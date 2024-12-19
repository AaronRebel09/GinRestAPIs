package main

func MyFunction(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("both arguments must be non-negative")
	}

	return a + b, nil
}

func TestMyFunction(t *testing.T) {
	// Test case 1: Positive inputs
	sum, err := MyFunction(2, 3)
	assert.Equal(t, 5, sum, "Sum of 2 and 3 is 5")
	assert.NoError(t, err, "There are no errors")

}