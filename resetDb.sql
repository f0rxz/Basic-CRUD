CREATE TABLE Card_Quest (
  id SERIAL PRIMARY KEY,
  Title VARCHAR(255),
  Description TEXT,
  Rating INTEGER,
  Duration VARCHAR(255),
  Difficulty INTEGER,
  Type VARCHAR(255),
  Image BYTEA,
  Genres VARCHAR(255)
);
CREATE TABLE Card_Quest_Full (
  id SERIAL PRIMARY KEY,
  Title VARCHAR(255),
  Description TEXT,
  Images BYTEA,
  Genres VARCHAR(255),
  Rating INTEGER,
  Difficulty INTEGER,
  Duration VARCHAR(255),
  Type VARCHAR(255),
  Card_Quest_id INTEGER REFERENCES Card_Quest(id)
);
CREATE TABLE Comments (
  id SERIAL PRIMARY KEY,
  Author VARCHAR(255),
  Rating INTEGER,
  Text TEXT,
  Quest_ID INTEGER REFERENCES Card_Quest(id)
);