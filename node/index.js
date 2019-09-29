const http = require('http');
const hitMeController = require('./hitme')
const fibonacciController = require('./fibonacci')
const readFileController = require('./read-file')
const sortController = require('./sort')

const handler = (req, res) => {
    console.log(`Node: ${new Date().toLocaleTimeString()} - ${req.method} - ${req.url}`)

    res.writeHead(200, { 'Content-Type': 'text/html' });

    if(req.url.startsWith("/hitme")){
        return hitMeController(req, res);
    }

    if(req.url.startsWith("/fibonacci")){
        return fibonacciController(req, res);
    }

    if(req.url.startsWith("/readfile")){
        return readFileController(req, res);
    }

    if(req.url.startsWith("/sort")){
        return sortController(req, res);
    }
    
    res.write('hit success');
    res.end();
};

http.createServer(handler).listen(8080, () => console.log('8080 ready'));
