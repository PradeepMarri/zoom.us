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

func ListcalllogsmetricsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("from=%v", val))
		}
		if val, ok := args["to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("to=%v", val))
		}
		if val, ok := args["site_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("site_id=%v", val))
		}
		if val, ok := args["quality_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("quality_type=%v", val))
		}
		if val, ok := args["page_size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page_size=%v", val))
		}
		if val, ok := args["next_page_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("next_page_token=%v", val))
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
		url := fmt.Sprintf("%s/phone/metrics/call_logs%s", cfg.BaseURL, queryString)
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

func CreateListcalllogsmetricsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_phone_metrics_call_logs",
		mcp.WithDescription("List call logs"),
		mcp.WithString("from", mcp.Description("Start date for the report in `yyyy-mm-dd` format. Specify a 30 day range using the `from` and `to` parameters as the response provides a maximum of a month worth of data per API request.")),
		mcp.WithString("to", mcp.Description("End date for the report in `yyyy-mm-dd` format.")),
		mcp.WithString("site_id", mcp.Description("Unique identifier of the [site](https://support.zoom.us/hc/en-us/articles/360020809672-Managing-multiple-sites). Use this query parameter if you have enabled multiple sites and would like to filter the response of this API call by call logs of a specific phone site.")),
		mcp.WithString("quality_type", mcp.Description("Filter call logs by voice quality. Zoom uses MOS of 3.5 as a general baseline to categorize calls by call quality. A MOS greater than or equal to 3.5 means good quality, while below 3.5 means poor quality. <br><br>The value of this field can be one of the following:<br>\n* `good`: Retrieve call logs of the call(s) with good quality of voice.<br>\n* `bad`: Retrieve call logs of the call(s) with good quality of voice.<br>\n* `all`: Retrieve all call logs without filtering by voice quality. \n\n\n\n")),
		mcp.WithNumber("page_size", mcp.Description("The number of records returned within a single call.")),
		mcp.WithString("next_page_token", mcp.Description("The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListcalllogsmetricsHandler(cfg),
	}
}
