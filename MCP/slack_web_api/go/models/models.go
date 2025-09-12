package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Objscomment represents the Objscomment schema from the OpenAPI specification
type Objscomment struct {
	Comment string `json:"comment"`
	Created int `json:"created"`
	Pinned_info Defspinnedinfo `json:"pinned_info,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	User string `json:"user"`
	Is_intro bool `json:"is_intro"`
	Is_starred bool `json:"is_starred,omitempty"`
	Id string `json:"id"`
	Num_stars int `json:"num_stars,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	Timestamp int `json:"timestamp"`
}

// Objsresponsemetadata represents the Objsresponsemetadata schema from the OpenAPI specification
type Objsresponsemetadata struct {
}

// Objsteam represents the Objsteam schema from the OpenAPI specification
type Objsteam struct {
	Has_compliance_export bool `json:"has_compliance_export,omitempty"`
	Name string `json:"name"`
	Over_integrations_limit bool `json:"over_integrations_limit,omitempty"`
	Pay_prod_cur string `json:"pay_prod_cur,omitempty"`
	Plan string `json:"plan,omitempty"`
	Enterprise_name string `json:"enterprise_name,omitempty"`
	Limit_ts int `json:"limit_ts,omitempty"`
	Domain string `json:"domain"`
	Enterprise_id string `json:"enterprise_id,omitempty"`
	Msg_edit_window_mins int `json:"msg_edit_window_mins,omitempty"`
	Date_create int `json:"date_create,omitempty"`
	Over_storage_limit bool `json:"over_storage_limit,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Is_assigned bool `json:"is_assigned,omitempty"`
	Locale string `json:"locale,omitempty"`
	Messages_count int `json:"messages_count,omitempty"`
	Id string `json:"id"`
	Is_over_storage_limit bool `json:"is_over_storage_limit,omitempty"`
	Primary_owner Objsprimaryowner `json:"primary_owner,omitempty"`
	Sso_provider map[string]interface{} `json:"sso_provider,omitempty"`
	Avatar_base_url string `json:"avatar_base_url,omitempty"`
	Created int `json:"created,omitempty"`
	External_org_migrations Objsexternalorgmigrations `json:"external_org_migrations,omitempty"`
	Is_enterprise int `json:"is_enterprise,omitempty"`
	Email_domain string `json:"email_domain"`
	Icon Objsicon `json:"icon"`
	Discoverable map[string]interface{} `json:"discoverable,omitempty"`
}

// Objsicon represents the Objsicon schema from the OpenAPI specification
type Objsicon struct {
	Image_44 string `json:"image_44,omitempty"`
	Image_68 string `json:"image_68,omitempty"`
	Image_88 string `json:"image_88,omitempty"`
	Image_default bool `json:"image_default,omitempty"`
	Image_102 string `json:"image_102,omitempty"`
	Image_132 string `json:"image_132,omitempty"`
	Image_230 string `json:"image_230,omitempty"`
	Image_34 string `json:"image_34,omitempty"`
}

// Objsmessage represents the Objsmessage schema from the OpenAPI specification
type Objsmessage struct {
	Text string `json:"text"`
	Thread_ts string `json:"thread_ts,omitempty"`
	Files []Objsfile `json:"files,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	Bot_id map[string]interface{} `json:"bot_id,omitempty"`
	Client_msg_id string `json:"client_msg_id,omitempty"`
	Parent_user_id string `json:"parent_user_id,omitempty"`
	Inviter string `json:"inviter,omitempty"`
	Username string `json:"username,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	File Objsfile `json:"file,omitempty"`
	Subscribed bool `json:"subscribed,omitempty"`
	Upload bool `json:"upload,omitempty"`
	Comment Objscomment `json:"comment,omitempty"`
	User string `json:"user,omitempty"`
	Attachments []map[string]interface{} `json:"attachments,omitempty"`
	Ts string `json:"ts"`
	Team string `json:"team,omitempty"`
	Icons map[string]interface{} `json:"icons,omitempty"`
	Is_delayed_message bool `json:"is_delayed_message,omitempty"`
	Display_as_bot bool `json:"display_as_bot,omitempty"`
	Reply_users []string `json:"reply_users,omitempty"`
	Reply_count int `json:"reply_count,omitempty"`
	Is_starred bool `json:"is_starred,omitempty"`
	User_profile Objsuserprofileshort `json:"user_profile,omitempty"`
	Purpose string `json:"purpose,omitempty"`
	Subtype string `json:"subtype,omitempty"`
	Name string `json:"name,omitempty"`
	Old_name string `json:"old_name,omitempty"`
	Bot_profile Objsbotprofile `json:"bot_profile,omitempty"`
	Latest_reply string `json:"latest_reply,omitempty"`
	User_team string `json:"user_team,omitempty"`
	Unread_count int `json:"unread_count,omitempty"`
	Permalink string `json:"permalink,omitempty"`
	Reply_users_count int `json:"reply_users_count,omitempty"`
	Last_read string `json:"last_read,omitempty"`
	Topic string `json:"topic,omitempty"`
	Source_team string `json:"source_team,omitempty"`
	Is_intro bool `json:"is_intro,omitempty"`
	Blocks []map[string]interface{} `json:"blocks,omitempty"` // This is a very loose definition, in the future, we'll populate this with deeper schema in this definition namespace.
	TypeField string `json:"type"`
}

// Objsteamprofilefield represents the Objsteamprofilefield schema from the OpenAPI specification
type Objsteamprofilefield struct {
	Is_hidden bool `json:"is_hidden,omitempty"`
	Label string `json:"label"`
	Options map[string]interface{} `json:"options,omitempty"`
	Ordering float64 `json:"ordering"`
	TypeField string `json:"type"`
	Hint string `json:"hint"`
	Id string `json:"id"`
}

// Objsteamprofilefieldoption represents the Objsteamprofilefieldoption schema from the OpenAPI specification
type Objsteamprofilefieldoption struct {
}

// Objschannel represents the Objschannel schema from the OpenAPI specification
type Objschannel struct {
	Unlinked int `json:"unlinked,omitempty"`
	Is_member bool `json:"is_member,omitempty"`
	Is_archived bool `json:"is_archived,omitempty"`
	Purpose map[string]interface{} `json:"purpose"`
	Previous_names []string `json:"previous_names,omitempty"`
	Creator string `json:"creator"`
	Is_mpim bool `json:"is_mpim"`
	Topic map[string]interface{} `json:"topic"`
	Is_non_threadable bool `json:"is_non_threadable,omitempty"`
	Is_pending_ext_shared bool `json:"is_pending_ext_shared,omitempty"`
	Is_channel bool `json:"is_channel"`
	Is_private bool `json:"is_private"`
	Latest map[string]interface{} `json:"latest,omitempty"`
	Is_general bool `json:"is_general,omitempty"`
	Id string `json:"id"`
	Is_read_only bool `json:"is_read_only,omitempty"`
	Is_frozen bool `json:"is_frozen,omitempty"`
	Pending_shared []string `json:"pending_shared,omitempty"`
	Is_org_shared bool `json:"is_org_shared"`
	Created int `json:"created"`
	Name string `json:"name"`
	Num_members int `json:"num_members,omitempty"`
	Priority float64 `json:"priority,omitempty"`
	Unread_count_display int `json:"unread_count_display,omitempty"`
	Accepted_user string `json:"accepted_user,omitempty"`
	Is_moved int `json:"is_moved,omitempty"`
	Unread_count int `json:"unread_count,omitempty"`
	Last_read string `json:"last_read,omitempty"`
	Members []string `json:"members"`
	Is_thread_only bool `json:"is_thread_only,omitempty"`
	Is_shared bool `json:"is_shared"`
	Name_normalized string `json:"name_normalized"`
}

// Objsuser represents the Objsuser schema from the OpenAPI specification
type Objsuser struct {
}

// Objsconversation represents the Objsconversation schema from the OpenAPI specification
type Objsconversation struct {
}

// Objsresources represents the Objsresources schema from the OpenAPI specification
type Objsresources struct {
	Excluded_ids []map[string]interface{} `json:"excluded_ids,omitempty"`
	Ids []map[string]interface{} `json:"ids"`
	Wildcard bool `json:"wildcard,omitempty"`
}

// Objsprimaryowner represents the Objsprimaryowner schema from the OpenAPI specification
type Objsprimaryowner struct {
	Email string `json:"email"`
	Id string `json:"id"`
}

// Objsuserprofile represents the Objsuserprofile schema from the OpenAPI specification
type Objsuserprofile struct {
	Status_text string `json:"status_text"`
	Always_active bool `json:"always_active,omitempty"`
	Status_default_text string `json:"status_default_text,omitempty"`
	Team string `json:"team,omitempty"`
	Is_app_user bool `json:"is_app_user,omitempty"`
	Skype string `json:"skype"`
	Title string `json:"title"`
	Avatar_hash string `json:"avatar_hash"`
	Bot_id string `json:"bot_id,omitempty"`
	Display_name_normalized string `json:"display_name_normalized"`
	Status_default_emoji string `json:"status_default_emoji,omitempty"`
	Last_avatar_image_hash string `json:"last_avatar_image_hash,omitempty"`
	Api_app_id string `json:"api_app_id,omitempty"`
	Memberships_count int `json:"memberships_count,omitempty"`
	Pronouns string `json:"pronouns,omitempty"`
	User_id string `json:"user_id,omitempty"`
	Real_name string `json:"real_name"`
	Real_name_normalized string `json:"real_name_normalized"`
	Updated int `json:"updated,omitempty"`
	Phone string `json:"phone"`
	Status_expiration int `json:"status_expiration,omitempty"`
	Status_emoji string `json:"status_emoji"`
	Is_custom_image bool `json:"is_custom_image,omitempty"`
	Display_name string `json:"display_name"`
}

// Objsexternalorgmigrations represents the Objsexternalorgmigrations schema from the OpenAPI specification
type Objsexternalorgmigrations struct {
	Date_updated int `json:"date_updated"`
	Current []map[string]interface{} `json:"current"`
}

// Objsreaction represents the Objsreaction schema from the OpenAPI specification
type Objsreaction struct {
	Name string `json:"name"`
	Users []string `json:"users"`
	Count int `json:"count"`
}

// Objsreminder represents the Objsreminder schema from the OpenAPI specification
type Objsreminder struct {
	Time int `json:"time,omitempty"`
	User string `json:"user"`
	Complete_ts int `json:"complete_ts,omitempty"`
	Creator string `json:"creator"`
	Id string `json:"id"`
	Recurring bool `json:"recurring"`
	Text string `json:"text"`
}

// Objssubteam represents the Objssubteam schema from the OpenAPI specification
type Objssubteam struct {
	Name string `json:"name"`
	User_count int `json:"user_count,omitempty"`
	Users []string `json:"users,omitempty"`
	Date_create int `json:"date_create"`
	Description string `json:"description"`
	Updated_by string `json:"updated_by"`
	Is_usergroup bool `json:"is_usergroup"`
	Auto_provision bool `json:"auto_provision"`
	Auto_type map[string]interface{} `json:"auto_type"`
	Date_delete int `json:"date_delete"`
	Enterprise_subteam_id string `json:"enterprise_subteam_id"`
	Is_external bool `json:"is_external"`
	Deleted_by map[string]interface{} `json:"deleted_by"`
	Handle string `json:"handle"`
	Is_subteam bool `json:"is_subteam"`
	Date_update int `json:"date_update"`
	Id string `json:"id"`
	Prefs map[string]interface{} `json:"prefs"`
	Team_id string `json:"team_id"`
	Created_by string `json:"created_by"`
	Channel_count int `json:"channel_count,omitempty"`
}

// Defspinnedinfo represents the Defspinnedinfo schema from the OpenAPI specification
type Defspinnedinfo struct {
}

// Objspaging represents the Objspaging schema from the OpenAPI specification
type Objspaging struct {
	Pages int `json:"pages,omitempty"`
	Per_page int `json:"per_page,omitempty"`
	Spill int `json:"spill,omitempty"`
	Total int `json:"total"`
	Count int `json:"count,omitempty"`
	Page int `json:"page"`
}

// Objsfile represents the Objsfile schema from the OpenAPI specification
type Objsfile struct {
	Permalink string `json:"permalink,omitempty"`
	Preview string `json:"preview,omitempty"`
	Thumb_720_w int `json:"thumb_720_w,omitempty"`
	Is_starred bool `json:"is_starred,omitempty"`
	Editable bool `json:"editable,omitempty"`
	Thumb_1024 string `json:"thumb_1024,omitempty"`
	Image_exif_rotation int `json:"image_exif_rotation,omitempty"`
	Ims []string `json:"ims,omitempty"`
	Filetype string `json:"filetype,omitempty"`
	Thumb_720 string `json:"thumb_720,omitempty"`
	Thumb_480_w int `json:"thumb_480_w,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Thumb_720_h int `json:"thumb_720_h,omitempty"`
	Thumb_800_w int `json:"thumb_800_w,omitempty"`
	User string `json:"user,omitempty"`
	Comments_count int `json:"comments_count,omitempty"`
	Thumb_360_h int `json:"thumb_360_h,omitempty"`
	Num_stars int `json:"num_stars,omitempty"`
	Thumb_360 string `json:"thumb_360,omitempty"`
	Url_private string `json:"url_private,omitempty"`
	Thumb_360_w int `json:"thumb_360_w,omitempty"`
	Permalink_public string `json:"permalink_public,omitempty"`
	Thumb_960_h int `json:"thumb_960_h,omitempty"`
	Original_h int `json:"original_h,omitempty"`
	Username string `json:"username,omitempty"`
	Url_private_download string `json:"url_private_download,omitempty"`
	State string `json:"state,omitempty"`
	Last_editor string `json:"last_editor,omitempty"`
	Pinned_info Defspinnedinfo `json:"pinned_info,omitempty"`
	Thumb_160 string `json:"thumb_160,omitempty"`
	Mode string `json:"mode,omitempty"`
	Thumb_480 string `json:"thumb_480,omitempty"`
	Display_as_bot bool `json:"display_as_bot,omitempty"`
	Thumb_800_h int `json:"thumb_800_h,omitempty"`
	External_url string `json:"external_url,omitempty"`
	Pinned_to []string `json:"pinned_to,omitempty"`
	Thumb_tiny string `json:"thumb_tiny,omitempty"`
	Non_owner_editable bool `json:"non_owner_editable,omitempty"`
	Is_external bool `json:"is_external,omitempty"`
	Created int `json:"created,omitempty"`
	External_type string `json:"external_type,omitempty"`
	Thumb_1024_w int `json:"thumb_1024_w,omitempty"`
	Thumb_800 string `json:"thumb_800,omitempty"`
	Thumb_960_w int `json:"thumb_960_w,omitempty"`
	Id string `json:"id,omitempty"`
	Channels []string `json:"channels,omitempty"`
	Source_team string `json:"source_team,omitempty"`
	Title string `json:"title,omitempty"`
	Public_url_shared bool `json:"public_url_shared,omitempty"`
	Has_rich_preview bool `json:"has_rich_preview,omitempty"`
	Thumb_480_h int `json:"thumb_480_h,omitempty"`
	External_id string `json:"external_id,omitempty"`
	Timestamp int `json:"timestamp,omitempty"`
	Is_tombstoned bool `json:"is_tombstoned,omitempty"`
	Mimetype string `json:"mimetype,omitempty"`
	Date_delete int `json:"date_delete,omitempty"`
	Original_w int `json:"original_w,omitempty"`
	Shares map[string]interface{} `json:"shares,omitempty"`
	Is_public bool `json:"is_public,omitempty"`
	Updated int `json:"updated,omitempty"`
	Name string `json:"name,omitempty"`
	Pretty_type string `json:"pretty_type,omitempty"`
	Thumb_64 string `json:"thumb_64,omitempty"`
	User_team string `json:"user_team,omitempty"`
	Editor string `json:"editor,omitempty"`
	Reactions []Objsreaction `json:"reactions,omitempty"`
	Thumb_1024_h int `json:"thumb_1024_h,omitempty"`
	Size int `json:"size,omitempty"`
	Thumb_960 string `json:"thumb_960,omitempty"`
	Thumb_80 string `json:"thumb_80,omitempty"`
}

// Objsenterpriseuser represents the Objsenterpriseuser schema from the OpenAPI specification
type Objsenterpriseuser struct {
	Is_admin bool `json:"is_admin"`
	Is_owner bool `json:"is_owner"`
	Teams []string `json:"teams"`
	Enterprise_id string `json:"enterprise_id"`
	Enterprise_name string `json:"enterprise_name"`
	Id string `json:"id"`
}

// Objsuserprofileshort represents the Objsuserprofileshort schema from the OpenAPI specification
type Objsuserprofileshort struct {
	Is_ultra_restricted bool `json:"is_ultra_restricted"`
	Display_name_normalized string `json:"display_name_normalized,omitempty"`
	Name string `json:"name"`
	Real_name string `json:"real_name"`
	Is_restricted bool `json:"is_restricted"`
	Team string `json:"team"`
	Display_name string `json:"display_name"`
	Real_name_normalized string `json:"real_name_normalized,omitempty"`
	Avatar_hash string `json:"avatar_hash"`
	Image_72 string `json:"image_72"`
}

// Objsbotprofile represents the Objsbotprofile schema from the OpenAPI specification
type Objsbotprofile struct {
	App_id string `json:"app_id"`
	Deleted bool `json:"deleted"`
	Icons map[string]interface{} `json:"icons"`
	Id string `json:"id"`
	Name string `json:"name"`
	Team_id string `json:"team_id"`
	Updated int `json:"updated"`
}
