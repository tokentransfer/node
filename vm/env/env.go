package env

import (
	"context"
	"time"

	"github.com/caivega/glog"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

func LoadModule(ctx context.Context, runtime wazero.Runtime) error {
	_, err := runtime.NewHostModuleBuilder("env").
		NewFunctionBuilder().WithFunc(log).Export("log").
		NewFunctionBuilder().WithFunc(runtimeTicks).Export("runtime.ticks").
		NewFunctionBuilder().WithFunc(runtimeSleepTicks).Export("runtime.sleepTicks").
		NewFunctionBuilder().WithFunc(stringVal).Export("syscall/js.stringVal").
		NewFunctionBuilder().WithFunc(valueGet).Export("syscall/js.valueGet").
		NewFunctionBuilder().WithFunc(valueSet).Export("syscall/js.valueSet").
		NewFunctionBuilder().WithFunc(valueIndex).Export("syscall/js.valueIndex").
		NewFunctionBuilder().WithFunc(valueSetIndex).Export("syscall/js.valueSetIndex").
		NewFunctionBuilder().WithFunc(valueCall).Export("syscall/js.valueCall").
		NewFunctionBuilder().WithFunc(valueNew).Export("syscall/js.valueNew").
		NewFunctionBuilder().WithFunc(valueLength).Export("syscall/js.valueLength").
		NewFunctionBuilder().WithFunc(valuePrepareString).Export("syscall/js.valuePrepareString").
		NewFunctionBuilder().WithFunc(valueLoadString).Export("syscall/js.valueLoadString").
		NewFunctionBuilder().WithFunc(finalizeRef).Export("syscall/js.finalizeRef").
		Instantiate(ctx)
	if err != nil {
		return err
	}
	return nil
}

//export runtimeTicks
func runtimeTicks(_ context.Context) float64 {
	t := float64(time.Now().UnixNano() / int64(time.Millisecond/time.Nanosecond))
	return t
}

//export runtimeSleepTicks
func runtimeSleepTicks(_ context.Context, a float64) {
	time.Sleep(time.Duration(a) * time.Millisecond)
}

//export stringVal
func stringVal(_ context.Context, sp1, sp2, sp3, sp4 int32) {
	// str := b.loadString(sp2, sp3)
	// b.storeValue(sp1, str)
}

//export valueGet
func valueGet(_ context.Context, sp1, sp2, sp3, sp4, sp5 int32) {
	// str := b.loadString(sp3, sp4)
	// id, val := b.loadValue(sp2)
	// obj, ok := val.(*object)
	// if !ok {
	// 	log.Fatalln("Object conversion error", str, id, val)
	// }
	// res, ok := obj.props[str]
	// if !ok {
	// 	log.Fatalln("Missing object property", val, str)
	// }
	// b.storeValue(sp1, res)
}

//export valueSet
func valueSet(_ context.Context, sp1, sp2, sp3, sp4, sp5 int32) {
	// _, obj := b.loadValue(sp1)
	// str1 := b.loadString(sp2, sp3)
	// _, str2 := b.loadValue(sp4)
	// obj.(*object).props[str1] = str2
}

func valueIndex(_ context.Context, sp1, sp2, sp3, sp4 int32) {
}

func valueSetIndex(_ context.Context, sp1, sp2, sp3, sp4, sp5 int32) {
}

//export valueCall
func valueCall(_ context.Context, sp1, sp2, sp3, sp4, sp5, sp6, sp7, sp8 int32) {
	// _, val := b.loadValue(sp2)
	// str := b.loadString(sp3, sp4)
	// args := b.loadSlice(sp5, sp6)
	// if reflect.TypeOf(val) == reflect.TypeOf(&object{}) {
	// 	obj := val.(*object)
	// 	var pfunc func(*stypedArray) *stypedArray
	// 	pfunc = obj.props[str].(func(*stypedArray) *stypedArray)
	// 	res := pfunc(args[0].(*stypedArray))
	// 	b.storeValue(sp1, res)
	// 	b.setUint8(sp1+8, 1)
	// } else if reflect.TypeOf(val) == reflect.TypeOf(&stypedArray{}) {
	// 	obj := val.(*stypedArray)
	// 	res := obj.toString()
	// 	b.storeValue(sp1, res)
	// 	b.setUint8(sp1+8, 1)
	// }
}

//export valueNew
func valueNew(_ context.Context, sp1, sp2, sp3, sp4, sp5, sp6, sp7 int32) {
	// id, val := b.loadValue(sp2)
	// args := b.loadSlice(sp3, sp4)
	// obj, ok := val.(*object)
	// if !ok {
	// 	log.Fatal("val is not an object", val)
	// }
	// ret := obj.new(args)
	// b.storeValue(sp1, ret)
	// b.setUint8(sp1+8, 1)
}

func valueLength(_ context.Context, sp1, sp2 int32) int32 {
	return 0
}

//export valuePrepareString
func valuePrepareString(_ context.Context, sp1, sp2, sp3 int32) {
	// _, v := b.loadValue(sp2)
	// b.storeString(sp1, v.(string))
}

//export valueLoadString
func valueLoadString(_ context.Context, sp1, sp2, sp3, sp4, sp5 int32) {
	// _, str := b.loadValue(sp1)
	// b.storeBytes(sp3, sp4, *str.(*[]byte))
}

//export finalizeRef
func finalizeRef(_ context.Context, sp1, sp2 int32) {

}

func log(_ context.Context, m api.Module, offset, byteCount uint32) {
	buf, ok := m.Memory().Read(offset, byteCount)
	if !ok {
		glog.Errorf("vm Memory.Read(%d, %d) out of range", offset, byteCount)
	} else {
		glog.Infoln("vm:" + string(buf))
	}
}
