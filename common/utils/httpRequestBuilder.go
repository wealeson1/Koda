package main

import (
	"strings"
)

// commonHeader interns common header strings.
var commonHeader = map[string]string{
	"Accept":                    "Accept",
	"Accept-Charset":            "Accept-Charset",
	"Accept-Encoding":           "Accept-Encoding",
	"Accept-Language":           "Accept-Language",
	"Accept-Ranges":             "Accept-Ranges",
	"Cache-Control":             "Cache-Control",
	"Cc":                        "Cc",
	"Connection":                "Connection",
	"Content-Id":                "Content-Id",
	"Content-Language":          "Content-Language",
	"Content-Length":            "Content-Length",
	"Content-Transfer-Encoding": "Content-Transfer-Encoding",
	"Content-Type":              "Content-Type",
	"Cookie":                    "Cookie",
	"Date":                      "Date",
	"Etag":                      "Etag",
	"Expires":                   "Expires",
	"From":                      "From",
	"Host":                      "Host",
	"If-Modified-Since":         "If-Modified-Since",
	"If-None-Match":             "If-None-Match",
	"In-Reply-To":               "In-Reply-To",
	"Last-Modified":             "Last-Modified",
	"Location":                  "Location",
	"Message-Id":                "Message-Id",
	"Mime-Version":              "Mime-Version",
	"Pragma":                    "Pragma",
	"Received":                  "Received",
	"Return-Path":               "Return-Path",
	"Server":                    "Server",
	"Set-Cookie":                "Set-Cookie",
	"Subject":                   "Subject",
	"To":                        "To",
	"User-Agent":                "User-Agent",
	"X-Forwarded-For":           "X-Forwarded-For",
	"X-Powered-By":              "X-Powered-By",
}

// ParseHTTPRequestFirstLine parses "GET /foo HTTP/1.1" into its three parts.
func ParseHTTPRequestFirstLine(line string) (method, requestURI, proto string, ok bool) {
	firstSpaceIndex := strings.Index(line, " ")
	if firstSpaceIndex == -1 {
		return "", "", "", false
	}
	secondSpaceIndex := strings.LastIndex(line[firstSpaceIndex+1:], " ")
	httpVersion := "HTTP/1.1"
	method = line[:firstSpaceIndex]
	if secondSpaceIndex == -1 {
		requestURI = line[firstSpaceIndex+1:]
		return method, requestURI, httpVersion, true
	}
	secondSpaceIndex = secondSpaceIndex + 1 + firstSpaceIndex
	requestURI = line[firstSpaceIndex+1 : secondSpaceIndex]
	proto = line[secondSpaceIndex+1:]
	return method, requestURI, httpVersion, true
}

func main() {
	ParseHTTPRequestFirstLine("GET /foo")
}
