# go_rest
A template for a REST API service core


# Purpose
The purpose of this project is to have a template for a REST API which covers most of the best practices, paths using routing, etc...  The server is optmized for Kubernetes, Docker and is always having in consideration the best performance. 

# Run
To run this project, just do - `go run .`


## Setting Up the Database

1. Install PostgreSQL:
   - On Ubuntu: `sudo apt install postgresql postgresql-contrib`
   - On macOS: `brew install postgresql`

2. Start the PostgreSQL service:
   - `sudo service postgresql start`

3. Create a new PostgreSQL user and database:
   - `sudo -u postgres createuser -P yourusername`
   - `sudo -u postgres createdb -O yourusername go_rest_db`

4. Set up environment variables:
   - Create a `.env` file in the project root:
     ```env
     DB_USER=yourusername
     DB_PASSWORD=yourpassword
     DB_NAME=go_rest_db
     DB_HOST=localhost
     DB_PORT=5432
     ```

5. Run the Go application:
   - `go run .`

