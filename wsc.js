/**
 * Thin program to send arbitrary data to a websocket.
 * Usage: pipe ws://localhost:4000 some_data
 */
const WebSocket = require('ws')
const args = process.argv.slice(2)

const ws = new WebSocket(args[0])
ws.addEventListener('open', event => {
  ws.send(args[1])
  ws.close()
})