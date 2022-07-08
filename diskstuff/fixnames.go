package diskstuff

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "io/fs"
  "log"
  "regexp"
  "strings"
  "syscall"
)

func FixNames(folder string, config *Config) error {
  theConfig = config
  if ConfigDryRun() {
    fmt.Printf("Dry Run.\n")
  }

	//fmt.Printf("\n")
  syscall.Umask(0)

  if len(folder) == 0 {
    folder = "."
  }

  re := regexp.MustCompile(`[^a-zA-Z0-9/.-]`)

  // Get file names
  filenames := map[string]string{}

  m := NewMultiBlast(folder)
  //fileSystem := os.DirFS(folder)
  err := fs.WalkDir(*m.Fs, folder, func(path string, d fs.DirEntry, err error) error {
    if path == ".git"  || path == ".idea" {
      return fs.SkipDir
    }
    processIt := false
    //fmt.Printf("path: %s, dirent: %v, error: %v\n", path, d, err)
    if strings.Contains(path, "1-Bit") {
      if path != "." {
        //fmt.Printf("!!!path: %s, dirent: %v, error: %v\n", path, d, err)
        processIt = true
        //return fs.SkipDir
      }
    }

    if processIt {
      newpath := strings.ReplaceAll(strings.ToLower(path), " ", "-")
      newpath = re.ReplaceAllString(newpath, "_")
      if d.IsDir() {
        if path != "." {
          //fmt.Printf("mkdir -p %s\n", newpath)
          //os.MkdirAll(newpath, 0755)
        }
      } else {
        //fmt.Printf("cp '%s' '%s' \n", path, newpath)
        filenames[path] = newpath
      }
    }

    return nil
  })
  if err != nil {
    log.Fatal(err)
  }

  err = m.cpMany(filenames)
  if err != nil {
    log.Fatal(err)
  }

  return err
}
