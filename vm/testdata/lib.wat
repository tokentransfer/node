(module
  (type (;0;) (func (param i32 i32) (result i32)))
  (type (;1;) (func (param f32 f32) (result f32)))
  (func (;0;) (type 0) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.add)
  (func (;1;) (type 1) (param f32 f32) (result f32)
    local.get 0
    local.get 1
    f32.add)
  (memory (;0;) 17)
  (export "memory" (memory 0))
  (export "add" (func 0))
  (export "sum" (func 1)))
