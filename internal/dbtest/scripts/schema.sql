PRAGMA case_sensitive_like=on;
PRAGMA foreign_keys=on;

DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS queues;
DROP TABLE IF EXISTS user_queues;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  uuid BLOB NOT NULL,
  email TEXT NOT NULL,
  "password" TEXT,
  
  CONSTRAINT pk_users PRIMARY KEY (uuid)
);

CREATE TABLE user_queues (
  user_uuid BLOB NOT NULL,
  queue_id INTEGER NOT NULL
);

CREATE TABLE queues (
  id INTEGER NOT NULL,
  name TEXT,
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp,
  consuming BOOLEAN
);

CREATE TABLE messages (
  id INTEGER NOT NULL,
  queue_id INTEGER,
  message TEXT,
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);

