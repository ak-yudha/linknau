// process_test.go
package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockProcessor adalah implementasi mock dari Processor untuk pengujian.
type MockProcessor struct {
	mock.Mock
}

func (m *MockProcessor) Process(number int) int {
	args := m.Called(number)
	return args.Int(0)
}

// TestProcessNumbers menguji fungsi ProcessNumbers
func TestProcessNumbers(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	// Membuat mock processor
	mockProcessor := new(MockProcessor)
	mockProcessor.On("Process", 1).Return(1)
	mockProcessor.On("Process", 2).Return(4)
	mockProcessor.On("Process", 3).Return(9)
	mockProcessor.On("Process", 4).Return(16)
	mockProcessor.On("Process", 5).Return(25)

	// Memanggil fungsi
	results := ProcessNumbers(numbers, mockProcessor)

	// Hasil yang diharapkan
	expectedResults := []int{1, 4, 9, 16, 25}

	// Verifikasi hasil
	assert.ElementsMatch(t, expectedResults, results)

	// Verifikasi bahwa mock dipanggil dengan argumen yang diharapkan
	mockProcessor.AssertExpectations(t)
}
