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

func DeletesiptrunkHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		accountIdVal, ok := args["accountId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: accountId"), nil
		}
		accountId, ok := accountIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: accountId"), nil
		}
		trunkIdVal, ok := args["trunkId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: trunkId"), nil
		}
		trunkId, ok := trunkIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: trunkId"), nil
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
		url := fmt.Sprintf("%s/accounts/%s/sip_trunk/trunks/%s%s", cfg.BaseURL, accountId, trunkId, queryString)
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

func CreateDeletesiptrunkTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_accounts_accountId_sip_trunk_trunks_trunkId",
		mcp.WithDescription("Delete a SIP trunk"),
		mcp.WithString("accountId", mcp.Required(), mcp.Description("Unique identifier of the sub account.")),
		mcp.WithString("trunkId", mcp.Required(), mcp.Description("Unique identifier of the SIP Trunk that was previously assigned to a sub account. To retrieve the value of this field, use the List SIP Trunks API.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletesiptrunkHandler(cfg),
	}
}
