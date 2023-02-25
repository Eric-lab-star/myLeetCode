# Two sum

## map을 이용한 해결법

```go
func twoSum(nums []int, target int) []int {
    s := make(map[int]int)
    for idx, num := range nums {

        if pos, ok := s[target-num]; ok {
            return []int{pos, idx}
        }
        s[num] = idx
    }
    return []int{}
}
```

인풋이 [3,2,4]라고 할 때, 맵을 이용해서 6이 나오는 두 수의 인덱스를 맵을 이용해서 구하려고 하면 먼저 자료구조를 맵 형태로 바꾸어 주어야힌다.

```go
    s := map[int]int{
        3 : 0,
        2: 1,
        4 :2,
    }
```

인풋의 정보를 s와 같은 형태로 만들어 주면, 합이 6이되는 두 숫자 x,y로부터 s[x], s[y]를 구할 수 있다. 이때 y = 6 - x 임으로,
s[x], s[6 - x]가 원하는 결과가 된다.
x를 찾기위한 알고리즘을 구현하면,

```go
    s := make(map[int]int)
    for idx, num := range nums {
        if pos, ok := s[target-num]; ok {
            return []int{pos, idx}
        }
    s[num] = idx
    }
```

이렇게 된다.

## for loop를 두 번 사용한 brute force 방법

```go
func twoSum(nums []int, target int) []int {
    ret := []int{}

    for i := 0; i < len(nums) -1; i++ {
        for j := i+1; j < len(nums); j++ {
            if nums[i]+nums[j] == target && i != j {
                return []int{i, j}
            }
        }
    }
    return ret
}
```
