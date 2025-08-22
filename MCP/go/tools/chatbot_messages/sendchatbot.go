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

func SendchatbotHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/im/chat/messages%s", cfg.BaseURL, queryString)
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

func CreateSendchatbotTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_im_chat_messages",
		mcp.WithDescription("Send chatbot messages"),
		mcp.WithObject("content", mcp.Required(), mcp.Description("Input parameter: JSON template describing how the message should be displayed for the user. For more information please see our [\"Send Message\" templates](https://marketplace.zoom.us/docs/guides/chatbots/sending-messages#example-request).")),
		mcp.WithBoolean("is_markdown_support", mcp.Description("Input parameter: **Optional**<br>\nApplies the markdown parser to your chatbot message if the value of this field is set to `true`.<br> To learn more, refer to the Chatbot message [markdown reference](https://marketplace.zoom.us/docs/guides/chatbots/customizing-messages/message-with-markdown).")),
		mcp.WithString("robot_jid", mcp.Required(), mcp.Description("Input parameter: Robot JID created when enabling chatbot features on your marketplace app.")),
		mcp.WithString("to_jid", mcp.Required(), mcp.Description("Input parameter: Unique JID of reciever. Can be a group or user.")),
		mcp.WithString("user_jid", mcp.Description("Input parameter: **Optional**<br>\nThe UserJID of the user on whose behalf the message is being sent. Use this field to prevent members of a channel from getting notifications that were set up by a user who has left the channel.")),
		mcp.WithString("visible_to_user", mcp.Description("Input parameter: **Optional**<br>Allow a Chatbot to send a message to a group channel, but have only one designated person in that group channel see the message by providing the person's UserID in this field.")),
		mcp.WithString("account_id", mcp.Required(), mcp.Description("Input parameter: Account ID of the authorized account.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    SendchatbotHandler(cfg),
	}
}
