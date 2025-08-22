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

func ListaccountphonenumbersHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["next_page_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("next_page_token=%v", val))
		}
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		if val, ok := args["extension_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("extension_type=%v", val))
		}
		if val, ok := args["page_size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page_size=%v", val))
		}
		if val, ok := args["number_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("number_type=%v", val))
		}
		if val, ok := args["pending_numbers"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pending_numbers=%v", val))
		}
		if val, ok := args["site_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("site_id=%v", val))
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
		url := fmt.Sprintf("%s/phone/numbers%s", cfg.BaseURL, queryString)
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

func CreateListaccountphonenumbersTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_phone_numbers",
		mcp.WithDescription("List phone numbers"),
		mcp.WithString("next_page_token", mcp.Description("The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.")),
		mcp.WithString("type", mcp.Description("Query response by number assignment. The value can be one of the following:\n<br>\n`assigned`: The number has been assigned to either a user, a call queue, an auto-receptionist or a common area phone in an account. <br>`unassigned`: The number is not assigned to anyone.<br>\n`all`: Include both assigned and unassigned numbers in the response.<br>\n`byoc`: Include Bring Your Own Carrier (BYOC) numbers only in the response.")),
		mcp.WithString("extension_type", mcp.Description("The type of assignee to whom the number is assigned. The value can be one of the following:<br>\n`user`<br> `callQueue`<br> `autoReceptionist`<br>\n`commonAreaPhone`")),
		mcp.WithNumber("page_size", mcp.Description("The number of records returned within a single API call.")),
		mcp.WithString("number_type", mcp.Description("The type of phone number. The value can be either `toll` or `tollfree`.")),
		mcp.WithBoolean("pending_numbers", mcp.Description("Include or exclude pending numbers in the response. The value can be either `true` or `false`.")),
		mcp.WithString("site_id", mcp.Description("Unique identifier of the site. Use this query parameter if you have enabled multiple sites and would like to filter the response of this API call by a specific phone site. See [Managing multiple sites](https://support.zoom.us/hc/en-us/articles/360020809672-Managing-multiple-sites) or [Adding a site](https://support.zoom.us/hc/en-us/articles/360020809672-Managing-multiple-sites#h_05c88e35-1593-491f-b1a8-b7139a75dc15) for details.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListaccountphonenumbersHandler(cfg),
	}
}
