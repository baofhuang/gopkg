package main

import (
	"fmt"
)

type Config struct {
	name   string
	age    int
	gender string
}

// ParseConfig 直接解析；要求Config内部字段可导出
// func ParseConfig(config Config) {
// 	fmt.Printf("name:%v,age:%v,gender:%v", config.Name, config.Age, config.Gender)
// }

// 定义接口类型；需要提供参数解析函数
type Opt interface {
	apply(*Config)
}

type option struct{ fn func(*Config) }

// 实现接口的方法
func (opt option) apply(cfg *Config) { opt.fn(cfg) }

// ParseConfigWithOption 使用函数解析；
func ParseConfigWithOption(opts ...Opt) {
	cfg := Config{}
	for _, opt := range opts {
		opt.apply(&cfg)
	}
	fmt.Printf("name:%v,age:%v,gender:%v", cfg.name, cfg.age, cfg.gender)
}

// WithName 填充姓名;名字填充可以放到此时，或者放到apply内部，实际一样
func WithName(name string) option {
	return option{func(cfg *Config) {
		cfg.name = name
	},
	}
}

// WithAge 填充年龄
func WithAge(age int) option {
	return option{func(cfg *Config) {
		cfg.age = age
	},
	}
}

func main() {
	opts := []Opt{
		WithName("ZhangSan"),
		WithAge(18),
	}
	ParseConfigWithOption(opts...)
}
