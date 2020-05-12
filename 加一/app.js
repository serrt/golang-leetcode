
/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function(digits) {
    for (let i = digits.length - 1; i >= 0; i--) {
        var subNum = digits[i]
        subNum++
        if (subNum >= 10) {
            digits[i] = 0
            if (i === 0) {
                digits.unshift(1)
            }
        } else {
            digits[i] = subNum
            break;
        }
    }
    return digits
}

console.log(plusOne([1,2,3]))
