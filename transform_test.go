package sqlGoString

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
  testCases := []struct {
    queryName string
    expected  string
    isError   bool
  }{
    {
      queryName: "query_test",
      expected:  "SELECT * FROM test_table",
      isError:   false,
    },
    {
      queryName: "query_test2",
      expected:  "SELECT * FROM test_table2\nWHERE id = $1",
      isError:   false,
    },
    {
      queryName: "query_test3",
      expected:  "SELECT * FROM test_table3 t\nWHERE id = $1\nAND t.name = $2",
      isError:   false,
    },
    {
      queryName: "non_existing_query",
      expected:  "",
      isError:   true,
    },
  }

  for _, tc := range testCases {
    t.Run(tc.queryName, func(t *testing.T) {
      sqlQuery, err := Transform(tc.queryName)

      if tc.isError {
        assert.Error(t, err)
      } else {
        assert.NoError(t, err)
        assert.Equal(t, tc.expected, sqlQuery)
      }
    })
  }
}
