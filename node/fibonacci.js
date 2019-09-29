var url = require('url');

const fibonacci = (n) => {
    let a1 = 0,
        a2 = 1,
        a = 0;

    for (let i = 0; i <= n; i++) {
        a = a1 + a2;
        a1 = a2;
        a2 = a;
    }

    return a;
};

const fibonacciController = (req, res) => {
    res.writeHead(200, { 'Content-Type': 'text/html' });
    query = url.parse(req.url, true).query;
    result = fibonacci(query.n);
    res.write(result.toString());
    res.end();
};

module.exports = fibonacciController;
