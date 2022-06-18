CREATE TABLE STATION_MATSTER_TB (
    station_id INTEGER PRIMARY KEY,
    station_name VARCHAR(256) NOT NULL,
    line_name VARCHAR(256) NOT NULL
);

CREATE TABLE NEAREST_STATION_TB (
    bar_id INTEGER,
    station_id INTEGER,
    on_foot_time INTEGER NOT NULL,
    PRIMARY KEY(bar_id, station_id)
);

CREATE TABLE BAR_MASTER_TB (
    bar_id INTEGER PRIMARY KEY,
    bar_name VARCHAR(256) NOT NULL,
    number_of_taps INTEGER,
    address VARCHAR(256) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    official_link TEXT,
    tabelog_link TEXT,
    food VARCHAR(256),
    style VARCHAR(256),
    list_image_url VARCHAR(256)
);

CREATE TABLE BREWERY_HANDLED_TB (
    bar_id INTEGER,
    brewery_id INTEGER,
    standard BOOLEAN NOT NULL,
    PRIMARY KEY(bar_id, brewery_id)
);

CREATE TABLE BREWERY_MASTER_TB (
    brewery_id INTEGER PRIMARY KEY,
    brewery_name VARCHAR(256) NOT NULL,
    country VARCHAR(128) NOT NULL,
    area VARCHAR(128)
);