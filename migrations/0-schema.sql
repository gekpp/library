CREATE TABLE IF NOT EXISTS book (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    annotation TEXT NOT NULL,
    author TEXT NOT NULL,
    images JSON NOT NULL
);
