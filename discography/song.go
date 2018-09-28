package discography

type Song struct {
	Name        string   `yaml:"name"`        // FIXME: Or Title?
	Composition string   `yaml:"composition"` // 作曲
	Arrangement string   `yaml:"arrangement"` // 編曲
	Lyrics      string   `yaml:"lyrics"`      // 作詞
	Topics      []string `yaml:"topics"`
}
