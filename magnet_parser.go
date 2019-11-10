package main

import (
	"log"
	"net/url"
	"strings"
)

func parseValue(value string) string {

	decodedValue, err := url.QueryUnescape(value)
	if err != nil {
		log.Fatal(err)
	}
	return decodedValue
}

func parse(link string) map[string]string {

	droppedPrefix := strings.TrimPrefix(link, "magnet:?")
	params := make(map[string]string)
	split := strings.Split(droppedPrefix, "&")

	for _, param := range split {
		items := strings.Split(param, "=")

		if len(items) != 2 {
			continue
		}

		val := parseValue(items[1])
		params[items[0]] = val
	}

	return params
}

func infoHash(xt string) string {
	return strings.TrimPrefix(xt, "urn:btih:")
}
