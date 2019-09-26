const http = require('http');
const fibonanceController = require('./fibonance')
const readFileController = require('./read-file')

const handler = (req, res) => {
    res.writeHead(200, { 'Content-Type': 'text/html' });
    if(req.url.startsWith("/fibonance")){
        return fibonanceController(req, res);
    }

    if(req.url.startsWith("/readfile")){
        return readFileController(req, res);
    }

    res.write('hit success');
    res.end();
};

http.createServer(handler).listen(8080, () => console.log('8080 ready'));
