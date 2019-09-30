const whileTrueController = (req, res) => {
    let i = 0;
    setTimeout(() => {
        while (true) {
            console.log(i + ' ');
            i++;
        }
    }, 0);

    res.write('Opsss!!!');
    res.end();
};

module.exports = whileTrueController;
