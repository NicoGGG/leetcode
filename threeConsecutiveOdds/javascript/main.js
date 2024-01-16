/**
 * @param {number[]} arr
 * @return {boolean}
 */
function threeConsecutiveOdds(arr) {
    let n = 0
    for (let i = 0; i < arr.length; i++) {
        if (arr[i] % 2 != 0) {
            n++;
        } else {
            n = 0;
        }
        if (n == 3) {
            return true;
        }
    }
    return false;
};

console.log(threeConsecutiveOdds([2,6,4,1]))
console.log(threeConsecutiveOdds([1,2,34,3,4,5,7,23,12]))
