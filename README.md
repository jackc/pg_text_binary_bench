# PostgreSQL text vs binary format for text types

This test compares the performance of selecting text type values with the text format or the binary format. It creates a temporary table of 5 text columns populated with 1000 rows where each text value is 16 characters.

Run the test with `go test -bench=.`. The database connection can be controlled with the standard `PG*` environment variables.

The binary format seems to have a significant performance penalty.

```
jack@glados ~/dev/pg_text_binary_bench » psql -c 'select version()'
                                                      version
───────────────────────────────────────────────────────────────────────────────────────────────────────────────────
 PostgreSQL 12.3 on x86_64-apple-darwin19.4.0, compiled by Apple clang version 11.0.3 (clang-1103.0.32.59), 64-bit
(1 row)

jack@glados ~/dev/pg_text_binary_bench » go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/jackc/pg_text_binary_bench
BenchmarkSelectText/text_format-16         	    2989	    349023 ns/op
BenchmarkSelectText/binary_format-16       	    2767	    404590 ns/op
PASS
ok  	github.com/jackc/pg_text_binary_bench	2.514s
```
