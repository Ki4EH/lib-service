import psycopg2
from psycopg2 import Error
from random import randint

def connect_db(HOST):
    connection = psycopg2.connect(user=HOST["user"],
                                password=HOST["password"],
                                host=HOST["host"],
                                port=HOST["port"], 
                                dbname=HOST["dbname"])
    cursor = connection.cursor()
    print("Информация о сервере PostgreSQL")
    print(connection.get_dsn_parameters(), "\n")
    cursor.execute("SELECT version();")
    record = cursor.fetchone()
    print("Вы подключены к - ", record, "\n")
    cursor.close()
    return connection


def fetch_data(cursor, connection=None):
    cursor.execute("SELECT * from tracing")
    return cursor.fetchall(), cursor.close()


def creator(connection):
    last = list(range(30))
    for i in range(10):
        last = list(range(30))
        for j in range(10):
            cursor = connection.cursor()
            cursor.execute(f"""INSERT INTO tracing (id, user_id, book_id) VALUES ({i * 10 + j}, {i + 1}, {randomiser(last)})""")
            connection.commit()
            cursor.close()
    print("OK")


def randomiser(last):
    op = randint(0, len(last) - 1)
    zen = last[op]
    del last[op]
    return zen