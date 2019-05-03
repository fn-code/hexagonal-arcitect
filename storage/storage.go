package storage

//Database service interface
type Database interface {
	UsersService
	DeviceService
	Close() error
}

// UsersService is users interface
type UsersService interface {
	AddUser(User, int8) (string, error)
	GetUsers() ([]UserInfo, error)
	GetUserByID(string) (*UserInfo, error)
	UpdateUser(string, UserInfo) error
	DeleteUser(string) (string, error)
	GetUserByUsername(string) (*User, error)
}

// DeviceService is device interface
type DeviceService interface {
	GetDevice() ([]Device, error)
	GetDeviceByID(string) (*Device, error)
	AddDevice(Device) (string, error)
	DeleteDevice(string) (string, error)
	UpdateDevice(string, Device) error
}
