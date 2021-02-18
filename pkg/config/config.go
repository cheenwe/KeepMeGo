package config

import (
	"KeepMeGo/pkg/util"
	"bytes"
	"log"
	"strings"

	"github.com/Unknwon/goconfig"
)

func RequestUrl()(url string )  {

	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}

	env_value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "env")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "request_url", err)
	}
	log.Printf("======= 当前程序运行环境: %s =======", env_value)

	value, err := cfg.GetValue(env_value, "request_url")

	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "request_url", err)
	}
	log.Printf("======= 请求接口: %s =======", value)
	url  = value
	return 
}


func Conf(name string)(url string )  {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}

	env_value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "env")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "request_url", err)
	}
	log.Printf("======= 当前程序运行环境: %s =======", env_value)

	value, err := cfg.GetValue(env_value, name)

	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "request_url", err)
	}
	log.Printf("======= 请求接口: %s =======", value)
	url  = value
	return 
}

func GetIntIni(name string)(url int )  {

	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}

	env_value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "env")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "request_url", err)
	}
	log.Printf("======= 当前程序运行环境: %s =======", env_value)

	value, err := cfg.Int(env_value, name)

	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "request_url", err)
	}
	log.Printf("======= 请求接口: %s =======", value)
	url  = value
	return 
}



func SetIni(abbr string, value string )(url string )  {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}
	env_value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "env")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "env", err)
	}
	log.Printf("======= 当前程序运行环境: %s =======", env_value)

	log.Printf("abbr  %v",abbr)
	log.Printf("value  %v",value)

	cfg.SetValue(env_value, abbr, value)
	
	// 保存 ConfigFile 对象到文件系统，保存后的键顺序与读取时的一样
	err = goconfig.SaveConfigFile(cfg, "config.ini")
	if err != nil {
		log.Fatalf("无法保存配置文件：%s", err)
	}
	url = ""
	return 
}


func RequestInterval()(num int64 )  {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}
	
	env_value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "env")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "env", err)
	}
	log.Printf("======= 当前程序运行环境: %s =======", env_value)

	vInt, err := cfg.Int64(env_value, "interval")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "int", err)
	}
	log.Printf("%s > %s: %v", "must", "int", vInt)
	num = vInt
	return 
}

func WatchAllPid()(num int )  {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}
	
	env_value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "env")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "env", err)
	}
	log.Printf("======= 当前程序运行环境: %s =======", env_value)
	vInt,_ := cfg.Int(env_value, "watch_all")
	num = vInt
	return 
}


func Pids()(num string )  {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}
	
	env_value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "env")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "env", err)
	}
	log.Printf("======= 当前程序运行环境: %s =======", env_value)

	vInt, err := cfg.GetValue(env_value, "watch_abbr")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "watch_abbr", err)
	}
	log.Printf("%s > %s: %v", "watch_abbr", "is", vInt)

	num = vInt
	return 
}


func PidAdd(abbr string)(url string )  {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}
	env_value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "env")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "env", err)
	}
	log.Printf("======= 当前程序运行环境: %s =======", env_value)


	config_abbrs, err := cfg.GetValue(env_value, "watch_abbr")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "watch_abbr", err)
	}
	log.Printf("%s > %s: %v", "watch_abbr", "is", config_abbrs)

    //查找 包含
    var isCon bool = strings.Contains(config_abbrs, abbr)
	// fmt.Println(isCon) //true
	if isCon {
		url = "该值已存在，请勿重复添加！"
	} else {
				
		//需要先导入bytes包
		s1 := ";"
		//定义Buffer类型
		var bt bytes.Buffer
		// 向bt中写入字符串 
		bt.WriteString(config_abbrs)
		bt.WriteString(abbr)
		bt.WriteString(s1)
		//获得拼接后的字符串
		s2 := bt.String()
		
		// 设置某个键的注释，返回值为 true 时表示注释被插入或删除（空字符串），false 表示注释被重写
		v := cfg.SetValue(env_value, "watch_abbr", s2)
		log.Printf("键 %s 的注释被插入或删除：%v", "New Value", v)

		value, err := cfg.GetValue(env_value, "watch_abbr")

		if err != nil {
			log.Fatalf("无法获取键值（%s）：%s", "watch_abbr", err)
		}
		log.Printf("======= watch_abbr: %s =======", value)
		url  = value

		// 保存 ConfigFile 对象到文件系统，保存后的键顺序与读取时的一样
		err = goconfig.SaveConfigFile(cfg, "config.ini")
		if err != nil {
			log.Fatalf("无法保存配置文件：%s", err)
		}
	}
	return 
}
func PidDel(abbr string)(url string )  {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}
	env_value, err := cfg.GetValue(goconfig.DEFAULT_SECTION, "env")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "env", err)
	}
	log.Printf("======= 当前程序运行环境: %s =======", env_value)

	config_abbrs, err := cfg.GetValue(env_value, "watch_abbr")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "watch_abbr", err)
	}
	log.Printf("%s > %s: %v", "watch_abbr", "is", config_abbrs)

    //查找 包含
    var isCon bool = strings.Contains(config_abbrs, abbr)
	// fmt.Println(isCon) //true
	if isCon {

		//需要先导入bytes包
		s1 := ";"
		//定义Buffer类型
		var bt bytes.Buffer
		// 向bt中写入字符串 
		bt.WriteString(abbr)
		bt.WriteString(s1)
		//获得拼接后的字符串
		s2 := bt.String()
		
		new_str := strings.Replace(config_abbrs, s2, "", -1)
		// 设置某个键的注释，返回值为 true 时表示注释被插入或删除（空字符串），false 表示注释被重写
		v := cfg.SetValue(env_value, "watch_abbr", new_str)
		log.Printf("键 %s 的注释被插入或删除：%v", "New Value", v)

		value, err := cfg.GetValue(env_value, "watch_abbr")

		if err != nil {
			log.Fatalf("无法获取键值（%s）：%s", "watch_abbr", err)
		}
		log.Printf("======= watch_abbr: %s =======", value)
		url  = value

		// 保存 ConfigFile 对象到文件系统，保存后的键顺序与读取时的一样
		err = goconfig.SaveConfigFile(cfg, "config.ini")
		if err != nil {
			log.Fatalf("无法保存配置文件：%s", err)
		}
	} else {
				
		url = "该值不已存在，请勿重复删除！"
	}
	return 
}


func InitConfigFile()(url string )  {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}
	env_value := "auth"
	// 设置默认用户名密码。
	v := cfg.SetValue(env_value, "username", "admin")
	log.Println("set username %v", v)

	v = cfg.SetValue(env_value, "password", "admin")
	log.Println("set password %v", v)


	token := util.RandomString(24)
	v = cfg.SetValue(env_value, "token", token)

	v = cfg.SetValue("config", "port", "5554")
	log.Println("set port %v", v)
	
	v = cfg.SetValue("config", "db", "data.db")
	log.Println("set port %v", v)
	
	// 保存 ConfigFile 对象到文件系统，保存后的键顺序与读取时的一样
	err = goconfig.SaveConfigFile(cfg, "config.ini")
	if err != nil {
		log.Fatalf("无法保存配置文件：%s", err)
	}
	return 
}


func GetAuthIni(name string)(url string )  {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}

	env_value := "auth"

	value, err := cfg.GetValue(env_value, name)
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", name, err)
	}
	url  = value
	return 
}

func GetConfigIni(env string, name string)(url string )  {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件：%s", err)
	}

	value, err := cfg.GetValue(env, name)
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", name, err)
	}
	url  = value
	return 
}