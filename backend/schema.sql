-- This schema defines the database structure for the AI Prompt Gallery application

-- Users table stores user data
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    last_login_at TIMESTAMP WITH TIME ZONE,
    profile_image VARCHAR(255),
    bio TEXT,
    is_admin BOOLEAN DEFAULT FALSE
);

-- Categories table for organizing prompts
CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    icon_url VARCHAR(255),
    color VARCHAR(7), -- Hex color code (e.g., #FF5733)
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Tags table for prompt tagging
CREATE TABLE IF NOT EXISTS tags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    color VARCHAR(7), -- Hex color code (e.g., #FF5733)
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- Prompts table stores AI prompts
CREATE TABLE IF NOT EXISTS prompts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(100) NOT NULL,
    description TEXT,
    content TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    view_count INTEGER NOT NULL DEFAULT 0,
    favorite_count INTEGER NOT NULL DEFAULT 0,
    image_url VARCHAR(255),
    
    -- Full-text search index
    CONSTRAINT prompt_title_content_check CHECK (title != '' AND content != '')
);

-- Prompt_Tags junction table for many-to-many relationship
CREATE TABLE IF NOT EXISTS prompt_tags (
    prompt_id UUID NOT NULL REFERENCES prompts(id) ON DELETE CASCADE,
    tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    PRIMARY KEY (prompt_id, tag_id)
);

-- Favorites table to track user favorites
CREATE TABLE IF NOT EXISTS favorites (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    prompt_id UUID NOT NULL REFERENCES prompts(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, prompt_id)
);

-- Views table to track prompt views
CREATE TABLE IF NOT EXISTS views (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    prompt_id UUID NOT NULL REFERENCES prompts(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    ip VARCHAR(45), -- IPv4/IPv6 address
    user_agent TEXT
);

-- Create indexes for performance
CREATE INDEX IF NOT EXISTS idx_prompts_user_id ON prompts(user_id);
CREATE INDEX IF NOT EXISTS idx_prompts_category_id ON prompts(category_id);
CREATE INDEX IF NOT EXISTS idx_views_prompt_id ON views(prompt_id);
CREATE INDEX IF NOT EXISTS idx_favorites_prompt_id ON favorites(prompt_id);
CREATE INDEX IF NOT EXISTS idx_favorites_user_id ON favorites(user_id);

-- Full-text search indexes for searching prompts
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE INDEX IF NOT EXISTS idx_prompts_title_trgm ON prompts USING gin (title gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_prompts_content_trgm ON prompts USING gin (content gin_trgm_ops);

-- Add some initial categories
INSERT INTO categories (id, name, description, color) VALUES 
(gen_random_uuid(), 'GPT-4', 'Prompts designed for GPT-4 and similar models', '#FF5733'),
(gen_random_uuid(), 'Midjourney', 'Prompts for Midjourney image generation', '#33FF57'),
(gen_random_uuid(), 'DALL-E', 'Prompts for DALL-E image generation', '#3357FF'),
(gen_random_uuid(), 'Stable Diffusion', 'Prompts for Stable Diffusion models', '#F3FF33'),
(gen_random_uuid(), 'Claude', 'Prompts designed for Anthropic''s Claude models', '#FF33E9')
ON CONFLICT (name) DO NOTHING;

-- Add some initial tags
INSERT INTO tags (id, name, description, color) VALUES 
(gen_random_uuid(), 'Productivity', 'Prompts for enhancing productivity', '#4287f5'),
(gen_random_uuid(), 'Creative Writing', 'Prompts for creative writing assistance', '#f542a7'),
(gen_random_uuid(), 'Art', 'Prompts for generating artistic content', '#42f56f'),
(gen_random_uuid(), 'Business', 'Prompts for business use cases', '#f5d442'),
(gen_random_uuid(), 'Programming', 'Prompts for coding and development', '#8d42f5'),
(gen_random_uuid(), 'Academic', 'Prompts for academic and educational purposes', '#4245f5'),
(gen_random_uuid(), 'Roleplay', 'Prompts for character roleplay scenarios', '#f54242'),
(gen_random_uuid(), 'Fantasy', 'Prompts for fantasy themed content', '#42d3f5')
ON CONFLICT (name) DO NOTHING;

-- Functions for updating timestamps
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Triggers to automatically update updated_at
CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER update_categories_updated_at
BEFORE UPDATE ON categories
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER update_tags_updated_at
BEFORE UPDATE ON tags
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER update_prompts_updated_at
BEFORE UPDATE ON prompts
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

-- Function to increment view count
CREATE OR REPLACE FUNCTION increment_view_count()
RETURNS TRIGGER AS $$
BEGIN
   UPDATE prompts
   SET view_count = view_count + 1
   WHERE id = NEW.prompt_id;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to increment view count on new view
CREATE TRIGGER increment_prompt_view_count
AFTER INSERT ON views
FOR EACH ROW
EXECUTE FUNCTION increment_view_count();

-- Function to update favorite count
CREATE OR REPLACE FUNCTION update_favorite_count()
RETURNS TRIGGER AS $$
BEGIN
   IF TG_OP = 'INSERT' THEN
      UPDATE prompts
      SET favorite_count = favorite_count + 1
      WHERE id = NEW.prompt_id;
      RETURN NEW;
   ELSIF TG_OP = 'DELETE' THEN
      UPDATE prompts
      SET favorite_count = favorite_count - 1
      WHERE id = OLD.prompt_id;
      RETURN OLD;
   END IF;
END;
$$ LANGUAGE plpgsql;

-- Triggers to update favorite count
CREATE TRIGGER increment_favorite_count
AFTER INSERT ON favorites
FOR EACH ROW
EXECUTE FUNCTION update_favorite_count();

CREATE TRIGGER decrement_favorite_count
AFTER DELETE ON favorites
FOR EACH ROW
EXECUTE FUNCTION update_favorite_count();