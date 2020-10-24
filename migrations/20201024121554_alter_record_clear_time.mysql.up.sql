ALTER TABLE records DROP COLUMN min_clear_time;
ALTER TABLE records ADD COLUMN clear_time float COMMENT '最短クリア時間';
