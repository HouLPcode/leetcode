Backtracking 回溯法，常用来求所有的解，排列，组合，n皇后，数独等

template：

``` golang
func backtrack(candidate) {
    if find_solution(candidate) {
        output(candidate)
        return
    }

    // iterate all possible candidates.
     for next_candidate := range list_of_candidates {
         if is_valid(next_candidate) == false {
             continue
         }

        // try this partial candidate solution
        place(next_candidate)
        // given the candidate, explore further.
        backtrack(next_candidate)
        // backtrack
        remove(next_candidate)
     }
}
```

根据模板，对应的排列代码

// 排列


// 组合


// N皇后


// 数独



