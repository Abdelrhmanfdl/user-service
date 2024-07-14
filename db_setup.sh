#!/bin/bash



# Define ScyllaDB connection parameters
SCYLLA_HOST="127.0.0.1"
SCYLLA_PORT="9042"

# Define the keyspace and table creation CQL commands
CREATE_KEYSPACE_CQL="CREATE KEYSPACE IF NOT EXISTS chatchatgo WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'};"

CREATE_TABLE_CQL="CREATE TABLE IF NOT EXISTS chatchatgo.users (    id UUID,    username TEXT,    email TEXT,    password TEXT, PRIMARY KEY(id, email));"

# Function to execute CQL commands
execute_cql() {
    local cql_command=$1
    echo "exec " $cql_command 
    echo $cql_command | cqlsh $SCYLLA_HOST $SCYLLA_PORT
}

# Create keyspace
echo "Creating keyspace..."
execute_cql "$CREATE_KEYSPACE_CQL"

# Create table
echo "Creating table..."
execute_cql "$CREATE_TABLE_CQL"

echo "Keyspace and table created successfully."