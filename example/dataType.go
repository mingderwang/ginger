//go:generate ginger $GOFILE
package service

//@ginger
type SlackUser struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Deleted  bool        `json:"deleted"`
	Status   interface{} `json:"status"`
	Color    string      `json:"color"`
	RealName string      `json:"real_name"`
	Tz       string      `json:"tz"`
	TzLabel  string      `json:"tz_label"`
	TzOffset int         `json:"tz_offset"`
	Profile  struct {
		RealName           string `json:"real_name"`
		RealNameNormalized string `json:"real_name_normalized"`
		Email              string `json:"email"`
		Image24            string `json:"image_24"`
		Image32            string `json:"image_32"`
		Image48            string `json:"image_48"`
		Image72            string `json:"image_72"`
		Image192           string `json:"image_192"`
	} `json:"profile"`
	IsAdmin           bool `json:"is_admin"`
	IsOwner           bool `json:"is_owner"`
	IsPrimaryOwner    bool `json:"is_primary_owner"`
	IsRestricted      bool `json:"is_restricted"`
	IsUltraRestricted bool `json:"is_ultra_restricted"`
	IsBot             bool `json:"is_bot"`
	HasFiles          bool `json:"has_files"`
	Has2Fa            bool `json:"has_2fa"`
}
