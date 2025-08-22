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

func GetchatmessagesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["to_contact"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("to_contact=%v", val))
		}
		if val, ok := args["to_channel"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("to_channel=%v", val))
		}
		if val, ok := args["date"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("date=%v", val))
		}
		if val, ok := args["page_size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page_size=%v", val))
		}
		if val, ok := args["next_page_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("next_page_token=%v", val))
		}
		if val, ok := args["include_deleted_and_edited_message"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_deleted_and_edited_message=%v", val))
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
		url := fmt.Sprintf("%s/chat/users/%s/messages%s", cfg.BaseURL, queryString)
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

func CreateGetchatmessagesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_chat_users_userId_messages",
		mcp.WithDescription("List user's chat messages"),
		mcp.WithString("to_contact", mcp.Description("The email address of a chat contact with whom the current user chatted. Messages that were sent and/or received between the user and the contact is displayed.\n\nNote: You must provide either `contact` or `channel` as a query parameter to retrieve messages either from an individual or a chat channel. ")),
		mcp.WithString("to_channel", mcp.Description("The channel Id of a channel inside which the current user had chat conversations. Messages that were sent and/or received between the user and the channel is displayed.\n\nNote: You must provide either `contact` or `channel` as a query parameter to retrieve messages either from an individual or a chat channel. ")),
		mcp.WithString("date", mcp.Description("The query date for which you would like to get the chat messages.")),
		mcp.WithNumber("page_size", mcp.Description("The number of records returned with a single API call. ")),
		mcp.WithString("next_page_token", mcp.Description("The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.")),
		mcp.WithString("include_deleted_and_edited_message", mcp.Description("**Optional** <br>\nSet the value of this field to `true` to include edited and deleted messages in the response.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetchatmessagesHandler(cfg),
	}
}
