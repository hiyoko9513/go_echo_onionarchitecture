package util

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
)

type ID string

var defaultEntropySource *ulid.MonotonicEntropy

func init() {
	defaultEntropySource = ulid.Monotonic(rand.Reader, 0)
}

func NewULID() ID {
	return ID(fmt.Sprint(ulid.MustNew(ulid.Timestamp(time.Now()), defaultEntropySource)))
}
