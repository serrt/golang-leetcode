
/**
 * @param {number[][]} grid
 * @param {number} k
 * @return {number[][]}
 */
var shiftGrid = function(grid, k) {
    var arr = [];
    var len = grid[0].length;
    for (let i = 0; i < grid.length; i++) {
        arr = arr.concat(...grid[i])
    }
    // var start = arr.slice(0, k)
    if (k > arr.length) {
        k = k % arr.length
    }
    var end = arr.slice(0 - k)
    arr.splice(0 - k, k)
    arr.unshift(...end)
    var result = []
    for (let i = 0; i < grid.length; i++) {
        result.push(arr.slice(i * len, (i + 1) * len))
    }
    return result;
}

const grid = [[1],[2],[3],[4],[7],[6],[5]]

console.log(shiftGrid(grid, 23))
