package models

import "github.com/gobuffalo/nulls"

type Record struct {
    ID              int             `db:"id"`
    PlayerID        int             `db:"player_id"`
    StageId         int             `db:"stage_id"`
    IsClear         bool            `db:"is_clear"`
    PlayTimes       int             `db:"play_times"`
    FirstClearTimes nulls.Int       `db:"first_clear_times"`
    MinClearTime    nulls.String    `db:"min_clear_time"`
}