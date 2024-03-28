import psycopg2
import uuid
import datetime

# Define your connection parameters
conn_params = {
    "host": "localhost",
    "database": "fowardapi",
    "user": "revolutionize",
    "password": "docker"
}

# Connect to the PostgreSQL database
conn = psycopg2.connect(**conn_params)
cursor = conn.cursor()

# Get the current timestamp
current_timestamp = datetime.datetime.now()

# Calculate the expiration date by adding one month to the current timestamp
expiration_date = current_timestamp + datetime.timedelta(days=30)

# Insert API key data
api_key_insert_query = f"""
    INSERT INTO api_key (id, key, ip_address, num_proxies, expiration_date)
    VALUES (%s,
            'key', '192.168.65.1', 10, %s)
"""
api_key_id = uuid.uuid4()
cursor.execute(api_key_insert_query, (str(api_key_id), expiration_date,))

# Sample data to be inserted
data_to_insert = [

]

# Define the SQL query to insert proxy data
insert_query = """
    INSERT INTO proxy (id, address, username, password, scheme)
    VALUES (%s, %s, %s, %s, %s)
"""

proxies_id = []
# Iterate over the data and insert it into the table
for item in data_to_insert:
    # Generate UUID for id field
    id_value = uuid.uuid4()
    proxies_id.append(id_value)
    # Execute the insert query
    cursor.execute(insert_query, (str(id_value), *item))

for id in proxies_id:
    cursor.execute(
        "INSERT INTO api_key_proxy (api_key_id, proxy_id) VALUES (%s, %s)", (str(api_key_id), str(id),))

client_insert_query = """
    INSERT INTO client (id, name, api_key_id)
    VALUES (%s, %s, %s)
"""
cursor.execute(client_insert_query, (str(
    uuid.uuid4()), 'zoloft', str(api_key_id),))

# Commit the transaction
conn.commit()

# Close the cursor and the connection
cursor.close()
conn.close()
