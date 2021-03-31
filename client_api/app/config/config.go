package config

var Env EnvVariables

func init()  {
	Env = newEnv()
	err := Env.load()
	if err != nil{
		panic(err)
	}
}