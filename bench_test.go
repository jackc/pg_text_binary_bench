package pg_text_binary_bench_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgconn"
)

func createTestData(b *testing.B, conn *pgconn.PgConn) {
	mrr := conn.Exec(context.Background(), `
create temporary table strings (
	a text not null,
	b text not null,
	c text not null,
	d text not null,
	e text not null
);

insert into strings(a, b, c, d, e)
select '0123456789012345', '0123456789012345', '0123456789012345', '0123456789012345', '0123456789012345'
from generate_series(1, 1000);
`)
	err := mrr.Close()
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkSelectText(b *testing.B) {
	conn, err := pgconn.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close(context.Background())

	createTestData(b, conn)

	formats := []struct {
		name string
		code int16
	}{
		{"text", 0},
		{"binary", 1},
	}
	for _, format := range formats {
		b.Run(fmt.Sprintf("%s format", format.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rr := conn.ExecParams(context.Background(), `select a, b, c, d, e from strings`, nil, nil, nil, []int16{format.code})
				_, err = rr.Close()
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
