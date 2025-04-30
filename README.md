# QueBot

A Discord bot built with Go using the Goscord library.

## Features

- Slash command support
- Event handling system
- Voice state tracking
- Message content monitoring
- Guild member management
- Basic ping command
- Booru image search command (supports Safebooru, Danbooru, and Gelbooru)

## Prerequisites

- Go 1.23.4 or higher
- Discord Bot Token
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/NeroQue/QueBot.git
cd QueBot
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file in the root directory and add your Discord bot token:
```
BOT_TOKEN=your_discord_bot_token_here
```

## Project Structure

- `main.go` - Main entry point of the application
- `command/` - Command handler and management
- `event/` - Event handlers for Discord events

## Running the Bot

To start the bot, simply run:
```bash
go run main.go
```

## Features in Detail

### Command System
The bot uses a command manager to handle slash commands and interactions.

#### Available Commands
- `/ping` - Get the bot's latency
- `/booru` - Search for images on various booru sites
  - Provider options: Safebooru, Danbooru, Gelbooru
  - Tag-based search functionality

### Event Handling
- Ready event handling
- Interaction creation handling
- Guild member add event handling
- Message creation event handling (includes a basic ping command)

### Intents
The bot uses the following Discord intents:
- Guilds
- Guild Messages
- Guild Members
- Guild Voice States
- Message Content

## Dependencies

- [Goscord](https://github.com/Goscord/goscord) - Discord API wrapper for Go
- [godotenv](https://github.com/joho/godotenv) - Environment variable management

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
