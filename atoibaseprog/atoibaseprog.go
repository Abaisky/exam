package exam

func AtoiBase(s string, base string) int {
    res := 0
    n := len(base)
    if len(base) < 2 {
        return res
    }
    for _, i := range base {
        if i == '-' || i == '+' {
            return res
        }
    }
    for i := 0; i < n-1; i++ {
        for j := i + 1; j < n; j++ {
            if base[i] == base[j] {
                return res
            }
        }
    }
    count := len(s) - 1
    for _, i := range s {
        for index, j := range base {
            if i == j {
                res = index*RecursivePower(n, count) + res
                count--
            }
        }
    }
    return res
}

func RecursivePower(nb int, power int) int {
    if power < 0 {
        return 0
    } else {
        if power == 0 {
            return 1
        } else {
            return nb * RecursivePower(nb, power-1)
        }
    }
}
