CREATE TABLE books_cat_1_5 (CHECK ( category_id >= 1 OR category_id <= 5 )) INHERITS ( books );

CREATE RULE books_insert_to_category_1_5 AS ON INSERT TO books
WHERE ( category_id >= 1 AND category_id <= 5 )
DO INSTEAD INSERT INTO books_cat_1_5 VALUES (NEW.*);

CREATE TABLE books_cat_6_10 (CHECK ( category_id >= 6 OR category_id <= 10 )) INHERITS ( books );

CREATE RULE books_insert_to_category_6_10 AS ON INSERT TO books
WHERE ( category_id >= 6 AND category_id <= 10 )
DO INSTEAD INSERT INTO books_cat_6_10 VALUES (NEW.*);