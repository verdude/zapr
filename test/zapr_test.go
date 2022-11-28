package test

import "testing"

func TestTopLevelLogs(t *testing.T) {
  Init()
  D("Debug")
  I("Info")
  W("Warn")
  E("Error")
  t.Log("Initialized and logged successfully")
}

func TestVLevel(t *testing.T) {
  // Idk how to do this properly
  V(5).I("Test")
}
