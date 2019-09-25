const fibonance = (n) => {
    if (n < 2) return n;

    return fibonance(n - 1) + fibonance(n - 2);
};

module.exports = fibonance;
