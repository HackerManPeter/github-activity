# GitHub Activity CLI

A simple command-line tool that fetches and displays the recent activity of any GitHub user in your terminal.

## Features

- Fetch and display a GitHub user's recent activity
- Simple, clean command-line interface
- Minimal dependencies
- Fast performance

## Installation

### Option 1: Using Go

If you have Go installed, you can install directly from the source:

```bash
go install github.com/hackermanpeter/github-activity@latest
```

### Option 2: Binary Download

Download the pre-compiled binary for your platform:

1. Go to the [releases page](https://github.com/yourusername/github-activity/releases)
2. Download the binary for your operating system (Windows, macOS, Linux)
3. Extract and add to your PATH

### Option 3: Build from Source

```bash
# Clone the repository
git clone https://github.com/yourusername/github-activity.git

# Navigate to the project directory
cd github-activity

# Build the application
go build -o github-activity

# Move to a directory in your PATH (optional)
mv github-activity /usr/local/bin/
```

## Usage

```bash
# Basic usage
github-activity <username>

# Example
github-activity kamranahmedse
```

### Example Output

```
Recent GitHub Activity for kamranahmedse:
- Pushed 3 commits to kamranahmedse/developer-roadmap
- Opened a new issue in kamranahmedse/developer-roadmap
- Starred octocat/Hello-World
- Created a new repository kamranahmedse/sample-project
- Commented on issue #42 in kamranahmedse/developer-roadmap
```

### Environment Variables

You can configure the application using the following environment variables:

- `GITHUB_TOKEN`: Your GitHub personal access token (optional, increases API rate limits)
- `ACTIVITY_LIMIT`: Number of activities to display (default: 10)

Example:

```bash
export GITHUB_TOKEN=your_personal_access_token
export ACTIVITY_LIMIT=20
github-activity kamranahmedse
```

## Upcoming Features

- 🚧 Activity filtering by type (commits, issues, PRs, etc.)
- 🚧 Time-based filtering (last day, week, month)
- 🚧 Interactive mode with paging and detailed views
- 🚧 Custom output formatting (JSON, CSV)
- 🚧 Watch mode to monitor activity in real-time
- 🚧 Support for organization activity
- 🚧 Activity statistics and summaries

## Technical Details

This tool uses the GitHub REST API v3 to fetch user events. By default, it works without authentication but has limited API rate. For increased limits, you can provide a GitHub personal access token through the `GITHUB_TOKEN` environment variable.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
