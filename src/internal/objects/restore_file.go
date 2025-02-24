package objects

import "time"

//goland:noinspection GoNameStartsWithPackageName
type RestoreFile struct {
	TimeStamp time.Time                 `json:"time_stamp"`
	Items     map[string]*SimplyProduct `json:"items"`
}
