package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// AccountOptions represents the AccountOptions schema from the OpenAPI specification
type AccountOptions struct {
	Pay_mode string `json:"pay_mode,omitempty"` // Payee:<br>`master` - master account holder pays.<br>`sub` - Sub account holder pays.
	Room_connector_list []string `json:"room_connector_list,omitempty"` // Specify the IP addresses of the Room Connectors that you would like to share with the sub account. Multiple values can be separated by comma. If no value is provided in this field, all the Room Connectors of a master account will be shared with the sub account. **Note:** This option can only be used if the value of `share_rc` is set to `true`.
	Share_mc bool `json:"share_mc,omitempty"` // Enable/disable the option for a sub account to use shared [Meeting Connector(s)](https://support.zoom.us/hc/en-us/articles/201363093-Getting-Started-with-the-Meeting-Connector) that are set up by the master account. Meeting Connectors can only be used by On-prem users.
	Share_rc bool `json:"share_rc,omitempty"` // Enable/disable the option for a sub account to use shared [Virtual Room Connector(s)](https://support.zoom.us/hc/en-us/articles/202134758-Getting-Started-With-Virtual-Room-Connector) that are set up by the master account. Virtual Room Connectors can only be used by On-prem users.
	Billing_auto_renew bool `json:"billing_auto_renew,omitempty"` // Toggle whether automatic billing renewal is on or off.
	Meeting_connector_list []string `json:"meeting_connector_list,omitempty"` // Specify the IP addresses of the Meeting Connectors that you would like to share with the sub account. Multiple values can be separated by comma. If no value is provided in this field, all the Meeting Connectors of a master account will be shared with the sub account. **Note:** This option can only be used if the value of `share_mc` is set to `true`.
}

// MeetingCreate represents the MeetingCreate schema from the OpenAPI specification
type MeetingCreate struct {
	Schedule_for string `json:"schedule_for,omitempty"` // If you would like to schedule this meeting for someone else in your account, provide the Zoom user id or email address of the user here.
	Settings map[string]interface{} `json:"settings,omitempty"` // Meeting settings.
	Tracking_fields []map[string]interface{} `json:"tracking_fields,omitempty"` // Tracking fields
	TypeField int `json:"type,omitempty"` // Meeting Type:<br>`1` - Instant meeting.<br>`2` - Scheduled meeting.<br>`3` - Recurring meeting with no fixed time.<br>`8` - Recurring meeting with fixed time.
	Duration int `json:"duration,omitempty"` // Meeting duration (minutes). Used for scheduled meetings only.
	Recurrence map[string]interface{} `json:"recurrence,omitempty"` // Recurrence object. Use this object only for a meeting with type `8` i.e., a recurring meeting with fixed time.
	Template_id string `json:"template_id,omitempty"` // Unique identifier of the **admin meeting template**. To create admin meeting templates, contact the Zoom support team. Use this field if you would like to [schedule the meeting from a admin meeting template](https://support.zoom.us/hc/en-us/articles/360036559151-Meeting-templates#h_86f06cff-0852-4998-81c5-c83663c176fb). You can retrieve the value of this field by calling the [List meeting templates](https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/listmeetingtemplates) API.
	Topic string `json:"topic,omitempty"` // Meeting topic.
	Timezone string `json:"timezone,omitempty"` // Time zone to format start_time. For example, "America/Los_Angeles". For scheduled meetings only. Please reference our [time zone](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones) list for supported time zones and their formats.
	Agenda string `json:"agenda,omitempty"` // Meeting description.
	Password string `json:"password,omitempty"` // Passcode to join the meeting. By default, passcode may only contain the following characters: [a-z A-Z 0-9 @ - _ *] and can have a maximum of 10 characters. **Note:** If the account owner or the admin has configured [minimum passcode requirement settings](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords#h_a427384b-e383-4f80-864d-794bf0a37604), the passcode value provided here must meet those requirements. <br><br>If the requirements are enabled, you can view those requirements by calling either the [Get User Settings API](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usersettings) or the [Get Account Settings](https://marketplace.zoom.us/docs/api-reference/zoom-api/accounts/accountsettings) API.
	Start_time string `json:"start_time,omitempty"` // Meeting start time. We support two formats for `start_time` - local time and GMT.<br> To set time as GMT the format should be `yyyy-MM-dd`T`HH:mm:ssZ`. Example: "2020-03-31T12:02:00Z" To set time using a specific timezone, use `yyyy-MM-dd`T`HH:mm:ss` format and specify the timezone [ID](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones) in the `timezone` field OR leave it blank and the timezone set on your Zoom account will be used. You can also set the time as UTC as the timezone field. The `start_time` should only be used for scheduled and / or recurring webinars with fixed time.
}

// IMGroup represents the IMGroup schema from the OpenAPI specification
type IMGroup struct {
	Total_members int `json:"total_members,omitempty"` // Total number of members in this group.
	Name string `json:"name,omitempty"` // Group name.
	Search_by_account bool `json:"search_by_account,omitempty"` // Members can search for others under same account.
	Search_by_domain bool `json:"search_by_domain,omitempty"` // Members can search for others in the same email domain.
	Search_by_ma_account bool `json:"search_by_ma_account,omitempty"` // Members can search for others under same master account - including all sub accounts.
	TypeField string `json:"type,omitempty"` // IM Group types:<br>`normal` - Only members can see the other members in the group. Other people can search for members in the group.<br>`shared` - Everyone in the account can see the group and members. <br>`restricted` - No one except group members can see the group or search for other group members.
}

// BillingContact represents the BillingContact schema from the OpenAPI specification
type BillingContact struct {
	Email string `json:"email,omitempty"` // Billing Contact's email address.
	First_name string `json:"first_name,omitempty"` // Billing Contact's first name.
	Phone_number string `json:"phone_number,omitempty"` // Billing Contact's phone number.
	State string `json:"state,omitempty"` // Billing Contact's state.
	Zip string `json:"zip,omitempty"` // Billing Contact's zip/postal code.
	Country string `json:"country,omitempty"` // Billing Contact's country.
	Apt string `json:"apt,omitempty"` // Billing Contact's apartment/suite.
	Address string `json:"address,omitempty"` // Billing Contact's address.
	Last_name string `json:"last_name,omitempty"` // Billing Contact's last name.
	City string `json:"city,omitempty"` // Billing Contact's city.
}

// QOSParticipant represents the QOSParticipant schema from the OpenAPI specification
type QOSParticipant struct {
	Version string `json:"version,omitempty"` // Participant's Zoom Client version.
	Harddisk_id string `json:"harddisk_id,omitempty"` // Participant's hard disk ID.
	Leave_time string `json:"leave_time,omitempty"` // The time at which participant left the meeting.
	Mac_addr string `json:"mac_addr,omitempty"` // Participant's MAC address.
	Pc_name string `json:"pc_name,omitempty"` // Participant's PC name.
	User_id string `json:"user_id,omitempty"` // Participant ID.
	Device string `json:"device,omitempty"` // The type of device using which the participant joined the meeting.
	Domain string `json:"domain,omitempty"` // Participant's PC domain.
	Ip_address string `json:"ip_address,omitempty"` // Participant's IP address.
	Join_time string `json:"join_time,omitempty"` // The time at which participant joined the meeting.
	Location string `json:"location,omitempty"` // Participant's location.
	User_name string `json:"user_name,omitempty"` // Participant display name.
	User_qos []map[string]interface{} `json:"user_qos,omitempty"` // Quality of service provided to the user.
}

// MeetingRegistrantList represents the MeetingRegistrantList schema from the OpenAPI specification
type MeetingRegistrantList struct {
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_number int `json:"page_number,omitempty"` // **Deprecated**: This field has been deprecated. Please use the "next_page_token" field for pagination instead of this field. The page number of the current results.
	Page_size int `json:"page_size,omitempty"` // The number of records returned with a single API call.
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
	Registrants []interface{} `json:"registrants,omitempty"` // List of registrant objects.
}

// Session represents the Session schema from the OpenAPI specification
type Session struct {
	Tracking_fields []interface{} `json:"tracking_fields,omitempty"` // Tracking fields
	TypeField int `json:"type,omitempty"` // Meeting Type:<br>`1` - Instant meeting.<br>`2` - Scheduled meeting.<br>`3` - Recurring meeting with no fixed time.<br>`8` - Recurring meeting with fixed time.
	Duration int `json:"duration,omitempty"` // Meeting duration (minutes). Used for scheduled meetings only.
	Timezone string `json:"timezone,omitempty"` // Time zone to format start_time. For example, "America/Los_Angeles". For scheduled meetings only. Please reference our [time zone](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones) list for supported time zones and their formats.
	Agenda string `json:"agenda,omitempty"` // Meeting description.
	Password string `json:"password,omitempty"` // Password to join the meeting. Password may only contain the following characters: [a-z A-Z 0-9 @ - _ *]. Max of 10 characters.
	Topic string `json:"topic,omitempty"` // Meeting topic.
	Recurrence map[string]interface{} `json:"recurrence,omitempty"` // Recurrence object. Use this object only for a meeting with type `8` i.e., a recurring meeting with fixed time.
	Settings map[string]interface{} `json:"settings,omitempty"` // Meeting settings.
	Start_time string `json:"start_time,omitempty"` // Meeting start time. When using a format like "yyyy-MM-dd'T'HH:mm:ss'Z'", always use GMT time. When using a format like "yyyy-MM-dd'T'HH:mm:ss", you should use local time and specify the time zone. This is only used for scheduled meetings and recurring meetings with a fixed time.
}

// MeetingLiveStreamStatus represents the MeetingLiveStreamStatus schema from the OpenAPI specification
type MeetingLiveStreamStatus struct {
	Action string `json:"action,omitempty"` // Update the status of a livestream. The value can be one of the following:<br> `start`: Start a live stream. <br> `stop`: Stop an ongoing live stream.
	Settings map[string]interface{} `json:"settings,omitempty"` // Update the settings of a live streaming session. The settings can only be updated for a live stream that has been stopped. You can not update the settings of an ongoing live stream.
}

// WebinarRegistrantQuestions represents the WebinarRegistrantQuestions schema from the OpenAPI specification
type WebinarRegistrantQuestions struct {
	Custom_questions []interface{} `json:"custom_questions,omitempty"` // Array of Registrant Custom Questions.
	Questions []interface{} `json:"questions,omitempty"` // Array of registration fields whose values should be provided by registrants during registration.
}

// Device represents the Device schema from the OpenAPI specification
type Device struct {
	Name string `json:"name"` // Device name.
	Protocol string `json:"protocol"` // Device protocol:<br>`H.323` - H.323.<br>`SIP` - SIP.
	Encryption string `json:"encryption"` // Device encryption:<br>`auto` - auto.<br>`yes` - yes.<br>`no` - no.
	Ip string `json:"ip"` // Device IP.
}

// AccountSettingsZoomRooms represents the AccountSettingsZoomRooms schema from the OpenAPI specification
type AccountSettingsZoomRooms struct {
	Hide_host_information bool `json:"hide_host_information,omitempty"` // Hide host and meeting ID from private meetings.
	List_meetings_with_calendar bool `json:"list_meetings_with_calendar,omitempty"` // Display meeting list with calendar integration.
	Start_airplay_manually bool `json:"start_airplay_manually,omitempty"` // Start AirPlay service manually.
	Force_private_meeting bool `json:"force_private_meeting,omitempty"` // Shift all meetings to private.
	Ultrasonic bool `json:"ultrasonic,omitempty"` // Automatic direct sharing using an ultrasonic proximity signal.
	Zr_post_meeting_feedback bool `json:"zr_post_meeting_feedback,omitempty"` // Zoom Room post meeting feedback.
	Auto_start_stop_scheduled_meetings bool `json:"auto_start_stop_scheduled_meetings,omitempty"` // Automatic start and stop for scheduled meetings.
	Upcoming_meeting_alert bool `json:"upcoming_meeting_alert,omitempty"` // Upcoming meeting alert.
	Cmr_for_instant_meeting bool `json:"cmr_for_instant_meeting,omitempty"` // Cloud recording for instant meetings.
	Weekly_system_restart bool `json:"weekly_system_restart,omitempty"` // Weekly system restart.
}

// UserSettingsUpdate represents the UserSettingsUpdate schema from the OpenAPI specification
type UserSettingsUpdate struct {
	Tsp map[string]interface{} `json:"tsp,omitempty"` // Account Settings: TSP.
	Email_notification map[string]interface{} `json:"email_notification,omitempty"`
	Feature map[string]interface{} `json:"feature,omitempty"`
	In_meeting map[string]interface{} `json:"in_meeting,omitempty"`
	Profile map[string]interface{} `json:"profile,omitempty"`
	Recording map[string]interface{} `json:"recording,omitempty"`
	Schedule_meeting map[string]interface{} `json:"schedule_meeting,omitempty"`
	Telephony map[string]interface{} `json:"telephony,omitempty"`
}

// RecordingMeetingList represents the RecordingMeetingList schema from the OpenAPI specification
type RecordingMeetingList struct {
	From string `json:"from,omitempty"` // Start Date.
	To string `json:"to,omitempty"` // End Date.
	Page_size int `json:"page_size,omitempty"` // The number of records returned within a single API call.
	Total_records int `json:"total_records,omitempty"` // The number of all records available across pages.
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Meetings []interface{} `json:"meetings,omitempty"` // List of recordings.
}

// QOSVideo represents the QOSVideo schema from the OpenAPI specification
type QOSVideo struct {
	Max_loss string `json:"max_loss,omitempty"` // Max loss: the max amount of packet loss, that is the max percentage of packets that fail to arrive at their destination.
	Avg_loss string `json:"avg_loss,omitempty"` // The average amount of packet loss, i.e., the percentage of packets that fail to arrive at their destination.
	Bitrate string `json:"bitrate,omitempty"` // The number of bits per second that can be transmitted along a digital network. The value of this field is expressed in kbps.
	Jitter string `json:"jitter,omitempty"` // The variation in the delay of received packets. The value of this field is expressed in milliseconds.
	Latency string `json:"latency,omitempty"` // The amount of time it takes for a packet to travel from one point to another. The value of this field is expressed in milliseconds.
	Frame_rate string `json:"frame_rate,omitempty"` // The rate at which your video camera can produce unique images, or frames. Zoom supports a frame rate of up to 30fps.
	Resolution string `json:"resolution,omitempty"` // The number of pixels in each dimension that can be displayed by your video camera.
}

// MeetingRegistrantQuestions represents the MeetingRegistrantQuestions schema from the OpenAPI specification
type MeetingRegistrantQuestions struct {
	Custom_questions []map[string]interface{} `json:"custom_questions,omitempty"` // Array of Registrant Custom Questions
	Questions []map[string]interface{} `json:"questions,omitempty"` // Array of Registrant Questions
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Monthly_week int `json:"monthly_week,omitempty"` // Use this field **only if you're scheduling a recurring webinar of type** `3` to state the week of the month when the webinar should recur. If you use this field, **you must also use the `monthly_week_day` field to state the day of the week when the webinar should recur.** <br>`-1` - Last week of the month.<br>`1` - First week of the month.<br>`2` - Second week of the month.<br>`3` - Third week of the month.<br>`4` - Fourth week of the month.
	Monthly_week_day int `json:"monthly_week_day,omitempty"` // Use this field **only if you're scheduling a recurring webinar of type** `3` to state a specific day in a week when the monthly webinar should recur. To use this field, you must also use the `monthly_week` field. <br>`1` - Sunday.<br>`2` - Monday.<br>`3` - Tuesday.<br>`4` - Wednesday.<br>`5` - Thursday.<br>`6` - Friday.<br>`7` - Saturday.
	Repeat_interval int `json:"repeat_interval,omitempty"` // Define the interval at which the webinar should recur. For instance, if you would like to schedule a Webinar that recurs every two months, you must set the value of this field as `2` and the value of the `type` parameter as `3`. For a daily webinar, the maximum interval you can set is `90` days. For a weekly webinar, the maximum interval that you can set is `12` weeks. For a monthly webinar, the maximum interval that you can set is `3` months.
	TypeField int `json:"type"` // Recurrence webinar types:<br>`1` - Daily.<br>`2` - Weekly.<br>`3` - Monthly.
	Weekly_days string `json:"weekly_days,omitempty"` // Use this field **only if you're scheduling a recurring webinar of type** `2` to state which day(s) of the week the webinar should repeat. <br> The value for this field could be a number between `1` to `7` in string format. For instance, if the Webinar should recur on Sunday, provide `"1"` as the value of this field. <br><br> **Note:** If you would like the webinar to occur on multiple days of a week, you should provide comma separated values for this field. For instance, if the Webinar should recur on Sundays and Tuesdays provide `"1,3"` as the value of this field. <br>`1` - Sunday. <br>`2` - Monday.<br>`3` - Tuesday.<br>`4` - Wednesday.<br>`5` - Thursday.<br>`6` - Friday.<br>`7` - Saturday.
	End_date_time string `json:"end_date_time,omitempty"` // Select a date when the webinar will recur before it is canceled. Should be in UTC time, such as 2017-11-25T12:00:00Z. (Cannot be used with "end_times".)
	End_times int `json:"end_times,omitempty"` // Select how many times the webinar will recur before it is canceled. (Cannot be used with "end_date_time".)
	Monthly_day int `json:"monthly_day,omitempty"` // Use this field **only if you're scheduling a recurring webinar of type** `3` to state which day in a month, the webinar should recur. The value range is from 1 to 31. For instance, if you would like the webinar to recur on 23rd of each month, provide `23` as the value of this field and `1` as the value of the `repeat_interval` field. Instead, if you would like the webinar to recur once every three months, on 23rd of the month, change the value of the `repeat_interval` field to `3`.
}

// WebinarUpdate represents the WebinarUpdate schema from the OpenAPI specification
type WebinarUpdate struct {
	Tracking_fields []map[string]interface{} `json:"tracking_fields,omitempty"` // Tracking fields
	Password string `json:"password,omitempty"` // [Webinar passcode](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords). By default, passcode may only contain the following characters: [a-z A-Z 0-9 @ - _ * !] and can have a maximum of 10 characters. **Note:** If the account owner or the admin has configured [minimum passcode requirement settings](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords#h_a427384b-e383-4f80-864d-794bf0a37604), the passcode value provided here must meet those requirements. <br><br>If the requirements are enabled, you can view those requirements by calling either the [Get User Settings API](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usersettings) or the [Get Account Settings](https://marketplace.zoom.us/docs/api-reference/zoom-api/accounts/accountsettings) API. If "**Require a passcode when scheduling new meetings**" setting has been **enabled** **and** [locked](https://support.zoom.us/hc/en-us/articles/115005269866-Using-Tiered-Settings#locked) for the user, the passcode field will be autogenerated for the Webinar in the response even if it is not provided in the API request. <br><br>
	Settings interface{} `json:"settings,omitempty"`
	Topic string `json:"topic,omitempty"` // Webinar topic.
	TypeField int `json:"type,omitempty"` // Webinar Types:<br>`5` - webinar.<br>`6` - Recurring webinar with no fixed time.<br>`9` - Recurring webinar with a fixed time.
	Duration int `json:"duration,omitempty"` // Webinar duration (minutes). Used for scheduled webinar only.
	Recurrence map[string]interface{} `json:"recurrence,omitempty"` // Recurrence object. Use this object only for a meeting with type `8` i.e., a recurring meeting with fixed time.
	Start_time string `json:"start_time,omitempty"` // Webinar start time, in the format "yyyy-MM-dd'T'HH:mm:ss'Z'." Should be in GMT time. In the format "yyyy-MM-dd'T'HH:mm:ss." This should be in local time and the timezone should be specified. Only used for scheduled webinars and recurring webinars with a fixed time.
	Timezone string `json:"timezone,omitempty"` // Time zone to format start_time. For example, "America/Los_Angeles". For scheduled meetings only. Please reference our [time zone](#timezones) list for supported time zones and their formats.
	Agenda string `json:"agenda,omitempty"` // Webinar description.
}

// AccountPlanBaseRequired represents the AccountPlanBaseRequired schema from the OpenAPI specification
type AccountPlanBaseRequired struct {
	Hosts int `json:"hosts"` // Account base plan number of hosts. For a Pro Plan please select a value between 1 and 9. For a Business Plan please select a value between 10 and 49. For a Education Plan please select a value between 20 and 149. For a Free Trial Plan please select a value between 1 and 9999.
	TypeField string `json:"type"` // Account base <a href="https://marketplace.zoom.us/docs/api-reference/other-references/plans">plan type.</a>
}

// QOSAudio represents the QOSAudio schema from the OpenAPI specification
type QOSAudio struct {
	Max_loss string `json:"max_loss,omitempty"` // Max loss: the max amount of packet loss, that is the max percentage of packets that fail to arrive at their destination.
	Avg_loss string `json:"avg_loss,omitempty"` // The average amount of packet loss, i.e., the percentage of packets that fail to arrive at their destination.
	Bitrate string `json:"bitrate,omitempty"` // The number of bits per second that can be transmitted along a digital network. The value of this field is expressed in kbps.
	Jitter string `json:"jitter,omitempty"` // The variation in the delay of received packets. The value of this field is expressed in milliseconds.
	Latency string `json:"latency,omitempty"` // The amount of time it takes for a packet to travel from one point to another. The value of this field is expressed in milliseconds.
}

// GroupList represents the GroupList schema from the OpenAPI specification
type GroupList struct {
	Groups []interface{} `json:"groups,omitempty"` // List of Group objects.
	Total_records int `json:"total_records,omitempty"` // Total records.
}

// IMGroupList represents the IMGroupList schema from the OpenAPI specification
type IMGroupList struct {
	Total_records int `json:"total_records,omitempty"` // Total number of records returned.
	Groups []interface{} `json:"groups,omitempty"` // List of group objects.
}

// MeetingSettings represents the MeetingSettings schema from the OpenAPI specification
type MeetingSettings struct {
	Approved_or_denied_countries_or_regions map[string]interface{} `json:"approved_or_denied_countries_or_regions,omitempty"` // Approve or block users from specific regions/countries from joining this meeting.
	Jbh_time int `json:"jbh_time,omitempty"` // If the value of "join_before_host" field is set to true, this field can be used to indicate time limits within which a participant may join a meeting before a host. The value of this field can be one of the following: * `0`: Allow participant to join anytime. * `5`: Allow participant to join 5 minutes before meeting start time. * `10`: Allow participant to join 10 minutes before meeting start time.
	Auto_recording string `json:"auto_recording,omitempty"` // Automatic recording:<br>`local` - Record on local.<br>`cloud` - Record on cloud.<br>`none` - Disabled.
	Show_share_button bool `json:"show_share_button,omitempty"` // Show social share buttons on the meeting registration page. This setting only works for meetings that require [registration](https://support.zoom.us/hc/en-us/articles/211579443-Setting-up-registration-for-a-meeting).
	Use_pmi bool `json:"use_pmi,omitempty"` // Use a personal meeting ID. Only used for scheduled meetings and recurring meetings with no fixed time.
	Registration_type int `json:"registration_type,omitempty"` // Registration type. Used for recurring meeting with fixed time only. <br>`1` Attendees register once and can attend any of the occurrences.<br>`2` Attendees need to register for each occurrence to attend.<br>`3` Attendees register once and can choose one or more occurrences to attend.
	Watermark bool `json:"watermark,omitempty"` // Add watermark when viewing a shared screen.
	Global_dial_in_countries []string `json:"global_dial_in_countries,omitempty"` // List of global dial-in countries
	Authentication_domains string `json:"authentication_domains,omitempty"` // If user has configured ["Sign Into Zoom with Specified Domains"](https://support.zoom.us/hc/en-us/articles/360037117472-Authentication-Profiles-for-Meetings-and-Webinars#h_5c0df2e1-cfd2-469f-bb4a-c77d7c0cca6f) option, this will list the domains that are authenticated.
	Global_dial_in_numbers []map[string]interface{} `json:"global_dial_in_numbers,omitempty"` // Global Dial-in Countries/Regions
	Enforce_login_domains string `json:"enforce_login_domains,omitempty"` // Only signed in users with specified domains can join meetings. **This field is deprecated and will not be supported in the future.** <br><br>As an alternative, use the "meeting_authentication", "authentication_option" and "authentication_domains" fields to understand the [authentication configurations](https://support.zoom.us/hc/en-us/articles/360037117472-Authentication-Profiles-for-Meetings-and-Webinars) set for the meeting.
	Participant_video bool `json:"participant_video,omitempty"` // Start video when participants join the meeting.
	Waiting_room bool `json:"waiting_room,omitempty"` // Enable waiting room
	Approval_type int `json:"approval_type,omitempty"` // Enable registration and set approval for the registration. Note that this feature requires the host to be of **Licensed** user type. **Registration cannot be enabled for a basic user.** <br><br> `0` - Automatically approve.<br>`1` - Manually approve.<br>`2` - No registration required.
	Custom_keys []map[string]interface{} `json:"custom_keys,omitempty"` // Custom keys and values assigned to the meeting.
	Contact_name string `json:"contact_name,omitempty"` // Contact name for registration
	In_meeting bool `json:"in_meeting,omitempty"` // Host meeting in India.
	Authentication_exception []map[string]interface{} `json:"authentication_exception,omitempty"` // The participants added here will receive unique meeting invite links and bypass authentication.
	Meeting_authentication bool `json:"meeting_authentication,omitempty"` // `true`- Only authenticated users can join meetings.
	Enforce_login bool `json:"enforce_login,omitempty"` // Only signed in users can join this meeting. **This field is deprecated and will not be supported in the future.** <br><br>As an alternative, use the "meeting_authentication", "authentication_option" and "authentication_domains" fields to understand the [authentication configurations](https://support.zoom.us/hc/en-us/articles/360037117472-Authentication-Profiles-for-Meetings-and-Webinars) set for the meeting.
	Audio string `json:"audio,omitempty"` // Determine how participants can join the audio portion of the meeting.<br>`both` - Both Telephony and VoIP.<br>`telephony` - Telephony only.<br>`voip` - VoIP only.
	Authentication_name string `json:"authentication_name,omitempty"` // Authentication name set in the [authentication profile](https://support.zoom.us/hc/en-us/articles/360037117472-Authentication-Profiles-for-Meetings-and-Webinars#h_5c0df2e1-cfd2-469f-bb4a-c77d7c0cca6f).
	Cn_meeting bool `json:"cn_meeting,omitempty"` // Host meeting in China.
	Host_video bool `json:"host_video,omitempty"` // Start video when the host joins the meeting.
	Registrants_confirmation_email bool `json:"registrants_confirmation_email,omitempty"` // Send confirmation email to registrants upon successful registration.
	Join_before_host bool `json:"join_before_host,omitempty"` // Allow participants to join the meeting before the host starts the meeting. Only used for scheduled or recurring meetings.
	Allow_multiple_devices bool `json:"allow_multiple_devices,omitempty"` // Allow attendees to join the meeting from multiple devices. This setting only works for meetings that require [registration](https://support.zoom.us/hc/en-us/articles/211579443-Setting-up-registration-for-a-meeting).
	Close_registration bool `json:"close_registration,omitempty"` // Close registration after event date
	Alternative_hosts string `json:"alternative_hosts,omitempty"` // Alternative host's emails or IDs: multiple values are separated by a semicolon.
	Authentication_option string `json:"authentication_option,omitempty"` // Meeting authentication option id.
	Alternative_hosts_email_notification bool `json:"alternative_hosts_email_notification,omitempty"` // Flag to determine whether to send email notifications to alternative hosts, default value is true.
	Breakout_room map[string]interface{} `json:"breakout_room,omitempty"` // Setting to [pre-assign breakout rooms](https://support.zoom.us/hc/en-us/articles/360032752671-Pre-assigning-participants-to-breakout-rooms#h_36f71353-4190-48a2-b999-ca129861c1f4).
	Contact_email string `json:"contact_email,omitempty"` // Contact email for registration
	Encryption_type string `json:"encryption_type,omitempty"` // Choose between enhanced encryption and [end-to-end encryption](https://support.zoom.us/hc/en-us/articles/360048660871) when starting or a meeting. When using end-to-end encryption, several features (e.g. cloud recording, phone/SIP/H.323 dial-in) will be **automatically disabled**. <br><br>The value of this field can be one of the following:<br> `enhanced_encryption`: Enhanced encryption. Encryption is stored in the cloud if you enable this option. <br> `e2ee`: [End-to-end encryption](https://support.zoom.us/hc/en-us/articles/360048660871). The encryption key is stored in your local device and can not be obtained by anyone else. Enabling this setting also **disables** the following features: join before host, cloud recording, streaming, live transcription, breakout rooms, polling, 1:1 private chat, and meeting reactions.
	Registrants_email_notification bool `json:"registrants_email_notification,omitempty"` // Send email notifications to registrants about approval, cancellation, denial of the registration. The value of this field must be set to true in order to use the `registrants_confirmation_email` field.
	Mute_upon_entry bool `json:"mute_upon_entry,omitempty"` // Mute participants upon entry.
	Language_interpretation map[string]interface{} `json:"language_interpretation,omitempty"`
}

// Panelist represents the Panelist schema from the OpenAPI specification
type Panelist struct {
	Email string `json:"email,omitempty"` // Panelist's email.
	Name string `json:"name,omitempty"` // Panelist's full name.
}

// QoSPhone represents the QoSPhone schema from the OpenAPI specification
type QoSPhone struct {
	Mos string `json:"mos,omitempty"` // Mean Opinion Score(MOS) measures voice quality on a scale of 1 to 5. A MOS greater than or equal to 3.5 means good quality, while below 3.5 means poor quality.
	Network_delay string `json:"network_delay,omitempty"` // The amount of time it takes for a VoIP packet to travel from one point to another.
	Avg_loss string `json:"avg_loss,omitempty"` // The average amount of packet loss, i.e., the percentage of packets that fail to arrive at their destination.
	Bitrate string `json:"bitrate,omitempty"` // The number of bits per second that can be transmitted along a digital network.
	Jitter string `json:"jitter,omitempty"` // The variation in the delay of received packets.
	Max_loss string `json:"max_loss,omitempty"` // The max amount of packet loss, i.e., the max percentage of packets that fail to arrive at their destination.
}

// AccountSettingsAuthenticationUpdate represents the AccountSettingsAuthenticationUpdate schema from the OpenAPI specification
type AccountSettingsAuthenticationUpdate struct {
}

// AccountSettingsTelephony represents the AccountSettingsTelephony schema from the OpenAPI specification
type AccountSettingsTelephony struct {
	Third_party_audio bool `json:"third_party_audio,omitempty"` // Users can join the meeting using the existing third party audio configuration.
	Audio_conference_info string `json:"audio_conference_info,omitempty"` // Third party audio conference info.
	Telephony_regions map[string]interface{} `json:"telephony_regions,omitempty"` // Indicates where most of the participants call into or call from duriing a meeting.
}

// DomainsList represents the DomainsList schema from the OpenAPI specification
type DomainsList struct {
	Domains []map[string]interface{} `json:"domains,omitempty"` // List of managed domain objects.
	Total_records int `json:"total_records,omitempty"` // Total records.
}

// UserUpdate represents the UserUpdate schema from the OpenAPI specification
type UserUpdate struct {
	Job_title string `json:"job_title,omitempty"` // User's job title.
	First_name string `json:"first_name,omitempty"` // User's first name. Cannot contain more than 5 Chinese characters.
	Language string `json:"language,omitempty"` // language
	Phone_country string `json:"phone_country,omitempty"` // **Note:** This field has been **deprecated** and will not be supported in the future. Use the **country** field of the **phone_numbers** object instead to select the country for the phone number. [Country ID](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#countries) of the phone number. For example, if the phone number provided in the `phone_number` field is a Brazil based number, the value of the `phone_country` field should be `BR`.
	Pmi int `json:"pmi,omitempty"` // Personal meeting ID: length must be 10.
	Timezone string `json:"timezone,omitempty"` // The time zone ID for a user profile. For this parameter value please refer to the ID value in the [timezone](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones) list.
	TypeField int `json:"type,omitempty"` // User types:<br>`1` - Basic.<br>`2` - Licensed.<br>`3` - On-prem.<br>`99` - None (this can only be set with `ssoCreate`).
	Custom_attributes map[string]interface{} `json:"custom_attributes,omitempty"` // Custom attribute(s) of the user.
	Phone_number string `json:"phone_number,omitempty"` // **Note:** This field has been **deprecated** and will not be supported in the future. Use the **phone_numbers** field instead to assign phone number(s) to a user. Phone number of the user. To update a phone number, you must also provide the `phone_country` field.
	Vanity_name string `json:"vanity_name,omitempty"` // Personal meeting room name.
	Last_name string `json:"last_name,omitempty"` // User's last name. Cannot contain more than 5 Chinese characters.
	Group_id string `json:"group_id,omitempty"` // Provide unique identifier of the group that you would like to add a [pending user](https://support.zoom.us/hc/en-us/articles/201363183-Managing-users#h_13c87a2a-ecd6-40ad-be61-a9935e660edb) to. The value of this field can be retrieved from [List Groups](https://marketplace.zoom.us/docs/api-reference/zoom-api/groups/groups) API.
	Cms_user_id string `json:"cms_user_id,omitempty"` // Kaltura user ID.
	Dept string `json:"dept,omitempty"` // Department for user profile: use for report.
	Manager string `json:"manager,omitempty"` // The manager for the user.
	Phone_numbers map[string]interface{} `json:"phone_numbers,omitempty"`
	Company string `json:"company,omitempty"` // User's company.
	Host_key string `json:"host_key,omitempty"` // Host key. It should be a 6-10 digit number.
	Location string `json:"location,omitempty"` // User's location.
	Use_pmi bool `json:"use_pmi,omitempty"` // Use Personal Meeting ID for instant meetings.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Audio_url string `json:"audio_url,omitempty"` // The global dial-in URL for a TSP enabled account. The URL must be valid with a max-length of 512 characters.
}

// CustomQuestion represents the CustomQuestion schema from the OpenAPI specification
type CustomQuestion struct {
	Value string `json:"value,omitempty"`
	Title string `json:"title,omitempty"`
}

// PAC represents the PAC schema from the OpenAPI specification
type PAC struct {
	Conference_id int `json:"conference_id,omitempty"` // Conference ID.
	Dedicated_dial_in_number []map[string]interface{} `json:"dedicated_dial_in_number,omitempty"` // List of dedicated dial-in numbers.
	Global_dial_in_numbers []map[string]interface{} `json:"global_dial_in_numbers,omitempty"` // List of global dial-in numbers.
	Listen_only_password string `json:"listen_only_password,omitempty"` // Listen-Only passcode: numeric value - length is less than 6.
	Participant_password string `json:"participant_password,omitempty"` // Participant passcode: numeric value - length is less than 6.
}

// WebinarList represents the WebinarList schema from the OpenAPI specification
type WebinarList struct {
	Page_size int `json:"page_size,omitempty"` // The number of records returned with a single API call.
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_number int `json:"page_number,omitempty"` // **Deprecated**: This field has been deprecated. Please use the "next_page_token" field for pagination instead of this field. The page number of the current results.
	Webinars []map[string]interface{} `json:"webinars,omitempty"` // List of webinar objects.
}

// AccountSettingsInMeeting represents the AccountSettingsInMeeting schema from the OpenAPI specification
type AccountSettingsInMeeting struct {
	Meeting_reactions bool `json:"meeting_reactions,omitempty"` // Enable or disable meeting reactions. <br> `true`: Allow meeting participants to communicate without interrupting by reacting with an emoji that shows on their video.<br> `false`: Do not enable meeting reactions.
	Dscp_marking bool `json:"dscp_marking,omitempty"` // DSCP marking.
	Screen_sharing bool `json:"screen_sharing,omitempty"` // Allow screen sharing.
	Use_html_format_email bool `json:"use_html_format_email,omitempty"` // Use HTML formatted email for the Outlook plugin.
	Virtual_background_settings map[string]interface{} `json:"virtual_background_settings,omitempty"` // Settings to manage virtual background.
	File_transfer bool `json:"file_transfer,omitempty"` // Indicates whether [in-meeting file transfer](https://support.zoom.us/hc/en-us/articles/209605493-In-meeting-file-transfer) setting has been enabled on the account or not.
	Custom_live_streaming_service bool `json:"custom_live_streaming_service,omitempty"` // Custom live streaming.
	Allow_participants_to_rename bool `json:"allow_participants_to_rename,omitempty"` // If the value of this field is set to `true`, meeting participants and webinar panelists can be allowed to rename themselves during a meeting or a webinar.
	Breakout_room bool `json:"breakout_room,omitempty"` // Allow host to split meeting participants into separate, smaller rooms.
	Polling bool `json:"polling,omitempty"` // Add "Polls" to the meeting controls.
	Dscp_video int `json:"dscp_video,omitempty"` // DSCP video.
	Group_hd bool `json:"group_hd,omitempty"` // Activate higher quality video for host and participants. Please note: This will use more bandwidth.
	Annotation bool `json:"annotation,omitempty"` // Allow participants to use annotation tools to add information to shared screens.
	Watermark bool `json:"watermark,omitempty"` // Add a watermark when viewing a shared screen.
	Who_can_share_screen string `json:"who_can_share_screen,omitempty"` // Indicates who can share their screen or content during meetings. The value can be one of the following: <br> `host`: Only host can share the screen.<br> `all`: Both hosts and attendees can share their screen during meetings. For Webinar, the hosts and panelists can start screen sharing, but not the attendees.
	Ports_range string `json:"ports_range,omitempty"` // The listening ports range, separated by a comma (ex 55,56). The ports range must be between 1 to 65535.
	Anonymous_question_answer bool `json:"anonymous_question_answer,omitempty"` // Allow an anonymous Q&A in a webinar.
	P2p_ports bool `json:"p2p_ports,omitempty"` // Peer to peer listening ports range.
	Attendee_on_hold bool `json:"attendee_on_hold,omitempty"` // Allow host to put attendee on hold. **This field has been deprecated and is no longer supported.**
	Original_audio bool `json:"original_audio,omitempty"` // Allow users to select original sound in their client settings.
	Post_meeting_feedback bool `json:"post_meeting_feedback,omitempty"` // Display a thumbs up or down survey at the end of each meeting.
	Show_a_join_from_your_browser_link bool `json:"show_a_join_from_your_browser_link,omitempty"` // If the value of this field is set to `true`, you will allow participants to join a meeting directly from their browser and bypass the Zoom application download process. This is a workaround for participants who are unable to download, install, or run applications. Note that the meeting experience from the browser is limited.
	Auto_saving_chat bool `json:"auto_saving_chat,omitempty"` // Automatically save all in-meeting chats so that the host does not need to manually save the chat transcript after the meeting starts.
	Virtual_background bool `json:"virtual_background,omitempty"` // Allow users to replace their background with any selected image. Choose or upload an image in the Zoom desktop application settings.
	Private_chat bool `json:"private_chat,omitempty"` // Allow a meeting participant to send a private message to another participant.
	Remote_control bool `json:"remote_control,omitempty"` // Allow users to request remote control.
	Custom_service_instructions string `json:"custom_service_instructions,omitempty"` // Custom service instructions.
	Far_end_camera_control bool `json:"far_end_camera_control,omitempty"` // Allow another user to take control of your camera during a meeting.
	Allow_show_zoom_windows bool `json:"allow_show_zoom_windows,omitempty"` // Show the Zoom desktop application when sharing screens.
	Auto_answer bool `json:"auto_answer,omitempty"` // Enable users to see and add contacts to the "auto-answer group" in the chat contact list. Any call from members of this group will automatically be answered.
	Co_host bool `json:"co_host,omitempty"` // Allow the host to add co-hosts.
	Data_center_regions []string `json:"data_center_regions,omitempty"` // If you have set the value of `custom_data_center_regions` to `true`, specify the data center regions that you would like to opt in to (country codes from among: ["EU", "HK", "AU", "IN", "LA", "TY", "CN", "US", "CA"]).
	Entry_exit_chime string `json:"entry_exit_chime,omitempty"` // Play sound when participants join or leave.<br>`host` - Heard by host only.<br>`all` - Heard by host and all attendees.<br>`none` - Disable.
	Chat bool `json:"chat,omitempty"` // Allow meeting participants to send a message that is visible to all participants.
	Stereo_audio bool `json:"stereo_audio,omitempty"` // Allow users to select stereo audio in their client settings.
	Feedback bool `json:"feedback,omitempty"` // Add a "Feedback" tab to the Windows Settings or Mac Preferences dialog. Enable users to provide feedback to Zoom at the end of the meeting.
	Workplace_by_facebook bool `json:"workplace_by_facebook,omitempty"` // Workplace by facebook.
	Closed_caption bool `json:"closed_caption,omitempty"` // Allow a host to type closed captions. Enable a host to assign a participant or third party device to add closed captions.
	P2p_connetion bool `json:"p2p_connetion,omitempty"` // Peer to peer connection while only two people are in a meeting.
	Record_play_own_voice bool `json:"record_play_own_voice,omitempty"` // Record and play their own voice.
	Webinar_question_answer bool `json:"webinar_question_answer,omitempty"` // Allow a Q&A in a webinar.
	Dscp_audio int `json:"dscp_audio,omitempty"` // DSCP audio.
	E2e_encryption bool `json:"e2e_encryption,omitempty"` // Zoom requires encryption for all data between the Zoom cloud, Zoom client, and Zoom Room. Require encryption for 3rd party endpoints (H323/SIP).
	Sending_default_email_invites bool `json:"sending_default_email_invites,omitempty"` // Only show the default email when sending email invites.
	Show_meeting_control_toolbar bool `json:"show_meeting_control_toolbar,omitempty"` // Always show the meeting control toolbar.
	Whiteboard bool `json:"whiteboard,omitempty"` // Allow participants to share a whiteboard that includes annotation tools.
	Request_permission_to_unmute bool `json:"request_permission_to_unmute,omitempty"` // Indicates whether the [**Request permission to unmute participants**](https://support.zoom.us/hc/en-us/articles/203435537-Muting-and-unmuting-participants-in-a-meeting#h_01EGK4XFWS1SJGZ71MYGKF7260) option has been enabled for the account or not.
	Allow_live_streaming bool `json:"allow_live_streaming,omitempty"` // Allow live streaming.
	Custom_data_center_regions bool `json:"custom_data_center_regions,omitempty"` // If set to `true`, account owners and admins on paid accounts can [select data center regions](https://support.zoom.us/hc/en-us/articles/360042411451-Selecting-data-center-regions-for-hosted-meetings-and-webinars) to use for hosting their real-time meeting and webinar traffic. These regions can be provided in the `data_center_regions` field. If set to `false`, the regions cannot be customized and the default regions will be used.
	Alert_guest_join bool `json:"alert_guest_join,omitempty"` // Identify guest participants in a meeting or webinar.
	Who_can_share_screen_when_someone_is_sharing string `json:"who_can_share_screen_when_someone_is_sharing,omitempty"` // Indicates who is allowed to start sharing screen when someone else in the meeting is sharing their screen. The value can be one of the following:<br> `host`: Only a host can share the screen when someone else is sharing.<br> `all`: Anyone in the meeting is allowed to start sharing their screen when someone else is sharing. For Webinar, the hosts and panelists can start screen sharing, but not the attendees.
}

// RecordingSettings represents the RecordingSettings schema from the OpenAPI specification
type RecordingSettings struct {
	Password string `json:"password,omitempty"` // Enable password protection for the recording by setting a password. The password must have a minimum of **eight** characters with a mix of numbers, letters and special characters.<br><br> **Note:** If the account owner or the admin has set minimum password strength requirements for recordings via Account Settings, the password value provided here must meet those requirements. <br><br>If the requirements are enabled, you can view those requirements by calling either the [Get User Settings API](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usersettings) or the [Get Account Settings](https://marketplace.zoom.us/docs/api-reference/zoom-api/accounts/accountsettings) API.
	Recording_authentication bool `json:"recording_authentication,omitempty"` // Only authenticated users can view.
	Send_email_to_host bool `json:"send_email_to_host,omitempty"` // Send an email to host when someone registers to view the recording. This applies for On-demand recordings only.
	Viewer_download bool `json:"viewer_download,omitempty"` // Determine whether a viewer can download the recording file or not.
	Approval_type int `json:"approval_type,omitempty"` // Approval type for the registration.<br> `0`- Automatically approve the registration when a user registers.<br> `1` - Manually approve or deny the registration of a user.<br> `2` - No registration required to view the recording.
	Show_social_share_buttons bool `json:"show_social_share_buttons,omitempty"` // Show social share buttons on registration page. This applies for On-demand recordings only.
	Authentication_option string `json:"authentication_option,omitempty"` // Authentication Options.
	Authentication_domains string `json:"authentication_domains,omitempty"` // Authentication domains.
	Share_recording string `json:"share_recording,omitempty"` // Determine how the meeting recording is shared.
	Topic string `json:"topic,omitempty"` // Name of the recording.
	On_demand bool `json:"on_demand,omitempty"` // Determine whether registration isrequired to view the recording.
}

// WebinarSettings represents the WebinarSettings schema from the OpenAPI specification
type WebinarSettings struct {
	Registration_type int `json:"registration_type,omitempty"` // Registration types. Only used for recurring webinars with a fixed time.<br>`1` - Attendees register once and can attend any of the webinar sessions.<br>`2` - Attendees need to register for each session in order to attend.<br>`3` - Attendees register once and can choose one or more sessions to attend.
	Host_video bool `json:"host_video,omitempty"` // Start video when host joins webinar.
	Follow_up_absentees_email_notification map[string]interface{} `json:"follow_up_absentees_email_notification,omitempty"` // Send follow-up email to absentees.
	Hd_video bool `json:"hd_video,omitempty"` // Default to HD video.
	Email_language string `json:"email_language,omitempty"` // Set the email language to one of the following: `en-US`,`de-DE`,`es-ES`,`fr-FR`,`jp-JP`,`pt-PT`,`ru-RU`,`zh-CN`, `zh-TW`, `ko-KO`, `it-IT`, `vi-VN`.
	Audio string `json:"audio,omitempty"` // Determine how participants can join the audio portion of the webinar.
	Close_registration bool `json:"close_registration,omitempty"` // Close registration after event date.
	Attendees_and_panelists_reminder_email_notification map[string]interface{} `json:"attendees_and_panelists_reminder_email_notification,omitempty"` // Send reminder email to attendees and panelists.
	Approval_type int `json:"approval_type,omitempty"` // `0` - Automatically approve.<br>`1` - Manually approve.<br>`2` - No registration required.
	Enforce_login bool `json:"enforce_login,omitempty"` // Only signed in users can join this meeting. **This field is deprecated and will not be supported in the future.** <br><br>As an alternative, use the "meeting_authentication", "authentication_option" and "authentication_domains" fields to understand the [authentication configurations](https://support.zoom.us/hc/en-us/articles/360037117472-Authentication-Profiles-for-Meetings-and-Webinars) set for the Webinar.
	Alternative_hosts string `json:"alternative_hosts,omitempty"` // Alternative host emails or IDs. Multiple values separated by comma.
	Meeting_authentication bool `json:"meeting_authentication,omitempty"` // `true`- Only authenticated users can join Webinar.
	Registrants_email_notification bool `json:"registrants_email_notification,omitempty"` // Send email notifications to registrants about approval, cancellation, denial of the registration. The value of this field must be set to true in order to use the `registrants_confirmation_email` field.
	Contact_email string `json:"contact_email,omitempty"` // Contact email for registration
	Survey_url string `json:"survey_url,omitempty"` // Survey url for post webinar survey
	Contact_name string `json:"contact_name,omitempty"` // Contact name for registration
	Show_share_button bool `json:"show_share_button,omitempty"` // Show social share buttons on the registration page.
	Follow_up_attendees_email_notification map[string]interface{} `json:"follow_up_attendees_email_notification,omitempty"` // Send follow-up email to attendees.
	Practice_session bool `json:"practice_session,omitempty"` // Enable practice session.
	Allow_multiple_devices bool `json:"allow_multiple_devices,omitempty"` // Allow attendees to join from multiple devices.
	Registrants_confirmation_email bool `json:"registrants_confirmation_email,omitempty"` // Send confirmation email to registrants
	Notify_registrants bool `json:"notify_registrants,omitempty"` // Send notification email to registrants when the host updates a webinar.
	Registrants_restrict_number int `json:"registrants_restrict_number,omitempty"` // Restrict number of registrants for a webinar. By default, it is set to `0`. A `0` value means that the restriction option is disabled. Provide a number higher than 0 to restrict the webinar registrants by the that number.
	Question_and_answer map[string]interface{} `json:"question_and_answer,omitempty"` // [Q&A](https://support.zoom.us/hc/en-us/articles/203686015-Using-Q-A-as-the-webinar-host#:~:text=Overview,and%20upvote%20each%20other's%20questions.) for webinar.
	On_demand bool `json:"on_demand,omitempty"` // Make the webinar on-demand
	Authentication_domains string `json:"authentication_domains,omitempty"` // If user has configured ["Sign Into Zoom with Specified Domains"](https://support.zoom.us/hc/en-us/articles/360037117472-Authentication-Profiles-for-Meetings-and-Webinars#h_5c0df2e1-cfd2-469f-bb4a-c77d7c0cca6f) option, this will list the domains that are authenticated.
	Global_dial_in_countries []string `json:"global_dial_in_countries,omitempty"` // List of global dial-in countries
	Auto_recording string `json:"auto_recording,omitempty"` // Automatic recording:<br>`local` - Record on local.<br>`cloud` - Record on cloud.<br>`none` - Disabled.
	Authentication_option string `json:"authentication_option,omitempty"` // Webinar authentication option id.
	Panelists_video bool `json:"panelists_video,omitempty"` // Start video when panelists join webinar.
	Post_webinar_survey bool `json:"post_webinar_survey,omitempty"` // Zoom will open a survey page in attendees' browsers after leaving the webinar
	Authentication_name string `json:"authentication_name,omitempty"` // Authentication name set in the [authentication profile](https://support.zoom.us/hc/en-us/articles/360037117472-Authentication-Profiles-for-Meetings-and-Webinars#h_5c0df2e1-cfd2-469f-bb4a-c77d7c0cca6f).
	Enforce_login_domains string `json:"enforce_login_domains,omitempty"` // Only signed in users with specified domains can join meetings. **This field is deprecated and will not be supported in the future.** <br><br>As an alternative, use the "meeting_authentication", "authentication_option" and "authentication_domains" fields to understand the [authentication configurations](https://support.zoom.us/hc/en-us/articles/360037117472-Authentication-Profiles-for-Meetings-and-Webinars) set for the Webinar.
	Panelists_invitation_email_notification bool `json:"panelists_invitation_email_notification,omitempty"` // * `true`: Send invitation email to panelists. * `false`: Do not send invitation email to panelists.
}

// ZoomRoomList represents the ZoomRoomList schema from the OpenAPI specification
type ZoomRoomList struct {
	Page_size int `json:"page_size,omitempty"` // The number of records returned with a single API call.
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
	Next_page_token string `json:"next_page_token,omitempty"`
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_number int `json:"page_number,omitempty"` // The page number of the current results.
	Zoom_rooms []map[string]interface{} `json:"zoom_rooms,omitempty"` // Array of Zoom Rooms
}

// UserSettingsRecording represents the UserSettingsRecording schema from the OpenAPI specification
type UserSettingsRecording struct {
	Ip_address_access_control map[string]interface{} `json:"ip_address_access_control,omitempty"` // Setting to allow cloud recording access only from specific IP address ranges.
	Ask_host_to_confirm_disclaimer bool `json:"ask_host_to_confirm_disclaimer,omitempty"` // Ask host to confirm the disclaimer.
	Record_audio_file bool `json:"record_audio_file,omitempty"` // Record an audio only file.
	Auto_delete_cmr_days int `json:"auto_delete_cmr_days,omitempty"` // A specified number of days of auto delete cloud recordings.
	Host_pause_stop_recording bool `json:"host_pause_stop_recording,omitempty"` // Host can pause/stop the auto recording in the cloud.
	Local_recording bool `json:"local_recording,omitempty"` // Local recording.
	Record_speaker_view bool `json:"record_speaker_view,omitempty"` // Record the active speaker view.
	Recording_disclaimer bool `json:"recording_disclaimer,omitempty"` // Show a disclaimer to participants before a recording starts
	Auto_recording string `json:"auto_recording,omitempty"` // Automatic recording:<br>`local` - Record on local.<br>`cloud` - Record on cloud.<br>`none` - Disabled.
	Save_chat_text bool `json:"save_chat_text,omitempty"` // Save chat text from the meeting.
	Show_timestamp bool `json:"show_timestamp,omitempty"` // Show timestamp on video.
	Cloud_recording bool `json:"cloud_recording,omitempty"` // Cloud recording.
	Recording_audio_transcript bool `json:"recording_audio_transcript,omitempty"` // Audio transcript.
	Record_gallery_view bool `json:"record_gallery_view,omitempty"` // Record the gallery view.
	Ask_participants_to_consent_disclaimer bool `json:"ask_participants_to_consent_disclaimer,omitempty"` // This field can be used if `recording_disclaimer` is set to true. This field indicates whether or not you would like to ask participants for consent when a recording starts. The value can be one of the following:<br> * `true`: Ask participants for consent when a recording starts. <br> * `false`: Do not ask participants for consent when a recording starts.
	Recording_password_requirement map[string]interface{} `json:"recording_password_requirement,omitempty"` // This object represents the minimum passcode requirements set for recordings via Account Recording Settings.
	Auto_delete_cmr bool `json:"auto_delete_cmr,omitempty"` // Auto delete cloud recordings.
}

// SessionWebinarUpdate represents the SessionWebinarUpdate schema from the OpenAPI specification
type SessionWebinarUpdate struct {
	Agenda string `json:"agenda,omitempty"` // Webinar description.
	Recurrence map[string]interface{} `json:"recurrence,omitempty"` // Recurrence object. Use this object only for a meeting with type `8` i.e., a recurring meeting with fixed time.
	Topic string `json:"topic,omitempty"` // Webinar topic.
	Tracking_fields []map[string]interface{} `json:"tracking_fields,omitempty"` // Tracking fields
	Password string `json:"password,omitempty"` // [Webinar passcode](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords). By default, passcode may only contain the following characters: [a-z A-Z 0-9 @ - _ * !] and can have a maximum of 10 characters. **Note:** If the account owner or the admin has configured [minimum passcode requirement settings](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords#h_a427384b-e383-4f80-864d-794bf0a37604), the passcode value provided here must meet those requirements. <br><br>If the requirements are enabled, you can view those requirements by calling either the [Get User Settings API](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usersettings) or the [Get Account Settings](https://marketplace.zoom.us/docs/api-reference/zoom-api/accounts/accountsettings) API. If "**Require a passcode when scheduling new meetings**" setting has been **enabled** **and** [locked](https://support.zoom.us/hc/en-us/articles/115005269866-Using-Tiered-Settings#locked) for the user, the passcode field will be autogenerated for the Webinar in the response even if it is not provided in the API request. <br><br>
	Settings interface{} `json:"settings,omitempty"`
	Start_time string `json:"start_time,omitempty"` // Webinar start time, in the format "yyyy-MM-dd'T'HH:mm:ss'Z'." Should be in GMT time. In the format "yyyy-MM-dd'T'HH:mm:ss." This should be in local time and the timezone should be specified. Only used for scheduled webinars and recurring webinars with a fixed time.
	Timezone string `json:"timezone,omitempty"` // Time zone to format start_time. For example, "America/Los_Angeles". For scheduled meetings only. Please reference our [time zone](#timezones) list for supported time zones and their formats.
	TypeField int `json:"type,omitempty"` // Webinar Types:<br>`5` - webinar.<br>`6` - Recurring webinar with no fixed time.<br>`9` - Recurring webinar with a fixed time.
	Duration int `json:"duration,omitempty"` // Webinar duration (minutes). Used for scheduled webinar only.
}

// UserSettingsInMeeting represents the UserSettingsInMeeting schema from the OpenAPI specification
type UserSettingsInMeeting struct {
	Attendee_on_hold bool `json:"attendee_on_hold,omitempty"` // Allow host to put attendee on hold. **This field has been deprecated and is no longer supported.**
	Annotation bool `json:"annotation,omitempty"` // Allow participants to use annotation tools.
	File_transfer bool `json:"file_transfer,omitempty"` // Indicates whether [in-meeting file transfer](https://support.zoom.us/hc/en-us/articles/209605493-In-meeting-file-transfer) setting has been enabled for the user or not.
	Far_end_camera_control bool `json:"far_end_camera_control,omitempty"` // Allow another user to take control of the camera.
	Polling bool `json:"polling,omitempty"` // Add polls to the meeting controls.
	Private_chat bool `json:"private_chat,omitempty"` // Enable 1:1 private chat between participants during meetings.
	Who_can_share_screen_when_someone_is_sharing string `json:"who_can_share_screen_when_someone_is_sharing,omitempty"` // Indicates who is allowed to start sharing screen when someone else in the meeting is sharing their screen. The value can be one of the following:<br> `host`: Only a host can share the screen when someone else is sharing.<br> `all`: Anyone in the meeting is allowed to start sharing their screen when someone else is sharing. For Webinar, the hosts and panelists can start screen sharing, but not the attendees.
	Allow_live_streaming bool `json:"allow_live_streaming,omitempty"` // Allow live streaming.
	Chat bool `json:"chat,omitempty"` // Enable chat during meeting for all participants.
	Remote_control bool `json:"remote_control,omitempty"` // Enable remote control during screensharing.
	Record_play_voice bool `json:"record_play_voice,omitempty"` // Record and play their own voice.
	Screen_sharing bool `json:"screen_sharing,omitempty"` // Allow host and participants to share their screen or content during meetings
	Share_dual_camera bool `json:"share_dual_camera,omitempty"` // Share dual camera (deprecated).
	Breakout_room bool `json:"breakout_room,omitempty"` // Allow host to split meeting participants into separate breakout rooms.
	Waiting_room bool `json:"waiting_room,omitempty"` // Enable Waiting room - if enabled, attendees can only join after host approves.
	Auto_saving_chat bool `json:"auto_saving_chat,omitempty"` // Auto save all in-meeting chats.
	Request_permission_to_unmute bool `json:"request_permission_to_unmute,omitempty"` // Indicates whether the [**Request permission to unmute participants**](https://support.zoom.us/hc/en-us/articles/203435537-Muting-and-unmuting-participants-in-a-meeting#h_01EGK4XFWS1SJGZ71MYGKF7260) option has been enabled for the user or not.
	Non_verbal_feedback bool `json:"non_verbal_feedback,omitempty"` // Enable non-verbal feedback through screens.
	Who_can_share_screen string `json:"who_can_share_screen,omitempty"` // Indicates who can share their screen or content during meetings. The value can be one of the following: <br> `host`: Only host can share the screen.<br> `all`: Both hosts and attendees can share their screen during meetings. For Webinar, the hosts and panelists can start screen sharing, but not the attendees.
	Workplace_by_facebook bool `json:"workplace_by_facebook,omitempty"` // Allow livestreaming by host through Workplace by Facebook.
	Feedback bool `json:"feedback,omitempty"` // Enable option to send feedback to Zoom at the end of the meeting.
	Virtual_background bool `json:"virtual_background,omitempty"` // Enable virtual background.
	Virtual_background_settings map[string]interface{} `json:"virtual_background_settings,omitempty"` // Settings to manage virtual background.
	Group_hd bool `json:"group_hd,omitempty"` // Enable group HD video.
	E2e_encryption bool `json:"e2e_encryption,omitempty"` // Zoom requires encryption for all data between the Zoom cloud, Zoom client, and Zoom Room. Require encryption for 3rd party endpoints (H323/SIP).
	Entry_exit_chime string `json:"entry_exit_chime,omitempty"` // Play sound when participants join or leave:<br>`host` - When host joins or leaves.<br>`all` - When any participant joins or leaves.<br>`none` - No join or leave sound.
	Show_meeting_control_toolbar bool `json:"show_meeting_control_toolbar,omitempty"` // Always show meeting controls during a meeting.
	Closed_caption bool `json:"closed_caption,omitempty"` // Enable closed captions.
	Custom_service_instructions string `json:"custom_service_instructions,omitempty"` // Custom service instructions.
	Custom_live_streaming_service bool `json:"custom_live_streaming_service,omitempty"` // Allow custom live streaming.
	Remote_support bool `json:"remote_support,omitempty"` // Allow host to provide 1:1 remote support to a participant.
	Data_center_regions []string `json:"data_center_regions,omitempty"` // If you have set the value of `custom_data_center_regions` to `true`, specify the data center regions that you would like to opt in to (country codes from among: ["DE", "NL", "HK", "AU", "IN", "LA", "TY", "CN", "US", "CA"]).
	Co_host bool `json:"co_host,omitempty"` // Allow the host to add co-hosts.
	Custom_data_center_regions bool `json:"custom_data_center_regions,omitempty"` // If set to `true`, you can [select data center regions](https://support.zoom.us/hc/en-us/articles/360042411451-Selecting-data-center-regions-for-hosted-meetings-and-webinars) to use for hosting your real-time meeting and webinar traffic. These regions can be provided in the `data_center_regions` field. If set to `false`, the regions cannot be customized and the default regions will be used.
}

// Meeting represents the Meeting schema from the OpenAPI specification
type Meeting struct {
	Topic string `json:"topic,omitempty"` // Meeting topic.
	Tracking_fields []interface{} `json:"tracking_fields,omitempty"` // Tracking fields
	Agenda string `json:"agenda,omitempty"` // Meeting description.
	TypeField int `json:"type,omitempty"` // Meeting Type:<br>`1` - Instant meeting.<br>`2` - Scheduled meeting.<br>`3` - Recurring meeting with no fixed time.<br>`8` - Recurring meeting with fixed time.
	Settings map[string]interface{} `json:"settings,omitempty"` // Meeting settings.
	Duration int `json:"duration,omitempty"` // Meeting duration (minutes). Used for scheduled meetings only.
	Password string `json:"password,omitempty"` // Password to join the meeting. Password may only contain the following characters: [a-z A-Z 0-9 @ - _ *]. Max of 10 characters.
	Recurrence map[string]interface{} `json:"recurrence,omitempty"` // Recurrence object. Use this object only for a meeting with type `8` i.e., a recurring meeting with fixed time.
	Start_time string `json:"start_time,omitempty"` // Meeting start time. When using a format like "yyyy-MM-dd'T'HH:mm:ss'Z'", always use GMT time. When using a format like "yyyy-MM-dd'T'HH:mm:ss", you should use local time and specify the time zone. This is only used for scheduled meetings and recurring meetings with a fixed time.
	Timezone string `json:"timezone,omitempty"` // Time zone to format start_time. For example, "America/Los_Angeles". For scheduled meetings only. Please reference our [time zone](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones) list for supported time zones and their formats.
}

// AccountSettingsTSP represents the AccountSettingsTSP schema from the OpenAPI specification
type AccountSettingsTSP struct {
	Call_out_countries []interface{} `json:"call_out_countries,omitempty"` // Call Out Countries/Regions
	Display_toll_free_numbers bool `json:"display_toll_free_numbers,omitempty"` // Display toll-free numbers
	Show_international_numbers_link bool `json:"show_international_numbers_link,omitempty"` // Show international numbers link on the invitation email
	Call_out bool `json:"call_out,omitempty"` // Call Out
}

// AccountUpdateSettings represents the AccountUpdateSettings schema from the OpenAPI specification
type AccountUpdateSettings struct {
	Security map[string]interface{} `json:"security,omitempty"` // [Security settings](https://support.zoom.us/hc/en-us/articles/360034675592-Advanced-security-settings#h_bf8a25f6-9a66-447a-befd-f02ed3404f89) of an Account.
	Telephony map[string]interface{} `json:"telephony,omitempty"` // Account Settings Update: Telephony.
	Tsp map[string]interface{} `json:"tsp,omitempty"` // Account Settings: TSP.
	Zoom_rooms map[string]interface{} `json:"zoom_rooms,omitempty"` // Account Settings: Zoom Rooms.
	Email_notification map[string]interface{} `json:"email_notification,omitempty"` // Account Settings: Notification.
	Schedule_meeting map[string]interface{} `json:"schedule_meeting,omitempty"` // Account Settings: Schedule Meeting.
	Profile map[string]interface{} `json:"profile,omitempty"`
	Feature map[string]interface{} `json:"feature,omitempty"` // Account Settings: Feature.
	In_meeting map[string]interface{} `json:"in_meeting,omitempty"` // Account Settings: In Meeting.
	Integration map[string]interface{} `json:"integration,omitempty"` // Account Settings: Integration.
	Recording map[string]interface{} `json:"recording,omitempty"` // Account Settings: Recording.
}

// RoleList represents the RoleList schema from the OpenAPI specification
type RoleList struct {
	Roles []interface{} `json:"roles,omitempty"` // List of Roles objects
	Total_records int `json:"total_records,omitempty"` // The number of all records available across pages
}

// Recording represents the Recording schema from the OpenAPI specification
type Recording struct {
	Deleted_time string `json:"deleted_time,omitempty"` // The time at which recording was deleted. Returned in the response only for trash query.
	File_size float64 `json:"file_size,omitempty"` // The recording file size.
	Meeting_id string `json:"meeting_id,omitempty"` // The meeting ID.
	Recording_start string `json:"recording_start,omitempty"` // The recording start time.
	Download_url string `json:"download_url,omitempty"` // The URL using which the recording file can be downloaded. **To access a private or password protected cloud recording of a user in your account, you can use a [Zoom JWT App Type](https://marketplace.zoom.us/docs/guides/getting-started/app-types/create-jwt-app). Use the generated JWT token as the value of the `access_token` query parameter and include this query parameter at the end of the URL as shown in the example.** <br> Example: `https://api.zoom.us/recording/download/{{ Download Path }}?access_token={{ JWT Token }}` **Similarly, if the user has installed your OAuth app that contains recording scope(s), you can also use the user's [OAuth access token](https://marketplace.zoom.us/docs/guides/auth/oauth) to download the Cloud Recording.**<br> Example: `https://api.zoom.us/recording/download/{{ Download Path }}?access_token={{ OAuth Access Token }}`
	Recording_end string `json:"recording_end,omitempty"` // The recording end time. Response in general query.
	Recording_type string `json:"recording_type,omitempty"` // The recording type. The value of this field can be one of the following:<br>`shared_screen_with_speaker_view(CC)`<br>`shared_screen_with_speaker_view`<br>`shared_screen_with_gallery_view`<br>`speaker_view`<br>`gallery_view`<br>`shared_screen`<br>`audio_only`<br>`audio_transcript`<br>`chat_file`<br>`active_speaker`<br>`poll`
	Status string `json:"status,omitempty"` // The recording status.
	File_type string `json:"file_type,omitempty"` // The recording file type. The value of this field could be one of the following:<br> `MP4`: Video file of the recording.<br>`M4A` Audio-only file of the recording.<br>`TIMELINE`: Timestamp file of the recording in JSON file format. To get a timeline file, the "Add a timestamp to the recording" setting must be enabled in the [recording settings](https://support.zoom.us/hc/en-us/articles/203741855-Cloud-recording#h_3f14c3a4-d16b-4a3c-bbe5-ef7d24500048). The time will display in the host's timezone, set on their Zoom profile. <br> `TRANSCRIPT`: Transcription file of the recording in VTT format.<br> `CHAT`: A TXT file containing in-meeting chat messages that were sent during the meeting.<br>`CC`: File containing closed captions of the recording in VTT file format.<br>`CSV`: File containing polling data in csv format. <br> A recording file object with file type of either `CC` or `TIMELINE` **does not have** the following properties:<br> 	`id`, `status`, `file_size`, `recording_type`, and `play_url`.
	Id string `json:"id,omitempty"` // The recording file ID. Included in the response of general query.
	Play_url string `json:"play_url,omitempty"` // The URL using which a recording file can be played.
}

// Group represents the Group schema from the OpenAPI specification
type Group struct {
	Name string `json:"name,omitempty"` // Group name.
	Total_members int `json:"total_members,omitempty"` // Total number of members in this group.
}

// UserSettingsFeatureUpdate represents the UserSettingsFeatureUpdate schema from the OpenAPI specification
type UserSettingsFeatureUpdate struct {
	Webinar_capacity int `json:"webinar_capacity,omitempty"` // Set the Webinar capacity for a user who has the Webinar feature enabled. The value of this field can be 100, 500, 1000, 3000, 5000 or 10000.
	Zoom_phone bool `json:"zoom_phone,omitempty"` // Zoom phone feature.
	Large_meeting bool `json:"large_meeting,omitempty"` // Enable [large meeting](https://support.zoom.us/hc/en-us/articles/201362823-What-is-a-Large-Meeting-) feature for the user.
	Large_meeting_capacity int `json:"large_meeting_capacity,omitempty"` // Set the meeting capacity for the user if the user has **Large meeting** feature enabled. The value for the field can be either 500 or 1000.
	Meeting_capacity int `json:"meeting_capacity,omitempty"` // Set a user's meeting capacity. User’s meeting capacity denotes the maximum number of participants that can join a meeting scheduled by the user.
	Webinar bool `json:"webinar,omitempty"` // Enable Webinar feature for the user.
}

// AccountSettingsRecordingAuthenticationUpdate represents the AccountSettingsRecordingAuthenticationUpdate schema from the OpenAPI specification
type AccountSettingsRecordingAuthenticationUpdate struct {
	Authentication_option map[string]interface{} `json:"authentication_option,omitempty"`
	Recording_authentication bool `json:"recording_authentication,omitempty"`
}

// WebinarPanelistList represents the WebinarPanelistList schema from the OpenAPI specification
type WebinarPanelistList struct {
	Panelists []interface{} `json:"panelists,omitempty"` // List of panelist objects.
	Total_records int `json:"total_records,omitempty"` // Total records.
}

// UserSettingsTSP represents the UserSettingsTSP schema from the OpenAPI specification
type UserSettingsTSP struct {
	Call_out bool `json:"call_out,omitempty"` // Call Out
	Call_out_countries []interface{} `json:"call_out_countries,omitempty"` // Call Out Countries/Regions
	Show_international_numbers_link bool `json:"show_international_numbers_link,omitempty"` // Show international numbers link on the invitation email
}

// UserList represents the UserList schema from the OpenAPI specification
type UserList struct {
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_number int `json:"page_number,omitempty"` // **Deprecated**: This field has been deprecated. Please use the "next_page_token" field for pagination instead of this field. The page number of the current results.
	Page_size int `json:"page_size,omitempty"` // The number of records returned with a single API call.
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
	Users []interface{} `json:"users,omitempty"` // List of user objects.
}

// MeetingRegistrant represents the MeetingRegistrant schema from the OpenAPI specification
type MeetingRegistrant struct {
	First_name string `json:"first_name"` // Registrant's first name.
	Purchasing_time_frame string `json:"purchasing_time_frame,omitempty"` // This field can be included to gauge interest of webinar attendees towards buying your product or service. Purchasing Time Frame:<br>`Within a month`<br>`1-3 months`<br>`4-6 months`<br>`More than 6 months`<br>`No timeframe`
	No_of_employees string `json:"no_of_employees,omitempty"` // Number of Employees:<br>`1-20`<br>`21-50`<br>`51-100`<br>`101-500`<br>`500-1,000`<br>`1,001-5,000`<br>`5,001-10,000`<br>`More than 10,000`
	Zip string `json:"zip,omitempty"` // Registrant's Zip/Postal Code.
	Email string `json:"email"` // A valid email address of the registrant.
	Comments string `json:"comments,omitempty"` // A field that allows registrants to provide any questions or comments that they might have.
	State string `json:"state,omitempty"` // Registrant's State/Province.
	Job_title string `json:"job_title,omitempty"` // Registrant's job title.
	Org string `json:"org,omitempty"` // Registrant's Organization.
	Industry string `json:"industry,omitempty"` // Registrant's Industry.
	Phone string `json:"phone,omitempty"` // Registrant's Phone number.
	Address string `json:"address,omitempty"` // Registrant's address.
	Last_name string `json:"last_name,omitempty"` // Registrant's last name.
	City string `json:"city,omitempty"` // Registrant's city.
	Custom_questions []map[string]interface{} `json:"custom_questions,omitempty"` // Custom questions.
	Role_in_purchase_process string `json:"role_in_purchase_process,omitempty"` // Role in Purchase Process:<br>`Decision Maker`<br>`Evaluator/Recommender`<br>`Influencer`<br>`Not involved`
	Country string `json:"country,omitempty"` // Registrant's country. The value of this field must be in two-letter abbreviated form and must match the ID field provided in the [Countries](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#countries) table.
}

// WebinarInfo represents the WebinarInfo schema from the OpenAPI specification
type WebinarInfo struct {
	Start_url string `json:"start_url,omitempty"` // <br><aside>The <code>start_url</code> of a Webinar is a URL using which a host or an alternative host can start the Webinar. This URL should only be used by the host of the meeting and should not be shared with anyone other than the host of the Webinar. The expiration time for the <code>start_url</code> field listed in the response of [Create a Webinar API](https://marketplace.zoom.us/docs/api-reference/zoom-api/webinars/webinarcreate) is two hours for all regular users. 	 For users created using the <code>custCreate</code> option via the [Create Users](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usercreate) API, the expiration time of the <code>start_url</code> field is 90 days. 	 For security reasons, to retrieve the latest value for the <code>start_url</code> field programmatically (after expiry), you must call the [Retrieve a Webinar API](https://marketplace.zoom.us/docs/api-reference/zoom-api/webinars/webinar) and refer to the value of the <code>start_url</code> field in the response.</aside><br><br><br>
	TypeField int `json:"type,omitempty"` // Webinar Types:<br>`5` - Webinar.<br>`6` - Recurring webinar with no fixed time.<br>`9` - Recurring webinar with a fixed time.
	Recurrence map[string]interface{} `json:"recurrence,omitempty"` // Recurrence object. Use this object only for a webinar of type `9` i.e., a recurring webinar with fixed time.
	Duration int `json:"duration,omitempty"` // Webinar duration.
	Join_url string `json:"join_url,omitempty"` // URL to join the Webinar. This URL should only be shared with the users who should be invited to the Webinar.
	Occurrences []map[string]interface{} `json:"occurrences,omitempty"` // Array of occurrence objects.
	Settings map[string]interface{} `json:"settings,omitempty"` // Webinar settings.
	Start_time string `json:"start_time,omitempty"` // Webinar start time in GMT/UTC.
	Agenda string `json:"agenda,omitempty"` // Webinar agenda.
	Password string `json:"password,omitempty"` // Webinar passcode. If "Require a passcode when scheduling new meetings" setting has been **enabled** **and** [locked](https://support.zoom.us/hc/en-us/articles/115005269866-Using-Tiered-Settings#locked) for the user, the passcode field will be autogenerated for the Webinar in the response even if it is not provided in the API request. <br><br> **Note:** If the account owner or the admin has configured [minimum passcode requirement settings](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords#h_a427384b-e383-4f80-864d-794bf0a37604), the passcode value provided here must meet those requirements. <br><br>If the requirements are enabled, you can view those requirements by calling either the [Get User Settings API](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usersettings) or the [Get Account Settings](https://marketplace.zoom.us/docs/api-reference/zoom-api/accounts/accountsettings) API.
	Timezone string `json:"timezone,omitempty"` // Time zone to format start_time.
	Topic string `json:"topic,omitempty"` // Webinar topic.
	Tracking_fields []map[string]interface{} `json:"tracking_fields,omitempty"` // Tracking fields
	Created_at string `json:"created_at,omitempty"` // Create time.
}

// SessionUpdate represents the SessionUpdate schema from the OpenAPI specification
type SessionUpdate struct {
	Agenda string `json:"agenda,omitempty"` // Meeting description.
	Duration int `json:"duration,omitempty"` // Meeting duration (minutes). Used for scheduled meetings only.
	Timezone string `json:"timezone,omitempty"` // Time zone to format start_time. For example, "America/Los_Angeles". For scheduled meetings only. Please reference our [time zone](#timezones) list for supported time zones and their formats.
	Password string `json:"password,omitempty"` // Meeting passcode. Passcode may only contain the following characters: [a-z A-Z 0-9 @ - _ *] and can have a maximum of 10 characters. **Note:** If the account owner or the admin has configured [minimum passcode requirement settings](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords#h_a427384b-e383-4f80-864d-794bf0a37604), the passcode value provided here must meet those requirements. <br><br>If the requirements are enabled, you can view those requirements by calling either the [Get User Settings API](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usersettings) or the [Get Account Settings](https://marketplace.zoom.us/docs/api-reference/zoom-api/accounts/accountsettings) API.
	Settings interface{} `json:"settings,omitempty"`
	Start_time string `json:"start_time,omitempty"` // Meeting start time. When using a format like "yyyy-MM-dd'T'HH:mm:ss'Z'", always use GMT time. When using a format like "yyyy-MM-dd'T'HH:mm:ss", you should use local time and specify the time zone. Only used for scheduled meetings and recurring meetings with a fixed time.
	Topic string `json:"topic,omitempty"` // Meeting topic.
	Template_id string `json:"template_id,omitempty"` // Unique identifier of the meeting template. Use this field if you would like to [schedule the meeting from a meeting template](https://support.zoom.us/hc/en-us/articles/360036559151-Meeting-templates#h_86f06cff-0852-4998-81c5-c83663c176fb). You can retrieve the value of this field by calling the [List meeting templates]() API.
	Tracking_fields []map[string]interface{} `json:"tracking_fields,omitempty"` // Tracking fields
	Recurrence map[string]interface{} `json:"recurrence,omitempty"` // Recurrence object. Use this object only for a meeting with type `8` i.e., a recurring meeting with fixed time.
	TypeField int `json:"type,omitempty"` // Meeting Types:<br>`1` - Instant meeting.<br>`2` - Scheduled meeting.<br>`3` - Recurring meeting with no fixed time.<br>`8` - Recurring meeting with a fixed time.
}

// WebinarMetric represents the WebinarMetric schema from the OpenAPI specification
type WebinarMetric struct {
	Has_voip bool `json:"has_voip,omitempty"` // Indicates whether or not VoIP was used for the Webinar.
	Dept string `json:"dept,omitempty"` // Department of the host.
	Has_recording bool `json:"has_recording,omitempty"` // Indicates whether or not recording was used for the Webinar.
	Start_time string `json:"start_time,omitempty"` // Webinar start time.
	Custom_keys []map[string]interface{} `json:"custom_keys,omitempty"` // Custom keys and values assigned to the Webinar.
	Has_screen_share bool `json:"has_screen_share,omitempty"` // Indicates whether or not screen sharing was used for the Webinar.
	Participants int `json:"participants,omitempty"` // Webinar participant count.
	Has_sip bool `json:"has_sip,omitempty"` // Indicates whether or not SIP was used for the Webinar.
	Duration string `json:"duration,omitempty"` // Webinar duration, formatted as hh:mm:ss, for example: `10:00` for ten minutes.
	End_time string `json:"end_time,omitempty"` // Webinar end time.
	Host string `json:"host,omitempty"` // User display name.
	Uuid string `json:"uuid,omitempty"` // Webinar UUID.
	Topic string `json:"topic,omitempty"` // Webinar topic.
	Has_3rd_party_audio bool `json:"has_3rd_party_audio,omitempty"` // Indicates whether or not TSP was used for the Webinar.
	Has_pstn bool `json:"has_pstn,omitempty"` // Indicates whether or not PSTN was used for the Webinar.
	Id int64 `json:"id,omitempty"` // Webinar ID in "**long**" format(represented as int64 data type in JSON), also known as the webinar number.
	Email string `json:"email,omitempty"` // User email.
	Has_video bool `json:"has_video,omitempty"` // Indicates whether or not video was used for the Webinar.
	User_type string `json:"user_type,omitempty"` // User type.
}

// Authenticationusersettings represents the Authenticationusersettings schema from the OpenAPI specification
type Authenticationusersettings struct {
}

// BillingContactRequired represents the BillingContactRequired schema from the OpenAPI specification
type BillingContactRequired struct {
	City string `json:"city"` // Billing Contact's city.
	Email string `json:"email"` // Billing Contact's email address.
	Apt string `json:"apt,omitempty"` // Billing Contact's apartment/suite.
	Country string `json:"country"` // Billing Contact's Country [ID](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#countries) in abbreviated format.
	First_name string `json:"first_name"` // Billing Contact's first name.
	Last_name string `json:"last_name"` // Billing Contact's last name.
	Address string `json:"address"` // Billing Contact's address.
	Phone_number string `json:"phone_number"` // Billing Contact's phone number.
	State string `json:"state"` // Billing Contact's state.
	Zip string `json:"zip"` // Billing Contact's zip/postal code.
}

// RoleMembersAdd represents the RoleMembersAdd schema from the OpenAPI specification
type RoleMembersAdd struct {
	Members []interface{} `json:"members,omitempty"` // List of Role's members
}

// UserSettingsTelephony represents the UserSettingsTelephony schema from the OpenAPI specification
type UserSettingsTelephony struct {
	Telephony_regions map[string]interface{} `json:"telephony_regions,omitempty"` // Indicates where most of the participants call into or call from duriing a meeting.
	Third_party_audio bool `json:"third_party_audio,omitempty"` // Third party audio conference.
	Audio_conference_info string `json:"audio_conference_info,omitempty"` // Third party audio conference info.
	Show_international_numbers_link bool `json:"show_international_numbers_link,omitempty"` // Show the international numbers link on the invitation email.
}

// WebinarRegistrant represents the WebinarRegistrant schema from the OpenAPI specification
type WebinarRegistrant struct {
	Comments string `json:"comments,omitempty"` // A field that allows registrants to provide any questions or comments that they might have.
	No_of_employees string `json:"no_of_employees,omitempty"` // Number of Employees:<br>`1-20`<br>`21-50`<br>`51-100`<br>`101-500`<br>`500-1,000`<br>`1,001-5,000`<br>`5,001-10,000`<br>`More than 10,000`
	Last_name string `json:"last_name,omitempty"` // Registrant's last name.
	Email string `json:"email"` // A valid email address of the registrant.
	Custom_questions []map[string]interface{} `json:"custom_questions,omitempty"` // Custom questions.
	First_name string `json:"first_name"` // Registrant's first name.
	Job_title string `json:"job_title,omitempty"` // Registrant's job title.
	Org string `json:"org,omitempty"` // Registrant's Organization.
	Purchasing_time_frame string `json:"purchasing_time_frame,omitempty"` // This field can be included to gauge interest of webinar attendees towards buying your product or service. Purchasing Time Frame:<br>`Within a month`<br>`1-3 months`<br>`4-6 months`<br>`More than 6 months`<br>`No timeframe`
	Address string `json:"address,omitempty"` // Registrant's address.
	Role_in_purchase_process string `json:"role_in_purchase_process,omitempty"` // Role in Purchase Process:<br>`Decision Maker`<br>`Evaluator/Recommender`<br>`Influencer`<br>`Not involved`
	Country string `json:"country,omitempty"` // Registrant's country. The value of this field must be in two-letter abbreviated form and must match the ID field provided in the [Countries](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#countries) table.
	State string `json:"state,omitempty"` // Registrant's State/Province.
	City string `json:"city,omitempty"` // Registrant's city.
	Industry string `json:"industry,omitempty"` // Registrant's Industry.
	Zip string `json:"zip,omitempty"` // Registrant's Zip/Postal Code.
	Phone string `json:"phone,omitempty"` // Registrant's Phone number.
}

// AccountSettingsEmailNotification represents the AccountSettingsEmailNotification schema from the OpenAPI specification
type AccountSettingsEmailNotification struct {
	Jbh_reminder bool `json:"jbh_reminder,omitempty"` // Notify the host when participants join the meeting before them.
	Low_host_count_reminder bool `json:"low_host_count_reminder,omitempty"` // Notify user when host licenses are running low.
	Schedule_for_reminder bool `json:"schedule_for_reminder,omitempty"` // Notify the host there is a meeting is scheduled, rescheduled, or cancelled.
	Alternative_host_reminder bool `json:"alternative_host_reminder,omitempty"` // Notify when an alternative host is set or removed from a meeting.
	Cancel_meeting_reminder bool `json:"cancel_meeting_reminder,omitempty"` // Notify the host and participants when a meeting is cancelled.
	Cloud_recording_avaliable_reminder bool `json:"cloud_recording_avaliable_reminder,omitempty"` // Notify host when cloud recording is available.
}

// QOSParticipantList represents the QOSParticipantList schema from the OpenAPI specification
type QOSParticipantList struct {
	Page_count int64 `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_size int `json:"page_size,omitempty"` // The number of items per page.
	Total_records int64 `json:"total_records,omitempty"` // The number of all records available across pages.
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceed the current page size. The expiration period for this token is 15 minutes.
	Participants []map[string]interface{} `json:"participants,omitempty"` // Array of user objects.
}

// GroupMemberList represents the GroupMemberList schema from the OpenAPI specification
type GroupMemberList struct {
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_number int `json:"page_number,omitempty"` // **Deprecated**: This field has been deprecated. Please use the "next_page_token" field for pagination instead of this field. The page number of the current results.
	Page_size int `json:"page_size,omitempty"` // The number of records returned with a single API call.
	Members []map[string]interface{} `json:"members,omitempty"` // List of Group member objects.
}

// MeetingInfoGet represents the MeetingInfoGet schema from the OpenAPI specification
type MeetingInfoGet struct {
	H323_password string `json:"h323_password,omitempty"` // H.323/SIP room system passcode.
	Occurrences []map[string]interface{} `json:"occurrences,omitempty"` // Array of occurrence objects.
	Status string `json:"status,omitempty"` // Meeting status
	Topic string `json:"topic,omitempty"` // Meeting topic.
	Duration int `json:"duration,omitempty"` // Meeting duration.
	Pmi int64 `json:"pmi,omitempty"` // Personal Meeting Id. Only used for scheduled meetings and recurring meetings with no fixed time.
	Recurrence map[string]interface{} `json:"recurrence,omitempty"` // Recurrence object. Use this object only for a meeting with type `8` i.e., a recurring meeting with fixed time.
	Tracking_fields []map[string]interface{} `json:"tracking_fields,omitempty"` // Tracking fields
	Password string `json:"password,omitempty"` // Meeting passcode.
	Created_at string `json:"created_at,omitempty"` // Time of creation.
	Join_url string `json:"join_url,omitempty"` // URL for participants to join the meeting. This URL should only be shared with users that you would like to invite for the meeting.
	Timezone string `json:"timezone,omitempty"` // Timezone to format the meeting start time on the .
	Encrypted_password string `json:"encrypted_password,omitempty"` // Encrypted passcode for third party endpoints (H323/SIP).
	Start_time string `json:"start_time,omitempty"` // Meeting start time in GMT/UTC. Start time will not be returned if the meeting is an **instant** meeting.
	Start_url string `json:"start_url,omitempty"` // <br><aside>The <code>start_url</code> of a Meeting is a URL using which a host or an alternative host can start the Meeting. The expiration time for the <code>start_url</code> field listed in the response of [Create a Meeting API](https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingcreate) is two hours for all regular users. 	 For users created using the <code>custCreate</code> option via the [Create Users](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usercreate) API, the expiration time of the <code>start_url</code> field is 90 days. 	 For security reasons, to retrieve the updated value for the <code>start_url</code> field programmatically (after the expiry time), you must call the [Retrieve a Meeting API](https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meeting) and refer to the value of the <code>start_url</code> field in the response.</aside><br>This URL should only be used by the host of the meeting and **should not be shared with anyone other than the host** of the meeting as anyone with this URL will be able to login to the Zoom Client as the host of the meeting.
	Settings map[string]interface{} `json:"settings,omitempty"` // Meeting settings.
	Agenda string `json:"agenda,omitempty"` // Meeting description
	TypeField int `json:"type,omitempty"` // Meeting Types:<br>`1` - Instant meeting.<br>`2` - Scheduled meeting.<br>`3` - Recurring meeting with no fixed time.<br>`4` - PMI Meeting<br> `8` - Recurring meeting with a fixed time.
}

// WebinarInstances represents the WebinarInstances schema from the OpenAPI specification
type WebinarInstances struct {
	Webinars []map[string]interface{} `json:"webinars,omitempty"` // List of ended webinar instances.
}

// AccountSettingsUpdateTelephony represents the AccountSettingsUpdateTelephony schema from the OpenAPI specification
type AccountSettingsUpdateTelephony struct {
	Telephony_regions map[string]interface{} `json:"telephony_regions,omitempty"` // Indicates where most of the participants call into or call from duriing a meeting.
	Third_party_audio bool `json:"third_party_audio,omitempty"` // Users can join the meeting using the existing third party audio configuration.
	Audio_conference_info string `json:"audio_conference_info,omitempty"` // Third party audio conference info.
}

// AccountSettingsAuthentication represents the AccountSettingsAuthentication schema from the OpenAPI specification
type AccountSettingsAuthentication struct {
}

// MeetingSecuritySettings represents the MeetingSecuritySettings schema from the OpenAPI specification
type MeetingSecuritySettings struct {
	Meeting_security map[string]interface{} `json:"meeting_security,omitempty"`
}

// RegistrantList represents the RegistrantList schema from the OpenAPI specification
type RegistrantList struct {
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_number int `json:"page_number,omitempty"` // **Deprecated**: This field has been deprecated. Please use the "next_page_token" field for pagination instead of this field. The page number of the current results.
	Page_size int `json:"page_size,omitempty"` // The number of records returned with a single API call.
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Registrants []interface{} `json:"registrants,omitempty"` // List of registrant objects.
}

// RecordingRegistrantList represents the RecordingRegistrantList schema from the OpenAPI specification
type RecordingRegistrantList struct {
	Page_number int `json:"page_number,omitempty"` // **Deprecated**: This field has been deprecated. Please use the "next_page_token" field for pagination instead of this field. The page number of the current results.
	Page_size int `json:"page_size,omitempty"` // The number of records returned with a single API call.
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Registrants []interface{} `json:"registrants,omitempty"` // List of Registrant objects
}

// UserSettingsScheduleMeeting represents the UserSettingsScheduleMeeting schema from the OpenAPI specification
type UserSettingsScheduleMeeting struct {
	Audio_type string `json:"audio_type,omitempty"` // Determine how participants can join the audio portion of the meeting:<br>`both` - Telephony and VoIP.<br>`telephony` - Audio PSTN telephony only.<br>`voip` - VoIP only.<br>`thirdParty` - Third party audio conference.
	Require_password_for_instant_meetings bool `json:"require_password_for_instant_meetings,omitempty"` // Require a passcode for instant meetings. If you use PMI for your instant meetings, this option will be disabled. This setting is always enabled for free accounts and Pro accounts with a single host and cannot be modified for these accounts.
	Participants_video bool `json:"participants_video,omitempty"` // Start meetings with participants video on.
	Use_pmi_for_scheduled_meetings bool `json:"use_pmi_for_scheduled_meetings,omitempty"` // Use Personal Meeting ID (PMI) when scheduling a meeting
	Require_password_for_pmi_meetings string `json:"require_password_for_pmi_meetings,omitempty"` // Require a passcode for Personal Meeting ID (PMI). This setting is always enabled for free accounts and Pro accounts with a single host and cannot be modified for these accounts.
	Require_password_for_scheduling_new_meetings bool `json:"require_password_for_scheduling_new_meetings,omitempty"` // Require a passcode when scheduling new meetings.This setting is always enabled for free accounts and Pro accounts with a single host and cannot be modified for these accounts.
	Use_pmi_for_instant_meetings bool `json:"use_pmi_for_instant_meetings,omitempty"` // Use Personal Meeting ID (PMI) when starting an instant meeting
	Host_video bool `json:"host_video,omitempty"` // Start meetings with host video on.
	Force_pmi_jbh_password bool `json:"force_pmi_jbh_password,omitempty"` // Require a passcode for personal meetings if attendees can join before host.
	Embed_password_in_join_link bool `json:"embed_password_in_join_link,omitempty"` // If the value is set to `true`, the meeting passcode will be encrypted and included in the join meeting link to allow participants to join with just one click without having to enter the passcode.
	Personal_meeting bool `json:"personal_meeting,omitempty"` // Personal Meeting Setting.<br><br> `true`: Indicates that the **"Enable Personal Meeting ID"** setting is turned on. Users can choose to use personal meeting ID for their meetings. <br><br> `false`: Indicates that the **"Enable Personal Meeting ID"** setting is [turned off](https://support.zoom.us/hc/en-us/articles/201362843-Personal-meeting-ID-PMI-and-personal-link#h_aa0335c8-3b06-41bc-bc1f-a8b84ef17f2a). If this setting is disabled, meetings that were scheduled with PMI will be invalid. Scheduled meetings will need to be manually updated. For Zoom Phone only:If a user has been assigned a desk phone, **"Elevate to Zoom Meeting"** on desk phone will be disabled.
	Pmi_password string `json:"pmi_password,omitempty"` // PMI passcode
	Require_password_for_scheduled_meetings bool `json:"require_password_for_scheduled_meetings,omitempty"` // Require a passcode for meetings which have already been scheduled
	Meeting_password_requirement map[string]interface{} `json:"meeting_password_requirement,omitempty"` // Account wide meeting/webinar [passcode requirements](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords#h_a427384b-e383-4f80-864d-794bf0a37604).
	Default_password_for_scheduled_meetings string `json:"default_password_for_scheduled_meetings,omitempty"` // Passcode for already scheduled meetings
	Join_before_host bool `json:"join_before_host,omitempty"` // Join the meeting before host arrives.
	Pstn_password_protected bool `json:"pstn_password_protected,omitempty"` // Generate and require passcode for participants joining by phone.
}

// MeetingLiveStream represents the MeetingLiveStream schema from the OpenAPI specification
type MeetingLiveStream struct {
	Stream_url string `json:"stream_url"` // Streaming URL.
	Page_url string `json:"page_url,omitempty"` // The livestream page URL.
	Stream_key string `json:"stream_key"` // Stream name and key.
}

// AccountPlan represents the AccountPlan schema from the OpenAPI specification
type AccountPlan struct {
	TypeField string `json:"type,omitempty"` // Account <a href="https://marketplace.zoom.us/docs/api-reference/other-references/plans">plan type.</a>
	Hosts int `json:"hosts,omitempty"` // Account plan number of hosts.
}

// Poll represents the Poll schema from the OpenAPI specification
type Poll struct {
	Title string `json:"title,omitempty"` // Title for the poll.
	Questions []map[string]interface{} `json:"questions,omitempty"` // Array of Polls
}

// GroupUserSettingsAuthentication represents the GroupUserSettingsAuthentication schema from the OpenAPI specification
type GroupUserSettingsAuthentication struct {
}

// Pagination represents the Pagination schema from the OpenAPI specification
type Pagination struct {
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_number int `json:"page_number,omitempty"` // **Deprecated**: This field has been deprecated. Please use the "next_page_token" field for pagination instead of this field. The page number of the current results.
	Page_size int `json:"page_size,omitempty"` // The number of records returned with a single API call.
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
}

// UserSettingsFeature represents the UserSettingsFeature schema from the OpenAPI specification
type UserSettingsFeature struct {
	Meeting_capacity int `json:"meeting_capacity,omitempty"` // User’s meeting capacity.
	Webinar bool `json:"webinar,omitempty"` // Webinar feature.
	Webinar_capacity int `json:"webinar_capacity,omitempty"` // Webinar capacity: can be 100, 500, 1000, 3000, 5000 or 10000, depending on if the user has a webinar capacity plan subscription or not.
	Zoom_phone bool `json:"zoom_phone,omitempty"` // Zoom phone feature.
	Cn_meeting bool `json:"cn_meeting,omitempty"` // Host meeting in China.
	In_meeting bool `json:"in_meeting,omitempty"` // Host meeting in India.
	Large_meeting bool `json:"large_meeting,omitempty"` // Large meeting feature.
	Large_meeting_capacity int `json:"large_meeting_capacity,omitempty"` // Large meeting capacity: can be 500 or 1000, depending on if the user has a large meeting capacity plan subscription or not.
}

// CloudArchivedFiles represents the CloudArchivedFiles schema from the OpenAPI specification
type CloudArchivedFiles struct {
	Host_id string `json:"host_id"` // The ID of the user who set as the host of the meeting.
	Recording_count int `json:"recording_count"` // Number of recording files returned in the response of this API call.
	Start_time string `json:"start_time"` // Meeting start time.
	Timezone string `json:"timezone"` // Timezone to format the meeting start time.
	Topic string `json:"topic"` // The meeting topic.
	Total_size int `json:"total_size"` // Total size of the archive.
	Uuid string `json:"uuid"` // The Unique Meeting ID. Each meeting instance will generate its own Meeting UUID.
	Archive_files []map[string]interface{} `json:"archive_files"` // An explanation about the purpose of this instance.
	Id int `json:"id"` // The Meeting ID, also known as the meeting number in long (int64) format.
	TypeField int `json:"type"` // The meeting type:<br>`1` - Instant meeting.<br>`2` - Scheduled meeting.<br>`3` - Recurring meeting with no fixed time.<br>`8` - Recurring meeting with fixed time.
	Duration int `json:"duration"` // The duration.
}

// WebinarRegistrantList represents the WebinarRegistrantList schema from the OpenAPI specification
type WebinarRegistrantList struct {
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_number int `json:"page_number,omitempty"` // **Deprecated**: This field has been deprecated. Please use the "next_page_token" field for pagination instead of this field. The page number of the current results.
	Page_size int `json:"page_size,omitempty"` // The number of records returned with a single API call.
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
	Registrants []interface{} `json:"registrants,omitempty"` // List of registrant objects.
}

// AccountSettingsSecurity represents the AccountSettingsSecurity schema from the OpenAPI specification
type AccountSettingsSecurity struct {
	Admin_change_name_pic bool `json:"admin_change_name_pic,omitempty"` // Only account administrators can change a user's username and picture.
	Sign_again_period_for_inactivity_on_web int `json:"sign_again_period_for_inactivity_on_web,omitempty"` // Settings for User Sign In interval requirements after a period of inactivity. If enabled, this setting forces automatic logout of users in Zoom Web Portal after a set amount of time. <br> If this setting is disabled, the value of this field will be `0`. If the setting is enabled, the value of this field will indicate the **period of inactivity** in minutes after which, an inactive user will be automatically logged out of the Zoom Web Portal. The value for the period of inactivity can be one of the following:<br> `5`: 5 minutes<br> `10`: 10 minutes<br> `15`: 15 minutes<br> `30`: 30 minutes<br> `60`: 60 minutes<br> `120`: 120 minutes
	Sign_in_with_two_factor_auth_groups []string `json:"sign_in_with_two_factor_auth_groups,omitempty"` // This field contains group IDs of groups that have 2FA enabled. This field is only returned if the value of `sign_in_with_two_factor_auth` is `group`
	Sign_in_with_two_factor_auth_roles []string `json:"sign_in_with_two_factor_auth_roles,omitempty"` // This field contains role IDs of roles that have 2FA enabled. This field is only returned if the value of `sign_in_with_two_factor_auth` is `role`.
	Hide_billing_info bool `json:"hide_billing_info,omitempty"` // Hide billing information.
	Sign_in_with_two_factor_auth string `json:"sign_in_with_two_factor_auth,omitempty"` // Settings for 2FA( [two factor authentication](https://support.zoom.us/hc/en-us/articles/360038247071) ). The value can be one of the following: `all`: Two factor authentication will be enabled for all users in the account.<br> `none`: Two factor authentication is disabled.<br> `group`: Two factor authentication will be enabled for users belonging to specific groups. If 2FA is enabled for certain groups, the group IDs of the group(s) will be provided in the `sign_in_with_two_factor_auth_groups` field.<br> `role`: Two factor authentication will be enabled only for users assigned with specific roles in the account. If 2FA is enabled for specific roles, the role IDs will be provided in the `sign_in_with_two_factor_auth_roles` field.
	Import_photos_from_devices bool `json:"import_photos_from_devices,omitempty"` // Allow users to import photos from a photo library on a device.
	Password_requirement map[string]interface{} `json:"password_requirement,omitempty"` // This object refers to the [enhanced password rules](https://support.zoom.us/hc/en-us/articles/360034675592-Advanced-security-settings#h_bf8a25f6-9a66-447a-befd-f02ed3404f89) that allows Zoom account admins and owners to apply extra requiremets to the users' Zoom login password.
	Sign_again_period_for_inactivity_on_client int `json:"sign_again_period_for_inactivity_on_client,omitempty"` // Settings for User Sign In interval requirements after a period of inactivity. If enabled, this setting forces automatic logout of users in Zoom Client app after a set amount of time. <br> If this setting is disabled, the value of this field will be `0`. If the setting is enabled, the value of this field will indicate the **period of inactivity** in minutes after which, an inactive user will be automatically logged out of the Zoom Client. The value for the period of inactivity can be one of the following:<br> `5`: 5 minutes<br> `10`: 10 minutes<br> `15`: 15 minutes<br> `30`: 30 minutes<br> `45`: 45 minutes<br> `60`: 60 minutes<br> `90`: 90 minutes<br> `120`: 120 minutes
}

// MeetingInvitation represents the MeetingInvitation schema from the OpenAPI specification
type MeetingInvitation struct {
	Invitation string `json:"invitation,omitempty"` // Meeting invitation.
}

// AccountSettingsScheduleMeeting represents the AccountSettingsScheduleMeeting schema from the OpenAPI specification
type AccountSettingsScheduleMeeting struct {
	Personal_meeting bool `json:"personal_meeting,omitempty"` // Personal Meeting Setting.<br><br> `true`: Indicates that the **"Enable Personal Meeting ID"** setting is turned on. Users can choose to use personal meeting ID for their meetings. <br><br> `false`: Indicates that the **"Enable Personal Meeting ID"** setting is [turned off](https://support.zoom.us/hc/en-us/articles/201362843-Personal-meeting-ID-PMI-and-personal-link#h_aa0335c8-3b06-41bc-bc1f-a8b84ef17f2a). If this setting is disabled, meetings that were scheduled with PMI will be invalid. Scheduled meetings will need to be manually updated. For Zoom Phone only:If a user has been assigned a desk phone, **"Elevate to Zoom Meeting"** on desk phone will be disabled.
	Require_password_for_instant_meetings bool `json:"require_password_for_instant_meetings,omitempty"` // Require a password for instant meetings. If you use PMI for your instant meetings, this option will be disabled. This setting is always enabled for free accounts and Pro accounts with a single host and cannot be modified for these accounts.
	Require_password_for_scheduled_meetings bool `json:"require_password_for_scheduled_meetings,omitempty"` // Require a password for meetings which have already been scheduled
	Enforce_login_domains string `json:"enforce_login_domains,omitempty"` // Only signed in users with a specified domain can join the meeting.
	Enforce_login bool `json:"enforce_login,omitempty"` // Only Zoom users who are signed in can join meetings.
	Join_before_host bool `json:"join_before_host,omitempty"` // Allow participants to join the meeting before the host arrives.
	Use_pmi_for_instant_meetings bool `json:"use_pmi_for_instant_meetings,omitempty"` // Use Personal Meeting ID (PMI) when starting an instant meeting
	Meeting_password_requirement map[string]interface{} `json:"meeting_password_requirement,omitempty"` // Account wide meeting/webinar [password requirements](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords#h_a427384b-e383-4f80-864d-794bf0a37604).
	Participant_video bool `json:"participant_video,omitempty"` // Start meetings with the participant video on. Participants can change this setting during the meeting.
	Host_video bool `json:"host_video,omitempty"` // Start meetings with the host video on.
	Enforce_login_with_domains bool `json:"enforce_login_with_domains,omitempty"` // Only signed in users with a specific domain can join meetings.
	Force_pmi_jbh_password bool `json:"force_pmi_jbh_password,omitempty"` // Require a password for Personal Meetings if attendees can join before host.
	Not_store_meeting_topic bool `json:"not_store_meeting_topic,omitempty"` // Always display "Zoom Meeting" as the meeting topic.
	Audio_type string `json:"audio_type,omitempty"` // Determine how participants can join the audio portion of the meeting.<br>`both` - Telephony and VoIP.<br>`telephony` - Audio PSTN telephony only.<br>`voip` - VoIP only.<br>`thirdParty` - 3rd party audio conference.
	Use_pmi_for_scheduled_meetings bool `json:"use_pmi_for_scheduled_meetings,omitempty"` // Use Personal Meeting ID (PMI) when scheduling a meeting
	Require_password_for_scheduling_new_meetings bool `json:"require_password_for_scheduling_new_meetings,omitempty"` // Require a password when scheduling new meetings. This setting applies for regular meetings that do not use PMI. If enabled, a password will be generated while a host schedules a new meeting and participants will be required to enter the password before they can join the meeting. This setting is always enabled for free accounts and Pro accounts with a single host and cannot be modified for these accounts.
	Require_password_for_pmi_meetings string `json:"require_password_for_pmi_meetings,omitempty"` // Require a password for a meeting held using Personal Meeting ID (PMI) This setting is always enabled for free accounts and Pro accounts with a single host and cannot be modified for these accounts.
}

// RecordingMeeting represents the RecordingMeeting schema from the OpenAPI specification
type RecordingMeeting struct {
	Account_id string `json:"account_id,omitempty"` // Unique Identifier of the user account.
	Duration int `json:"duration,omitempty"` // Meeting duration.
	Host_id string `json:"host_id,omitempty"` // ID of the user set as host of meeting.
	Id string `json:"id,omitempty"` // Meeting ID - also known as the meeting number.
	Recording_count string `json:"recording_count,omitempty"` // Number of recording files returned in the response of this API call.
	TypeField string `json:"type,omitempty"` // Type of the meeting that was recorded. Meeting Types:<br>`1` - Instant meeting.<br>`2` - Scheduled meeting.<br>`3` - Recurring meeting with no fixed time.<br>`8` - Recurring meeting with fixed time.
	Topic string `json:"topic,omitempty"` // Meeting topic.
	Start_time string `json:"start_time,omitempty"` // The time at which the meeting started.
	Total_size string `json:"total_size,omitempty"` // Total size of the recording.
	Uuid string `json:"uuid,omitempty"` // Unique Meeting Identifier. Each instance of the meeting will have its own UUID.
	Recording_files []interface{} `json:"recording_files,omitempty"` // List of recording file.
}

// AccountPlans represents the AccountPlans schema from the OpenAPI specification
type AccountPlans struct {
	Plan_large_meeting []map[string]interface{} `json:"plan_large_meeting,omitempty"` // Additional large meeting Plans.
	Plan_phone map[string]interface{} `json:"plan_phone,omitempty"` // Phone Plan Object
	Plan_recording string `json:"plan_recording,omitempty"` // Additional cloud recording plan.
	Plan_room_connector map[string]interface{} `json:"plan_room_connector,omitempty"` // Account plan object.
	Plan_webinar []map[string]interface{} `json:"plan_webinar,omitempty"` // Additional webinar plans.
	Plan_zoom_rooms map[string]interface{} `json:"plan_zoom_rooms,omitempty"` // Account plan object.
	Plan_audio map[string]interface{} `json:"plan_audio,omitempty"` // Additional audio conferencing <a href="https://marketplace.zoom.us/docs/api-reference/other-references/plans#audio-conferencing-plans">plan type</a>.
	Plan_base map[string]interface{} `json:"plan_base"` // Account base plan object.
}

// Createwebinar represents the Createwebinar schema from the OpenAPI specification
type Createwebinar struct {
	Password string `json:"password,omitempty"` // Webinar passcode. Passcode may only contain the following characters: [a-z A-Z 0-9 @ - _ * !]. Max of 10 characters. If "Require a passcode when scheduling new meetings" setting has been **enabled** **and** [locked](https://support.zoom.us/hc/en-us/articles/115005269866-Using-Tiered-Settings#locked) for the user, the passcode field will be autogenerated for the Webinar in the response even if it is not provided in the API request. <br><br> **Note:** If the account owner or the admin has configured [minimum passcode requirement settings](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords#h_a427384b-e383-4f80-864d-794bf0a37604), the passcode value provided here must meet those requirements. <br><br>If the requirements are enabled, you can view those requirements by calling [Get Account Settings](https://marketplace.zoom.us/docs/api-reference/zoom-api/accounts/accountsettings) API.
	Recurrence map[string]interface{} `json:"recurrence,omitempty"` // Recurrence object. Use this object only for a webinar of type `9` i.e., a recurring webinar with fixed time.
	Settings map[string]interface{} `json:"settings,omitempty"` // Create Webinar settings.
	Start_time string `json:"start_time,omitempty"` // Webinar start time. We support two formats for `start_time` - local time and GMT.<br> To set time as GMT the format should be `yyyy-MM-dd`T`HH:mm:ssZ`. To set time using a specific timezone, use `yyyy-MM-dd`T`HH:mm:ss` format and specify the timezone [ID](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones) in the `timezone` field OR leave it blank and the timezone set on your Zoom account will be used. You can also set the time as UTC as the timezone field. The `start_time` should only be used for scheduled and / or recurring webinars with fixed time.
	Timezone string `json:"timezone,omitempty"` // Time zone to format start_time. For example, "America/Los_Angeles". For scheduled meetings only. Please reference our [timezone](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones) list for supported time zones and their formats.
	Topic string `json:"topic,omitempty"` // Webinar topic.
	Tracking_fields []map[string]interface{} `json:"tracking_fields,omitempty"` // Tracking fields
	TypeField int `json:"type,omitempty"` // Webinar Types:<br>`5` - Webinar.<br>`6` - Recurring webinar with no fixed time.<br>`9` - Recurring webinar with a fixed time.
	Agenda string `json:"agenda,omitempty"` // Webinar description.
	Duration int `json:"duration,omitempty"` // Webinar duration (minutes). Used for scheduled webinars only.
}

// WebinarPanelist represents the WebinarPanelist schema from the OpenAPI specification
type WebinarPanelist struct {
	Panelists []interface{} `json:"panelists,omitempty"` // List of panelist objects.
}

// Listmeetingmetrics represents the Listmeetingmetrics schema from the OpenAPI specification
type Listmeetingmetrics struct {
	Custom_keys []map[string]interface{} `json:"custom_keys,omitempty"` // Custom keys and values assigned to the meeting.
	Duration string `json:"duration,omitempty"` // Meeting duration. Formatted as hh:mm:ss, for example: `16:08` for 16 minutes and 8 seconds.
	Topic string `json:"topic,omitempty"` // Meeting topic.
	Uuid string `json:"uuid,omitempty"` // Meeting UUID. Please double encode your UUID when using it for API calls if the UUID begins with a '/'or contains '//' in it.
	Start_time string `json:"start_time,omitempty"` // Meeting start time.
	Dept string `json:"dept,omitempty"` // Department of the host.
	Has_pstn bool `json:"has_pstn,omitempty"` // Indicates whether or not the PSTN was used in the meeting.
	Email string `json:"email,omitempty"` // Email address of the host.
	Has_screen_share bool `json:"has_screen_share,omitempty"` // Indicates whether or not screenshare feature was used in the meeting.
	Has_sip bool `json:"has_sip,omitempty"` // Indicates whether or not someone joined the meeting using SIP.
	Has_video bool `json:"has_video,omitempty"` // Indicates whether or not video was used in the meeting.
	Participants int `json:"participants,omitempty"` // Meeting participant count.
	Has_3rd_party_audio bool `json:"has_3rd_party_audio,omitempty"` // Indicates whether or not [third party audio](https://support.zoom.us/hc/en-us/articles/202470795-3rd-Party-Audio-Conference) was used in the meeting.
	Has_recording bool `json:"has_recording,omitempty"` // Indicates whether or not the recording feature was used in the meeting.
	Has_voip bool `json:"has_voip,omitempty"` // Indicates whether or not VoIP was used in the meeting.
	Id int64 `json:"id,omitempty"` // [Meeting ID](https://support.zoom.us/hc/en-us/articles/201362373-What-is-a-Meeting-ID-): Unique identifier of the meeting in "**long**" format(represented as int64 data type in JSON), also known as the meeting number.
	Tracking_fields []map[string]interface{} `json:"tracking_fields,omitempty"` // Tracking fields and values assigned to the meeting.
	End_time string `json:"end_time,omitempty"` // Meeting end time.
	User_type string `json:"user_type,omitempty"` // License type of the user.
	Host string `json:"host,omitempty"` // Host display name.
	In_room_participants int `json:"in_room_participants,omitempty"` // The number of Zoom Room participants in the meeting.
}

// UserPermissions represents the UserPermissions schema from the OpenAPI specification
type UserPermissions struct {
	Permissions []string `json:"permissions,omitempty"` // List of user permissions.
}

// MeetingInfo represents the MeetingInfo schema from the OpenAPI specification
type MeetingInfo struct {
	Agenda string `json:"agenda,omitempty"` // Agenda
	Settings map[string]interface{} `json:"settings,omitempty"` // Meeting settings.
	Password string `json:"password,omitempty"` // Meeting password. Password may only contain the following characters: `[a-z A-Z 0-9 @ - _ * !]` If "Require a password when scheduling new meetings" setting has been **enabled** **and** [locked](https://support.zoom.us/hc/en-us/articles/115005269866-Using-Tiered-Settings#locked) for the user, the password field will be autogenerated in the response even if it is not provided in the API request.
	Start_time string `json:"start_time,omitempty"` // Meeting start date-time in UTC/GMT. Example: "2020-03-31T12:02:00Z"
	Join_url string `json:"join_url,omitempty"` // URL for participants to join the meeting. This URL should only be shared with users that you would like to invite for the meeting.
	H323_password string `json:"h323_password,omitempty"` // H.323/SIP room system password
	Pmi int64 `json:"pmi,omitempty"` // Personal Meeting Id. Only used for scheduled meetings and recurring meetings with no fixed time.
	TypeField int `json:"type,omitempty"` // Meeting Type
	Duration int `json:"duration,omitempty"` // Meeting duration.
	Occurrences []map[string]interface{} `json:"occurrences,omitempty"` // Array of occurrence objects.
	Start_url string `json:"start_url,omitempty"` // URL to start the meeting. This URL should only be used by the host of the meeting and **should not be shared with anyone other than the host** of the meeting as anyone with this URL will be able to login to the Zoom Client as the host of the meeting.
	Tracking_fields []map[string]interface{} `json:"tracking_fields,omitempty"` // Tracking fields
	Topic string `json:"topic,omitempty"` // Meeting topic
	Recurrence map[string]interface{} `json:"recurrence,omitempty"` // Recurrence object. Use this object only for a meeting with type `8` i.e., a recurring meeting with fixed time.
	Created_at string `json:"created_at,omitempty"` // The date and time at which this meeting was created.
	Timezone string `json:"timezone,omitempty"` // Timezone to format start_time
}

// MeetingList represents the MeetingList schema from the OpenAPI specification
type MeetingList struct {
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_number int `json:"page_number,omitempty"` // **Deprecated**: This field has been deprecated. Please use the "next_page_token" field for pagination instead of this field. The page number of the current results.
	Page_size int `json:"page_size,omitempty"` // The number of records returned with a single API call.
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Meetings []map[string]interface{} `json:"meetings,omitempty"` // List of Meeting objects.
}

// Webinar represents the Webinar schema from the OpenAPI specification
type Webinar struct {
	Agenda string `json:"agenda,omitempty"` // Webinar description.
	Recurrence map[string]interface{} `json:"recurrence,omitempty"` // Recurrence object. Use this object only for a webinar of type `9` i.e., a recurring webinar with fixed time.
	Timezone string `json:"timezone,omitempty"` // Time zone to format start_time. For example, "America/Los_Angeles". For scheduled meetings only. Please reference our [time zone](#timezones) list for supported time zones and their formats.
	Tracking_fields []map[string]interface{} `json:"tracking_fields,omitempty"` // Tracking fields
	Duration int `json:"duration,omitempty"` // Webinar duration (minutes). Used for scheduled webinars only.
	Password string `json:"password,omitempty"` // Webinar Passcode. Passcode may only contain the following characters: [a-z A-Z 0-9 @ - _ *]. Max of 10 characters.
	Settings map[string]interface{} `json:"settings,omitempty"` // Webinar settings.
	Topic string `json:"topic,omitempty"` // Webinar topic.
	Start_time string `json:"start_time,omitempty"` // Webinar start time. We support two formats for `start_time` - local time and GMT.<br> To set time as GMT the format should be `yyyy-MM-dd`T`HH:mm:ssZ`. To set time using a specific timezone, use `yyyy-MM-dd`T`HH:mm:ss` format and specify the timezone [ID](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones) in the `timezone` field OR leave it blank and the timezone set on your Zoom account will be used. You can also set the time as UTC as the timezone field. The `start_time` should only be used for scheduled and / or recurring webinars with fixed time.
	TypeField int `json:"type,omitempty"` // Webinar Types:<br>`5` - Webinar.<br>`6` - Recurring webinar with no fixed time.<br>`9` - Recurring webinar with a fixed time.
}

// PaginationToken4IMChat represents the PaginationToken4IMChat schema from the OpenAPI specification
type PaginationToken4IMChat struct {
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of the available result list exceeds the page size. The expiration period is 15 minutes.
	Page_size int `json:"page_size,omitempty"` // The amount of records returns within a single API call.
}

// PollList represents the PollList schema from the OpenAPI specification
type PollList struct {
	Polls []interface{} `json:"polls,omitempty"` // Array of Polls
	Total_records int `json:"total_records,omitempty"` // The number of all records available across pages
}

// SettingsUpdateTelephony represents the SettingsUpdateTelephony schema from the OpenAPI specification
type SettingsUpdateTelephony struct {
	Show_international_numbers_link bool `json:"show_international_numbers_link,omitempty"` // Show the international numbers link on the invitation email.
	Telephony_regions map[string]interface{} `json:"telephony_regions,omitempty"` // Indicates where most of the participants call into or call from duriing a meeting.
	Third_party_audio bool `json:"third_party_audio,omitempty"` // Third party audio conference.
	Audio_conference_info string `json:"audio_conference_info,omitempty"` // Third party audio conference info.
}

// User represents the User schema from the OpenAPI specification
type User struct {
	Pmi int64 `json:"pmi,omitempty"` // Personal meeting ID.
	TypeField int `json:"type"` // User's plan type:<br>`1` - Basic.<br>`2` - Licensed.<br>`3` - On-prem.<br>`99` - None (this can only be set with `ssoCreate`).
	Created_at string `json:"created_at,omitempty"` // User create time.
	Dept string `json:"dept,omitempty"` // Department.
	First_name string `json:"first_name,omitempty"` // User's first name.
	Last_client_version string `json:"last_client_version,omitempty"` // User last login client version.
	Role_name string `json:"role_name,omitempty"` // User's [role](https://support.zoom.us/hc/en-us/articles/115001078646-Role-Based-Access-Control) name.
	Timezone string `json:"timezone,omitempty"` // The time zone of the user.
	Use_pmi bool `json:"use_pmi,omitempty"` // Use Personal Meeting ID for instant meetings.
	Email string `json:"email"` // User's email address.
	Last_login_time string `json:"last_login_time,omitempty"` // User last login time.
	Last_name string `json:"last_name,omitempty"` // User's last name.
}

// MeetingMetric represents the MeetingMetric schema from the OpenAPI specification
type MeetingMetric struct {
	Duration string `json:"duration,omitempty"` // Meeting duration.
	Has_voip bool `json:"has_voip,omitempty"` // Indicates whether or not VoIP was used in the meeting.
	Dept string `json:"dept,omitempty"` // Department of the host.
	Has_pstn bool `json:"has_pstn,omitempty"` // Indicates whether or not the PSTN was used in the meeting.
	Has_recording bool `json:"has_recording,omitempty"` // Indicates whether or not the recording feature was used in the meeting.
	Has_sip bool `json:"has_sip,omitempty"` // Indicates whether or not someone joined the meeting using SIP.
	Custom_keys []map[string]interface{} `json:"custom_keys,omitempty"` // Custom keys and values assigned to the meeting.
	Uuid string `json:"uuid,omitempty"` // Meeting UUID. Please double encode your UUID when using it for API calls if the UUID begins with a '/'or contains '//' in it.
	Has_screen_share bool `json:"has_screen_share,omitempty"` // Indicates whether or not screenshare feature was used in the meeting.
	Id int64 `json:"id,omitempty"` // [Meeting ID](https://support.zoom.us/hc/en-us/articles/201362373-What-is-a-Meeting-ID-): Unique identifier of the meeting in "**long**" format(represented as int64 data type in JSON), also known as the meeting number.
	In_room_participants int `json:"in_room_participants,omitempty"` // The number of Zoom Room participants in the meeting.
	End_time string `json:"end_time,omitempty"` // Meeting end time.
	Has_video bool `json:"has_video,omitempty"` // Indicates whether or not video was used in the meeting.
	Participants int `json:"participants,omitempty"` // Meeting participant count.
	Start_time string `json:"start_time,omitempty"` // Meeting start time.
	User_type string `json:"user_type,omitempty"` // License type of the user.
	Email string `json:"email,omitempty"` // Email address of the host.
	Has_3rd_party_audio bool `json:"has_3rd_party_audio,omitempty"` // Indicates whether or not [third party audio](https://support.zoom.us/hc/en-us/articles/202470795-3rd-Party-Audio-Conference) was used in the meeting.
	Host string `json:"host,omitempty"` // Host display name.
	Topic string `json:"topic,omitempty"` // Meeting topic.
}

// GroupMember represents the GroupMember schema from the OpenAPI specification
type GroupMember struct {
	First_name string `json:"first_name,omitempty"` // User first name.
	Id string `json:"id,omitempty"` // User ID.
	Last_name string `json:"last_name,omitempty"` // User last name.
	TypeField int `json:"type,omitempty"` // User type.<br> `1` - Basic<br> `2` - Licensed<br> `3` - On-prem
	Email string `json:"email,omitempty"` // User email.
}

// AccountSettingsIntegration represents the AccountSettingsIntegration schema from the OpenAPI specification
type AccountSettingsIntegration struct {
	Box bool `json:"box,omitempty"` // Enable users who join a meeting from their mobile device to share content from their Box account.
	Dropbox bool `json:"dropbox,omitempty"` // Enable users who join a meeting from their mobile device to share content from their Dropbox account.
	Google_calendar bool `json:"google_calendar,omitempty"` // Enable meetings to be scheduled using Google Calendar.
	Google_drive bool `json:"google_drive,omitempty"` // Enable users who join a meeting from their mobile device to share content from their Google Drive.
	Kubi bool `json:"kubi,omitempty"` // Enable users to control a connected Kubi device from within a Zoom meeting.
	Microsoft_one_drive bool `json:"microsoft_one_drive,omitempty"` // Enable users who join a meeting from their mobile device to share content from their Microsoft OneDrive account.
}

// DeviceList represents the DeviceList schema from the OpenAPI specification
type DeviceList struct {
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_number int `json:"page_number,omitempty"` // **Deprecated**: This field has been deprecated. Please use the "next_page_token" field for pagination instead of this field. The page number of the current results.
	Page_size int `json:"page_size,omitempty"` // The number of records returned with a single API call.
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Devices []interface{} `json:"devices,omitempty"` // List of H.323/SIP Device objects.
}

// SessionWebinar represents the SessionWebinar schema from the OpenAPI specification
type SessionWebinar struct {
	Recurrence map[string]interface{} `json:"recurrence,omitempty"` // Recurrence object. Use this object only for a webinar of type `9` i.e., a recurring webinar with fixed time.
	Timezone string `json:"timezone,omitempty"` // Time zone to format start_time. For example, "America/Los_Angeles". For scheduled meetings only. Please reference our [time zone](#timezones) list for supported time zones and their formats.
	Duration int `json:"duration,omitempty"` // Webinar duration (minutes). Used for scheduled webinars only.
	Agenda string `json:"agenda,omitempty"` // Webinar description.
	Tracking_fields []map[string]interface{} `json:"tracking_fields,omitempty"` // Tracking fields
	Settings map[string]interface{} `json:"settings,omitempty"` // Webinar settings.
	Start_time string `json:"start_time,omitempty"` // Webinar start time. We support two formats for `start_time` - local time and GMT.<br> To set time as GMT the format should be `yyyy-MM-dd`T`HH:mm:ssZ`. To set time using a specific timezone, use `yyyy-MM-dd`T`HH:mm:ss` format and specify the timezone [ID](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#timezones) in the `timezone` field OR leave it blank and the timezone set on your Zoom account will be used. You can also set the time as UTC as the timezone field. The `start_time` should only be used for scheduled and / or recurring webinars with fixed time.
	Topic string `json:"topic,omitempty"` // Webinar topic.
	TypeField int `json:"type,omitempty"` // Webinar Types:<br>`5` - Webinar.<br>`6` - Recurring webinar with no fixed time.<br>`9` - Recurring webinar with a fixed time.
	Password string `json:"password,omitempty"` // Webinar Passcode. Passcode may only contain the following characters: [a-z A-Z 0-9 @ - _ *]. Max of 10 characters.
}

// CreateWebinarSettings represents the CreateWebinarSettings schema from the OpenAPI specification
type CreateWebinarSettings struct {
	Authentication_domains string `json:"authentication_domains,omitempty"` // Meeting authentication domains. This option, allows you to specify the rule so that Zoom users, whose email address contains a certain domain, can join the Webinar. You can either provide multiple domains, using a comma in between and/or use a wildcard for listing domains.
	Global_dial_in_countries []string `json:"global_dial_in_countries,omitempty"` // List of global dial-in countries
	Follow_up_absentees_email_notification map[string]interface{} `json:"follow_up_absentees_email_notification,omitempty"` // Send follow-up email to absentees.
	Alternative_hosts string `json:"alternative_hosts,omitempty"` // Alternative host emails or IDs. Multiple values separated by comma.
	Post_webinar_survey bool `json:"post_webinar_survey,omitempty"` // Zoom will open a survey page in attendees' browsers after leaving the webinar
	Survey_url string `json:"survey_url,omitempty"` // Survey url for post webinar survey
	Email_language string `json:"email_language,omitempty"` // Set the email language to one of the following: `en-US`,`de-DE`,`es-ES`,`fr-FR`,`jp-JP`,`pt-PT`,`ru-RU`,`zh-CN`, `zh-TW`, `ko-KO`, `it-IT`, `vi-VN`.
	Enforce_login_domains string `json:"enforce_login_domains,omitempty"` // Only signed-in users with specified domains can join meetings. **This field is deprecated and will not be supported in future.** <br><br> Instead of this field, use the "authentication_domains" field for this Webinar.
	Auto_recording string `json:"auto_recording,omitempty"` // Automatic recording:<br>`local` - Record on local.<br>`cloud` - Record on cloud.<br>`none` - Disabled.
	Contact_name string `json:"contact_name,omitempty"` // Contact name for registration
	Approval_type int `json:"approval_type,omitempty"` // The default value is `2`. To enable registration required, set the approval type to `0` or `1`. Values include:<br> `0` - Automatically approve.<br>`1` - Manually approve.<br>`2` - No registration required.
	Attendees_and_panelists_reminder_email_notification map[string]interface{} `json:"attendees_and_panelists_reminder_email_notification,omitempty"` // Send reminder email to attendees and panelists.
	Registrants_email_notification bool `json:"registrants_email_notification,omitempty"` // Send email notifications to registrants about approval, cancellation, denial of the registration. The value of this field must be set to true in order to use the `registrants_confirmation_email` field.
	Show_share_button bool `json:"show_share_button,omitempty"` // Show social share buttons on the registration page.
	Hd_video bool `json:"hd_video,omitempty"` // Default to HD video.
	Question_and_answer map[string]interface{} `json:"question_and_answer,omitempty"` // [Q&A](https://support.zoom.us/hc/en-us/articles/203686015-Using-Q-A-as-the-webinar-host#:~:text=Overview,and%20upvote%20each%20other's%20questions.) for webinar.
	Registrants_restrict_number int `json:"registrants_restrict_number,omitempty"` // Restrict number of registrants for a webinar. By default, it is set to `0`. A `0` value means that the restriction option is disabled. Provide a number higher than 0 to restrict the webinar registrants by the that number.
	Panelists_video bool `json:"panelists_video,omitempty"` // Start video when panelists join webinar.
	Enforce_login bool `json:"enforce_login,omitempty"` // Only signed-in users can join this meeting. **This field is deprecated and will not be supported in future.** <br><br> Instead of this field, use the "meeting_authentication", "authentication_option" and/or "authentication_domains" fields to establish the authentication mechanism for this Webinar.
	Close_registration bool `json:"close_registration,omitempty"` // Close registration after event date.
	Audio string `json:"audio,omitempty"` // Determine how participants can join the audio portion of the meeting.
	Allow_multiple_devices bool `json:"allow_multiple_devices,omitempty"` // Allow attendees to join from multiple devices.
	Registration_type int `json:"registration_type,omitempty"` // Registration types. Only used for recurring webinars with a fixed time.<br>`1` - Attendees register once and can attend any of the webinar sessions.<br>`2` - Attendees need to register for each session in order to attend.<br>`3` - Attendees register once and can choose one or more sessions to attend.
	Meeting_authentication bool `json:"meeting_authentication,omitempty"` // Only [authenticated](https://support.zoom.us/hc/en-us/articles/360037117472-Authentication-Profiles-for-Meetings-and-Webinars) users can join meeting if the value of this field is set to `true`.
	Authentication_option string `json:"authentication_option,omitempty"` // Specify the authentication type for users to join a Webinar with`meeting_authentication` setting set to `true`. The value of this field can be retrieved from the `id` field within `authentication_options` array in the response of [Get User Settings API](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usersettings).
	Follow_up_attendees_email_notification map[string]interface{} `json:"follow_up_attendees_email_notification,omitempty"` // Send follow-up email to attendees.
	Practice_session bool `json:"practice_session,omitempty"` // Enable practice session.
	Host_video bool `json:"host_video,omitempty"` // Start video when host joins webinar.
	Contact_email string `json:"contact_email,omitempty"` // Contact email for registration
	On_demand bool `json:"on_demand,omitempty"` // Make the webinar on-demand
	Panelists_invitation_email_notification bool `json:"panelists_invitation_email_notification,omitempty"` // * `true`: Send invitation email to panelists. * `false`: Do not send invitation email to panelists.
}

// UserSettings represents the UserSettings schema from the OpenAPI specification
type UserSettings struct {
	Tsp map[string]interface{} `json:"tsp,omitempty"` // Account Settings: TSP.
	Email_notification map[string]interface{} `json:"email_notification,omitempty"`
	Feature map[string]interface{} `json:"feature,omitempty"`
	In_meeting map[string]interface{} `json:"in_meeting,omitempty"`
	Profile map[string]interface{} `json:"profile,omitempty"`
	Recording map[string]interface{} `json:"recording,omitempty"`
	Schedule_meeting map[string]interface{} `json:"schedule_meeting,omitempty"`
	Telephony map[string]interface{} `json:"telephony,omitempty"`
}

// AccountListItem represents the AccountListItem schema from the OpenAPI specification
type AccountListItem struct {
	Accounts []map[string]interface{} `json:"accounts,omitempty"` // List of Account objects.
}

// PaginationToken represents the PaginationToken schema from the OpenAPI specification
type PaginationToken struct {
	Total_records int `json:"total_records,omitempty"` // The number of all records available across pages.
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_size int `json:"page_size,omitempty"` // The number of records returned within a single API call.
}

// AccountPlanRequired represents the AccountPlanRequired schema from the OpenAPI specification
type AccountPlanRequired struct {
	Hosts int `json:"hosts"` // Number of hosts for this plan.
	TypeField string `json:"type"` // Account <a href="https://marketplace.zoom.us/docs/api-reference/other-references/plans">plan type.</a>
}

// PaginationToken4Qos represents the PaginationToken4Qos schema from the OpenAPI specification
type PaginationToken4Qos struct {
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceed the current page size. The expiration period for this token is 15 minutes.
	Page_count int64 `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_size int `json:"page_size,omitempty"` // The number of items per page.
	Total_records int64 `json:"total_records,omitempty"` // The number of all records available across pages.
}

// Channel represents the Channel schema from the OpenAPI specification
type Channel struct {
	Id string `json:"id,omitempty"` // Channel Id.
	Name string `json:"name,omitempty"` // Name of the channel.
	TypeField int `json:"type,omitempty"` // Type of the channel. The value can be one of the following:<br> `1`: Private channel. In this type of channel, members must be invited to join a channel.<br> `2`: Private channel with members that belong to one Zoom account. Members in this channel should be invited and the members should be from the same organization.<br> `3`: Public channel. Anyone can search for this channel and join the channel.<br>
}

// PanelistList represents the PanelistList schema from the OpenAPI specification
type PanelistList struct {
	Panelists []interface{} `json:"panelists,omitempty"` // List of panelist objects.
	Total_records int `json:"total_records,omitempty"` // Total records.
}

// PhonePlan represents the PhonePlan schema from the OpenAPI specification
type PhonePlan struct {
	Plan_base map[string]interface{} `json:"plan_base,omitempty"` // Additional phone base plans.
	Plan_calling []map[string]interface{} `json:"plan_calling,omitempty"` // Additional phone calling plans.
	Plan_number []map[string]interface{} `json:"plan_number,omitempty"` // Additional phone number plans.
}

// RegistrantStatus represents the RegistrantStatus schema from the OpenAPI specification
type RegistrantStatus struct {
	Action string `json:"action"` // Registrant Status:<br>`approve` - Approve registrant.<br>`cancel` - Cancel previously approved registrant's registration.<br>`deny` - Deny registrant.
	Registrants []map[string]interface{} `json:"registrants,omitempty"` // List of registrants.
}

// UserSettingsEmailNotification represents the UserSettingsEmailNotification schema from the OpenAPI specification
type UserSettingsEmailNotification struct {
	Alternative_host_reminder bool `json:"alternative_host_reminder,omitempty"` // When an alternative host is set or removed from a meeting.
	Cancel_meeting_reminder bool `json:"cancel_meeting_reminder,omitempty"` // When a meeting is cancelled.
	Jbh_reminder bool `json:"jbh_reminder,omitempty"` // When attendees join meeting before host.
	Schedule_for_reminder bool `json:"schedule_for_reminder,omitempty"` // Notify the host there is a meeting is scheduled, rescheduled, or cancelled.
}

// UserAssistantsList represents the UserAssistantsList schema from the OpenAPI specification
type UserAssistantsList struct {
	Assistants []map[string]interface{} `json:"assistants,omitempty"` // List of User's assistants.
}

// RoleMembersList represents the RoleMembersList schema from the OpenAPI specification
type RoleMembersList struct {
	Page_size int `json:"page_size,omitempty"` // The number of records returned within a single API call.
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
	Members []map[string]interface{} `json:"members,omitempty"` // List of a Role Members
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_number int `json:"page_number,omitempty"` // The page number of the current results.
}

// ZoomRoom represents the ZoomRoom schema from the OpenAPI specification
type ZoomRoom struct {
	Status string `json:"status,omitempty"` // Zoom room status.
	Account_type string `json:"account_type,omitempty"` // Zoom room email type.
	Email string `json:"email,omitempty"` // Zoom room email.
	Microphone string `json:"microphone,omitempty"` // Zoom room microphone.
	Room_name string `json:"room_name,omitempty"` // Zoom room name.
	Speaker string `json:"speaker,omitempty"` // Zoom room speaker.
	Device_ip string `json:"device_ip,omitempty"` // Zoom room device IP.
	Calender_name string `json:"calender_name,omitempty"` // Zoom calendar name.
	Id string `json:"id,omitempty"` // Zoom room ID.
	Location string `json:"location,omitempty"` // Zoom room location.
	Camera string `json:"camera,omitempty"` // Zoom room camera.
	Health string `json:"health,omitempty"`
	Issues []string `json:"issues,omitempty"` // Zoom Room issues.
	Last_start_time string `json:"last_start_time,omitempty"` // Zoom room last start time.
}

// DateTime represents the DateTime schema from the OpenAPI specification
type DateTime struct {
	To string `json:"to,omitempty"` // End Date.
	From string `json:"from,omitempty"` // Start Date.
}

// UserSchedulersList represents the UserSchedulersList schema from the OpenAPI specification
type UserSchedulersList struct {
	Schedulers []map[string]interface{} `json:"schedulers,omitempty"` // List of users for whom the current user can schedule meetings.
}

// MeetingInstances represents the MeetingInstances schema from the OpenAPI specification
type MeetingInstances struct {
	Meetings []map[string]interface{} `json:"meetings,omitempty"` // List of ended meeting instances.
}

// MeetingUpdate represents the MeetingUpdate schema from the OpenAPI specification
type MeetingUpdate struct {
	Agenda string `json:"agenda,omitempty"` // Meeting description.
	Settings interface{} `json:"settings,omitempty"`
	Template_id string `json:"template_id,omitempty"` // Unique identifier of the meeting template. Use this field if you would like to [schedule the meeting from a meeting template](https://support.zoom.us/hc/en-us/articles/360036559151-Meeting-templates#h_86f06cff-0852-4998-81c5-c83663c176fb). You can retrieve the value of this field by calling the [List meeting templates]() API.
	Timezone string `json:"timezone,omitempty"` // Time zone to format start_time. For example, "America/Los_Angeles". For scheduled meetings only. Please reference our [time zone](#timezones) list for supported time zones and their formats.
	Recurrence map[string]interface{} `json:"recurrence,omitempty"` // Recurrence object. Use this object only for a meeting with type `8` i.e., a recurring meeting with fixed time.
	Tracking_fields []map[string]interface{} `json:"tracking_fields,omitempty"` // Tracking fields
	Duration int `json:"duration,omitempty"` // Meeting duration (minutes). Used for scheduled meetings only.
	Password string `json:"password,omitempty"` // Meeting passcode. Passcode may only contain the following characters: [a-z A-Z 0-9 @ - _ *] and can have a maximum of 10 characters. **Note:** If the account owner or the admin has configured [minimum passcode requirement settings](https://support.zoom.us/hc/en-us/articles/360033559832-Meeting-and-webinar-passwords#h_a427384b-e383-4f80-864d-794bf0a37604), the passcode value provided here must meet those requirements. <br><br>If the requirements are enabled, you can view those requirements by calling either the [Get User Settings API](https://marketplace.zoom.us/docs/api-reference/zoom-api/users/usersettings) or the [Get Account Settings](https://marketplace.zoom.us/docs/api-reference/zoom-api/accounts/accountsettings) API.
	Start_time string `json:"start_time,omitempty"` // Meeting start time. When using a format like "yyyy-MM-dd'T'HH:mm:ss'Z'", always use GMT time. When using a format like "yyyy-MM-dd'T'HH:mm:ss", you should use local time and specify the time zone. Only used for scheduled meetings and recurring meetings with a fixed time.
	TypeField int `json:"type,omitempty"` // Meeting Types:<br>`1` - Instant meeting.<br>`2` - Scheduled meeting.<br>`3` - Recurring meeting with no fixed time.<br>`8` - Recurring meeting with a fixed time.
	Topic string `json:"topic,omitempty"` // Meeting topic.
}

// MeetingRecordingRegistrantList represents the MeetingRecordingRegistrantList schema from the OpenAPI specification
type MeetingRecordingRegistrantList struct {
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_number int `json:"page_number,omitempty"` // **Deprecated**: This field has been deprecated. Please use the "next_page_token" field for pagination instead of this field. The page number of the current results.
	Page_size int `json:"page_size,omitempty"` // The number of records returned with a single API call.
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Registrants []interface{} `json:"registrants,omitempty"` // List of Registrant objects
}

// AccountSettingsRecording represents the AccountSettingsRecording schema from the OpenAPI specification
type AccountSettingsRecording struct {
	Auto_delete_cmr bool `json:"auto_delete_cmr,omitempty"` // Allow Zoom to permanantly delete recordings automatically after a specified number of days.
	Record_audio_file bool `json:"record_audio_file,omitempty"` // Record an audio only file.
	Record_gallery_view bool `json:"record_gallery_view,omitempty"` // Record the gallery view with a shared screen.
	Archive map[string]interface{} `json:"archive,omitempty"` // [Archiving solution](https://support.zoom.us/hc/en-us/articles/360050431572-Archiving-Meeting-and-Webinar-data) settings. This setting can only be used if you have been granted with archiving solution access by the Zoom support team.
	Prevent_host_access_recording bool `json:"prevent_host_access_recording,omitempty"` // If set to `true`, meeting hosts cannot view their meeting cloud recordings. Only the admins who have recording management privilege can access them.
	Save_chat_text bool `json:"save_chat_text,omitempty"` // Save the chat text from the meeting.
	Cloud_recording_download bool `json:"cloud_recording_download,omitempty"` // Cloud recording downloads.
	Auto_delete_cmr_days int `json:"auto_delete_cmr_days,omitempty"` // When `auto_delete_cmr` function is 'true' this value will set the number of days before the auto deletion of cloud recordings.
	Recording_audio_transcript bool `json:"recording_audio_transcript,omitempty"` // Automatically transcribe the audio of the meeting or webinar to the cloud.
	Recording_password_requirement map[string]interface{} `json:"recording_password_requirement,omitempty"` // This object represents the minimum password requirements set for recordings via Account Recording Settings.
	Cloud_recording bool `json:"cloud_recording,omitempty"` // Allow hosts to record and save the meeting in the cloud.
	Ip_address_access_control map[string]interface{} `json:"ip_address_access_control,omitempty"` // Setting to allow cloud recording access only from specific IP address ranges.
	Recording_disclaimer bool `json:"recording_disclaimer,omitempty"` // Show a disclaimer to participants before a recording starts
	Host_delete_cloud_recording bool `json:"host_delete_cloud_recording,omitempty"` // If the value of this field is set to `true`, hosts will be able to delete the recordings. If this option is set to `false`, the recordings cannot be deleted by the host and only admin can delete them.
	Record_speaker_view bool `json:"record_speaker_view,omitempty"` // Record the active speaker with a shared screen.
	Allow_recovery_deleted_cloud_recordings bool `json:"allow_recovery_deleted_cloud_recordings,omitempty"` // Allow recovery of deleted cloud recordings from trash. If the value of this field is set to `true`, deleted cloud recordings will be kept in trash for 30 days after deletion and can be recovered within that period.
	Local_recording bool `json:"local_recording,omitempty"` // Allow hosts and participants to record the meeting using a local file.
	Required_password_for_existing_cloud_recordings bool `json:"required_password_for_existing_cloud_recordings,omitempty"` // Require a passcode to access existing cloud recordings.
	Auto_recording string `json:"auto_recording,omitempty"` // Automatic recording:<br>`local` - Record on local.<br>`cloud` - Record on cloud.<br>`none` - Disabled.
	Cloud_recording_download_host bool `json:"cloud_recording_download_host,omitempty"` // Only the host can download cloud recordings.
	Show_timestamp bool `json:"show_timestamp,omitempty"` // Add a timestamp to the recording.
	Account_user_access_recording bool `json:"account_user_access_recording,omitempty"` // Cloud recordings are only accessible to account members. People outside of your organization cannot open links that provide access to cloud recordings.
}

// Profile represents the Profile schema from the OpenAPI specification
type Profile struct {
	Recording_storage_location map[string]interface{} `json:"recording_storage_location,omitempty"`
}

// Account represents the Account schema from the OpenAPI specification
type Account struct {
	First_name string `json:"first_name"` // User's first name.
	Last_name string `json:"last_name"` // User's last name.
	Options map[string]interface{} `json:"options,omitempty"` // Account options object.
	Password string `json:"password"` // User's password.
	Vanity_url string `json:"vanity_url,omitempty"` // Account Vanity URL
	Email string `json:"email"` // User's email address.
}

// RecordingRegistrantStatus represents the RecordingRegistrantStatus schema from the OpenAPI specification
type RecordingRegistrantStatus struct {
	Registrants []interface{} `json:"registrants,omitempty"` // List of registrants
	Action string `json:"action"`
}

// RecordingRegistrantQuestions represents the RecordingRegistrantQuestions schema from the OpenAPI specification
type RecordingRegistrantQuestions struct {
	Custom_questions []map[string]interface{} `json:"custom_questions,omitempty"` // Array of Registrant Custom Questions
	Questions []map[string]interface{} `json:"questions,omitempty"` // Array of Registrant Questions
}

// Occurrence represents the Occurrence schema from the OpenAPI specification
type Occurrence struct {
	Start_time string `json:"start_time,omitempty"` // Start time.
	Status string `json:"status,omitempty"` // Occurrence status.
	Duration int `json:"duration,omitempty"` // Duration.
	Occurrence_id string `json:"occurrence_id,omitempty"` // Occurrence ID: Unique Identifier that identifies an occurrence of a recurring webinar. [Recurring webinars](https://support.zoom.us/hc/en-us/articles/216354763-How-to-Schedule-A-Recurring-Webinar) can have a maximum of 50 occurrences.
}

// TrackingFieldList represents the TrackingFieldList schema from the OpenAPI specification
type TrackingFieldList struct {
	Total_records int `json:"total_records,omitempty"` // The number of all records available across pages
	Tracking_fields []interface{} `json:"tracking_fields,omitempty"` // Array of Tracking Fields
}

// RecordingList represents the RecordingList schema from the OpenAPI specification
type RecordingList struct {
	Recording_files []interface{} `json:"recording_files,omitempty"` // List of recording file.
}

// TrackingField represents the TrackingField schema from the OpenAPI specification
type TrackingField struct {
	Field string `json:"field,omitempty"` // Label/ Name for the tracking field.
	Recommended_values []string `json:"recommended_values,omitempty"` // Array of recommended values
	Required bool `json:"required,omitempty"` // Tracking Field Required
	Visible bool `json:"visible,omitempty"` // Tracking Field Visible
}

// GroupUserSettingsAuthenticationUpdate represents the GroupUserSettingsAuthenticationUpdate schema from the OpenAPI specification
type GroupUserSettingsAuthenticationUpdate struct {
}

// TSP represents the TSP schema from the OpenAPI specification
type TSP struct {
	Conference_code string `json:"conference_code"` // Conference code: numeric value, length is less than 16.
	Dial_in_numbers []map[string]interface{} `json:"dial_in_numbers,omitempty"` // List of dial in numbers.
	Leader_pin string `json:"leader_pin"` // Leader PIN: numeric value, length is less than 16.
	Tsp_bridge string `json:"tsp_bridge,omitempty"` // Telephony bridge
}

// Registrant represents the Registrant schema from the OpenAPI specification
type Registrant struct {
	Org string `json:"org,omitempty"` // Registrant's Organization.
	No_of_employees string `json:"no_of_employees,omitempty"` // Number of Employees:<br>`1-20`<br>`21-50`<br>`51-100`<br>`101-500`<br>`500-1,000`<br>`1,001-5,000`<br>`5,001-10,000`<br>`More than 10,000`
	Purchasing_time_frame string `json:"purchasing_time_frame,omitempty"` // This field can be included to gauge interest of webinar attendees towards buying your product or service. Purchasing Time Frame:<br>`Within a month`<br>`1-3 months`<br>`4-6 months`<br>`More than 6 months`<br>`No timeframe`
	Zip string `json:"zip,omitempty"` // Registrant's Zip/Postal Code.
	Industry string `json:"industry,omitempty"` // Registrant's Industry.
	Job_title string `json:"job_title,omitempty"` // Registrant's job title.
	Phone string `json:"phone,omitempty"` // Registrant's Phone number.
	Last_name string `json:"last_name,omitempty"` // Registrant's last name.
	City string `json:"city,omitempty"` // Registrant's city.
	Country string `json:"country,omitempty"` // Registrant's country. The value of this field must be in two-letter abbreviated form and must match the ID field provided in the [Countries](https://marketplace.zoom.us/docs/api-reference/other-references/abbreviation-lists#countries) table.
	State string `json:"state,omitempty"` // Registrant's State/Province.
	Comments string `json:"comments,omitempty"` // A field that allows registrants to provide any questions or comments that they might have.
	Custom_questions []map[string]interface{} `json:"custom_questions,omitempty"` // Custom questions.
	Role_in_purchase_process string `json:"role_in_purchase_process,omitempty"` // Role in Purchase Process:<br>`Decision Maker`<br>`Evaluator/Recommender`<br>`Influencer`<br>`Not involved`
	First_name string `json:"first_name"` // Registrant's first name.
	Address string `json:"address,omitempty"` // Registrant's address.
	Email string `json:"email"` // A valid email address of the registrant.
}

// AccountSettingsFeature represents the AccountSettingsFeature schema from the OpenAPI specification
type AccountSettingsFeature struct {
	Meeting_capacity int `json:"meeting_capacity,omitempty"` // Set the maximum number of participants a host can have in a single meeting.
}

// AccountList represents the AccountList schema from the OpenAPI specification
type AccountList struct {
	Page_size int `json:"page_size,omitempty"` // The number of records returned with a single API call.
	Total_records int `json:"total_records,omitempty"` // The total number of all the records available across pages.
	Next_page_token string `json:"next_page_token,omitempty"` // The next page token is used to paginate through large result sets. A next page token will be returned whenever the set of available results exceeds the current page size. The expiration period for this token is 15 minutes.
	Page_count int `json:"page_count,omitempty"` // The number of pages returned for the request made.
	Page_number int `json:"page_number,omitempty"` // **Deprecated**: This field has been deprecated. Please use the "next_page_token" field for pagination instead of this field. The page number of the current results.
	Accounts []map[string]interface{} `json:"accounts,omitempty"` // List of Account objects.
}

// Recurrence represents the Recurrence schema from the OpenAPI specification
type Recurrence struct {
	Monthly_day int `json:"monthly_day,omitempty"` // Use this field **only if you're scheduling a recurring meeting of type** `3` to state which day in a month, the meeting should recur. The value range is from 1 to 31. For instance, if you would like the meeting to recur on 23rd of each month, provide `23` as the value of this field and `1` as the value of the `repeat_interval` field. Instead, if you would like the meeting to recur every three months, on 23rd of the month, change the value of the `repeat_interval` field to `3`.
	Monthly_week int `json:"monthly_week,omitempty"` // Use this field **only if you're scheduling a recurring meeting of type** `3` to state the week of the month when the meeting should recur. If you use this field, **you must also use the `monthly_week_day` field to state the day of the week when the meeting should recur.** <br>`-1` - Last week of the month.<br>`1` - First week of the month.<br>`2` - Second week of the month.<br>`3` - Third week of the month.<br>`4` - Fourth week of the month.
	Monthly_week_day int `json:"monthly_week_day,omitempty"` // Use this field **only if you're scheduling a recurring meeting of type** `3` to state a specific day in a week when the monthly meeting should recur. To use this field, you must also use the `monthly_week` field. <br>`1` - Sunday.<br>`2` - Monday.<br>`3` - Tuesday.<br>`4` - Wednesday.<br>`5` - Thursday.<br>`6` - Friday.<br>`7` - Saturday.
	Repeat_interval int `json:"repeat_interval,omitempty"` // Define the interval at which the meeting should recur. For instance, if you would like to schedule a meeting that recurs every two months, you must set the value of this field as `2` and the value of the `type` parameter as `3`. For a daily meeting, the maximum interval you can set is `90` days. For a weekly meeting the maximum interval that you can set is of `12` weeks. For a monthly meeting, there is a maximum of `3` months.
	TypeField int `json:"type"` // Recurrence meeting types:<br>`1` - Daily.<br>`2` - Weekly.<br>`3` - Monthly.
	Weekly_days string `json:"weekly_days,omitempty"` // This field is required **if you're scheduling a recurring meeting of type** `2` to state which day(s) of the week the meeting should repeat. <br> <br> The value for this field could be a number between `1` to `7` in string format. For instance, if the meeting should recur on Sunday, provide `"1"` as the value of this field.<br><br> **Note:** If you would like the meeting to occur on multiple days of a week, you should provide comma separated values for this field. For instance, if the meeting should recur on Sundays and Tuesdays provide `"1,3"` as the value of this field. <br>`1` - Sunday. <br>`2` - Monday.<br>`3` - Tuesday.<br>`4` - Wednesday.<br>`5` - Thursday.<br>`6` - Friday.<br>`7` - Saturday.
	End_date_time string `json:"end_date_time,omitempty"` // Select the final date on which the meeting will recur before it is canceled. Should be in UTC time, such as 2017-11-25T12:00:00Z. (Cannot be used with "end_times".)
	End_times int `json:"end_times,omitempty"` // Select how many times the meeting should recur before it is canceled. (Cannot be used with "end_date_time".)
}

// AccountSettings represents the AccountSettings schema from the OpenAPI specification
type AccountSettings struct {
	Email_notification map[string]interface{} `json:"email_notification,omitempty"` // Account Settings: Notification.
	Feature map[string]interface{} `json:"feature,omitempty"` // Account Settings: Feature.
	Integration map[string]interface{} `json:"integration,omitempty"` // Account Settings: Integration.
	Profile map[string]interface{} `json:"profile,omitempty"`
	Recording map[string]interface{} `json:"recording,omitempty"` // Account Settings: Recording.
	Tsp map[string]interface{} `json:"tsp,omitempty"` // Account Settings: TSP.
	Schedule_meeting map[string]interface{} `json:"schedule_meeting,omitempty"` // Account Settings: Schedule Meeting.
	Security map[string]interface{} `json:"security,omitempty"` // [Security settings](https://support.zoom.us/hc/en-us/articles/360034675592-Advanced-security-settings#h_bf8a25f6-9a66-447a-befd-f02ed3404f89) of an Account.
	Zoom_rooms map[string]interface{} `json:"zoom_rooms,omitempty"` // Account Settings: Zoom Rooms.
	In_meeting map[string]interface{} `json:"in_meeting,omitempty"` // Account Settings: In Meeting.
	Telephony map[string]interface{} `json:"telephony,omitempty"` // Account Settings: Telephony.
}
