package testutils

import (
	"fmt"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

type MockUuidGenerator struct {
	callNum int
}

func NewMockUuidGenerator() MockUuidGenerator {
	return MockUuidGenerator{0}
}
func (gen *MockUuidGenerator) GenerateMockUUID() string {
	uuid := fmt.Sprintf("FOOBAR%d", gen.callNum)
	gen.callNum += 1
	return uuid
}

func Testdata(t *testing.T, fileName string) string {
	t.Helper()

	var (
		urlPath string
		err     error

		_, thisFile, _, _ = runtime.Caller(0)
	)

	if strings.Contains(thisFile, "vendor") {
		urlPath, err = filepath.Abs(path.Join(thisFile, "../../../../../..", "resources", "testdata", fileName))
	} else {
		urlPath, err = filepath.Abs(path.Join(thisFile, "../..", "resources", "testdata", fileName))
	}

	Expect(err).NotTo(HaveOccurred())
	Expect(urlPath).To(BeAnExistingFile())

	buf, err := ioutil.ReadFile(urlPath)
	Expect(err).NotTo(HaveOccurred())
	return string(buf)
}
