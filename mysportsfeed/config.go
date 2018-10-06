package mysportsfeed

import "os"

var mySportsFeedAPIKey = os.Getenv("MSF_API_KEY")
var mySportsFeedPassword = os.Getenv("MSF_PASSWORD")

// EnvConfig is teh default configuration with the apitoken and password
// obtained from ENV variables.
func EnvConfig() Config {
	return Config{
		BaseURL:  "https://api.mysportsfeeds.com",
		APIToken: mySportsFeedAPIKey,
		Password: mySportsFeedPassword,
		Version:  V1_2,
		Sport:    "nfl",
	}
}
