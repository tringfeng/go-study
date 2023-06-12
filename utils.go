package main

func main() {
    
}

// 转换小写字母
func toLowerCase(str string) string {
    rune_arr := []rune(str)
    for i, _ := range rune_arr {
        if rune_arr[i] >= 65 && rune_arr[i] <= 90 {
            rune_arr[i] += 32
        }
    }
    return string(rune_arr)
}
