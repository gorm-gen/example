package global

type Log struct {
	Path string `json:"path" yaml:"path" mapstructure:"path"`
}

type Mysql struct {
	Debug    bool   `json:"debug" yaml:"debug" mapstructure:"debug"`
	Host     string `json:"host" yaml:"host" mapstructure:"host"`
	Port     int    `json:"port" yaml:"port" mapstructure:"port"`
	User     string `json:"user" yaml:"user" mapstructure:"user"`
	Password string `json:"password" yaml:"password" mapstructure:"password"`
	DB       string `json:"db" yaml:"db" mapstructure:"db"`
}

type Conf struct {
	Log   Log   `json:"log" yaml:"log" mapstructure:"log"`
	Mysql Mysql `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
}
