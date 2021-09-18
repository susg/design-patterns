package adapter

type MacPort interface {
	InsertIntoLightningPort() string
}

type LightningPort struct{}

func (lp LightningPort) InsertIntoLightningPort() string {
	return "inserted into lightning port"
}
