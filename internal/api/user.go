package api

import (
	"CTF-backend/internal/client"
	"CTF-backend/internal/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {

	entity := Entity{
		Code: int(OperateFail),
		Msg:  OperateFail.String(),
		Data: "Wrong username or password",
	}
	var user, user1 mysql.User
	if err := c.ShouldBindJSON(&user); err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	isExisted, err := client.IsExisted(user.UserID)
	if err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if !isExisted {
		entity.Msg = OperateFail.String()
		entity.Data = "The user does not exist"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if err := mysql.DB.Where(mysql.User{
		UserID:   user.UserID,
		Password: user.Password,
	}).First(&user1).Error; err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if user1.UserID != "" {
		entity = Entity{
			Code:    http.StatusOK,
			Success: true,
			Msg:     OperateOk.String(),
			Data:    "Login successfully",
		}
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}

func UserRegister(c *gin.Context) {
	entity := Entity{
		Code:  int(OperateFail),
		Msg:   OperateFail.String(),
		Total: 0,
	}
	var user mysql.User
	if err := c.ShouldBindJSON(&user); err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	isExisted, err := client.IsExisted(user.UserID)
	if err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if isExisted {
		entity.Msg = OperateFail.String()
		entity.Data = "The user is already existed"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if mysql.DB.Create(&user).Error != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err.Error() + "can not add user"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	} else {
		entity.Code = int(OperateOk)
		entity.Msg = OperateOk.String()
		entity.Success = true
		entity.Data = "Register successful"
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}
func UserEdit(c *gin.Context) {
	entity := Entity{
		Code:  int(OperateFail),
		Msg:   OperateFail.String(),
		Total: 0,
	}
	var user mysql.User
	if err := c.ShouldBindJSON(&user); err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	isExisted, err := client.IsExisted(user.UserID)
	if err != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if isExisted {
		entity.Msg = OperateFail.String()
		entity.Data = "The user is already existed"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	if mysql.DB.Save(&user).Error != nil {
		entity.Msg = OperateFail.String()
		entity.Data = err.Error() + "can not save user data"
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	} else {
		entity.Code = int(OperateOk)
		entity.Msg = OperateOk.String()
		entity.Success = true
		entity.Data = "Edit successfully"
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}
