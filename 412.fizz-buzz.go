/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 */
func fizzBuzz(n int) []string {
	rnt := make([]string, 0, n)
	for i:=1; i < n+1; i++{
		if i%3 ==0 && i%5 != 0{
			rnt = append(rnt, "Fizz")
		}else if i%5 == 0 && i%3 != 0{
			rnt = append(rnt, "Buzz")
		}else if i%5 == 0 && i%3 == 0{
			rnt = append(rnt, "FizzBuzz")
		}else{
			rnt = append(rnt, fmt.Sprint(i)) 
		}
	}
	return rnt
}

