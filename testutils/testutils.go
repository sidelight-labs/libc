package testutils

import (
	"fmt"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strings"
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

func FileToString(fileName string) (string, error) {
	_, thisFile, _, _ := runtime.Caller(0)

	var (
		urlPath string
		err     error
	)
	if strings.Contains(thisFile, "vendor") {
		urlPath, err = filepath.Abs(path.Join(thisFile, "../../../../../..", "resources", "testdata", fileName))
	} else {
		urlPath, err = filepath.Abs(path.Join(thisFile, "../..", "resources", "testdata", fileName))
	}

	if err != nil {
		return "", err
	}

	Expect(urlPath).To(BeAnExistingFile())

	buf, err := ioutil.ReadFile(urlPath)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
