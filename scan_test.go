package sqlGoString

import (
  "errors"
  "path/filepath"
  "testing"

  "github.com/stretchr/testify/assert"
)

func mockErrorWalk(root string, walkFn filepath.WalkFunc) error {
  return errors.New("Simulated error during walk")
}

func TestScan(t *testing.T) {
  t.Run("Scan", func(t *testing.T) {
    scan, err := scanSQLFiles()
    assert.Nil(t, err)
    assert.NoError(t, err)
    assert.Equal(t, []string{"query_test.sql"}, scan)
  })

  t.Run("Transform", func(t *testing.T) {
    scan, err := transformSQLFiles("query_test", []string{"query_test.sql"})
    assert.Nil(t, err)
    assert.NoError(t, err)
    assert.Equal(t, "SELECT * FROM test_table", scan)
  })
}
