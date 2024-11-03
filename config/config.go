package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var ConfigInstance Config

var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0"

type Config struct {
	PushPlusToken  string `yaml:"pushplus_token"`
	RefreshToken   string `yaml:"refresh_token"`
	BilibiliCookie string `yaml:"bilibili_cookie"`
	KKCookie       string `yaml:"_UP_A4A_11_=wb9651e7d5a84e1ea33e5df412bd1143; b-user-id=24a9f570-cd2f-d5e8-247b-ea15d1e6ff28; _UP_30C_6A_=st96a6201b7fmr8qaf8u1vp63v33340j; _UP_TS_=sg1337bd897950b7c2ebcc134d96a8a503a; _UP_E37_B7_=sg1337bd897950b7c2ebcc134d96a8a503a; _UP_TG_=st96a6201b7fmr8qaf8u1vp63v33340j; _UP_335_2B_=1; __pus=1bea49f61d449a62778ccbc3a18f1ecfAAS+Qg5HrfEFZV24cSAwaT4bN1MCX8YRUs4hcKaBgVQak8JWLjC+18YciKSPEaSrCPZ7TEI3XaXMfdfUy75ePGlr; __kp=d18ba6d0-93ec-11ef-920b-db563eb3e006; __kps=AARZpDhF7KxbHoLDFI2RRmTY; __ktd=2NAOyuassUPU7thaE2nozA==; __uid=AARZpDhF7KxbHoLDFI2RRmTY; isg=BK2ta9aYtigWp1MpC40xYKzLvEknCuHcpldD6e-yLcSRZswYt1tGrvp1UDqAZvmU; __puus=c9f2f3dd0e5dd27cba74fbd1b38179e8AAQ9EJvClq4zXUhjjEGGAW43xzW5mYiUX7LKhYsZVzbgdIHx+f9QolBwGFGCFo5ZmlzfczFzogy7wkZ1hv3dPr3VNTFs8+jgiwx7ym7VG0vnxmv/nSFeRRXH5XUBsWpnFKag3YyrMz92wEMRwStU6wUqQWKQ6OY9FqctFvpmzgp0Um/H1G0R1PNGLBEsSdy1qdEvxkL+czaG+9Bd/mSayvI/"`
	JdCookie       string `yaml:"jd_cookie"`
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	confFIle, err := os.ReadFile("./config.yaml")
	if err != nil {
		panic(err.Error())
	}
	config := Config{}
	yaml.Unmarshal(confFIle, &config)
	ConfigInstance = config
}
