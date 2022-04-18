package mysql

import (
	"time"
)

type Question struct {
	QuestionID            int       `json:"questionID,omitempty"`
	QuestionTitle         string    `json:"questionTitle,omitempty"`
	Type                  string    `json:"type,omitempty"`
	Describe              string    `json:"describe,omitempty"`
	Point                 int       `json:"point,omitempty"`
	IsCollection          bool      `json:"isCollection,omitempty"`
	SuccessfulPersonCount int       `json:"successfulPersonCount,omitempty"`
	TryPersonCount        int       `json:"tryPersonCount,omitempty"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

type User struct {
	UserID    string    `json:"UserID"`
	Name      string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IfLogin   bool      `json:"ifLogin"`
}

type UserQuestionMessage struct {
	UserID     string
	QuestionID int
	Point      int
}
type UserDailyData struct {
	UserID      string    `json:"userID"`
	UserPoint   int       `json:"userPoint"`
	UserRanking int       `json:"userRanking"`
	RankingDate time.Time `json:"rankingDate"`
}
