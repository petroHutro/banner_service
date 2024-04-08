package config

func NewFlags() Flags {
	return Flags{

		NetAddress: NetAddress{
			Host: "localhost",
			Port: 8080,
		},

		Logger: Logger{
			LoggerFilePath:  "file.log",
			LoggerFileFlag:  false,
			LoggerMultiFlag: false,
		},
	}
}
