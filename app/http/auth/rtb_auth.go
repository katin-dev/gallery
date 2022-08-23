package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type (
	RtbAuth struct {
		Url string
	}
	RtbUser struct {
		ID          string   `json:"id"`
		Permissions []string `json:"permissions"`
	}
)

func NewRtbAuth(url string) *RtbAuth {
	return &RtbAuth{url}
}

func (a *RtbAuth) Auth(ctx *gin.Context) {
	// ctx.Writer.Header().Set("Location", o.Endpoint.AuthURL)
	authorization := ctx.Request.Header.Get("Authorization")
	if len(authorization) == 0 {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("Authorization required"))
	}

	authorizationParts := strings.Split(authorization, " ")
	if len(authorizationParts) != 2 {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("Authorization required"))
	}

	httpClient := http.Client{}

	url := a.Url + "/api/v1/user?_with=permissions"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", authorization)
	// req.Header.Set("X-Proxy-Client-Id", "gallery-client")

	res, err := httpClient.Do(req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("Unable to check authorization"))
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("Authorization failed"))
		return
	}

	resContent, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("Unable to read authorization response"))
		return
	}

	user := RtbUser{}
	if err = json.Unmarshal(resContent, &user); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("Unable to unmarshall authorization response"))
		return
	}

	ctx.Set("user", user)
}
