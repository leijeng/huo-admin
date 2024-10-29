package start

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/leijeng/huo-admin/common/config"
	"github.com/leijeng/huo-admin/common/middleware"
	coreCfg "github.com/leijeng/huo-core/config"
	"github.com/leijeng/huo-core/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"path"
	"time"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:     "start",
		Short:   "Get Application config info",
		Example: "huo start -c resources/config.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "resources/config.yml", "Start server with provided configuration file")
}

func run() {
	if configYml == "" {
		panic("找不到配置文件")
	}

	v := viper.New()
	v.SetConfigFile(configYml)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file: %v \n", err))
	}

	var cfg coreCfg.AppCfg

	if err = v.Unmarshal(&cfg); err != nil {
		fmt.Println(err)
	}

	if cfg.Server.RemoteEnable {

		configpath := cfg.Remote.Path
		configpath = path.Join(cfg.Server.Mode, cfg.Server.Name, configpath)

		rviper := viper.New()
		if cfg.Remote.SecretKeyring == "" {
			err = rviper.AddRemoteProvider(cfg.Remote.Provider, cfg.Remote.Endpoint, configpath)
		} else {
			err = rviper.AddSecureRemoteProvider(cfg.Remote.Provider, cfg.Remote.Endpoint, configpath, cfg.Remote.SecretKeyring)
		}
		if err != nil {
			panic(fmt.Sprintf("Fatal error remote config : %v \n", err))
		}
		rviper.SetConfigType(cfg.Remote.GetConfigType())
		err = rviper.ReadRemoteConfig()
		if err != nil {
			panic(fmt.Sprintf("Fatal error remote config : %v \n", err))
		}
		var remoteCfg coreCfg.AppCfg
		rviper.Unmarshal(&remoteCfg)

		log.Printf("------------------remoteCfg------------%#v-----------------------\n", remoteCfg)
		log.Printf("------------------AllSettings------------%#v-----------------------", rviper.AllSettings())
		mergeCfg(&cfg, &remoteCfg)

		extend := rviper.Sub("extend")
		if extend != nil {
			extend.Unmarshal(&config.Ext)
		}
		go func() {
			for {
				time.Sleep(time.Second * 5) // delay after each request
				err := rviper.WatchRemoteConfig()
				if err != nil {
					fmt.Println(err)
					continue
				}
				rviper.Unmarshal(&remoteCfg)

				mergeCfg(&cfg, &remoteCfg)

				extend := rviper.Sub("extend")
				if extend != nil {
					extend.Unmarshal(&config.Ext)
				}
			}
		}()
	} else {
		mergeCfg(&cfg, nil)
		v.Sub("extend").Unmarshal(&config.Ext)
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config file changed:", e.String())
			if err = v.Unmarshal(cfg); err != nil {
				fmt.Println(err)
			}
			mergeCfg(&cfg, nil)
			extend := v.Sub("extend")
			if extend != nil {
				extend.Unmarshal(&config.Ext)
			}
		})
	}

	core.Init()

	r := core.GetGinEngine()
	middleware.InitMiddleware(r, &core.Cfg)

	//初始化路由
	for _, f := range AppRouters {
		f()
	}

	core.Run()
}

// 服务启动后要初始化的资源
func startedInit() {
	core.Log.Debug("服务启动，初始化执行完成")
}

// 服务关闭要释放的资源
func toClose() {
	core.Log.Debug("服务关闭需要释放的资源")
}

func mergeCfg(local, remote *coreCfg.AppCfg) {
	if remote != nil {
		core.Cfg = *local
		core.Cfg = *remote
		core.Cfg.Server.Mode = local.Server.Mode
		core.Cfg.Server.RemoteEnable = local.Server.RemoteEnable
		core.Cfg.Remote = local.Remote
		core.Cfg.Server.Name = local.Server.Name
		core.Cfg.Server.Port = local.Server.Port
		core.Cfg.Server.Host = local.Server.Host
	} else {
		core.Cfg = *local
	}
}
