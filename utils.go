package utils

import (
	"fmt"
	"io"
	"io/fs"
	"os"
)

// CopyFile copies file from src to dst.
// If dst not exists create it with mode perm (before umask).
// If exists, then truncate it.
func CopyFile(src, dst string, perm fs.FileMode) error {

	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)

	return err
}

// ToStringSlice convert any type of slice to []string.
func ToStringSlice(s []fmt.Stringer) []string {

	r := make([]string, 0)

	for i := range s {
		r = append(r, s[i].String())
	}

	return r
}
