package api

type PersonalCompletionRate struct {
	UserID string `json:"userID"`
	Count  int64  `json:"count"`
	All    int64  `json:"all"`
}
