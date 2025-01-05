# URL Shortener Project

## Overview

The **URL Shortener** project involves creating a system to store long URLs and their corresponding shortened codes in a database. The project tests performance by inserting large datasets (1,000, 1 million, and 10 million rows) and measuring system response times for queries. It also explores the impact of growing data on database performance and storage requirements.

This project uses Go (Golang) to insert data into a MySQL database, perform queries, and measure performance metrics.

## Requirements

1. **Database Setup**: Set up a MySQL database to store URLs and short codes.
2. **Table Creation**: Create a table to store original URLs, their short codes, and a creation timestamp.
3. **Inserting Data**: Insert 1,000 rows initially, then scale to 1 million and 10 million rows.
4. **Performance Testing**: Measure how much time it takes to insert large amounts of data and run queries.
5. **Querying**: Measure the performance of fetching original URLs from the database using their short codes.

## Technologies Used

- **MySQL**: Relational database for storing the URLs and short codes.
- **Go (Golang)**: Programming language for writing scripts to interact with the database, handle data insertion, and performance testing.
- **SQL**: To manage and query the database.

## Steps Taken

### 1. Database Setup
- **Database Name**: `url_shortner`
- **Table Name**: `url_shortner`
- The table stores:
  - `id`: Auto-incremented primary key.
  - `original_url`: The long URL.
  - `short_code`: The unique short code.
  - `created_at`: Timestamp of when the URL was shortened.

### 2. Inserting Data
- Data is inserted into the database using a Go script. Initially, 1,000 rows were inserted, then scaled up to 1 million and 10 million rows.
- Performance was measured for both insertion time and query speed.

### 3. Query Performance
- A Go script was used to run multiple queries to fetch original URLs using their short codes.
- The system's performance was measured by running 1 million queries and logging the time it took to retrieve the data.

### 4. Performance and Metrics
- The time taken to insert 1,000, 1 million, and 10 million rows was measured and reported.
- Query performance was tested by running 1 million iterations of querying for specific short codes.

### 5. Database Size Monitoring
- The database size was tracked as the number of rows increased, allowing us to measure the impact on storage requirements as the dataset grew.

## Future Enhancements

- **Web Interface**: Implement a web-based interface for URL shortening and retrieval.
- **Security**: Enhance security by implementing token-based short codes or encryption.
- **Query Optimization**: Implement caching strategies to speed up query performance.
- **Logging & Monitoring**: Add detailed logging and monitoring to track performance and errors.

## License

This project is open-source and free to modify or extend for any purpose. Contributions are welcome!




## How to Run

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/url-shortener.git
   cd url-shortener


2. Install Go and MySQL:

Ensure that Go and MySQL are installed on your system.

Create the database:

Log into MySQL and create the database and table:

sql
Copy code
CREATE DATABASE url_shortner;
USE url_shortner;

CREATE TABLE url_shortner (
    id INT auto_increment PRIMARY KEY,
    original_url VARCHAR(255) NOT NULL,
    short_code VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp )

4. Run the Go script:

Insert data by running the Go script to populate the database.
Query performance testing can also be done using Go scripts provided in the repository.

### Key Sections in the README:
- **Overview**: High-level description of the project and its purpose.
- **Requirements**: Technologies used and their purpose.
- **Steps Taken**: Explanation of the key stages of the project (database setup, data insertion, performance testing).
- **Future Enhancements**: Possible improvements you plan to make to the system.
- **License**: Information about the open-source nature of the project.
- **How to Run**: Instructions for setting up and running the project locally.

Feel free to modify the README as needed, and it will be ready for use on your GitHub repository!