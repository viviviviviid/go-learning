package main

/// @title struct
// object와 비슷해서 map보다 더 유연함
// map을 이용하면, key와 value에 대한 타입을 정해줘야하기 때문에, 아래와 같은 유연한 이용이 불가능 하다.
//
//	{
//		"name": "nico",
//		"age": 18,
//		"favFood": ["strong"]
//	}
//
// 그래서 struct를 이용 // solidty에서 보던 그대로
// type person struct {
// 	name    string
// 	age     int
// 	favFood []string
// }
//
// func main() {
// 	favFood := []string{"kimchi", "ramen"}
// 	// nico := person{"nico", 18, favFood} // 이렇게 진행할 수 있지만, 어떤 내용이 어떤 필드인지 확인 불가
// 	nico := person{name: "nico", age: 18, favFood: favFood} // 그래서 여기와 같이 필드명을 적어주면 직관적임
// 	fmt.Println(nico)
// }

/// @title map
// func main() {
// 	// map[key_type]value_type 으로 map 선언
// 	nico := map[string]string{"name": "nico", "age": "12"}
// 	fmt.Println(nico)
// 	// nico := map[string]string{"name": "nico", "age": 12}	// map 선언시 key와 value 둘다 string 타입으로 선언. 즉 마지막 12는 오류가 발생
// 	// map도 range를 이용하여 반복문에 이용할 수 있음
// 	for key, value := range nico {
// 		fmt.Println(key, value)
// 	}
// 	// 이전과 같이 return값을 ignore 시킬 수도 있음.
// 	for _, value := range nico {
// 		fmt.Println(value)
// 	}
// }

/// @title Arrays -> 길이 제한이 없는 배열 : "slice" => "[]"
// func main() {
// // 길이가 제한된 배열 생성
// names := [5]string{"nico", "lynn", "dal"} // 길이 명시, 타입 명시
// names[3] = "alalal"
// names[4] = "alalal"
// names[5] = "alalal" // 길이가 5인 배열에 6개를 넣으려 했으므로 오류
// // 길이 제한이 없는 배열 생성 -> "slice"
// names := []string{"nico", "lynn", "dal"} // 길이 상관 없음, 타입은 명시
// names[3] = "alalal"                      // slice에 값을 추가할때는 이러한 방식으로 불가
// // append를 이용하여 배열에 값 추가 가능 // append(slice의 이름, 추가하고 싶은 값)
// append(names, "alalal")                // 하지만 이 append가 배열을 수정하는건 아님. // 수정된 내용을 return 해주는 것 뿐
// fmt.Println((append(names, "alalal"))) // 즉 이런식으로 사용해야함
// names = append(names, "alalal")		  // 또는 이런식
// 즉 수정해주는 js와 py와는 다른 시스템
// }

/// @title 포인터
// func main() {
// 	a := 2
// 	b := a  // a 값을 복사하는 것
// 	c := &a // a 메모리 주소를 복사하는 것 // a가 변하는 건 곧 c도 변하는 것
// 	a = 10  // b는 영향 없음 -> 메모리 주소 내용이 바뀐게 아니기 때문 // c에는 영향 있음
// 	fmt.Println(a, b, c)
// 	fmt.Println(&a, &b) // 메모리 주소 보기
// 	fmt.Println(*c)     // * 는 메모리 주소를 살펴볼 수 있음 // 현재 c에는 정수와 같은 값이 아니라, 메모리주소번지가 들어있기에 그 주소를 살펴보는것
// 	*c = 20             // c를 가지고 a를 수정하는 법 : c에는 a의 메모리주소번지가 있기 때문에 *를 통해 c의 값인 a 주소를 포인팅하고, 그 내용을 변경
// 	fmt.Println(a)
// }

/// @title switch
// // 기본 형태
//
//	func canIDrink(age int) bool {
//		switch age {
//		case 10:
//			return false
//		case 18:
//			return true
//		}
//		return false
//	}
//
// // if-else if의 반복을 깔끔하게 처리하기
//
//	func canIDrink(age int) bool {
//		switch {
//		case age < 18:
//			return false
//		case age == 18:
//			return true
//		case age > 50:
//			return false
//		}
//		return false
//	}
//
// // variable expression과 같이 사용하기
// func canIDrink(age int) bool {
// 	switch koreanAge := age + 2; koreanAge {
// 	case 10:
// 		return false
// 	case 18:
// 		return true
// 	}
// 	return false
// }

/// @title variable expression : if 조건에서 변수 생성해서 쓰기
// // if, else
// func canIDrink(age int) bool {
// 	// // 기본 형태
// 	// if age < 18 {
// 	// 	return false
// 	// }
// 	// return true
// 	// if문 내에서 변수 생성 // 세미콜론 이후로 생성된 변수 사용가능
// 	// 파이썬과 같이 한줄에 쓸 경우 세미콜론 // 그게 아니라면 세미콜론 없이 다음줄에 써도 가능
// 	if koreanAge := age + 2; koreanAge < 20 {
// 		fmt.Println(koreanAge)
// 		return false
// 	}
// 	return true
// }
// func main() {
// 	fmt.Println(canIDrink(16))
// }

/// @title go랭에는 for만 존재하고 파생상품 map, for in 들은 전혀 없다.
// // 그래도 순회는 있음 -> range
// // 기본 for문
// func superAdd(numbers ...int) {
// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Println(numbers[i])
// 	}
// }
// // range 사용
// func superAdd(numbers ...int) int {
// 	// range는 forEach 같이 순회하는 역할
// 	// range는 (index, element)를 리턴
// 	// 이때 인자를 하나만 넣으면 index만 출력되는데, 이는 0부터 시작
// 	for index, number := range numbers {
//	for '_', number := range numbers { // ignore
// 		fmt.Println((index))
// 		fmt.Println(number)
// 	}
// 	return 1 // 뭔가는 리턴해야해서
// }
// func main() {
// 	superAdd(1, 2, 3, 4, 5, 6)
// }

/// @title defer -> js의 on이라고 보면 편함. 함수코드가 return까지 실행되고 나서 끝나느게 아니라,
// // defer가 있는 코드라인으로 가서, 함수가 끝난 이후 실행해주는 메소드
// // 함수가 끝나고, 이미지같은 파일을 열던가 닫던가 api 요청을 보내던지 하는 확장성이 있음.
// func lenAndUpper(name string) (length int, uppercase string) {
// 	defer fmt.Println("I'm done") // 원래 main에 있던 코드. // defer로 인해 함수가 종료되고도 특정행위 가능
// 	length = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

/// @title go랭 func returnx 관련한 naked return.
// func lenAndUpper(name string) (length int, uppercase string) {
// 	length = len(name) // 위 return 타입 넣어주는 곳에서 선언함
// 	uppercase = strings.ToUpper(name)
// 	return // 위 return에서 선언을 해줬고, 밑 라인에서 변수 초기화를 해줬으므로,
// 	// return 안해줘도 되는 'naked return' (return할 variable을 굳이 꼭 명시하지 않아도 된다.)
// 	// 해줘도 에러가 안나긴 함
// }
// func main() {
// 	len, up := lenAndUpper("nico")
// 	fmt.Println(len, up)
// }

/// @title go랭에서 원하는 양만큼의 인자 넣기
// // 파라미터 타입앞에 온점 세개
// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }
// func main() {
// 	repeatMe("nico", "lynn", "dal")
// }

/// @title go언어는 return을 여러개 줄 수 있다.
// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }
// func main() {
// 	totalLength, upperName := lenAndUpper("nico")
//	//  하나만 받고싶은 경우
//  // 	totalLength, _ := lenAndUpper("nico")
// 	//  "_" 는 무시된 값
// 	fmt.Println(totalLength, upperName)
// 	fmt.Println(totalLength, upperName)
// }

/// @title 함수 만들기
// func multiply(a, b int) int {
// 	return a * b
// }
// func main() {
// 	fmt.Println((multiply(2, 2)))
// }

/// @title 변수 선언, 축약형
// func main() {
// 	// 프로그램의 진입점
// 	const name string = "vivivid"
// 	var major string = "intelligence"
// 	// 축약형 // func안에서만 가능하고 var(변수)에만 사용 가능
// 	// go가 자동으로 type 찾아서 정해줌
// 	age := 14
// 	// name = "hey"
// 	fmt.Println((age))
// 	fmt.Println(major)
// }

/// @title "=" vs ":="
// 고랭기지에서 "="와 ":="의 차이점은 변수 할당 방식에서 나타납니다.
// "="는 일반적인 프로그래밍 언어에서 사용되는 할당 연산자입니다. 이 연산자를 사용하여 변수에 값을 할당할 수 있습니다. 예를 들어:
// x = 5
// 위 코드는 변수 x에 값 5를 할당합니다. 이는 기존의 변수에 값을 할당하는 일반적인 방법입니다.
// 반면, ":="은 고랭기지(Golang)에서 사용되는 특수한 할당 연산자입니다. 이 연산자는 변수에 값을 할당하는 동시에 변수를 선언하는 역할을 합니다. 예를 들어:
// x := 5
// 위 코드는 변수 x를 선언하고 값을 5로 할당합니다. 이는 변수를 선언하면서 바로 초기값을 할당하는 간편한 방법입니다. 이러한 방식으로 변수를 선언하면 자료형 추론을 통해 변수의 타입을 자동으로 결정할 수 있습니다.
// 따라서, "="는 기존 변수에 값을 할당하는 데 사용되고, ":="는 변수를 선언하면서 값을 할당하는 데 사용됩니다. ":="은 변수 선언과 초기화를 동시에 처리하는 편리한 문법적 요소입니다.
