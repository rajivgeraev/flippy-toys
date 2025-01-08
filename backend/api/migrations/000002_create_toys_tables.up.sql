CREATE TYPE toy_condition AS ENUM ('new', 'like_new', 'good', 'acceptable');
CREATE TYPE toy_status AS ENUM ('active', 'reserved', 'exchanged', 'deleted');
CREATE TYPE toy_category AS ENUM (
    'construction_toys',
    'dolls',
    'vehicles',
    'educational',
    'outdoor',
    'board_games',
    'electronic',
    'stuffed_animals',
    'action_figures',
    'arts_crafts',
    'musical',
    'other'
);

-- Таблица игрушек
CREATE TABLE toys (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    age_range JSONB,
    condition toy_condition NOT NULL,
    category toy_category NOT NULL,
    status toy_status NOT NULL DEFAULT 'active',
    is_deleted TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    -- Ограничения
    CONSTRAINT valid_age_range CHECK (
        age_range IS NULL OR (
            age_range ? 'min' AND 
            age_range ? 'max' AND 
            (age_range->>'min')::int >= 0 AND 
            (age_range->>'max')::int >= (age_range->>'min')::int
        )
    ),
    CONSTRAINT title_length CHECK (char_length(title) >= 3),
    CONSTRAINT description_length CHECK (char_length(description) >= 10)
);

-- Индекс для ускорения поиска по user_id
CREATE INDEX idx_toys_user_id ON toys(user_id);

-- Таблица фотографий игрушек
CREATE TABLE toy_photos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    toy_id UUID NOT NULL REFERENCES toys(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    cloudinary_id VARCHAR(255) NOT NULL,
    asset_id VARCHAR(255),
    is_main BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Индекс для ускорения поиска фотографий по toy_id
CREATE INDEX idx_toy_photos_toy_id ON toy_photos(toy_id);
