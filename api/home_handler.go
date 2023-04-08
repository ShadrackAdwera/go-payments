package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

type Profile struct {
	Iss      string   `json:"iss"`
	Sub      string   `json:"sub"`
	Aud      []string `json:"aud"`
	Iat      int64    `json:"iat"`
	Exp      int64    `json:"exp"`
	Azp      string   `json:"azp"`
	Scope    string   `json:"scope"`
	Name     string   `json:"name"`
	Nickname string   `json:"nickname"`
	Picture  string   `json:"picture"`
	Sid      string   `json:"sid"`
}

func (s *Server) home(ctx *gin.Context) {
	var profileData Profile
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	if profile == nil {
		ctx.JSON(http.StatusAccepted, gin.H{"message": "please login to proceed"})
		return
	}

	mapstructure.Decode(profile, &profileData)

	ctx.JSON(http.StatusAccepted, gin.H{"user": profileData})
}
