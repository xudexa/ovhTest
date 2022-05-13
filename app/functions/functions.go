package functions

import (
	"strings"

	"github.com/gofrs/uuid"
)

func NewUUID() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}

func RemoveDuplicate(xs *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *xs {
		if strings.HasPrefix(x, "-") {
			x = x[1:]
		}
		if !found[strings.ToLower(x)] {
			found[strings.ToLower(x)] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}
