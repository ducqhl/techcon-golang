var fs = require('fs'),
    path = require('path'),
    url = require('url');

const readFileController = (req, res) => {
    res.writeHead(200, { 'Content-Type': 'text/html' });
    let query = url.parse(req.url, true).query;
    let quota = query.n || '500KB';

    filePath = path.join(__dirname, `../data/${quota}.txt`);

    fs.readFile(filePath, (err, data) => {
        if (err) {
            console.log(err);
            res.writeHead(500, { 'Content-Type': 'text/html' });
            res.write('Read Failed');
            res.end();
        } else {
            res.writeHead(200, { 'Content-Type': 'text/html' });
            res.write('Read Success: ' + data.length);
            res.end();
        }
    });
};

module.exports = readFileController;
