var url = require('url');

const fibonance = (n) => {
    a1 = 1;
    a2 = 1;

    if (n == 1 || n == 2) return 1;

    a = 0;

    for (let i = 3; i <= n; ++i) {
        a = a1 + a2;
        a1 = a2;
        a2 = a;
    }

    return a;
};

const fibonanceController = (req, res) => {
    res.writeHead(200, { 'Content-Type': 'text/html' });
    query = url.parse(req.url, true).query;
    result = fibonance(query.n);
    res.write(result.toString());
    res.end();
};

module.exports = fibonanceController;