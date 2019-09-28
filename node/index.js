const http = require('http');
const hitMeController = require('./hitme')
const fibonanceController = require('./fibonance')
const readFileController = require('./read-file')
const nthPrimeController = require('./nthprime')

const handler = (req, res) => {
    console.log(`${new Date()} - ${req.method} - ${req.url}`)

    res.writeHead(200, { 'Content-Type': 'text/html' });

    if(req.url.startsWith("/hitme")){
        return hitMeController(req, res);
    }

    if(req.url.startsWith("/fibonance")){
        return fibonanceController(req, res);
    }

    if(req.url.startsWith("/readfile")){
        return readFileController(req, res);
    }

    if(req.url.startsWith("/nthprime")){
        return nthPrimeController(req, res);
    }

    res.write('hit success');
    res.end();
};

http.createServer(handler).listen(8080, () => console.log('8080 ready'));
