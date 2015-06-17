package main

import (
  "testing"
)

type tests struct {
  Input []uint
  Exptd uint
}

func TestAnswer(t *testing.T) {
  tsts := []tests{
    {Input: []uint{1, 0}, Exptd: 2},
    {Input: []uint{1, 2, 1}, Exptd: 2},
    {Input: []uint{1, 3, 0, 1}, Exptd: 2},
    {Input: []uint{2, 1, 3, 2}, Exptd: 2},
    {Input: []uint{2, 2, 4, 4, 0}, Exptd: 3},
    {Input: []uint{2, 3, 1, 4, 5, 0}, Exptd: 6},
  }
  for _, tt := range tsts {
    output := Answer(tt.Input)
    if output != tt.Exptd {
      t.Errorf("Answer(%d) = %d, expected %d", tt.Input, output, tt.Exptd)
    }
  }
}
