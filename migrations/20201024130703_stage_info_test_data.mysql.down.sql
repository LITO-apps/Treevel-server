DELETE FROM players WHERE id IN (3, 4, 5, 6);
DELETE FROM records WHERE id IN (9000, 9001, 9002, 9003, 9004, 9005, 9006, 9007);

UPDATE records SET stage_id = '1' WHERE stage_id = 'Spring-1-1';
UPDATE records SET stage_id = '2' WHERE stage_id = 'Spring-1-2';

ALTER TABLE records MODIFY stage_id INT NOT NULL;
