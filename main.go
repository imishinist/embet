package embet

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	// DirPerm is defined to write embed files.
	// Because embed file system is read-only.
	DirPerm  = os.FileMode(0755)
	FilePerm = os.FileMode(0644)
)

// WriteEmbedFiles write embed files to the real file system.
func WriteEmbedFiles(dir embed.FS, prefix, dest string) error {
	walkFunc := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		dpath := filepath.Join(dest, path)
		if d.IsDir() {
			if err := os.Mkdir(dpath, DirPerm); err != nil {
				return err
			}
			return nil
		}

		data, err := fs.ReadFile(dir, path)
		if err != nil {
			return err
		}
		if err := os.WriteFile(dpath, data, FilePerm); err != nil {
			return err
		}
		return nil
	}
	if err := fs.WalkDir(dir, prefix, walkFunc); err != nil {
		return err
	}
	return nil
}
