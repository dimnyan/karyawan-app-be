package api

import (
	"database/sql"
	"fmt"
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
		Name:     "token",
		Value:    stringToken,
		HttpOnly: true,
	})
	// Token: Give Response
	var loginResponse LoginResponse
	loginResponse.Token = stringToken
	ctx.JSON(http.StatusOK, loginResponse)
}

type TokenRequest struct {
	Token string `json:"token"`
}

func (server *Server) CheckToken(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("token")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorMessage("Invalid Token on Cookie"))
		ctx.Abort()
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret-key"), nil
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		ctx.Abort()
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorMessage("Unauthorized Token"))
	}

	// Extract role (assuming it's a simple integer)
	roleValue, ok := claims["role"].(map[string]interface{})
	if ok {
		roleInt, ok := roleValue["Int64"].(float64)
		if ok {
			fmt.Println("Role:", int(roleInt))
		} else {
			fmt.Println("Error extracting role")
		}
	} else {
		fmt.Println("Error extracting role")
	}

}

func (server *Server) Logout(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().AddDate(-1, 0, 0),
		HttpOnly: true,
	})
}
