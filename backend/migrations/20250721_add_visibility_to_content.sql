-- Add visibility column to content tables
ALTER TABLE posts
    ADD COLUMN IF NOT EXISTS visibility VARCHAR(20) DEFAULT 'public' CHECK (visibility IN ('public', 'private', 'followers_only'));

ALTER TABLE reels
    ADD COLUMN IF NOT EXISTS visibility VARCHAR(20) DEFAULT 'public' CHECK (visibility IN ('public', 'private', 'followers_only'));

ALTER TABLE videos
    ADD COLUMN IF NOT EXISTS visibility VARCHAR(20) DEFAULT 'public' CHECK (visibility IN ('public', 'private', 'followers_only'));

ALTER TABLE streams
    ADD COLUMN IF NOT EXISTS visibility VARCHAR(20) DEFAULT 'public' CHECK (visibility IN ('public', 'private', 'followers_only'));
