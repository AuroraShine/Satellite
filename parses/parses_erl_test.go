package parses

import (
	"bytes"
	"fmt"
	"testing"
)

func TestExtractOneElement(t *testing.T) {
	s := []byte("hello,1,[1,2,3],{simple,12}")
	r, sub, err := extractOneElement(s)
	if err != nil {
		t.Fatal("Error extract one element:", err)
	}
	fmt.Println("extract return:", string(r))
	fmt.Println("extract substring:", string(sub))
	if !bytes.Equal(r, []byte("hello")) {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("hello,1,[1,2,3],{simple,12}")
		r, sub, err := extractOneElement(s)
		if err != nil {
			b.Fatal("Error extract one element:", err)
		}
		if !bytes.Equal(r, []byte("hello")) {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
			b.Fatal("Error extract substring")
		}
	}
}

func TestExtractOneString(t *testing.T) {
	s := []byte("hello,1,[1,2,3],{simple,12}")
	r, sub, err := extractOneString(s)
	if err != nil {
		t.Fatal("Error extract one string:", err)
	}
	fmt.Println("extract return:", r)
	fmt.Println("extract substring:", string(sub))
	if r != "hello" {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("hello,1,[1,2,3],{simple,12}")
		r, sub, err := extractOneString(s)
		if err != nil {
			b.Fatal("Error extract one string:", err)
		}
		if r != "hello" {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
			b.Fatal("Error extract substring")
		}
	}
}

func TestExtractOneInt(t *testing.T) {
	s := []byte("15,atomsphere,[1,2,3],{simple,12}")
	r, sub, err := extractOneInt(s)
	if err != nil {
		t.Fatal("Error extract one int:", err)
	}
	fmt.Println("extract return:", r)
	fmt.Println("extract substring:", string(sub))
	if r != 15 {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("atomsphere,[1,2,3],{simple,12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("15,atomsphere,[1,2,3],{simple,12}")
		r, sub, err := extractOneInt(s)
		if err != nil {
			b.Fatal("Error extract one int:", err)
		}
		if r != 15 {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("atomsphere,[1,2,3],{simple,12}")) {
			b.Fatal("Error extract substring")
		}
	}
}

func TestExtractOneFloat64(t *testing.T) {
	s := []byte("3.14,atomsphere,[1,2,3],{simple,12}")
	r, sub, err := extractOneFloat64(s)
	if err != nil {
		t.Fatal("Error extract one float64:", err)
	}
	fmt.Println("extract return:", r)
	fmt.Println("extract substring:", string(sub))
	if r != 3.14 {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("atomsphere,[1,2,3],{simple,12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("3.14,atomsphere,[1,2,3],{simple,12}")
		r, sub, err := extractOneFloat64(s)
		if err != nil {
			b.Fatal("Error extract one float64:", err)
		}
		if r != 3.14 {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("atomsphere,[1,2,3],{simple,12}")) {
			b.Fatal("Error extract substring")
		}
	}
}

func TestExtractOneBool(t *testing.T) {
	s := []byte("true,atomsphere,[1,2,3],{simple,12}")
	r, sub, err := extractOneBool(s)
	if err != nil {
		t.Fatal("Error extract one bool:", err)
	}
	fmt.Println("extract return:", r)
	fmt.Println("extract substring:", string(sub))
	if r != true {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("atomsphere,[1,2,3],{simple,12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("true,atomsphere,[1,2,3],{simple,12}")
		r, sub, err := extractOneBool(s)
		if err != nil {
			b.Fatal("Error extract one bool:", err)
		}
		if r != true {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("atomsphere,[1,2,3],{simple,12}")) {
			b.Fatal("Error extract substring")
		}
	}
}

func TestExtractOneList(t *testing.T) {
	s := []byte("[apple,orange],1,[1,2,3],{simple,12}")
	r, sub, err := extractOneList(s)
	if err != nil {
		t.Fatal("Error extract one list:", err)
	}
	fmt.Println("extract return:", string(r))
	fmt.Println("extract substring:", string(sub))
	if !bytes.Equal(r, []byte("[apple,orange]")) {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("[apple,orange],1,[1,2,3],{simple,12}")
		r, sub, err := extractOneList(s)
		if err != nil {
			b.Fatal("Error extract one list:", err)
		}
		if !bytes.Equal(r, []byte("[apple,orange]")) {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
			b.Fatal("Error extract substring")
		}
	}
}

func TestExtractOneTuple(t *testing.T) {
	s := []byte("{stars,514},1,[1,2,3],{simple,12}")
	r, sub, err := extractOneTuple(s)
	if err != nil {
		t.Fatal("Error extract one tuple:", err)
	}
	fmt.Println("extract return:", string(r))
	fmt.Println("extract substring:", string(sub))
	if !bytes.Equal(r, []byte("{stars,514}")) {
		t.Fatal("Error extract return")
	}
	if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
		t.Fatal("Error extract substring")
	}
}

func BenchmarkExtractOneTuple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("{stars,514},1,[1,2,3],{simple,12}")
		r, sub, err := extractOneTuple(s)
		if err != nil {
			b.Fatal("Error extract one tuple:", err)
		}
		if !bytes.Equal(r, []byte("{stars,514}")) {
			b.Fatal("Error extract return")
		}
		if !bytes.Equal(sub, []byte("1,[1,2,3],{simple,12}")) {
			b.Fatal("Error extract substring")
		}
	}
}

func TestTrimList(t *testing.T) {
	s := []byte("[apple,orange]")
	r, err := trimList(s)
	if err != nil {
		t.Fatal("Error trim list:", err)
	}
	fmt.Println("trim return:", string(r))
	if !bytes.Equal(r, []byte("apple,orange")) {
		t.Fatal("Error trim result")
	}
}

func BenchmarkTrimList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("[apple,orange]")
		r, err := trimList(s)
		if err != nil {
			b.Fatal("Error trim list:", err)
		}
		if !bytes.Equal(r, []byte("apple,orange")) {
			b.Fatal("Error trim result")
		}
	}
}

func TestTrimTuple(t *testing.T) {
	s := []byte("{stars,514}")
	r, err := trimTuple(s)
	if err != nil {
		t.Fatal("Error trim tuple:", err)
	}
	fmt.Println("trim return:", string(r))
	if !bytes.Equal(r, []byte("stars,514")) {
		t.Fatal("Error trim result")
	}
}

func BenchmarkTrimTuple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("{stars,514}")
		r, err := trimTuple(s)
		if err != nil {
			b.Fatal("Error trim tuple:", err)
		}
		if !bytes.Equal(r, []byte("stars,514")) {
			b.Fatal("Error trim result")
		}
	}
}

func TestRepairTrim(t *testing.T) {
	s := []byte("apple,1")
	r := repairTrim(s)
	if !bytes.Equal(r, []byte("apple,1,")) {
		t.Fatal("Error repair trim list or tuple")
	}
}

func BenchmarkRepairTrim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []byte("apple,1")
		r := repairTrim(s)
		if !bytes.Equal(r, []byte("apple,1,")) {
			b.Fatal("Error repair trim list or tuple")
		}
	}
}

func TestDecodeOneParameter(t *testing.T) {
	type subtest struct {
		Name   string `erl:"string"`
		Number int    `erl:"int"`
	}
	type test struct {
		Name    string  `erl:"string"`
		Content string  `erl:"string"`
		Number  int     `erl:"int"`
		List    []int   `erl:"list"`
		Tuple   subtest `erl:"tuple"`
	}
	in := []byte("test_name,hello,1,[1,2,3],{simple,12}")
	out := test{}
	err := decodeOneParameter(in, &out)
	if err != nil {
		t.Fatal("Error decode one parameter:", err)
	}
	fmt.Println(out)
}

func BenchmarkDecodeOneParameter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type subtest struct {
			Name   string `erl:"string"`
			Number int    `erl:"int"`
		}
		type test struct {
			Name    string  `erl:"string"`
			Content string  `erl:"string"`
			Number  int     `erl:"int"`
			List    []int   `erl:"list"`
			Tuple   subtest `erl:"tuple"`
		}
		in := []byte("test_name,hello,1,[1,2,3],{simple,12}")
		out := test{}
		err := decodeOneParameter(in, &out)
		if err != nil {
			b.Fatal("Error decode one parameter:", err)
		}
	}
}

func TestDecodeOneParameter2(t *testing.T) {
	type subtest struct {
		Name   string `erl:"string"`
		Number int    `erl:"int"`
	}
	type test struct {
		Name    string    `erl:"string"`
		Content string    `erl:"string"`
		Number  int       `erl:"int"`
		List    []subtest `erl:"list"`
	}
	in := []byte("asteroid,earth,9,[{apple,1},{orange,2},{lemon,5}]")
	out := test{}
	err := decodeOneParameter(in, &out)
	if err != nil {
		t.Fatal("Error decode one parameter:", err)
	}
	fmt.Println(out)
}

func BenchmarkDecodeOneParameter2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type subtest struct {
			Name   string `erl:"string"`
			Number int    `erl:"int"`
		}
		type test struct {
			Name    string    `erl:"string"`
			Content string    `erl:"string"`
			Number  int       `erl:"int"`
			List    []subtest `erl:"list"`
		}
		in := []byte("asteroid,earth,9,[{apple,1},{orange,2},{lemon,5}]")
		out := test{}
		err := decodeOneParameter(in, &out)
		if err != nil {
			b.Fatal("Error decode one parameter:", err)
		}
	}
}

func TestDecodeOneParameter3(t *testing.T) {
	type subtest struct {
		Name   string `erl:"string"`
		Number int    `erl:"int"`
		List   []int  `erl:"list"`
	}
	type test struct {
		List []subtest `erl:"list"`
	}
	in := []byte("[{apple,1,[2,3]},{orange,2,[1,5,4]},{lemon,5,[]}]")
	out := test{}
	err := decodeOneParameter(in, &out)
	if err != nil {
		t.Fatal("Error decode one parameter:", err)
	}
	fmt.Println(out)
}

func BenchmarkDecodeOneParameter3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type subtest struct {
			Name   string `erl:"string"`
			Number int    `erl:"int"`
			List   []int  `erl:"list"`
		}
		type test struct {
			List []subtest `erl:"list"`
		}
		in := []byte("[{apple,1,[2,3]},{orange,2,[1,5,4]},{lemon,5,[]}]")
		out := test{}
		err := decodeOneParameter(in, &out)
		if err != nil {
			b.Fatal("Error decode one parameter:", err)
		}
	}
}

func TestDecodeOneParameter4(t *testing.T) {
	type subsub struct {
		Name    string `erl:"string"`
		Content string `erl:"string"`
		Value   int    `erl:"int"`
	}
	type subtest1 struct {
		Name  string   `erl:"string"`
		Index int      `erl:"int"`
		Sub   []subsub `erl:"list"`
	}
	type subtest2 struct {
		Name    string `erl:"string"`
		SubName string `erl:"string"`
	}
	type subtest3 struct {
		Name string `erl:"string"`
		Up   int    `erl:"int"`
		Down int    `erl:"int"`
		Sub  subsub `erl:"tuple"`
	}
	type test struct {
		Name    string        `erl:"string"`
		Index   int           `erl:"int"`
		Options []interface{} `erl:"list"`
	}
	in := []byte("test,1,[{sub_1,2,[{subs,apple,3},{subs,lemon,5},{subs,banana,1}]},{sub_3,2,1,{subs,peach,7}},{sub_2,asteroid}]")
	out := test{}
	fmt.Println(out)
	MParsesErl = make(map[string]interface{})
	MParsesErl["sub_1"] = subtest1{}
	MParsesErl["sub_2"] = subtest2{}
	MParsesErl["sub_3"] = subtest3{}
	err := decodeOneParameter(in, &out)
	if err != nil {
		t.Fatal("Error decode one parameter:", err)
	}
	fmt.Println(out)
}

func BenchmarkDecodeOneParameter4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type subsub struct {
			Name    string `erl:"string"`
			Content string `erl:"string"`
			Value   int    `erl:"int"`
		}
		type subtest1 struct {
			Name  string   `erl:"string"`
			Index int      `erl:"int"`
			Sub   []subsub `erl:"list"`
		}
		type subtest2 struct {
			Name    string `erl:"string"`
			SubName string `erl:"string"`
		}
		type subtest3 struct {
			Name string `erl:"string"`
			Up   int    `erl:"int"`
			Down int    `erl:"int"`
			Sub  subsub `erl:"tuple"`
		}
		type test struct {
			Name    string        `erl:"string"`
			Index   int           `erl:"int"`
			Options []interface{} `erl:"list"`
		}
		in := []byte("test,1,[{sub_1,2,[{subs,apple,3},{subs,lemon,5},{subs,banana,1}]},{sub_3,2,1,{subs,peach,7}},{sub_2,asteroid}]")
		out := test{}
		MParsesErl = make(map[string]interface{})
		MParsesErl["sub_1"] = subtest1{}
		MParsesErl["sub_2"] = subtest2{}
		MParsesErl["sub_3"] = subtest3{}
		err := decodeOneParameter(in, &out)
		if err != nil {
			b.Fatal("Error decode one parameter:", err)
		}
	}
}

func TestDecode(t *testing.T) {
	type subsub struct {
		Name    string `erl:"string" json:"name"`
		Content string `erl:"string" json:"content"`
		Value   int    `erl:"int" json:"value"`
	}
	type subtest1 struct {
		Name  string   `erl:"string" json:"name"`
		Index int      `erl:"int" json:"index"`
		Sub   []subsub `erl:"list" json:"sub"`
	}
	type subtest2 struct {
		Name    string `erl:"string" json:"name"`
		SubName string `erl:"string" json:"subname"`
	}
	type subtest3 struct {
		Name string `erl:"string" json:"name"`
		Up   int    `erl:"int" json:"up"`
		Down int    `erl:"int" json:"down"`
		Sub  subsub `erl:"tuple" json:"subsub"`
	}
	type test struct {
		Name    string        `erl:"string" json:"name"`
		Index   int           `erl:"int" json:"index"`
		Options []interface{} `erl:"list" json:"options"`
	}
	type other struct {
		Use   bool    `erl:"bool" json:"use"`
		Speed float64 `erl:"float64" json:"speed"`
	}
	type testlist struct {
		Tests  []test  `erl:"list" json:"tests"`
		Others []other `erl:"list" json:"others"`
	}
	in := []byte("test,1,[{sub_1,2,[{subs,apple,3},{subs,lemon,5},{subs,banana,1}]},{sub_3,2,1,{subs,peach,7}},{sub_2,asteroid}]")
	out := testlist{}
	fmt.Println(out)
	MParsesErl = make(map[string]interface{})
	MParsesErl["test"] = test{}
	MParsesErl["sub_1"] = subtest1{}
	MParsesErl["sub_2"] = subtest2{}
	MParsesErl["sub_3"] = subtest3{}
	err := decode(in, &out)
	if err != nil {
		t.Fatal("Error decode:", err)
	}
	fmt.Println(out)
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type subsub struct {
			Name    string `erl:"string" json:"name"`
			Content string `erl:"string" json:"content"`
			Value   int    `erl:"int" json:"value"`
		}
		type subtest1 struct {
			Name  string   `erl:"string" json:"name"`
			Index int      `erl:"int" json:"index"`
			Sub   []subsub `erl:"list" json:"sub"`
		}
		type subtest2 struct {
			Name    string `erl:"string" json:"name"`
			SubName string `erl:"string" json:"subname"`
		}
		type subtest3 struct {
			Name string `erl:"string" json:"name"`
			Up   int    `erl:"int" json:"up"`
			Down int    `erl:"int" json:"down"`
			Sub  subsub `erl:"tuple" json:"subsub"`
		}
		type test struct {
			Name    string        `erl:"string" json:"name"`
			Index   int           `erl:"int" json:"index"`
			Options []interface{} `erl:"list" json:"options"`
		}
		type other struct {
			Use   bool    `erl:"bool" json:"use"`
			Speed float64 `erl:"float64" json:"speed"`
		}
		type testlist struct {
			Tests  []test  `erl:"list" json:"tests"`
			Others []other `erl:"list" json:"others"`
		}
		in := []byte("test,1,[{sub_1,2,[{subs,apple,3},{subs,lemon,5},{subs,banana,1}]},{sub_3,2,1,{subs,peach,7}},{sub_2,asteroid}]")
		out := testlist{}
		MParsesErl = make(map[string]interface{})
		MParsesErl["test"] = test{}
		MParsesErl["sub_1"] = subtest1{}
		MParsesErl["sub_2"] = subtest2{}
		MParsesErl["sub_3"] = subtest3{}
		err := decode(in, &out)
		if err != nil {
			b.Fatal("Error decode:", err)
		}
	}
}

func TestUnmarshal(t *testing.T) {
	type subsub struct {
		Name    string `erl:"string" json:"name"`
		Content string `erl:"string" json:"content"`
		Value   int    `erl:"int" json:"value"`
	}
	type subtest1 struct {
		Name  string   `erl:"string" json:"name"`
		Index int      `erl:"int" json:"index"`
		Sub   []subsub `erl:"list" json:"sub"`
	}
	type subtest2 struct {
		Name    string `erl:"string" json:"name"`
		SubName string `erl:"string" json:"subname"`
	}
	type subtest3 struct {
		Name string `erl:"string" json:"name"`
		Up   int    `erl:"int" json:"up"`
		Down int    `erl:"int" json:"down"`
		Sub  subsub `erl:"tuple" json:"subsub"`
	}
	type test struct {
		Name    string        `erl:"string" json:"name"`
		Index   int           `erl:"int" json:"index"`
		Options []interface{} `erl:"list" json:"options"`
	}
	type other struct {
		Name  string  `erl:"string" json:"name"`
		Use   bool    `erl:"bool" json:"use"`
		Speed float64 `erl:"float64" json:"speed"`
	}
	type testlist struct {
		Tests  []test  `erl:"list" json:"tests"`
		Others []other `erl:"list" json:"others"`
	}
	in := []byte("{test,1,[{sub_1,2,[{subs,apple,3},{subs,lemon,5},{subs,banana,1}]},{sub_3,2,1,{subs,peach,7}},{sub_2,asteroid}]}.\r\n{other,true,3.66}.\r\n{test,2,[{sub_1,2,[{subs,apple,3},{subs,lemon,5},{subs,banana,1}]},{sub_3,2,1,{subs,peach,7}},{sub_2,asteroid}]}.\r\n{other,false,12.96}.\r\n{test,3,[{sub_1,2,[{subs,apple,3},{subs,lemon,5},{subs,banana,1}]},{sub_3,2,1,{subs,peach,7}},{sub_2,asteroid}]}.")
	out := testlist{}
	fmt.Println(out)
	MParsesErl = make(map[string]interface{})
	MParsesErl["test"] = test{}
	MParsesErl["other"] = other{}
	MParsesErl["sub_1"] = subtest1{}
	MParsesErl["sub_2"] = subtest2{}
	MParsesErl["sub_3"] = subtest3{}
	err := unmarshal(in, &out)
	if err != nil {
		t.Fatal("Error unmarshal:", err)
	}
	fmt.Println(out)
}

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type subsub struct {
			Name    string `erl:"string" json:"name"`
			Content string `erl:"string" json:"content"`
			Value   int    `erl:"int" json:"value"`
		}
		type subtest1 struct {
			Name  string   `erl:"string" json:"name"`
			Index int      `erl:"int" json:"index"`
			Sub   []subsub `erl:"list" json:"sub"`
		}
		type subtest2 struct {
			Name    string `erl:"string" json:"name"`
			SubName string `erl:"string" json:"subname"`
		}
		type subtest3 struct {
			Name string `erl:"string" json:"name"`
			Up   int    `erl:"int" json:"up"`
			Down int    `erl:"int" json:"down"`
			Sub  subsub `erl:"tuple" json:"subsub"`
		}
		type test struct {
			Name    string        `erl:"string" json:"name"`
			Index   int           `erl:"int" json:"index"`
			Options []interface{} `erl:"list" json:"options"`
		}
		type other struct {
			Name  string  `erl:"string" json:"name"`
			Use   bool    `erl:"bool" json:"use"`
			Speed float64 `erl:"float64" json:"speed"`
		}
		type testlist struct {
			Tests  []test  `erl:"list" json:"tests"`
			Others []other `erl:"list" json:"others"`
		}
		in := []byte("{test,1,[{sub_1,2,[{subs,apple,3},{subs,lemon,5},{subs,banana,1}]},{sub_3,2,1,{subs,peach,7}},{sub_2,asteroid}]}.\r\n{other,true,3.66}.\r\n{test,2,[{sub_1,2,[{subs,apple,3},{subs,lemon,5},{subs,banana,1}]},{sub_3,2,1,{subs,peach,7}},{sub_2,asteroid}]}.\r\n{other,false,12.96}.\r\n{test,3,[{sub_1,2,[{subs,apple,3},{subs,lemon,5},{subs,banana,1}]},{sub_3,2,1,{subs,peach,7}},{sub_2,asteroid}]}.")
		out := testlist{}
		MParsesErl = make(map[string]interface{})
		MParsesErl["test"] = test{}
		MParsesErl["other"] = other{}
		MParsesErl["sub_1"] = subtest1{}
		MParsesErl["sub_2"] = subtest2{}
		MParsesErl["sub_3"] = subtest3{}
		err := unmarshal(in, &out)
		if err != nil {
			b.Fatal("Error unmarshal:", err)
		}
	}
}

func TestWrapOneElement(t *testing.T) {
	var s = "hello"
	r, err := wrapOneElement(s)
	if err != nil {
		t.Fatal("Error wrap one element:", err)
	}
	fmt.Println(string(r))
}

func TestWrapOneElementString(t *testing.T) {
	var s = "apple"
	r, err := wrapOneElement(s)
	if err != nil {
		t.Fatal("Error wrap one element:", err)
	}
	fmt.Println(string(r))
}

func TestWrapOneElementInt(t *testing.T) {
	var s = 2
	r, err := wrapOneElement(s)
	if err != nil {
		t.Fatal("Error wrap one element:", err)
	}
	fmt.Println(string(r))
}

func TestWrapOneElementBool(t *testing.T) {
	var s = false
	r, err := wrapOneElement(s)
	if err != nil {
		t.Fatal("Error wrap one element:", err)
	}
	fmt.Println(string(r))
}

func TestWrapOneElementFloat64(t *testing.T) {
	var s = 3.1415926
	r, err := wrapOneElement(s)
	if err != nil {
		t.Fatal("Error wrap one element:", err)
	}
	fmt.Println(string(r))
}