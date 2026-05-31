# Weather MCP Server

A Model Context Protocol (MCP) server written in Go that integrates with the National Weather Service (NWS) API to provide real-time weather forecasts and active alerts.

This server is designed to work seamlessly with MCP-compatible clients (such as Claude Desktop, Cursor, etc.) to enable LLMs to fetch localized weather information dynamically.

## Features

- **Get Weather Forecast (`get_forecast`)**: Fetches detailed weather forecasts (for the next 5 forecast periods, including temperature, wind speed/direction, and conditions) for any US location specified by latitude and longitude.
- **Get Weather Alerts (`get_alerts`)**: Retrieves active weather alerts, warnings, and advisories for a specified US state.

## Installation

### Prerequisites

- [Go](https://go.dev/) 1.25 or higher.

### Build

Clone the repository and build the binary:

```bash
git clone https://github.com/dakshcodez/weather-mcp.git
cd weather-mcp
go build -o weather ./cmd
```

This will create a `weather` executable in the root of the project.

## Configuration

To use this server with an MCP client, add it to your client's configuration file (e.g., `claude_desktop_config.json` on macOS/Windows/Linux).

Replace `/absolute/path/to/weather-mcp/weather` with the actual path to your compiled binary:

```json
{
  "mcpServers": {
    "weather": {
      "command": "/absolute/path/to/weather-mcp/weather",
      "args": []
    }
  }
}
```

Alternatively, you can run the server directly using Go without compiling first:

```json
{
  "mcpServers": {
    "weather": {
      "command": "go",
      "args": [
        "run",
        "/absolute/path/to/weather-mcp/cmd"
      ]
    }
  }
}
```

## Available Tools

### 1. `get_forecast`
Get the weather forecast for a specific location.
- **Arguments:**
  - `latitude` (number, required): Latitude of the location (e.g., `34.0522`).
  - `longitude` (number, required): Longitude of the location (e.g., `-118.2437`).

### 2. `get_alerts`
Get active weather alerts for a US state.
- **Arguments:**
  - `state` (string, required): Two-letter US state code (e.g., `CA`, `NY`, `TX`).

## Architecture & Codebase Structure

- `cmd/main.go`: The main entrypoint. Initializes the MCP server, registers the weather tools, and runs the server using standard I/O (stdio) transport.
- `internal/tools/`: Contains the handler functions (`GetForecast` and `GetAlerts`) that interface with the NWS API.
- `internal/services/`: Services for making HTTP requests to the NWS API and formatting the responses.
- `internal/types/`: Go struct definitions for tool inputs/outputs and API responses.
- `internal/config/`: Configuration values such as base URLs and User-Agents.

## License

MIT
