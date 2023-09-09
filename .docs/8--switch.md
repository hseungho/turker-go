# switch문

## switch문 동작 원리
switch문은 값에 따라 다른 로직을 수행할 때 사용한다.
```go
switch 비굣값 {
case 값1: 
  문장
case 값2:
  문장
default:
  문장
}
```

`switch` 키워드 다음으로 비굣값이 온다. 첫 번째 `case`부터 값을 검사해가게 되며, 만일 비굣값과 `case` 값이 같으면
해당 `case` 문장을 수행하고 `switch`문을 종료한다. 모든 `case`값과 일치하지 않으면 `default` 문장을 수행하게 되며,
`default`는 생략 가능하다.

[> 코드로 가기](../ch10/ex10.1/ex10.1.go)

## switch문은 언제 쓰는가?
`switch`문을 이용하면 복잡한 `if else`문을 보기 좋게 정리할 수 있다.

[> `if else` 코드로 가기](../ch10/ex10.2/ex10.2.go)  
[> `switch` 코드로 가기](../ch10/ex10.3/ex10.3.go)

## 다양한 switch문 형태
### 한 번에 여러 값 비교
하나의 `case`는 하나 이상의 값을 비교할 수 있다. 각 값은 `,`쉼표로 구분한다.

[> 코드로 가기](../ch10/ex10.4/ex10.4.go)

### 조건문 비교
`switch`문의 동작을 응용하면 단순히 값만 비교하는게 아닌 `if`문처럼 `true`가 되는 조건문을 검사할 수 있다.
`switch`문은 비굣값과 `case`의 값이 같아지는 경우를 찾는 구문이기 때문에 비굣값을 `true`로 할 경우 `case`의 조건문이
`true`가 되는 경우가 실행됩니다.

`switch` 다음에 비굣값을 적지 않는 경우 default값으로 `true`를 사용합니다.
```go
switch true {
  ...
}
```
위 식을 다음과 같이 줄여 쓸 수 있다.
```go
switch {
  ...
}
```
[> 코드로 가기](../ch10/ex10.5/ex10.5.go)

### switch 초기문
`if`문과 마찬가지로 `switch`문에서도 초기문을 넣을 수 있다.
```go
switch 초기문; 비굣값 {
case 값1:
  ...
case 값2:
  ...
default:
}
```
초기문에서 선언한 변수는 `switch`문이 종료되면 소멸되기 때문에 `switch`문 밖에서 사용할 수 없다.

[> 코드로 가기 - 1](../ch10/ex10.6/ex10.6.go)
[> 코드로 가기 - 2](../ch10/ex10.7/ex10.7.go)

## const 열거값과 switch
`const` 열거값에 따라 수행되는 로직을 변경할 때 `switch`문을 주로 사용한다.

[> 코드로 가기](../ch10/ex10.8/ex10.8.go)

> 위 예제 코드에서 만일 Violet 같은 새로운 색깔이 추가되면 colorToString() 함수도 수정해줘야 한다.
> 이런 경우를 커플링됐다고 하거나 결합되어 있다고 말한다. 즉 열거값이 수정되면 연관된 모든 switch case문도 수정해줘야 한다.
> 그래서 열거값에 연관된 switch case가 많아질수록 작은 수정에도 많은 코드가 변경되어야 하는 산탄총 수술 문제가 발생한다.
> 그래서 하나의 열거값에 연관된 switch case는 최대한 줄이는 게 좋다.

## `break`와 `fallthrough` 키워드
일반적으로 다른 언어에서는 `switch`문의 각 `case` 종료 시에 `break`문을 사용해야 다음 `case`로 코드가 이어서 실행되지 않는다.
하지만 Go언어에서는 `break`를 사용하지 않아도 `case `하나를 실행 후 자동으로 `switch`문을 빠져나가게 된다.

[> 코드로 가기](../ch10/ex10.9/ex10.9.go)

그런데 만약 하나의 `case`문 실행 후 다음 `case`문까지 실행하고 싶을 때는 `fallthrough` 키워드를 사용하면 된다. 
`case` 마지막에 `fallthrough` 키워드를 사용하면 다음 `case`까지 실행된다.

[> 코드로 가기](../ch10/ex10.10/ex10.10.go)

> 하지만 fallthrough 키워드는 코드를 보는 사람에게 혼동을 일으킬 수 있으니 되도록 사용하지 않는 것을 권장한다.

## 핵심요약
switch 문은 값에 따라서 다른 문장을 실행할 때 사용되는 구문이다.
```go
switch 비굣값 {
case 값1:
  문장
case 값2:
  문장
default:
  문장
}
```
```go
switch {
case 조건문1:
  문장
case 조건문2:
  문장
default:
  문장
}
```
```go
switch 초기문; 비굣값 {
case 값1:
  문장
case 값2:
  문장
default:
  문장
}
```
