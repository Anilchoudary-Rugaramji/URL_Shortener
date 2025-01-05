-- Create a database named 'url_shortner' to store URLs and their short codes
CREATE DATABASE url_shortner;

-- Switch to the newly created database
USE url_shortner;

-- Create a table called 'url_shortner' with the following fields:
-- 'id' as a primary key with auto increment
-- 'original_url' to store the original URL
-- 'short_code' to store the unique shortened code
-- 'created_at' to automatically capture the time when the URL was added
CREATE TABLE url_shortner (
    id INT auto_increment primary Key,        -- Auto-incremented ID for each URL entry
    original_url VARCHAR(255) NOT NULL,       -- Original URL field, cannot be NULL
    short_code VARCHAR(50) UNIQUE NOT NULL,  -- Unique short code for the URL, cannot be NULL
    created_at TIMESTAMP DEFAULT current_timestamp -- Timestamp for when the entry is created
);

-- Insert sample data into the table
-- The following values represent original URLs with their respective short codes
INSERT INTO url_shortner (original_url, short_code)
VALUES 
    ('https://example.com', 'example1'),   -- Shortened version of 'https://example.com'
    ('https://openai.com', 'openai2'),      -- Shortened version of 'https://openai.com'
    ('https://google.com', 'google3');      -- Shortened version of 'https://google.com'

-- Select all records from the 'url_shortner' table to verify the inserted data
SELECT * FROM url_shortner;

-- Count the total number of records in the 'url_shortner' table
SELECT COUNT(*) FROM url_shortner;

-- Select the first 10 records from the 'url_shortner' table (useful for paging or previewing data)
SELECT * FROM url_shortner LIMIT 10;

-- Get the size of the 'url_shortner' table in kilobytes (KB)
-- This query calculates the size of the table by adding up the data and index lengths
SELECT 
    table_name AS "Table", 
    ROUND((data_length + index_length) / 1024, 2) AS "Size (KB)"  -- Size in KB (rounded to 2 decimal places)
FROM 
    information_schema.TABLES 
WHERE 
    table_schema = 'url_shortner'  -- Specify the database schema
    AND table_name = 'url_shortner';  -- Specify the table name

-- Show the status of the 'url_shortner' table, including info like the number of rows, data size, etc.
SHOW TABLE STATUS WHERE Name = 'url_shortner';

-- Remove all records from the 'url_shortner' table but keep the structure intact
-- This is useful when you want to clear the data but keep the table for future use
TRUNCATE TABLE url_shortner;

-- Select the 'short_code' of the first 5 records (this returns an empty result after the truncate)
SELECT short_code FROM url_shortner LIMIT 5;

-- Enable profiling to measure performance of SQL queries
SET profiling = 1;

-- Query to retrieve the original URLs corresponding to certain short codes
SELECT original_url
FROM url_shortner
WHERE short_code IN ('0000ShIg', '0004ChRr', '0005w7q9');  -- Search for multiple short codes

-- Show the profile of the query execution (after profiling is enabled)
-- This will give detailed information about the query's performance such as time spent, rows examined, etc.
SHOW PROFILE FOR QUERY 1;
