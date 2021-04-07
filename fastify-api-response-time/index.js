// Require the framework and instantiate it
const fastify = require('fastify')({ logger: true })

const knex = require('knex')({
  client: 'mysql2',
  connection: {
    host: '127.0.0.1',
    user: 'root',
    password: 'secret',
    database: 'api_response_time',
  },
})

// Declare a route
fastify.get('/v1/clients/', async (request, reply) => {
  const result = await knex.select('*').from('clients')

  return reply
    .code(200)
    .header('Content-Type', 'application/json; charset=utf-8')
    .send(result)
})

// Run the server!
const start = async () => {
  try {
    await fastify.listen(8000)
  } catch (err) {
    fastify.log.error(err)
    process.exit(1)
  }
}
start()
