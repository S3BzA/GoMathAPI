package math

// Adds numbers and returns the sum
func Add(numbers ...float64) (sum float64) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// Subtracts numbers and returns the difference
func Subtract(numbers ...float64) (difference float64) {
	for _, number := range numbers {
		difference -= number
	}
	return
}

// Multiplies numbers and returns the product
func Multiply(numbers ...float64) (product float64) {
	for _, number := range numbers {
		product *= number
	}
	return
}

// Calculates the nth power (int) of the base number (float) and returns the power
func Power(base float64, exponent int64) (product float64) {
	product = 1
	
	if exponent == 0 {
		return
	}

	if exponent < 0 {
		base = 1 / base
		exponent *= -1
	}

	for i := 0; i < int(exponent); i++ {
		product *= base
	}

	return
}

// Divides a numerator by a denominator and returns the quotient
func Divide(numerator, denominator float64) (quotient float64) {
	quotient = numerator / denominator
	return
}

// Divides a dividend by a divisor and returns the modulus (remainder)
func Modulo(dividend, divisor int64) (modulus int64) {
	modulus = dividend % divisor
	return
}