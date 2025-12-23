package main

import (
	"testing"
	"time"
)

func TestConvertMarkdownToHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple italic",
			input:    "*italic text*",
			expected: "<i>italic text</i>",
		},
		{
			name:     "Simple bold",
			input:    "**bold text**",
			expected: "<b>bold text</b>",
		},
		{
			name:     "Russian text with guillemets",
			input:    "*«Чтобы оставаться трезвым».*",
			expected: "<i>«Чтобы оставаться трезвым».</i>",
		},
		{
			name:     "Multiple italics",
			input:    "*first* and *second*",
			expected: "<i>first</i> and <i>second</i>",
		},
		{
			name:     "Mixed formatting",
			input:    "*italic* and **bold**",
			expected: "<i>italic</i> and <b>bold</b>",
		},
		{
			name:     "No formatting",
			input:    "Plain text without formatting",
			expected: "Plain text without formatting",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertMarkdownToHTML(tt.input)
			if result != tt.expected {
				t.Errorf("convertMarkdownToHTML() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFormatDateHeader(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected string
	}{
		{
			name:     "December 23",
			date:     time.Date(2024, 12, 23, 0, 0, 0, 0, time.UTC),
			expected: "<b>23 Декабря</b>\n\n",
		},
		{
			name:     "January 1",
			date:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: "<b>1 Января</b>\n\n",
		},
		{
			name:     "May 9",
			date:     time.Date(2024, 5, 9, 0, 0, 0, 0, time.UTC),
			expected: "<b>9 Мая</b>\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatDateHeader(tt.date)
			if result != tt.expected {
				t.Errorf("formatDateHeader() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestGetRussianMonthName(t *testing.T) {
	expected := map[time.Month]string{
		time.January:   "Января",
		time.February:  "Февраля",
		time.March:     "Марта",
		time.April:     "Апреля",
		time.May:       "Мая",
		time.June:      "Июня",
		time.July:      "Июля",
		time.August:    "Августа",
		time.September: "Сентября",
		time.October:   "Октября",
		time.November:  "Ноября",
		time.December:  "Декабря",
	}

	for month, expectedName := range expected {
		t.Run(expectedName, func(t *testing.T) {
			result := getRussianMonthName(month)
			if result != expectedName {
				t.Errorf("getRussianMonthName(%v) = %q, want %q", month, result, expectedName)
			}
		})
	}
}
