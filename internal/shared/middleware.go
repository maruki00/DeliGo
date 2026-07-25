package shared

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const API_AUTH = "http://127.0.0.1:8081/api/v1/auth"

func Request(body []byte) (bool, error) {
	req, err := http.NewRequest("POST", API_AUTH, bytes.NewReader(body))
	if err != nil {
		return false, nil
	}

	client := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200, nil
}
func GetToken(req *http.Request) string {
	authHeader := req.Header.Get("Authorization")
	strings.Split(authHeader, "")
}
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := GetToken(ctx.Request)
		body := fmt.Sprintf("{\"token\":\"%s\"}", token)
		if isAuth, err := Request([]byte(body)); !isAuth || err != nil {
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
