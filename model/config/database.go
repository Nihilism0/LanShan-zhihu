package config

import (
	"fmt"
	"time"
)

type Database struct {
	Mysql   *Mysql   `mapstructure:"mysql" yaml:"mysql"`
	Redis   *Redis   `mapstructure:"redis" yaml:"redis"`
	MongoDB *MongoDB `mapstructure:"MongoDB" yaml:"MongoDB"`
}

type Mysql struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Port     string `mapstructure:"port" yaml:"port"`
	Db       string `mapstructure:"db" yaml:"db"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Charset  string `mapstructure:"charset" yaml:"charset"`

	ConnMaxIdleTime string `mapstructure:"connMaxIdleTime" yaml:"connMaxIdleTime"`
	ConnMaxLifeTime string `mapstructure:"connMaxLifeTime" yaml:"connMaxLifeTime"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns" yaml:"maxOpenConns"`
}

func (m *Mysql) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Australia%%2FMelbourne",
		m.Username,
		m.Password,
		m.Addr,
		m.Port,
		m.Db,
		m.Charset)
}

func (m *Mysql) GetConnMaxIdleTime() time.Duration {
	t, _ := time.ParseDuration(m.ConnMaxIdleTime)
	return t
}

func (m *Mysql) GetConnMaxLifeTime() time.Duration {
	t, _ := time.ParseDuration(m.ConnMaxLifeTime)
	return t
}

type Redis struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Port     string `mapstructure:"port" yaml:"port"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Db       int    `mapstructure:"db" yaml:"db"`
	PoolSize int    `mapstructure:"poolSize" yaml:"poolSize"`
}
type MongoDB struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Port     string `mapstructure:"port" yaml:"port"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
}
