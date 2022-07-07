package diskstuff

/* Copyright © 2022 sparksb -- MIT (see LICENSE file) */
/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "fmt"
  "os"
  "path"
)

func Blast(folder string) {
  var count int

  if len(folder) == 0 {
    folder = "."
  }

  count = countGoFiles(folder)
  fmt.Printf("Count: %d\n", count)
  fmt.Printf("Count2: %d\n", countGoFiles2(folder, 0))
}

func countGoFiles2(folder string, count int) int {
  files, err := os.ReadDir(folder)
  if err != nil {
    fmt.Printf("Error in countGoFiles: %v\n", err)
    return count
  }
  for _, file := range files {
    if file.IsDir() {
      count = countGoFiles2(folder+"/"+file.Name(), count)
    }
    if path.Ext(file.Name()) == ".go" {
      count += 1
    }
  }
  return count
}

func countGoFiles(folder string) (count int) {
  files, err := os.ReadDir(folder)
  if err != nil {
    fmt.Printf("Error in countGoFiles: %v\n", err)
    return 0
  }

  for _, file := range files {
    if path.Ext(file.Name()) == ".go" {
      count += 1
    }
  }

  return count
}
