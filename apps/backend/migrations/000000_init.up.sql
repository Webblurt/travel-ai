-- Able extension for uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS itineraries (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    city TEXT,
    cost TEXT,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    created_by TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by TEXT,
    deleted_at TIMESTAMP,
    deleted_by TEXT
);

CREATE TABLE IF NOT EXISTS itinerary_days (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    itinerary_id UUID REFERENCES itineraries(id) ON DELETE SET NULL,
    day_number SMALLINT CHECK (day_number > 0),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    created_by TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by TEXT,
    deleted_at TIMESTAMP,
    deleted_by TEXT
);

CREATE TABLE IF NOT EXISTS day_activities (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    day_id UUID REFERENCES itinerary_days(id) ON DELETE SET NULL,
    time_of_day TEXT NOT NULL,
    place TEXT,
    activity_description TEXT,
    activity_type TEXT,
    cost TEXT,
    duration TEXT,
    activity_address TEXT,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    created_by TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by TEXT,
    deleted_at TIMESTAMP,
    deleted_by TEXT
);

CREATE TABLE IF NOT EXISTS activity_details (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    activity_id UUID REFERENCES day_activities(id) ON DELETE SET NULL,
    short_description TEXT,
    history_of_background TEXT,
    opening_hours TEXT,
    best_time_to_visit TEXT,
    average_visit_duration TEXT,
    entry_fee TEXT,
    website TEXT,
    contact_phone TEXT,
    neighborhood TEXT,
    nearest_transport TEXT,
    distance_from_city_center TEXT,
    map_link TEXT,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    created_by TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by TEXT,
    deleted_at TIMESTAMP,
    deleted_by TEXT
);

CREATE TABLE IF NOT EXISTS detail_images (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    detail_id UUID REFERENCES activity_details(id) ON DELETE SET NULL,
    image_url TEXT NOT NULL,
    image_description TEXT,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    created_by TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by TEXT,
    deleted_at TIMESTAMP,
    deleted_by TEXT
);

CREATE TABLE IF NOT EXISTS detail_facts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    detail_id UUID REFERENCES activity_details(id) ON DELETE SET NULL,
    fact TEXT,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    created_by TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by TEXT,
    deleted_at TIMESTAMP,
    deleted_by TEXT
);

CREATE TABLE IF NOT EXISTS detail_to_bring (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    detail_id UUID REFERENCES activity_details(id) ON DELETE SET NULL,
    what_to_bring TEXT,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    created_by TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by TEXT,
    deleted_at TIMESTAMP,
    deleted_by TEXT
);

CREATE TABLE IF NOT EXISTS detail_local_tips (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    detail_id UUID REFERENCES activity_details(id) ON DELETE SET NULL,
    tip TEXT,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    created_by TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by TEXT,
    deleted_at TIMESTAMP,
    deleted_by TEXT
);

CREATE TABLE IF NOT EXISTS detail_safety_notes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    detail_id UUID REFERENCES activity_details(id) ON DELETE SET NULL,
    note TEXT,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    created_by TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW(),
    updated_by TEXT,
    deleted_at TIMESTAMP,
    deleted_by TEXT
);