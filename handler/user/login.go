package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	. "template/handler"
	"template/model"
	"template/pkg/auth"
	"template/pkg/errno"
	"template/pkg/token"
)

func Login(c *gin.Context) {
	// Binding the data with the user struct.
	var u model.UserModel
	if err := c.ShouldBindJSON(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	passwd, _ := auth.Encrypt("123456")
	d := model.UserModel{Password: passwd}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{ID: u.Id, Username: u.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})
}

func Get(c *gin.Context) {
	t, ok := c.Get("token")
	log.Infof("%v", ok)
	log.Infof("===%s", t.(*token.Context).Username)

	SendResponse(c, nil, "ok")
}
