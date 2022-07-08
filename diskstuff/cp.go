package diskstuff

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

func cp(src, dest string) (int64, error) {
  mb := NewMultiBlast(dest)
  return mb.cp(src, dest)
}

