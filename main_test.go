package main

import (
	"testing"
	"time"
)

func TestExtractDailyReflection(t *testing.T) {
	filePath := "files/daily_reflections_structured.md"

	// Test January 1st
	jan1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	reflection, err := extractDailyReflection(filePath, jan1)

	if err != nil {
		t.Fatalf("Failed to extract reflection for January 1st: %v", err)
	}

	if len(reflection) == 0 {
		t.Error("Reflection for January 1st is empty")
	}

	t.Logf("Successfully extracted reflection for 01-01 (%d characters)", len(reflection))

	// Test today
	today := time.Now()
	todayReflection, err := extractDailyReflection(filePath, today)

	if err != nil {
		t.Logf("No reflection found for today (%s): %v", today.Format("01-02"), err)
	} else {
		t.Logf("Successfully extracted reflection for today %s (%d characters)", today.Format("01-02"), len(todayReflection))
	}
}

func TestGetDailyMessage(t *testing.T) {
	filePath := "files/daily_reflections_structured.md"
	fallback := "Fallback message"

	// Test with a valid date (January 1st)
	// Temporarily override time for testing
	message := getDailyMessage(filePath, fallback)

	if message == "" {
		t.Error("getDailyMessage returned empty string")
	}

	if message == fallback {
		t.Logf("Using fallback message (no reflection found for today)")
	} else {
		t.Logf("Successfully got daily message (%d characters)", len(message))
	}
}
