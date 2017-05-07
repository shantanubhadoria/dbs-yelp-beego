package models

import (
	"bytes"
  "strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"
  "net/url"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Restaurant struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	ImageUrl  string  `json:"image_url"`
	Rating    float64 `json:"rating"`
	Phone     string  `json:"phone"`
}

type RestaurantResponse struct {
	Businesses []Restaurant `json:businesses`
	Total 		 int64        `json:total`
}

type AuthorizationResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int64    `json:"expires_in"`
	TokenType string   `json:"token_type"`
}

func init() {
	orm.RegisterModel(new(Restaurant))
}

// GetAllRestaurant retrieves all Restaurant matches certain condition. Returns empty list if
// no records exist
func GetAllRestaurant(keywords string, latitude string, longitude string) (resp RestaurantResponse, err error) {
	logs.GetLogger("models.Restaurant#GetAllRestaurant")
	// Get Auth token
	authResp := getAuthToken()

	request, _ := http.NewRequest("GET", "https://api.yelp.com/v3/businesses/search", nil)
	query := request.URL.Query()
	query.Add("term", keywords)
	query.Add("latitude", latitude)
	query.Add("longitude", longitude)
	request.URL.RawQuery = query.Encode()

	request1, _ := http.NewRequest("GET", request.URL.String(), nil)
	request1.Header.Add("Authorization", authResp.TokenType + " " + authResp.AccessToken)
	client := &http.Client{}
	response, err := client.Do(request1)
	body, _ := ioutil.ReadAll(response.Body)

	var restaurants RestaurantResponse
	json.Unmarshal(body, &restaurants)
	logs.Debug("%#v", restaurants)

	return restaurants, err
}

func getAuthToken() (r AuthorizationResponse) {
	// Get Auth token
	data := url.Values{}
	data.Set("client_id", "f6dRSiywqXr8joviPkvinw")
	data.Set("client_secret", "yHGQNjnWwnTYc1H3RfExUyq8HQ12a2yJmI7rfsLEnTt2bLEV9vnjzRFdeURHhdEU")
	data.Set("grant_type", "client_credentials")
	client := &http.Client{}
	request, _ := http.NewRequest("POST", "https://api.yelp.com/oauth2/token", bytes.NewBufferString(data.Encode()))
  request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)

	var authResp AuthorizationResponse
	json.Unmarshal(body, &authResp)

	return authResp
}
