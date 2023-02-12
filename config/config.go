package config

type Configuration struct {
	App   App   `mapstructure:"app"·json:"app"·yaml:"app"`
	Mysql Mysql `mapstructure:"mysql"·json:"mysql"·yaml:"mysql"`
	Log   Log   `mapstructure:"log" json:"log" yaml:"log"`
	Jwt   Jwt   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

type App struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}

type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Database     string `mapstructure:"database" json:"database" yaml:"database"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Charset      string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdleConns string `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns string `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
}

type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	RootDir    string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine   bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"` // MB
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`    // day
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type Jwt struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl int64  `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"` // token 有效期（秒）
}
