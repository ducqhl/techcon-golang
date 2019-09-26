const url = require('url');

// prime check
const isPrime = (number) => {
    if (number < 2) {
        return false;
    }

    for (i = 2; i < number; i++) {
        if (number % i == 0) {
            return false;
        }
    }

    return true;
};

const getNthPrime = (nth) => {
    if (nth < 1) {
        return -1;
    }

    let nthPrime = 0;
    let prime = -1;

    for (let i = 2; nthPrime < nth; i++) {
        if (isPrime(i)) {
            prime = i;
            nthPrime++;
        }
    }

    return prime;
};

const nthPrimeController = (req, res) => {
    res.writeHead(200, { 'Content-Type': 'text/html' });
    let query = url.parse(req.url, true).query;
    let result = getNthPrime(parseInt(query.n));
    res.write(result.toString());
    res.end();
};

module.exports = nthPrimeController;
