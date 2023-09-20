package db

import "errors"

var ErrRecordNotFound = errors.New("sql: no rows in result set")
