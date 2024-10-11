package oauth2

// import (
// 	"context"
// 	"crypto/rand"
// 	"encoding/base64"
// 	"fmt"
// 	"github/lambda-microservice/config"
// 	"io/ioutil"
// 	"net/http"

// 	"github.com/sirupsen/logrus"
// 	"golang.org/x/oauth2"
// 	"golang.org/x/oauth2/google"
// )

// var googleOauthConfig = &oauth2.Config{
// 	RedirectURL:  "http://localhost:8080/api/public/auth/google/callback",
// 	ClientID:     config.Value.GoogleClientID,
// 	ClientSecret: config.Value.GoogleClientSecret,
// 	Scopes: []string{
// 		"https://www.googleapis.com/auth/userinfo.profile",
// 		"https://www.googleapis.com/auth/userinfo.email",
// 	},
// 	Endpoint: google.Endpoint,
// }

// func OauthGoogleLogin(w http.ResponseWriter, r *http.Request) {
// 	oauthState := generateStateOauthCookie(w)
// 	u := googleOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
// 	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
// }

// func generateStateOauthCookie(w http.ResponseWriter) string {
// 	b := make([]byte, 16)
// 	rand.Read(b)
// 	state := base64.URLEncoding.EncodeToString(b)

// 	return state
// }

// func OauthGoogleCallBack(w http.ResponseWriter, r *http.Request) {
// 	data, err := getUserDataFromGoogle(r.FormValue("code"))
// 	if err != nil {
// 		logrus.Println(err.Error())
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return
// 	}
// 	cookie := http.Cookie{
// 		Name:  "jwt-token",
// 		Value: string(data),
// 		Path:  "/",
// 	}
// 	http.SetCookie(w, &cookie)
// 	http.Redirect(w, r, "http://localhost:3000/dashboard", http.StatusPermanentRedirect)
// }

// func getUserDataFromGoogle(code string) ([]byte, error) {
// 	token, err := googleOauthConfig.Exchange(context.Background(), code)

// 	if err != nil {
// 		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
// 	}
// 	req, err := http.NewRequest(http.MethodGet, "https://people.googleapis.com/v1/people/me?personFields=names%2CemailAddresses", nil)
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
// 	if err != nil {
// 		return nil, err
// 	}
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
// 	}
// 	logrus.Infof("response  %v \n", res)

// 	defer res.Body.Close()
// 	contents, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed read response: %s", err.Error())
// 	}
// 	return contents, nil
// }
