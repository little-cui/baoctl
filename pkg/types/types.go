package types

type Options struct {
	FilePath string
}

type Config struct {
	Xlsx Xlsx `yaml:"xlsx"`
}

type Xlsx struct {
	FilePath     string   `yaml:"path"`
	Password     string   `yaml:"password"`
	SheetN       int      `yaml:"sheetNumber"`
	FixedHeaders []string `yaml:"fixedHeaders"`
	FixedHeader  string   `yaml:"fixedHeader"`
	AddedHeaders []string `yaml:"addedHeaders"`
}
