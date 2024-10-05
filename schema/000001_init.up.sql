CREATE TABLE songs
(
    "id"           SERIAL        PRIMARY KEY,
    "group_name"   varchar(255)  NOT NULL,
    "song"         varchar(255)  NOT NULL,
    "release_date" TEXT,
    "text"         TEXT,
    "link"         TEXT
);