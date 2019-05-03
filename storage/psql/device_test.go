package psql

import (
	"context"
	"testing"
)

func TestGetDevice(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db := dbConnection(ctx, t)

	defer func() {
		if err := db.Close(); err != nil {
			t.Errorf("failed close database connection: %v\n", err)
		}
	}()

	devs, err := db.GetDevice()
	if err != nil {
		t.Errorf("failed get device data:%v\n", err)
	}
	if devs == nil {
		t.Error("failed get device data is empty")
	}

}

func TestGetDeviceByID(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name string
		val  string
		err  error
	}{
		{name: "with device id", val: "1107239f3645765fdgfgfsdf1107"},
		{name: "invalid device id", val: "1107239f3645765fdgfgfsdf110710020kk", err: ErrDataNotFound},
		{name: "empty device id", val: "", err: ErrDataNotFound},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db := dbConnection(ctx, t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Errorf("failed close database connection: %v\n", err)
		}
	}()

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dev, err := db.GetDeviceByID(tc.val)
			if tc.err != nil {
				if tc.err != err {
					t.Errorf("failed get device id, expect %v got %v\n", tc.err, err)
				}
				return
			}
			if dev == nil {
				t.Error("failed get device id, device id is empty")
			}
		})

	}

}
