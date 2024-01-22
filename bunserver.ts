import Bun from 'bun'

Bun.serve({
	fetch(request) {
		console.log(request)
		return new Response("Bun!")
	},
	port: 3000
})
