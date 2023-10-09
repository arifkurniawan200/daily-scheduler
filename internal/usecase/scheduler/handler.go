package scheduler

import "github.com/labstack/gommon/log"

func (c cronHandler) GenerateBirthdayCampaign() {
	log.Infof("start cron GenerateBirthdayCampaign")
	err := c.u.CreateCampaignForBirthdayUser()
	if err != nil {
		return
	}
	log.Infof("finish cron GenerateBirthdayCampaign")
}
