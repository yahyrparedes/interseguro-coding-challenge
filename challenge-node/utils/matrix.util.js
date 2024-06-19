// validate is matrix diagonal
export function isDiagonal (matrix)   {
    if (!isMatrix(matrix)) {
        return false;
    }

    for (let i = 0; i < matrix.length; i++) {
        if (matrix[i][i] === 0) {
            return false;
        }
    }

    return true;
}

// get max value to matrix bi-dimensional
export function getMax  (matrix)  {
    return Math.max(...matrix.flat());
}

// get min value to matrix bi-dimensional
export function getMin   (matrix)  {
    return Math.min(...matrix.flat());
}

/// get sum all values to matrix bi-dimensional
export function sum  (matrix)   {
    return matrix.flat().reduce((a, b) => a + b);
}

// validate is matrix bi-dimensional
export function isMatrix   (matrix)  {
    if (!Array.isArray(matrix)) {
        return false;
    }

    for (let i = 0; i < matrix.length; i++) {
        if (!Array.isArray(matrix[i])) {
            return false;
        }
    }

    return true;
}

// calculate average to matrix bi-dimensional
export function calculateAverage  (matrix)  {
    let flattened = matrix.flat();
    let sum = flattened.reduce((a, b) => a + b, 0);
    return sum / flattened.length;
}
