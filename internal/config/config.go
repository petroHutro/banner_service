package config

import "time"

// TO DO добавить чтения из файла и ENV
func NewFlags() Flags {
	return Flags{
		NetAddress: NetAddress{
			Host: "0.0.0.0",
			Port: 8080,
		},

		Logger: Logger{
			LoggerFilePath:  "file.log",
			LoggerFileFlag:  false,
			LoggerMultiFlag: false,
		},

		Storage: Storage{
			DatabaseDSN: "host=banner-bd port=5432 user=intern password=fyr8as4da6 dbname=banner_db sslmode=disable",
		},

		Caches: Caches{
			Url: "chaches:6379",
		},

		Token: Token{
			TokenSecretKey: "",
			TokenTime: TokenTime{
				Time:     3,
				TokenEXP: time.Hour * 3,
			},
		},
	}
}
