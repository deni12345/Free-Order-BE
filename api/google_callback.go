package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	people "google.golang.org/api/people/v1"
)

var (
	googleUserIdAPI = "https://people.googleapis.com/v1/people/me?personFields=names%2CemailAddresses"
)

func (s Server) GoogleCallBack(w http.ResponseWriter, r *http.Request) {
	data, err := s.getUserDataFromGoogle(r.Context(), r.FormValue("code"))
	if err != nil {
		logrus.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	logrus.Println(string(data))
	redirectURL := "http://localhost:3000/dashboard?token=" + "testing"
	http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
}

func (s Server) getUserDataFromGoogle(ctx context.Context, code string) ([]byte, error) {
	token, err := s.googleOauth.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	client, err := people.NewService(ctx, option.WithTokenSource(s.googleOauth.TokenSource(ctx, token)))
	if err != nil {
		return nil, fmt.Errorf("client wrong: %s", err.Error())
	}
	resp, err := client.People.Get("people/me").PersonFields("names,emailAddresses").Do()
	if err != nil {
		return nil, fmt.Errorf("client wrong: %s", err.Error())
	}

	return resp.MarshalJSON()
}
