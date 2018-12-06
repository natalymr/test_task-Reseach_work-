package main

import (
    "fmt"
    "math"
)

type Pair struct {
    first int
    second int
}

func main() {

    var T, N int
    
    // read from console max time T and the number of tests N
    fmt.Scanf("%d %d", &T, &N)
  
    // create slices for storing the values of test duration and probability of their fall
    t_arr := make([]int, 0, N)
    p_arr := make ([]float64, 0, N)
    // create a slices for storing more accurate probability values that will be needed in the solution
    tmp := make([]float64, 0, N)
    
    var i, j, t int
    var tmp_i int
    var tmp_f float64
    
    // fill from console t_arr, p_arr
    for i = 0; i < N; i++ {
        fmt.Scanf("%d %g", &tmp_i, &tmp_f)
        t_arr = append(t_arr, tmp_i)
        p_arr = append(p_arr, tmp_f)
    }
    
    // create a slice in which we store the solutions of subtasks
    mtx := make([]float64, 0, (T + 1) * (N + 1))
  
    // create a slice in which we will store pairs of values,
    // where we came from and where we are, 
    // this array will be needed to restore the solution
    path := make([]Pair, 0, (T + 1) * (N + 1))
    
    // create a slice in which we will store a set of tests that be contained in an answer
    set_of_tests := make([]int, 0, N)
    
    // fill the tmp array
    for i = 0; i < N; i++ {
    
        tmp = append(tmp, math.Log(1 - p_arr[i]))
    }

    // fill mtx and path of default values 
    for i = 0; i < T + 1; i++ {
        for j = 0; j < N + 1; j++ {

            mtx = append(mtx, 2)
            path = append(path, Pair{-1, -1})
        }
    } 
    
    // the first value of main mtx
    mtx[0] = 0

    for j = 1; j < N + 1; j++ {
        for t = 0; t < T + 1; t++ {
            
            // that if needed for filling path and further recovering an answer
            if mtx[t * (N + 1) + j] > mtx[t * (N + 1) + (j - 1)] {
                path[t * (N + 1) + j] = Pair{t, j - 1}
            }
            
            // in this dynamics there are three transition
      
            // the first transition: a case where we don't take a test
            mtx[t * (N + 1) + j] = math.Min(mtx[t * (N + 1) + j], mtx[t * (N + 1) + (j - 1)])
            
            // the second transition: a case where an answer for the previous time is better than for the current one
            if t != 0 {
                if mtx[t * (N + 1) + j] > mtx[(t - 1) * (N + 1) + j] {
                    path[t * (N + 1) + j] = Pair{t - 1, j}
                }
                
                mtx[t * (N + 1) + j] = math.Min(mtx[t * (N + 1) + j], mtx[(t - 1) * (N + 1) + j])
            }
            
            // the third transition: a case where we take a test
            if t_arr[j - 1] <= t {
                if mtx[t * (N + 1) + j] > mtx[(t - t_arr[j - 1]) * (N + 1) + (j - 1)] + tmp[j - 1] {
                    path[t * (N + 1) + j] = Pair{t - t_arr[j - 1], j - 1}
                }
                
                mtx[t * (N + 1) + j] = math.Min(mtx[t * (N + 1) + j], mtx[(t - t_arr[j - 1]) * (N + 1) + (j - 1)] + tmp[j - 1])
            }
        }
    }
    
    // restore an answer
    var from Pair
    cur_ind := Pair{T, N}
    for {
        from = path[cur_ind.first * (N + 1) + cur_ind.second]
        
        // we' ve reached the begiging of path
        if from.first == -1 {
            break
        }
        
        if from.first == cur_ind.first || from.second == cur_ind.second {
            cur_ind = from
        } else {    
            set_of_tests = append(set_of_tests, cur_ind.second)
            cur_ind = from
        }
    }
    
    // print an answer
    fmt.Println("The set of tests:")
    for i = len(set_of_tests) - 1; i >= 0; i-- {
        fmt.Printf("%d ", set_of_tests[i])
    }
    fmt.Println("\nThe probability:")
    fmt.Printf("%g\n", 1 - math.Pow(2.718281828, mtx[T * (N + 1) + N]))
}

