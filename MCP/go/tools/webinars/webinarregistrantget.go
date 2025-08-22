package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/zoom-api/mcp-server/config"
	"github.com/zoom-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func WebinarregistrantgetHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		webinarIdVal, ok := args["webinarId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: webinarId"), nil
		}
		webinarId, ok := webinarIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: webinarId"), nil
		}
		registrantIdVal, ok := args["registrantId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: registrantId"), nil
		}
		registrantId, ok := registrantIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: registrantId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["occurrence_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("occurrence_id=%v", val))
		}
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("next_page_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("page_number=%s", cfg.APIKey))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/webinars/%s/registrants/%s%s", cfg.BaseURL, webinarId, registrantId, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateWebinarregistrantgetTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_webinars_webinarId_registrants_registrantId",
		mcp.WithDescription("Get a webinar registrant"),
		mcp.WithNumber("webinarId", mcp.Required(), mcp.Description("The webinar ID in \"**long**\" format(represented as int64 data type in JSON). ")),
		mcp.WithString("registrantId", mcp.Required(), mcp.Description("The registrant ID.")),
		mcp.WithString("occurrence_id", mcp.Description("The meeting occurrence ID.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    WebinarregistrantgetHandler(cfg),
	}
}
