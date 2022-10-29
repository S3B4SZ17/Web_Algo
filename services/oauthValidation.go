package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/S3B4SZ17/Web_Algo/management"
	pbEmail "github.com/S3B4SZ17/Web_Algo/proto/email_user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ExtractToken(c *gin.Context) string {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.String(http.StatusForbidden, "No Authorization header provided")
		c.Abort()

	}
	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		c.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
		c.Abort()
		return ""
	}
	return token
}

func ValidateToken(c *gin.Context) error {
	token := ExtractToken(c)
	user, valid_token, _ := ExtractUser(&token)
	if user == nil {
		err := errors.New("User not authorized")
		return err
	}
	if !valid_token {
		// err := errors.New("User not authorized")
		// return err
	}
	return nil
}

func ExtractUser(token *string) (*pbEmail.EmailUser, bool, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + *token)
	if err != nil {
		management.Log.Error("Get: " + err.Error() + "\n")
		return nil, false, err
	}
	defer resp.Body.Close()

	user := &pbEmail.EmailUser{}

	response, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(response, &user)
	// Check your errors!
	fmt.Println(user.String())
	if user.String() == "" || err != nil {
		management.Log.Error("Error extracting user from response")
		return nil, false, err
	}
	valid := IsTokenValid(user.Email)
	return user, valid, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
