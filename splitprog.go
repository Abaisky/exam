func Split(str, charset string) []string {
    if charset == "" {
        ar := make([]string, len(str))
        i := 0
        for _, v := range str {
            ar[i] = string(v)
            i++
        }
        return ar
    }

    nw := 0
    a := 0
    for i := 0; i <= len(str)-len(charset); i++ {
        if str[i:i+len(charset)] == charset {
            a++
            i = i + len(charset)
        }
    }
    ar := make([]string, a+1)
    j := 0
    for i := 0; i <= len(str)-len(charset); i++ {
        if str[i:i+len(charset)] == charset {
            ar[j] = str[nw:i]
            j++
            nw = i + len(charset)
            i = i + len(charset)
        }
    }
    ar[j] = str[nw:]
    return ar
}
