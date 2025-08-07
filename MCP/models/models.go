package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Defspinnedinfo represents the Defspinnedinfo schema from the OpenAPI specification
type Defspinnedinfo struct {
}

// Objsexternalorgmigrations represents the Objsexternalorgmigrations schema from the OpenAPI specification
type Objsexternalorgmigrations struct {
	Current []map[string]interface{} `json:"current"`
	Date_updated int `json:"date_updated"`
}

// Objsprimaryowner represents the Objsprimaryowner schema from the OpenAPI specification
type Objsprimaryowner struct {
	Email string `json:"email"`
	Id string `json:"id"`
}

// Objsreaction represents the Objsreaction schema from the OpenAPI specification
type Objsreaction struct {
	Users []string `json:"users"`
	Count int `json:"count"`
	Name string `json:"name"`
}

// Objscomment represents the Objscomment schema from the OpenAPI specification
type Objscomment struct {
	Timestamp int `json:"timestamp"`
	User string `json:"user"`
	Created int `json:"created"`
	Num_stars int `json:"num_stars,omitempty"`
	Is_starred bool `json:"is_starred,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	Comment string `json:"comment"`
	Id string `json:"id"`
	Is_intro bool `json:"is_intro"`
	Pinned_info Defspinnedinfo `json:"pinned_info,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
}

// Objsmessage represents the Objsmessage schema from the OpenAPI specification
type Objsmessage struct {
	Old_name string `json:"old_name,omitempty"`
	Unread_count int `json:"unread_count,omitempty"`
	TypeField string `json:"type"`
	Inviter string `json:"inviter,omitempty"`
	User_team string `json:"user_team,omitempty"`
	Purpose string `json:"purpose,omitempty"`
	Ts string `json:"ts"`
	User string `json:"user,omitempty"`
	Blocks []map[string]interface{} `json:"blocks,omitempty"` // This is a very loose definition, in the future, we'll populate this with deeper schema in this definition namespace.
	Is_delayed_message bool `json:"is_delayed_message,omitempty"`
	Parent_user_id string `json:"parent_user_id,omitempty"`
	Reply_users []string `json:"reply_users,omitempty"`
	Is_starred bool `json:"is_starred,omitempty"`
	Username string `json:"username,omitempty"`
	Icons map[string]interface{} `json:"icons,omitempty"`
	Latest_reply string `json:"latest_reply,omitempty"`
	Name string `json:"name,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	Thread_ts string `json:"thread_ts,omitempty"`
	Reply_users_count int `json:"reply_users_count,omitempty"`
	Bot_profile Objsbotprofile `json:"bot_profile,omitempty"`
	Last_read string `json:"last_read,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	Client_msg_id string `json:"client_msg_id,omitempty"`
	Permalink string `json:"permalink,omitempty"`
	Reply_count int `json:"reply_count,omitempty"`
	Display_as_bot bool `json:"display_as_bot,omitempty"`
	Files []Objsfile `json:"files,omitempty"`
	File Objsfile `json:"file,omitempty"`
	Bot_id map[string]interface{} `json:"bot_id,omitempty"`
	User_profile Objsuserprofileshort `json:"user_profile,omitempty"`
	Text string `json:"text"`
	Topic string `json:"topic,omitempty"`
	Is_intro bool `json:"is_intro,omitempty"`
	Team string `json:"team,omitempty"`
	Source_team string `json:"source_team,omitempty"`
	Subscribed bool `json:"subscribed,omitempty"`
	Comment Objscomment `json:"comment,omitempty"`
	Subtype string `json:"subtype,omitempty"`
	Upload bool `json:"upload,omitempty"`
	Attachments []map[string]interface{} `json:"attachments,omitempty"`
}

// Objsfile represents the Objsfile schema from the OpenAPI specification
type Objsfile struct {
	Permalink_public string `json:"permalink_public,omitempty"`
	Has_rich_preview bool `json:"has_rich_preview,omitempty"`
	Thumb_1024_w int `json:"thumb_1024_w,omitempty"`
	State string `json:"state,omitempty"`
	Url_private string `json:"url_private,omitempty"`
	Title string `json:"title,omitempty"`
	Thumb_720_h int `json:"thumb_720_h,omitempty"`
	Original_w int `json:"original_w,omitempty"`
	Thumb_960 string `json:"thumb_960,omitempty"`
	Thumb_160 string `json:"thumb_160,omitempty"`
	Last_editor string `json:"last_editor,omitempty"`
	Thumb_960_h int `json:"thumb_960_h,omitempty"`
	Updated int `json:"updated,omitempty"`
	Thumb_64 string `json:"thumb_64,omitempty"`
	Thumb_tiny string `json:"thumb_tiny,omitempty"`
	Timestamp int `json:"timestamp,omitempty"`
	Date_delete int `json:"date_delete,omitempty"`
	Thumb_800 string `json:"thumb_800,omitempty"`
	Url_private_download string `json:"url_private_download,omitempty"`
	Thumb_480_h int `json:"thumb_480_h,omitempty"`
	Mode string `json:"mode,omitempty"`
	External_url string `json:"external_url,omitempty"`
	Is_starred bool `json:"is_starred,omitempty"`
	Thumb_360_w int `json:"thumb_360_w,omitempty"`
	External_id string `json:"external_id,omitempty"`
	Shares map[string]interface{} `json:"shares,omitempty"`
	Name string `json:"name,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Ims []string `json:"ims,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	Permalink string `json:"permalink,omitempty"`
	Preview string `json:"preview,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	Filetype string `json:"filetype,omitempty"`
	Public_url_shared bool `json:"public_url_shared,omitempty"`
	Size int `json:"size,omitempty"`
	User string `json:"user,omitempty"`
	Is_public bool `json:"is_public,omitempty"`
	Thumb_480_w int `json:"thumb_480_w,omitempty"`
	Comments_count int `json:"comments_count,omitempty"`
	Non_owner_editable bool `json:"non_owner_editable,omitempty"`
	Is_tombstoned bool `json:"is_tombstoned,omitempty"`
	Thumb_1024_h int `json:"thumb_1024_h,omitempty"`
	Pinned_info Defspinnedinfo `json:"pinned_info,omitempty"`
	Display_as_bot bool `json:"display_as_bot,omitempty"`
	Original_h int `json:"original_h,omitempty"`
	Source_team string `json:"source_team,omitempty"`
	Thumb_720 string `json:"thumb_720,omitempty"`
	Num_stars int `json:"num_stars,omitempty"`
	Thumb_360 string `json:"thumb_360,omitempty"`
	Thumb_480 string `json:"thumb_480,omitempty"`
	Thumb_360_h int `json:"thumb_360_h,omitempty"`
	Thumb_720_w int `json:"thumb_720_w,omitempty"`
	Thumb_960_w int `json:"thumb_960_w,omitempty"`
	Thumb_80 string `json:"thumb_80,omitempty"`
	Editable bool `json:"editable,omitempty"`
	Thumb_800_w int `json:"thumb_800_w,omitempty"`
	User_team string `json:"user_team,omitempty"`
	Thumb_800_h int `json:"thumb_800_h,omitempty"`
	External_type string `json:"external_type,omitempty"`
	Is_external bool `json:"is_external,omitempty"`
	Created int `json:"created,omitempty"`
	Thumb_1024 string `json:"thumb_1024,omitempty"`
	Channels []string `json:"channels,omitempty"`
	Editor string `json:"editor,omitempty"`
	Image_exif_rotation int `json:"image_exif_rotation,omitempty"`
	Username string `json:"username,omitempty"`
	Id string `json:"id,omitempty"`
	Pretty_type string `json:"pretty_type,omitempty"`
	Mimetype string `json:"mimetype,omitempty"`
}

// Objsicon represents the Objsicon schema from the OpenAPI specification
type Objsicon struct {
	Image_default bool `json:"image_default,omitempty"`
	Image_102 string `json:"image_102,omitempty"`
	Image_132 string `json:"image_132,omitempty"`
	Image_230 string `json:"image_230,omitempty"`
	Image_34 string `json:"image_34,omitempty"`
	Image_44 string `json:"image_44,omitempty"`
	Image_68 string `json:"image_68,omitempty"`
	Image_88 string `json:"image_88,omitempty"`
}

// Objsresponsemetadata represents the Objsresponsemetadata schema from the OpenAPI specification
type Objsresponsemetadata struct {
}

// Objsteamprofilefield represents the Objsteamprofilefield schema from the OpenAPI specification
type Objsteamprofilefield struct {
	Label string `json:"label"`
	Options map[string]interface{} `json:"options,omitempty"`
	Ordering float64 `json:"ordering"`
	TypeField string `json:"type"`
	Hint string `json:"hint"`
	Id string `json:"id"`
	Is_hidden bool `json:"is_hidden,omitempty"`
}

// Objsteamprofilefieldoption represents the Objsteamprofilefieldoption schema from the OpenAPI specification
type Objsteamprofilefieldoption struct {
}

// Objsuserprofileshort represents the Objsuserprofileshort schema from the OpenAPI specification
type Objsuserprofileshort struct {
	Image_72 string `json:"image_72"`
	Avatar_hash string `json:"avatar_hash"`
	Display_name string `json:"display_name"`
	Display_name_normalized string `json:"display_name_normalized,omitempty"`
	Real_name_normalized string `json:"real_name_normalized,omitempty"`
	Is_ultra_restricted bool `json:"is_ultra_restricted"`
	Name string `json:"name"`
	Real_name string `json:"real_name"`
	Team string `json:"team"`
	Is_restricted bool `json:"is_restricted"`
}

// Objspaging represents the Objspaging schema from the OpenAPI specification
type Objspaging struct {
	Spill int `json:"spill,omitempty"`
	Total int `json:"total"`
	Count int `json:"count,omitempty"`
	Page int `json:"page"`
	Pages int `json:"pages,omitempty"`
	Per_page int `json:"per_page,omitempty"`
}

// Objsuser represents the Objsuser schema from the OpenAPI specification
type Objsuser struct {
}

// Objsuserprofile represents the Objsuserprofile schema from the OpenAPI specification
type Objsuserprofile struct {
	Last_avatar_image_hash string `json:"last_avatar_image_hash,omitempty"`
	Phone string `json:"phone"`
	Real_name_normalized string `json:"real_name_normalized"`
	Status_default_text string `json:"status_default_text,omitempty"`
	Avatar_hash string `json:"avatar_hash"`
	Is_app_user bool `json:"is_app_user,omitempty"`
	Display_name string `json:"display_name"`
	Status_default_emoji string `json:"status_default_emoji,omitempty"`
	Is_custom_image bool `json:"is_custom_image,omitempty"`
	Title string `json:"title"`
	Pronouns string `json:"pronouns,omitempty"`
	Display_name_normalized string `json:"display_name_normalized"`
	Status_text string `json:"status_text"`
	Status_expiration int `json:"status_expiration,omitempty"`
	Always_active bool `json:"always_active,omitempty"`
	Api_app_id string `json:"api_app_id,omitempty"`
	Real_name string `json:"real_name"`
	Bot_id string `json:"bot_id,omitempty"`
	Skype string `json:"skype"`
	Updated int `json:"updated,omitempty"`
	Memberships_count int `json:"memberships_count,omitempty"`
	User_id string `json:"user_id,omitempty"`
	Status_emoji string `json:"status_emoji"`
	Team string `json:"team,omitempty"`
}

// Objsresources represents the Objsresources schema from the OpenAPI specification
type Objsresources struct {
	Excluded_ids []map[string]interface{} `json:"excluded_ids,omitempty"`
	Ids []map[string]interface{} `json:"ids"`
	Wildcard bool `json:"wildcard,omitempty"`
}

// Objsbotprofile represents the Objsbotprofile schema from the OpenAPI specification
type Objsbotprofile struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Team_id string `json:"team_id"`
	Updated int `json:"updated"`
	App_id string `json:"app_id"`
	Deleted bool `json:"deleted"`
	Icons map[string]interface{} `json:"icons"`
}

// Objssubteam represents the Objssubteam schema from the OpenAPI specification
type Objssubteam struct {
	Date_update int `json:"date_update"`
	User_count int `json:"user_count,omitempty"`
	Enterprise_subteam_id string `json:"enterprise_subteam_id"`
	Is_usergroup bool `json:"is_usergroup"`
	Prefs map[string]interface{} `json:"prefs"`
	Deleted_by map[string]interface{} `json:"deleted_by"`
	Updated_by string `json:"updated_by"`
	Name string `json:"name"`
	Is_external bool `json:"is_external"`
	Description string `json:"description"`
	Users []string `json:"users,omitempty"`
	Auto_type map[string]interface{} `json:"auto_type"`
	Channel_count int `json:"channel_count,omitempty"`
	Handle string `json:"handle"`
	Date_delete int `json:"date_delete"`
	Is_subteam bool `json:"is_subteam"`
	Date_create int `json:"date_create"`
	Id string `json:"id"`
	Team_id string `json:"team_id"`
	Created_by string `json:"created_by"`
	Auto_provision bool `json:"auto_provision"`
}

// Objsconversation represents the Objsconversation schema from the OpenAPI specification
type Objsconversation struct {
}

// Objsenterpriseuser represents the Objsenterpriseuser schema from the OpenAPI specification
type Objsenterpriseuser struct {
	Enterprise_name string `json:"enterprise_name"`
	Id string `json:"id"`
	Is_admin bool `json:"is_admin"`
	Is_owner bool `json:"is_owner"`
	Teams []string `json:"teams"`
	Enterprise_id string `json:"enterprise_id"`
}

// Objsteam represents the Objsteam schema from the OpenAPI specification
type Objsteam struct {
	Deleted bool `json:"deleted,omitempty"`
	Enterprise_id string `json:"enterprise_id,omitempty"`
	Msg_edit_window_mins int `json:"msg_edit_window_mins,omitempty"`
	Locale string `json:"locale,omitempty"`
	Name string `json:"name"`
	Over_storage_limit bool `json:"over_storage_limit,omitempty"`
	Primary_owner Objsprimaryowner `json:"primary_owner,omitempty"`
	Sso_provider map[string]interface{} `json:"sso_provider,omitempty"`
	Avatar_base_url string `json:"avatar_base_url,omitempty"`
	Icon Objsicon `json:"icon"`
	Discoverable map[string]interface{} `json:"discoverable,omitempty"`
	Is_assigned bool `json:"is_assigned,omitempty"`
	Has_compliance_export bool `json:"has_compliance_export,omitempty"`
	Id string `json:"id"`
	Archived bool `json:"archived,omitempty"`
	Date_create int `json:"date_create,omitempty"`
	Domain string `json:"domain"`
	Enterprise_name string `json:"enterprise_name,omitempty"`
	Email_domain string `json:"email_domain"`
	Is_over_storage_limit bool `json:"is_over_storage_limit,omitempty"`
	Messages_count int `json:"messages_count,omitempty"`
	Over_integrations_limit bool `json:"over_integrations_limit,omitempty"`
	Pay_prod_cur string `json:"pay_prod_cur,omitempty"`
	External_org_migrations Objsexternalorgmigrations `json:"external_org_migrations,omitempty"`
	Created int `json:"created,omitempty"`
	Limit_ts int `json:"limit_ts,omitempty"`
	Is_enterprise int `json:"is_enterprise,omitempty"`
	Plan string `json:"plan,omitempty"`
}

// Objsreminder represents the Objsreminder schema from the OpenAPI specification
type Objsreminder struct {
	Recurring bool `json:"recurring"`
	Text string `json:"text"`
	Time int `json:"time,omitempty"`
	User string `json:"user"`
	Complete_ts int `json:"complete_ts,omitempty"`
	Creator string `json:"creator"`
	Id string `json:"id"`
}

// Objschannel represents the Objschannel schema from the OpenAPI specification
type Objschannel struct {
	Unlinked int `json:"unlinked,omitempty"`
	Created int `json:"created"`
	Num_members int `json:"num_members,omitempty"`
	Members []string `json:"members"`
	Is_thread_only bool `json:"is_thread_only,omitempty"`
	Is_general bool `json:"is_general,omitempty"`
	Is_archived bool `json:"is_archived,omitempty"`
	Last_read string `json:"last_read,omitempty"`
	Topic map[string]interface{} `json:"topic"`
	Previous_names []string `json:"previous_names,omitempty"`
	Is_read_only bool `json:"is_read_only,omitempty"`
	Is_member bool `json:"is_member,omitempty"`
	Name string `json:"name"`
	Pending_shared []string `json:"pending_shared,omitempty"`
	Unread_count int `json:"unread_count,omitempty"`
	Is_pending_ext_shared bool `json:"is_pending_ext_shared,omitempty"`
	Is_channel bool `json:"is_channel"`
	Is_shared bool `json:"is_shared"`
	Unread_count_display int `json:"unread_count_display,omitempty"`
	Accepted_user string `json:"accepted_user,omitempty"`
	Latest map[string]interface{} `json:"latest,omitempty"`
	Is_moved int `json:"is_moved,omitempty"`
	Is_frozen bool `json:"is_frozen,omitempty"`
	Is_mpim bool `json:"is_mpim"`
	Is_private bool `json:"is_private"`
	Is_org_shared bool `json:"is_org_shared"`
	Purpose map[string]interface{} `json:"purpose"`
	Creator string `json:"creator"`
	Is_non_threadable bool `json:"is_non_threadable,omitempty"`
	Id string `json:"id"`
	Priority float64 `json:"priority,omitempty"`
	Name_normalized string `json:"name_normalized"`
}
