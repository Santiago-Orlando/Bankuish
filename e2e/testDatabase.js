const pg = require("pg")


const { Pool } = pg

const credentials = {
  host: "localhost",
  database: "test",
  port: 3002,
  user: "postgres",
  password: "your_password"
}

const pool = new Pool(credentials)

module.exports = pool