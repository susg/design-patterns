package adapter

type WindowsAdapter struct {
	port WindowsPort
}

func (wa WindowsAdapter) InsertIntoLightningPort() string {
	return wa.port.InsertIntoUSBPort()
}
