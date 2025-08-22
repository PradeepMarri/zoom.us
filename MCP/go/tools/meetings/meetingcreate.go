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

func MeetingcreateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		userIdVal, ok := args["userId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: userId"), nil
		}
		userId, ok := userIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: userId"), nil
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
		url := fmt.Sprintf("%s/users/%s/meetings%s", cfg.BaseURL, userId, queryString)
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
		var result interface{}
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

func CreateMeetingcreateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_users_userId_meetings",
		mcp.WithDescription("Create a meeting"),
		mcp.WithString("userId", mcp.Required(), mcp.Description("The user ID or email address of the user. For user-level apps, pass `me` as the value for userId.")),
		mcp.WithObject("recurrence", mcp.Description("Input parameter: Recurrence object. Use this object only for a meeting with type `8` i.e., a recurring meeting with fixed time. ")),
		mcp.WithString("template_id", mcp.Description("Input parameter: Unique identifier of the **admin meeting template**. To create admin meeting templates, contact the Zoom support team.\n\nUse this field if you would like to [schedule the meeting from a admin meeting template](https://support.zoom.us/hc/en-us/articles/360036559151-Meeting-templates#h_86f06cff-0852-4998-81c5-c83663c176fb). You can retrieve the value of this field by calling the [List meeting templates](https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/listmeetingtemplates) API.")),
		mcp.WithString("agenda", mcp.Description("Input parameter: Meeting description.")),
		mcp.WithString("schedule_for", mcp.Description("Input parameter: If you would like to schedule this meeting for someone else in your account, provide the Zoom user id or email address of the user here.")),
		mcp.WithObject("settings", mcp.Description("Input parameter: Meeting settings.")),
		mcp.WithString("start_time", mcp.Description("Input parameter: Meeting start time. We support two formats for `start_time` - local time and GMT.<br> \n\nTo set time as GMT the format should be `yyyy-MM-dd`T`HH:mm:ssZ`. Example: \"2020-03-31T12:02:00Z\"\n\nTo set time using a specific timezone, use `yyyy-MM-dd`T`HH:mm:ss` format and specify the timezone [ID](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones) in the `timezone` field OR leave it blank and the timezone set on your Zoom account will be used. You can also set the time as UTC as the timezone field.\n\nThe `start_time` should only be used for scheduled and / or recurring webinars with fixed time.")),
		mcp.WithArray("tracking_fields", mcp.Description("Input parameter: Tracking fields")),
		mcp.WithNumber("duration", mcp.Description("Input parameter: Meeting duration (minutes). Used for scheduled meetings only.")),
		mcp.WithString("password", mcp.Description("Input parameter: Passcode to join the meeting. By default, passcode may only contain the following characters: [a-z A-Z 0-9 @ - _ *] and can have a maximum of 10 characters.\n\n**Note:** If the account owner or the admin has configured [minimum passcode requirement settings](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords#h_a427384b-e383-4f80-864d-794bf0a37604), the passcode value provided here must meet those requirements. <br><br>If the requirements are enabled, you can view those requirements by calling either the [Get User Settings API](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usersettings) or the [Get Account Settings](https://marketplace.zoom.us/docs/api-reference/zoom-api/accounts/accountsettings) API. ")),
		mcp.WithString("topic", mcp.Description("Input parameter: Meeting topic.")),
		mcp.WithString("timezone", mcp.Description("Input parameter: Time zone to format start_time. For example, \"America/Los_Angeles\". For scheduled meetings only. Please reference our [time zone](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones) list for supported time zones and their formats.")),
		mcp.WithNumber("type", mcp.Description("Input parameter: Meeting Type:<br>`1` - Instant meeting.<br>`2` - Scheduled meeting.<br>`3` - Recurring meeting with no fixed time.<br>`8` - Recurring meeting with fixed time.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    MeetingcreateHandler(cfg),
	}
}
