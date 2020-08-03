package infrastructure

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "../interfaces/database"
  "fmt"
)

type SqlHandler struct {
  Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
  conn, err := sql.Open("mysql", "root:@tcp(db:3306)/go_api_practice")
  if err != nil {
    panic(err.Error)
  }
  sqlHandler := new(SqlHandler)
  sqlHandler.Conn = conn
  fmt.Print(conn)
  return sqlHandler
}


func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
  res := SqlResult{}
  result, err := handler.Conn.Exec(statement, args...)
  if err != nil {
    return res, err
  }
  res.Result = result
  return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
    fmt.Print(*handler)
    rows, err := handler.Conn.Query(statement, args...)
    if err != nil {
        fmt.Print(rows)
        return new(SqlRow), err
    }
    row := new(SqlRow)
    row.Rows = rows
    return row, nil
}

type SqlResult struct {
  Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
    return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
    return r.Result.RowsAffected()
}

type SqlRow struct {
    Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
    return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
    return r.Rows.Next()
}

func (r SqlRow) Close() error {
    fmt.Print("hogehoge")
    return r.Rows.Close()
}

