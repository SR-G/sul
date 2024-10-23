package sul

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func IsFileAvailable(s string) bool {
	if stat, err := os.Stat(s); err == nil && !stat.IsDir() {
		return true
	} else {
		return false
	}
}

func WriteGZippedFile(filename string, data []byte, perm fs.FileMode) error {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(data); err != nil {
		return fmt.Errorf("can't write content to [%s] : %s", filename, err)
	}
	if err := gz.Close(); err != nil {
		return fmt.Errorf("can't write content to [%s] : %s", filename, err)
	}
	return os.WriteFile(filename, b.Bytes(), perm)
}

func ReadContentFromGZ(filename string) ([]byte, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	fz, err := gzip.NewReader(fi)
	if err != nil {
		return nil, err
	}
	defer fz.Close()

	s, err := io.ReadAll(fz)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func ReadContent(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("can't read file content from [%s] %s", filename, err)
	}
	return content, nil
}

func AddTrailingSlashIfNeeded(s string) string {
	if strings.HasSuffix(s, "/") {
		return s
	} else {
		return s + "/"
	}
}
