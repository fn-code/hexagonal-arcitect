package psql

import (
	"context"
	"os"
	"testing"

	"github.com/fn-code/hexagonal-arcitec/storage"
)

func dbConnection(ctx context.Context, t *testing.T) storage.Database {
	db, err := NewPostgres(ctx, os.Getenv("PGUSERNAME"), os.Getenv("PGPASSWORD"), os.Getenv("PGPORT"), os.Getenv("PGDATABASE"))
	if err != nil {
		t.Errorf("failed open new connection: %v\n", err)
	}
	return db
}

func TestNewPostgres(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := NewPostgres(ctx, os.Getenv("PGUSERNAME"), os.Getenv("PGPASSWORD"), os.Getenv("PGPORT"), os.Getenv("PGDATABASE"))
	if err != nil {
		t.Errorf("failed open new connection: %v\n", err)
	}
	if err := db.Close(); err != nil {
		t.Errorf("failed close open connection:%v\n", err)
	}
}

// func BenchmarkMaxOpenCons(b *testing.B) {
// 	tc := []struct {
// 		name string
// 		val  int
// 	}{
// 		{name: "MaxOpenCons1", val: 1},
// 		{name: "MaxOpenCons2", val: 2},
// 		{name: "MaxOpenCons5", val: 5},
// 		{name: "MaxOpenCons10", val: 10},
// 		{name: "MaxOpenConsUnlimited", val: 0},
// 	}
// 	for _, tb := range tc {
// 		b.Run(tb.name, func(bb *testing.B) {
// 			db := NewPostgres(tb.val)
// 			conn, err := db.Connect()
// 			if err != nil {
// 				b.Errorf("error running database : %v\n", err)
// 			}
// 			defer conn.Close()
// 			bb.RunParallel(func(pb *testing.PB) {
// 				for pb.Next() {
// 					_, err := conn.GetDevice()
// 					if err != nil {
// 						b.Errorf("error get list of device : %v\n", err)
// 					}
// 				}
// 			})
// 		})
// 	}
// }
