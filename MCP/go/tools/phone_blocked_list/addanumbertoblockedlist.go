package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/zoom-api/mcp-server/config"
	"github.com/zoom-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func AddanumbertoblockedlistHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
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
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/phone/blocked_list%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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

func CreateAddanumbertoblockedlistTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_phone_blocked_list",
		mcp.WithDescription("Create a blocked list"),
		mcp.WithString("comment", mcp.Description("Input parameter: Provide a comment to help you identify the blocked number or prefix.")),
		mcp.WithString("match_type", mcp.Description("Input parameter: Specify the match type for the blocked list. The values can be one of the following:<br><br>\n`phoneNumber`: Choose this option (Phone Number Match) if you want to block a specific phone number. Then, in the `phone_number` field, provide the phone number along with the country code.<br><br>\n`prefix`: Choose this option (Prefix Match) if you want to block all numbers with a specific country code and area code. Next, in the `phone_number` field, enter a country code as part of the prefix. For example, entering 1907 blocks numbers with country code 1 and area code 907.")),
		mcp.WithString("phone_number", mcp.Description("Input parameter: The phone number to be blocked if you passed \"phoneNumber\" as the value for the `match_type` field. If you passed \"prefix\" as the value for the `match_type` field, provide the prefix of the phone number here including the country code. For example, entering 1905 blocks numbers with country code 1 and area code 905. ")),
		mcp.WithString("status", mcp.Description("Input parameter: Enable or disable the blocking. One of the following values are allowed:<br>\n`active`: Keep the blocking active.<br>\n`inactive`: Disable the blocking.")),
		mcp.WithString("block_type", mcp.Description("Input parameter: State whether you want the block type to be inbound or outbound.<br>\n`inbound`: Pass this value to prevent the blocked number or prefix from calling in to phone users.<br>\n`outbound`: Pass this value to prevent phone users from calling the blocked number or prefix.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AddanumbertoblockedlistHandler(cfg),
	}
}
