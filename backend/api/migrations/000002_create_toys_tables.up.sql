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
    age_range JSONB NOT NULL, -- {"min": 0, "max": 3}
    condition toy_condition NOT NULL,
    category toy_category NOT NULL,
    status toy_status DEFAULT 'active',
    is_deleted TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT valid_age_range CHECK (
        age_range ? 'min' AND 
        age_range ? 'max' AND 
        (age_range->>'min')::int >= 0 AND 
        (age_range->>'max')::int >= (age_range->>'min')::int
    )
);

CREATE TABLE toy_photos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    toy_id UUID NOT NULL REFERENCES toys(id),
    url VARCHAR(500) NOT NULL,
    cloudinary_id VARCHAR(255) NOT NULL,
    is_main BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT unique_main_photo UNIQUE (toy_id, is_main) 
        WHERE is_main = true
);

CREATE INDEX idx_toys_user ON toys(user_id);
CREATE INDEX idx_toys_category ON toys(category) WHERE status = 'active';
CREATE INDEX idx_toys_status ON toys(status) WHERE is_deleted IS NULL;
CREATE INDEX idx_toy_photos ON toy_photos(toy_id);