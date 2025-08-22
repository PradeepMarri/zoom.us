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

func WebinarupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/webinars/%s%s", cfg.BaseURL, webinarId, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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

func CreateWebinarupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_webinars_webinarId",
		mcp.WithDescription("Update a webinar"),
		mcp.WithNumber("webinarId", mcp.Required(), mcp.Description("The webinar ID in \"**long**\" format(represented as int64 data type in JSON). ")),
		mcp.WithString("occurrence_id", mcp.Description("Webinar occurrence id. Support change of agenda, start_time, duration, settings: {host_video, panelist_video, hd_video, watermark, auto_recording}")),
		mcp.WithString("start_time", mcp.Description("Input parameter: Webinar start time, in the format \"yyyy-MM-dd'T'HH:mm:ss'Z'.\" Should be in GMT time. In the format \"yyyy-MM-dd'T'HH:mm:ss.\" This should be in local time and the timezone should be specified. Only used for scheduled webinars and recurring webinars with a fixed time.")),
		mcp.WithString("topic", mcp.Description("Input parameter: Webinar topic.")),
		mcp.WithNumber("type", mcp.Description("Input parameter: Webinar Types:<br>`5` - webinar.<br>`6` - Recurring webinar with no fixed time.<br>`9` - Recurring webinar with a fixed time.")),
		mcp.WithNumber("duration", mcp.Description("Input parameter: Webinar duration (minutes). Used for scheduled webinar only.")),
		mcp.WithString("agenda", mcp.Description("Input parameter: Webinar description.")),
		mcp.WithObject("recurrence", mcp.Description("Input parameter: Recurrence object. Use this object only for a meeting with type `8` i.e., a recurring meeting with fixed time. ")),
		mcp.WithString("settings", mcp.Description("")),
		mcp.WithString("timezone", mcp.Description("Input parameter: Time zone to format start_time. For example, \"America/Los_Angeles\". For scheduled meetings only. Please reference our [time zone](#timezones) list for supported time zones and their formats.")),
		mcp.WithArray("tracking_fields", mcp.Description("Input parameter: Tracking fields")),
		mcp.WithString("password", mcp.Description("Input parameter: [Webinar passcode](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords). By default, passcode may only contain the following characters: [a-z A-Z 0-9 @ - _ * !] and can have a maximum of 10 characters.\n\n**Note:** If the account owner or the admin has configured [minimum passcode requirement settings](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords#h_a427384b-e383-4f80-864d-794bf0a37604), the passcode value provided here must meet those requirements. <br><br>If the requirements are enabled, you can view those requirements by calling either the [Get User Settings API](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usersettings) or the  [Get Account Settings](https://marketplace.zoom.us/docs/api-reference/zoom-api/accounts/accountsettings) API. \n\nIf \"**Require a passcode when scheduling new meetings**\" setting has been **enabled** **and** [locked](https://support.zoom.us/hc/en-us/articles/115005269866-Using-Tiered-Settings#locked) for the user, the passcode field will be autogenerated for the Webinar in the response even if it is not provided in the API request. <br><br>\n\n\n\n\n\n\n\n")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    WebinarupdateHandler(cfg),
	}
}
