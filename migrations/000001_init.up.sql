CREATE TYPE analytics_status AS ENUM ('pending', 'in_progress', 'completed', 'failed');

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) UNIQUE NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,
                       email VARCHAR(100),
                       created_at TIMESTAMPTZ DEFAULT NOW(),
                       updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE websites (
                          id SERIAL PRIMARY KEY,
                          url TEXT NOT NULL,
                          created_at TIMESTAMPTZ DEFAULT NOW(),
                          updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE user_websites (
                               user_id INTEGER NOT NULL REFERENCES users(id),
                               website_id INTEGER NOT NULL REFERENCES websites(id),
                               PRIMARY KEY (user_id, website_id)
);

CREATE TABLE analytics (
                           id SERIAL PRIMARY KEY,
                           website_id INTEGER NOT NULL REFERENCES websites(id),
                           data JSONB NOT NULL,
                           status analytics_status NOT NULL DEFAULT 'pending',
                           created_at TIMESTAMPTZ DEFAULT NOW(),
                           updated_at TIMESTAMPTZ DEFAULT NOW()
);
