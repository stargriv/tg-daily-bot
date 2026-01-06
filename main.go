package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	tele "gopkg.in/telebot.v3"
)

type Config struct {
	BotToken      string
	ChatID        int64
	MessageThread int
	CronSchedule  string
	Message       string
	FilePath      string
	RunOnce       bool
}

func loadConfig() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	chatID, err := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	if err != nil {
		return nil, err
	}

	messageThread, err := strconv.Atoi(os.Getenv("MESSAGE_THREAD_ID"))
	if err != nil {
		return nil, err
	}

	return &Config{
		BotToken:      os.Getenv("BOT_TOKEN"),
		ChatID:        chatID,
		MessageThread: messageThread,
		CronSchedule:  getEnvOrDefault("CRON_SCHEDULE", "0 9 * * *"), // Default: 9 AM daily
		Message:       getEnvOrDefault("MESSAGE", "Hello from your scheduled bot!"),
		FilePath:      getEnvOrDefault("FILE_PATH", "files/daily_reflections_structured.md"),
		RunOnce:       os.Getenv("RUN_ONCE") == "true",
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// extractDailyReflection reads the markdown file and extracts the content for the given date
func extractDailyReflection(filePath string, date time.Time) (string, error) {
	// Read the entire file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// Format the date as MM-DD
	dateStr := date.Format("01-02")

	// Create regex pattern to match the date header and capture content until next header
	// Pattern: ## MM-DD followed by any content until the next ## header
	pattern := fmt.Sprintf(`(?s)## %s\s*\n(.*?)(?:\n## \d{2}-\d{2}|\z)`, regexp.QuoteMeta(dateStr))
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(string(content))
	if len(matches) < 2 {
		return "", fmt.Errorf("no reflection found for date %s", dateStr)
	}

	// Extract and clean up the content
	reflection := strings.TrimSpace(matches[1])

	return reflection, nil
}

// getRussianMonthName returns the Russian genitive month name for dates
func getRussianMonthName(month time.Month) string {
	months := map[time.Month]string{
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
	return months[month]
}

// formatDateHeader returns a formatted date header in Russian (e.g., "23 Декабря")
func formatDateHeader(date time.Time) string {
	return fmt.Sprintf("<b>%d %s</b>\n\n", date.Day(), getRussianMonthName(date.Month()))
}

// getDailyMessage returns today's reflection with a date header, or falls back to the default message
func getDailyMessage(filePath string, fallbackMessage string) string {
	today := time.Now()
	reflection, err := extractDailyReflection(filePath, today)
	if err != nil {
		log.Printf("Warning: Could not extract daily reflection: %v. Using fallback message.", err)
		return fallbackMessage
	}

	// Prepend the date header
	header := formatDateHeader(today)
	return header + reflection
}

// convertMarkdownToHTML converts simple markdown formatting to HTML for Telegram
func convertMarkdownToHTML(text string) string {
	// Convert **text** to <b>text</b> for bold first (before single asterisks)
	re := regexp.MustCompile(`\*\*([^*]+)\*\*`)
	text = re.ReplaceAllString(text, "<b>$1</b>")

	// Convert *text* to <i>text</i> for italics
	re = regexp.MustCompile(`\*([^*]+)\*`)
	text = re.ReplaceAllString(text, "<i>$1</i>")

	return text
}

func sendMessageToTopic(bot *tele.Bot, chatID int64, threadID int, message string) error {
	chat := &tele.Chat{ID: chatID}

	// Convert markdown to HTML for better compatibility
	htmlMessage := convertMarkdownToHTML(message)

	// Try sending with HTML formatting
	_, err := bot.Send(chat, htmlMessage, &tele.SendOptions{
		ThreadID:  threadID,
		ParseMode: tele.ModeHTML,
	})

	// If HTML fails, try sending as plain text
	if err != nil {
		log.Printf("HTML parsing failed, sending as plain text: %v", err)
		_, err = bot.Send(chat, message, &tele.SendOptions{
			ThreadID: threadID,
		})
	}

	if err != nil {
		return err
	}

	log.Printf("Message sent successfully to chat %d, topic %d", chatID, threadID)
	return nil
}

func main() {
	log.Println("Starting Telegram bot...")

	// Load configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create bot instance
	pref := tele.Settings{
		Token:  config.BotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Me.Username)

	// Send message with today's reflection
	log.Println("Sending message with today's reflection...")
	dailyMessage := getDailyMessage(config.FilePath, config.Message)
	if err := sendMessageToTopic(bot, config.ChatID, config.MessageThread, dailyMessage); err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// If RUN_ONCE mode, exit after sending the message
	if config.RunOnce {
		log.Println("RUN_ONCE mode enabled. Message sent successfully. Exiting.")
		return
	}

	// Set up cron scheduler for continuous operation
	c := cron.New()
	_, err = c.AddFunc(config.CronSchedule, func() {
		log.Println("Executing scheduled message...")
		// Get today's reflection each time the cron job runs
		dailyMessage := getDailyMessage(config.FilePath, config.Message)
		if err := sendMessageToTopic(bot, config.ChatID, config.MessageThread, dailyMessage); err != nil {
			log.Printf("Error sending scheduled message: %v", err)
		}
	})

	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}

	c.Start()
	log.Printf("Scheduler started with cron: %s", config.CronSchedule)
	log.Println("Bot is running. Press Ctrl+C to stop.")

	// Keep the bot running
	bot.Start()
}
