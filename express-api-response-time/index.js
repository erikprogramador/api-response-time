const express = require('express')
const app = express()

const knex = require('knex')({
  client: 'mysql2',
  connection: {
    host: '127.0.0.1',
    user: 'root',
    password: 'secret',
    database: 'api_response_time',
  },
})

app.get('/v1/clients', async function (request, response) {
  const result = await knex.select('*').from('clients')

  response.status(200).json(result)
})

app.listen(8000)
