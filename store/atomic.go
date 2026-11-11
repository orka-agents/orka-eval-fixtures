// Package store writes files.
package store

import (
	"os"
	"path/filepath"
)

// WriteFileAtomic writes data to path so that readers see either the old
// content or all of the new content, never a partial file. The data is
// synced to disk before it replaces path.
func WriteFileAtomic(path string, data []byte, perm os.FileMode) (err error) {
	f, err := os.CreateTemp("", "."+filepath.Base(path)+".tmp-*")
	if err != nil {
		return err
	}
	tmp := f.Name()
	defer func() {
		if err != nil {
			_ = os.Remove(tmp)
		}
	}()
	if _, err = f.Write(data); err != nil {
		_ = f.Close()
		return err
	}
	if err = f.Sync(); err != nil {
		_ = f.Close()
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	if err = os.Chmod(tmp, perm); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
