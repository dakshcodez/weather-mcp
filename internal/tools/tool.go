package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/dakshcodez/weather-mcp/internal/services"
	"github.com/dakshcodez/weather-mcp/internal/types"
	"github.com/dakshcodez/weather-mcp/internal/config"
)

func GetForecast(ctx context.Context, req *mcp.CallToolRequest, input types.ForecastInput) (
	*mcp.CallToolResult, any, error,
) {
	// Get points data
	pointsURL := fmt.Sprintf("%s/points/%f,%f", config.NWSAPIBase, input.Latitude, input.Longitude)
	pointsData, err := services.MakeNWSRequest[types.PointsResponse](ctx, pointsURL)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch forecast data for this location."},
			},
		}, nil, nil
	}

	// Get forecast data
	forecastURL := pointsData.Properties.Forecast
	if forecastURL == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch forecast URL."},
			},
		}, nil, nil
	}

	forecastData, err := services.MakeNWSRequest[types.ForecastResponse](ctx, forecastURL)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch detailed forecast."},
			},
		}, nil, nil
	}

	// Format the periods
	periods := forecastData.Properties.Periods
	if len(periods) == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "No forecast periods available."},
			},
		}, nil, nil
	}

	// Show next 5 periods
	var forecasts []string
	for i := range min(5, len(periods)) {
		forecasts = append(forecasts, services.FormatPeriod(periods[i]))
	}

	result := strings.Join(forecasts, "\n---\n")

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: result},
		},
	}, nil, nil
}

func GetAlerts(ctx context.Context, req *mcp.CallToolRequest, input types.AlertsInput) (
	*mcp.CallToolResult, any, error,
) {
	// Build alerts URL
	stateCode := strings.ToUpper(input.State)
	alertsURL := fmt.Sprintf("%s/alerts/active/area/%s", config.NWSAPIBase, stateCode)

	alertsData, err := services.MakeNWSRequest[types.AlertsResponse](ctx, alertsURL)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Unable to fetch alerts or no alerts found."},
			},
		}, nil, nil
	}

	// Check if there are any alerts
	if len(alertsData.Features) == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "No active alerts for this state."},
			},
		}, nil, nil
	}

	// Format alerts
	var alerts []string
	for _, feature := range alertsData.Features {
		alerts = append(alerts, services.FormatAlert(feature))
	}

	result := strings.Join(alerts, "\n---\n")

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: result},
		},
	}, nil, nil
}