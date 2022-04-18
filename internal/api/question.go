package api

import (
	"CTF-backend/internal/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	failedEntity = Entity{
		Code:      int(OperateFail),
		Msg:       OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	successEntity = Entity{
		Code:      int(OperateOk),
		Msg:       OperateOk.String(),
		Success:   true,
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
)

func GetQuestionList(c *gin.Context) {
	entity := failedEntity
	var questions []mysql.Question
	if err := mysql.DB.Find(&questions).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = questions
	entity.Total = len(questions)
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func GetUserRecentRankData(c *gin.Context) {
	//获取某个人的最近题目排名数据
	entity := failedEntity
	var userDailyData []mysql.UserDailyData
	now := time.Now()
	zeroTime := time.Date(now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0, now.Location())
	sevenDaysAgo := zeroTime.AddDate(0, 0, -6)
	userId := c.Param("id")
	err := mysql.DB.Where("created_at > ?", sevenDaysAgo).Where("user_id=?", userId).Find(&userDailyData).Error
	if err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	entity = successEntity
	entity.Data = userDailyData
	entity.Total = len(userDailyData)
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return
}

func GetCompletionRate(c *gin.Context) {
	//获取某个人的题目完成率
	entity := failedEntity
	var count, all int64
	id := c.Param("id")
	if err := mysql.DB.Model(mysql.UserQuestionMessage{}).Where("user_id= ?", id).Count(&count).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if err := mysql.DB.Model(mysql.Question{}).Count(&all).Error; err != nil {
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	resp := PersonalCompletionRate{
		UserID: id,
		Count:  count,
		All:    all,
	}
	entity = successEntity
	entity.Data = resp
	c.JSON(http.StatusOK, gin.H{"entity": entity})
	return

}
