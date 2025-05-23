package config

import (
	"log"
	"os"
	"path"
	"sync"

	"github.com/spf13/viper"

	"example/internal/global"
)

var once sync.Once

func Init(fileName string) {
	once.Do(func() {
		viper.SetConfigFile(fileName)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
		global.Config = &global.Conf{}
		if err := viper.Unmarshal(global.Config); err != nil {
			log.Fatal(err)
		}
	})
}

func InitAuto(filename string, n ...int) {
	_n := 200
	if len(n) > 0 && n[0] > 0 {
		_n = n[0]
	}

	i := 0
	var notFund bool
	var filePath string

	for {
		if i == 0 {
			filePath = "./"
		}
		if i > 0 {
			filePath += "../"
		}
		if _, err := os.Stat(path.Join(filePath, filename)); err != nil {
			if i == 0 {
				filePath = ""
			}
			if i == _n {
				notFund = true
				break
			}
			i++
			continue
		} else {
			Init(path.Join(filePath, filename))
			break
		}
	}

	if notFund {
		log.Fatal("Not fund config file.")
	}
}
