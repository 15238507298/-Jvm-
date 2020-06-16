package classpath

import (
	"os"
	_ "os"
	"strings"
)
import _ "strings"

const pathListSeparater = string(os.PathListSeparator)

type Entry interface {
	readClass(classname string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparater) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
