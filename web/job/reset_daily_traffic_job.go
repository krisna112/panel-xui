package job

import (
	"time"

	"github.com/mhsanaei/3x-ui/v2/database"
	"github.com/mhsanaei/3x-ui/v2/logger"
	"github.com/mhsanaei/3x-ui/v2/web/service"
	"github.com/mhsanaei/3x-ui/v2/xray"
)

type ResetDailyTrafficJob struct {
	xrayService service.XrayService
}

func NewResetDailyTrafficJob() *ResetDailyTrafficJob {
	return &ResetDailyTrafficJob{}
}

func (j *ResetDailyTrafficJob) Run() {
	logger.Info("Run ResetDailyTrafficJob")
	db := database.GetDB()

	// Reset daily usage
	err := db.Model(&xray.ClientTraffic{}).Updates(map[string]any{
		"up_daily":   0,
		"down_daily": 0,
	}).Error
	if err != nil {
		logger.Error("ResetDailyTrafficJob reset usage failed:", err)
		return
	}

	// Re-enable clients that were disabled due to daily limit
	// Logic: Enable=false AND DepletedAt>0 AND TotalDaily>0
	// AND (Total=0 OR Up+Down < Total) -> Not total limit depleted
	// AND (ExpiryTime=0 OR ExpiryTime > Now) -> Not expired

	now := time.Now().Unix() * 1000
	result := db.Model(&xray.ClientTraffic{}).
		Where("enable = ? AND depleted_at > 0 AND total_daily > 0", false).
		Where("total = 0 OR up + down < total").
		Where("expiry_time = 0 OR expiry_time > ?", now).
		Updates(map[string]any{
			"enable":      true,
			"depleted_at": 0,
		})

	if result.Error != nil {
		logger.Error("ResetDailyTrafficJob re-enable clients failed:", result.Error)
	} else if result.RowsAffected > 0 {
		logger.Infof("ResetDailyTrafficJob re-enabled %d clients", result.RowsAffected)
		j.xrayService.SetToNeedRestart()
	}
}
