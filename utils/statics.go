package utils

var GET = "get"
var POST = "post"

type Config struct {
	
	Base_Url   string "base_url"
	Urls [] struct{
		Name string "name"
		Route string "route"
		Method string "method"
	} "yaml: urls"
}
