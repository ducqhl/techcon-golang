var fs = require('fs'),
    path = require('path'),
    filePath = path.join(__dirname, '/data/100MB.txt');

const readFileController = (req, res) => {
    res.writeHead(200, { 'Content-Type': 'text/html' });

    fs.readFile(filePath, 'utf-8', (data, err) => {
        if (err) {
            console.log(err);
            res.writeHead(500, { 'Content-Type': 'text/html' });
            res.write('Read Failed');
            res.end();
        } else {
            res.writeHead(200, { 'Content-Type': 'text/html' });
            res.write('Read Success');
            res.end();
        }
    });
};

module.exports = readFileController;
