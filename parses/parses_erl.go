package parses

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// MParsesErl global mapping struct search by type 'interface{}'
// put all struct which will be push into 'interface{}' before call function 'Marshal' or 'Unmarshal'(recommend initialized in function 'init()')
// ...
// example:
// MParsesErl = make(map[string]interface{})
// MParsesErl["sub_1"] = subtest1{}
// MParsesErl["sub_2"] = subtest2{}
// MParsesErl["sub_3"] = subtest3{}
// ...
var MParsesErl map[string]interface{}

// Unmarshal stream
func Unmarshal(in []byte, out interface{}) (err error) {
	return unmarshal(in, out)
}

// Unmarshal from file
func UnmarshalFrom(file string, out interface{}) (err error) {
	return unmarshalFrom(file, out)
}

// Marshal stream
func Marshal(in []byte, t interface{}) (out []byte, err error) {
	return marshal(in, t)
}

// Marshal to file
func MarshalTo(file string, t interface{}) (err error) {
	return marshalTo(file, t)
}

// decode
func extractOneElement(s []byte) (r []byte, sub []byte, err error) {
	// check bytes whether empty?
	if bytes.Equal(s, []byte("")) {
		err = errors.New("bytes is empty now")
		return r, sub, err
	}
	// erlang bytes contains ',' for element
	if !bytes.Contains(s, []byte(",")) {
		err = errors.New("bytes don't have element")
		return r, sub, err
	}
	// find element by ','
	index := bytes.Index(s, []byte(","))
	for i := 0; i < index; i++ {
		r = append(r, s[i])
	}
	// check element without '{}' or '[]'
	if bytes.ContainsAny(r, "{&}&[&]") {
		err = errors.New("bytes may extract list or tuple first")
		return r, sub, err
	}
	sub = bytes.TrimPrefix(s, r)
	sub = bytes.TrimPrefix(sub, []byte(","))
	return r, sub, err
}

func extractOneString(s []byte) (r string, sub []byte, err error) {
	// first, call function extract one element
	sr, sub, err := extractOneElement(s)
	if err != nil {
		return r, sub, err
	}
	// convert []byte -> string
	r = string(sr)
	return r, sub, err
}

func extractOneInt(s []byte) (r int, sub []byte, err error) {
	// first, call function extract one element
	sr, sub, err := extractOneElement(s)
	if err != nil {
		return r, sub, err
	}
	// convert []byte -> int
	r, err = strconv.Atoi(string(sr))
	if err != nil {
		err = errors.New("parameter may not integer")
		return r, sub, err
	}
	return r, sub, err
}

func extractOneFloat64(s []byte) (r float64, sub []byte, err error) {
	// first, call function extract one element
	sr, sub, err := extractOneElement(s)
	if err != nil {
		return r, sub, err
	}
	// convert []byte -> float64
	r, err = strconv.ParseFloat(string(sr), 64)
	if err != nil {
		err = errors.New("parameter may not float64")
		return r, sub, err
	}
	return r, sub, err
}

func extractOneBool(s []byte) (r bool, sub []byte, err error) {
	// first, call function extract one element
	sr, sub, err := extractOneElement(s)
	if err != nil {
		return r, sub, err
	}
	// convert []byte -> bool
	if string(sr) == "false" {
		r = false
	} else if string(sr) == "true" {
		r = true
	} else {
		err = errors.New("parameter may not bool")
		return r, sub, err
	}
	return r, sub, err
}

func extractOneList(s []byte) (r []byte, sub []byte, err error) {
	// erlang bytes contains '[]' for list
	if !bytes.Contains(s, []byte("[")) || !bytes.Contains(s, []byte("]")) {
		err = errors.New("bytes don't have list")
		return r, sub, err
	}
	// check '[' & ']' numbers should equal
	if bytes.Count(s, []byte("[")) != bytes.Count(s, []byte("]")) {
		err = errors.New("bytes list illegal, list may not intergrity")
		return r, sub, err
	}
	// check '[' & ']' sequence
	if bytes.Index(s, []byte("[")) > bytes.Index(s, []byte("]")) {
		err = errors.New("bytes list illegal, sequence of list may not right")
		return r, sub, err
	}
	// extract erlang list
	count := 0
	start := bytes.Index(s, []byte("["))
	for i := start; i < len(s); i++ {
		if s[i] == '[' {
			count++
		} else if s[i] == ']' {
			count--
		}
		r = append(r, s[i])
		if count == 0 {
			break
		}
	}
	sub = bytes.TrimPrefix(s, r)
	sub = bytes.TrimPrefix(sub, []byte(","))
	return r, sub, err
}

func extractOneTuple(s []byte) (r []byte, sub []byte, err error) {
	// erlang bytes contains '{}' for tuple
	if !bytes.Contains(s, []byte("{")) || !bytes.Contains(s, []byte("}")) {
		err = errors.New("bytes don't have tuple")
		return r, sub, err
	}
	// check '{' & '}' numbers should equal
	if bytes.Count(s, []byte("{")) != bytes.Count(s, []byte("}")) {
		err = errors.New("bytes tuple illegal, tuple may not intergrity")
		return r, sub, err
	}
	// check '{' & '}' sequence
	if bytes.Index(s, []byte("{")) > bytes.Index(s, []byte("}")) {
		err = errors.New("bytes tuple illegal, sequence of tuple may not right")
		return r, sub, err
	}
	// extract erlang tuple
	count := 0
	start := bytes.Index(s, []byte("{"))
	for i := start; i < len(s); i++ {
		if s[i] == '{' {
			count++
		} else if s[i] == '}' {
			count--
		}
		r = append(r, s[i])
		if count == 0 {
			break
		}
	}
	sub = bytes.TrimPrefix(s, r)
	sub = bytes.TrimPrefix(sub, []byte(","))
	return r, sub, err
}

func trimList(s []byte) (r []byte, err error) {
	// check list whether start '[' and end ']'
	if len(s) == 0 || s[0] != '[' || s[len(s)-1] != ']' {
		err = errors.New("bytes is not normal list")
		return r, err
	}
	// trim '[]' and return
	r = bytes.TrimPrefix(s, []byte("["))
	r = bytes.TrimSuffix(r, []byte("]"))
	return r, err
}

func trimTuple(s []byte) (r []byte, err error) {
	// check tuple whether start '{' and end '}'
	if len(s) == 0 || s[0] != '{' || s[len(s)-1] != '}' {
		err = errors.New("bytes is not normal tuple")
		return r, err
	}
	// trim '{}' and return
	r = bytes.TrimPrefix(s, []byte("{"))
	r = bytes.TrimSuffix(r, []byte("}"))
	return r, err
}

func repairTrim(s []byte) (r []byte) {
	// repair bytes after call trim function
	r = append(s, ',')
	return r
}

func decodeOneParameter(in []byte, out interface{}) (err error) {
	// get pointer's value...
	var rType = reflect.TypeOf(out)
	var rValue = reflect.ValueOf(out)
	// check the out type kind
	if rType.Kind() != reflect.Ptr {
		err = errors.New("out interface should be struct pointer")
		return err
	}
	// get real variable value...
	rType = rType.Elem()
	rValue = rValue.Elem()
	// switch the kind of type...
	switch rType.Kind() {
	case reflect.Struct:
		var rem = in
		// traverse struct fields
		for i := 0; i < rType.NumField(); i++ {
			// get struct field value...
			t := rType.Field(i)
			f := rValue.Field(i)
			// parse tag
			tag := t.Tag.Get("erl")
			fields := strings.Split(tag, ",")
			if len(fields) > 1 {
				tag = fields[0]
			}
			// swich the tag & parse element...
			switch tag {
			case "string":
				r, sub, err := extractOneString(rem)
				if err != nil {
					return err
				}
				rem = sub
				f.Set(reflect.ValueOf(r))
			case "int":
				r, sub, err := extractOneInt(rem)
				if err != nil {
					return err
				}
				rem = sub
				f.Set(reflect.ValueOf(r))
			case "float64":
				r, sub, err := extractOneFloat64(rem)
				if err != nil {
					return err
				}
				rem = sub
				f.Set(reflect.ValueOf(r))
			case "bool":
				r, sub, err := extractOneBool(rem)
				if err != nil {
					return err
				}
				rem = sub
				f.Set(reflect.ValueOf(r))
			case "list":
				r, sub, err := extractOneList(rem)
				if err != nil {
					return err
				}
				rem = sub
				r, err = trimList(r)
				if err != nil {
					return err
				}
				r = repairTrim(r)
				err = decodeOneParameter(r, f.Addr().Interface())
				if err != nil {
					return err
				}
			case "tuple":
				r, sub, err := extractOneTuple(rem)
				if err != nil {
					return err
				}
				rem = sub
				r, err = trimTuple(r)
				if err != nil {
					return err
				}
				r = repairTrim(r)
				err = decodeOneParameter(r, f.Addr().Interface())
				if err != nil {
					return err
				}
			default:
				err = errors.New("unrecognized struct field type")
				return err
			}
		}
	case reflect.Slice:
		var rem = in
		// switch the kind of sub type...
		switch rValue.Type().Elem().Kind() {
		case reflect.String:
			// traverse slice elements(string)
			var e error
			for {
				r, sub, e := extractOneString(rem)
				if e != nil {
					break
				}
				rValue = reflect.Append(rValue, reflect.ValueOf(r))
				rem = sub
			}
			// parse error leave
			if len(rem) != 0 {
				err = e
				return err
			}
			reflect.ValueOf(out).Elem().Set(rValue)
		case reflect.Int:
			// traverse slice elements(int)
			var e error
			for {
				r, sub, e := extractOneInt(rem)
				if e != nil {
					break
				}
				rValue = reflect.Append(rValue, reflect.ValueOf(r))
				rem = sub
			}
			// parse error leave
			if len(rem) != 0 {
				err = e
				return err
			}
			reflect.ValueOf(out).Elem().Set(rValue)
		case reflect.Float64:
			// traverse slice elements(float64)
			var e error
			for {
				r, sub, e := extractOneFloat64(rem)
				if e != nil {
					break
				}
				rValue = reflect.Append(rValue, reflect.ValueOf(r))
				rem = sub
			}
			// parse error leave
			if len(rem) != 0 {
				err = e
				return err
			}
			reflect.ValueOf(out).Elem().Set(rValue)
		case reflect.Bool:
			// traverse slice elements(bool)
			var e error
			for {
				r, sub, e := extractOneBool(rem)
				if e != nil {
					break
				}
				rValue = reflect.Append(rValue, reflect.ValueOf(r))
				rem = sub
			}
			// parse error leave
			if len(rem) != 0 {
				err = e
				return err
			}
			reflect.ValueOf(out).Elem().Set(rValue)
		case reflect.Struct:
			// traverse slice elements(struct)
			var e error
			for {
				r, sub, e := extractOneTuple(rem)
				if e != nil {
					break
				}
				rem = sub
				r, err = trimTuple(r)
				if err != nil {
					return err
				}
				r = repairTrim(r)
				// new struct value
				o := reflect.New(rValue.Type().Elem())
				err = decodeOneParameter(r, o.Interface())
				if err != nil {
					return err
				}
				rValue = reflect.Append(rValue, reflect.ValueOf(o.Elem().Interface()))
			}
			// parse error leave
			if len(rem) != 0 {
				err = e
				return err
			}
			reflect.ValueOf(out).Elem().Set(rValue)
		case reflect.Interface:
			// traverse slice elements(interface)
			var e error
			for {
				r, sub, e := extractOneTuple(rem)
				if e != nil {
					break
				}
				rem = sub
				r, err = trimTuple(r)
				if err != nil {
					return err
				}
				r = repairTrim(r)
				// check type by name of interface(extract type name)
				name, _, err := extractOneString(r)
				if err != nil {
					return err
				}
				// traverse map elements
				for k, v := range MParsesErl {
					if k == string(name) {
						// new struct value
						o := reflect.New(reflect.TypeOf(v))
						err = decodeOneParameter(r, o.Interface())
						if err != nil {
							return err
						}
						rValue = reflect.Append(rValue, reflect.ValueOf(o.Elem().Interface()))
						break
					}
				}
			}
			// parse error leave
			if len(rem) != 0 {
				err = e
				return err
			}
			reflect.ValueOf(out).Elem().Set(rValue)
		default:
			err = errors.New("unrecognized list element type")
			return err
		}
	/*case reflect.Interface:
	var rem = in
	// first, extract one parameter as string...
	r, sub, err := extractOneString(rem)
	if err != nil {
		return err
	}
	rem = sub*/
	default:
		err = errors.New("unrecognized reflect type")
		return err
	}
	return err
}

func decode(in []byte, out interface{}) (err error) {
	// get pointer's value...
	var rType = reflect.TypeOf(out)
	var rValue = reflect.ValueOf(out)
	// check the out type kind
	if rType.Kind() != reflect.Ptr {
		err = errors.New("out interface should be struct pointer")
		return err
	}
	// get real variable value...
	rType = rType.Elem()
	rValue = rValue.Elem()
	// type should be struct
	if rType.Kind() != reflect.Struct {
		err = errors.New("real variable should be struct")
		return err
	}
	// extract type name
	var r = in
	name, _, err := extractOneString(r)
	if err != nil {
		return err
	}
	// traverse struct fields
	for i := 0; i < rType.NumField(); i++ {
		// get struct field value...
		//t := rType.Field(i)
		f := rValue.Field(i)
		// check list elements type
		for k, v := range MParsesErl {
			if k == string(name) && reflect.TypeOf(v).Name() == f.Type().Elem().Name() {
				// new struct value
				o := reflect.New(reflect.TypeOf(v))
				err = decodeOneParameter(r, o.Interface())
				if err != nil {
					return err
				}
				f = reflect.Append(f, reflect.ValueOf(o.Elem().Interface()))
				break
			}
		}
		reflect.ValueOf(out).Elem().Field(i).Set(f)
	}
	return err
}

func unmarshal(in []byte, out interface{}) (err error) {
	var s [][]byte
	// split the stream by symbol '\n' in order to delete comments
	s1 := bytes.Split(in, []byte("\n"))
	for _, v := range s1 {
		// delete comments
		if bytes.Contains(v, []byte("%")) {
			index := bytes.Index(v, []byte("%"))
			v = append(v[:index], v[len(v):]...)
		}
		// delete C/C++ comments?
		if bytes.Contains(v, []byte("//")) {
			index := bytes.Index(v, []byte("//"))
			v = append(v[:index], v[len(v):]...)
		}
		// delete all space
		v = bytes.ReplaceAll(v, []byte(" "), []byte(""))
		// delete all return
		v = bytes.ReplaceAll(v, []byte("\r"), []byte(""))
		// delete all blank lines, packet to s slice
		if !bytes.Equal(v, []byte("")) {
			s = append(s, v)
		}
	}
	data := bytes.Join(s, []byte(""))
	// split the data by symbol '{' and '}.', according to syntax
	s = bytes.Split(data, []byte("}."))
	for _, v := range s {
		// valid line
		if !bytes.Contains(v, []byte("{")) {
			continue
		}
		// delete '{'
		index := bytes.Index(v, []byte("{"))
		v = append(v[:index], v[index+1:]...)
		// repair with ','
		v = repairTrim(v)
		// decode
		err = decode(v, out)
		if err != nil {
			return err
		}
	}
	return err
}

func unmarshalFrom(in string, out interface{}) (err error) {
	// try to open file...
	file, err := os.Open(in)
	if err != nil {
		return err
	}
	defer file.Close()
	// read the file...
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	// unmarshal
	err = unmarshal(data, out)
	if err != nil {
		return err
	}
	return err
}

// encode
func trimElement(s []byte) (r []byte) {
	r = bytes.TrimSuffix(s, []byte(","))
	return r
}

func repairList(s []byte) (r []byte) {
	ss := make([][]byte, 3)
	ss[0] = []byte("[")
	ss[1] = s
	ss[2] = []byte("]")
	r = bytes.Join(ss, []byte(""))
	return r
}

func repairTuple(s []byte) (r []byte) {
	ss := make([][]byte, 3)
	ss[0] = []byte("{")
	ss[1] = s
	ss[2] = []byte("}")
	r = bytes.Join(ss, []byte(""))
	return r
}

func encodeOneParameter(in interface{}) (r []byte, err error) {
	var rType = reflect.TypeOf(in)
	var rValue = reflect.ValueOf(in)
	// switch the s type kind
	switch rType.Kind() {
	case reflect.String:
		r = []byte(rValue.Interface().(string))
		r = append(r, ',')
	case reflect.Bool:
		f := rValue.Interface().(bool)
		if f == true {
			r = []byte("true")
		} else {
			r = []byte("false")
		}
		r = append(r, ',')
	case reflect.Int:
		r = []byte(strconv.Itoa(rValue.Interface().(int)))
		r = append(r, ',')
	case reflect.Float64:
		r = []byte(strconv.FormatFloat(rValue.Interface().(float64), 'f', -1, 64))
		r = append(r, ',')
	case reflect.Struct:
		var s [][]byte
		// traverse struct fields
		for i := 0; i < rType.NumField(); i++ {
			// get struct field value...
			t := rType.Field(i)
			f := rValue.Field(i)
			// parse tag
			tag := t.Tag.Get("erl")
			fields := strings.Split(tag, ",")
			if len(fields) > 1 {
				tag = fields[0]
			}
			// swich the tag & parse element...
			switch tag {
			case "string":
				fallthrough
			case "int":
				fallthrough
			case "float64":
				fallthrough
			case "bool":
				fallthrough
			case "tuple":
				fallthrough
			case "list":
				rs, err := encodeOneParameter(f.Interface())
				if err != nil {
					return r, err
				}
				if i == rType.NumField()-1 {
					rs = trimElement(rs)
				}
				s = append(s, rs)
			default:
				err = errors.New("unrecognized struct field type")
				return r, err
			}
		}
		// repair tuple...
		r = bytes.Join(s, []byte(""))
		r = repairTuple(r)
		r = repairTrim(r)
	case reflect.Slice:
		// traverse slice elements
		var s [][]byte
		for i := 0; i < rValue.Len(); i++ {
			rs, err := encodeOneParameter(rValue.Index(i).Interface())
			if err != nil {
				return r, err
			}
			if i == rValue.Len()-1 {
				rs = trimElement(rs)
			}
			s = append(s, rs)
		}
		// repair list...
		r = bytes.Join(s, []byte(""))
		r = repairList(r)
		r = repairTrim(r)
	default:
		err = errors.New("unrecognized element type")
		return r, err
	}
	return r, err
}

func encode(in interface{}) (out []byte, err error) {
	var rType = reflect.TypeOf(in)
	var rValue = reflect.ValueOf(in)
	// check the in type kind
	if rType.Kind() != reflect.Struct {
		err = errors.New("in interface should be struct")
		return nil, err
	}
	// traverse struct fields
	var s [][]byte
	for i := 0; i < rType.NumField(); i++ {
		// get struct field value...
		t := rType.Field(i)
		f := rValue.Field(i)
		// parse tag
		tag := t.Tag.Get("erl")
		fields := strings.Split(tag, ",")
		if len(fields) > 1 {
			tag = fields[0]
		}
		// check tag type(should be list only)
		if tag != "list" {
			err = errors.New("struct field tag should be list only")
			return out, err
		}
		// check list elements whether struct
		if f.Type().Elem().Kind() != reflect.Struct {
			err = errors.New("list element should be struct only")
			return out, err
		}
		// traverse list elements
		for j := 0; j < f.Len(); j++ {
			r, err := encodeOneParameter(f.Index(j).Interface())
			if err != nil {
				return out, err
			}
			// repair tuple...
			r = bytes.TrimSuffix(r, []byte(","))
			r = append(r, '.')
			r = append(r, '\n')
			s = append(s, r)
		}
	}
	out = bytes.Join(s, []byte(""))
	return out, err
}

func needClear(name []byte, in interface{}) (b bool, err error) {
	var rType = reflect.TypeOf(in)
	var rValue = reflect.ValueOf(in)
	// check the in type kind
	if rType.Kind() != reflect.Struct {
		err = errors.New("in interface should be struct")
		return b, err
	}
	// traverse struct fields
	for i := 0; i < rType.NumField(); i++ {
		// get struct field value...
		t := rType.Field(i)
		f := rValue.Field(i)
		// parse tag
		tag := t.Tag.Get("erl")
		fields := strings.Split(tag, ",")
		if len(fields) > 1 {
			tag = fields[0]
		}
		// check field whether empty?
		if f.IsNil() {
			return b, nil
		}
		// check tag type(should be list only)
		if tag != "list" {
			err = errors.New("struct field tag should be list only")
			return b, err
		}
		// check list elements whether struct
		if f.Type().Elem().Kind() != reflect.Struct {
			err = errors.New("list element should be struct only")
			return b, err
		}
		// check list elements type
		for k, v := range MParsesErl {
			if k == string(name) && reflect.TypeOf(v).Name() == f.Type().Elem().Name() {
				b = true
				return b, err
			}
		}
	}
	return b, err
}

func marshal(in []byte, t interface{}) (out []byte, err error) {
	var s [][]byte
	// split the stream by symbol '\n' in order to delete comments
	s1 := bytes.Split(in, []byte("\n"))
	for _, v := range s1 {
		// delete comments
		if bytes.Contains(v, []byte("%")) {
			index := bytes.Index(v, []byte("%"))
			v = append(v[:index], v[len(v):]...)
		}
		// delete C/C++ comments?
		if bytes.Contains(v, []byte("//")) {
			index := bytes.Index(v, []byte("//"))
			v = append(v[:index], v[len(v):]...)
		}
		// delete all space
		v = bytes.ReplaceAll(v, []byte(" "), []byte(""))
		// delete all return
		v = bytes.ReplaceAll(v, []byte("\r"), []byte(""))
		// delete all blank lines, packet to s slice
		if !bytes.Equal(v, []byte("")) {
			s = append(s, v)
		}
	}
	data := bytes.Join(s, []byte(""))
	// split the data by symbol '{' and '}.', according to syntax
	s = bytes.Split(data, []byte("}."))
	for i := 0; i < len(s); i++ {
		// valid line
		if !bytes.Contains(s[i], []byte("{")) {
			s = append(s[:i], s[i+1:]...)
			i--
			continue
		}
		// append all '}.'
		s[i] = append(s[i], '}')
		s[i] = append(s[i], '.')
		// find first syntax
		start := bytes.Index(s[i], []byte("{"))
		index := bytes.Index(s[i], []byte(","))
		name := append(s[i][start+1:index], s[i][len(s[i]):]...)
		// check parameter whether need clear?
		b, err := needClear(name, t)
		if err != nil {
			return out, err
		}
		if b {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	// encode parameters
	r1, err := encode(t)
	if err != nil {
		return out, err
	}
	r := bytes.Split(r1, []byte("\n"))
	for _, v := range r {
		// valid line
		if !bytes.Contains(v, []byte("{")) {
			continue
		}
		s = append(s, v)
	}
	// combine
	out = bytes.Join(s, []byte("\n"))
	return out, err
}

func marshalTo(in string, t interface{}) (err error) {
	// try to open file...
	file, err := os.Open(in)
	if err != nil {
		return err
	}
	defer file.Close()
	// read the file...
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	// marshal
	out, err := marshal(data, t)
	if err != nil {
		return err
	}
	// write to the file...
	err = ioutil.WriteFile(in, out, 0644)
	if err != nil {
		return err
	}
	return err
}
