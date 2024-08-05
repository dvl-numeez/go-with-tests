package reflection

import "reflect"


func walk(x interface{},fn func(input string)){
	val := getValue(x)
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice,reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _,key:= range val.MapKeys(){
			walk(val.MapIndex(key).Interface(),fn)
		}
	case reflect.Chan:
		for{
			if v,ok:=val.Recv();ok{
				walkValue(v)
			}else{
				break
			}
		}
	case reflect.Func:
		result:=val.Call(nil)
		for _,res:=range result{
			walkValue(res)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{})reflect.Value{
	val:=reflect.ValueOf(x)
	if val.Kind()==reflect.Pointer{
		val = val.Elem()
	}
	return val

}