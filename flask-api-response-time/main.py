from flask import Flask
from flask_mysqldb import MySQL
app = Flask(__name__)

app.config['MYSQL_HOST'] = 'localhost'
app.config['MYSQL_USER'] = 'root'
app.config['MYSQL_PASSWORD'] = 'secret'
app.config['MYSQL_DB'] = 'api_response_time'

mysql = MySQL(app)

@app.route('/v1/clients', methods=['GET'])
def hello_world():
    cur = mysql.connection.cursor()
    cur.execute('SELECT * FROM clients')
    result = cur.fetchall()
    cur.close()
    return str(result)

if __name__ == '__main__':
    app.run()
