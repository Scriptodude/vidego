package config

import "time"

type Configuration struct {
	// The ip address to open the server on, defaults to localhost
	IpAddress string

	// The port to open the server on, defaults to 6911
	Port int

	// The timeout in seconds before a read timeout occurs, defaults to infinity
	ReadTimeout time.Duration

	// The timeout in seconds before a write timeout occurs, defaults to infinity
	WriteTimeout time.Duration

	// The base folder where to fetch the html from, must end with a trailing slash '/'
	HtmlBaseFolder string

	isInit bool
}
