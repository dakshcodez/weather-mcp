package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/dakshcodez/weather-mcp/internal/tools"
)

func main() {
	// Create MCP server
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "weather",
		Version: "1.0.0",
	}, nil)

	// Add get_forecast tool
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_forecast",
		Description: "Get weather forecast for a location",
	}, tools.GetForecast)

	// Add get_alerts tool
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_alerts",
		Description: "Get weather alerts for a US state",
	}, tools.GetAlerts)

	// Run server on stdio transport
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}