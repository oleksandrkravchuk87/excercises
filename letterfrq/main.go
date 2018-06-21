package main

func Frequency(text string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range []rune(text) {
		m[v]++
	}
	return m
}

func ConcurrentFrequency(text []string) map[rune]int {
	ch := make(chan map[rune]int, len(text))

	for _, v := range text {

		go func(v string) {

			m := make(map[rune]int)
			for _, v := range []rune(v) {
				m[v]++
			}

			ch <- m

			return
		}(v)
	}

	resultMap := make(map[rune]int)
	for i := 0; i < len(text); i++ {
		m := <-ch
		for k, v := range m {
			resultMap[k] = resultMap[k] + v

		}
	}

	return resultMap
}

func main() {

}
