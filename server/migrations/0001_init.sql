-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  steam_id TEXT UNIQUE NOT NULL,
  username TEXT NOT NULL,
  avatar_url TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS maps (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS rounds (
  id BIGSERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  panorama_url TEXT NOT NULL,
  correct_latitude DOUBLE PRECISION NOT NULL,
  correct_longitude DOUBLE PRECISION NOT NULL,
  correct_year INT NOT NULL,
  min_year INT NOT NULL,
  max_year INT NOT NULL,
  location_name TEXT NOT NULL,
  description TEXT,
  difficulty TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS rooms (id BIGSERIAL PRIMARY KEY, code TEXT UNIQUE NOT NULL, host_user_id BIGINT, created_at TIMESTAMPTZ DEFAULT now());
CREATE TABLE IF NOT EXISTS matches (id BIGSERIAL PRIMARY KEY, room_id BIGINT, status TEXT NOT NULL, created_at TIMESTAMPTZ DEFAULT now());
CREATE TABLE IF NOT EXISTS match_rounds (id BIGSERIAL PRIMARY KEY, match_id BIGINT, round_id BIGINT, round_index INT NOT NULL);
CREATE TABLE IF NOT EXISTS player_guesses (id BIGSERIAL PRIMARY KEY, match_round_id BIGINT, user_id BIGINT, guessed_latitude DOUBLE PRECISION, guessed_longitude DOUBLE PRECISION, guessed_year INT, submitted_at TIMESTAMPTZ DEFAULT now());
CREATE TABLE IF NOT EXISTS player_scores (id BIGSERIAL PRIMARY KEY, match_round_id BIGINT, user_id BIGINT, location_score INT NOT NULL, time_score INT NOT NULL, total_score INT NOT NULL);
CREATE TABLE IF NOT EXISTS achievements (id BIGSERIAL PRIMARY KEY, user_id BIGINT, key TEXT NOT NULL, unlocked_at TIMESTAMPTZ DEFAULT now());
-- +goose Down
DROP TABLE IF EXISTS achievements, player_scores, player_guesses, match_rounds, matches, rooms, rounds, maps, users;
