CREATE TABLE todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    status TEXT NOT NULL,
    priority TEXT NOT NULL,
    category TEXT,
    task TEXT NOT NULL,
    lastmod TIMESTAMP
);

