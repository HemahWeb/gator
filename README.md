# Gator 🐊

A powerful RSS feed aggregator built in Go that helps you collect, organize, and browse RSS feeds from your favorite sources.

## Features

- 📰 **RSS Feed Aggregation**: Automatically fetch and store RSS feeds
- 👤 **User Management**: User registration, login, and authentication
- 🔗 **Feed Management**: Add, follow, and unfollow RSS feeds
- 📚 **Smart Browsing**: Browse posts with customizable limits
- 🗄️ **PostgreSQL Storage**: Robust data persistence with SQL
- ⚡ **CLI Interface**: Fast command-line operations
- 🔄 **Periodic Updates**: Configurable feed refresh intervals

## Prerequisites

Before running Gator, make sure you have the following installed:

- **Go 1.24.6 or later** - [Download Go](https://golang.org/dl/)
- **PostgreSQL 12 or later** - [Download PostgreSQL](https://www.postgresql.org/download/)

## Installation

### 1. Install the Gator CLI

```bash
go install github.com/HemahWeb/gator@latest
```

### 2. Verify Installation

```bash
gator --help
```

## Configuration

### 1. Create Configuration File

Create a `.env` file in your project directory:

```bash
# Database Configuration
DB_URL=postgres://username:password@localhost:5432/gator_db?sslmode=disable
```

### 2. Set Up Database

```bash
# Connect to PostgreSQL
psql -U postgres

# Create database
CREATE DATABASE gator_db;

# Create user (optional)
CREATE USER gator_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE gator_db TO gator_user;
```

## Usage

### Getting Started

1. **Register a new user:**
   ```bash
   gator register your_username
   ```

2. **Login:**
   ```bash
   gator login your_username
   ```

3. **Add your first RSS feed:**
   ```bash
   gator addfeed "Tech News" "https://example.com/feed.xml"
   ```

4. **Start aggregating feeds:**
   ```bash
   gator agg 5m
   ```

5. **Browse your posts:**
   ```bash
   gator browse 10
   ```

### Available Commands

#### User Management
- `gator register <username>` - Create a new user account
- `gator login <username>` - Authenticate as a user
- `gator reset` - Reset user data
- `gator users` - List all users

#### Feed Management
- `gator addfeed <name> <url>` - Add a new RSS feed to follow
- `gator feeds` - List all available feeds
- `gator follow <feed_id>` - Start following a specific feed
- `gator unfollow <feed_id>` - Stop following a feed
- `gator following` - Show feeds you're currently following

#### Content Aggregation
- `gator agg <interval>` - Start aggregating feeds (e.g., `5m`, `1h`, `24h`)
- `gator browse [limit]` - Browse posts (default: 2, max: your choice)

### Examples

```bash
# Add a tech blog feed
gator addfeed "TechCrunch" "https://techcrunch.com/feed/"

# Add a news feed
gator addfeed "BBC News" "http://feeds.bbci.co.uk/news/rss.xml"

# Start collecting feeds every 10 minutes
gator agg 10m

# Browse the latest 20 posts
gator browse 20

# Check what feeds you're following
gator following
```

## Development

### Project Structure

```
gator/
├── cmd/           # Command implementations
├── internal/      # Internal packages
│   ├── config/    # Configuration management
│   ├── database/  # Database operations
│   ├── rss/       # RSS feed handling
│   └── user/      # User management
├── sql/           # SQL migrations and queries
├── main.go        # Main application entry point
└── go.mod         # Go module dependencies
```

### Running from Source

```bash
# Clone the repository
git clone https://github.com/HemahWeb/gator.git
cd gator

# Install dependencies
go mod download

# Run the application
go run main.go <command> [args...]
```

### Building

```bash
# Build for your current platform
go build -o gator

# Build for specific platforms
GOOS=linux GOARCH=amd64 go build -o gator-linux
GOOS=darwin GOARCH=amd64 go build -o gator-macos
GOOS=windows GOARCH=amd64 go build -o gator-windows.exe
```

## Database Schema

The application uses PostgreSQL with the following main tables:
- `users` - User accounts and authentication
- `feeds` - RSS feed sources
- `posts` - Individual RSS feed items
- `user_feeds` - User-feed relationships

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

**Happy RSS aggregating! 🚀**
