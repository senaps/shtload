package utils

var GET = "get"
var POST = "post"

type Url struct {
	Route  string `default:"route"`
	Method string `default:"method"`
}

type Config struct {
	BaseUrl string `default:"base_url"`
	Urls    []Url  `default:"address"`
}
