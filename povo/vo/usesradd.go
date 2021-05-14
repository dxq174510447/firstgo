package vo

import "time"

type UsersAdd struct {
	Name       string     `column:"name" columnDefinition:""`
	Password   string     `column:"password" columnDefinition:""`
	Status     int        `column:"status" columnDefinition:""`
	Fee        float64    `column:"fee" columnDefinition:""`
	FeeStatus  int        `column:"fee_status" columnDefinition:""`
	FeeTotal   int64      `column:"fee_total" columnDefinition:""`
	CreateDate *time.Time `column:"create_date" columnDefinition:""`
	CreateTime *time.Time `column:"create_time" updatable:"false" columnDefinition:""`
}
