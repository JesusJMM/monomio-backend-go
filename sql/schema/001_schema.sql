CREATE TABLE users (
  id    BIGSERIAL NOT NULL,
  name  VARCHAR(255) NOT NULL,
  password text NOT NULL,
  img_url text
);

CREATE TABLE bios (
  id    BIGSERIAL NOT NULL,
  user_id    BIGSERIAL NOT NULL,
  bio   text NOT NULL
);

CREATE TABLE posts (
  id    BIGSERIAL NOT NULL,
  user_id    BIGSERIAL NOT NULL,
  create_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  title VARCHAR(255) NOT NULL,
  description text,
  content text
);
