package scheduler

import (
	"github.com/robfig/cron/v3"
	"template/internal/usecase"
	"time"
)

type cronHandler struct {
	t usecase.TransactionUcase
	u usecase.UserUcase
}

func RunCron(u usecase.UserUcase, t usecase.TransactionUcase) {
	c := cronHandler{
		t: t,
		u: u,
	}
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))
	{
		scheduler.AddFunc("0 9 * * *", c.GenerateBirthdayCampaign)
	}
	go scheduler.Start()
}
