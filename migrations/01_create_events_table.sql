-- +goose Up
CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);

--- INSERT INTO events (id, title, date, duration, user_id, notify_before, notify_at) VALUES 
---    ('fcfa069c-28a9-48d6-b48d-befc8133f2b4', 'Event 1', '2022-01-06 11:00:00', 900, 'U1', 0,   '2022-01-06 11:00:00'),
---    ('87a876ce-4488-45a9-bd36-ca72368a7185', 'Event 2', '2022-01-06 12:00:00', 900, 'U2', 900, '2022-01-06 11:45:00'),
---    ('5753e882-91e0-4e1a-a827-eef8d8271e50', 'Event 3', '2022-01-06 11:45:00', 900, 'U3', 0,   '2022-01-06 11:45:00')
---;

-- +goose Down
DROP TABLE files;