package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorrectFizzBuzz(t *testing.T) {
	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	got, _ := FizzBuzz("Fizz", "Buzz", 3, 5, 15)
	assert.Equal(t, want, got, "FizzBuzz(15) \n got: \n%v \n want: \n%v", got, want)
}

func TestSameFizzAndBuzzMultiple(t *testing.T) {
	_, err := FizzBuzz("Fizz", "Buzz", 3, 3, 15)
	assert.NotNil(t, err, "Error should be throw when we give the same multiple [%v]", err)
	want := "InvalidParameters: The 2 multiples should be different"
	got := err.Error()
	assert.Equal(t, want, got, "Not the good error message [want:%v,got:%v]", want, got)
}

func TestLimitLowerThanStart(t *testing.T) {
	_, err := FizzBuzz("Fizz", "Buzz", 3, 5, -15)
	assert.NotNil(t, err, "Error should be throw when limit is lower than 1 [%v]", err)

	want := "InvalidParameters: Limit should be greater than 1"
	got := err.Error()
	assert.Equal(t, want, got, "Not the good error message [want:%v,got:%v]", want, got)
}

func TestSameFizzAndBuzzWord(t *testing.T) {
	_, err := FizzBuzz("Fizz", "Fizz", 3, 5, 15)
	assert.NotNil(t, err, "Error should be throw when we give the same strings [%v]", err)

	want := "InvalidParameters: The 2 strings should be different"
	got := err.Error()
	assert.Equal(t, want, got, "Not the good error message [want:%v,got:%v]", want, got)
}
