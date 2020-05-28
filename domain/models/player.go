package models

import (
    "time"

    "github.com/gobuffalo/nulls"
)

type Player struct {
    ID              int         `db:"id"`
    Name            string      `db:"name"`
    PlayingStageID  nulls.Int   `db:"playing_stage_id"`
    CreateTime      time.Time   `db:"create_time"`
    LastLoginTime   time.Time   `db:"last_login_time"`
}
