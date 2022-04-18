package main

import (
	"CTF-backend/internal/mysql"
	"fmt"
	"time"
)

func main() {
	//model := new(mysql.User)
	//mysql.InitDB()
	//mysql.DB.AutoMigrate(model)

	//nowTomorrow := time.Now().AddDate(0, 0, -7)
	//zeroTime := time.Date(nowTomorrow.Year(), nowTomorrow.Month(), nowTomorrow.Day(),
	//	0, 0, 0, 0, nowTomorrow.Location())
	//
	//mysql.DB.Create(mysql.UserDailyData{
	//	UserID:      6,
	//	UserRanking: 2,
	//	UserPoint:   9,
	//	RankingDate: zeroTime,
	//})

	var userQuestionMessages []mysql.UserDailyData
	now := time.Now()
	zeroTime := time.Date(now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0, now.Location())
	sevenDaysAgo := zeroTime.AddDate(0, 0, -6)
	mysql.DB.Where("ranking_date > ?", sevenDaysAgo).Where("user_id = ?", 6).Find(&userQuestionMessages)
	fmt.Printf("%s", userQuestionMessages)
}
