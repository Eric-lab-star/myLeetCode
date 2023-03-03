알게 된 것

```go

func romanToInt(s string) int {
	romanNumberDict := map[rune]uint16{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var ret uint16
	var pre uint16

	for _, char := range s {
		cur := romanNumberDict[char]
		if cur > pre {
			ret -= pre * 2
		}
		ret += cur
		pre = cur
	}

	return int(ret)
}

```

go 에서는 글자를 저장하기위해서 rune을 사용한다.
uint16의 범위는 0에서 65535사이의 정수이다.

문자열을 for range 루프에 넣을 수 있다.
주어진 문제를 정확하게 이해할 수 있다면 if문을 줄일 수 있다.
pre = cur 과 같이 현제 값을 이전 값으로 넘겨주는 것이 자주사용된다.
