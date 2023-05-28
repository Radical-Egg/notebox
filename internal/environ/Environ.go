package environ

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)
type ConfigNoteBox struct {
	NoteLoc     string `yaml:"data_dir" env:"NOTEBOX_DATADIR" env-default:"./data"`
	DB_TYPE		string `yaml:"DB_TYPE" env:"NOTEBOX_DB_TYPE" env-default:"sqlite3"`
	DB_NAME		string `yaml:"DB_NAME" env:"NOTEBOX_DB_NAME" env-default:"notes.db"`
}

func LoadConfigs() ConfigNoteBox {
	var cfg ConfigNoteBox
	
	err := cleanenv.ReadConfig("./build/config.yml", &cfg)
	if err != nil {
		log.Fatalln(err)
		log.Fatalln("Unable to load environment configurations")
	}

	return cfg
}