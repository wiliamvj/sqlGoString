package sqlGoString

import (
  "bufio"
  "errors"
  "fmt"
  "os"
  "path/filepath"
  "strings"
)

var (
  TAG = "--transform: "
)

func scanSQLFiles() ([]string, error) {
  var sqlFiles []string

  err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err
    }
    if !info.IsDir() && filepath.Ext(path) == ".sql" {
      sqlFiles = append(sqlFiles, path)
    }
    return nil
  })

  if err != nil {
    return nil, fmt.Errorf("Erro ao ler os arquivos: %v", err)
  }

  return sqlFiles, nil
}

func transformSQLFiles(queryName string, sqlFiles []string) (string, error) {
  if len(sqlFiles) == 0 {
    return "", fmt.Errorf("No sql files found")
  }
  var sqlStatements []string

  for _, filePath := range sqlFiles {
    file, err := os.Open(filePath)
    if err != nil {
      return "", err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var isTransformBlock bool

    for scanner.Scan() {
      line := scanner.Text()

      if strings.Contains(line, TAG+queryName) {
        isTransformBlock = true
      } else if isTransformBlock && line == "" {
        isTransformBlock = false
      } else if isTransformBlock {
        sqlStatements = append(sqlStatements, line)
      }
    }

    if err := scanner.Err(); err != nil {
      return "", err
    }
  }

  if len(sqlStatements) <= 0 {
    return "", errors.New("No SQL statements found for " + queryName)
  }

  return strings.Join(sqlStatements, "\n"), nil
}
