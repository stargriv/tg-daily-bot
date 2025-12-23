# Telegram Topic Bot

A lightweight Telegram bot that sends daily reflections from a markdown file to group chat topics. Perfect for running on a tiny VPS.

## Features

- Automatically extracts and sends daily reflections based on today's date
- Parses markdown files with date-based headers (format: `## MM-DD`)
- Send messages to specific topics in Telegram group chats
- Scheduled message posting using cron expressions
- Minimal resource usage
- Easy configuration via environment variables
- Fallback to custom message if no reflection found for today

## Prerequisites

- Go 1.21 or higher
- A Telegram bot token from [@BotFather](https://t.me/BotFather)
- A Telegram group with topics enabled
- The bot must be added to the group as an admin

## Setup

### 1. Create a Telegram Bot

1. Open Telegram and search for [@BotFather](https://t.me/BotFather)
2. Send `/newbot` and follow the instructions
3. Save the bot token provided by BotFather

### 2. Get Your Chat ID

1. Add your bot to your Telegram group
2. Send a message in the group
3. Visit `https://api.telegram.org/bot<YOUR_BOT_TOKEN>/getUpdates` (replace `<YOUR_BOT_TOKEN>` with your actual token)
4. Look for the `"chat":{"id":` field - this is your chat ID (will be a negative number for groups)

### 3. Get the Topic/Thread ID

1. In your Telegram group, right-click on the topic you want to post to
2. Select "Copy Link"
3. The link will look like: `https://t.me/c/1234567890/1/123`
4. The number after `/1/` (e.g., `123`) is your MESSAGE_THREAD_ID

### 4. Configure the Bot

1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

2. Edit `.env` and fill in your values:
   ```env
   BOT_TOKEN=your_bot_token_here
   CHAT_ID=-1001234567890
   MESSAGE_THREAD_ID=123
   CRON_SCHEDULE=0 9 * * *
   FILE_PATH=files/daily_reflections_structured.md
   MESSAGE=Hello from your scheduled bot!
   ```

### 5. Add Your Daily Reflections File

The bot reads daily reflections from a markdown file. The file should be structured with date headers in `MM-DD` format:

```markdown
## 01-01

*Your reflection or quote here*

Main content for January 1st goes here.
This can be multiple paragraphs.

*Closing thought or meditation*

## 01-02

Content for January 2nd...
```

**Requirements:**
- Place your markdown file in the `files/` directory
- Each date entry must start with a header: `## MM-DD` (e.g., `## 01-15` for January 15)
- Content continues until the next date header
- Configure the file path in `.env` using the `FILE_PATH` variable

The bot will automatically extract and send the reflection for the current date.

### 6. Build and Run

```bash
# Build the bot
go build

# Run the bot
./saa-bot
```

## Cron Schedule Format

The cron schedule uses the standard 5-field format:

```
* * * * *
│ │ │ │ │
│ │ │ │ └─── Day of week (0-7, 0 and 7 are Sunday)
│ │ │ └───── Month (1-12)
│ │ └─────── Day of month (1-31)
│ └───────── Hour (0-23)
└─────────── Minute (0-59)
```

### Common Examples

- `0 9 * * *` - Every day at 9:00 AM
- `0 */6 * * *` - Every 6 hours
- `0 9 * * 1` - Every Monday at 9:00 AM
- `*/30 * * * *` - Every 30 minutes
- `0 0 * * 0` - Every Sunday at midnight
- `0 12 1 * *` - First day of every month at noon

## Deployment on VPS

### Using systemd (Recommended)

1. Build the bot:
   ```bash
   go build
   ```

2. Create a systemd service file `/etc/systemd/system/saa-bot.service`:
   ```ini
   [Unit]
   Description=Telegram SAA Bot
   After=network.target

   [Service]
   Type=simple
   User=youruser
   WorkingDirectory=/path/to/saa-bot
   ExecStart=/path/to/saa-bot/saa-bot
   Restart=always
   RestartSec=10

   [Install]
   WantedBy=multi-user.target
   ```

3. Enable and start the service:
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable saa-bot
   sudo systemctl start saa-bot
   ```

4. Check status:
   ```bash
   sudo systemctl status saa-bot
   ```

5. View logs:
   ```bash
   sudo journalctl -u saa-bot -f
   ```

### Using screen (Alternative)

```bash
screen -S saa-bot
./saa-bot
# Press Ctrl+A then D to detach
```

To reattach later:
```bash
screen -r saa-bot
```

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `BOT_TOKEN` | Yes | - | Your Telegram bot token from BotFather |
| `CHAT_ID` | Yes | - | The chat ID of your group (negative number) |
| `MESSAGE_THREAD_ID` | Yes | - | The topic/thread ID where messages will be sent |
| `CRON_SCHEDULE` | No | `0 9 * * *` | Cron expression for scheduling messages |
| `FILE_PATH` | No | `files/daily_reflections_structured.md` | Path to markdown file with daily reflections (format: `## MM-DD`) |
| `MESSAGE` | No | `Hello from your scheduled bot!` | Fallback message if no reflection found for today |

## Troubleshooting

### Bot can't send messages
- Ensure the bot is added to the group as an admin
- Verify the bot has permission to post in topics
- Check that the chat ID and thread ID are correct

### Bot fails to start
- Verify your `.env` file exists and has correct values
- Check that all required environment variables are set
- Ensure the bot token is valid

### Messages not being sent at scheduled time
- Verify your cron expression is correct
- Check the server timezone
- View logs to see if there are any errors

## License

MIT
