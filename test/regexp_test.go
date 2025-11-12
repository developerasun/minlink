package test

import (
	"regexp"
	"strings"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTargetStringToFind(t *testing.T) {
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
