#!/bin/sh

dbx schema -d pgx easypin.dbx .
dbx golang -d pgx -p dbx -t templates easypin.dbx .
( printf '%s\n' '//lint:file-ignore U1000,ST1012 generated file'; cat easypin.dbx.go ) > easypin.dbx.go.tmp && mv easypin.dbx.go.tmp easypin.dbx.go
gofmt -r "*sql.Tx -> tagsql.Tx" -w easypin.dbx.go
gofmt -r "*sql.Rows -> tagsql.Rows" -w easypin.dbx.go
perl -0777 -pi \
  -e 's,\t_ "github.com/jackc/pgx/v4/stdlib"\n\),\t_ "github.com/jackc/pgx/v4/stdlib"\n\n\t"storj.io/private/tagsql"\n\),' \
  easypin.dbx.go
perl -0777 -pi \
  -e 's/type DB struct \{\n\t\*sql\.DB/type DB struct \{\n\ttagsql.DB/' \
  easypin.dbx.go
perl -0777 -pi \
  -e 's/db = &DB\{\n\t\tDB: sql_db,/\tdb = &DB\{\n\t\tDB: tagsql.Wrap\(sql_db\),/' \
  easypin.dbx.go
