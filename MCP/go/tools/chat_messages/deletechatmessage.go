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

func DeletechatmessageHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		messageIdVal, ok := args["messageId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: messageId"), nil
		}
		messageId, ok := messageIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: messageId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["to_contact"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("to_contact=%v", val))
		}
		if val, ok := args["to_channel"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("to_channel=%v", val))
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
		url := fmt.Sprintf("%s/chat/users/%s/messages/%s%s", cfg.BaseURL, messageId, queryString)
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

func CreateDeletechatmessageTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_chat_users_userId_messages_messageId",
		mcp.WithDescription("Delete a message"),
		mcp.WithString("messageId", mcp.Required(), mcp.Description("Message ID")),
		mcp.WithString("to_contact", mcp.Description("The userId or email address of a chat contact to whom you previously sent the message.\n\nNote: You must provide either `to_contact` or `to_channel` as a query parameter to delete a message that was previously sent to either an individual or a chat channel respectively. ")),
		mcp.WithString("to_channel", mcp.Description("The channel Id of the channel where you would like to send the message.\n\nYou must provide either `to_contact` or `to_channel` as a query parameter to delete a message that was previously sent to either an individual or a chat channel ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletechatmessageHandler(cfg),
	}
}
