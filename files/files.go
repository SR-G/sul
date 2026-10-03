package files

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
)

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

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

func IsSymlinkAvailableOnFilesystem(symlink string) bool {
	info, err := os.Lstat(symlink)
	if err != nil {
		return false
	} else if info.Mode()&os.ModeSymlink != 0 {
		return true
	}
	return false
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

func WriteStringToFile(filename string, content string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		return err
	}
	return nil
}

func AddTrailingSlashIfNeeded(s string) string {
	if strings.HasSuffix(s, "/") {
		return s
	} else {
		return s + "/"
	}
}

func GetFileSize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}

	return info.Size(), nil
}

func ForceMkDirAllAndWipeBeforeIfNeeded(directory string, wipe bool) error {
	// Delete previous dir
	// - if --wipe is activated
	// - and if previously found on disk
	if wipe {
		if _, err := os.Stat(directory); err == nil {
			if err = os.RemoveAll(directory); err != nil {
				return fmt.Errorf("can't delete directory %q", directory)
			}
		}
	}

	// In all other situations, recreate the target dir
	// - directory was found on disk, with wipe activated, and has been deleted
	// - directory was found or not
	return ForceMkDirAll(directory)
}

func ForceMkDirAll(directory string) error {
	if err := os.MkdirAll(directory, 0o755); err != nil {
		return fmt.Errorf("unable to create output directory %q: %w", directory, err)
	}
	return nil
}
