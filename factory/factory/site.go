package factory

type socialNetworkSite interface {
	createProfile()
	addSection(section)
	getSections() []section
}

type baseSocialNetworkSite struct {
	sections []section
}

func (sn *baseSocialNetworkSite) addSection(s section) {
	sn.sections = append(sn.sections, s)
}

func (sn *baseSocialNetworkSite) getSections() []section {
	return sn.sections
}

type facebook struct {
	baseSocialNetworkSite
}

func (fb *facebook) createProfile() {
	fb.addSection(personalInfoSection{})
	fb.addSection(albumSection{})
}

type linkedin struct {
	baseSocialNetworkSite
}

func (l *linkedin) createProfile() {
	l.addSection(personalInfoSection{})
	l.addSection(achievementSection{})
}
