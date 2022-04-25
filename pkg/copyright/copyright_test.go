package copyright

import (
	"BloodPressure/pkg/config"
	"fmt"
	"testing"
)

func TestLoadCopyright(t *testing.T) {
	conf := config.Config{
		BasicinfoConfig: config.BasicinfoConfig{
			AppName:    "高血压测量记录小程序后端",
			Author:     "王博杰",
			AppCompany: "深圳始动科技有限公司",
			Version:    "v0.0.1",
			Copyright:  "@2022 Begining Power,Inc. All rights reserved. 深圳始动科技有限公司保留所有权利",
		},
	}
	GetInstance().LoadCopyright(conf.BasicinfoConfig)
	fmt.Println(GetInstance().GetCopyright())
}

func TestFmtPrintCopyright(t *testing.T) {
	// conf := config.Load("./config/config.ini")
	conf := config.Config{
		BasicinfoConfig: config.BasicinfoConfig{
			AppName:    "高血压测量记录小程序后端",
			Author:     "王博杰",
			AppCompany: "深圳始动科技有限公司",
			Version:    "v0.0.1",
			Copyright:  "@2022 Begining Power,Inc. All rights reserved. 深圳始动科技有限公司保留所有权利",
		},
	}
	FmtPrintCopyright(conf.BasicinfoConfig)
}

// test unsuccessful(log wrong)
func TestLogPrintCopyright(t *testing.T) {
	// conf := config.Load("./config/config.ini")
	conf := config.Config{
		BasicinfoConfig: config.BasicinfoConfig{
			AppName:    "高血压测量记录小程序后端",
			Author:     "王博杰",
			AppCompany: "深圳始动科技有限公司",
			Version:    "v0.0.1",
			Copyright:  "@2022 Begining Power,Inc. All rights reserved. 深圳始动科技有限公司保留所有权利",
		},
	}
	LogPrintCopyright(conf.BasicinfoConfig)
}
