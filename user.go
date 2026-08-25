package models

import "time"

type User struct {
	ID                     string        `json:"id" bson:"_id"`
	Email                  string        `json:"email" bson:"email"`
	FirstName              string        `json:"firstName" bson:"firstName"`
	LastName               string        `json:"lastName" bson:"lastName"`
	Picture                string        `json:"picture" bson:"picture"`
	Role                   []Role        `json:"role" bson:"role"`
	Permissions            []string      `json:"permissions" bson:"permissions"`
	LastLoginIP            string        `json:"lastLoginIp" bson:"lastLoginIp"`
	LastLoginCity          string        `json:"lastLoginCity" bson:"lastLoginCity"`
	LastLoginCountry       string        `json:"lastLoginCountry" bson:"lastLoginCountry"`
	LastLoginCountryCode   string        `json:"lastLoginCountryCode" bson:"lastLoginCountryCode"`
	LastLoginLatitude      *float64      `json:"lastLoginLatitude" bson:"lastLoginLatitude"`
	LastLoginLongitude     *float64      `json:"lastLoginLongitude" bson:"lastLoginLongitude"`
	LastLoginTimestamp    *time.Time    `json:"lastLoginTimestamp" bson:"lastLoginTimestamp"`
	LastLoginDeviceType    string        `json:"lastLoginDeviceType" bson:"lastLoginDeviceType"`
	LastLoginUserAgent     string        `json:"lastLoginUserAgent" bson:"lastLoginUserAgent"`
	CreatedAt              time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt              time.Time     `json:"updatedAt" bson:"updatedAt"`
	Disabled               bool          `json:"disabled" bson:"disabled"`
	ContactEmail           string        `json:"contactEmail" bson:"contactEmail"`
	ContactPhone           string        `json:"contactPhone" bson:"contactPhone"`
	AllowNotifications     bool          `json:"allowNotifications" bson:"allowNotifications"`
	Theme                  string        `json:"theme" bson:"theme"`
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetLastLoginLocation() string {
	if u.LastLoginCity != "" && u.LastLoginCountry != "" {
		return u.LastLoginCity + ", " + u.LastLoginCountry
	} else if u.LastLoginCountry != "" {
		return u.LastLoginCountry
	}
	return "Unknown location"
}

func (u *User) HasGeolocationData() bool {
	return u.LastLoginIP != ""
}
