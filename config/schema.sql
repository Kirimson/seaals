CREATE TABLE seals (
  id INTEGER PRIMARY KEY,
  path text NOT NULL,
  mime_type text NOT NULL,
  created_at datetime NOT NULL DEFAULT current_timestamp
);

CREATE TABLE tags (
  id INTEGER PRIMARY KEY,
  name text NOT NULL UNIQUE
);

CREATE TABLE seal_tags (
  id INTEGER PRIMARY KEY,
  seal_id INTEGER NOT NULL,
  tag_id INTEGER NOT NULL,
  FOREIGN KEY (seal_id) REFERENCES seals (id),
  FOREIGN KEY (tag_id) REFERENCES tags (id)
);
