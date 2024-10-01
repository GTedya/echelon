CREATE TABLE IF NOT EXISTS thumbnails
(
    video_id TEXT PRIMARY KEY,
    thumbnail BLOB,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
);