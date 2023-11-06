function add(a, b) {
    return a + b;
}
function isEven(a) {
    if (a % 2 == 0) {
        return true;
    }
    else {
        return false;
    }
}
var isOdd = function (a) {
    if (a % 2 == 0) {
        return false;
    }
    else {
        return true;
    }
};
console.log(add(10, 5));
console.log(isEven(13));
console.log(isEven(4));
console.log(isOdd(1));
