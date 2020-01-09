const express = require('express')
const app = express()
const port = 9002

function hello(req, res) {
    console.log(`server2 node hit`)
    res.send('Hello from server2 node!')
}

app.get('/', hello)
app.get('/server2', hello)

app.listen(port, () => console.log(`Example app listening on port ${port}!`))