package mocker

import "testing"

type Mocker[T any] interface {
	Mock(t *testing.T) T
}
