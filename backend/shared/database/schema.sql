-- schema.sql
-- PostgreSQL schema for Superapp backend

-- Drop schema to reset database (removes all tables and data)
DROP SCHEMA IF EXISTS public CASCADE;
CREATE SCHEMA public;

-- Enable CITEXT extension for case-insensitive text
CREATE EXTENSION IF NOT EXISTS citext;

-- Users table (profiles)
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    privacy VARCHAR(10) DEFAULT 'public' CHECK (privacy IN ('public', 'private')),
    verified BOOLEAN DEFAULT FALSE,
    follower_count INTEGER DEFAULT 0,
    following_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Channels table
CREATE TABLE IF NOT EXISTS channels (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    verified BOOLEAN DEFAULT FALSE,
    subscriber_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Videos table (long-form, channel-only)
CREATE TABLE IF NOT EXISTS videos (
    id SERIAL PRIMARY KEY,
    channel_id INTEGER REFERENCES channels(id) ON DELETE CASCADE NOT NULL,
    url VARCHAR(255) NOT NULL,
    title VARCHAR(100),
    caption TEXT,
    duration INTEGER,
    status VARCHAR(10) DEFAULT 'active' CHECK (status IN ('active', 'archived', 'deleted')),
    view_count INTEGER DEFAULT 0,
    like_count INTEGER DEFAULT 0,
    share_count INTEGER DEFAULT 0,
    location VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Posts table (images/reels, up to 10 items)
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    channel_id INTEGER REFERENCES channels(id) ON DELETE CASCADE,
    media JSONB NOT NULL, -- Array of {type: 'image'/'reel', url: string}
    linked_video_id INTEGER REFERENCES videos(id) ON DELETE SET NULL,
    caption TEXT,
    status VARCHAR(10) DEFAULT 'active' CHECK (status IN ('active', 'archived', 'deleted')),
    view_count INTEGER DEFAULT 0,
    like_count INTEGER DEFAULT 0,
    share_count INTEGER DEFAULT 0,
    location VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT check_media_count CHECK (jsonb_array_length(media) <= 10)
);

-- Index for posts
CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts(user_id);
CREATE INDEX IF NOT EXISTS idx_posts_channel_id ON posts(channel_id);

-- Reels table (single short video)
CREATE TABLE IF NOT EXISTS reels (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    channel_id INTEGER REFERENCES channels(id) ON DELETE CASCADE,
    url VARCHAR(255) NOT NULL,
    linked_video_id INTEGER REFERENCES videos(id) ON DELETE SET NULL,
    caption TEXT,
    duration INTEGER,
    status VARCHAR(10) DEFAULT 'active' CHECK (status IN ('active', 'archived', 'deleted')),
    view_count INTEGER DEFAULT 0,
    like_count INTEGER DEFAULT 0,
    share_count INTEGER DEFAULT 0,
    location VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index for reels
CREATE INDEX IF NOT EXISTS idx_reels_user_id ON reels(user_id);
CREATE INDEX IF NOT EXISTS idx_reels_channel_id ON reels(channel_id);

-- Streams table (placeholder for live streaming)
CREATE TABLE IF NOT EXISTS streams (
    id SERIAL PRIMARY KEY,
    channel_id INTEGER REFERENCES channels(id) ON DELETE CASCADE NOT NULL,
    url VARCHAR(255),
    status VARCHAR(10) DEFAULT 'inactive' CHECK (status IN ('active', 'inactive')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index for streams
CREATE INDEX IF NOT EXISTS idx_streams_channel_id ON streams(channel_id);

-- Tags table (using CITEXT for case-insensitive name)
CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY,
    name CITEXT NOT NULL UNIQUE
);

-- Pre-populate tags with 10 categories
INSERT INTO tags (name) VALUES
                            ('tech'),
                            ('lifestyle'),
                            ('food'),
                            ('travel'),
                            ('fashion'),
                            ('music'),
                            ('sports'),
                            ('art'),
                            ('gaming'),
                            ('education')
    ON CONFLICT (name) DO NOTHING;

-- Content-Tags junction table
CREATE TABLE IF NOT EXISTS content_tags (
    type VARCHAR(10) NOT NULL CHECK (type IN ('post', 'reel', 'video')),
    content_id INTEGER NOT NULL,
    tag_id INTEGER REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (type, content_id, tag_id)
    -- Note: Validate content_id references (posts.id, reels.id, videos.id) in application logic
);

-- Index for content_tags
CREATE INDEX IF NOT EXISTS idx_content_tags_tag_id ON content_tags(tag_id);

-- Comments table (nested)
CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY,
    type VARCHAR(10) NOT NULL CHECK (type IN ('post', 'reel', 'video')),
    content_id INTEGER NOT NULL,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    parent_id INTEGER REFERENCES comments(id) ON DELETE CASCADE,
    like_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    -- Note: Validate content_id references (posts.id, reels.id, videos.id) in application logic
);

-- Index for comments
CREATE INDEX IF NOT EXISTS idx_comments_content_id ON comments(content_id);

-- Likes table
CREATE TABLE IF NOT EXISTS likes (
    id SERIAL PRIMARY KEY,
    type VARCHAR(10) NOT NULL CHECK (type IN ('post', 'reel', 'video', 'comment')),
    content_id INTEGER NOT NULL,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (type, content_id, user_id)
    -- Note: Validate content_id references (posts.id, reels.id, videos.id, comments.id) in application logic
);

-- Index for likes
CREATE INDEX IF NOT EXISTS idx_likes_content_id ON likes(content_id);

-- Follows table (profile follows and follow requests)
CREATE TABLE IF NOT EXISTS follows (
    follower_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    followed_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    status VARCHAR(10) DEFAULT 'accepted' CHECK (status IN ('pending', 'accepted')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (follower_id, followed_id)
);

-- Index for follows
CREATE INDEX IF NOT EXISTS idx_follows_followed_id ON follows(followed_id);

-- Subscriptions table (channel subscriptions)
CREATE TABLE IF NOT EXISTS subscriptions (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    channel_id INTEGER REFERENCES channels(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, channel_id)
);

-- Index for subscriptions
CREATE INDEX IF NOT EXISTS idx_subscriptions_channel_id ON subscriptions(channel_id);

-- Notifications table
CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(20) NOT NULL CHECK (type IN ('follow_request', 'like', 'comment')),
    data JSONB NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index for notifications
CREATE INDEX IF NOT EXISTS idx_notifications_user_id ON notifications(user_id);

-- Chats table (placeholder)
CREATE TABLE IF NOT EXISTS chats (
    id SERIAL PRIMARY KEY,
    type VARCHAR(10) DEFAULT 'one_on_one' CHECK (type IN ('one_on_one', 'group')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Chat Participants table (placeholder)
CREATE TABLE IF NOT EXISTS chat_participants (
    chat_id INTEGER REFERENCES chats(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (chat_id, user_id)
);

-- Channel Communities table (placeholder)
CREATE TABLE IF NOT EXISTS communities (
    id SERIAL PRIMARY KEY,
    channel_id INTEGER REFERENCES channels(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);