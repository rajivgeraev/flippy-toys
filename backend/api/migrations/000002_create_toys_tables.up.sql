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

CREATE TABLE toys (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    age_range JSONB,
    condition toy_condition,
    category toy_category,
    status toy_status DEFAULT 'active',
    is_deleted TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    -- Изменяем проверку age_range, учитывая что поле может быть NULL
    CONSTRAINT valid_age_range CHECK (
        age_range IS NULL OR (
            age_range ? 'min' AND 
            age_range ? 'max' AND 
            (age_range->>'min')::int >= 0 AND 
            (age_range->>'max')::int >= (age_range->>'min')::int
        )
    )
);