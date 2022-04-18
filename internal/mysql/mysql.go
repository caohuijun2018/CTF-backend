package mysql

import (
	"CTF-backend/internal/utils"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

var DB *gorm.DB

func mysqlDSNFromENV(prefix string, params ...string) string {
	dsn := utils.MustGetenv("DSN")
	if len(params) > 0 {
		dsn = dsn + "?" + strings.Join(params, "&")
	}
	return dsn
}

func InitDB() error {
	var err error
	dsn := mysqlDSNFromENV("", "parseTime=true")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return err
	}

	if sqlDB, err := DB.DB(); err != nil {
		return err
	} else {
		sqlDB.SetMaxIdleConns(10)

		sqlDB.SetMaxOpenConns(100)

		sqlDB.SetConnMaxLifetime(time.Hour)
		return nil
	}
}

func IsMissing(tx *gorm.DB) bool {
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return true
	} else if tx.Error == nil {
		return false
	} else {
		panic(tx.Error)
	}
}

func CalculateRank() {
	//每次用户完成题目后必须调用该函数，更新排名
	var userDailyRank, newRank []UserDailyData
	nowTomorrow := time.Now().AddDate(0, 0, 1)
	zeroTime := time.Date(nowTomorrow.Year(), nowTomorrow.Month(), nowTomorrow.Day(),
		0, 0, 0, 0, nowTomorrow.Location())

	if err := DB.Where("ranking_date = ?", zeroTime).Order("user_point desc").Find(&userDailyRank).Error; err != nil {
		panic(err)
	}
	for i, data := range userDailyRank {
		model := UserDailyData{
			UserID:      data.UserID,
			UserPoint:   data.UserPoint,
			RankingDate: zeroTime,
			UserRanking: i + 1,
		}
		newRank = append(newRank, model)
	}
	if err := DB.Save(&newRank).Error; err != nil {
		panic(err)
	}
}
