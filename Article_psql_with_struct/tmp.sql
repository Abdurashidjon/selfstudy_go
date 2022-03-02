CREATE TABLE testing (
    "id" serial PRIMARY KEY,
    "firstname" VARCHAR (100) NOT NULL,
    "lastname" VARCHAR (100) NOT NULL,
    "description" text,
    "created_at" TIMESTAMP DEFAULT(NOW())
);

INSERT INTO testing(firstname,lastname,description) VALUES('Abdurashid','Gaybullaev','Talaba') RETURNING id;