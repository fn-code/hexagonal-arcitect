package storage

// Device return device data
type Device struct {
	DeviceID    string `json:"device_id,omitempty"`
	DeviceName  string `json:"device_name,omitempty"`
	DeviceToken string `json:"device_token,omitempty"`
}
