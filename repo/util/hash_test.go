package util

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestHash(t *testing.T) {
	i := Murmur3([]byte(uuid.NewString()))
	fmt.Println(i)
	fmt.Println(To62Str(i))
}
