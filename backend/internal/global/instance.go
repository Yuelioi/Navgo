package global

import (
	"backend/internal/config"
	"backend/internal/svc"
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
)

var ConfInst config.Config      // 配置实例
var SVCInst *svc.ServiceContext // 上下实例

func init() {
	var configFile = flag.String("f", "etc/main.yaml", "the config file")
	flag.Parse()

	conf.MustLoad(*configFile, &ConfInst)

	SVCInst = svc.NewServiceContext(ConfInst)
}
