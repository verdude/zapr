package zapr

import "testing"

func TestTopLevelLogs(t *testing.T) {
  Init()
  I("Initialized successfully")
  t.Log("Initialized and logged successfully")
}
