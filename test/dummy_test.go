package test

import (
	"log"
	"strings"

	"testing"
)

func TestAddressPrefix(t *testing.T) {
	t.Skip()
	address := "0x5a27fdA4A09B3feF34c5410de1c5F3497B8EBa11"
	after, found := strings.CutPrefix(address, "0x")

	if !found {
		t.Log("can't find eth address prefix: ", address)
		t.FailNow()
	}

	log.Println("sliced: ", after)
}
