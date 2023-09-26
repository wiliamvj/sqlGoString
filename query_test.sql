--transform: query_test
SELECT * FROM test_table

--name test :one --transform: query_test2
SELECT * FROM test_table2
WHERE id = $1

--name test :one --transform: query_test3 --other-anotation
SELECT * FROM test_table3 t
WHERE id = $1
AND t.name = $2
