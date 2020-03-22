package utils

var GET = "get"
var POST = "post"

type Url struct{
	Route string "route"
	Method string "method"
}

type Config struct {
	
	Base_Url   string "base_url"
	Urls []Url "address"
}
