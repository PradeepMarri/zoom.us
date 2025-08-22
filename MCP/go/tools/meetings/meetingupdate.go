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

func MeetingupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		meetingIdVal, ok := args["meetingId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: meetingId"), nil
		}
		meetingId, ok := meetingIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: meetingId"), nil
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
		var requestBody interface{}
		
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
		url := fmt.Sprintf("%s/meetings/%s%s", cfg.BaseURL, meetingId, queryString)
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

func CreateMeetingupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_meetings_meetingId",
		mcp.WithDescription("Update a meeting"),
		mcp.WithNumber("meetingId", mcp.Required(), mcp.Description("The meeting ID in **long** format. The data type of this field is \"long\"(represented as int64 in JSON).\n\nWhile storing it in your database, store it as a **long** data type and **not as an integer**, as the Meeting IDs can be longer than 10 digits.")),
		mcp.WithString("occurrence_id", mcp.Description("Meeting occurrence id. Support change of agenda, start_time, duration, settings: {host_video, participant_video, join_before_host, mute_upon_entry, waiting_room, watermark, auto_recording}")),
		mcp.WithString("schedule_for", mcp.Description("Input parameter: Email or userId if you want to schedule meeting for another user.")),
		mcp.WithString("agenda", mcp.Description("Input parameter: Meeting description.")),
		mcp.WithNumber("duration", mcp.Description("Input parameter: Meeting duration (minutes). Used for scheduled meetings only.")),
		mcp.WithNumber("type", mcp.Description("Input parameter: Meeting Types:<br>`1` - Instant meeting.<br>`2` - Scheduled meeting.<br>`3` - Recurring meeting with no fixed time.<br>`8` - Recurring meeting with a fixed time.")),
		mcp.WithObject("recurrence", mcp.Description("Input parameter: Recurrence object. Use this object only for a meeting with type `8` i.e., a recurring meeting with fixed time. ")),
		mcp.WithString("password", mcp.Description("Input parameter: Meeting passcode. Passcode may only contain the following characters: [a-z A-Z 0-9 @ - _ *] and can have a maximum of 10 characters.\n\n**Note:** If the account owner or the admin has configured [minimum passcode requirement settings](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords#h_a427384b-e383-4f80-864d-794bf0a37604), the passcode value provided here must meet those requirements. <br><br>If the requirements are enabled, you can view those requirements by calling either the [Get User Settings API](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usersettings) or the  [Get Account Settings](https://marketplace.zoom.us/docs/api-reference/zoom-api/accounts/accountsettings) API. ")),
		mcp.WithString("settings", mcp.Description("")),
		mcp.WithString("start_time", mcp.Description("Input parameter: Meeting start time. When using a format like \"yyyy-MM-dd'T'HH:mm:ss'Z'\", always use GMT time. When using a format like \"yyyy-MM-dd'T'HH:mm:ss\", you should use local time and  specify the time zone. Only used for scheduled meetings and recurring meetings with a fixed time.")),
		mcp.WithString("template_id", mcp.Description("Input parameter: Unique identifier of the meeting template. \n\nUse this field if you would like to [schedule the meeting from a meeting template](https://support.zoom.us/hc/en-us/articles/360036559151-Meeting-templates#h_86f06cff-0852-4998-81c5-c83663c176fb). You can retrieve the value of this field by calling the [List meeting templates]() API.")),
		mcp.WithString("timezone", mcp.Description("Input parameter: Time zone to format start_time. For example, \"America/Los_Angeles\". For scheduled meetings only. Please reference our [time zone](#timezones) list for supported time zones and their formats.")),
		mcp.WithString("topic", mcp.Description("Input parameter: Meeting topic.")),
		mcp.WithArray("tracking_fields", mcp.Description("Input parameter: Tracking fields")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    MeetingupdateHandler(cfg),
	}
}
