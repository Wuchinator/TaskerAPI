
BEGIN;

ALTER TABLE tasks ADD COLUMN user_id INT;
ALTER TABLE tasks ADD CONSTRAINT fk_user 
FOREIGN KEY (user_id) REFERENCES users(id);
UPDATE tasks SET user_id = 1 WHERE user_id IS NULL;

COMMIT;