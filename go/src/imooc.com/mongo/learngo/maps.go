package main

/**
Map的操作
1、创建: make(map[string]int)
2、获取元素: m[key]
3、key不存在时, 获得Value类型的初始值
4、用Value, ok:= m[key]来判断是否存在key
5、用delete删除一个key
Map的遍历
1、使用range遍历key, 或者遍历key, value对
2、不保证遍历顺序, 如需顺序, 需手动对key排序
3、使用len获得元素个数
Map的key
1、map使用哈希表, 必须可以比较相等
2、除了slice, map, function的内建类型都可以作为key
3、Struct类型不包含上述字段，也可以作为key

*/
import "fmt"

func main() {

	m1 := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) //empty map
	var m3 map[string]int      //nil

	fmt.Println(m1, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m1["course"]
	fmt.Println(courseName)

	fmt.Println("Deleting values")
	name, ok := m1["name"]
	fmt.Println(name, ok)

	delete(m1, "name")
	name, ok = m1["name"]
	fmt.Println(name, ok)
}
