package server

import "net/http/cookiejar"

var CookieJar, _ = cookiejar.New(nil)
