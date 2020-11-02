
function minTimeToVisitAllPoints(points) {
    let num = 0
    for (let i = 0; i < points.length - 1; i++) {
        let point = points[i]
        let point2 = points[i + 1]
        let length = Math.abs(point[0] - point2[0])
        let length2 = Math.abs(point[1] - point2[1])
        num += Math.max(length, length2)
    }

    return num
}

const num = minTimeToVisitAllPoints([
    [1, 1],
    [3, 4],
    [-1, 0],
    [10, 0]
])
console.log(num)
