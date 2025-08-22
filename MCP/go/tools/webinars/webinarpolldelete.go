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

func WebinarpolldeleteHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		pollIdVal, ok := args["pollId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: pollId"), nil
		}
		pollId, ok := pollIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: pollId"), nil
		}
		queryParams := make([]string, 0)
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
		url := fmt.Sprintf("%s/webinars/%s/polls/%s%s", cfg.BaseURL, webinarId, pollId, queryString)
		req, err := http.NewRequest("DELETE", url, nil)
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

func CreateWebinarpolldeleteTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_webinars_webinarId_polls_pollId",
		mcp.WithDescription("Delete a webinar poll"),
		mcp.WithNumber("webinarId", mcp.Required(), mcp.Description("The webinar ID in \"**long**\" format(represented as int64 data type in JSON). ")),
		mcp.WithString("pollId", mcp.Required(), mcp.Description("The poll ID")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    WebinarpolldeleteHandler(cfg),
	}
}
