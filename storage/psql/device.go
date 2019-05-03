package psql

import (
	"context"
	"log"
	"time"

	"github.com/fn-code/hexagonal-arcitec/storage"
)

func (db *postgresConn) getDevice(query string, args ...interface{}) ([]storage.Device, error) {
	ctx, cancel := context.WithTimeout(db.ctx, 10*time.Second)
	defer cancel()

	stmt, err := db.conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	res := make([]storage.Device, 0)
	for rows.Next() {
		addrs := storage.Device{}
		err := rows.Scan(
			&addrs.DeviceID,
			&addrs.DeviceName,
			&addrs.DeviceToken,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, addrs)
	}
	return res, nil
}

func (db *postgresConn) GetDevice() ([]storage.Device, error) {
	addrs, err := db.getDevice("SELECT device_id, device_name, device_token FROM device")
	if err != nil {
		return nil, err
	}
	return addrs, nil
}

func (db *postgresConn) GetDeviceByID(id string) (*storage.Device, error) {
	r, err := db.getDevice("SELECT device_id, device_name, device_token FROM device WHERE device_id=$1", id)
	if err != nil {
		return nil, err
	}
	dev := storage.Device{}
	if len(r) == 0 {
		return nil, ErrDataNotFound
	}
	dev = r[0]
	return &dev, nil
}

func (db *postgresConn) AddDevice(dev storage.Device) (string, error) {
	ctx, cancel := context.WithTimeout(db.ctx, 10*time.Second)
	defer cancel()

	stmt, err := db.conn.PrepareContext(ctx, "INSERT INTO device (device_id, device_name, device_token) VALUES($1, $2, $3) RETURNING device_id")
	if err != nil {
		return "", err
	}
	var id string
	err = stmt.QueryRowContext(db.ctx, dev.DeviceID, dev.DeviceName, dev.DeviceToken).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (db *postgresConn) DeleteDevice(id string) (string, error) {
	ctx, cancel := context.WithTimeout(db.ctx, 10*time.Second)
	defer cancel()

	stmt, err := db.conn.PrepareContext(ctx, "DELETE FROM device where device_id=$1 RETURNING device_id")
	if err != nil {
		return "", err
	}
	var ids string
	err = stmt.QueryRowContext(ctx, id).Scan(&ids)
	if err != nil {
		return "", err
	}
	return ids, nil
}

func (db *postgresConn) UpdateDevice(id string, dev storage.Device) error {
	ctx, cancel := context.WithTimeout(db.ctx, 10*time.Second)
	defer cancel()

	stmt, err := db.conn.PrepareContext(ctx, "UPDATE device SET device_name=$1, device_token=$2 WHERE device_id=$3 RETURNING device_id")
	if err != nil {
		log.Println(err)
		return err
	}
	var ids string
	err = stmt.QueryRowContext(ctx, dev.DeviceName, dev.DeviceToken, id).Scan(&ids)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
