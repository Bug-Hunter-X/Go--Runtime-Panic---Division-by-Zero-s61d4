# Go: Runtime Panic - Division by Zero

This repository demonstrates a common runtime error in Go: division by zero. The `bug.go` file contains a function that doesn't handle potential division by zero, resulting in a panic when executed.  The `bugSolution.go` file provides a corrected version that gracefully handles this error.

## Bug Description

The `calculate` function in `bug.go` attempts to divide `a` by `b` without checking if `b` is zero.  If `b` is zero, this results in a `runtime panic: integer divide by zero` error.

## Solution

The `bugSolution.go` file demonstrates how to avoid this error by adding a check to see if the divisor (`b`) is zero before performing the division. If `b` is zero, it returns an appropriate error or value, preventing the panic.