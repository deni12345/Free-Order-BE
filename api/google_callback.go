package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	googleUserIdAPI = "https://people.googleapis.com/v1/people/me?personFields=names%2CemailAddresses"
)

func (s Server) GoogleCallBack(w http.ResponseWriter, r *http.Request) {
	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		logrus.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	logrus.Println(string(data))
	redirectURL := "http://localhost:3000/dashboard?token=" + "testing"
	http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	req, err := http.NewRequest(http.MethodGet, googleUserIdAPI, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}
