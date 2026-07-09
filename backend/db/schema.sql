-- Afya Plus Database Schema

-- Users: both CHWs and Supervisors
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    pin_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('chw', 'supervisor')),
    zone VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Households: the families being tracked
CREATE TABLE households (
    id SERIAL PRIMARY KEY,
    head_name VARCHAR(100) NOT NULL,
    location VARCHAR(200),
    member_count INT DEFAULT 1,
    created_by INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Visits: each time a CHW checks in on a household
CREATE TABLE visits (
    id SERIAL PRIMARY KEY,
    household_id INT REFERENCES households(id),
    chw_id INT REFERENCES users(id),
    client_visit_id VARCHAR(64) UNIQUE NOT NULL,
    visit_type VARCHAR(30),
    child_weight_kg DECIMAL(4,1),
    child_age_months INT,
    vaccination_due_date DATE,
    vaccination_done BOOLEAN DEFAULT FALSE,
    notes TEXT,
    visited_at TIMESTAMP NOT NULL,
    synced_at TIMESTAMP DEFAULT NOW()
);

-- Flags: auto-generated when a visit trips a risk rule
CREATE TABLE flags (
    id SERIAL PRIMARY KEY,
    visit_id INT REFERENCES visits(id),
    household_id INT REFERENCES households(id),
    reason VARCHAR(100) NOT NULL,
    status VARCHAR(20) DEFAULT 'flagged' CHECK (status IN ('flagged', 'under_review', 'action_taken', 'resolved')),
    assigned_to INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Helpful indexes for common lookups
CREATE INDEX idx_visits_household ON visits(household_id);
CREATE INDEX idx_visits_chw ON visits(chw_id);
CREATE INDEX idx_flags_status ON flags(status);
CREATE INDEX idx_households_created_by ON households(created_by);
