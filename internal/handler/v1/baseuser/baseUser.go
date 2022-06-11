package baseuser

import (
	"BloodPressure/internal/service"
	"BloodPressure/pkg/constant"
	"BloodPressure/pkg/errors"
	"BloodPressure/pkg/errors/code"
	"BloodPressure/pkg/response"
	"context"

	"github.com/gin-gonic/gin"
)

// BaseUserHandler 用户业务handler
type BaseUserHandler struct {
	userSrv      service.BaseUserService        // 用户服务
	bprSrv       service.PatientBpRecordService // 血压相关记录服务
	trplanSrc    service.TreatmentPlanService   // 治疗方案服务
	pinfoService service.PatientInfoService     // 患者信息记录
	tonSrv       service.TongueDetailService    // 舌苔脉象记录
	// 我的检验指标
	m24bprSrv      service.MedicalReport24HoursbprService
	m24ecgSrv      service.MedicalReport24hoursecgService
	mecgSrv        service.MedicalReportEcgService
	mcreatinineSrv service.MedicalReportEchocardiographyService
	// 我的检查报告
	tbnpSrv        service.TestIndicatorBnpService
	tcreatinineSrv service.TestIndicatorCreatinineService
}

// 新建一个handler
func NewBaseUserHandler(
	_userSrv service.BaseUserService,
	_bprSrv service.PatientBpRecordService,
	_trplanSrc service.TreatmentPlanService,
	_pinfoService service.PatientInfoService,
	_tonSrv service.TongueDetailService,

	_m24bprSrv service.MedicalReport24HoursbprService,
	_m24ecgSrv service.MedicalReport24hoursecgService,
	_mecgSrv service.MedicalReportEcgService,
	_mcreatinineSrv service.MedicalReportEchocardiographyService,

	_tbnpSrv service.TestIndicatorBnpService,
	_tcreatinineSrv service.TestIndicatorCreatinineService,
) *BaseUserHandler {
	// 新建Handler
	return &BaseUserHandler{
		userSrv:      _userSrv,
		bprSrv:       _bprSrv,
		trplanSrc:    _trplanSrc,
		pinfoService: _pinfoService,
		tonSrv:       _tonSrv,
		// 我的检验指标
		m24bprSrv:      _m24bprSrv,
		m24ecgSrv:      _m24ecgSrv,
		mecgSrv:        _mecgSrv,
		mcreatinineSrv: _mcreatinineSrv,
		// 我的检查报告
		tbnpSrv:        _tbnpSrv,
		tcreatinineSrv: _tcreatinineSrv,
	}
}

// 获取用户信息
func (uh *BaseUserHandler) GetBaseUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetUint(constant.UserID)
		// 通过id找到基本用户
		baseUser, err := uh.userSrv.GetById(context.TODO(), uid)

		if err != nil {
			response.JSON(c, errors.Wrap(err, code.NotFoundErr, "用户信息为空"), nil)
		} else {
			response.JSON(c, nil, baseUser)
		}
	}
}
