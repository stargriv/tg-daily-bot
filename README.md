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
./tg-daily-bot
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

2. Create a systemd service file `/etc/systemd/system/tg-daily-bot.service`:
   ```ini
   [Unit]
   Description=Telegram Daily Bot
   After=network.target

   [Service]
   Type=simple
   User=youruser
   WorkingDirectory=/path/to/tg-daily-bot
   ExecStart=/path/to/tg-daily-bot/tg-daily-bot
   Restart=always
   RestartSec=10

   [Install]
   WantedBy=multi-user.target
   ```

3. Enable and start the service:
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable tg-daily-bot
   sudo systemctl start tg-daily-bot
   ```

4. Check status:
   ```bash
   sudo systemctl status tg-daily-bot
   ```

5. View logs:
   ```bash
   sudo journalctl -u tg-daily-bot -f
   ```

### Using screen (Alternative)

```bash
screen -S tg-daily-bot
./tg-daily-bot
# Press Ctrl+A then D to detach
```

To reattach later:
```bash
screen -r tg-daily-bot
```

## Deployment with GitHub Actions

You can also run the bot using GitHub Actions to send daily messages without maintaining a VPS.

### Setup

1. **Fork or push this repository to GitHub**

2. **Store your daily reflections** (choose one):

   **Option A: Include in repository (simpler)**
   - Keep your `files/daily_reflections_structured.md` file in the repo
   - Note: The file is git-ignored by default. Remove `files/` from `.gitignore` to commit it

   **Option B: Use private GitHub Gist (recommended for private content)**
   - Store your markdown file in a private gist
   - See [GIST_SETUP.md](GIST_SETUP.md) for detailed instructions
   - This keeps your reflections separate from the code repository

3. **Configure GitHub Secrets**

   Go to your repository Settings → Secrets and variables → Actions → New repository secret

   **Required secrets:**
   - `BOT_TOKEN`: Your Telegram bot token from BotFather
   - `CHAT_ID`: Your group chat ID (negative number)
   - `MESSAGE_THREAD_ID`: The topic/thread ID

   **Additional secrets (only if using Gist - Option B):**
   - `GIST_ID`: Your private gist ID
   - `GIST_TOKEN`: Personal access token with `gist` scope

4. **Customize the schedule** (optional)

   Edit `.github/workflows/daily-message.yml` to change the schedule:
   ```yaml
   schedule:
     - cron: '0 9 * * *'  # Change this to your desired time (UTC)
   ```

4. **Enable GitHub Actions**

   - Go to the "Actions" tab in your repository
   - Enable workflows if prompted
   - The workflow will run automatically according to the schedule
   - You can also trigger it manually using "Run workflow"

### How it Works

The GitHub Actions workflow:
1. Runs on a schedule (default: 9 AM UTC daily)
2. Checks out your code and daily reflections file
3. Builds and runs the bot with `RUN_ONCE=true`
4. Sends today's reflection to your Telegram topic
5. Exits automatically

**Note:** GitHub Actions uses UTC time. Adjust the cron schedule accordingly for your timezone.

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `BOT_TOKEN` | Yes | - | Your Telegram bot token from BotFather |
| `CHAT_ID` | Yes | - | The chat ID of your group (negative number) |
| `MESSAGE_THREAD_ID` | Yes | - | The topic/thread ID where messages will be sent |
| `CRON_SCHEDULE` | No | `0 9 * * *` | Cron expression for scheduling messages (only used when `RUN_ONCE` is false) |
| `FILE_PATH` | No | `files/daily_reflections_structured.md` | Path to markdown file with daily reflections (format: `## MM-DD`) |
| `MESSAGE` | No | `Hello from your scheduled bot!` | Fallback message if no reflection found for today |
| `RUN_ONCE` | No | `false` | Set to `true` to send one message and exit (for GitHub Actions/single execution) |

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
