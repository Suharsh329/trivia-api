package utils

import (
	"testing"
)

func TestCreateQueryFilters(t *testing.T) {
	// Arrange
	tests := []struct {
		filters    map[string]interface{}
		exactMatch bool
		expected   string
		accepted   string
	}{
		{
			filters:    map[string]interface{}{"name": "John", "age": 30},
			exactMatch: true,
			expected:   " WHERE name = 'John' AND age = 30",
			accepted:   " WHERE age = 30 AND name = 'John'",
		},
		{
			filters:    map[string]interface{}{"name": "John", "age": 30},
			exactMatch: false,
			expected:   " WHERE name LIKE '%John%' AND age = 30",
			accepted:   " WHERE age = 30 AND name LIKE '%John%'",
		},
		{
			filters:    map[string]interface{}{"category": "books", "price": 19.99},
			exactMatch: true,
			expected:   " WHERE category = 'books' AND price = 19.99",
			accepted:   " WHERE price = 19.99 AND category = 'books'",
		},
		{
			filters:    map[string]interface{}{"tags": []interface{}{"fiction", "bestseller"}},
			exactMatch: true,
			expected:   " WHERE tags IN ('fiction', 'bestseller')",
			accepted:   " WHERE tags IN ('bestseller', 'fiction')",
		},
	}

	for _, test := range tests {
		// Act
		result := CreateQueryFilters(test.filters, test.exactMatch)
		// Assert
		if result != test.expected && result != test.accepted {
			t.Errorf("Got %v; want %v, or %v", result, test.expected, test.accepted)
		}
	}
}
