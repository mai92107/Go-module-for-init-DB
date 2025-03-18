package init

func InitAll(){
	LoadEnvFromJSON("config.json")
	InitDB()
}