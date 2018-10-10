package johnny5

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var SESSION_VARAIBLE = "AOC_SESSION"
var URL = "https://adventofcode.com/%d/day/%d/input"
var USER_AGENT = "johnny5/v0.0.1"

var cacheDir = ".johnny5"
var cacheKey = "johnny5.%04d%02d.cache"
var maxCacheAge = 5 * time.Minute

// GetInput for a problem by year and day.
func GetInput(year int, day int) (string, error) {
	input, err := getInputFromCache(year, day)
	if err != nil {
		input, err = fetchInput(year, day)
		if err != nil {
			return "", fmt.Errorf("failed to get input: %v", err)
		}
		err = cacheInput(year, day, input)
		if err != nil {
			return "", fmt.Errorf("failed to write to cache: %v", err)
		}
	}
	return strings.TrimSpace(string(input)), nil
}

func getCacheKey(year int, day int) string {
	var name = fmt.Sprintf(cacheKey, year, day)
	return fmt.Sprintf("%s/%s", cacheDir, name)
}

func getInputFromCache(year int, day int) ([]byte, error) {
	var filename = getCacheKey(year, day)
	info, err := os.Stat(filename)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to stat cache file: %v", err)
	}
	if time.Since(info.ModTime()) > maxCacheAge {
		return []byte{}, fmt.Errorf("cache file too old")
	}
	inputBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, fmt.Errorf("could not read cache file: %v", err)
	}
	return inputBytes, nil
}

func cacheInput(year int, day int, input []byte) error {
	os.Mkdir(cacheDir, 0700)
	var filename = getCacheKey(year, day)
	return ioutil.WriteFile(filename, input, 0700)
}

func fetchInput(year int, day int) ([]byte, error) {
	var url = fmt.Sprintf(URL, year, day)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to fetch input: %v", err)
	}

	var session = os.Getenv(SESSION_VARAIBLE)
	var cookie = &http.Cookie{Name: "session", Value: session}
	request.AddCookie(cookie)
	request.Header.Add("User-Agent", USER_AGENT)

	var client = &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to fetch input: %v", err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to fetch input: %v", err)
	}

	return content, nil
}
