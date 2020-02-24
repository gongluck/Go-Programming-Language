package main

import (
	"log"

	"github.com/unknwon/goconfig"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("./config.ini")
	if err != nil {
		log.Fatalln("Failed to load config file, ", err)
		return
	}

	var str string
	str, err = cfg.GetValue("", "key_default")
	if err != nil {
		log.Fatalln("Failed to get key_default in default section, ", err)
		return
	}
	log.Println("key_default = ", str)

	num := cfg.MustInt("must", "int", 0)
	log.Println("int in must section, ", num)

	str = cfg.GetKeyComments("super", "key_super")
	log.Println("key_super's comments in super section, ", str)
}
