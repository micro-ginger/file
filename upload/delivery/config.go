package delivery

type config struct {
	Files []fileConfig

	filesMap map[string]fileConfig
}

func (c *config) initialize() {
	c.filesMap = make(map[string]fileConfig)
	for _, f := range c.Files {
		c.filesMap[f.Type] = f
	}
}

type fileConfig struct {
	Type    string
	MinSize int64
	MaxSize int64
}
