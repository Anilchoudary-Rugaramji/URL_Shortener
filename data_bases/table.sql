CREATE database url_shortner;

USE url_shortner;

CREATE TABLE url_shortner(
	id INT auto_increment primary Key,
    original_url varchar(255) not null,
    short_code varchar(50) unique not null,
    created_at timestamp default current_timestamp
);

INSERT INTO url_shortner (original_url, short_code)
values ('https://example.com', 'example1'),
('https://openai.com', 'openai2'),
('https://google.com', 'google3');


SELECT * FROM  url_shortner;