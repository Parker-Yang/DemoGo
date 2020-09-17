package main

import (
	"github.com/google/uuid"
	"fmt"
	"strings"
	"testing"
)

func TestRandom(t *testing.T) {
	u := uuid.New()
	t.Log(u.String())
	all := strings.ReplaceAll(u.String(), "-", "")
	fmt.Printf("%T-----%v\n",all,all)
}
