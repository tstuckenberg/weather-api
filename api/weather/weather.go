package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

//Conditions represents a weather conditions object
type Conditions struct {
	WindSpeed          float64 `json:"wind_speed"`
	TemperatureDegrees float64 `json:"temperature_degrees"`
}

//GetWeather fetches current weather for specified city
func GetWeather(city string) (Conditions, error) {

	//Create empty conditions object
	w := Conditions{}

	//Query for Yahoo weather API
	locationQuery := "select item.condition, wind from weather.forecast where woeid in (select woeid from geo.places where text='" + city + "' limit 1) and u='c'"
	//Format response URL
	locationURL := "http://query.yahooapis.com/v1/public/yql?q=" + url.QueryEscape(locationQuery) + "&format=json&env=store%3A%2F%2Fdatatables.org%2Falltableswithkeys"

	//Get response
	yahooResp, err := http.Get(locationURL)

	//If error or API is not response use open weather map
	if err != nil || yahooResp.StatusCode != 200 {

		//Query open weather map
		openweathermapResp, err := http.Get(`http://api.openweathermap.org/data/2.5/weather?q=` + city + `,AU&appid=2326504fb9b100bee21400190e4dbe6d`)

		if err != nil {
			return w, err
		}
		//Parse response to return
		body, err := ParseResponse(openweathermapResp)

		if err != nil {
			return w, err
		}
		//Format response
		w.FormatOpenweathermapResp(body)

		//Return
		return w, err

	}

	//Parse response to return
	body, err := ParseResponse(yahooResp)

	if err != nil {
		return w, err
	}

	//Format repsonse
	w.FormatYahooResponse(body)

	return w, err

}

//This was the wrong approach and stucts  should have been defined for responses; a bit of an experiment
//ParseResponse pareses http response body returning a JSON
func ParseResponse(resp *http.Response) (map[string]interface{}, error) {

	var data map[string]interface{}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return data, err
	}

	json.Unmarshal([]byte(body), &data)

	if data["query"].(map[string]interface{})["results"] == nil {

		return data, errors.New("Location not found")
	}

	return data, err

}

//FormatYahooResponse formats Yahoo Weather API response to weather conditionns
func (w *Conditions) FormatYahooResponse(body map[string]interface{}) Conditions {

	if windString, ok := body["query"].(map[string]interface{})["results"].(map[string]interface{})["channel"].(map[string]interface{})["wind"].(map[string]interface{})["speed"].(string); ok {

		windFloat, err := strconv.ParseFloat(windString, 64)
		if err != nil {
			return *w
		}

		w.WindSpeed = windFloat
	}
	if tempString, ok := body["query"].(map[string]interface{})["results"].(map[string]interface{})["channel"].(map[string]interface{})["item"].(map[string]interface{})["condition"].(map[string]interface{})["temp"].(string); ok {

		tempFloat, err := strconv.ParseFloat(tempString, 64)
		if err != nil {
			return *w
		}
		w.TemperatureDegrees = tempFloat
	}

	//return weather conditions
	return *w
}

//FormatOpenweathermapResp formats Yahoo Weather API response to weather conditionns
func (w *Conditions) FormatOpenweathermapResp(body map[string]interface{}) Conditions {
	fmt.Println(body)

	windString := body["wind"].(map[string]interface{})["speed"].(float64)
	tempString := body["main"].(map[string]interface{})["temp"].(float64)

	//As Open Weather Map return temperature in Kelvin we need to convert to degrees celsius
	kelvin := 273.15

	w.WindSpeed = windString
	w.TemperatureDegrees = (tempString - kelvin)

	//return weather conditions
	return *w
}
