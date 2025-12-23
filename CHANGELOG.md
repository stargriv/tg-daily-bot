# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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

[1.0.0]: https://github.com/stargriv/tg-daily-bot/releases/tag/v1.0.0
