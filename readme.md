# sqlGoString

This package generates SQL queries in string for database queries

###### Install
```sh
go get github.com/wiliamvj/sqlGoString
```

### How to use:
put a tag mark with ```--transform: my_query```

Example:
```sql
--transform: my_query
SELECT * FROM user
```

```go
queryString, err := sqlGoString.Transform("my_query")
if err != nil {
  return err
}
db.Query(queryString)
```
sqlGoString will search for any ```.sql``` file and the ```my_query tag```.
You can have multiple sql files or use all queries in the same file.
