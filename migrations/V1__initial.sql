CREATE TABLE public.book
 (
    Name CHARACTER VARYING(30)  PRIMARY KEY,
    Author CHARACTER VARYING(64) NOT NULL,
    Genre CHARACTER VARYING(32)  NOT NULL,
    Year integer  NOT NULL
)