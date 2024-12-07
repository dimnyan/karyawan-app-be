package api

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"karyawan-app-be/utils"
	"net/http"
	"time"
)

type Auth struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (server *Server) Login(ctx *gin.Context) {
	var auth Auth
	if err := ctx.ShouldBind(&auth); err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	// Get User
	username := sql.NullString{String: auth.Username, Valid: true}
	user, err := server.store.GetUserByUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusUnauthorized, utils.ErrorMessage("Credentials Not Match"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	// Check Password
	valid := utils.CheckPasswordHash(auth.Password, user.Password.String)
	if !valid {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorMessage("Credentials Not Match"))
		return
	}
	// Token: Create Claims
	claims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.RolesID,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
	}
	// Token: Sign Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Token: Sign Secret Key
	stringToken, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Error Signing Token"))
	}
	// Token : Set Cookie
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:  "token",
		Value: stringToken,
	})
	// Token: Give Response
	var loginResponse LoginResponse
	loginResponse.Token = stringToken
	ctx.JSON(http.StatusOK, loginResponse)
}
