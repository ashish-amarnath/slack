package types

// ResponseRtmStart represents rtm start message
type ResponseRtmStart struct {
	Ok    bool        `json:"ok"`
	Error string      `json:"error"`
	URL   string      `json:"url"`
	Bot   BotID       `json:"self"`
	Users []SlackUser `json:"users"`
}

// BotID represents the object storing the user ID
type BotID struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Message represents messages written to and read from web socket.
type Message struct {
	ID      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
	User    string `json:"user"`
}

// SlackUser represents a slack user
type SlackUser struct {
	ID       string `json:"id"`
	TeamID   string `json:"team_id"`
	Name     string `json:"name"`
	Deleted  bool   `json:"deleted"`
	RealName string `json:"real_name"`
	Tz       string `json:"tz"`
	TzLabel  string `json:"tz_label"`
	TzOffset int    `json:"tz_offset"`
	Profile  struct {
		FirstName             string `json:"first_name"`
		LastName              string `json:"last_name"`
		RealName              string `json:"real_name"`
		DisplayName           string `json:"display_name"`
		RealNameNormalized    string `json:"real_name_normalized"`
		DisplayNameNormalized string `json:"display_name_normalized"`
		Email                 string `json:"email"`
		Team                  string `json:"team"`
	} `json:"profile"`
	IsAdmin           bool   `json:"is_admin"`
	IsOwner           bool   `json:"is_owner"`
	IsPrimaryOwner    bool   `json:"is_primary_owner"`
	IsRestricted      bool   `json:"is_restricted"`
	IsUltraRestricted bool   `json:"is_ultra_restricted"`
	IsBot             bool   `json:"is_bot"`
	Updated           int    `json:"updated"`
	IsAppUser         bool   `json:"is_app_user"`
	Presence          string `json:"presence"`
}
