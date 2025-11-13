package test

import (
	"io"
	"net/http"
	"regexp"
	"strings"
	// "time"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTargetStringToFind(t *testing.T) {
	t.Skip()
	assert := assert.New(t)
	target := "The current value of U.S. Dollar Index is 1555.22"
	paragraph := "Track the index more closely on the U.S. Dollar Index chart. What is U.S. Dollar Index highest value ever?U.S. Dollar Index reached its highest quote on Feb 25, 1985 — 164.720 USD. See more data on the U.S. Dollar Index chart.What is U.S. Dollar Index lowest value ever?The lowest ever quote of U.S. Dollar Index is 70.698 USD." + target

	re := regexp.MustCompile(target)
	match := re.FindString(paragraph)
	extracted := strings.Replace(target, "The current value of U.S. Dollar Index is", "", 1)

	t.Log("match: ", match)
	t.Log("extracted: ", extracted)

	assert.Equal(target, match)
}

func TestMultipleEndpoints(t *testing.T) {
	assert := assert.New(t)
	endpoints := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/users",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/comments/46",
		"https://jsonplaceholder.typicode.com/photos/22",
	}

	result := make(chan string, 5)
	for _, v := range endpoints {
		go func(v string) {
			request, _ := http.NewRequest(http.MethodGet, v, nil)
			client := http.Client{}
			response, _ := client.Do(request)
			defer response.Body.Close() // @dev prevent FD leak

			raw, _ := io.ReadAll(response.Body)
			data := string(raw)
			result <- data
		}(v)
	}

	var final []string
	for i := 0; i < len(endpoints); i++ {
		final = append(final, <-result)
	}

	t.Logf("final: %s", final)
	assert.Equal(len(endpoints), 5)
}
