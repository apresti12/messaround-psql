DROP TABLE IF EXISTS secretSanta;
CREATE TABLE secretSanta (
   id           SERIAL PRIMARY KEY,
   name         VARCHAR(128) NOT NULL,
   wishList     text[] NOT NULL,
   rejectPerson VARCHAR(128) NOT NULL
);

INSERT INTO secretSanta
    (name, wishList, rejectPerson)
VALUES
    ('Mario',
     '{"xbox", "ps5"}',
     'Antonio'),
    ('Dominic',
     '{"pc", "games"}',
     'Mario'),
    ('Antonio',
     '{"Snowboard", "Truck"}',
     'Sophie'),
    ('Sophie',
     '{"clothes", "makeup"}',
     'Dominic')