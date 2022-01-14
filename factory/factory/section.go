package factory

type section interface {
	description() string
}

type achievementSection struct{}

func (ps achievementSection) description() string {
	return "this is achievement section"
}

type albumSection struct{}

func (as albumSection) description() string {
	return "this is album section"
}

type personalInfoSection struct{}

func (pis personalInfoSection) description() string {
	return "this is personal info section"
}
