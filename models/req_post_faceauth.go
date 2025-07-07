package models

import "time"

type PostFaceAuthRequest struct {
	Detected time.Time `json:"detected"`
	Status   string    `json:"status"`
}
