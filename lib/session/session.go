package session

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fasthttp/session"
)

// Data describes data which saved to session
type Data struct {
	UserID    uint      `json:"userId"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	LoginTime time.Time `json:"loginTime"`
}

// SetAuth will set session auth for user
func SetAuth(store session.Storer, data *Data) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	store.Set("auth", string(bytes))
	return nil
}

// GetAuth will provide authentication info for logged user
func GetAuth(store session.Storer) (interface{}, error) {
	authData := store.Get("auth")

	if authData != nil {
		str := fmt.Sprintf("%v", authData)
		bytes := []byte(str)

		var data Data

		err := json.Unmarshal(bytes, &data)

		if err != nil {
			return nil, err
		}

		return data, nil
	}

	return authData, nil
}
