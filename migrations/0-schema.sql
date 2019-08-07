CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE book (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    annotation TEXT NOT NULL,
    author TEXT NOT NULL,
    images JSONB NOT NULL
);

CREATE TABLE book_status_change_event (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    book_id BIGINT NULL,
    subscriber_id UUID NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
    created_by UUID NOT NULL,
    last_modified_at TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);
CREATE INDEX idx_book_status_change_event_book_id_subscriber_id ON book_status_change_event (book_id, subscriber_id, created_at DESC);