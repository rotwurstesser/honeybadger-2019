package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	// Cache it for the whole lifetime of the backend
	WeatherCache string
)

func main() {

	// Initial data to prevent api problems
	WeatherCache = "{\"Currently\":{\"time\":1566713236,\"summary\":\"Clear\",\"icon\":\"sunny\",\"nearestStormDistance\":0,\"precipIntensity\":0,\"precipProbability\":0,\"precipType\":\"\",\"temperature\":48.38,\"apparentTemperature\":44.42,\"dewPoint\":29.91,\"humidity\":0.49,\"pressure\":1021.79,\"windSpeed\":8.91,\"windGust\":8.91,\"windBearing\":48,\"cloudCover\":0,\"uvIndex\":0,\"visibility\":8.17,\"ozone\":298.2},\"Future\":{\"time\":1566943200,\"summary\":\"Possible light rain in the evening.\",\"icon\":\"rainy\",\"nearestStormDistance\":0,\"precipIntensity\":0.0046,\"precipProbability\":0.5,\"precipType\":\"rain\",\"temperature\":0,\"apparentTemperature\":0,\"dewPoint\":38.95,\"humidity\":0.58,\"pressure\":1018.83,\"windSpeed\":2.26,\"windGust\":7.16,\"windBearing\":140,\"cloudCover\":0.94,\"uvIndex\":4,\"visibility\":8.53,\"ozone\":286.7}}"

	http.HandleFunc("/instagram", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)

		resp, err := http.Get("https://www.instagram.com/explore/tags/" + r.URL.Query().Get("hashtag") + "/?__a=1")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var response InstagramResponse
		if err := json.Unmarshal(body, &response); err != nil {
			log.Fatal(err)
			return
		}

		var srcResponse []string

		for i := 0; i < 9; i++ {
			srcResponse = append(srcResponse, response.Graphql.Hashtag.EdgeHashtagToTopPosts.Edges[i].Node.ThumbnailResources[4].Src)
		}

		value, _ := json.Marshal(srcResponse)

		fmt.Fprintln(w, string(value))

	})

	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)

		fmt.Fprintln(w, WeatherCache)
	})

	go StartWeatherThread()

	log.Fatal(http.ListenAndServe(":1569", nil))
}

func StartWeatherThread() {
	for {
		GetWeather()
		time.Sleep(time.Duration(time.Hour * 4))
	}
}

func GetWeather() {
	resp, err := http.Get("https://api.darksky.net/forecast/5375e67e1f87323c4b9b9ef18c29f5db/46.020714,7.749117")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var weatherResponse WeatherResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		log.Fatal(err)
		return
	}

	weatherResponse.Currently.Icon = GetIcon(weatherResponse.Currently.Icon)

	weatherResponse.Daily.Data[3].Icon = GetIcon(weatherResponse.Daily.Data[3].Icon)

	weatherResponseStruct := WeatherResponseStruct{Currently: weatherResponse.Currently, Future: weatherResponse.Daily.Data[3]}

	value, _ := json.Marshal(weatherResponseStruct)

	WeatherCache = string(value)
}

func GetIcon(icon string) string {
	switch icon {
	case "clear-day", "clear-night", "sleet", "wind", "partly-cloudy-day", "partly-cloudy-night", "cloudy":
		{
			return "sunny"
			break
		}
	case "rain", "hail", "thunderstorm", "tornado":
		{
			return "rainy"
			break
		}
	case "fog":
		{
			return "foggy"
			break
		}
	case "snow": {
		return "snow"
		break
	}
	}
	return "sunny"
}

// Just for testing purposes
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type InstagramResponse struct {
	Graphql struct {
		Hashtag struct {
			ID                 string `json:"id"`
			Name               string `json:"name"`
			AllowFollowing     bool   `json:"allow_following"`
			IsFollowing        bool   `json:"is_following"`
			IsTopMediaOnly     bool   `json:"is_top_media_only"`
			ProfilePicURL      string `json:"profile_pic_url"`
			EdgeHashtagToMedia struct {
				Count    int `json:"count"`
				PageInfo struct {
					HasNextPage bool   `json:"has_next_page"`
					EndCursor   string `json:"end_cursor"`
				} `json:"page_info"`
			} `json:"edge_hashtag_to_media"`
			EdgeHashtagToTopPosts struct {
				Edges []struct {
					Node struct {
						Typename           string `json:"__typename"`
						ID                 string `json:"id"`
						EdgeMediaToCaption struct {
							Edges []struct {
								Node struct {
									Text string `json:"text"`
								} `json:"node"`
							} `json:"edges"`
						} `json:"edge_media_to_caption"`
						Shortcode          string `json:"shortcode"`
						EdgeMediaToComment struct {
							Count int `json:"count"`
						} `json:"edge_media_to_comment"`
						TakenAtTimestamp int `json:"taken_at_timestamp"`
						Dimensions       struct {
							Height int `json:"height"`
							Width  int `json:"width"`
						} `json:"dimensions"`
						DisplayURL  string `json:"display_url"`
						EdgeLikedBy struct {
							Count int `json:"count"`
						} `json:"edge_liked_by"`
						EdgeMediaPreviewLike struct {
							Count int `json:"count"`
						} `json:"edge_media_preview_like"`
						Owner struct {
							ID string `json:"id"`
						} `json:"owner"`
						ThumbnailSrc       string `json:"thumbnail_src"`
						ThumbnailResources []struct {
							Src          string `json:"src"`
							ConfigWidth  int    `json:"config_width"`
							ConfigHeight int    `json:"config_height"`
						} `json:"thumbnail_resources"`
						IsVideo              bool   `json:"is_video"`
						AccessibilityCaption string `json:"accessibility_caption"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_hashtag_to_top_posts"`
			EdgeHashtagToContentAdvisory struct {
				Count int           `json:"count"`
				Edges []interface{} `json:"edges"`
			} `json:"edge_hashtag_to_content_advisory"`
		} `json:"hashtag"`
	} `json:"graphql"`
}

type WeatherResponseStruct struct {
	Currently WeatherStruct
	Future    WeatherStruct
}

type WeatherStruct struct {
	Time                 int     `json:"time"`
	Summary              string  `json:"summary"`
	Icon                 string  `json:"icon"`
	NearestStormDistance int     `json:"nearestStormDistance"`
	PrecipIntensity      float64 `json:"precipIntensity"`
	PrecipProbability    float64 `json:"precipProbability"`
	PrecipType           string  `json:"precipType"`
	Temperature          float64 `json:"temperature"`
	ApparentTemperature  float64 `json:"apparentTemperature"`
	DewPoint             float64 `json:"dewPoint"`
	Humidity             float64 `json:"humidity"`
	Pressure             float64 `json:"pressure"`
	WindSpeed            float64 `json:"windSpeed"`
	WindGust             float64 `json:"windGust"`
	WindBearing          int     `json:"windBearing"`
	CloudCover           float64 `json:"cloudCover"`
	UvIndex              int     `json:"uvIndex"`
	Visibility           float64 `json:"visibility"`
	Ozone                float64 `json:"ozone"`
}

type WeatherResponse struct {
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	Timezone  string        `json:"timezone"`
	Currently WeatherStruct `json:"currently"`
	Hourly    WeatherStruct `json:"hourly"`
	Daily     struct {
		Summary string          `json:"summary"`
		Icon    string          `json:"icon"`
		Data    []WeatherStruct `json:"data"`
	} `json:"daily"`
	Flags struct {
		Sources           []string `json:"sources"`
		MeteoalarmLicense string   `json:"meteoalarm-license"`
		NearestStation    float64  `json:"nearest-station"`
		Units             string   `json:"units"`
	} `json:"flags"`
	Offset int `json:"offset"`
}
