package api

import (
	"context"
	"fmt"
	"github/free-order-be/config"
	"github/free-order-be/internal/auth"
	"github/free-order-be/models"
	"net/http"
	"strings"

	m "github/free-order-be/api/middleware"

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
		m.HttpJSONError(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	user, err := s.logic.CreateUser(r.Context(), data)
	if err != nil {
		logrus.Println(err.Error())
		m.HttpJSONError(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := auth.CreateToken(user)
	if err != nil {
		logrus.Println(err.Error())
		m.HttpJSONError(w, err.Error(), http.StatusInternalServerError)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "jwt-token", Value: token, Secure: true, HttpOnly: true, Domain: "localhost:3000"})
	http.Redirect(w, r, config.Values.RedirectURL, http.StatusPermanentRedirect)
}

func (s Server) getUserDataFromGoogle(ctx context.Context, code string) (*models.User, error) {
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

	return &models.User{
		Email:    resp.EmailAddresses[0].Value,
		UserName: resp.Names[0].DisplayName,
		GoogleID: strings.Split(resp.ResourceName, "/")[1],
		IsActive: true,
	}, nil
}
