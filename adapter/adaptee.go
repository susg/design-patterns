package adapter

type WindowsPort interface {
	InsertIntoUSBPort() string
}

type USBPort struct {}

func (u USBPort) InsertIntoUSBPort() string {
	return "inserted into USB port"
}
