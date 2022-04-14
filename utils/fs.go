package utils

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// IsFile returns if a path is a file
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) || !info.Mode().IsRegular() {
		return false
	}

	return true
}

// IsDir returns if a path is a directory
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) || !info.IsDir() {
		return false
	}

	return true
}

// Copy will copy a file from src to dst
func Copy(src, dst string) error {
	in, err := os.Open(filepath.Clean(src))
	if err != nil {
		return err
	}
	/* #nosec G307 */
	defer in.Close()

	out, err := os.Create(filepath.Clean(dst))
	if err != nil {
		return err
	}
	/* #nosec G307 */
	// defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Close()
}

// FileGetContents will get the content of a file
// stripping out the source map (if set)
func FileGetContents(inFile string) (string, error) {
	fi, err := os.Open(filepath.Clean(inFile))
	if err != nil {
		return "", err
	}
	/* #nosec G307 */
	defer fi.Close()

	b, err := ioutil.ReadAll(fi)
	if err != nil {
		return "", err
	}

	// remove sourcemaps if any
	re := regexp.MustCompile(`(?Ui)/\*#\s+sourceMappingURL.*\*/`)

	return re.ReplaceAllString(string(b), ""), nil
}
