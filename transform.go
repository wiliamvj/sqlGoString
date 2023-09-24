package sqlGoString

import "fmt"

func Transform(queryName string) (string, error) {
  sqlFiles, err := scanSQLFiles()
  if err != nil {
    return "", fmt.Errorf("Error to scan SQL files: %v", err)
  }

  sqlQuery, err := transformSQLFiles(queryName, sqlFiles)
  if err != nil {
    return "", fmt.Errorf("Error to transform SQL files: %v", err)
  }

  return sqlQuery, nil
}
