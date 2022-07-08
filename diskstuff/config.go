package diskstuff

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
	"fmt"
)

func ShowConfig() {
	fmt.Printf("\n")
}

type Config struct {
  Dryrun bool
  NoClobber bool
  Verbose bool
}

var theConfig *Config

func init() {
  theConfig = &Config{}
}

func NewConfig(dryrun bool) *Config {
  return &Config{
    Dryrun: dryrun,
  }
}

func (c *Config) SetSafe(safe bool) {
  c.NoClobber = safe
}

func (c *Config) SetVerbose(verbose bool) {
  c.Verbose = verbose
}

func ConfigDryRun() bool {
  return theConfig.Dryrun
}

func ConfigNoClobber() bool {
  return theConfig.NoClobber
}

func ConfigVerbose() bool {
  return theConfig.Verbose
}

