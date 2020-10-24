ALTER TABLE records DROP COLUMN clear_time;
ALTER TABLE records ADD COLUMN min_clear_time TIME COMMENT '最短クリア時間';
