package main

import (
	"os"
	"testing"
	"time"
)

func TestCheckIfAlreadySentToday(t *testing.T) {
	// Use a test-specific tracking file
	testTrackingFile := ".test_last_sent_date"

	// Clean up after test
	defer os.Remove(testTrackingFile)

	// Override the tracking file for testing
	// Note: In real implementation, we'd need to make trackingFile configurable
	// For now, we'll test the logic manually

	t.Run("First time sending", func(t *testing.T) {
		// Remove any existing tracking file
		os.Remove(".last_sent_date")
		defer os.Remove(".last_sent_date")

		today := time.Now()
		alreadySent, err := checkIfAlreadySentToday(today)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if alreadySent {
			t.Error("Expected alreadySent to be false on first run")
		}
	})

	t.Run("Already sent today", func(t *testing.T) {
		// Remove any existing tracking file
		os.Remove(".last_sent_date")
		defer os.Remove(".last_sent_date")

		today := time.Now()

		// Record that we sent today
		err := recordSentDate(today)
		if err != nil {
			t.Fatalf("Failed to record sent date: %v", err)
		}

		// Check if already sent
		alreadySent, err := checkIfAlreadySentToday(today)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if !alreadySent {
			t.Error("Expected alreadySent to be true after recording today's date")
		}
	})

	t.Run("Sent yesterday, not today", func(t *testing.T) {
		// Remove any existing tracking file
		os.Remove(".last_sent_date")
		defer os.Remove(".last_sent_date")

		yesterday := time.Now().AddDate(0, 0, -1)
		today := time.Now()

		// Record that we sent yesterday
		err := recordSentDate(yesterday)
		if err != nil {
			t.Fatalf("Failed to record sent date: %v", err)
		}

		// Check if already sent today
		alreadySent, err := checkIfAlreadySentToday(today)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if alreadySent {
			t.Error("Expected alreadySent to be false when last sent was yesterday")
		}
	})
}

func TestRecordSentDate(t *testing.T) {
	// Remove any existing tracking file
	os.Remove(".last_sent_date")
	defer os.Remove(".last_sent_date")

	today := time.Now()
	err := recordSentDate(today)

	if err != nil {
		t.Fatalf("Failed to record sent date: %v", err)
	}

	// Verify the file was created and contains the correct date
	data, err := os.ReadFile(".last_sent_date")
	if err != nil {
		t.Fatalf("Failed to read tracking file: %v", err)
	}

	expectedDate := today.Format("2006-01-02")
	actualDate := string(data[:len(data)-1]) // Remove trailing newline

	if actualDate != expectedDate {
		t.Errorf("Expected date %s, got %s", expectedDate, actualDate)
	}
}
