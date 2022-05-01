package relpaginator_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPaginator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Paginator Suite")
}
