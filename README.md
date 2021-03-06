# Ginger - for go generate a micro service from data schema
Add following line in top of your data schema file 
// go:generate gigner $GOFILE
and also add following line above each data type you want to build api
// @ginger
then, type "go generate"

# Example data schema file
You can generate following data type from JSON-to-Go online tool with a sample of your data object in Json format.  (http://mholt.github.io/json-to-go/)

Ginger will generate main.go for you, only need a dataType.go in the current folder. 

Please add following 2 lines in your data schema
    Ginger_Created    int32  `json:"ginger_created,omitempty"`
    Ginger_Id         int32  `json:"ginger_id,omitempty"`

For example, User Schema
```go
//go:generate ginger $GOFILE
package main
//@ginger
type User struct {
    Ginger_Created    int32  `json:"ginger_created"`
    Ginger_Id         int32  `json:"ginger_id" gorm:"primary_key"`
	ID string `json:"id"`
	Name string `json:"name"`
	Deleted bool `json:"deleted"`
	Status interface{} `json:"status"`
	Color string `json:"color"`
	RealName string `json:"real_name"`
	Tz string `json:"tz"`
	TzLabel string `json:"tz_label"`
	TzOffset int `json:"tz_offset"`
	Profile struct {
		RealName string `json:"real_name"`
		RealNameNormalized string `json:"real_name_normalized"`
		Email string `json:"email"`
		Image24 string `json:"image_24"`
		Image32 string `json:"image_32"`
		Image48 string `json:"image_48"`
		Image72 string `json:"image_72"`
		Image192 string `json:"image_192"`
	} `json:"profile"`
	IsAdmin bool `json:"is_admin"`
	IsOwner bool `json:"is_owner"`
	IsPrimaryOwner bool `json:"is_primary_owner"`
	IsRestricted bool `json:"is_restricted"`
	IsUltraRestricted bool `json:"is_ultra_restricted"`
	IsBot bool `json:"is_bot"`
	HasFiles bool `json:"has_files"`
	Has2Fa bool `json:"has_2fa"`
}
```

#Demo
refer to github.com/mingderwang/goslack
