package diskstuff

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
)

func first(m map[string]string) (string, string, bool) {
  for k, v := range m {
    return k, v, true
  }
  return "", "", false
}

func foo() {
	fmt.Printf("\n")
}

func asserter(test bool) bool {
  if !test {
    breakout("", true)
  }
  return !test
}

func assertMsg(test bool, msg string) {
  if !test {
    breakout(msg, false)
  }
}

func assert(test bool) {
  if !test {
    breakout("", false)
  }
}

func breakout(msg string, quiet bool) {
  if !quiet {
    fmt.Printf("  ------------ BREAKOUT!! %v !!\n", msg)
  }
}
