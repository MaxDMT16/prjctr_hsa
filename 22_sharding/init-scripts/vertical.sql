CREATE TABLE books_1 (CHECK ( category_id= 1 )) INHERITS ( books );

CREATE RULE books_insert_to_category_1 AS ON INSERT TO books
WHERE ( category_id = 1 )
DO INSTEAD INSERT INTO books_1 VALUES (NEW.*);