package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	// Cache it for the whole lifetime of the backend
	WeatherCache string
)

func main() {

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

		if WeatherCache != "" {
			fmt.Fprintln(w, WeatherCache)
			return
		}

		resp, err := http.Get("https://api.darksky.net/forecast/5375e67e1f87323c4b9b9ef18c29f5db/46.020714,7.749117")

		if err != nil {
			fmt.Println(err.Error())
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err.Error())
		}

		var weatherResponse WeatherResponse
		if err := json.Unmarshal(body, &weatherResponse); err != nil {
			log.Fatal(err)
		}

		switch weatherResponse.Currently.Icon {
		case "clear-day", "clear-night", "sleet", "wind", "partly-cloudy-day", "partly-cloudy-night", "cloudy":
			{
				weatherResponse.Currently.Icon = "sunny"
				break
			}
		case "rain", "snow", "hail", "thunderstorm", "tornado":
			{
				weatherResponse.Currently.Icon = "rainy"
				break
			}
		case "fog":
			{
				weatherResponse.Currently.Icon = "foggy"
				break
			}
		default:
			{
				weatherResponse.Currently.Icon = "sunny"
			}
		}

		value, _ := json.Marshal(weatherResponse.Currently)

		WeatherCache = string(value)

		fmt.Fprintln(w, string(value))
	})

	log.Fatal(http.ListenAndServe(":1569", nil))
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

type WeatherResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Currently struct {
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
	} `json:"currently"`
	Hourly struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                int     `json:"time"`
			Summary             string  `json:"summary"`
			Icon                string  `json:"icon"`
			PrecipIntensity     float64 `json:"precipIntensity"`
			PrecipProbability   float64 `json:"precipProbability"`
			Temperature         float64 `json:"temperature"`
			ApparentTemperature float64 `json:"apparentTemperature"`
			DewPoint            float64 `json:"dewPoint"`
			Humidity            float64 `json:"humidity"`
			Pressure            float64 `json:"pressure"`
			WindSpeed           float64 `json:"windSpeed"`
			WindGust            float64 `json:"windGust"`
			WindBearing         int     `json:"windBearing"`
			CloudCover          float64 `json:"cloudCover"`
			UvIndex             int     `json:"uvIndex"`
			Visibility          float64 `json:"visibility"`
			Ozone               float64 `json:"ozone"`
			PrecipType          string  `json:"precipType,omitempty"`
		} `json:"data"`
	} `json:"hourly"`
	Daily struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                        int     `json:"time"`
			Summary                     string  `json:"summary"`
			Icon                        string  `json:"icon"`
			SunriseTime                 int     `json:"sunriseTime"`
			SunsetTime                  int     `json:"sunsetTime"`
			MoonPhase                   float64 `json:"moonPhase"`
			PrecipIntensity             float64 `json:"precipIntensity"`
			PrecipIntensityMax          float64 `json:"precipIntensityMax"`
			PrecipIntensityMaxTime      int     `json:"precipIntensityMaxTime"`
			PrecipProbability           float64 `json:"precipProbability"`
			PrecipType                  string  `json:"precipType"`
			TemperatureHigh             float64 `json:"temperatureHigh"`
			TemperatureHighTime         int     `json:"temperatureHighTime"`
			TemperatureLow              float64 `json:"temperatureLow"`
			TemperatureLowTime          int     `json:"temperatureLowTime"`
			ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
			ApparentTemperatureHighTime int     `json:"apparentTemperatureHighTime"`
			ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
			ApparentTemperatureLowTime  int     `json:"apparentTemperatureLowTime"`
			DewPoint                    float64 `json:"dewPoint"`
			Humidity                    float64 `json:"humidity"`
			Pressure                    float64 `json:"pressure"`
			WindSpeed                   float64 `json:"windSpeed"`
			WindGust                    float64 `json:"windGust"`
			WindGustTime                int     `json:"windGustTime"`
			WindBearing                 int     `json:"windBearing"`
			CloudCover                  float64 `json:"cloudCover"`
			UvIndex                     int     `json:"uvIndex"`
			UvIndexTime                 int     `json:"uvIndexTime"`
			Visibility                  float64 `json:"visibility"`
			Ozone                       float64 `json:"ozone"`
			TemperatureMin              float64 `json:"temperatureMin"`
			TemperatureMinTime          int     `json:"temperatureMinTime"`
			TemperatureMax              float64 `json:"temperatureMax"`
			TemperatureMaxTime          int     `json:"temperatureMaxTime"`
			ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
			ApparentTemperatureMinTime  int     `json:"apparentTemperatureMinTime"`
			ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
			ApparentTemperatureMaxTime  int     `json:"apparentTemperatureMaxTime"`
		} `json:"data"`
	} `json:"daily"`
	Flags struct {
		Sources           []string `json:"sources"`
		MeteoalarmLicense string   `json:"meteoalarm-license"`
		NearestStation    float64  `json:"nearest-station"`
		Units             string   `json:"units"`
	} `json:"flags"`
	Offset int `json:"offset"`
}
