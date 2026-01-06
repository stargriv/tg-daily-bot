# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- Simplified duplicate detection to use file-based tracking only
- GitHub Actions workflow now uses cache to persist `.last_sent_date` between runs
- Removed complex channel-checking logic (Telegram Bot API limitations)
- Improved logging with emoji indicators for better visibility

### Added
- GitHub Actions cache integration for tracking file persistence
- Cache restore and save steps in workflow

## [1.2.0] - 2026-01-06

### Added
- Duplicate message detection to prevent sending the same date's message multiple times
- `.last_sent_date` tracking file for persistent duplicate detection
- Comprehensive tests for duplicate detection logic
- Helper function `formatDateHeaderPlainText()` for date comparison
- GitHub Actions cache support for state persistence between runs

### Changed
- Bot now checks for duplicates before sending messages
- Records sent date after successful message delivery
- Improved logging for duplicate detection status with emoji indicators
- GitHub Actions workflow now downloads pre-built binary from releases instead of building
- Reduced workflow execution time by eliminating Go build step

### Features
- **File-Based Duplicate Detection**: Simple and reliable tracking
- **GitHub Actions Cache**: Persists tracking file between workflow runs
- **VPS Compatible**: Works across bot restarts
- **Smart Scheduling**: Safe to run multiple times per day without spamming
- **Faster Workflow**: Downloads pre-built binary from releases (~3-5 seconds vs ~30+ seconds for build)

## [1.1.0] - 2026-01-06

### Added
- GitHub Actions workflow for scheduled daily message sending
- `RUN_ONCE` mode to send a single message and exit (perfect for GitHub Actions)
- Private GitHub Gist integration for storing daily reflections separately
- Comprehensive Gist setup guide (GIST_SETUP.md)
- Support for both repository-based and Gist-based reflection storage
- Automatic Gist file download during workflow execution
- Manual workflow trigger support via GitHub Actions UI
- files/README.md documentation for reflection file management

### Changed
- Bot now sends message immediately on startup in RUN_ONCE mode
- Updated README with GitHub Actions deployment section
- Enhanced .env.example with RUN_ONCE configuration
- Improved error handling for initial message sending

### Features
- **GitHub Actions Deployment**: Run the bot on GitHub's infrastructure without a VPS
- **Gist Integration**: Store private content in GitHub Gists (up to unlimited size)
- **Flexible Storage**: Choose between committing files or using private Gists
- **Run-Once Mode**: Execute single message and exit for scheduled workflows
- **UTC Scheduling**: Default 9 AM UTC daily schedule (customizable)
- **Manual Triggers**: Run workflow on-demand from GitHub UI

## [1.0.0] - 2024-12-23

### Added
- Initial release of Telegram Daily Bot
- Automatic daily reflection posting from markdown files to Telegram group topics
- Markdown file parser with date-based headers (MM-DD format)
- Russian date headers (e.g., "23 Декабря") automatically prepended to messages
- Markdown to HTML conversion for proper formatting in Telegram
- Support for italic (*text*) and bold (**text**) formatting
- Scheduled message posting via cron expressions
- Fallback to custom message if reflection not found for today
- Cross-platform builds (Linux x86_64, x86 32-bit, ARM64, macOS, Windows)
- Comprehensive test coverage for date parsing and markdown conversion
- Makefile with build automation and deployment helpers
- Systemd service configuration via Makefile
- Environment-based configuration via .env file
- Minimal resource usage suitable for tiny VPS deployments

### Features
- **Date Extraction**: Automatically finds and sends today's reflection
- **Russian Localization**: Date headers in Russian (genitive case)
- **HTML Formatting**: Converts markdown to HTML for reliable Telegram rendering
- **Cron Scheduling**: Flexible scheduling with standard cron expressions
- **Multi-Platform**: Builds for Linux (x86_64, x86, ARM), macOS, and Windows
- **Easy Deployment**: One-command systemd service setup
- **Graceful Fallback**: Uses custom message if daily reflection not found

[1.2.0]: https://github.com/stargriv/tg-daily-bot/releases/tag/v1.2.0
[1.1.0]: https://github.com/stargriv/tg-daily-bot/releases/tag/v1.1.0
[1.0.0]: https://github.com/stargriv/tg-daily-bot/releases/tag/v1.0.0
