package utils

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
    "syscall"
)

type directory struct{}

func NewDirectory() *directory {
    return &directory{}
}

func (u *directory) CopyDirectory(scrDir, dest string) error {
    entries, err := os.ReadDir(scrDir)
    if err != nil {
        return err
    }
    for _, entry := range entries {
        sourcePath := filepath.Join(scrDir, entry.Name())
        destPath := filepath.Join(dest, entry.Name())

        var fileInfo os.FileInfo
        fileInfo, err = os.Stat(sourcePath)
        if err != nil {
            return err
        }

        stat, ok := fileInfo.Sys().(*syscall.Stat_t)
        if !ok {
            return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
        }

        switch fileInfo.Mode() & os.ModeType {
        case os.ModeDir:
            if err = u.CreateIfNotExists(destPath, 0755); err != nil {
                return err
            }
            if err = u.CopyDirectory(sourcePath, destPath); err != nil {
                return err
            }
        case os.ModeSymlink:
            if err = u.CopySymLink(sourcePath, destPath); err != nil {
                return err
            }
        default:
            if err = u.Copy(sourcePath, destPath); err != nil {
                return err
            }
        }

        if err = os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
            return err
        }

        var fInfo os.FileInfo
        fInfo, err = entry.Info()
        if err != nil {
            return err
        }

        isSymlink := fInfo.Mode()&os.ModeSymlink != 0
        if !isSymlink {
            if err = os.Chmod(destPath, fInfo.Mode()); err != nil {
                return err
            }
        }
    }
    return nil
}

func (u *directory) Copy(srcFile, dstFile string) error {
    out, err := os.Create(dstFile)
    if err != nil {
        return err
    }

    defer func(out *os.File) {
        _ = out.Close()
    }(out)

    var in *os.File
    in, err = os.Open(srcFile)
    defer func(in *os.File) {
        _ = in.Close()
    }(in)

    if err != nil {
        return err
    }

    _, err = io.Copy(out, in)
    if err != nil {
        return err
    }

    return nil
}

func (u *directory) Exists(filePath string) bool {
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        return false
    }

    return true
}

func (u *directory) CreateIfNotExists(dir string, perm os.FileMode) error {
    if u.Exists(dir) {
        return nil
    }

    if err := os.MkdirAll(dir, perm); err != nil {
        return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
    }

    return nil
}

func (u *directory) CopySymLink(source, dest string) error {
    link, err := os.Readlink(source)
    if err != nil {
        return err
    }
    return os.Symlink(link, dest)
}
