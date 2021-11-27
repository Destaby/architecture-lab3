-- CREATE USER user1 WITH ENCRYPTED PASSWORD 'pass1';
-- CREATE DATABASE system;
-- GRANT ALL PRIVILEGES ON DATABASE system TO user1;
-- psql -U user1 -d system -a -f db/schema.sql

DROP TABLE IF EXISTS machines;
DROP TABLE IF EXISTS balancers;
-- Create tables.
CREATE TABLE balancers
(
    id          SERIAL PRIMARY KEY,
    job         VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE machines
(
    id          SERIAL PRIMARY KEY,
    working     BOOLEAN NOT NULL DEFAULT FALSE,
    balancer_id int NOT NULL,
    FOREIGN KEY (balancer_id) REFERENCES balancers(id) ON DELETE CASCADE
);

-- Insert demo data.
INSERT INTO balancers (job) VALUES ('writing');
INSERT INTO balancers (job) VALUES ('reading');
INSERT INTO balancers (job) VALUES ('taking');

INSERT INTO machines (working, balancer_id) VALUES (TRUE, 1);
INSERT INTO machines (working, balancer_id) VALUES (FALSE, 1);
INSERT INTO machines (working, balancer_id) VALUES (TRUE, 1);
INSERT INTO machines (working, balancer_id) VALUES (TRUE, 1);
INSERT INTO machines (working, balancer_id) VALUES (FALSE, 1);
INSERT INTO machines (working, balancer_id) VALUES (TRUE, 1);
INSERT INTO machines (working, balancer_id) VALUES (TRUE, 1);
INSERT INTO machines (working, balancer_id) VALUES (FALSE, 1);
INSERT INTO machines (working, balancer_id) VALUES (TRUE, 2);
INSERT INTO machines (working, balancer_id) VALUES (FALSE, 2);
INSERT INTO machines (working, balancer_id) VALUES (TRUE, 2);
INSERT INTO machines (working, balancer_id) VALUES (TRUE, 2);
INSERT INTO machines (working, balancer_id) VALUES (FALSE, 2);
INSERT INTO machines (working, balancer_id) VALUES (TRUE, 2);
INSERT INTO machines (working, balancer_id) VALUES (TRUE, 3);
INSERT INTO machines (working, balancer_id) VALUES (TRUE, 3);
INSERT INTO machines (working, balancer_id) VALUES (FALSE, 3);
INSERT INTO machines (working, balancer_id) VALUES (TRUE, 3);
