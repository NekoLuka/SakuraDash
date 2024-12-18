package main

type Config struct {
	Title      string           `yaml:"title"`
	CacheIcons bool             `yaml:"cache_icons"`
	Categories []CategoryConfig `yaml:"categories"`
}

type CategoryConfig struct {
	Name      string       `yaml:"name"`
	Items     []ItemConfig `yaml:"items"`
	ViewGroup string       `yaml:"view_group"`
}

type ItemConfig struct {
	Name string `yaml:"name"`
	Icon string `yaml:"icon"`
	URL  string `yaml:"url"`
}
