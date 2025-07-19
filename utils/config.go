package utils

import "github.com/BurntSushi/toml"

type Config struct {
	Player struct {
		Speed       float64 `toml:"speed"`
		JumpBy      float64 `toml:"jump_by"`
		PaddleWdith float64 `toml:"paddle_wdith"`
	} `toml:"player"`

	Ball struct {
		Speed float64 `toml:"speed"`
	} `toml:"ball"`

	Brick struct {
		Level int `toml:"level"`
	} `toml:"bricks"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
