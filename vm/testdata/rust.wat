(module
  (type (;0;) (func (param i32 i32)))
  (type (;1;) (func (param i32 i32 i32) (result i32)))
  (type (;2;) (func (param i32 i32) (result i32)))
  (type (;3;) (func (param i32 i32 i32)))
  (type (;4;) (func (param i64 i32 i32) (result i64)))
  (type (;5;) (func (param i32 i64 i32) (result i32)))
  (type (;6;) (func (param i32 i32 i64 i32)))
  (type (;7;) (func (param f32 f64) (result f64)))
  (type (;8;) (func (param i32 f32 i32 f64)))
  (type (;9;) (func (param i32 i32 i32 i32) (result i32)))
  (type (;10;) (func (param i32) (result i64)))
  (type (;11;) (func (param i32)))
  (type (;12;) (func (param i32 i32 i32 i32)))
  (type (;13;) (func (param i32) (result i32)))
  (type (;14;) (func (result i32)))
  (type (;15;) (func (param i32 i32 i32 i32 i32)))
  (type (;16;) (func (param i32 i32 i32 i32 i32) (result i32)))
  (type (;17;) (func))
  (type (;18;) (func (param i32 i32 i32 i32 i32 i32 i64 i64 i64)))
  (type (;19;) (func (param i32 i32 i32 i32 i32 i32 i32)))
  (type (;20;) (func (param i32 i32 i32 i32 i32 i32) (result i32)))
  (type (;21;) (func (param i32 f32 i32 i32) (result i32)))
  (type (;22;) (func (param i32 f64 i32 i32) (result i32)))
  (type (;23;) (func (param i32 i32 i32 i32 i32 i32 i32) (result i32)))
  (type (;24;) (func (param i64 i32 i32) (result i32)))
  (func $_ZN60_$LT$alloc..string..String$u20$as$u20$core..fmt..Display$GT$3fmt17h14fb461a74ffd756E (type 2) (param i32 i32) (result i32)
    local.get 0
    i32.load offset=4
    local.get 0
    i32.load offset=8
    local.get 1
    call $_ZN42_$LT$str$u20$as$u20$core..fmt..Display$GT$3fmt17h100f1ae78e96c78dE)
  (func $add (type 2) (param i32 i32) (result i32)
    local.get 1
    local.get 0
    i32.add)
  (func $test_i32_string (type 2) (param i32 i32) (result i32)
    (local i32 i32)
    local.get 1
    i32.load offset=8
    local.set 2
    block  ;; label = @1
      local.get 1
      i32.load
      local.tee 3
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      i32.const 4
      i32.add
      i32.load
      local.get 3
      i32.const 1
      call $__rust_dealloc
    end
    local.get 2
    local.get 0
    i32.add)
  (func $test_string_i32 (type 2) (param i32 i32) (result i32)
    (local i32 i32)
    local.get 0
    i32.load offset=8
    local.set 2
    block  ;; label = @1
      local.get 0
      i32.load
      local.tee 3
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.const 4
      i32.add
      i32.load
      local.get 3
      i32.const 1
      call $__rust_dealloc
    end
    local.get 2
    local.get 1
    i32.add)
  (func $test_string_i32_ret_string (type 3) (param i32 i32 i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    local.get 3
    local.get 2
    i32.store offset=4
    local.get 3
    i32.const 8
    i32.add
    i32.const 12
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 28
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 32
    i32.add
    i32.const 12
    i32.add
    i32.const 1
    i32.store
    local.get 3
    i32.const 1048580
    i32.store offset=16
    local.get 3
    i32.const 0
    i32.store offset=8
    local.get 3
    i32.const 2
    i32.store offset=36
    local.get 3
    local.get 1
    i32.store offset=32
    local.get 3
    local.get 3
    i32.const 32
    i32.add
    i32.store offset=24
    local.get 3
    local.get 3
    i32.const 4
    i32.add
    i32.store offset=40
    local.get 0
    local.get 3
    i32.const 8
    i32.add
    call $_ZN5alloc3fmt6format12format_inner17hf9dd33246b0be2a8E
    block  ;; label = @1
      local.get 1
      i32.load
      local.tee 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      i32.const 4
      i32.add
      i32.load
      local.get 2
      i32.const 1
      call $__rust_dealloc
    end
    local.get 3
    i32.const 48
    i32.add
    global.set $__stack_pointer)
  (func $test_i64_string_i32 (type 4) (param i64 i32 i32) (result i64)
    (local i64 i32)
    local.get 1
    i64.load32_u offset=8
    local.set 3
    block  ;; label = @1
      local.get 1
      i32.load
      local.tee 4
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      i32.const 4
      i32.add
      i32.load
      local.get 4
      i32.const 1
      call $__rust_dealloc
    end
    local.get 2
    i64.extend_i32_s
    local.get 0
    i64.add
    local.get 3
    i64.add)
  (func $test_string_i64_i32 (type 5) (param i32 i64 i32) (result i32)
    (local i32 i32)
    local.get 0
    i32.load offset=8
    local.set 3
    block  ;; label = @1
      local.get 0
      i32.load
      local.tee 4
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.const 4
      i32.add
      i32.load
      local.get 4
      i32.const 1
      call $__rust_dealloc
    end
    local.get 1
    i32.wrap_i64
    local.get 2
    i32.add
    local.get 3
    i32.add)
  (func $test_bytes_i64_i32_ret_string (type 6) (param i32 i32 i64 i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 80
    i32.sub
    local.tee 4
    global.set $__stack_pointer
    local.get 4
    local.get 3
    i32.store offset=20
    local.get 4
    local.get 2
    i64.store offset=8
    local.get 4
    i32.const 24
    i32.add
    i32.const 12
    i32.add
    i32.const 3
    i32.store
    local.get 4
    i32.const 24
    i32.add
    i32.const 20
    i32.add
    i32.const 3
    i32.store
    local.get 4
    i32.const 48
    i32.add
    i32.const 20
    i32.add
    i32.const 1
    i32.store
    local.get 4
    i32.const 48
    i32.add
    i32.const 12
    i32.add
    i32.const 3
    i32.store
    local.get 4
    i32.const 1048596
    i32.store offset=32
    local.get 4
    i32.const 0
    i32.store offset=24
    local.get 4
    i32.const 4
    i32.store offset=52
    local.get 4
    local.get 1
    i32.load offset=8
    i32.store offset=76
    local.get 4
    local.get 4
    i32.const 48
    i32.add
    i32.store offset=40
    local.get 4
    local.get 4
    i32.const 20
    i32.add
    i32.store offset=64
    local.get 4
    local.get 4
    i32.const 8
    i32.add
    i32.store offset=56
    local.get 4
    local.get 4
    i32.const 76
    i32.add
    i32.store offset=48
    local.get 0
    local.get 4
    i32.const 24
    i32.add
    call $_ZN5alloc3fmt6format12format_inner17hf9dd33246b0be2a8E
    block  ;; label = @1
      local.get 1
      i32.load
      local.tee 3
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      i32.const 4
      i32.add
      i32.load
      local.get 3
      i32.const 1
      call $__rust_dealloc
    end
    local.get 4
    i32.const 80
    i32.add
    global.set $__stack_pointer)
  (func $test_string_string_ret_string (type 3) (param i32 i32 i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    local.get 3
    i32.const 32
    i32.add
    i32.const 12
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 8
    i32.add
    i32.const 12
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 28
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 1048580
    i32.store offset=16
    local.get 3
    i32.const 0
    i32.store offset=8
    local.get 3
    local.get 2
    i32.store offset=40
    local.get 3
    i32.const 2
    i32.store offset=36
    local.get 3
    local.get 1
    i32.store offset=32
    local.get 3
    local.get 3
    i32.const 32
    i32.add
    i32.store offset=24
    local.get 0
    local.get 3
    i32.const 8
    i32.add
    call $_ZN5alloc3fmt6format12format_inner17hf9dd33246b0be2a8E
    block  ;; label = @1
      local.get 2
      i32.load
      local.tee 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 2
      i32.const 4
      i32.add
      i32.load
      local.get 0
      i32.const 1
      call $__rust_dealloc
    end
    block  ;; label = @1
      local.get 1
      i32.load
      local.tee 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      i32.const 4
      i32.add
      i32.load
      local.get 2
      i32.const 1
      call $__rust_dealloc
    end
    local.get 3
    i32.const 48
    i32.add
    global.set $__stack_pointer)
  (func $test_f32_f64_ret_f64 (type 7) (param f32 f64) (result f64)
    local.get 0
    f64.promote_f32
    local.get 1
    f64.add)
  (func $test_f32_string_f64_ret_string (type 8) (param i32 f32 i32 f64)
    (local i32)
    global.get $__stack_pointer
    i32.const 64
    i32.sub
    local.tee 4
    global.set $__stack_pointer
    local.get 4
    local.get 3
    f64.store offset=8
    local.get 4
    local.get 1
    f32.store offset=4
    local.get 4
    i32.const 16
    i32.add
    i32.const 12
    i32.add
    i32.const 3
    i32.store
    local.get 4
    i32.const 16
    i32.add
    i32.const 20
    i32.add
    i32.const 3
    i32.store
    local.get 4
    i32.const 40
    i32.add
    i32.const 20
    i32.add
    i32.const 5
    i32.store
    local.get 4
    i32.const 40
    i32.add
    i32.const 12
    i32.add
    i32.const 2
    i32.store
    local.get 4
    i32.const 1048596
    i32.store offset=24
    local.get 4
    i32.const 0
    i32.store offset=16
    local.get 4
    local.get 2
    i32.store offset=48
    local.get 4
    i32.const 6
    i32.store offset=44
    local.get 4
    local.get 4
    i32.const 40
    i32.add
    i32.store offset=32
    local.get 4
    local.get 4
    i32.const 8
    i32.add
    i32.store offset=56
    local.get 4
    local.get 4
    i32.const 4
    i32.add
    i32.store offset=40
    local.get 0
    local.get 4
    i32.const 16
    i32.add
    call $_ZN5alloc3fmt6format12format_inner17hf9dd33246b0be2a8E
    block  ;; label = @1
      local.get 2
      i32.load
      local.tee 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 2
      i32.const 4
      i32.add
      i32.load
      local.get 0
      i32.const 1
      call $__rust_dealloc
    end
    local.get 4
    i32.const 64
    i32.add
    global.set $__stack_pointer)
  (func $process (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call $_ZN60_$LT$alloc..string..String$u20$as$u20$core..clone..Clone$GT$5clone17h48fea206591b489eE
    block  ;; label = @1
      local.get 1
      i32.load
      local.tee 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      i32.const 4
      i32.add
      i32.load
      local.get 0
      i32.const 1
      call $__rust_dealloc
    end)
  (func $call_winner (type 2) (param i32 i32) (result i32)
    i32.const 0)
  (func $__rust_alloc (type 2) (param i32 i32) (result i32)
    (local i32)
    local.get 0
    local.get 1
    call $__rdl_alloc
    local.set 2
    local.get 2
    return)
  (func $__rust_dealloc (type 3) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call $__rdl_dealloc
    return)
  (func $__rust_realloc (type 9) (param i32 i32 i32 i32) (result i32)
    (local i32)
    local.get 0
    local.get 1
    local.get 2
    local.get 3
    call $__rdl_realloc
    local.set 4
    local.get 4
    return)
  (func $__rust_alloc_error_handler (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call $_ZN3std5alloc8rust_oom17h528475c4a036b714E
    return)
  (func $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17h38941f43dd32bc22E (type 10) (param i32) (result i64)
    i64.const 726813373705916399)
  (func $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17ha86bcc140bde4c2fE (type 10) (param i32) (result i64)
    i64.const -8527728395957036344)
  (func $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17hb75db078a9b00dbdE (type 10) (param i32) (result i64)
    i64.const 3382899207125954301)
  (func $_ZN5alloc7raw_vec19RawVec$LT$T$C$A$GT$7reserve21do_reserve_and_handle17h68fb4a3403bc3a9fE (type 3) (param i32 i32 i32)
    (local i32 i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        local.get 2
        i32.add
        local.tee 2
        local.get 1
        i32.lt_u
        br_if 0 (;@2;)
        local.get 0
        i32.load
        local.tee 1
        i32.const 1
        i32.shl
        local.tee 4
        local.get 2
        local.get 4
        local.get 2
        i32.gt_u
        select
        local.tee 2
        i32.const 8
        local.get 2
        i32.const 8
        i32.gt_u
        select
        local.tee 2
        i32.const -1
        i32.xor
        i32.const 31
        i32.shr_u
        local.set 4
        block  ;; label = @3
          block  ;; label = @4
            local.get 1
            i32.eqz
            br_if 0 (;@4;)
            local.get 3
            i32.const 1
            i32.store offset=24
            local.get 3
            local.get 1
            i32.store offset=20
            local.get 3
            local.get 0
            i32.const 4
            i32.add
            i32.load
            i32.store offset=16
            br 1 (;@3;)
          end
          local.get 3
          i32.const 0
          i32.store offset=24
        end
        local.get 3
        local.get 2
        local.get 4
        local.get 3
        i32.const 16
        i32.add
        call $_ZN5alloc7raw_vec11finish_grow17hc2a10121f1c44e8dE
        local.get 3
        i32.load offset=4
        local.set 1
        block  ;; label = @3
          local.get 3
          i32.load
          br_if 0 (;@3;)
          local.get 0
          local.get 2
          i32.store
          local.get 0
          local.get 1
          i32.store offset=4
          br 2 (;@1;)
        end
        local.get 3
        i32.const 8
        i32.add
        i32.load
        local.tee 0
        i32.const -2147483647
        i32.eq
        br_if 1 (;@1;)
        local.get 0
        i32.eqz
        br_if 0 (;@2;)
        local.get 1
        local.get 0
        call $_ZN5alloc5alloc18handle_alloc_error17hc1ba5222dd907d5cE
        unreachable
      end
      call $_ZN5alloc7raw_vec17capacity_overflow17h25579b3a3c90486eE
      unreachable
    end
    local.get 3
    i32.const 32
    i32.add
    global.set $__stack_pointer)
  (func $_ZN4core3ptr100drop_in_place$LT$$RF$mut$u20$std..io..Write..write_fmt..Adapter$LT$alloc..vec..Vec$LT$u8$GT$$GT$$GT$17h30f4c8c66d1c1575E (type 11) (param i32))
  (func $_ZN4core3ptr42drop_in_place$LT$alloc..string..String$GT$17hc02ce641e9e8d8a8E (type 11) (param i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      i32.load
      local.tee 1
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.const 4
      i32.add
      i32.load
      local.get 1
      i32.const 1
      call $__rust_dealloc
    end)
  (func $_ZN4core3ptr70drop_in_place$LT$core..option..Option$LT$alloc..string..String$GT$$GT$17h1c38fec998b374e7E (type 11) (param i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      i32.const 4
      i32.add
      i32.load
      local.tee 1
      i32.eqz
      br_if 0 (;@1;)
      local.get 0
      i32.load
      local.tee 0
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      local.get 0
      i32.const 1
      call $__rust_dealloc
    end)
  (func $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$10write_char17hd5818ebd4677eb1aE (type 2) (param i32 i32) (result i32)
    (local i32 i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    local.get 0
    i32.load
    local.set 0
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.const 127
        i32.gt_u
        br_if 0 (;@2;)
        block  ;; label = @3
          local.get 0
          i32.load offset=8
          local.tee 3
          local.get 0
          i32.load
          i32.ne
          br_if 0 (;@3;)
          local.get 0
          local.get 3
          call $_ZN5alloc7raw_vec19RawVec$LT$T$C$A$GT$16reserve_for_push17hcd46213a7c2e29a0E
          local.get 0
          i32.load offset=8
          local.set 3
        end
        local.get 0
        local.get 3
        i32.const 1
        i32.add
        i32.store offset=8
        local.get 0
        i32.load offset=4
        local.get 3
        i32.add
        local.get 1
        i32.store8
        br 1 (;@1;)
      end
      local.get 2
      i32.const 0
      i32.store offset=12
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.const 2048
          i32.lt_u
          br_if 0 (;@3;)
          block  ;; label = @4
            local.get 1
            i32.const 65536
            i32.lt_u
            br_if 0 (;@4;)
            local.get 2
            local.get 1
            i32.const 63
            i32.and
            i32.const 128
            i32.or
            i32.store8 offset=15
            local.get 2
            local.get 1
            i32.const 6
            i32.shr_u
            i32.const 63
            i32.and
            i32.const 128
            i32.or
            i32.store8 offset=14
            local.get 2
            local.get 1
            i32.const 12
            i32.shr_u
            i32.const 63
            i32.and
            i32.const 128
            i32.or
            i32.store8 offset=13
            local.get 2
            local.get 1
            i32.const 18
            i32.shr_u
            i32.const 7
            i32.and
            i32.const 240
            i32.or
            i32.store8 offset=12
            i32.const 4
            local.set 1
            br 2 (;@2;)
          end
          local.get 2
          local.get 1
          i32.const 63
          i32.and
          i32.const 128
          i32.or
          i32.store8 offset=14
          local.get 2
          local.get 1
          i32.const 12
          i32.shr_u
          i32.const 224
          i32.or
          i32.store8 offset=12
          local.get 2
          local.get 1
          i32.const 6
          i32.shr_u
          i32.const 63
          i32.and
          i32.const 128
          i32.or
          i32.store8 offset=13
          i32.const 3
          local.set 1
          br 1 (;@2;)
        end
        local.get 2
        local.get 1
        i32.const 63
        i32.and
        i32.const 128
        i32.or
        i32.store8 offset=13
        local.get 2
        local.get 1
        i32.const 6
        i32.shr_u
        i32.const 192
        i32.or
        i32.store8 offset=12
        i32.const 2
        local.set 1
      end
      block  ;; label = @2
        local.get 0
        i32.load
        local.get 0
        i32.load offset=8
        local.tee 3
        i32.sub
        local.get 1
        i32.ge_u
        br_if 0 (;@2;)
        local.get 0
        local.get 3
        local.get 1
        call $_ZN5alloc7raw_vec19RawVec$LT$T$C$A$GT$7reserve21do_reserve_and_handle17h68fb4a3403bc3a9fE
        local.get 0
        i32.load offset=8
        local.set 3
      end
      local.get 0
      i32.load offset=4
      local.get 3
      i32.add
      local.get 2
      i32.const 12
      i32.add
      local.get 1
      call $memcpy
      drop
      local.get 0
      local.get 3
      local.get 1
      i32.add
      i32.store offset=8
    end
    local.get 2
    i32.const 16
    i32.add
    global.set $__stack_pointer
    i32.const 0)
  (func $_ZN5alloc7raw_vec19RawVec$LT$T$C$A$GT$16reserve_for_push17hcd46213a7c2e29a0E (type 0) (param i32 i32)
    (local i32 i32 i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.const 1
        i32.add
        local.tee 1
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        i32.load
        local.tee 3
        i32.const 1
        i32.shl
        local.tee 4
        local.get 1
        local.get 4
        local.get 1
        i32.gt_u
        select
        local.tee 1
        i32.const 8
        local.get 1
        i32.const 8
        i32.gt_u
        select
        local.tee 1
        i32.const -1
        i32.xor
        i32.const 31
        i32.shr_u
        local.set 4
        block  ;; label = @3
          block  ;; label = @4
            local.get 3
            i32.eqz
            br_if 0 (;@4;)
            local.get 2
            i32.const 1
            i32.store offset=24
            local.get 2
            local.get 3
            i32.store offset=20
            local.get 2
            local.get 0
            i32.const 4
            i32.add
            i32.load
            i32.store offset=16
            br 1 (;@3;)
          end
          local.get 2
          i32.const 0
          i32.store offset=24
        end
        local.get 2
        local.get 1
        local.get 4
        local.get 2
        i32.const 16
        i32.add
        call $_ZN5alloc7raw_vec11finish_grow17hc2a10121f1c44e8dE
        local.get 2
        i32.load offset=4
        local.set 3
        block  ;; label = @3
          local.get 2
          i32.load
          br_if 0 (;@3;)
          local.get 0
          local.get 1
          i32.store
          local.get 0
          local.get 3
          i32.store offset=4
          br 2 (;@1;)
        end
        local.get 2
        i32.const 8
        i32.add
        i32.load
        local.tee 0
        i32.const -2147483647
        i32.eq
        br_if 1 (;@1;)
        local.get 0
        i32.eqz
        br_if 0 (;@2;)
        local.get 3
        local.get 0
        call $_ZN5alloc5alloc18handle_alloc_error17hc1ba5222dd907d5cE
        unreachable
      end
      call $_ZN5alloc7raw_vec17capacity_overflow17h25579b3a3c90486eE
      unreachable
    end
    local.get 2
    i32.const 32
    i32.add
    global.set $__stack_pointer)
  (func $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_fmt17h2452ec875f334425E (type 2) (param i32 i32) (result i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    local.get 0
    i32.load
    local.set 0
    local.get 2
    i32.const 8
    i32.add
    i32.const 16
    i32.add
    local.get 1
    i32.const 16
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    i32.const 8
    i32.add
    i32.const 8
    i32.add
    local.get 1
    i32.const 8
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    local.get 1
    i64.load align=4
    i64.store offset=8
    local.get 2
    local.get 0
    i32.store offset=4
    local.get 2
    i32.const 4
    i32.add
    i32.const 1048620
    local.get 2
    i32.const 8
    i32.add
    call $_ZN4core3fmt5write17h7d7e1e065b0161edE
    local.set 1
    local.get 2
    i32.const 32
    i32.add
    global.set $__stack_pointer
    local.get 1)
  (func $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_str17h7147b2416db1cf04E (type 1) (param i32 i32 i32) (result i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      i32.load
      local.tee 0
      i32.load
      local.get 0
      i32.load offset=8
      local.tee 3
      i32.sub
      local.get 2
      i32.ge_u
      br_if 0 (;@1;)
      local.get 0
      local.get 3
      local.get 2
      call $_ZN5alloc7raw_vec19RawVec$LT$T$C$A$GT$7reserve21do_reserve_and_handle17h68fb4a3403bc3a9fE
      local.get 0
      i32.load offset=8
      local.set 3
    end
    local.get 0
    i32.load offset=4
    local.get 3
    i32.add
    local.get 1
    local.get 2
    call $memcpy
    drop
    local.get 0
    local.get 3
    local.get 2
    i32.add
    i32.store offset=8
    i32.const 0)
  (func $_ZN5alloc7raw_vec11finish_grow17hc2a10121f1c44e8dE (type 12) (param i32 i32 i32 i32)
    (local i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.eqz
        br_if 0 (;@2;)
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 1
                  i32.const -1
                  i32.le_s
                  br_if 0 (;@7;)
                  local.get 3
                  i32.load offset=8
                  i32.eqz
                  br_if 2 (;@5;)
                  local.get 3
                  i32.load offset=4
                  local.tee 4
                  br_if 1 (;@6;)
                  local.get 1
                  br_if 3 (;@4;)
                  local.get 2
                  local.set 3
                  br 4 (;@3;)
                end
                local.get 0
                i32.const 8
                i32.add
                i32.const 0
                i32.store
                br 5 (;@1;)
              end
              local.get 3
              i32.load
              local.get 4
              local.get 2
              local.get 1
              call $__rust_realloc
              local.set 3
              br 2 (;@3;)
            end
            local.get 1
            br_if 0 (;@4;)
            local.get 2
            local.set 3
            br 1 (;@3;)
          end
          local.get 1
          local.get 2
          call $__rust_alloc
          local.set 3
        end
        block  ;; label = @3
          local.get 3
          i32.eqz
          br_if 0 (;@3;)
          local.get 0
          local.get 3
          i32.store offset=4
          local.get 0
          i32.const 8
          i32.add
          local.get 1
          i32.store
          local.get 0
          i32.const 0
          i32.store
          return
        end
        local.get 0
        local.get 1
        i32.store offset=4
        local.get 0
        i32.const 8
        i32.add
        local.get 2
        i32.store
        br 1 (;@1;)
      end
      local.get 0
      local.get 1
      i32.store offset=4
      local.get 0
      i32.const 8
      i32.add
      i32.const 0
      i32.store
    end
    local.get 0
    i32.const 1
    i32.store)
  (func $_ZN8dlmalloc17Dlmalloc$LT$A$GT$6malloc17h7a9e47086a367e16E (type 2) (param i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 1
              i32.const 9
              i32.lt_u
              br_if 0 (;@5;)
              i32.const 16
              i32.const 8
              call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
              local.get 1
              i32.gt_u
              br_if 1 (;@4;)
              br 2 (;@3;)
            end
            local.get 0
            call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$6malloc17hfc360c49b49c4ccfE
            local.set 2
            br 2 (;@2;)
          end
          i32.const 16
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          local.set 1
        end
        call $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE
        local.tee 3
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        local.set 4
        i32.const 20
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        local.set 5
        i32.const 16
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        local.set 6
        i32.const 0
        local.set 2
        i32.const 0
        i32.const 16
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        i32.const 2
        i32.shl
        i32.sub
        local.tee 7
        local.get 3
        local.get 6
        local.get 4
        local.get 5
        i32.add
        i32.add
        i32.sub
        i32.const -65544
        i32.add
        i32.const -9
        i32.and
        i32.const -3
        i32.add
        local.tee 3
        local.get 7
        local.get 3
        i32.lt_u
        select
        local.get 1
        i32.sub
        local.get 0
        i32.le_u
        br_if 0 (;@2;)
        local.get 1
        i32.const 16
        local.get 0
        i32.const 4
        i32.add
        i32.const 16
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        i32.const -5
        i32.add
        local.get 0
        i32.gt_u
        select
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        local.tee 4
        i32.add
        i32.const 16
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        i32.add
        i32.const -4
        i32.add
        call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$6malloc17hfc360c49b49c4ccfE
        local.tee 3
        i32.eqz
        br_if 0 (;@2;)
        local.get 3
        call $_ZN8dlmalloc8dlmalloc5Chunk8from_mem17hb5c408f151442e99E
        local.set 0
        block  ;; label = @3
          block  ;; label = @4
            local.get 1
            i32.const -1
            i32.add
            local.tee 2
            local.get 3
            i32.and
            br_if 0 (;@4;)
            local.get 0
            local.set 1
            br 1 (;@3;)
          end
          local.get 2
          local.get 3
          i32.add
          i32.const 0
          local.get 1
          i32.sub
          i32.and
          call $_ZN8dlmalloc8dlmalloc5Chunk8from_mem17hb5c408f151442e99E
          local.set 2
          i32.const 16
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          local.set 3
          local.get 0
          call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
          local.get 2
          i32.const 0
          local.get 1
          local.get 2
          local.get 0
          i32.sub
          local.get 3
          i32.gt_u
          select
          i32.add
          local.tee 1
          local.get 0
          i32.sub
          local.tee 2
          i32.sub
          local.set 3
          block  ;; label = @4
            local.get 0
            call $_ZN8dlmalloc8dlmalloc5Chunk7mmapped17h69bc6c14adea384fE
            br_if 0 (;@4;)
            local.get 1
            local.get 3
            call $_ZN8dlmalloc8dlmalloc5Chunk9set_inuse17h5f3bbfae424d3974E
            local.get 0
            local.get 2
            call $_ZN8dlmalloc8dlmalloc5Chunk9set_inuse17h5f3bbfae424d3974E
            local.get 0
            local.get 2
            call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$13dispose_chunk17he67418b997d9e729E
            br 1 (;@3;)
          end
          local.get 0
          i32.load
          local.set 0
          local.get 1
          local.get 3
          i32.store offset=4
          local.get 1
          local.get 0
          local.get 2
          i32.add
          i32.store
        end
        local.get 1
        call $_ZN8dlmalloc8dlmalloc5Chunk7mmapped17h69bc6c14adea384fE
        br_if 1 (;@1;)
        local.get 1
        call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
        local.tee 0
        i32.const 16
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        local.get 4
        i32.add
        i32.le_u
        br_if 1 (;@1;)
        local.get 1
        local.get 4
        call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
        local.set 2
        local.get 1
        local.get 4
        call $_ZN8dlmalloc8dlmalloc5Chunk9set_inuse17h5f3bbfae424d3974E
        local.get 2
        local.get 0
        local.get 4
        i32.sub
        local.tee 0
        call $_ZN8dlmalloc8dlmalloc5Chunk9set_inuse17h5f3bbfae424d3974E
        local.get 2
        local.get 0
        call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$13dispose_chunk17he67418b997d9e729E
        br 1 (;@1;)
      end
      local.get 2
      return
    end
    local.get 1
    call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
    local.set 0
    local.get 1
    call $_ZN8dlmalloc8dlmalloc5Chunk7mmapped17h69bc6c14adea384fE
    drop
    local.get 0)
  (func $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$6malloc17hfc360c49b49c4ccfE (type 13) (param i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i64)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 1
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 0
                  i32.const 245
                  i32.lt_u
                  br_if 0 (;@7;)
                  call $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE
                  local.tee 2
                  i32.const 8
                  call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                  local.set 3
                  i32.const 20
                  i32.const 8
                  call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                  local.set 4
                  i32.const 16
                  i32.const 8
                  call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                  local.set 5
                  i32.const 0
                  local.set 6
                  i32.const 0
                  i32.const 16
                  i32.const 8
                  call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                  i32.const 2
                  i32.shl
                  i32.sub
                  local.tee 7
                  local.get 2
                  local.get 5
                  local.get 3
                  local.get 4
                  i32.add
                  i32.add
                  i32.sub
                  i32.const -65544
                  i32.add
                  i32.const -9
                  i32.and
                  i32.const -3
                  i32.add
                  local.tee 2
                  local.get 7
                  local.get 2
                  i32.lt_u
                  select
                  local.get 0
                  i32.le_u
                  br_if 6 (;@1;)
                  local.get 0
                  i32.const 4
                  i32.add
                  i32.const 8
                  call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                  local.set 2
                  i32.const 0
                  i32.load offset=1056284
                  i32.eqz
                  br_if 5 (;@2;)
                  i32.const 0
                  local.set 8
                  block  ;; label = @8
                    local.get 2
                    i32.const 256
                    i32.lt_u
                    br_if 0 (;@8;)
                    i32.const 31
                    local.set 8
                    local.get 2
                    i32.const 16777215
                    i32.gt_u
                    br_if 0 (;@8;)
                    local.get 2
                    i32.const 6
                    local.get 2
                    i32.const 8
                    i32.shr_u
                    i32.clz
                    local.tee 0
                    i32.sub
                    i32.shr_u
                    i32.const 1
                    i32.and
                    local.get 0
                    i32.const 1
                    i32.shl
                    i32.sub
                    i32.const 62
                    i32.add
                    local.set 8
                  end
                  i32.const 0
                  local.get 2
                  i32.sub
                  local.set 6
                  local.get 8
                  i32.const 2
                  i32.shl
                  i32.const 1055872
                  i32.add
                  i32.load
                  local.tee 3
                  br_if 1 (;@6;)
                  i32.const 0
                  local.set 0
                  i32.const 0
                  local.set 4
                  br 2 (;@5;)
                end
                i32.const 16
                local.get 0
                i32.const 4
                i32.add
                i32.const 16
                i32.const 8
                call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                i32.const -5
                i32.add
                local.get 0
                i32.gt_u
                select
                i32.const 8
                call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                local.set 2
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            block  ;; label = @13
                              i32.const 0
                              i32.load offset=1056280
                              local.tee 4
                              local.get 2
                              i32.const 3
                              i32.shr_u
                              local.tee 6
                              i32.shr_u
                              local.tee 0
                              i32.const 3
                              i32.and
                              br_if 0 (;@13;)
                              local.get 2
                              i32.const 0
                              i32.load offset=1056288
                              i32.le_u
                              br_if 11 (;@2;)
                              local.get 0
                              br_if 1 (;@12;)
                              i32.const 0
                              i32.load offset=1056284
                              local.tee 0
                              i32.eqz
                              br_if 11 (;@2;)
                              local.get 0
                              call $_ZN8dlmalloc8dlmalloc9least_bit17h8a739bad8428dbdfE
                              i32.ctz
                              i32.const 2
                              i32.shl
                              i32.const 1055872
                              i32.add
                              i32.load
                              local.tee 3
                              call $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE
                              call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
                              local.get 2
                              i32.sub
                              local.set 6
                              block  ;; label = @14
                                local.get 3
                                call $_ZN8dlmalloc8dlmalloc9TreeChunk14leftmost_child17he1deea80b772637aE
                                local.tee 0
                                i32.eqz
                                br_if 0 (;@14;)
                                loop  ;; label = @15
                                  local.get 0
                                  call $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE
                                  call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
                                  local.get 2
                                  i32.sub
                                  local.tee 4
                                  local.get 6
                                  local.get 4
                                  local.get 6
                                  i32.lt_u
                                  local.tee 4
                                  select
                                  local.set 6
                                  local.get 0
                                  local.get 3
                                  local.get 4
                                  select
                                  local.set 3
                                  local.get 0
                                  call $_ZN8dlmalloc8dlmalloc9TreeChunk14leftmost_child17he1deea80b772637aE
                                  local.tee 0
                                  br_if 0 (;@15;)
                                end
                              end
                              local.get 3
                              call $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE
                              local.tee 0
                              local.get 2
                              call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                              local.set 4
                              local.get 3
                              call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18unlink_large_chunk17h66e0805d86ff34c0E
                              local.get 6
                              i32.const 16
                              i32.const 8
                              call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                              i32.lt_u
                              br_if 5 (;@8;)
                              local.get 4
                              call $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE
                              local.set 4
                              local.get 0
                              local.get 2
                              call $_ZN8dlmalloc8dlmalloc5Chunk34set_size_and_pinuse_of_inuse_chunk17h54212a19225d5a32E
                              local.get 4
                              local.get 6
                              call $_ZN8dlmalloc8dlmalloc5Chunk33set_size_and_pinuse_of_free_chunk17h6fb36837ea4ab16bE
                              i32.const 0
                              i32.load offset=1056288
                              local.tee 7
                              i32.eqz
                              br_if 4 (;@9;)
                              local.get 7
                              i32.const -8
                              i32.and
                              i32.const 1056016
                              i32.add
                              local.set 5
                              i32.const 0
                              i32.load offset=1056296
                              local.set 3
                              i32.const 0
                              i32.load offset=1056280
                              local.tee 8
                              i32.const 1
                              local.get 7
                              i32.const 3
                              i32.shr_u
                              i32.shl
                              local.tee 7
                              i32.and
                              i32.eqz
                              br_if 2 (;@11;)
                              local.get 5
                              i32.load offset=8
                              local.set 7
                              br 3 (;@10;)
                            end
                            block  ;; label = @13
                              block  ;; label = @14
                                local.get 0
                                i32.const -1
                                i32.xor
                                i32.const 1
                                i32.and
                                local.get 6
                                i32.add
                                local.tee 2
                                i32.const 3
                                i32.shl
                                local.tee 3
                                i32.const 1056024
                                i32.add
                                i32.load
                                local.tee 0
                                i32.const 8
                                i32.add
                                i32.load
                                local.tee 6
                                local.get 3
                                i32.const 1056016
                                i32.add
                                local.tee 3
                                i32.eq
                                br_if 0 (;@14;)
                                local.get 6
                                local.get 3
                                i32.store offset=12
                                local.get 3
                                local.get 6
                                i32.store offset=8
                                br 1 (;@13;)
                              end
                              i32.const 0
                              local.get 4
                              i32.const -2
                              local.get 2
                              i32.rotl
                              i32.and
                              i32.store offset=1056280
                            end
                            local.get 0
                            local.get 2
                            i32.const 3
                            i32.shl
                            call $_ZN8dlmalloc8dlmalloc5Chunk20set_inuse_and_pinuse17h134eb0365548205aE
                            local.get 0
                            call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
                            local.set 6
                            br 11 (;@1;)
                          end
                          block  ;; label = @12
                            block  ;; label = @13
                              i32.const 1
                              local.get 6
                              i32.const 31
                              i32.and
                              local.tee 6
                              i32.shl
                              call $_ZN8dlmalloc8dlmalloc9left_bits17hf9924c098a7140faE
                              local.get 0
                              local.get 6
                              i32.shl
                              i32.and
                              call $_ZN8dlmalloc8dlmalloc9least_bit17h8a739bad8428dbdfE
                              i32.ctz
                              local.tee 6
                              i32.const 3
                              i32.shl
                              local.tee 4
                              i32.const 1056024
                              i32.add
                              i32.load
                              local.tee 0
                              i32.const 8
                              i32.add
                              i32.load
                              local.tee 3
                              local.get 4
                              i32.const 1056016
                              i32.add
                              local.tee 4
                              i32.eq
                              br_if 0 (;@13;)
                              local.get 3
                              local.get 4
                              i32.store offset=12
                              local.get 4
                              local.get 3
                              i32.store offset=8
                              br 1 (;@12;)
                            end
                            i32.const 0
                            i32.const 0
                            i32.load offset=1056280
                            i32.const -2
                            local.get 6
                            i32.rotl
                            i32.and
                            i32.store offset=1056280
                          end
                          local.get 0
                          local.get 2
                          call $_ZN8dlmalloc8dlmalloc5Chunk34set_size_and_pinuse_of_inuse_chunk17h54212a19225d5a32E
                          local.get 0
                          local.get 2
                          call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                          local.tee 4
                          local.get 6
                          i32.const 3
                          i32.shl
                          local.get 2
                          i32.sub
                          local.tee 5
                          call $_ZN8dlmalloc8dlmalloc5Chunk33set_size_and_pinuse_of_free_chunk17h6fb36837ea4ab16bE
                          block  ;; label = @12
                            i32.const 0
                            i32.load offset=1056288
                            local.tee 3
                            i32.eqz
                            br_if 0 (;@12;)
                            local.get 3
                            i32.const -8
                            i32.and
                            i32.const 1056016
                            i32.add
                            local.set 6
                            i32.const 0
                            i32.load offset=1056296
                            local.set 2
                            block  ;; label = @13
                              block  ;; label = @14
                                i32.const 0
                                i32.load offset=1056280
                                local.tee 7
                                i32.const 1
                                local.get 3
                                i32.const 3
                                i32.shr_u
                                i32.shl
                                local.tee 3
                                i32.and
                                i32.eqz
                                br_if 0 (;@14;)
                                local.get 6
                                i32.load offset=8
                                local.set 3
                                br 1 (;@13;)
                              end
                              i32.const 0
                              local.get 7
                              local.get 3
                              i32.or
                              i32.store offset=1056280
                              local.get 6
                              local.set 3
                            end
                            local.get 6
                            local.get 2
                            i32.store offset=8
                            local.get 3
                            local.get 2
                            i32.store offset=12
                            local.get 2
                            local.get 6
                            i32.store offset=12
                            local.get 2
                            local.get 3
                            i32.store offset=8
                          end
                          i32.const 0
                          local.get 4
                          i32.store offset=1056296
                          i32.const 0
                          local.get 5
                          i32.store offset=1056288
                          local.get 0
                          call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
                          local.set 6
                          br 10 (;@1;)
                        end
                        i32.const 0
                        local.get 8
                        local.get 7
                        i32.or
                        i32.store offset=1056280
                        local.get 5
                        local.set 7
                      end
                      local.get 5
                      local.get 3
                      i32.store offset=8
                      local.get 7
                      local.get 3
                      i32.store offset=12
                      local.get 3
                      local.get 5
                      i32.store offset=12
                      local.get 3
                      local.get 7
                      i32.store offset=8
                    end
                    i32.const 0
                    local.get 4
                    i32.store offset=1056296
                    i32.const 0
                    local.get 6
                    i32.store offset=1056288
                    br 1 (;@7;)
                  end
                  local.get 0
                  local.get 6
                  local.get 2
                  i32.add
                  call $_ZN8dlmalloc8dlmalloc5Chunk20set_inuse_and_pinuse17h134eb0365548205aE
                end
                local.get 0
                call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
                local.tee 6
                br_if 5 (;@1;)
                br 4 (;@2;)
              end
              local.get 2
              local.get 8
              call $_ZN8dlmalloc8dlmalloc24leftshift_for_tree_index17hf578475f2b108550E
              i32.shl
              local.set 5
              i32.const 0
              local.set 0
              i32.const 0
              local.set 4
              loop  ;; label = @6
                block  ;; label = @7
                  local.get 3
                  call $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE
                  call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
                  local.tee 7
                  local.get 2
                  i32.lt_u
                  br_if 0 (;@7;)
                  local.get 7
                  local.get 2
                  i32.sub
                  local.tee 7
                  local.get 6
                  i32.ge_u
                  br_if 0 (;@7;)
                  local.get 7
                  local.set 6
                  local.get 3
                  local.set 4
                  local.get 7
                  br_if 0 (;@7;)
                  i32.const 0
                  local.set 6
                  local.get 3
                  local.set 4
                  local.get 3
                  local.set 0
                  br 3 (;@4;)
                end
                local.get 3
                i32.const 20
                i32.add
                i32.load
                local.tee 7
                local.get 0
                local.get 7
                local.get 3
                local.get 5
                i32.const 29
                i32.shr_u
                i32.const 4
                i32.and
                i32.add
                i32.const 16
                i32.add
                i32.load
                local.tee 3
                i32.ne
                select
                local.get 0
                local.get 7
                select
                local.set 0
                local.get 5
                i32.const 1
                i32.shl
                local.set 5
                local.get 3
                br_if 0 (;@6;)
              end
            end
            block  ;; label = @5
              local.get 0
              local.get 4
              i32.or
              br_if 0 (;@5;)
              i32.const 0
              local.set 4
              i32.const 1
              local.get 8
              i32.shl
              call $_ZN8dlmalloc8dlmalloc9left_bits17hf9924c098a7140faE
              i32.const 0
              i32.load offset=1056284
              i32.and
              local.tee 0
              i32.eqz
              br_if 3 (;@2;)
              local.get 0
              call $_ZN8dlmalloc8dlmalloc9least_bit17h8a739bad8428dbdfE
              i32.ctz
              i32.const 2
              i32.shl
              i32.const 1055872
              i32.add
              i32.load
              local.set 0
            end
            local.get 0
            i32.eqz
            br_if 1 (;@3;)
          end
          loop  ;; label = @4
            local.get 0
            local.get 4
            local.get 0
            call $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE
            call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
            local.tee 3
            local.get 2
            i32.ge_u
            local.get 3
            local.get 2
            i32.sub
            local.tee 3
            local.get 6
            i32.lt_u
            i32.and
            local.tee 5
            select
            local.set 4
            local.get 3
            local.get 6
            local.get 5
            select
            local.set 6
            local.get 0
            call $_ZN8dlmalloc8dlmalloc9TreeChunk14leftmost_child17he1deea80b772637aE
            local.tee 0
            br_if 0 (;@4;)
          end
        end
        local.get 4
        i32.eqz
        br_if 0 (;@2;)
        block  ;; label = @3
          i32.const 0
          i32.load offset=1056288
          local.tee 0
          local.get 2
          i32.lt_u
          br_if 0 (;@3;)
          local.get 6
          local.get 0
          local.get 2
          i32.sub
          i32.ge_u
          br_if 1 (;@2;)
        end
        local.get 4
        call $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE
        local.tee 0
        local.get 2
        call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
        local.set 3
        local.get 4
        call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18unlink_large_chunk17h66e0805d86ff34c0E
        block  ;; label = @3
          block  ;; label = @4
            local.get 6
            i32.const 16
            i32.const 8
            call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
            i32.lt_u
            br_if 0 (;@4;)
            local.get 0
            local.get 2
            call $_ZN8dlmalloc8dlmalloc5Chunk34set_size_and_pinuse_of_inuse_chunk17h54212a19225d5a32E
            local.get 3
            local.get 6
            call $_ZN8dlmalloc8dlmalloc5Chunk33set_size_and_pinuse_of_free_chunk17h6fb36837ea4ab16bE
            block  ;; label = @5
              local.get 6
              i32.const 256
              i32.lt_u
              br_if 0 (;@5;)
              local.get 3
              local.get 6
              call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18insert_large_chunk17h8bf72eeabb70ddaaE
              br 2 (;@3;)
            end
            local.get 6
            i32.const -8
            i32.and
            i32.const 1056016
            i32.add
            local.set 4
            block  ;; label = @5
              block  ;; label = @6
                i32.const 0
                i32.load offset=1056280
                local.tee 5
                i32.const 1
                local.get 6
                i32.const 3
                i32.shr_u
                i32.shl
                local.tee 6
                i32.and
                i32.eqz
                br_if 0 (;@6;)
                local.get 4
                i32.load offset=8
                local.set 6
                br 1 (;@5;)
              end
              i32.const 0
              local.get 5
              local.get 6
              i32.or
              i32.store offset=1056280
              local.get 4
              local.set 6
            end
            local.get 4
            local.get 3
            i32.store offset=8
            local.get 6
            local.get 3
            i32.store offset=12
            local.get 3
            local.get 4
            i32.store offset=12
            local.get 3
            local.get 6
            i32.store offset=8
            br 1 (;@3;)
          end
          local.get 0
          local.get 6
          local.get 2
          i32.add
          call $_ZN8dlmalloc8dlmalloc5Chunk20set_inuse_and_pinuse17h134eb0365548205aE
        end
        local.get 0
        call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
        local.tee 6
        br_if 1 (;@1;)
      end
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      i32.const 0
                      i32.load offset=1056288
                      local.tee 6
                      local.get 2
                      i32.ge_u
                      br_if 0 (;@9;)
                      i32.const 0
                      i32.load offset=1056292
                      local.tee 0
                      local.get 2
                      i32.gt_u
                      br_if 2 (;@7;)
                      local.get 1
                      i32.const 1055872
                      local.get 2
                      call $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE
                      local.tee 0
                      i32.sub
                      local.get 0
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      i32.add
                      i32.const 20
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      i32.add
                      i32.const 16
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      i32.add
                      i32.const 8
                      i32.add
                      i32.const 65536
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      call $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$5alloc17hb11e962cef68a804E
                      local.get 1
                      i32.load
                      local.tee 6
                      br_if 1 (;@8;)
                      i32.const 0
                      local.set 6
                      br 8 (;@1;)
                    end
                    i32.const 0
                    i32.load offset=1056296
                    local.set 0
                    block  ;; label = @9
                      local.get 6
                      local.get 2
                      i32.sub
                      local.tee 6
                      i32.const 16
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      i32.ge_u
                      br_if 0 (;@9;)
                      i32.const 0
                      i32.const 0
                      i32.store offset=1056296
                      i32.const 0
                      i32.load offset=1056288
                      local.set 2
                      i32.const 0
                      i32.const 0
                      i32.store offset=1056288
                      local.get 0
                      local.get 2
                      call $_ZN8dlmalloc8dlmalloc5Chunk20set_inuse_and_pinuse17h134eb0365548205aE
                      local.get 0
                      call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
                      local.set 6
                      br 8 (;@1;)
                    end
                    local.get 0
                    local.get 2
                    call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                    local.set 3
                    i32.const 0
                    local.get 6
                    i32.store offset=1056288
                    i32.const 0
                    local.get 3
                    i32.store offset=1056296
                    local.get 3
                    local.get 6
                    call $_ZN8dlmalloc8dlmalloc5Chunk33set_size_and_pinuse_of_free_chunk17h6fb36837ea4ab16bE
                    local.get 0
                    local.get 2
                    call $_ZN8dlmalloc8dlmalloc5Chunk34set_size_and_pinuse_of_inuse_chunk17h54212a19225d5a32E
                    local.get 0
                    call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
                    local.set 6
                    br 7 (;@1;)
                  end
                  local.get 1
                  i32.load offset=8
                  local.set 8
                  i32.const 0
                  i32.const 0
                  i32.load offset=1056304
                  local.get 1
                  i32.load offset=4
                  local.tee 5
                  i32.add
                  local.tee 0
                  i32.store offset=1056304
                  i32.const 0
                  i32.const 0
                  i32.load offset=1056308
                  local.tee 3
                  local.get 0
                  local.get 3
                  local.get 0
                  i32.gt_u
                  select
                  i32.store offset=1056308
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          i32.const 0
                          i32.load offset=1056300
                          i32.eqz
                          br_if 0 (;@11;)
                          i32.const 1056000
                          local.set 0
                          loop  ;; label = @12
                            local.get 6
                            local.get 0
                            call $_ZN8dlmalloc8dlmalloc7Segment3top17h07497f7196b9a54fE
                            i32.eq
                            br_if 2 (;@10;)
                            local.get 0
                            i32.load offset=8
                            local.tee 0
                            br_if 0 (;@12;)
                            br 3 (;@9;)
                          end
                        end
                        i32.const 0
                        i32.load offset=1056316
                        local.tee 0
                        i32.eqz
                        br_if 5 (;@5;)
                        local.get 6
                        local.get 0
                        i32.lt_u
                        br_if 5 (;@5;)
                        br 7 (;@3;)
                      end
                      local.get 0
                      call $_ZN8dlmalloc8dlmalloc7Segment9is_extern17hfc6d4e628d2b8313E
                      br_if 0 (;@9;)
                      local.get 0
                      call $_ZN8dlmalloc8dlmalloc7Segment9sys_flags17h1c26c3dd33e84b43E
                      local.get 8
                      i32.ne
                      br_if 0 (;@9;)
                      local.get 0
                      i32.const 0
                      i32.load offset=1056300
                      call $_ZN8dlmalloc8dlmalloc7Segment5holds17h02477c6e3bcb4213E
                      br_if 1 (;@8;)
                    end
                    i32.const 0
                    i32.const 0
                    i32.load offset=1056316
                    local.tee 0
                    local.get 6
                    local.get 6
                    local.get 0
                    i32.gt_u
                    select
                    i32.store offset=1056316
                    local.get 6
                    local.get 5
                    i32.add
                    local.set 3
                    i32.const 1056000
                    local.set 0
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          loop  ;; label = @12
                            local.get 0
                            i32.load
                            local.get 3
                            i32.eq
                            br_if 1 (;@11;)
                            local.get 0
                            i32.load offset=8
                            local.tee 0
                            br_if 0 (;@12;)
                            br 2 (;@10;)
                          end
                        end
                        local.get 0
                        call $_ZN8dlmalloc8dlmalloc7Segment9is_extern17hfc6d4e628d2b8313E
                        br_if 0 (;@10;)
                        local.get 0
                        call $_ZN8dlmalloc8dlmalloc7Segment9sys_flags17h1c26c3dd33e84b43E
                        local.get 8
                        i32.eq
                        br_if 1 (;@9;)
                      end
                      i32.const 0
                      i32.load offset=1056300
                      local.set 3
                      i32.const 1056000
                      local.set 0
                      block  ;; label = @10
                        loop  ;; label = @11
                          block  ;; label = @12
                            local.get 0
                            i32.load
                            local.get 3
                            i32.gt_u
                            br_if 0 (;@12;)
                            local.get 0
                            call $_ZN8dlmalloc8dlmalloc7Segment3top17h07497f7196b9a54fE
                            local.get 3
                            i32.gt_u
                            br_if 2 (;@10;)
                          end
                          local.get 0
                          i32.load offset=8
                          local.tee 0
                          br_if 0 (;@11;)
                        end
                        i32.const 0
                        local.set 0
                      end
                      local.get 0
                      call $_ZN8dlmalloc8dlmalloc7Segment3top17h07497f7196b9a54fE
                      local.tee 4
                      i32.const 20
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      local.tee 9
                      i32.sub
                      i32.const -23
                      i32.add
                      local.set 0
                      local.get 3
                      local.get 0
                      local.get 0
                      call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
                      local.tee 7
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      local.get 7
                      i32.sub
                      i32.add
                      local.tee 0
                      local.get 0
                      local.get 3
                      i32.const 16
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      i32.add
                      i32.lt_u
                      select
                      local.tee 7
                      call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
                      local.set 10
                      local.get 7
                      local.get 9
                      call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                      local.set 0
                      call $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE
                      local.tee 11
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      local.set 12
                      i32.const 20
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      local.set 13
                      i32.const 16
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      local.set 14
                      i32.const 0
                      local.get 6
                      local.get 6
                      call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
                      local.tee 15
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      local.get 15
                      i32.sub
                      local.tee 16
                      call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                      local.tee 15
                      i32.store offset=1056300
                      i32.const 0
                      local.get 11
                      local.get 5
                      i32.add
                      local.get 14
                      local.get 12
                      local.get 13
                      i32.add
                      i32.add
                      local.get 16
                      i32.add
                      i32.sub
                      local.tee 11
                      i32.store offset=1056292
                      local.get 15
                      local.get 11
                      i32.const 1
                      i32.or
                      i32.store offset=4
                      call $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE
                      local.tee 12
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      local.set 13
                      i32.const 20
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      local.set 14
                      i32.const 16
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      local.set 16
                      local.get 15
                      local.get 11
                      call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                      local.get 16
                      local.get 14
                      local.get 13
                      local.get 12
                      i32.sub
                      i32.add
                      i32.add
                      i32.store offset=4
                      i32.const 0
                      i32.const 2097152
                      i32.store offset=1056312
                      local.get 7
                      local.get 9
                      call $_ZN8dlmalloc8dlmalloc5Chunk34set_size_and_pinuse_of_inuse_chunk17h54212a19225d5a32E
                      i32.const 0
                      i64.load offset=1056000 align=4
                      local.set 17
                      local.get 10
                      i32.const 8
                      i32.add
                      i32.const 0
                      i64.load offset=1056008 align=4
                      i64.store align=4
                      local.get 10
                      local.get 17
                      i64.store align=4
                      i32.const 0
                      local.get 8
                      i32.store offset=1056012
                      i32.const 0
                      local.get 5
                      i32.store offset=1056004
                      i32.const 0
                      local.get 6
                      i32.store offset=1056000
                      i32.const 0
                      local.get 10
                      i32.store offset=1056008
                      loop  ;; label = @10
                        local.get 0
                        i32.const 4
                        call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                        local.set 6
                        local.get 0
                        call $_ZN8dlmalloc8dlmalloc5Chunk14fencepost_head17hfc6527a8ab7d96eaE
                        i32.store offset=4
                        local.get 6
                        local.set 0
                        local.get 6
                        i32.const 4
                        i32.add
                        local.get 4
                        i32.lt_u
                        br_if 0 (;@10;)
                      end
                      local.get 7
                      local.get 3
                      i32.eq
                      br_if 7 (;@2;)
                      local.get 7
                      local.get 3
                      i32.sub
                      local.set 0
                      local.get 3
                      local.get 0
                      local.get 3
                      local.get 0
                      call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                      call $_ZN8dlmalloc8dlmalloc5Chunk20set_free_with_pinuse17h56042b51bf6983deE
                      block  ;; label = @10
                        local.get 0
                        i32.const 256
                        i32.lt_u
                        br_if 0 (;@10;)
                        local.get 3
                        local.get 0
                        call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18insert_large_chunk17h8bf72eeabb70ddaaE
                        br 8 (;@2;)
                      end
                      local.get 0
                      i32.const -8
                      i32.and
                      i32.const 1056016
                      i32.add
                      local.set 6
                      block  ;; label = @10
                        block  ;; label = @11
                          i32.const 0
                          i32.load offset=1056280
                          local.tee 4
                          i32.const 1
                          local.get 0
                          i32.const 3
                          i32.shr_u
                          i32.shl
                          local.tee 0
                          i32.and
                          i32.eqz
                          br_if 0 (;@11;)
                          local.get 6
                          i32.load offset=8
                          local.set 0
                          br 1 (;@10;)
                        end
                        i32.const 0
                        local.get 4
                        local.get 0
                        i32.or
                        i32.store offset=1056280
                        local.get 6
                        local.set 0
                      end
                      local.get 6
                      local.get 3
                      i32.store offset=8
                      local.get 0
                      local.get 3
                      i32.store offset=12
                      local.get 3
                      local.get 6
                      i32.store offset=12
                      local.get 3
                      local.get 0
                      i32.store offset=8
                      br 7 (;@2;)
                    end
                    local.get 0
                    i32.load
                    local.set 4
                    local.get 0
                    local.get 6
                    i32.store
                    local.get 0
                    local.get 0
                    i32.load offset=4
                    local.get 5
                    i32.add
                    i32.store offset=4
                    local.get 6
                    call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
                    local.tee 0
                    i32.const 8
                    call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                    local.set 3
                    local.get 4
                    call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
                    local.tee 5
                    i32.const 8
                    call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                    local.set 7
                    local.get 6
                    local.get 3
                    local.get 0
                    i32.sub
                    i32.add
                    local.tee 6
                    local.get 2
                    call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                    local.set 3
                    local.get 6
                    local.get 2
                    call $_ZN8dlmalloc8dlmalloc5Chunk34set_size_and_pinuse_of_inuse_chunk17h54212a19225d5a32E
                    local.get 4
                    local.get 7
                    local.get 5
                    i32.sub
                    i32.add
                    local.tee 0
                    local.get 2
                    local.get 6
                    i32.add
                    i32.sub
                    local.set 2
                    block  ;; label = @9
                      local.get 0
                      i32.const 0
                      i32.load offset=1056300
                      i32.eq
                      br_if 0 (;@9;)
                      local.get 0
                      i32.const 0
                      i32.load offset=1056296
                      i32.eq
                      br_if 3 (;@6;)
                      local.get 0
                      call $_ZN8dlmalloc8dlmalloc5Chunk5inuse17h824c71bc0f61bb75E
                      br_if 5 (;@4;)
                      block  ;; label = @10
                        block  ;; label = @11
                          local.get 0
                          call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
                          local.tee 4
                          i32.const 256
                          i32.lt_u
                          br_if 0 (;@11;)
                          local.get 0
                          call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18unlink_large_chunk17h66e0805d86ff34c0E
                          br 1 (;@10;)
                        end
                        block  ;; label = @11
                          local.get 0
                          i32.const 12
                          i32.add
                          i32.load
                          local.tee 5
                          local.get 0
                          i32.const 8
                          i32.add
                          i32.load
                          local.tee 7
                          i32.eq
                          br_if 0 (;@11;)
                          local.get 7
                          local.get 5
                          i32.store offset=12
                          local.get 5
                          local.get 7
                          i32.store offset=8
                          br 1 (;@10;)
                        end
                        i32.const 0
                        i32.const 0
                        i32.load offset=1056280
                        i32.const -2
                        local.get 4
                        i32.const 3
                        i32.shr_u
                        i32.rotl
                        i32.and
                        i32.store offset=1056280
                      end
                      local.get 4
                      local.get 2
                      i32.add
                      local.set 2
                      local.get 0
                      local.get 4
                      call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                      local.set 0
                      br 5 (;@4;)
                    end
                    i32.const 0
                    local.get 3
                    i32.store offset=1056300
                    i32.const 0
                    i32.const 0
                    i32.load offset=1056292
                    local.get 2
                    i32.add
                    local.tee 0
                    i32.store offset=1056292
                    local.get 3
                    local.get 0
                    i32.const 1
                    i32.or
                    i32.store offset=4
                    local.get 6
                    call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
                    local.set 6
                    br 7 (;@1;)
                  end
                  local.get 0
                  local.get 0
                  i32.load offset=4
                  local.get 5
                  i32.add
                  i32.store offset=4
                  i32.const 0
                  i32.load offset=1056300
                  i32.const 0
                  i32.load offset=1056292
                  local.get 5
                  i32.add
                  call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$8init_top17h195f4d1bd1c1d152E
                  br 5 (;@2;)
                end
                i32.const 0
                local.get 0
                local.get 2
                i32.sub
                local.tee 6
                i32.store offset=1056292
                i32.const 0
                i32.const 0
                i32.load offset=1056300
                local.tee 0
                local.get 2
                call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                local.tee 3
                i32.store offset=1056300
                local.get 3
                local.get 6
                i32.const 1
                i32.or
                i32.store offset=4
                local.get 0
                local.get 2
                call $_ZN8dlmalloc8dlmalloc5Chunk34set_size_and_pinuse_of_inuse_chunk17h54212a19225d5a32E
                local.get 0
                call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
                local.set 6
                br 5 (;@1;)
              end
              i32.const 0
              local.get 3
              i32.store offset=1056296
              i32.const 0
              i32.const 0
              i32.load offset=1056288
              local.get 2
              i32.add
              local.tee 0
              i32.store offset=1056288
              local.get 3
              local.get 0
              call $_ZN8dlmalloc8dlmalloc5Chunk33set_size_and_pinuse_of_free_chunk17h6fb36837ea4ab16bE
              local.get 6
              call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
              local.set 6
              br 4 (;@1;)
            end
            i32.const 0
            local.get 6
            i32.store offset=1056316
            br 1 (;@3;)
          end
          local.get 3
          local.get 2
          local.get 0
          call $_ZN8dlmalloc8dlmalloc5Chunk20set_free_with_pinuse17h56042b51bf6983deE
          block  ;; label = @4
            local.get 2
            i32.const 256
            i32.lt_u
            br_if 0 (;@4;)
            local.get 3
            local.get 2
            call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18insert_large_chunk17h8bf72eeabb70ddaaE
            local.get 6
            call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
            local.set 6
            br 3 (;@1;)
          end
          local.get 2
          i32.const -8
          i32.and
          i32.const 1056016
          i32.add
          local.set 0
          block  ;; label = @4
            block  ;; label = @5
              i32.const 0
              i32.load offset=1056280
              local.tee 4
              i32.const 1
              local.get 2
              i32.const 3
              i32.shr_u
              i32.shl
              local.tee 2
              i32.and
              i32.eqz
              br_if 0 (;@5;)
              local.get 0
              i32.load offset=8
              local.set 2
              br 1 (;@4;)
            end
            i32.const 0
            local.get 4
            local.get 2
            i32.or
            i32.store offset=1056280
            local.get 0
            local.set 2
          end
          local.get 0
          local.get 3
          i32.store offset=8
          local.get 2
          local.get 3
          i32.store offset=12
          local.get 3
          local.get 0
          i32.store offset=12
          local.get 3
          local.get 2
          i32.store offset=8
          local.get 6
          call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
          local.set 6
          br 2 (;@1;)
        end
        i32.const 0
        i32.const 4095
        i32.store offset=1056320
        i32.const 0
        local.get 8
        i32.store offset=1056012
        i32.const 0
        local.get 5
        i32.store offset=1056004
        i32.const 0
        local.get 6
        i32.store offset=1056000
        i32.const 0
        i32.const 1056016
        i32.store offset=1056028
        i32.const 0
        i32.const 1056024
        i32.store offset=1056036
        i32.const 0
        i32.const 1056016
        i32.store offset=1056024
        i32.const 0
        i32.const 1056032
        i32.store offset=1056044
        i32.const 0
        i32.const 1056024
        i32.store offset=1056032
        i32.const 0
        i32.const 1056040
        i32.store offset=1056052
        i32.const 0
        i32.const 1056032
        i32.store offset=1056040
        i32.const 0
        i32.const 1056048
        i32.store offset=1056060
        i32.const 0
        i32.const 1056040
        i32.store offset=1056048
        i32.const 0
        i32.const 1056056
        i32.store offset=1056068
        i32.const 0
        i32.const 1056048
        i32.store offset=1056056
        i32.const 0
        i32.const 1056064
        i32.store offset=1056076
        i32.const 0
        i32.const 1056056
        i32.store offset=1056064
        i32.const 0
        i32.const 1056072
        i32.store offset=1056084
        i32.const 0
        i32.const 1056064
        i32.store offset=1056072
        i32.const 0
        i32.const 1056080
        i32.store offset=1056092
        i32.const 0
        i32.const 1056072
        i32.store offset=1056080
        i32.const 0
        i32.const 1056080
        i32.store offset=1056088
        i32.const 0
        i32.const 1056088
        i32.store offset=1056100
        i32.const 0
        i32.const 1056088
        i32.store offset=1056096
        i32.const 0
        i32.const 1056096
        i32.store offset=1056108
        i32.const 0
        i32.const 1056096
        i32.store offset=1056104
        i32.const 0
        i32.const 1056104
        i32.store offset=1056116
        i32.const 0
        i32.const 1056104
        i32.store offset=1056112
        i32.const 0
        i32.const 1056112
        i32.store offset=1056124
        i32.const 0
        i32.const 1056112
        i32.store offset=1056120
        i32.const 0
        i32.const 1056120
        i32.store offset=1056132
        i32.const 0
        i32.const 1056120
        i32.store offset=1056128
        i32.const 0
        i32.const 1056128
        i32.store offset=1056140
        i32.const 0
        i32.const 1056128
        i32.store offset=1056136
        i32.const 0
        i32.const 1056136
        i32.store offset=1056148
        i32.const 0
        i32.const 1056136
        i32.store offset=1056144
        i32.const 0
        i32.const 1056144
        i32.store offset=1056156
        i32.const 0
        i32.const 1056152
        i32.store offset=1056164
        i32.const 0
        i32.const 1056144
        i32.store offset=1056152
        i32.const 0
        i32.const 1056160
        i32.store offset=1056172
        i32.const 0
        i32.const 1056152
        i32.store offset=1056160
        i32.const 0
        i32.const 1056168
        i32.store offset=1056180
        i32.const 0
        i32.const 1056160
        i32.store offset=1056168
        i32.const 0
        i32.const 1056176
        i32.store offset=1056188
        i32.const 0
        i32.const 1056168
        i32.store offset=1056176
        i32.const 0
        i32.const 1056184
        i32.store offset=1056196
        i32.const 0
        i32.const 1056176
        i32.store offset=1056184
        i32.const 0
        i32.const 1056192
        i32.store offset=1056204
        i32.const 0
        i32.const 1056184
        i32.store offset=1056192
        i32.const 0
        i32.const 1056200
        i32.store offset=1056212
        i32.const 0
        i32.const 1056192
        i32.store offset=1056200
        i32.const 0
        i32.const 1056208
        i32.store offset=1056220
        i32.const 0
        i32.const 1056200
        i32.store offset=1056208
        i32.const 0
        i32.const 1056216
        i32.store offset=1056228
        i32.const 0
        i32.const 1056208
        i32.store offset=1056216
        i32.const 0
        i32.const 1056224
        i32.store offset=1056236
        i32.const 0
        i32.const 1056216
        i32.store offset=1056224
        i32.const 0
        i32.const 1056232
        i32.store offset=1056244
        i32.const 0
        i32.const 1056224
        i32.store offset=1056232
        i32.const 0
        i32.const 1056240
        i32.store offset=1056252
        i32.const 0
        i32.const 1056232
        i32.store offset=1056240
        i32.const 0
        i32.const 1056248
        i32.store offset=1056260
        i32.const 0
        i32.const 1056240
        i32.store offset=1056248
        i32.const 0
        i32.const 1056256
        i32.store offset=1056268
        i32.const 0
        i32.const 1056248
        i32.store offset=1056256
        i32.const 0
        i32.const 1056264
        i32.store offset=1056276
        i32.const 0
        i32.const 1056256
        i32.store offset=1056264
        i32.const 0
        i32.const 1056264
        i32.store offset=1056272
        call $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE
        local.tee 3
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        local.set 4
        i32.const 20
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        local.set 7
        i32.const 16
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        local.set 8
        i32.const 0
        local.get 6
        local.get 6
        call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
        local.tee 0
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        local.get 0
        i32.sub
        local.tee 10
        call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
        local.tee 0
        i32.store offset=1056300
        i32.const 0
        local.get 3
        local.get 5
        i32.add
        local.get 8
        local.get 4
        local.get 7
        i32.add
        i32.add
        local.get 10
        i32.add
        i32.sub
        local.tee 6
        i32.store offset=1056292
        local.get 0
        local.get 6
        i32.const 1
        i32.or
        i32.store offset=4
        call $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE
        local.tee 3
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        local.set 4
        i32.const 20
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        local.set 5
        i32.const 16
        i32.const 8
        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
        local.set 7
        local.get 0
        local.get 6
        call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
        local.get 7
        local.get 5
        local.get 4
        local.get 3
        i32.sub
        i32.add
        i32.add
        i32.store offset=4
        i32.const 0
        i32.const 2097152
        i32.store offset=1056312
      end
      i32.const 0
      local.set 6
      i32.const 0
      i32.load offset=1056292
      local.tee 0
      local.get 2
      i32.le_u
      br_if 0 (;@1;)
      i32.const 0
      local.get 0
      local.get 2
      i32.sub
      local.tee 6
      i32.store offset=1056292
      i32.const 0
      i32.const 0
      i32.load offset=1056300
      local.tee 0
      local.get 2
      call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
      local.tee 3
      i32.store offset=1056300
      local.get 3
      local.get 6
      i32.const 1
      i32.or
      i32.store offset=4
      local.get 0
      local.get 2
      call $_ZN8dlmalloc8dlmalloc5Chunk34set_size_and_pinuse_of_inuse_chunk17h54212a19225d5a32E
      local.get 0
      call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
      local.set 6
    end
    local.get 1
    i32.const 16
    i32.add
    global.set $__stack_pointer
    local.get 6)
  (func $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$13dispose_chunk17he67418b997d9e729E (type 0) (param i32 i32)
    (local i32 i32 i32 i32)
    local.get 0
    local.get 1
    call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
    local.set 2
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          call $_ZN8dlmalloc8dlmalloc5Chunk6pinuse17hde09ac7fe521978fE
          br_if 0 (;@3;)
          local.get 0
          i32.load
          local.set 3
          block  ;; label = @4
            block  ;; label = @5
              local.get 0
              call $_ZN8dlmalloc8dlmalloc5Chunk7mmapped17h69bc6c14adea384fE
              br_if 0 (;@5;)
              local.get 3
              local.get 1
              i32.add
              local.set 1
              local.get 0
              local.get 3
              call $_ZN8dlmalloc8dlmalloc5Chunk12minus_offset17h1bc6fc1549394b74E
              local.tee 0
              i32.const 0
              i32.load offset=1056296
              i32.ne
              br_if 1 (;@4;)
              local.get 2
              i32.load offset=4
              i32.const 3
              i32.and
              i32.const 3
              i32.ne
              br_if 2 (;@3;)
              i32.const 0
              local.get 1
              i32.store offset=1056288
              local.get 0
              local.get 1
              local.get 2
              call $_ZN8dlmalloc8dlmalloc5Chunk20set_free_with_pinuse17h56042b51bf6983deE
              return
            end
            i32.const 1055872
            local.get 0
            local.get 3
            i32.sub
            local.get 3
            local.get 1
            i32.add
            i32.const 16
            i32.add
            local.tee 0
            call $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$4free17h447c2248b31fe256E
            i32.eqz
            br_if 2 (;@2;)
            i32.const 0
            i32.const 0
            i32.load offset=1056304
            local.get 0
            i32.sub
            i32.store offset=1056304
            return
          end
          block  ;; label = @4
            local.get 3
            i32.const 256
            i32.lt_u
            br_if 0 (;@4;)
            local.get 0
            call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18unlink_large_chunk17h66e0805d86ff34c0E
            br 1 (;@3;)
          end
          block  ;; label = @4
            local.get 0
            i32.const 12
            i32.add
            i32.load
            local.tee 4
            local.get 0
            i32.const 8
            i32.add
            i32.load
            local.tee 5
            i32.eq
            br_if 0 (;@4;)
            local.get 5
            local.get 4
            i32.store offset=12
            local.get 4
            local.get 5
            i32.store offset=8
            br 1 (;@3;)
          end
          i32.const 0
          i32.const 0
          i32.load offset=1056280
          i32.const -2
          local.get 3
          i32.const 3
          i32.shr_u
          i32.rotl
          i32.and
          i32.store offset=1056280
        end
        block  ;; label = @3
          local.get 2
          call $_ZN8dlmalloc8dlmalloc5Chunk6cinuse17hf9aa71d0c0fa920bE
          i32.eqz
          br_if 0 (;@3;)
          local.get 0
          local.get 1
          local.get 2
          call $_ZN8dlmalloc8dlmalloc5Chunk20set_free_with_pinuse17h56042b51bf6983deE
          br 2 (;@1;)
        end
        block  ;; label = @3
          block  ;; label = @4
            local.get 2
            i32.const 0
            i32.load offset=1056300
            i32.eq
            br_if 0 (;@4;)
            local.get 2
            i32.const 0
            i32.load offset=1056296
            i32.ne
            br_if 1 (;@3;)
            i32.const 0
            local.get 0
            i32.store offset=1056296
            i32.const 0
            i32.const 0
            i32.load offset=1056288
            local.get 1
            i32.add
            local.tee 1
            i32.store offset=1056288
            local.get 0
            local.get 1
            call $_ZN8dlmalloc8dlmalloc5Chunk33set_size_and_pinuse_of_free_chunk17h6fb36837ea4ab16bE
            return
          end
          i32.const 0
          local.get 0
          i32.store offset=1056300
          i32.const 0
          i32.const 0
          i32.load offset=1056292
          local.get 1
          i32.add
          local.tee 1
          i32.store offset=1056292
          local.get 0
          local.get 1
          i32.const 1
          i32.or
          i32.store offset=4
          local.get 0
          i32.const 0
          i32.load offset=1056296
          i32.ne
          br_if 1 (;@2;)
          i32.const 0
          i32.const 0
          i32.store offset=1056288
          i32.const 0
          i32.const 0
          i32.store offset=1056296
          return
        end
        local.get 2
        call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
        local.tee 3
        local.get 1
        i32.add
        local.set 1
        block  ;; label = @3
          block  ;; label = @4
            local.get 3
            i32.const 256
            i32.lt_u
            br_if 0 (;@4;)
            local.get 2
            call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18unlink_large_chunk17h66e0805d86ff34c0E
            br 1 (;@3;)
          end
          block  ;; label = @4
            local.get 2
            i32.const 12
            i32.add
            i32.load
            local.tee 4
            local.get 2
            i32.const 8
            i32.add
            i32.load
            local.tee 2
            i32.eq
            br_if 0 (;@4;)
            local.get 2
            local.get 4
            i32.store offset=12
            local.get 4
            local.get 2
            i32.store offset=8
            br 1 (;@3;)
          end
          i32.const 0
          i32.const 0
          i32.load offset=1056280
          i32.const -2
          local.get 3
          i32.const 3
          i32.shr_u
          i32.rotl
          i32.and
          i32.store offset=1056280
        end
        local.get 0
        local.get 1
        call $_ZN8dlmalloc8dlmalloc5Chunk33set_size_and_pinuse_of_free_chunk17h6fb36837ea4ab16bE
        local.get 0
        i32.const 0
        i32.load offset=1056296
        i32.ne
        br_if 1 (;@1;)
        i32.const 0
        local.get 1
        i32.store offset=1056288
      end
      return
    end
    block  ;; label = @1
      local.get 1
      i32.const 256
      i32.lt_u
      br_if 0 (;@1;)
      local.get 0
      local.get 1
      call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18insert_large_chunk17h8bf72eeabb70ddaaE
      return
    end
    local.get 1
    i32.const -8
    i32.and
    i32.const 1056016
    i32.add
    local.set 2
    block  ;; label = @1
      block  ;; label = @2
        i32.const 0
        i32.load offset=1056280
        local.tee 3
        i32.const 1
        local.get 1
        i32.const 3
        i32.shr_u
        i32.shl
        local.tee 1
        i32.and
        i32.eqz
        br_if 0 (;@2;)
        local.get 2
        i32.load offset=8
        local.set 1
        br 1 (;@1;)
      end
      i32.const 0
      local.get 3
      local.get 1
      i32.or
      i32.store offset=1056280
      local.get 2
      local.set 1
    end
    local.get 2
    local.get 0
    i32.store offset=8
    local.get 1
    local.get 0
    i32.store offset=12
    local.get 0
    local.get 2
    i32.store offset=12
    local.get 0
    local.get 1
    i32.store offset=8)
  (func $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18unlink_large_chunk17h66e0805d86ff34c0E (type 11) (param i32)
    (local i32 i32 i32 i32 i32)
    local.get 0
    i32.load offset=24
    local.set 1
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          call $_ZN8dlmalloc8dlmalloc9TreeChunk4next17h395c3bed0a2e5517E
          local.get 0
          i32.ne
          br_if 0 (;@3;)
          local.get 0
          i32.const 20
          i32.const 16
          local.get 0
          i32.const 20
          i32.add
          local.tee 2
          i32.load
          local.tee 3
          select
          i32.add
          i32.load
          local.tee 4
          br_if 1 (;@2;)
          i32.const 0
          local.set 2
          br 2 (;@1;)
        end
        local.get 0
        call $_ZN8dlmalloc8dlmalloc9TreeChunk4prev17h67b7d7330692af94E
        local.tee 4
        local.get 0
        call $_ZN8dlmalloc8dlmalloc9TreeChunk4next17h395c3bed0a2e5517E
        local.tee 2
        call $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE
        i32.store offset=12
        local.get 2
        local.get 4
        call $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE
        i32.store offset=8
        br 1 (;@1;)
      end
      local.get 2
      local.get 0
      i32.const 16
      i32.add
      local.get 3
      select
      local.set 3
      loop  ;; label = @2
        local.get 3
        local.set 5
        local.get 4
        local.tee 2
        i32.const 20
        i32.add
        local.tee 4
        local.get 2
        i32.const 16
        i32.add
        local.get 4
        i32.load
        local.tee 4
        select
        local.set 3
        local.get 2
        i32.const 20
        i32.const 16
        local.get 4
        select
        i32.add
        i32.load
        local.tee 4
        br_if 0 (;@2;)
      end
      local.get 5
      i32.const 0
      i32.store
    end
    block  ;; label = @1
      local.get 1
      i32.eqz
      br_if 0 (;@1;)
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          i32.load offset=28
          i32.const 2
          i32.shl
          i32.const 1055872
          i32.add
          local.tee 4
          i32.load
          local.get 0
          i32.eq
          br_if 0 (;@3;)
          local.get 1
          i32.const 16
          i32.const 20
          local.get 1
          i32.load offset=16
          local.get 0
          i32.eq
          select
          i32.add
          local.get 2
          i32.store
          local.get 2
          i32.eqz
          br_if 2 (;@1;)
          br 1 (;@2;)
        end
        local.get 4
        local.get 2
        i32.store
        local.get 2
        br_if 0 (;@2;)
        i32.const 0
        i32.const 0
        i32.load offset=1056284
        i32.const -2
        local.get 0
        i32.load offset=28
        i32.rotl
        i32.and
        i32.store offset=1056284
        return
      end
      local.get 2
      local.get 1
      i32.store offset=24
      block  ;; label = @2
        local.get 0
        i32.load offset=16
        local.tee 4
        i32.eqz
        br_if 0 (;@2;)
        local.get 2
        local.get 4
        i32.store offset=16
        local.get 4
        local.get 2
        i32.store offset=24
      end
      local.get 0
      i32.const 20
      i32.add
      i32.load
      local.tee 4
      i32.eqz
      br_if 0 (;@1;)
      local.get 2
      i32.const 20
      i32.add
      local.get 4
      i32.store
      local.get 4
      local.get 2
      i32.store offset=24
      return
    end)
  (func $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18insert_large_chunk17h8bf72eeabb70ddaaE (type 0) (param i32 i32)
    (local i32 i32 i32 i32 i32)
    i32.const 0
    local.set 2
    block  ;; label = @1
      local.get 1
      i32.const 256
      i32.lt_u
      br_if 0 (;@1;)
      i32.const 31
      local.set 2
      local.get 1
      i32.const 16777215
      i32.gt_u
      br_if 0 (;@1;)
      local.get 1
      i32.const 6
      local.get 1
      i32.const 8
      i32.shr_u
      i32.clz
      local.tee 2
      i32.sub
      i32.shr_u
      i32.const 1
      i32.and
      local.get 2
      i32.const 1
      i32.shl
      i32.sub
      i32.const 62
      i32.add
      local.set 2
    end
    local.get 0
    i64.const 0
    i64.store offset=16 align=4
    local.get 0
    local.get 2
    i32.store offset=28
    local.get 2
    i32.const 2
    i32.shl
    i32.const 1055872
    i32.add
    local.set 3
    local.get 0
    call $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE
    local.set 4
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              i32.const 0
              i32.load offset=1056284
              local.tee 5
              i32.const 1
              local.get 2
              i32.shl
              local.tee 6
              i32.and
              i32.eqz
              br_if 0 (;@5;)
              local.get 3
              i32.load
              local.set 5
              local.get 2
              call $_ZN8dlmalloc8dlmalloc24leftshift_for_tree_index17hf578475f2b108550E
              local.set 2
              local.get 5
              call $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE
              call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
              local.get 1
              i32.ne
              br_if 1 (;@4;)
              local.get 5
              local.set 2
              br 2 (;@3;)
            end
            i32.const 0
            local.get 5
            local.get 6
            i32.or
            i32.store offset=1056284
            local.get 3
            local.get 0
            i32.store
            local.get 0
            local.get 3
            i32.store offset=24
            br 3 (;@1;)
          end
          local.get 1
          local.get 2
          i32.shl
          local.set 3
          loop  ;; label = @4
            local.get 5
            local.get 3
            i32.const 29
            i32.shr_u
            i32.const 4
            i32.and
            i32.add
            i32.const 16
            i32.add
            local.tee 6
            i32.load
            local.tee 2
            i32.eqz
            br_if 2 (;@2;)
            local.get 3
            i32.const 1
            i32.shl
            local.set 3
            local.get 2
            local.set 5
            local.get 2
            call $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE
            call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
            local.get 1
            i32.ne
            br_if 0 (;@4;)
          end
        end
        local.get 2
        call $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE
        local.tee 2
        i32.load offset=8
        local.tee 3
        local.get 4
        i32.store offset=12
        local.get 2
        local.get 4
        i32.store offset=8
        local.get 4
        local.get 2
        i32.store offset=12
        local.get 4
        local.get 3
        i32.store offset=8
        local.get 0
        i32.const 0
        i32.store offset=24
        return
      end
      local.get 6
      local.get 0
      i32.store
      local.get 0
      local.get 5
      i32.store offset=24
    end
    local.get 4
    local.get 4
    i32.store offset=8
    local.get 4
    local.get 4
    i32.store offset=12)
  (func $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$23release_unused_segments17hb148d0d922675a0eE (type 14) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    i32.const 0
    local.set 0
    i32.const 0
    local.set 1
    block  ;; label = @1
      i32.const 0
      i32.load offset=1056008
      local.tee 2
      i32.eqz
      br_if 0 (;@1;)
      i32.const 1056000
      local.set 3
      i32.const 0
      local.set 1
      i32.const 0
      local.set 0
      loop  ;; label = @2
        local.get 2
        local.tee 4
        i32.load offset=8
        local.set 2
        local.get 4
        i32.load offset=4
        local.set 5
        local.get 4
        i32.load
        local.set 6
        block  ;; label = @3
          block  ;; label = @4
            i32.const 1055872
            local.get 4
            i32.const 12
            i32.add
            i32.load
            i32.const 1
            i32.shr_u
            call $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$16can_release_part17h79bb8fcc390ac09aE
            i32.eqz
            br_if 0 (;@4;)
            local.get 4
            call $_ZN8dlmalloc8dlmalloc7Segment9is_extern17hfc6d4e628d2b8313E
            br_if 0 (;@4;)
            local.get 6
            local.get 6
            call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
            local.tee 7
            i32.const 8
            call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
            local.get 7
            i32.sub
            i32.add
            local.tee 7
            call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
            local.set 8
            call $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE
            local.tee 9
            i32.const 8
            call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
            local.set 10
            i32.const 20
            i32.const 8
            call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
            local.set 11
            i32.const 16
            i32.const 8
            call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
            local.set 12
            local.get 7
            call $_ZN8dlmalloc8dlmalloc5Chunk5inuse17h824c71bc0f61bb75E
            br_if 0 (;@4;)
            local.get 7
            local.get 8
            i32.add
            local.get 6
            local.get 9
            local.get 5
            i32.add
            local.get 10
            local.get 11
            i32.add
            local.get 12
            i32.add
            i32.sub
            i32.add
            i32.lt_u
            br_if 0 (;@4;)
            block  ;; label = @5
              block  ;; label = @6
                local.get 7
                i32.const 0
                i32.load offset=1056296
                i32.eq
                br_if 0 (;@6;)
                local.get 7
                call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18unlink_large_chunk17h66e0805d86ff34c0E
                br 1 (;@5;)
              end
              i32.const 0
              i32.const 0
              i32.store offset=1056288
              i32.const 0
              i32.const 0
              i32.store offset=1056296
            end
            block  ;; label = @5
              i32.const 1055872
              local.get 6
              local.get 5
              call $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$4free17h447c2248b31fe256E
              br_if 0 (;@5;)
              local.get 7
              local.get 8
              call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18insert_large_chunk17h8bf72eeabb70ddaaE
              br 1 (;@4;)
            end
            i32.const 0
            i32.const 0
            i32.load offset=1056304
            local.get 5
            i32.sub
            i32.store offset=1056304
            local.get 3
            local.get 2
            i32.store offset=8
            local.get 5
            local.get 1
            i32.add
            local.set 1
            br 1 (;@3;)
          end
          local.get 4
          local.set 3
        end
        local.get 0
        i32.const 1
        i32.add
        local.set 0
        local.get 2
        br_if 0 (;@2;)
      end
    end
    i32.const 0
    local.get 0
    i32.const 4095
    local.get 0
    i32.const 4095
    i32.gt_u
    select
    i32.store offset=1056320
    local.get 1)
  (func $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$4free17ha64c8bc3953e2079E (type 11) (param i32)
    (local i32 i32 i32 i32 i32 i32)
    local.get 0
    call $_ZN8dlmalloc8dlmalloc5Chunk8from_mem17hb5c408f151442e99E
    local.set 0
    local.get 0
    local.get 0
    call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
    local.tee 1
    call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
    local.set 2
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          call $_ZN8dlmalloc8dlmalloc5Chunk6pinuse17hde09ac7fe521978fE
          br_if 0 (;@3;)
          local.get 0
          i32.load
          local.set 3
          block  ;; label = @4
            block  ;; label = @5
              local.get 0
              call $_ZN8dlmalloc8dlmalloc5Chunk7mmapped17h69bc6c14adea384fE
              br_if 0 (;@5;)
              local.get 3
              local.get 1
              i32.add
              local.set 1
              local.get 0
              local.get 3
              call $_ZN8dlmalloc8dlmalloc5Chunk12minus_offset17h1bc6fc1549394b74E
              local.tee 0
              i32.const 0
              i32.load offset=1056296
              i32.ne
              br_if 1 (;@4;)
              local.get 2
              i32.load offset=4
              i32.const 3
              i32.and
              i32.const 3
              i32.ne
              br_if 2 (;@3;)
              i32.const 0
              local.get 1
              i32.store offset=1056288
              local.get 0
              local.get 1
              local.get 2
              call $_ZN8dlmalloc8dlmalloc5Chunk20set_free_with_pinuse17h56042b51bf6983deE
              return
            end
            i32.const 1055872
            local.get 0
            local.get 3
            i32.sub
            local.get 3
            local.get 1
            i32.add
            i32.const 16
            i32.add
            local.tee 0
            call $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$4free17h447c2248b31fe256E
            i32.eqz
            br_if 2 (;@2;)
            i32.const 0
            i32.const 0
            i32.load offset=1056304
            local.get 0
            i32.sub
            i32.store offset=1056304
            return
          end
          block  ;; label = @4
            local.get 3
            i32.const 256
            i32.lt_u
            br_if 0 (;@4;)
            local.get 0
            call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18unlink_large_chunk17h66e0805d86ff34c0E
            br 1 (;@3;)
          end
          block  ;; label = @4
            local.get 0
            i32.const 12
            i32.add
            i32.load
            local.tee 4
            local.get 0
            i32.const 8
            i32.add
            i32.load
            local.tee 5
            i32.eq
            br_if 0 (;@4;)
            local.get 5
            local.get 4
            i32.store offset=12
            local.get 4
            local.get 5
            i32.store offset=8
            br 1 (;@3;)
          end
          i32.const 0
          i32.const 0
          i32.load offset=1056280
          i32.const -2
          local.get 3
          i32.const 3
          i32.shr_u
          i32.rotl
          i32.and
          i32.store offset=1056280
        end
        block  ;; label = @3
          block  ;; label = @4
            local.get 2
            call $_ZN8dlmalloc8dlmalloc5Chunk6cinuse17hf9aa71d0c0fa920bE
            i32.eqz
            br_if 0 (;@4;)
            local.get 0
            local.get 1
            local.get 2
            call $_ZN8dlmalloc8dlmalloc5Chunk20set_free_with_pinuse17h56042b51bf6983deE
            br 1 (;@3;)
          end
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 2
                  i32.const 0
                  i32.load offset=1056300
                  i32.eq
                  br_if 0 (;@7;)
                  local.get 2
                  i32.const 0
                  i32.load offset=1056296
                  i32.ne
                  br_if 1 (;@6;)
                  i32.const 0
                  local.get 0
                  i32.store offset=1056296
                  i32.const 0
                  i32.const 0
                  i32.load offset=1056288
                  local.get 1
                  i32.add
                  local.tee 1
                  i32.store offset=1056288
                  local.get 0
                  local.get 1
                  call $_ZN8dlmalloc8dlmalloc5Chunk33set_size_and_pinuse_of_free_chunk17h6fb36837ea4ab16bE
                  return
                end
                i32.const 0
                local.get 0
                i32.store offset=1056300
                i32.const 0
                i32.const 0
                i32.load offset=1056292
                local.get 1
                i32.add
                local.tee 1
                i32.store offset=1056292
                local.get 0
                local.get 1
                i32.const 1
                i32.or
                i32.store offset=4
                local.get 0
                i32.const 0
                i32.load offset=1056296
                i32.eq
                br_if 1 (;@5;)
                br 2 (;@4;)
              end
              local.get 2
              call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
              local.tee 3
              local.get 1
              i32.add
              local.set 1
              block  ;; label = @6
                block  ;; label = @7
                  local.get 3
                  i32.const 256
                  i32.lt_u
                  br_if 0 (;@7;)
                  local.get 2
                  call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18unlink_large_chunk17h66e0805d86ff34c0E
                  br 1 (;@6;)
                end
                block  ;; label = @7
                  local.get 2
                  i32.const 12
                  i32.add
                  i32.load
                  local.tee 4
                  local.get 2
                  i32.const 8
                  i32.add
                  i32.load
                  local.tee 2
                  i32.eq
                  br_if 0 (;@7;)
                  local.get 2
                  local.get 4
                  i32.store offset=12
                  local.get 4
                  local.get 2
                  i32.store offset=8
                  br 1 (;@6;)
                end
                i32.const 0
                i32.const 0
                i32.load offset=1056280
                i32.const -2
                local.get 3
                i32.const 3
                i32.shr_u
                i32.rotl
                i32.and
                i32.store offset=1056280
              end
              local.get 0
              local.get 1
              call $_ZN8dlmalloc8dlmalloc5Chunk33set_size_and_pinuse_of_free_chunk17h6fb36837ea4ab16bE
              local.get 0
              i32.const 0
              i32.load offset=1056296
              i32.ne
              br_if 2 (;@3;)
              i32.const 0
              local.get 1
              i32.store offset=1056288
              br 3 (;@2;)
            end
            i32.const 0
            i32.const 0
            i32.store offset=1056288
            i32.const 0
            i32.const 0
            i32.store offset=1056296
          end
          i32.const 0
          i32.load offset=1056312
          local.get 1
          i32.ge_u
          br_if 1 (;@2;)
          call $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE
          local.tee 0
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          local.set 1
          i32.const 20
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          local.set 2
          i32.const 16
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          local.set 3
          i32.const 0
          i32.const 16
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          i32.const 2
          i32.shl
          i32.sub
          local.tee 4
          local.get 0
          local.get 3
          local.get 1
          local.get 2
          i32.add
          i32.add
          i32.sub
          i32.const -65544
          i32.add
          i32.const -9
          i32.and
          i32.const -3
          i32.add
          local.tee 0
          local.get 4
          local.get 0
          i32.lt_u
          select
          i32.eqz
          br_if 1 (;@2;)
          i32.const 0
          i32.load offset=1056300
          i32.eqz
          br_if 1 (;@2;)
          call $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE
          local.tee 0
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          local.set 1
          i32.const 20
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          local.set 3
          i32.const 16
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          local.set 4
          i32.const 0
          local.set 2
          block  ;; label = @4
            i32.const 0
            i32.load offset=1056292
            local.tee 5
            local.get 4
            local.get 3
            local.get 1
            local.get 0
            i32.sub
            i32.add
            i32.add
            local.tee 0
            i32.le_u
            br_if 0 (;@4;)
            local.get 5
            local.get 0
            i32.sub
            i32.const 65535
            i32.add
            i32.const -65536
            i32.and
            local.tee 4
            i32.const -65536
            i32.add
            local.set 3
            i32.const 0
            i32.load offset=1056300
            local.set 1
            i32.const 1056000
            local.set 0
            block  ;; label = @5
              loop  ;; label = @6
                block  ;; label = @7
                  local.get 0
                  i32.load
                  local.get 1
                  i32.gt_u
                  br_if 0 (;@7;)
                  local.get 0
                  call $_ZN8dlmalloc8dlmalloc7Segment3top17h07497f7196b9a54fE
                  local.get 1
                  i32.gt_u
                  br_if 2 (;@5;)
                end
                local.get 0
                i32.load offset=8
                local.tee 0
                br_if 0 (;@6;)
              end
              i32.const 0
              local.set 0
            end
            i32.const 0
            local.set 2
            local.get 0
            call $_ZN8dlmalloc8dlmalloc7Segment9is_extern17hfc6d4e628d2b8313E
            br_if 0 (;@4;)
            i32.const 1055872
            local.get 0
            i32.const 12
            i32.add
            i32.load
            i32.const 1
            i32.shr_u
            call $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$16can_release_part17h79bb8fcc390ac09aE
            i32.eqz
            br_if 0 (;@4;)
            local.get 0
            i32.load offset=4
            local.get 3
            i32.lt_u
            br_if 0 (;@4;)
            i32.const 1056000
            local.set 1
            loop  ;; label = @5
              local.get 0
              local.get 1
              call $_ZN8dlmalloc8dlmalloc7Segment5holds17h02477c6e3bcb4213E
              br_if 1 (;@4;)
              local.get 1
              i32.load offset=8
              local.tee 1
              br_if 0 (;@5;)
            end
            i32.const 1055872
            local.get 0
            i32.load
            local.get 0
            i32.load offset=4
            local.tee 1
            local.get 1
            local.get 3
            i32.sub
            call $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$9free_part17h38a6e48730b8faaeE
            i32.eqz
            br_if 0 (;@4;)
            local.get 3
            i32.eqz
            br_if 0 (;@4;)
            local.get 0
            local.get 0
            i32.load offset=4
            local.get 3
            i32.sub
            i32.store offset=4
            i32.const 0
            i32.const 0
            i32.load offset=1056304
            local.get 3
            i32.sub
            i32.store offset=1056304
            i32.const 0
            i32.load offset=1056292
            local.set 1
            i32.const 0
            i32.load offset=1056300
            local.set 0
            i32.const 0
            local.get 0
            local.get 0
            call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
            local.tee 2
            i32.const 8
            call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
            local.get 2
            i32.sub
            local.tee 2
            call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
            local.tee 0
            i32.store offset=1056300
            i32.const 0
            local.get 1
            local.get 4
            local.get 2
            i32.add
            i32.sub
            i32.const 65536
            i32.add
            local.tee 1
            i32.store offset=1056292
            local.get 0
            local.get 1
            i32.const 1
            i32.or
            i32.store offset=4
            call $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE
            local.tee 2
            i32.const 8
            call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
            local.set 4
            i32.const 20
            i32.const 8
            call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
            local.set 5
            i32.const 16
            i32.const 8
            call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
            local.set 6
            local.get 0
            local.get 1
            call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
            local.get 6
            local.get 5
            local.get 4
            local.get 2
            i32.sub
            i32.add
            i32.add
            i32.store offset=4
            i32.const 0
            i32.const 2097152
            i32.store offset=1056312
            local.get 3
            local.set 2
          end
          call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$23release_unused_segments17hb148d0d922675a0eE
          i32.const 0
          local.get 2
          i32.sub
          i32.ne
          br_if 1 (;@2;)
          i32.const 0
          i32.load offset=1056292
          i32.const 0
          i32.load offset=1056312
          i32.le_u
          br_if 1 (;@2;)
          i32.const 0
          i32.const -1
          i32.store offset=1056312
          return
        end
        local.get 1
        i32.const 256
        i32.lt_u
        br_if 1 (;@1;)
        local.get 0
        local.get 1
        call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18insert_large_chunk17h8bf72eeabb70ddaaE
        i32.const 0
        i32.const 0
        i32.load offset=1056320
        i32.const -1
        i32.add
        local.tee 0
        i32.store offset=1056320
        local.get 0
        br_if 0 (;@2;)
        call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$23release_unused_segments17hb148d0d922675a0eE
        drop
        return
      end
      return
    end
    local.get 1
    i32.const -8
    i32.and
    i32.const 1056016
    i32.add
    local.set 2
    block  ;; label = @1
      block  ;; label = @2
        i32.const 0
        i32.load offset=1056280
        local.tee 3
        i32.const 1
        local.get 1
        i32.const 3
        i32.shr_u
        i32.shl
        local.tee 1
        i32.and
        i32.eqz
        br_if 0 (;@2;)
        local.get 2
        i32.load offset=8
        local.set 1
        br 1 (;@1;)
      end
      i32.const 0
      local.get 3
      local.get 1
      i32.or
      i32.store offset=1056280
      local.get 2
      local.set 1
    end
    local.get 2
    local.get 0
    i32.store offset=8
    local.get 1
    local.get 0
    i32.store offset=12
    local.get 0
    local.get 2
    i32.store offset=12
    local.get 0
    local.get 1
    i32.store offset=8)
  (func $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$8init_top17h195f4d1bd1c1d152E (type 0) (param i32 i32)
    (local i32 i32 i32 i32)
    local.get 0
    local.get 0
    call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E
    local.tee 2
    i32.const 8
    call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
    local.get 2
    i32.sub
    local.tee 2
    call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
    local.set 0
    i32.const 0
    local.get 1
    local.get 2
    i32.sub
    local.tee 1
    i32.store offset=1056292
    i32.const 0
    local.get 0
    i32.store offset=1056300
    local.get 0
    local.get 1
    i32.const 1
    i32.or
    i32.store offset=4
    call $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE
    local.tee 2
    i32.const 8
    call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
    local.set 3
    i32.const 20
    i32.const 8
    call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
    local.set 4
    i32.const 16
    i32.const 8
    call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
    local.set 5
    local.get 0
    local.get 1
    call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
    local.get 5
    local.get 4
    local.get 3
    local.get 2
    i32.sub
    i32.add
    i32.add
    i32.store offset=4
    i32.const 0
    i32.const 2097152
    i32.store offset=1056312)
  (func $_ZN3std10sys_common9backtrace26__rust_end_short_backtrace17hf9832bd8d880a01bE (type 11) (param i32)
    local.get 0
    call $_ZN3std9panicking19begin_panic_handler28_$u7b$$u7b$closure$u7d$$u7d$17h199731297267a14eE
    unreachable)
  (func $_ZN3std9panicking19begin_panic_handler28_$u7b$$u7b$closure$u7d$$u7d$17h199731297267a14eE (type 11) (param i32)
    (local i32 i32 i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 1
    global.set $__stack_pointer
    local.get 0
    i32.load
    local.tee 2
    i32.const 20
    i32.add
    i32.load
    local.set 3
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 2
            i32.const 12
            i32.add
            i32.load
            br_table 0 (;@4;) 1 (;@3;) 3 (;@1;)
          end
          local.get 3
          br_if 2 (;@1;)
          i32.const 1048644
          local.set 2
          i32.const 0
          local.set 3
          br 1 (;@2;)
        end
        local.get 3
        br_if 1 (;@1;)
        local.get 2
        i32.load offset=8
        local.tee 2
        i32.load offset=4
        local.set 3
        local.get 2
        i32.load
        local.set 2
      end
      local.get 1
      local.get 3
      i32.store offset=4
      local.get 1
      local.get 2
      i32.store
      local.get 1
      i32.const 1048892
      local.get 0
      i32.load offset=4
      local.tee 2
      call $_ZN4core5panic10panic_info9PanicInfo7message17hdce92b3e180291bdE
      local.get 0
      i32.load offset=8
      local.get 2
      call $_ZN4core5panic10panic_info9PanicInfo10can_unwind17hd9c7a52d1ea179a2E
      call $_ZN3std9panicking20rust_panic_with_hook17hf8f067983b30da10E
      unreachable
    end
    local.get 1
    i32.const 0
    i32.store offset=4
    local.get 1
    local.get 2
    i32.store offset=12
    local.get 1
    i32.const 1048872
    local.get 0
    i32.load offset=4
    local.tee 2
    call $_ZN4core5panic10panic_info9PanicInfo7message17hdce92b3e180291bdE
    local.get 0
    i32.load offset=8
    local.get 2
    call $_ZN4core5panic10panic_info9PanicInfo10can_unwind17hd9c7a52d1ea179a2E
    call $_ZN3std9panicking20rust_panic_with_hook17hf8f067983b30da10E
    unreachable)
  (func $_ZN3std5alloc24default_alloc_error_hook17h8183fc7022a4a0e7E (type 0) (param i32 i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    block  ;; label = @1
      i32.const 0
      i32.load8_u offset=1055844
      i32.eqz
      br_if 0 (;@1;)
      local.get 2
      i32.const 20
      i32.add
      i32.const 2
      i32.store
      local.get 2
      i32.const 28
      i32.add
      i32.const 1
      i32.store
      local.get 2
      i32.const 1048724
      i32.store offset=16
      local.get 2
      i32.const 0
      i32.store offset=8
      local.get 2
      i32.const 4
      i32.store offset=36
      local.get 2
      local.get 0
      i32.store offset=44
      local.get 2
      local.get 2
      i32.const 32
      i32.add
      i32.store offset=24
      local.get 2
      local.get 2
      i32.const 44
      i32.add
      i32.store offset=32
      local.get 2
      i32.const 8
      i32.add
      i32.const 1048764
      call $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E
      unreachable
    end
    local.get 2
    i32.const 48
    i32.add
    global.set $__stack_pointer)
  (func $__rdl_alloc (type 2) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    call $_ZN8dlmalloc17Dlmalloc$LT$A$GT$6malloc17h7a9e47086a367e16E)
  (func $__rdl_dealloc (type 3) (param i32 i32 i32)
    local.get 0
    call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$4free17ha64c8bc3953e2079E)
  (func $__rdl_realloc (type 9) (param i32 i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 2
            i32.const 9
            i32.lt_u
            br_if 0 (;@4;)
            local.get 3
            local.get 2
            call $_ZN8dlmalloc17Dlmalloc$LT$A$GT$6malloc17h7a9e47086a367e16E
            local.tee 2
            br_if 1 (;@3;)
            i32.const 0
            return
          end
          call $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE
          local.tee 1
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          local.set 4
          i32.const 20
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          local.set 5
          i32.const 16
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          local.set 6
          i32.const 0
          local.set 2
          i32.const 0
          i32.const 16
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          i32.const 2
          i32.shl
          i32.sub
          local.tee 7
          local.get 1
          local.get 6
          local.get 4
          local.get 5
          i32.add
          i32.add
          i32.sub
          i32.const -65544
          i32.add
          i32.const -9
          i32.and
          i32.const -3
          i32.add
          local.tee 1
          local.get 7
          local.get 1
          i32.lt_u
          select
          local.get 3
          i32.le_u
          br_if 1 (;@2;)
          i32.const 16
          local.get 3
          i32.const 4
          i32.add
          i32.const 16
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          i32.const -5
          i32.add
          local.get 3
          i32.gt_u
          select
          i32.const 8
          call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
          local.set 4
          local.get 0
          call $_ZN8dlmalloc8dlmalloc5Chunk8from_mem17hb5c408f151442e99E
          local.set 1
          local.get 1
          local.get 1
          call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
          local.tee 5
          call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
          local.set 6
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          local.get 1
                          call $_ZN8dlmalloc8dlmalloc5Chunk7mmapped17h69bc6c14adea384fE
                          br_if 0 (;@11;)
                          local.get 5
                          local.get 4
                          i32.ge_u
                          br_if 1 (;@10;)
                          local.get 6
                          i32.const 0
                          i32.load offset=1056300
                          i32.eq
                          br_if 2 (;@9;)
                          local.get 6
                          i32.const 0
                          i32.load offset=1056296
                          i32.eq
                          br_if 3 (;@8;)
                          local.get 6
                          call $_ZN8dlmalloc8dlmalloc5Chunk6cinuse17hf9aa71d0c0fa920bE
                          br_if 7 (;@4;)
                          local.get 6
                          call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
                          local.tee 7
                          local.get 5
                          i32.add
                          local.tee 5
                          local.get 4
                          i32.lt_u
                          br_if 7 (;@4;)
                          local.get 5
                          local.get 4
                          i32.sub
                          local.set 8
                          local.get 7
                          i32.const 256
                          i32.lt_u
                          br_if 4 (;@7;)
                          local.get 6
                          call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$18unlink_large_chunk17h66e0805d86ff34c0E
                          br 5 (;@6;)
                        end
                        local.get 1
                        call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
                        local.set 5
                        local.get 4
                        i32.const 256
                        i32.lt_u
                        br_if 6 (;@4;)
                        block  ;; label = @11
                          local.get 5
                          local.get 4
                          i32.const 4
                          i32.add
                          i32.lt_u
                          br_if 0 (;@11;)
                          local.get 5
                          local.get 4
                          i32.sub
                          i32.const 131073
                          i32.lt_u
                          br_if 6 (;@5;)
                        end
                        i32.const 1055872
                        local.get 1
                        local.get 1
                        i32.load
                        local.tee 6
                        i32.sub
                        local.get 5
                        local.get 6
                        i32.add
                        i32.const 16
                        i32.add
                        local.tee 7
                        local.get 4
                        i32.const 31
                        i32.add
                        i32.const 1055872
                        call $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$9page_size17h1fc4fb2be482df9eE
                        call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                        local.tee 5
                        i32.const 1
                        call $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$5remap17h7f179872f7b3a976E
                        local.tee 4
                        i32.eqz
                        br_if 6 (;@4;)
                        local.get 4
                        local.get 6
                        i32.add
                        local.tee 1
                        local.get 5
                        local.get 6
                        i32.sub
                        local.tee 3
                        i32.const -16
                        i32.add
                        local.tee 2
                        i32.store offset=4
                        call $_ZN8dlmalloc8dlmalloc5Chunk14fencepost_head17hfc6527a8ab7d96eaE
                        local.set 0
                        local.get 1
                        local.get 2
                        call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                        local.get 0
                        i32.store offset=4
                        local.get 1
                        local.get 3
                        i32.const -12
                        i32.add
                        call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                        i32.const 0
                        i32.store offset=4
                        i32.const 0
                        i32.const 0
                        i32.load offset=1056304
                        local.get 5
                        local.get 7
                        i32.sub
                        i32.add
                        local.tee 3
                        i32.store offset=1056304
                        i32.const 0
                        i32.const 0
                        i32.load offset=1056316
                        local.tee 2
                        local.get 4
                        local.get 4
                        local.get 2
                        i32.gt_u
                        select
                        i32.store offset=1056316
                        i32.const 0
                        i32.const 0
                        i32.load offset=1056308
                        local.tee 2
                        local.get 3
                        local.get 2
                        local.get 3
                        i32.gt_u
                        select
                        i32.store offset=1056308
                        br 9 (;@1;)
                      end
                      local.get 5
                      local.get 4
                      i32.sub
                      local.tee 5
                      i32.const 16
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      i32.lt_u
                      br_if 4 (;@5;)
                      local.get 1
                      local.get 4
                      call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                      local.set 6
                      local.get 1
                      local.get 4
                      call $_ZN8dlmalloc8dlmalloc5Chunk9set_inuse17h5f3bbfae424d3974E
                      local.get 6
                      local.get 5
                      call $_ZN8dlmalloc8dlmalloc5Chunk9set_inuse17h5f3bbfae424d3974E
                      local.get 6
                      local.get 5
                      call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$13dispose_chunk17he67418b997d9e729E
                      br 4 (;@5;)
                    end
                    i32.const 0
                    i32.load offset=1056292
                    local.get 5
                    i32.add
                    local.tee 5
                    local.get 4
                    i32.le_u
                    br_if 4 (;@4;)
                    local.get 1
                    local.get 4
                    call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                    local.set 6
                    local.get 1
                    local.get 4
                    call $_ZN8dlmalloc8dlmalloc5Chunk9set_inuse17h5f3bbfae424d3974E
                    local.get 6
                    local.get 5
                    local.get 4
                    i32.sub
                    local.tee 4
                    i32.const 1
                    i32.or
                    i32.store offset=4
                    i32.const 0
                    local.get 4
                    i32.store offset=1056292
                    i32.const 0
                    local.get 6
                    i32.store offset=1056300
                    br 3 (;@5;)
                  end
                  i32.const 0
                  i32.load offset=1056288
                  local.get 5
                  i32.add
                  local.tee 5
                  local.get 4
                  i32.lt_u
                  br_if 3 (;@4;)
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 5
                      local.get 4
                      i32.sub
                      local.tee 6
                      i32.const 16
                      i32.const 8
                      call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                      i32.ge_u
                      br_if 0 (;@9;)
                      local.get 1
                      local.get 5
                      call $_ZN8dlmalloc8dlmalloc5Chunk9set_inuse17h5f3bbfae424d3974E
                      i32.const 0
                      local.set 6
                      i32.const 0
                      local.set 5
                      br 1 (;@8;)
                    end
                    local.get 1
                    local.get 4
                    call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                    local.tee 5
                    local.get 6
                    call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                    local.set 7
                    local.get 1
                    local.get 4
                    call $_ZN8dlmalloc8dlmalloc5Chunk9set_inuse17h5f3bbfae424d3974E
                    local.get 5
                    local.get 6
                    call $_ZN8dlmalloc8dlmalloc5Chunk33set_size_and_pinuse_of_free_chunk17h6fb36837ea4ab16bE
                    local.get 7
                    call $_ZN8dlmalloc8dlmalloc5Chunk12clear_pinuse17h47112d10bff2c398E
                  end
                  i32.const 0
                  local.get 5
                  i32.store offset=1056296
                  i32.const 0
                  local.get 6
                  i32.store offset=1056288
                  br 2 (;@5;)
                end
                block  ;; label = @7
                  local.get 6
                  i32.const 12
                  i32.add
                  i32.load
                  local.tee 9
                  local.get 6
                  i32.const 8
                  i32.add
                  i32.load
                  local.tee 6
                  i32.eq
                  br_if 0 (;@7;)
                  local.get 6
                  local.get 9
                  i32.store offset=12
                  local.get 9
                  local.get 6
                  i32.store offset=8
                  br 1 (;@6;)
                end
                i32.const 0
                i32.const 0
                i32.load offset=1056280
                i32.const -2
                local.get 7
                i32.const 3
                i32.shr_u
                i32.rotl
                i32.and
                i32.store offset=1056280
              end
              block  ;; label = @6
                local.get 8
                i32.const 16
                i32.const 8
                call $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E
                i32.lt_u
                br_if 0 (;@6;)
                local.get 1
                local.get 4
                call $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE
                local.set 5
                local.get 1
                local.get 4
                call $_ZN8dlmalloc8dlmalloc5Chunk9set_inuse17h5f3bbfae424d3974E
                local.get 5
                local.get 8
                call $_ZN8dlmalloc8dlmalloc5Chunk9set_inuse17h5f3bbfae424d3974E
                local.get 5
                local.get 8
                call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$13dispose_chunk17he67418b997d9e729E
                br 1 (;@5;)
              end
              local.get 1
              local.get 5
              call $_ZN8dlmalloc8dlmalloc5Chunk9set_inuse17h5f3bbfae424d3974E
            end
            local.get 1
            br_if 3 (;@1;)
          end
          local.get 3
          call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$6malloc17hfc360c49b49c4ccfE
          local.tee 4
          i32.eqz
          br_if 1 (;@2;)
          local.get 4
          local.get 0
          local.get 1
          call $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E
          i32.const -8
          i32.const -4
          local.get 1
          call $_ZN8dlmalloc8dlmalloc5Chunk7mmapped17h69bc6c14adea384fE
          select
          i32.add
          local.tee 2
          local.get 3
          local.get 2
          local.get 3
          i32.lt_u
          select
          call $memcpy
          local.set 3
          local.get 0
          call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$4free17ha64c8bc3953e2079E
          local.get 3
          return
        end
        local.get 2
        local.get 0
        local.get 1
        local.get 3
        local.get 1
        local.get 3
        i32.lt_u
        select
        call $memcpy
        drop
        local.get 0
        call $_ZN8dlmalloc8dlmalloc17Dlmalloc$LT$A$GT$4free17ha64c8bc3953e2079E
      end
      local.get 2
      return
    end
    local.get 1
    call $_ZN8dlmalloc8dlmalloc5Chunk7mmapped17h69bc6c14adea384fE
    drop
    local.get 1
    call $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E)
  (func $rust_begin_unwind (type 11) (param i32)
    (local i32 i32 i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 1
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        call $_ZN4core5panic10panic_info9PanicInfo8location17h669fabfe65fb79d4E
        local.tee 2
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        call $_ZN4core5panic10panic_info9PanicInfo7message17hdce92b3e180291bdE
        local.tee 3
        i32.eqz
        br_if 1 (;@1;)
        local.get 1
        local.get 2
        i32.store offset=8
        local.get 1
        local.get 0
        i32.store offset=4
        local.get 1
        local.get 3
        i32.store
        local.get 1
        call $_ZN3std10sys_common9backtrace26__rust_end_short_backtrace17hf9832bd8d880a01bE
        unreachable
      end
      i32.const 1048644
      i32.const 43
      i32.const 1048824
      call $_ZN4core9panicking5panic17he2fadc50abd7e245E
      unreachable
    end
    i32.const 1048644
    i32.const 43
    i32.const 1048808
    call $_ZN4core9panicking5panic17he2fadc50abd7e245E
    unreachable)
  (func $_ZN90_$LT$std..panicking..begin_panic_handler..PanicPayload$u20$as$u20$core..panic..BoxMeUp$GT$8take_box17h94ddafce7f810e6dE (type 0) (param i32 i32)
    (local i32 i32 i32 i64)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    block  ;; label = @1
      local.get 1
      i32.load offset=4
      br_if 0 (;@1;)
      local.get 1
      i32.load offset=12
      local.set 3
      local.get 2
      i32.const 16
      i32.add
      i32.const 8
      i32.add
      local.tee 4
      i32.const 0
      i32.store
      local.get 2
      i64.const 4294967296
      i64.store offset=16
      local.get 2
      local.get 2
      i32.const 16
      i32.add
      i32.store offset=28
      local.get 2
      i32.const 28
      i32.add
      i32.const 1048620
      local.get 3
      call $_ZN4core3fmt5write17h7d7e1e065b0161edE
      drop
      local.get 1
      i32.const 8
      i32.add
      local.get 4
      i32.load
      i32.store
      local.get 1
      local.get 2
      i64.load offset=16
      i64.store align=4
    end
    local.get 1
    i64.load align=4
    local.set 5
    local.get 1
    i64.const 4294967296
    i64.store align=4
    local.get 2
    i32.const 8
    i32.add
    local.tee 3
    local.get 1
    i32.const 8
    i32.add
    local.tee 1
    i32.load
    i32.store
    local.get 1
    i32.const 0
    i32.store
    local.get 2
    local.get 5
    i64.store
    block  ;; label = @1
      i32.const 12
      i32.const 4
      call $__rust_alloc
      local.tee 1
      br_if 0 (;@1;)
      i32.const 12
      i32.const 4
      call $_ZN5alloc5alloc18handle_alloc_error17hc1ba5222dd907d5cE
      unreachable
    end
    local.get 1
    local.get 2
    i64.load
    i64.store align=4
    local.get 1
    i32.const 8
    i32.add
    local.get 3
    i32.load
    i32.store
    local.get 0
    i32.const 1048840
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store
    local.get 2
    i32.const 32
    i32.add
    global.set $__stack_pointer)
  (func $_ZN90_$LT$std..panicking..begin_panic_handler..PanicPayload$u20$as$u20$core..panic..BoxMeUp$GT$3get17h570cda3222383d8dE (type 0) (param i32 i32)
    (local i32 i32 i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    block  ;; label = @1
      local.get 1
      i32.load offset=4
      br_if 0 (;@1;)
      local.get 1
      i32.load offset=12
      local.set 3
      local.get 2
      i32.const 8
      i32.add
      local.tee 4
      i32.const 0
      i32.store
      local.get 2
      i64.const 4294967296
      i64.store
      local.get 2
      local.get 2
      i32.store offset=12
      local.get 2
      i32.const 12
      i32.add
      i32.const 1048620
      local.get 3
      call $_ZN4core3fmt5write17h7d7e1e065b0161edE
      drop
      local.get 1
      i32.const 8
      i32.add
      local.get 4
      i32.load
      i32.store
      local.get 1
      local.get 2
      i64.load
      i64.store align=4
    end
    local.get 0
    i32.const 1048840
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store
    local.get 2
    i32.const 16
    i32.add
    global.set $__stack_pointer)
  (func $_ZN93_$LT$std..panicking..begin_panic_handler..StrPanicPayload$u20$as$u20$core..panic..BoxMeUp$GT$8take_box17hccf548b0011d7a03E (type 0) (param i32 i32)
    (local i32 i32)
    local.get 1
    i32.load offset=4
    local.set 2
    local.get 1
    i32.load
    local.set 3
    block  ;; label = @1
      i32.const 8
      i32.const 4
      call $__rust_alloc
      local.tee 1
      br_if 0 (;@1;)
      i32.const 8
      i32.const 4
      call $_ZN5alloc5alloc18handle_alloc_error17hc1ba5222dd907d5cE
      unreachable
    end
    local.get 1
    local.get 2
    i32.store offset=4
    local.get 1
    local.get 3
    i32.store
    local.get 0
    i32.const 1048856
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func $_ZN93_$LT$std..panicking..begin_panic_handler..StrPanicPayload$u20$as$u20$core..panic..BoxMeUp$GT$3get17h13f86f9299c9908bE (type 0) (param i32 i32)
    local.get 0
    i32.const 1048856
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func $_ZN3std9panicking20rust_panic_with_hook17hf8f067983b30da10E (type 15) (param i32 i32 i32 i32 i32)
    (local i32 i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 5
    global.set $__stack_pointer
    i32.const 0
    i32.const 0
    i32.load offset=1055868
    local.tee 6
    i32.const 1
    i32.add
    i32.store offset=1055868
    block  ;; label = @1
      block  ;; label = @2
        local.get 6
        i32.const 0
        i32.lt_s
        br_if 0 (;@2;)
        i32.const 0
        i32.const 0
        i32.load offset=1056324
        i32.const 1
        i32.add
        local.tee 6
        i32.store offset=1056324
        local.get 6
        i32.const 2
        i32.gt_u
        br_if 0 (;@2;)
        local.get 5
        local.get 4
        i32.store8 offset=24
        local.get 5
        local.get 3
        i32.store offset=20
        local.get 5
        local.get 2
        i32.store offset=16
        local.get 5
        i32.const 1048912
        i32.store offset=12
        local.get 5
        i32.const 1048644
        i32.store offset=8
        i32.const 0
        i32.load offset=1055852
        local.tee 3
        i32.const -1
        i32.le_s
        br_if 0 (;@2;)
        i32.const 0
        local.get 3
        i32.const 1
        i32.add
        i32.store offset=1055852
        block  ;; label = @3
          i32.const 0
          i32.load offset=1055860
          i32.eqz
          br_if 0 (;@3;)
          local.get 5
          local.get 0
          local.get 1
          i32.load offset=16
          call_indirect (type 0)
          local.get 5
          local.get 5
          i64.load
          i64.store offset=8
          i32.const 0
          i32.load offset=1055860
          local.get 5
          i32.const 8
          i32.add
          i32.const 0
          i32.load offset=1055864
          i32.load offset=20
          call_indirect (type 0)
          i32.const 0
          i32.load offset=1055852
          i32.const -1
          i32.add
          local.set 3
        end
        i32.const 0
        local.get 3
        i32.store offset=1055852
        local.get 6
        i32.const 1
        i32.gt_u
        br_if 0 (;@2;)
        local.get 4
        br_if 1 (;@1;)
      end
      unreachable
      unreachable
    end
    local.get 0
    local.get 1
    call $rust_panic
    unreachable)
  (func $rust_panic (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call $__rust_start_panic
    drop
    unreachable
    unreachable)
  (func $_ZN3std5alloc8rust_oom17h528475c4a036b714E (type 0) (param i32 i32)
    (local i32)
    local.get 0
    local.get 1
    i32.const 0
    i32.load offset=1055848
    local.tee 2
    i32.const 7
    local.get 2
    select
    call_indirect (type 0)
    unreachable
    unreachable)
  (func $__rust_start_panic (type 2) (param i32 i32) (result i32)
    unreachable
    unreachable)
  (func $_ZN8dlmalloc8dlmalloc8align_up17h4e755a0b7abf0eb0E (type 2) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.add
    i32.const -1
    i32.add
    i32.const 0
    local.get 1
    i32.sub
    i32.and)
  (func $_ZN8dlmalloc8dlmalloc9left_bits17hf9924c098a7140faE (type 13) (param i32) (result i32)
    local.get 0
    i32.const 1
    i32.shl
    local.tee 0
    i32.const 0
    local.get 0
    i32.sub
    i32.or)
  (func $_ZN8dlmalloc8dlmalloc9least_bit17h8a739bad8428dbdfE (type 13) (param i32) (result i32)
    i32.const 0
    local.get 0
    i32.sub
    local.get 0
    i32.and)
  (func $_ZN8dlmalloc8dlmalloc24leftshift_for_tree_index17hf578475f2b108550E (type 13) (param i32) (result i32)
    i32.const 0
    i32.const 25
    local.get 0
    i32.const 1
    i32.shr_u
    i32.sub
    local.get 0
    i32.const 31
    i32.eq
    select)
  (func $_ZN8dlmalloc8dlmalloc5Chunk14fencepost_head17hfc6527a8ab7d96eaE (type 14) (result i32)
    i32.const 7)
  (func $_ZN8dlmalloc8dlmalloc5Chunk4size17h6460db3a1f5ef376E (type 13) (param i32) (result i32)
    local.get 0
    i32.load offset=4
    i32.const -8
    i32.and)
  (func $_ZN8dlmalloc8dlmalloc5Chunk6cinuse17hf9aa71d0c0fa920bE (type 13) (param i32) (result i32)
    local.get 0
    i32.load8_u offset=4
    i32.const 2
    i32.and
    i32.const 1
    i32.shr_u)
  (func $_ZN8dlmalloc8dlmalloc5Chunk6pinuse17hde09ac7fe521978fE (type 13) (param i32) (result i32)
    local.get 0
    i32.load offset=4
    i32.const 1
    i32.and)
  (func $_ZN8dlmalloc8dlmalloc5Chunk12clear_pinuse17h47112d10bff2c398E (type 11) (param i32)
    local.get 0
    local.get 0
    i32.load offset=4
    i32.const -2
    i32.and
    i32.store offset=4)
  (func $_ZN8dlmalloc8dlmalloc5Chunk5inuse17h824c71bc0f61bb75E (type 13) (param i32) (result i32)
    local.get 0
    i32.load offset=4
    i32.const 3
    i32.and
    i32.const 1
    i32.ne)
  (func $_ZN8dlmalloc8dlmalloc5Chunk7mmapped17h69bc6c14adea384fE (type 13) (param i32) (result i32)
    local.get 0
    i32.load8_u offset=4
    i32.const 3
    i32.and
    i32.eqz)
  (func $_ZN8dlmalloc8dlmalloc5Chunk9set_inuse17h5f3bbfae424d3974E (type 0) (param i32 i32)
    local.get 0
    local.get 0
    i32.load offset=4
    i32.const 1
    i32.and
    local.get 1
    i32.or
    i32.const 2
    i32.or
    i32.store offset=4
    local.get 0
    local.get 1
    i32.add
    local.tee 0
    local.get 0
    i32.load offset=4
    i32.const 1
    i32.or
    i32.store offset=4)
  (func $_ZN8dlmalloc8dlmalloc5Chunk20set_inuse_and_pinuse17h134eb0365548205aE (type 0) (param i32 i32)
    local.get 0
    local.get 1
    i32.const 3
    i32.or
    i32.store offset=4
    local.get 0
    local.get 1
    i32.add
    local.tee 1
    local.get 1
    i32.load offset=4
    i32.const 1
    i32.or
    i32.store offset=4)
  (func $_ZN8dlmalloc8dlmalloc5Chunk34set_size_and_pinuse_of_inuse_chunk17h54212a19225d5a32E (type 0) (param i32 i32)
    local.get 0
    local.get 1
    i32.const 3
    i32.or
    i32.store offset=4)
  (func $_ZN8dlmalloc8dlmalloc5Chunk33set_size_and_pinuse_of_free_chunk17h6fb36837ea4ab16bE (type 0) (param i32 i32)
    local.get 0
    local.get 1
    i32.const 1
    i32.or
    i32.store offset=4
    local.get 0
    local.get 1
    i32.add
    local.get 1
    i32.store)
  (func $_ZN8dlmalloc8dlmalloc5Chunk20set_free_with_pinuse17h56042b51bf6983deE (type 3) (param i32 i32 i32)
    local.get 2
    local.get 2
    i32.load offset=4
    i32.const -2
    i32.and
    i32.store offset=4
    local.get 0
    local.get 1
    i32.const 1
    i32.or
    i32.store offset=4
    local.get 0
    local.get 1
    i32.add
    local.get 1
    i32.store)
  (func $_ZN8dlmalloc8dlmalloc5Chunk11plus_offset17h2d6162d80f46e59fE (type 2) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.add)
  (func $_ZN8dlmalloc8dlmalloc5Chunk12minus_offset17h1bc6fc1549394b74E (type 2) (param i32 i32) (result i32)
    local.get 0
    local.get 1
    i32.sub)
  (func $_ZN8dlmalloc8dlmalloc5Chunk6to_mem17had321cb1231dd337E (type 13) (param i32) (result i32)
    local.get 0
    i32.const 8
    i32.add)
  (func $_ZN8dlmalloc8dlmalloc5Chunk10mem_offset17h09955c191b66d08aE (type 14) (result i32)
    i32.const 8)
  (func $_ZN8dlmalloc8dlmalloc5Chunk8from_mem17hb5c408f151442e99E (type 13) (param i32) (result i32)
    local.get 0
    i32.const -8
    i32.add)
  (func $_ZN8dlmalloc8dlmalloc9TreeChunk14leftmost_child17he1deea80b772637aE (type 13) (param i32) (result i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      i32.load offset=16
      local.tee 1
      br_if 0 (;@1;)
      local.get 0
      i32.const 20
      i32.add
      i32.load
      local.set 1
    end
    local.get 1)
  (func $_ZN8dlmalloc8dlmalloc9TreeChunk5chunk17h7e6dfa11f2da074dE (type 13) (param i32) (result i32)
    local.get 0)
  (func $_ZN8dlmalloc8dlmalloc9TreeChunk4next17h395c3bed0a2e5517E (type 13) (param i32) (result i32)
    local.get 0
    i32.load offset=12)
  (func $_ZN8dlmalloc8dlmalloc9TreeChunk4prev17h67b7d7330692af94E (type 13) (param i32) (result i32)
    local.get 0
    i32.load offset=8)
  (func $_ZN8dlmalloc8dlmalloc7Segment9is_extern17hfc6d4e628d2b8313E (type 13) (param i32) (result i32)
    local.get 0
    i32.load offset=12
    i32.const 1
    i32.and)
  (func $_ZN8dlmalloc8dlmalloc7Segment9sys_flags17h1c26c3dd33e84b43E (type 13) (param i32) (result i32)
    local.get 0
    i32.load offset=12
    i32.const 1
    i32.shr_u)
  (func $_ZN8dlmalloc8dlmalloc7Segment5holds17h02477c6e3bcb4213E (type 2) (param i32 i32) (result i32)
    (local i32 i32)
    i32.const 0
    local.set 2
    block  ;; label = @1
      local.get 0
      i32.load
      local.tee 3
      local.get 1
      i32.gt_u
      br_if 0 (;@1;)
      local.get 3
      local.get 0
      i32.load offset=4
      i32.add
      local.get 1
      i32.gt_u
      local.set 2
    end
    local.get 2)
  (func $_ZN8dlmalloc8dlmalloc7Segment3top17h07497f7196b9a54fE (type 13) (param i32) (result i32)
    local.get 0
    i32.load
    local.get 0
    i32.load offset=4
    i32.add)
  (func $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$5alloc17hb11e962cef68a804E (type 3) (param i32 i32 i32)
    (local i32)
    local.get 2
    i32.const 16
    i32.shr_u
    memory.grow
    local.set 3
    local.get 0
    i32.const 0
    i32.store offset=8
    local.get 0
    i32.const 0
    local.get 2
    i32.const -65536
    i32.and
    local.get 3
    i32.const -1
    i32.eq
    local.tee 2
    select
    i32.store offset=4
    local.get 0
    i32.const 0
    local.get 3
    i32.const 16
    i32.shl
    local.get 2
    select
    i32.store)
  (func $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$5remap17h7f179872f7b3a976E (type 16) (param i32 i32 i32 i32 i32) (result i32)
    i32.const 0)
  (func $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$9free_part17h38a6e48730b8faaeE (type 9) (param i32 i32 i32 i32) (result i32)
    i32.const 0)
  (func $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$4free17h447c2248b31fe256E (type 1) (param i32 i32 i32) (result i32)
    i32.const 0)
  (func $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$16can_release_part17h79bb8fcc390ac09aE (type 2) (param i32 i32) (result i32)
    i32.const 0)
  (func $_ZN61_$LT$dlmalloc..sys..System$u20$as$u20$dlmalloc..Allocator$GT$9page_size17h1fc4fb2be482df9eE (type 13) (param i32) (result i32)
    i32.const 65536)
  (func $_ZN4core3ptr26drop_in_place$LT$usize$GT$17h052ff139fb0757f4E (type 11) (param i32))
  (func $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$10write_char17h0f937bd2ea63dfa4E (type 2) (param i32 i32) (result i32)
    local.get 0
    i32.load
    local.get 1
    call $_ZN5alloc6string6String4push17hf7eba259bb689f33E
    i32.const 0)
  (func $_ZN5alloc6string6String4push17hf7eba259bb689f33E (type 0) (param i32 i32)
    (local i32 i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.const 127
        i32.gt_u
        br_if 0 (;@2;)
        block  ;; label = @3
          local.get 0
          i32.load offset=8
          local.tee 3
          local.get 0
          i32.load
          i32.ne
          br_if 0 (;@3;)
          local.get 0
          local.get 3
          call $_ZN5alloc7raw_vec19RawVec$LT$T$C$A$GT$16reserve_for_push17hf03d0ee1756a2143E
          local.get 0
          i32.load offset=8
          local.set 3
        end
        local.get 0
        local.get 3
        i32.const 1
        i32.add
        i32.store offset=8
        local.get 0
        i32.load offset=4
        local.get 3
        i32.add
        local.get 1
        i32.store8
        br 1 (;@1;)
      end
      local.get 2
      i32.const 0
      i32.store offset=12
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.const 2048
          i32.lt_u
          br_if 0 (;@3;)
          block  ;; label = @4
            local.get 1
            i32.const 65536
            i32.lt_u
            br_if 0 (;@4;)
            local.get 2
            local.get 1
            i32.const 63
            i32.and
            i32.const 128
            i32.or
            i32.store8 offset=15
            local.get 2
            local.get 1
            i32.const 6
            i32.shr_u
            i32.const 63
            i32.and
            i32.const 128
            i32.or
            i32.store8 offset=14
            local.get 2
            local.get 1
            i32.const 12
            i32.shr_u
            i32.const 63
            i32.and
            i32.const 128
            i32.or
            i32.store8 offset=13
            local.get 2
            local.get 1
            i32.const 18
            i32.shr_u
            i32.const 7
            i32.and
            i32.const 240
            i32.or
            i32.store8 offset=12
            i32.const 4
            local.set 1
            br 2 (;@2;)
          end
          local.get 2
          local.get 1
          i32.const 63
          i32.and
          i32.const 128
          i32.or
          i32.store8 offset=14
          local.get 2
          local.get 1
          i32.const 12
          i32.shr_u
          i32.const 224
          i32.or
          i32.store8 offset=12
          local.get 2
          local.get 1
          i32.const 6
          i32.shr_u
          i32.const 63
          i32.and
          i32.const 128
          i32.or
          i32.store8 offset=13
          i32.const 3
          local.set 1
          br 1 (;@2;)
        end
        local.get 2
        local.get 1
        i32.const 63
        i32.and
        i32.const 128
        i32.or
        i32.store8 offset=13
        local.get 2
        local.get 1
        i32.const 6
        i32.shr_u
        i32.const 192
        i32.or
        i32.store8 offset=12
        i32.const 2
        local.set 1
      end
      block  ;; label = @2
        local.get 0
        i32.load
        local.get 0
        i32.load offset=8
        local.tee 3
        i32.sub
        local.get 1
        i32.ge_u
        br_if 0 (;@2;)
        local.get 0
        local.get 3
        local.get 1
        call $_ZN5alloc7raw_vec19RawVec$LT$T$C$A$GT$7reserve21do_reserve_and_handle17hb9fe56016b07faebE
        local.get 0
        i32.load offset=8
        local.set 3
      end
      local.get 0
      i32.load offset=4
      local.get 3
      i32.add
      local.get 2
      i32.const 12
      i32.add
      local.get 1
      call $memcpy
      drop
      local.get 0
      local.get 3
      local.get 1
      i32.add
      i32.store offset=8
    end
    local.get 2
    i32.const 16
    i32.add
    global.set $__stack_pointer)
  (func $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_fmt17h2389dd1df7b5adb2E (type 2) (param i32 i32) (result i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    local.get 0
    i32.load
    local.set 0
    local.get 2
    i32.const 8
    i32.add
    i32.const 16
    i32.add
    local.get 1
    i32.const 16
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    i32.const 8
    i32.add
    i32.const 8
    i32.add
    local.get 1
    i32.const 8
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    local.get 1
    i64.load align=4
    i64.store offset=8
    local.get 2
    local.get 0
    i32.store offset=4
    local.get 2
    i32.const 4
    i32.add
    i32.const 1048928
    local.get 2
    i32.const 8
    i32.add
    call $_ZN4core3fmt5write17h7d7e1e065b0161edE
    local.set 1
    local.get 2
    i32.const 32
    i32.add
    global.set $__stack_pointer
    local.get 1)
  (func $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_str17hd86cd7fd76c297c4E (type 1) (param i32 i32 i32) (result i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      i32.load
      local.tee 0
      i32.load
      local.get 0
      i32.load offset=8
      local.tee 3
      i32.sub
      local.get 2
      i32.ge_u
      br_if 0 (;@1;)
      local.get 0
      local.get 3
      local.get 2
      call $_ZN5alloc7raw_vec19RawVec$LT$T$C$A$GT$7reserve21do_reserve_and_handle17hb9fe56016b07faebE
      local.get 0
      i32.load offset=8
      local.set 3
    end
    local.get 0
    i32.load offset=4
    local.get 3
    i32.add
    local.get 1
    local.get 2
    call $memcpy
    drop
    local.get 0
    local.get 3
    local.get 2
    i32.add
    i32.store offset=8
    i32.const 0)
  (func $_ZN5alloc7raw_vec19RawVec$LT$T$C$A$GT$7reserve21do_reserve_and_handle17hb9fe56016b07faebE (type 3) (param i32 i32 i32)
    (local i32 i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        local.get 2
        i32.add
        local.tee 2
        local.get 1
        i32.lt_u
        br_if 0 (;@2;)
        local.get 0
        i32.load
        local.tee 1
        i32.const 1
        i32.shl
        local.tee 4
        local.get 2
        local.get 4
        local.get 2
        i32.gt_u
        select
        local.tee 2
        i32.const 8
        local.get 2
        i32.const 8
        i32.gt_u
        select
        local.tee 2
        i32.const -1
        i32.xor
        i32.const 31
        i32.shr_u
        local.set 4
        block  ;; label = @3
          block  ;; label = @4
            local.get 1
            i32.eqz
            br_if 0 (;@4;)
            local.get 3
            i32.const 1
            i32.store offset=24
            local.get 3
            local.get 1
            i32.store offset=20
            local.get 3
            local.get 0
            i32.const 4
            i32.add
            i32.load
            i32.store offset=16
            br 1 (;@3;)
          end
          local.get 3
          i32.const 0
          i32.store offset=24
        end
        local.get 3
        local.get 2
        local.get 4
        local.get 3
        i32.const 16
        i32.add
        call $_ZN5alloc7raw_vec11finish_grow17h97427dc56ae39043E
        local.get 3
        i32.load offset=4
        local.set 1
        block  ;; label = @3
          local.get 3
          i32.load
          br_if 0 (;@3;)
          local.get 0
          local.get 2
          i32.store
          local.get 0
          local.get 1
          i32.store offset=4
          br 2 (;@1;)
        end
        local.get 3
        i32.const 8
        i32.add
        i32.load
        local.tee 0
        i32.const -2147483647
        i32.eq
        br_if 1 (;@1;)
        local.get 0
        i32.eqz
        br_if 0 (;@2;)
        local.get 1
        local.get 0
        call $_ZN5alloc5alloc18handle_alloc_error17hc1ba5222dd907d5cE
        unreachable
      end
      call $_ZN5alloc7raw_vec17capacity_overflow17h25579b3a3c90486eE
      unreachable
    end
    local.get 3
    i32.const 32
    i32.add
    global.set $__stack_pointer)
  (func $_ZN5alloc7raw_vec11finish_grow17h97427dc56ae39043E (type 12) (param i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.eqz
        br_if 0 (;@2;)
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 1
                    i32.const -1
                    i32.le_s
                    br_if 0 (;@8;)
                    local.get 3
                    i32.load offset=8
                    br_if 1 (;@7;)
                    local.get 1
                    br_if 2 (;@6;)
                    i32.const 1
                    local.set 2
                    br 4 (;@4;)
                  end
                  local.get 0
                  i32.const 8
                  i32.add
                  i32.const 0
                  i32.store
                  br 6 (;@1;)
                end
                block  ;; label = @7
                  local.get 3
                  i32.load offset=4
                  local.tee 2
                  br_if 0 (;@7;)
                  block  ;; label = @8
                    local.get 1
                    br_if 0 (;@8;)
                    i32.const 1
                    local.set 2
                    br 4 (;@4;)
                  end
                  local.get 1
                  i32.const 1
                  call $__rust_alloc
                  local.set 2
                  br 2 (;@5;)
                end
                local.get 3
                i32.load
                local.get 2
                i32.const 1
                local.get 1
                call $__rust_realloc
                local.set 2
                br 1 (;@5;)
              end
              local.get 1
              i32.const 1
              call $__rust_alloc
              local.set 2
            end
            local.get 2
            i32.eqz
            br_if 1 (;@3;)
          end
          local.get 0
          local.get 2
          i32.store offset=4
          local.get 0
          i32.const 8
          i32.add
          local.get 1
          i32.store
          local.get 0
          i32.const 0
          i32.store
          return
        end
        local.get 0
        local.get 1
        i32.store offset=4
        local.get 0
        i32.const 8
        i32.add
        i32.const 1
        i32.store
        local.get 0
        i32.const 1
        i32.store
        return
      end
      local.get 0
      local.get 1
      i32.store offset=4
      local.get 0
      i32.const 8
      i32.add
      i32.const 0
      i32.store
    end
    local.get 0
    i32.const 1
    i32.store)
  (func $_ZN5alloc5alloc18handle_alloc_error17hc1ba5222dd907d5cE (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call $_ZN5alloc5alloc18handle_alloc_error8rt_error17h318c45b3aa9d998fE
    unreachable)
  (func $_ZN5alloc7raw_vec17capacity_overflow17h25579b3a3c90486eE (type 17)
    (local i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 0
    global.set $__stack_pointer
    local.get 0
    i32.const 20
    i32.add
    i32.const 1
    i32.store
    local.get 0
    i32.const 28
    i32.add
    i32.const 0
    i32.store
    local.get 0
    i32.const 1049000
    i32.store offset=16
    local.get 0
    i32.const 1048952
    i32.store offset=24
    local.get 0
    i32.const 0
    i32.store offset=8
    local.get 0
    i32.const 8
    i32.add
    i32.const 1049008
    call $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E
    unreachable)
  (func $_ZN5alloc7raw_vec19RawVec$LT$T$C$A$GT$16reserve_for_push17hf03d0ee1756a2143E (type 0) (param i32 i32)
    (local i32 i32 i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.const 1
        i32.add
        local.tee 1
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        i32.load
        local.tee 3
        i32.const 1
        i32.shl
        local.tee 4
        local.get 1
        local.get 4
        local.get 1
        i32.gt_u
        select
        local.tee 1
        i32.const 8
        local.get 1
        i32.const 8
        i32.gt_u
        select
        local.tee 1
        i32.const -1
        i32.xor
        i32.const 31
        i32.shr_u
        local.set 4
        block  ;; label = @3
          block  ;; label = @4
            local.get 3
            i32.eqz
            br_if 0 (;@4;)
            local.get 2
            i32.const 1
            i32.store offset=24
            local.get 2
            local.get 3
            i32.store offset=20
            local.get 2
            local.get 0
            i32.const 4
            i32.add
            i32.load
            i32.store offset=16
            br 1 (;@3;)
          end
          local.get 2
          i32.const 0
          i32.store offset=24
        end
        local.get 2
        local.get 1
        local.get 4
        local.get 2
        i32.const 16
        i32.add
        call $_ZN5alloc7raw_vec11finish_grow17h97427dc56ae39043E
        local.get 2
        i32.load offset=4
        local.set 3
        block  ;; label = @3
          local.get 2
          i32.load
          br_if 0 (;@3;)
          local.get 0
          local.get 1
          i32.store
          local.get 0
          local.get 3
          i32.store offset=4
          br 2 (;@1;)
        end
        local.get 2
        i32.const 8
        i32.add
        i32.load
        local.tee 0
        i32.const -2147483647
        i32.eq
        br_if 1 (;@1;)
        local.get 0
        i32.eqz
        br_if 0 (;@2;)
        local.get 3
        local.get 0
        call $_ZN5alloc5alloc18handle_alloc_error17hc1ba5222dd907d5cE
        unreachable
      end
      call $_ZN5alloc7raw_vec17capacity_overflow17h25579b3a3c90486eE
      unreachable
    end
    local.get 2
    i32.const 32
    i32.add
    global.set $__stack_pointer)
  (func $_ZN5alloc5alloc18handle_alloc_error8rt_error17h318c45b3aa9d998fE (type 0) (param i32 i32)
    local.get 0
    local.get 1
    call $__rust_alloc_error_handler
    unreachable)
  (func $_ZN5alloc3fmt6format12format_inner17hf9dd33246b0be2a8E (type 0) (param i32 i32)
    (local i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.const 12
        i32.add
        i32.load
        local.tee 3
        i32.eqz
        br_if 0 (;@2;)
        local.get 1
        i32.load offset=8
        local.set 4
        local.get 3
        i32.const -1
        i32.add
        i32.const 536870911
        i32.and
        local.tee 3
        i32.const 1
        i32.add
        local.tee 5
        i32.const 7
        i32.and
        local.set 6
        block  ;; label = @3
          block  ;; label = @4
            local.get 3
            i32.const 7
            i32.ge_u
            br_if 0 (;@4;)
            i32.const 0
            local.set 5
            local.get 4
            local.set 3
            br 1 (;@3;)
          end
          local.get 4
          i32.const 60
          i32.add
          local.set 3
          local.get 5
          i32.const 1073741816
          i32.and
          local.set 7
          i32.const 0
          local.set 5
          loop  ;; label = @4
            local.get 3
            i32.load
            local.get 3
            i32.const -8
            i32.add
            i32.load
            local.get 3
            i32.const -16
            i32.add
            i32.load
            local.get 3
            i32.const -24
            i32.add
            i32.load
            local.get 3
            i32.const -32
            i32.add
            i32.load
            local.get 3
            i32.const -40
            i32.add
            i32.load
            local.get 3
            i32.const -48
            i32.add
            i32.load
            local.get 3
            i32.const -56
            i32.add
            i32.load
            local.get 5
            i32.add
            i32.add
            i32.add
            i32.add
            i32.add
            i32.add
            i32.add
            i32.add
            local.set 5
            local.get 3
            i32.const 64
            i32.add
            local.set 3
            local.get 7
            i32.const -8
            i32.add
            local.tee 7
            br_if 0 (;@4;)
          end
          local.get 3
          i32.const -60
          i32.add
          local.set 3
        end
        block  ;; label = @3
          local.get 6
          i32.eqz
          br_if 0 (;@3;)
          local.get 3
          i32.const 4
          i32.add
          local.set 3
          loop  ;; label = @4
            local.get 3
            i32.load
            local.get 5
            i32.add
            local.set 5
            local.get 3
            i32.const 8
            i32.add
            local.set 3
            local.get 6
            i32.const -1
            i32.add
            local.tee 6
            br_if 0 (;@4;)
          end
        end
        block  ;; label = @3
          local.get 1
          i32.const 20
          i32.add
          i32.load
          i32.eqz
          br_if 0 (;@3;)
          i32.const 1
          local.set 3
          i32.const 0
          local.set 6
          local.get 5
          i32.const 0
          i32.lt_s
          br_if 2 (;@1;)
          local.get 5
          i32.const 16
          i32.lt_u
          local.get 4
          i32.load offset=4
          i32.eqz
          i32.and
          br_if 2 (;@1;)
          local.get 5
          i32.const 1
          i32.shl
          local.set 5
        end
        local.get 5
        i32.eqz
        br_if 0 (;@2;)
        block  ;; label = @3
          block  ;; label = @4
            local.get 5
            i32.const -1
            i32.le_s
            br_if 0 (;@4;)
            local.get 5
            i32.const 1
            call $__rust_alloc
            local.tee 3
            i32.eqz
            br_if 1 (;@3;)
            local.get 5
            local.set 6
            br 3 (;@1;)
          end
          call $_ZN5alloc7raw_vec17capacity_overflow17h25579b3a3c90486eE
          unreachable
        end
        local.get 5
        i32.const 1
        call $_ZN5alloc5alloc18handle_alloc_error17hc1ba5222dd907d5cE
        unreachable
      end
      i32.const 1
      local.set 3
      i32.const 0
      local.set 6
    end
    local.get 0
    i32.const 0
    i32.store offset=8
    local.get 0
    local.get 3
    i32.store offset=4
    local.get 0
    local.get 6
    i32.store
    local.get 2
    i32.const 16
    i32.add
    i32.const 16
    i32.add
    local.get 1
    i32.const 16
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    i32.const 16
    i32.add
    i32.const 8
    i32.add
    local.get 1
    i32.const 8
    i32.add
    i64.load align=4
    i64.store
    local.get 2
    local.get 1
    i64.load align=4
    i64.store offset=16
    local.get 2
    local.get 0
    i32.store offset=12
    block  ;; label = @1
      local.get 2
      i32.const 12
      i32.add
      i32.const 1048928
      local.get 2
      i32.const 16
      i32.add
      call $_ZN4core3fmt5write17h7d7e1e065b0161edE
      i32.eqz
      br_if 0 (;@1;)
      i32.const 1049024
      i32.const 51
      local.get 2
      i32.const 40
      i32.add
      i32.const 1049076
      i32.const 1049116
      call $_ZN4core6result13unwrap_failed17h80851617ed2e05abE
      unreachable
    end
    local.get 2
    i32.const 48
    i32.add
    global.set $__stack_pointer)
  (func $_ZN60_$LT$alloc..string..String$u20$as$u20$core..clone..Clone$GT$5clone17h48fea206591b489eE (type 0) (param i32 i32)
    (local i32 i32)
    local.get 1
    i32.const 4
    i32.add
    i32.load
    local.set 2
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 1
            i32.const 8
            i32.add
            i32.load
            local.tee 1
            br_if 0 (;@4;)
            i32.const 1
            local.set 3
            br 1 (;@3;)
          end
          local.get 1
          i32.const -1
          i32.le_s
          br_if 1 (;@2;)
          local.get 1
          i32.const 1
          call $__rust_alloc
          local.tee 3
          i32.eqz
          br_if 2 (;@1;)
        end
        local.get 3
        local.get 2
        local.get 1
        call $memcpy
        local.set 2
        local.get 0
        local.get 1
        i32.store offset=8
        local.get 0
        local.get 2
        i32.store offset=4
        local.get 0
        local.get 1
        i32.store
        return
      end
      call $_ZN5alloc7raw_vec17capacity_overflow17h25579b3a3c90486eE
      unreachable
    end
    local.get 1
    i32.const 1
    call $_ZN5alloc5alloc18handle_alloc_error17hc1ba5222dd907d5cE
    unreachable)
  (func $_ZN4core3ops8function6FnOnce9call_once17h936625e43bb00ae7E (type 2) (param i32 i32) (result i32)
    local.get 0
    i32.load
    drop
    loop (result i32)  ;; label = @1
      br 0 (;@1;)
    end)
  (func $_ZN4core3ptr102drop_in_place$LT$$RF$core..iter..adapters..copied..Copied$LT$core..slice..iter..Iter$LT$u8$GT$$GT$$GT$17h4bf2435251908228E (type 11) (param i32))
  (func $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E (type 0) (param i32 i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    local.get 2
    i32.const 1
    i32.store8 offset=24
    local.get 2
    local.get 1
    i32.store offset=20
    local.get 2
    local.get 0
    i32.store offset=16
    local.get 2
    i32.const 1052248
    i32.store offset=12
    local.get 2
    i32.const 1052176
    i32.store offset=8
    local.get 2
    i32.const 8
    i32.add
    call $rust_begin_unwind
    unreachable)
  (func $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E (type 3) (param i32 i32 i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    local.get 3
    local.get 1
    i32.store offset=4
    local.get 3
    local.get 0
    i32.store
    local.get 3
    i32.const 8
    i32.add
    i32.const 12
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 28
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 32
    i32.add
    i32.const 12
    i32.add
    i32.const 4
    i32.store
    local.get 3
    i32.const 1052316
    i32.store offset=16
    local.get 3
    i32.const 0
    i32.store offset=8
    local.get 3
    i32.const 4
    i32.store offset=36
    local.get 3
    local.get 3
    i32.const 32
    i32.add
    i32.store offset=24
    local.get 3
    local.get 3
    i32.store offset=40
    local.get 3
    local.get 3
    i32.const 4
    i32.add
    i32.store offset=32
    local.get 3
    i32.const 8
    i32.add
    local.get 2
    call $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E
    unreachable)
  (func $_ZN4core5slice5index26slice_start_index_len_fail17h61955cceb128f310E (type 3) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call $_ZN4core5slice5index29slice_start_index_len_fail_rt17he819c5908e9fdb1fE
    unreachable)
  (func $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E (type 3) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call $_ZN4core5slice5index27slice_end_index_len_fail_rt17h480eb8e5cf49cf70E
    unreachable)
  (func $_ZN4core3fmt9Formatter3pad17h7e3d356fb87d6820E (type 1) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i32.load offset=8
        local.tee 3
        local.get 0
        i32.load offset=16
        local.tee 4
        i32.or
        i32.eqz
        br_if 0 (;@2;)
        block  ;; label = @3
          local.get 4
          i32.eqz
          br_if 0 (;@3;)
          local.get 1
          local.get 2
          i32.add
          local.set 5
          local.get 0
          i32.const 20
          i32.add
          i32.load
          i32.const 1
          i32.add
          local.set 6
          i32.const 0
          local.set 7
          local.get 1
          local.set 8
          block  ;; label = @4
            loop  ;; label = @5
              local.get 8
              local.set 4
              local.get 6
              i32.const -1
              i32.add
              local.tee 6
              i32.eqz
              br_if 1 (;@4;)
              local.get 4
              local.get 5
              i32.eq
              br_if 2 (;@3;)
              block  ;; label = @6
                block  ;; label = @7
                  local.get 4
                  i32.load8_s
                  local.tee 9
                  i32.const -1
                  i32.le_s
                  br_if 0 (;@7;)
                  local.get 4
                  i32.const 1
                  i32.add
                  local.set 8
                  local.get 9
                  i32.const 255
                  i32.and
                  local.set 9
                  br 1 (;@6;)
                end
                local.get 4
                i32.load8_u offset=1
                i32.const 63
                i32.and
                local.set 10
                local.get 9
                i32.const 31
                i32.and
                local.set 8
                block  ;; label = @7
                  local.get 9
                  i32.const -33
                  i32.gt_u
                  br_if 0 (;@7;)
                  local.get 8
                  i32.const 6
                  i32.shl
                  local.get 10
                  i32.or
                  local.set 9
                  local.get 4
                  i32.const 2
                  i32.add
                  local.set 8
                  br 1 (;@6;)
                end
                local.get 10
                i32.const 6
                i32.shl
                local.get 4
                i32.load8_u offset=2
                i32.const 63
                i32.and
                i32.or
                local.set 10
                block  ;; label = @7
                  local.get 9
                  i32.const -16
                  i32.ge_u
                  br_if 0 (;@7;)
                  local.get 10
                  local.get 8
                  i32.const 12
                  i32.shl
                  i32.or
                  local.set 9
                  local.get 4
                  i32.const 3
                  i32.add
                  local.set 8
                  br 1 (;@6;)
                end
                local.get 10
                i32.const 6
                i32.shl
                local.get 4
                i32.load8_u offset=3
                i32.const 63
                i32.and
                i32.or
                local.get 8
                i32.const 18
                i32.shl
                i32.const 1835008
                i32.and
                i32.or
                local.tee 9
                i32.const 1114112
                i32.eq
                br_if 3 (;@3;)
                local.get 4
                i32.const 4
                i32.add
                local.set 8
              end
              local.get 7
              local.get 4
              i32.sub
              local.get 8
              i32.add
              local.set 7
              local.get 9
              i32.const 1114112
              i32.ne
              br_if 0 (;@5;)
              br 2 (;@3;)
            end
          end
          local.get 4
          local.get 5
          i32.eq
          br_if 0 (;@3;)
          block  ;; label = @4
            local.get 4
            i32.load8_s
            local.tee 8
            i32.const -1
            i32.gt_s
            br_if 0 (;@4;)
            local.get 8
            i32.const -32
            i32.lt_u
            br_if 0 (;@4;)
            local.get 8
            i32.const -16
            i32.lt_u
            br_if 0 (;@4;)
            local.get 4
            i32.load8_u offset=2
            i32.const 63
            i32.and
            i32.const 6
            i32.shl
            local.get 4
            i32.load8_u offset=1
            i32.const 63
            i32.and
            i32.const 12
            i32.shl
            i32.or
            local.get 4
            i32.load8_u offset=3
            i32.const 63
            i32.and
            i32.or
            local.get 8
            i32.const 255
            i32.and
            i32.const 18
            i32.shl
            i32.const 1835008
            i32.and
            i32.or
            i32.const 1114112
            i32.eq
            br_if 1 (;@3;)
          end
          block  ;; label = @4
            block  ;; label = @5
              local.get 7
              i32.eqz
              br_if 0 (;@5;)
              block  ;; label = @6
                local.get 7
                local.get 2
                i32.lt_u
                br_if 0 (;@6;)
                i32.const 0
                local.set 4
                local.get 7
                local.get 2
                i32.eq
                br_if 1 (;@5;)
                br 2 (;@4;)
              end
              i32.const 0
              local.set 4
              local.get 1
              local.get 7
              i32.add
              i32.load8_s
              i32.const -64
              i32.lt_s
              br_if 1 (;@4;)
            end
            local.get 1
            local.set 4
          end
          local.get 7
          local.get 2
          local.get 4
          select
          local.set 2
          local.get 4
          local.get 1
          local.get 4
          select
          local.set 1
        end
        block  ;; label = @3
          local.get 3
          br_if 0 (;@3;)
          local.get 0
          i32.load
          local.get 1
          local.get 2
          local.get 0
          i32.load offset=4
          i32.load offset=12
          call_indirect (type 1)
          return
        end
        local.get 0
        i32.const 12
        i32.add
        i32.load
        local.set 7
        block  ;; label = @3
          block  ;; label = @4
            local.get 2
            i32.const 16
            i32.lt_u
            br_if 0 (;@4;)
            local.get 1
            local.get 2
            call $_ZN4core3str5count14do_count_chars17hffbd23dea83a0c54E
            local.set 8
            br 1 (;@3;)
          end
          block  ;; label = @4
            local.get 2
            br_if 0 (;@4;)
            i32.const 0
            local.set 8
            br 1 (;@3;)
          end
          local.get 2
          i32.const 3
          i32.and
          local.set 9
          block  ;; label = @4
            block  ;; label = @5
              local.get 2
              i32.const 4
              i32.ge_u
              br_if 0 (;@5;)
              i32.const 0
              local.set 8
              local.get 1
              local.set 4
              br 1 (;@4;)
            end
            local.get 2
            i32.const -4
            i32.and
            local.set 6
            i32.const 0
            local.set 8
            local.get 1
            local.set 4
            loop  ;; label = @5
              local.get 8
              local.get 4
              i32.load8_s
              i32.const -65
              i32.gt_s
              i32.add
              local.get 4
              i32.load8_s offset=1
              i32.const -65
              i32.gt_s
              i32.add
              local.get 4
              i32.load8_s offset=2
              i32.const -65
              i32.gt_s
              i32.add
              local.get 4
              i32.load8_s offset=3
              i32.const -65
              i32.gt_s
              i32.add
              local.set 8
              local.get 4
              i32.const 4
              i32.add
              local.set 4
              local.get 6
              i32.const -4
              i32.add
              local.tee 6
              br_if 0 (;@5;)
            end
          end
          local.get 9
          i32.eqz
          br_if 0 (;@3;)
          loop  ;; label = @4
            local.get 8
            local.get 4
            i32.load8_s
            i32.const -65
            i32.gt_s
            i32.add
            local.set 8
            local.get 4
            i32.const 1
            i32.add
            local.set 4
            local.get 9
            i32.const -1
            i32.add
            local.tee 9
            br_if 0 (;@4;)
          end
        end
        block  ;; label = @3
          local.get 7
          local.get 8
          i32.le_u
          br_if 0 (;@3;)
          local.get 7
          local.get 8
          i32.sub
          local.tee 8
          local.set 7
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                i32.const 0
                local.get 0
                i32.load8_u offset=32
                local.tee 4
                local.get 4
                i32.const 3
                i32.eq
                select
                local.tee 4
                br_table 2 (;@4;) 0 (;@6;) 1 (;@5;) 2 (;@4;)
              end
              i32.const 0
              local.set 7
              local.get 8
              local.set 4
              br 1 (;@4;)
            end
            local.get 8
            i32.const 1
            i32.shr_u
            local.set 4
            local.get 8
            i32.const 1
            i32.add
            i32.const 1
            i32.shr_u
            local.set 7
          end
          local.get 4
          i32.const 1
          i32.add
          local.set 4
          local.get 0
          i32.const 4
          i32.add
          i32.load
          local.set 9
          local.get 0
          i32.load offset=28
          local.set 8
          local.get 0
          i32.load
          local.set 6
          block  ;; label = @4
            loop  ;; label = @5
              local.get 4
              i32.const -1
              i32.add
              local.tee 4
              i32.eqz
              br_if 1 (;@4;)
              local.get 6
              local.get 8
              local.get 9
              i32.load offset=16
              call_indirect (type 2)
              i32.eqz
              br_if 0 (;@5;)
            end
            i32.const 1
            return
          end
          i32.const 1
          local.set 4
          local.get 8
          i32.const 1114112
          i32.eq
          br_if 2 (;@1;)
          local.get 6
          local.get 1
          local.get 2
          local.get 9
          i32.load offset=12
          call_indirect (type 1)
          br_if 2 (;@1;)
          i32.const 0
          local.set 4
          loop  ;; label = @4
            block  ;; label = @5
              local.get 7
              local.get 4
              i32.ne
              br_if 0 (;@5;)
              local.get 7
              local.get 7
              i32.lt_u
              return
            end
            local.get 4
            i32.const 1
            i32.add
            local.set 4
            local.get 6
            local.get 8
            local.get 9
            i32.load offset=16
            call_indirect (type 2)
            i32.eqz
            br_if 0 (;@4;)
          end
          local.get 4
          i32.const -1
          i32.add
          local.get 7
          i32.lt_u
          return
        end
        local.get 0
        i32.load
        local.get 1
        local.get 2
        local.get 0
        i32.load offset=4
        i32.load offset=12
        call_indirect (type 1)
        return
      end
      local.get 0
      i32.load
      local.get 1
      local.get 2
      local.get 0
      i32.load offset=4
      i32.load offset=12
      call_indirect (type 1)
      local.set 4
    end
    local.get 4)
  (func $_ZN4core9panicking13assert_failed17hca637033118b5d2eE (type 3) (param i32 i32 i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    local.get 3
    local.get 1
    i32.store offset=4
    local.get 3
    local.get 0
    i32.store
    local.get 3
    i32.const 8
    i32.add
    i32.const 16
    i32.add
    local.get 2
    i32.const 16
    i32.add
    i64.load align=4
    i64.store
    local.get 3
    i32.const 8
    i32.add
    i32.const 8
    i32.add
    local.get 2
    i32.const 8
    i32.add
    i64.load align=4
    i64.store
    local.get 3
    local.get 2
    i64.load align=4
    i64.store offset=8
    i32.const 0
    local.get 3
    i32.const 1052332
    local.get 3
    i32.const 4
    i32.add
    i32.const 1052332
    local.get 3
    i32.const 8
    i32.add
    i32.const 1049216
    call $_ZN4core9panicking19assert_failed_inner17h6bb952400e68e912E
    unreachable)
  (func $_ZN4core9panicking5panic17he2fadc50abd7e245E (type 3) (param i32 i32 i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    local.get 3
    i32.const 12
    i32.add
    i32.const 1
    i32.store
    local.get 3
    i32.const 20
    i32.add
    i32.const 0
    i32.store
    local.get 3
    i32.const 1052176
    i32.store offset=16
    local.get 3
    i32.const 0
    i32.store
    local.get 3
    local.get 1
    i32.store offset=28
    local.get 3
    local.get 0
    i32.store offset=24
    local.get 3
    local.get 3
    i32.const 24
    i32.add
    i32.store offset=8
    local.get 3
    local.get 2
    call $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E
    unreachable)
  (func $_ZN4core3num7flt2dec8strategy6dragon9mul_pow1017h3a8b434637fe22e2E (type 2) (param i32 i32) (result i32)
    (local i32 i32 i64 i32 i32 i64 i32)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 1
                  i32.const 7
                  i32.and
                  local.tee 2
                  i32.eqz
                  br_if 0 (;@7;)
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        local.get 0
                        i32.load offset=160
                        local.tee 3
                        i32.const 41
                        i32.ge_u
                        br_if 0 (;@10;)
                        block  ;; label = @11
                          local.get 3
                          br_if 0 (;@11;)
                          i32.const 0
                          local.set 3
                          br 3 (;@8;)
                        end
                        local.get 2
                        i32.const 2
                        i32.shl
                        i32.const 1049232
                        i32.add
                        i64.load32_u
                        local.set 4
                        local.get 3
                        i32.const -1
                        i32.add
                        i32.const 1073741823
                        i32.and
                        local.tee 2
                        i32.const 1
                        i32.add
                        local.tee 5
                        i32.const 3
                        i32.and
                        local.set 6
                        block  ;; label = @11
                          local.get 2
                          i32.const 3
                          i32.ge_u
                          br_if 0 (;@11;)
                          i64.const 0
                          local.set 7
                          local.get 0
                          local.set 2
                          br 2 (;@9;)
                        end
                        local.get 5
                        i32.const 2147483644
                        i32.and
                        local.set 5
                        i64.const 0
                        local.set 7
                        local.get 0
                        local.set 2
                        loop  ;; label = @11
                          local.get 2
                          local.get 2
                          i64.load32_u
                          local.get 4
                          i64.mul
                          local.get 7
                          i64.add
                          local.tee 7
                          i64.store32
                          local.get 2
                          i32.const 4
                          i32.add
                          local.tee 8
                          local.get 8
                          i64.load32_u
                          local.get 4
                          i64.mul
                          local.get 7
                          i64.const 32
                          i64.shr_u
                          i64.add
                          local.tee 7
                          i64.store32
                          local.get 2
                          i32.const 8
                          i32.add
                          local.tee 8
                          local.get 8
                          i64.load32_u
                          local.get 4
                          i64.mul
                          local.get 7
                          i64.const 32
                          i64.shr_u
                          i64.add
                          local.tee 7
                          i64.store32
                          local.get 2
                          i32.const 12
                          i32.add
                          local.tee 8
                          local.get 8
                          i64.load32_u
                          local.get 4
                          i64.mul
                          local.get 7
                          i64.const 32
                          i64.shr_u
                          i64.add
                          local.tee 7
                          i64.store32
                          local.get 7
                          i64.const 32
                          i64.shr_u
                          local.set 7
                          local.get 2
                          i32.const 16
                          i32.add
                          local.set 2
                          local.get 5
                          i32.const -4
                          i32.add
                          local.tee 5
                          br_if 0 (;@11;)
                          br 2 (;@9;)
                        end
                      end
                      local.get 3
                      i32.const 40
                      i32.const 1054848
                      call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                      unreachable
                    end
                    block  ;; label = @9
                      local.get 6
                      i32.eqz
                      br_if 0 (;@9;)
                      loop  ;; label = @10
                        local.get 2
                        local.get 2
                        i64.load32_u
                        local.get 4
                        i64.mul
                        local.get 7
                        i64.add
                        local.tee 7
                        i64.store32
                        local.get 2
                        i32.const 4
                        i32.add
                        local.set 2
                        local.get 7
                        i64.const 32
                        i64.shr_u
                        local.set 7
                        local.get 6
                        i32.const -1
                        i32.add
                        local.tee 6
                        br_if 0 (;@10;)
                      end
                    end
                    local.get 7
                    i32.wrap_i64
                    local.tee 2
                    i32.eqz
                    br_if 0 (;@8;)
                    local.get 3
                    i32.const 39
                    i32.gt_u
                    br_if 2 (;@6;)
                    local.get 0
                    local.get 3
                    i32.const 2
                    i32.shl
                    i32.add
                    local.get 2
                    i32.store
                    local.get 3
                    i32.const 1
                    i32.add
                    local.set 3
                  end
                  local.get 0
                  local.get 3
                  i32.store offset=160
                end
                local.get 1
                i32.const 8
                i32.and
                i32.eqz
                br_if 4 (;@2;)
                local.get 0
                i32.load offset=160
                local.tee 3
                i32.const 41
                i32.ge_u
                br_if 1 (;@5;)
                block  ;; label = @7
                  local.get 3
                  br_if 0 (;@7;)
                  i32.const 0
                  local.set 3
                  br 4 (;@3;)
                end
                local.get 3
                i32.const -1
                i32.add
                i32.const 1073741823
                i32.and
                local.tee 2
                i32.const 1
                i32.add
                local.tee 5
                i32.const 3
                i32.and
                local.set 6
                block  ;; label = @7
                  local.get 2
                  i32.const 3
                  i32.ge_u
                  br_if 0 (;@7;)
                  i64.const 0
                  local.set 4
                  local.get 0
                  local.set 2
                  br 3 (;@4;)
                end
                local.get 5
                i32.const 2147483644
                i32.and
                local.set 5
                i64.const 0
                local.set 4
                local.get 0
                local.set 2
                loop  ;; label = @7
                  local.get 2
                  local.get 2
                  i64.load32_u
                  i64.const 100000000
                  i64.mul
                  local.get 4
                  i64.add
                  local.tee 4
                  i64.store32
                  local.get 2
                  i32.const 4
                  i32.add
                  local.tee 8
                  local.get 8
                  i64.load32_u
                  i64.const 100000000
                  i64.mul
                  local.get 4
                  i64.const 32
                  i64.shr_u
                  i64.add
                  local.tee 4
                  i64.store32
                  local.get 2
                  i32.const 8
                  i32.add
                  local.tee 8
                  local.get 8
                  i64.load32_u
                  i64.const 100000000
                  i64.mul
                  local.get 4
                  i64.const 32
                  i64.shr_u
                  i64.add
                  local.tee 4
                  i64.store32
                  local.get 2
                  i32.const 12
                  i32.add
                  local.tee 8
                  local.get 8
                  i64.load32_u
                  i64.const 100000000
                  i64.mul
                  local.get 4
                  i64.const 32
                  i64.shr_u
                  i64.add
                  local.tee 4
                  i64.store32
                  local.get 4
                  i64.const 32
                  i64.shr_u
                  local.set 4
                  local.get 2
                  i32.const 16
                  i32.add
                  local.set 2
                  local.get 5
                  i32.const -4
                  i32.add
                  local.tee 5
                  br_if 0 (;@7;)
                  br 3 (;@4;)
                end
              end
              local.get 3
              i32.const 40
              i32.const 1054848
              call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
              unreachable
            end
            local.get 3
            i32.const 40
            i32.const 1054848
            call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
            unreachable
          end
          block  ;; label = @4
            local.get 6
            i32.eqz
            br_if 0 (;@4;)
            loop  ;; label = @5
              local.get 2
              local.get 2
              i64.load32_u
              i64.const 100000000
              i64.mul
              local.get 4
              i64.add
              local.tee 4
              i64.store32
              local.get 2
              i32.const 4
              i32.add
              local.set 2
              local.get 4
              i64.const 32
              i64.shr_u
              local.set 4
              local.get 6
              i32.const -1
              i32.add
              local.tee 6
              br_if 0 (;@5;)
            end
          end
          local.get 4
          i32.wrap_i64
          local.tee 2
          i32.eqz
          br_if 0 (;@3;)
          local.get 3
          i32.const 39
          i32.gt_u
          br_if 2 (;@1;)
          local.get 0
          local.get 3
          i32.const 2
          i32.shl
          i32.add
          local.get 2
          i32.store
          local.get 3
          i32.const 1
          i32.add
          local.set 3
        end
        local.get 0
        local.get 3
        i32.store offset=160
      end
      block  ;; label = @2
        local.get 1
        i32.const 16
        i32.and
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        i32.const 1049312
        i32.const 2
        call $_ZN4core3num6bignum8Big32x4010mul_digits17h90fa0468b79b6f4cE
        drop
      end
      block  ;; label = @2
        local.get 1
        i32.const 32
        i32.and
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        i32.const 1049320
        i32.const 4
        call $_ZN4core3num6bignum8Big32x4010mul_digits17h90fa0468b79b6f4cE
        drop
      end
      block  ;; label = @2
        local.get 1
        i32.const 64
        i32.and
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        i32.const 1049336
        i32.const 7
        call $_ZN4core3num6bignum8Big32x4010mul_digits17h90fa0468b79b6f4cE
        drop
      end
      block  ;; label = @2
        local.get 1
        i32.const 128
        i32.and
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        i32.const 1049364
        i32.const 14
        call $_ZN4core3num6bignum8Big32x4010mul_digits17h90fa0468b79b6f4cE
        drop
      end
      block  ;; label = @2
        local.get 1
        i32.const 256
        i32.and
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        i32.const 1049420
        i32.const 27
        call $_ZN4core3num6bignum8Big32x4010mul_digits17h90fa0468b79b6f4cE
        drop
      end
      local.get 0
      return
    end
    local.get 3
    i32.const 40
    i32.const 1054848
    call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
    unreachable)
  (func $_ZN4core3num6bignum8Big32x4010mul_digits17h90fa0468b79b6f4cE (type 1) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i64 i64 i32)
    global.get $__stack_pointer
    i32.const 160
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    local.get 3
    i32.const 0
    i32.const 160
    call $memset
    local.set 4
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 0
          i32.load offset=160
          local.tee 5
          local.get 2
          i32.lt_u
          br_if 0 (;@3;)
          block  ;; label = @4
            local.get 5
            i32.const 41
            i32.ge_u
            br_if 0 (;@4;)
            local.get 1
            local.get 2
            i32.const 2
            i32.shl
            i32.add
            local.set 6
            local.get 5
            i32.eqz
            br_if 2 (;@2;)
            local.get 5
            i32.const 1
            i32.add
            local.set 7
            local.get 5
            i32.const 2
            i32.shl
            local.set 2
            i32.const 0
            local.set 8
            i32.const 0
            local.set 9
            loop  ;; label = @5
              local.get 4
              local.get 8
              i32.const 2
              i32.shl
              i32.add
              local.set 10
              loop  ;; label = @6
                local.get 8
                local.set 11
                local.get 10
                local.set 3
                local.get 1
                local.get 6
                i32.eq
                br_if 5 (;@1;)
                local.get 3
                i32.const 4
                i32.add
                local.set 10
                local.get 11
                i32.const 1
                i32.add
                local.set 8
                local.get 1
                i32.load
                local.set 12
                local.get 1
                i32.const 4
                i32.add
                local.tee 13
                local.set 1
                local.get 12
                i32.eqz
                br_if 0 (;@6;)
              end
              local.get 12
              i64.extend_i32_u
              local.set 14
              i64.const 0
              local.set 15
              local.get 2
              local.set 12
              local.get 11
              local.set 1
              local.get 0
              local.set 10
              block  ;; label = @6
                block  ;; label = @7
                  loop  ;; label = @8
                    local.get 1
                    i32.const 39
                    i32.gt_u
                    br_if 1 (;@7;)
                    local.get 3
                    local.get 15
                    local.get 3
                    i64.load32_u
                    i64.add
                    local.get 10
                    i64.load32_u
                    local.get 14
                    i64.mul
                    i64.add
                    local.tee 15
                    i64.store32
                    local.get 15
                    i64.const 32
                    i64.shr_u
                    local.set 15
                    local.get 3
                    i32.const 4
                    i32.add
                    local.set 3
                    local.get 1
                    i32.const 1
                    i32.add
                    local.set 1
                    local.get 10
                    i32.const 4
                    i32.add
                    local.set 10
                    local.get 12
                    i32.const -4
                    i32.add
                    local.tee 12
                    br_if 0 (;@8;)
                  end
                  local.get 5
                  local.set 3
                  local.get 15
                  i32.wrap_i64
                  local.tee 1
                  i32.eqz
                  br_if 1 (;@6;)
                  block  ;; label = @8
                    local.get 11
                    local.get 5
                    i32.add
                    local.tee 3
                    i32.const 39
                    i32.gt_u
                    br_if 0 (;@8;)
                    local.get 4
                    local.get 3
                    i32.const 2
                    i32.shl
                    i32.add
                    local.get 1
                    i32.store
                    local.get 7
                    local.set 3
                    br 2 (;@6;)
                  end
                  local.get 3
                  i32.const 40
                  i32.const 1054848
                  call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                  unreachable
                end
                local.get 1
                i32.const 40
                i32.const 1054848
                call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                unreachable
              end
              local.get 9
              local.get 3
              local.get 11
              i32.add
              local.tee 3
              local.get 9
              local.get 3
              i32.gt_u
              select
              local.set 9
              local.get 13
              local.set 1
              br 0 (;@5;)
            end
          end
          local.get 5
          i32.const 40
          i32.const 1054848
          call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
          unreachable
        end
        block  ;; label = @3
          local.get 5
          i32.const 41
          i32.ge_u
          br_if 0 (;@3;)
          local.get 2
          i32.const 2
          i32.shl
          local.set 7
          local.get 2
          i32.const 1
          i32.add
          local.set 16
          local.get 0
          local.get 5
          i32.const 2
          i32.shl
          i32.add
          local.set 13
          i32.const 0
          local.set 11
          local.get 0
          local.set 10
          i32.const 0
          local.set 9
          loop  ;; label = @4
            local.get 4
            local.get 11
            i32.const 2
            i32.shl
            i32.add
            local.set 8
            loop  ;; label = @5
              local.get 11
              local.set 12
              local.get 8
              local.set 3
              local.get 10
              local.get 13
              i32.eq
              br_if 4 (;@1;)
              local.get 3
              i32.const 4
              i32.add
              local.set 8
              local.get 12
              i32.const 1
              i32.add
              local.set 11
              local.get 10
              i32.load
              local.set 6
              local.get 10
              i32.const 4
              i32.add
              local.tee 5
              local.set 10
              local.get 6
              i32.eqz
              br_if 0 (;@5;)
            end
            local.get 6
            i64.extend_i32_u
            local.set 14
            i64.const 0
            local.set 15
            local.get 7
            local.set 6
            local.get 12
            local.set 10
            local.get 1
            local.set 8
            block  ;; label = @5
              block  ;; label = @6
                loop  ;; label = @7
                  local.get 10
                  i32.const 39
                  i32.gt_u
                  br_if 1 (;@6;)
                  local.get 3
                  local.get 15
                  local.get 3
                  i64.load32_u
                  i64.add
                  local.get 8
                  i64.load32_u
                  local.get 14
                  i64.mul
                  i64.add
                  local.tee 15
                  i64.store32
                  local.get 15
                  i64.const 32
                  i64.shr_u
                  local.set 15
                  local.get 3
                  i32.const 4
                  i32.add
                  local.set 3
                  local.get 10
                  i32.const 1
                  i32.add
                  local.set 10
                  local.get 8
                  i32.const 4
                  i32.add
                  local.set 8
                  local.get 6
                  i32.const -4
                  i32.add
                  local.tee 6
                  br_if 0 (;@7;)
                end
                local.get 2
                local.set 3
                local.get 15
                i32.wrap_i64
                local.tee 10
                i32.eqz
                br_if 1 (;@5;)
                block  ;; label = @7
                  local.get 12
                  local.get 2
                  i32.add
                  local.tee 3
                  i32.const 39
                  i32.gt_u
                  br_if 0 (;@7;)
                  local.get 4
                  local.get 3
                  i32.const 2
                  i32.shl
                  i32.add
                  local.get 10
                  i32.store
                  local.get 16
                  local.set 3
                  br 2 (;@5;)
                end
                local.get 3
                i32.const 40
                i32.const 1054848
                call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                unreachable
              end
              local.get 10
              i32.const 40
              i32.const 1054848
              call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
              unreachable
            end
            local.get 9
            local.get 3
            local.get 12
            i32.add
            local.tee 3
            local.get 9
            local.get 3
            i32.gt_u
            select
            local.set 9
            local.get 5
            local.set 10
            br 0 (;@4;)
          end
        end
        local.get 5
        i32.const 40
        i32.const 1054848
        call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
        unreachable
      end
      i32.const 0
      local.set 9
      i32.const 0
      local.set 3
      loop  ;; label = @2
        local.get 1
        local.get 6
        i32.eq
        br_if 1 (;@1;)
        local.get 3
        i32.const 1
        i32.add
        local.set 3
        local.get 1
        i32.load
        local.set 10
        local.get 1
        i32.const 4
        i32.add
        local.tee 8
        local.set 1
        local.get 10
        i32.eqz
        br_if 0 (;@2;)
        local.get 9
        local.get 3
        i32.const -1
        i32.add
        local.tee 1
        local.get 9
        local.get 1
        i32.gt_u
        select
        local.set 9
        local.get 8
        local.set 1
        br 0 (;@2;)
      end
    end
    local.get 0
    local.get 4
    i32.const 160
    call $memcpy
    local.tee 3
    local.get 9
    i32.store offset=160
    local.get 4
    i32.const 160
    i32.add
    global.set $__stack_pointer
    local.get 3)
  (func $_ZN4core3num7flt2dec8strategy6dragon15format_shortest17hb3d88f8cf15eff79E (type 12) (param i32 i32 i32 i32)
    (local i32 i64 i64 i64 i64 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 1344
    i32.sub
    local.tee 4
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            block  ;; label = @13
                              block  ;; label = @14
                                block  ;; label = @15
                                  block  ;; label = @16
                                    block  ;; label = @17
                                      block  ;; label = @18
                                        block  ;; label = @19
                                          block  ;; label = @20
                                            block  ;; label = @21
                                              block  ;; label = @22
                                                block  ;; label = @23
                                                  local.get 1
                                                  i64.load
                                                  local.tee 5
                                                  i64.eqz
                                                  br_if 0 (;@23;)
                                                  local.get 1
                                                  i64.load offset=8
                                                  local.tee 6
                                                  i64.eqz
                                                  br_if 1 (;@22;)
                                                  local.get 1
                                                  i64.load offset=16
                                                  local.tee 7
                                                  i64.eqz
                                                  br_if 2 (;@21;)
                                                  local.get 5
                                                  local.get 7
                                                  i64.add
                                                  local.tee 8
                                                  local.get 5
                                                  i64.lt_u
                                                  br_if 3 (;@20;)
                                                  local.get 5
                                                  local.get 6
                                                  i64.lt_u
                                                  br_if 4 (;@19;)
                                                  local.get 3
                                                  i32.const 17
                                                  i32.lt_u
                                                  br_if 5 (;@18;)
                                                  local.get 1
                                                  i32.load8_s offset=26
                                                  local.set 9
                                                  local.get 1
                                                  i32.load16_u offset=24
                                                  local.set 1
                                                  local.get 4
                                                  local.get 5
                                                  i64.store32
                                                  local.get 4
                                                  i32.const 1
                                                  i32.const 2
                                                  local.get 5
                                                  i64.const 4294967296
                                                  i64.lt_u
                                                  local.tee 10
                                                  select
                                                  i32.store offset=160
                                                  local.get 4
                                                  i32.const 0
                                                  local.get 5
                                                  i64.const 32
                                                  i64.shr_u
                                                  i32.wrap_i64
                                                  local.get 10
                                                  select
                                                  i32.store offset=4
                                                  local.get 4
                                                  i32.const 8
                                                  i32.add
                                                  i32.const 0
                                                  i32.const 152
                                                  call $memset
                                                  drop
                                                  local.get 4
                                                  local.get 6
                                                  i64.store32 offset=168
                                                  local.get 4
                                                  i32.const 1
                                                  i32.const 2
                                                  local.get 6
                                                  i64.const 4294967296
                                                  i64.lt_u
                                                  local.tee 10
                                                  select
                                                  i32.store offset=328
                                                  local.get 4
                                                  i32.const 0
                                                  local.get 6
                                                  i64.const 32
                                                  i64.shr_u
                                                  i32.wrap_i64
                                                  local.get 10
                                                  select
                                                  i32.store offset=172
                                                  local.get 4
                                                  i32.const 168
                                                  i32.add
                                                  i32.const 8
                                                  i32.add
                                                  i32.const 0
                                                  i32.const 152
                                                  call $memset
                                                  drop
                                                  local.get 4
                                                  local.get 7
                                                  i64.store32 offset=336
                                                  local.get 4
                                                  i32.const 1
                                                  i32.const 2
                                                  local.get 7
                                                  i64.const 4294967296
                                                  i64.lt_u
                                                  local.tee 10
                                                  select
                                                  i32.store offset=496
                                                  local.get 4
                                                  i32.const 0
                                                  local.get 7
                                                  i64.const 32
                                                  i64.shr_u
                                                  i32.wrap_i64
                                                  local.get 10
                                                  select
                                                  i32.store offset=340
                                                  local.get 4
                                                  i32.const 336
                                                  i32.add
                                                  i32.const 8
                                                  i32.add
                                                  i32.const 0
                                                  i32.const 152
                                                  call $memset
                                                  drop
                                                  local.get 4
                                                  i32.const 504
                                                  i32.add
                                                  i32.const 4
                                                  i32.or
                                                  i32.const 0
                                                  i32.const 156
                                                  call $memset
                                                  drop
                                                  local.get 4
                                                  i32.const 1
                                                  i32.store offset=504
                                                  local.get 4
                                                  i32.const 1
                                                  i32.store offset=664
                                                  local.get 1
                                                  i64.extend_i32_u
                                                  i64.extend16_s
                                                  local.get 8
                                                  i64.const -1
                                                  i64.add
                                                  i64.clz
                                                  i64.sub
                                                  i64.const 1292913986
                                                  i64.mul
                                                  i64.const 82746495104
                                                  i64.add
                                                  i64.const 32
                                                  i64.shr_u
                                                  i32.wrap_i64
                                                  local.tee 10
                                                  i32.extend16_s
                                                  local.set 11
                                                  block  ;; label = @24
                                                    block  ;; label = @25
                                                      local.get 1
                                                      i32.extend16_s
                                                      local.tee 12
                                                      i32.const 0
                                                      i32.lt_s
                                                      br_if 0 (;@25;)
                                                      local.get 4
                                                      local.get 1
                                                      call $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE
                                                      drop
                                                      local.get 4
                                                      i32.const 168
                                                      i32.add
                                                      local.get 1
                                                      call $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE
                                                      drop
                                                      local.get 4
                                                      i32.const 336
                                                      i32.add
                                                      local.get 1
                                                      call $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE
                                                      drop
                                                      br 1 (;@24;)
                                                    end
                                                    local.get 4
                                                    i32.const 504
                                                    i32.add
                                                    i32.const 0
                                                    local.get 12
                                                    i32.sub
                                                    i32.extend16_s
                                                    call $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE
                                                    drop
                                                  end
                                                  block  ;; label = @24
                                                    block  ;; label = @25
                                                      local.get 11
                                                      i32.const -1
                                                      i32.gt_s
                                                      br_if 0 (;@25;)
                                                      local.get 4
                                                      i32.const 0
                                                      local.get 11
                                                      i32.sub
                                                      i32.extend16_s
                                                      local.tee 1
                                                      call $_ZN4core3num7flt2dec8strategy6dragon9mul_pow1017h3a8b434637fe22e2E
                                                      drop
                                                      local.get 4
                                                      i32.const 168
                                                      i32.add
                                                      local.get 1
                                                      call $_ZN4core3num7flt2dec8strategy6dragon9mul_pow1017h3a8b434637fe22e2E
                                                      drop
                                                      local.get 4
                                                      i32.const 336
                                                      i32.add
                                                      local.get 1
                                                      call $_ZN4core3num7flt2dec8strategy6dragon9mul_pow1017h3a8b434637fe22e2E
                                                      drop
                                                      br 1 (;@24;)
                                                    end
                                                    local.get 4
                                                    i32.const 504
                                                    i32.add
                                                    local.get 10
                                                    i32.const 65535
                                                    i32.and
                                                    call $_ZN4core3num7flt2dec8strategy6dragon9mul_pow1017h3a8b434637fe22e2E
                                                    drop
                                                  end
                                                  local.get 4
                                                  local.get 4
                                                  i32.load offset=160
                                                  local.tee 13
                                                  i32.store offset=1336
                                                  local.get 4
                                                  i32.const 1176
                                                  i32.add
                                                  local.get 4
                                                  i32.const 160
                                                  call $memcpy
                                                  drop
                                                  local.get 13
                                                  local.get 4
                                                  i32.load offset=496
                                                  local.tee 14
                                                  local.get 13
                                                  local.get 14
                                                  i32.gt_u
                                                  select
                                                  local.tee 15
                                                  i32.const 40
                                                  i32.gt_u
                                                  br_if 6 (;@17;)
                                                  block  ;; label = @24
                                                    local.get 15
                                                    br_if 0 (;@24;)
                                                    i32.const 0
                                                    local.set 15
                                                    br 9 (;@15;)
                                                  end
                                                  local.get 15
                                                  i32.const 1
                                                  i32.and
                                                  local.set 16
                                                  block  ;; label = @24
                                                    local.get 15
                                                    i32.const 1
                                                    i32.ne
                                                    br_if 0 (;@24;)
                                                    i32.const 0
                                                    local.set 17
                                                    i32.const 0
                                                    local.set 18
                                                    br 8 (;@16;)
                                                  end
                                                  local.get 15
                                                  i32.const -2
                                                  i32.and
                                                  local.set 19
                                                  i32.const 0
                                                  local.set 17
                                                  local.get 4
                                                  i32.const 1176
                                                  i32.add
                                                  local.set 1
                                                  local.get 4
                                                  i32.const 336
                                                  i32.add
                                                  local.set 10
                                                  i32.const 0
                                                  local.set 18
                                                  loop  ;; label = @24
                                                    local.get 1
                                                    local.get 1
                                                    i32.load
                                                    local.tee 20
                                                    local.get 10
                                                    i32.load
                                                    i32.add
                                                    local.tee 12
                                                    local.get 17
                                                    i32.const 1
                                                    i32.and
                                                    i32.add
                                                    local.tee 21
                                                    i32.store
                                                    local.get 1
                                                    i32.const 4
                                                    i32.add
                                                    local.tee 17
                                                    local.get 17
                                                    i32.load
                                                    local.tee 22
                                                    local.get 10
                                                    i32.const 4
                                                    i32.add
                                                    i32.load
                                                    i32.add
                                                    local.tee 17
                                                    local.get 12
                                                    local.get 20
                                                    i32.lt_u
                                                    local.get 21
                                                    local.get 12
                                                    i32.lt_u
                                                    i32.or
                                                    i32.add
                                                    local.tee 12
                                                    i32.store
                                                    local.get 17
                                                    local.get 22
                                                    i32.lt_u
                                                    local.get 12
                                                    local.get 17
                                                    i32.lt_u
                                                    i32.or
                                                    local.set 17
                                                    local.get 10
                                                    i32.const 8
                                                    i32.add
                                                    local.set 10
                                                    local.get 1
                                                    i32.const 8
                                                    i32.add
                                                    local.set 1
                                                    local.get 19
                                                    local.get 18
                                                    i32.const 2
                                                    i32.add
                                                    local.tee 18
                                                    i32.ne
                                                    br_if 0 (;@24;)
                                                    br 8 (;@16;)
                                                  end
                                                end
                                                i32.const 1049575
                                                i32.const 28
                                                i32.const 1049604
                                                call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                                unreachable
                                              end
                                              i32.const 1049620
                                              i32.const 29
                                              i32.const 1049652
                                              call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                              unreachable
                                            end
                                            i32.const 1049668
                                            i32.const 28
                                            i32.const 1049696
                                            call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                            unreachable
                                          end
                                          i32.const 1049712
                                          i32.const 54
                                          i32.const 1049768
                                          call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                          unreachable
                                        end
                                        i32.const 1049784
                                        i32.const 55
                                        i32.const 1049840
                                        call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                        unreachable
                                      end
                                      i32.const 1049856
                                      i32.const 45
                                      i32.const 1049904
                                      call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                      unreachable
                                    end
                                    local.get 15
                                    i32.const 40
                                    i32.const 1054848
                                    call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                    unreachable
                                  end
                                  block  ;; label = @16
                                    local.get 16
                                    i32.eqz
                                    br_if 0 (;@16;)
                                    local.get 4
                                    i32.const 1176
                                    i32.add
                                    local.get 18
                                    i32.const 2
                                    i32.shl
                                    local.tee 1
                                    i32.add
                                    local.tee 10
                                    local.get 10
                                    i32.load
                                    local.tee 10
                                    local.get 4
                                    i32.const 336
                                    i32.add
                                    local.get 1
                                    i32.add
                                    i32.load
                                    i32.add
                                    local.tee 1
                                    local.get 17
                                    i32.add
                                    local.tee 12
                                    i32.store
                                    local.get 1
                                    local.get 10
                                    i32.lt_u
                                    local.get 12
                                    local.get 1
                                    i32.lt_u
                                    i32.or
                                    local.set 17
                                  end
                                  local.get 17
                                  i32.const 1
                                  i32.and
                                  i32.eqz
                                  br_if 0 (;@15;)
                                  local.get 15
                                  i32.const 39
                                  i32.gt_u
                                  br_if 1 (;@14;)
                                  local.get 4
                                  i32.const 1176
                                  i32.add
                                  local.get 15
                                  i32.const 2
                                  i32.shl
                                  i32.add
                                  i32.const 1
                                  i32.store
                                  local.get 15
                                  i32.const 1
                                  i32.add
                                  local.set 15
                                end
                                local.get 4
                                local.get 15
                                i32.store offset=1336
                                local.get 4
                                i32.load offset=664
                                local.tee 18
                                local.get 15
                                local.get 18
                                local.get 15
                                i32.gt_u
                                select
                                local.tee 1
                                i32.const 41
                                i32.ge_u
                                br_if 1 (;@13;)
                                local.get 1
                                i32.const 2
                                i32.shl
                                local.set 1
                                block  ;; label = @15
                                  block  ;; label = @16
                                    loop  ;; label = @17
                                      local.get 1
                                      i32.eqz
                                      br_if 1 (;@16;)
                                      i32.const -1
                                      local.get 1
                                      i32.const -4
                                      i32.add
                                      local.tee 1
                                      local.get 4
                                      i32.const 1176
                                      i32.add
                                      i32.add
                                      i32.load
                                      local.tee 10
                                      local.get 1
                                      local.get 4
                                      i32.const 504
                                      i32.add
                                      i32.add
                                      i32.load
                                      local.tee 12
                                      i32.ne
                                      local.get 10
                                      local.get 12
                                      i32.gt_u
                                      select
                                      local.tee 10
                                      i32.eqz
                                      br_if 0 (;@17;)
                                      br 2 (;@15;)
                                    end
                                  end
                                  i32.const -1
                                  i32.const 0
                                  local.get 1
                                  select
                                  local.set 10
                                end
                                block  ;; label = @15
                                  local.get 10
                                  local.get 9
                                  i32.lt_s
                                  br_if 0 (;@15;)
                                  local.get 13
                                  i32.const 41
                                  i32.ge_u
                                  br_if 3 (;@12;)
                                  block  ;; label = @16
                                    local.get 13
                                    br_if 0 (;@16;)
                                    i32.const 0
                                    local.set 13
                                    br 6 (;@10;)
                                  end
                                  local.get 13
                                  i32.const -1
                                  i32.add
                                  i32.const 1073741823
                                  i32.and
                                  local.tee 1
                                  i32.const 1
                                  i32.add
                                  local.tee 12
                                  i32.const 3
                                  i32.and
                                  local.set 10
                                  block  ;; label = @16
                                    local.get 1
                                    i32.const 3
                                    i32.ge_u
                                    br_if 0 (;@16;)
                                    local.get 4
                                    local.set 1
                                    i64.const 0
                                    local.set 5
                                    br 5 (;@11;)
                                  end
                                  local.get 12
                                  i32.const 2147483644
                                  i32.and
                                  local.set 12
                                  local.get 4
                                  local.set 1
                                  i64.const 0
                                  local.set 5
                                  loop  ;; label = @16
                                    local.get 1
                                    local.get 1
                                    i64.load32_u
                                    i64.const 10
                                    i64.mul
                                    local.get 5
                                    i64.add
                                    local.tee 5
                                    i64.store32
                                    local.get 1
                                    i32.const 4
                                    i32.add
                                    local.tee 17
                                    local.get 17
                                    i64.load32_u
                                    i64.const 10
                                    i64.mul
                                    local.get 5
                                    i64.const 32
                                    i64.shr_u
                                    i64.add
                                    local.tee 5
                                    i64.store32
                                    local.get 1
                                    i32.const 8
                                    i32.add
                                    local.tee 17
                                    local.get 17
                                    i64.load32_u
                                    i64.const 10
                                    i64.mul
                                    local.get 5
                                    i64.const 32
                                    i64.shr_u
                                    i64.add
                                    local.tee 5
                                    i64.store32
                                    local.get 1
                                    i32.const 12
                                    i32.add
                                    local.tee 17
                                    local.get 17
                                    i64.load32_u
                                    i64.const 10
                                    i64.mul
                                    local.get 5
                                    i64.const 32
                                    i64.shr_u
                                    i64.add
                                    local.tee 5
                                    i64.store32
                                    local.get 5
                                    i64.const 32
                                    i64.shr_u
                                    local.set 5
                                    local.get 1
                                    i32.const 16
                                    i32.add
                                    local.set 1
                                    local.get 12
                                    i32.const -4
                                    i32.add
                                    local.tee 12
                                    br_if 0 (;@16;)
                                    br 5 (;@11;)
                                  end
                                end
                                local.get 11
                                i32.const 1
                                i32.add
                                local.set 11
                                br 12 (;@2;)
                              end
                              local.get 15
                              i32.const 40
                              i32.const 1054848
                              call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                              unreachable
                            end
                            local.get 1
                            i32.const 40
                            i32.const 1054848
                            call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                            unreachable
                          end
                          local.get 13
                          i32.const 40
                          i32.const 1054848
                          call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                          unreachable
                        end
                        block  ;; label = @11
                          local.get 10
                          i32.eqz
                          br_if 0 (;@11;)
                          loop  ;; label = @12
                            local.get 1
                            local.get 1
                            i64.load32_u
                            i64.const 10
                            i64.mul
                            local.get 5
                            i64.add
                            local.tee 5
                            i64.store32
                            local.get 1
                            i32.const 4
                            i32.add
                            local.set 1
                            local.get 5
                            i64.const 32
                            i64.shr_u
                            local.set 5
                            local.get 10
                            i32.const -1
                            i32.add
                            local.tee 10
                            br_if 0 (;@12;)
                          end
                        end
                        local.get 5
                        i32.wrap_i64
                        local.tee 1
                        i32.eqz
                        br_if 0 (;@10;)
                        local.get 13
                        i32.const 39
                        i32.gt_u
                        br_if 1 (;@9;)
                        local.get 4
                        local.get 13
                        i32.const 2
                        i32.shl
                        i32.add
                        local.get 1
                        i32.store
                        local.get 13
                        i32.const 1
                        i32.add
                        local.set 13
                      end
                      local.get 4
                      local.get 13
                      i32.store offset=160
                      local.get 4
                      i32.load offset=328
                      local.tee 20
                      i32.const 41
                      i32.ge_u
                      br_if 1 (;@8;)
                      block  ;; label = @10
                        local.get 20
                        br_if 0 (;@10;)
                        i32.const 0
                        local.set 20
                        br 4 (;@6;)
                      end
                      local.get 20
                      i32.const -1
                      i32.add
                      i32.const 1073741823
                      i32.and
                      local.tee 1
                      i32.const 1
                      i32.add
                      local.tee 12
                      i32.const 3
                      i32.and
                      local.set 10
                      block  ;; label = @10
                        local.get 1
                        i32.const 3
                        i32.ge_u
                        br_if 0 (;@10;)
                        local.get 4
                        i32.const 168
                        i32.add
                        local.set 1
                        i64.const 0
                        local.set 5
                        br 3 (;@7;)
                      end
                      local.get 12
                      i32.const 2147483644
                      i32.and
                      local.set 12
                      local.get 4
                      i32.const 168
                      i32.add
                      local.set 1
                      i64.const 0
                      local.set 5
                      loop  ;; label = @10
                        local.get 1
                        local.get 1
                        i64.load32_u
                        i64.const 10
                        i64.mul
                        local.get 5
                        i64.add
                        local.tee 5
                        i64.store32
                        local.get 1
                        i32.const 4
                        i32.add
                        local.tee 17
                        local.get 17
                        i64.load32_u
                        i64.const 10
                        i64.mul
                        local.get 5
                        i64.const 32
                        i64.shr_u
                        i64.add
                        local.tee 5
                        i64.store32
                        local.get 1
                        i32.const 8
                        i32.add
                        local.tee 17
                        local.get 17
                        i64.load32_u
                        i64.const 10
                        i64.mul
                        local.get 5
                        i64.const 32
                        i64.shr_u
                        i64.add
                        local.tee 5
                        i64.store32
                        local.get 1
                        i32.const 12
                        i32.add
                        local.tee 17
                        local.get 17
                        i64.load32_u
                        i64.const 10
                        i64.mul
                        local.get 5
                        i64.const 32
                        i64.shr_u
                        i64.add
                        local.tee 5
                        i64.store32
                        local.get 5
                        i64.const 32
                        i64.shr_u
                        local.set 5
                        local.get 1
                        i32.const 16
                        i32.add
                        local.set 1
                        local.get 12
                        i32.const -4
                        i32.add
                        local.tee 12
                        br_if 0 (;@10;)
                        br 3 (;@7;)
                      end
                    end
                    local.get 13
                    i32.const 40
                    i32.const 1054848
                    call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                    unreachable
                  end
                  local.get 20
                  i32.const 40
                  i32.const 1054848
                  call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                  unreachable
                end
                block  ;; label = @7
                  local.get 10
                  i32.eqz
                  br_if 0 (;@7;)
                  loop  ;; label = @8
                    local.get 1
                    local.get 1
                    i64.load32_u
                    i64.const 10
                    i64.mul
                    local.get 5
                    i64.add
                    local.tee 5
                    i64.store32
                    local.get 1
                    i32.const 4
                    i32.add
                    local.set 1
                    local.get 5
                    i64.const 32
                    i64.shr_u
                    local.set 5
                    local.get 10
                    i32.const -1
                    i32.add
                    local.tee 10
                    br_if 0 (;@8;)
                  end
                end
                local.get 5
                i32.wrap_i64
                local.tee 1
                i32.eqz
                br_if 0 (;@6;)
                local.get 20
                i32.const 39
                i32.gt_u
                br_if 1 (;@5;)
                local.get 4
                i32.const 168
                i32.add
                local.get 20
                i32.const 2
                i32.shl
                i32.add
                local.get 1
                i32.store
                local.get 20
                i32.const 1
                i32.add
                local.set 20
              end
              local.get 4
              local.get 20
              i32.store offset=328
              local.get 14
              i32.const 41
              i32.ge_u
              br_if 1 (;@4;)
              block  ;; label = @6
                local.get 14
                br_if 0 (;@6;)
                local.get 4
                i32.const 0
                i32.store offset=496
                br 4 (;@2;)
              end
              local.get 14
              i32.const -1
              i32.add
              i32.const 1073741823
              i32.and
              local.tee 1
              i32.const 1
              i32.add
              local.tee 12
              i32.const 3
              i32.and
              local.set 10
              block  ;; label = @6
                local.get 1
                i32.const 3
                i32.ge_u
                br_if 0 (;@6;)
                local.get 4
                i32.const 336
                i32.add
                local.set 1
                i64.const 0
                local.set 5
                br 3 (;@3;)
              end
              local.get 12
              i32.const 2147483644
              i32.and
              local.set 12
              local.get 4
              i32.const 336
              i32.add
              local.set 1
              i64.const 0
              local.set 5
              loop  ;; label = @6
                local.get 1
                local.get 1
                i64.load32_u
                i64.const 10
                i64.mul
                local.get 5
                i64.add
                local.tee 5
                i64.store32
                local.get 1
                i32.const 4
                i32.add
                local.tee 17
                local.get 17
                i64.load32_u
                i64.const 10
                i64.mul
                local.get 5
                i64.const 32
                i64.shr_u
                i64.add
                local.tee 5
                i64.store32
                local.get 1
                i32.const 8
                i32.add
                local.tee 17
                local.get 17
                i64.load32_u
                i64.const 10
                i64.mul
                local.get 5
                i64.const 32
                i64.shr_u
                i64.add
                local.tee 5
                i64.store32
                local.get 1
                i32.const 12
                i32.add
                local.tee 17
                local.get 17
                i64.load32_u
                i64.const 10
                i64.mul
                local.get 5
                i64.const 32
                i64.shr_u
                i64.add
                local.tee 5
                i64.store32
                local.get 5
                i64.const 32
                i64.shr_u
                local.set 5
                local.get 1
                i32.const 16
                i32.add
                local.set 1
                local.get 12
                i32.const -4
                i32.add
                local.tee 12
                br_if 0 (;@6;)
                br 3 (;@3;)
              end
            end
            local.get 20
            i32.const 40
            i32.const 1054848
            call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
            unreachable
          end
          local.get 14
          i32.const 40
          i32.const 1054848
          call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
          unreachable
        end
        block  ;; label = @3
          local.get 10
          i32.eqz
          br_if 0 (;@3;)
          loop  ;; label = @4
            local.get 1
            local.get 1
            i64.load32_u
            i64.const 10
            i64.mul
            local.get 5
            i64.add
            local.tee 5
            i64.store32
            local.get 1
            i32.const 4
            i32.add
            local.set 1
            local.get 5
            i64.const 32
            i64.shr_u
            local.set 5
            local.get 10
            i32.const -1
            i32.add
            local.tee 10
            br_if 0 (;@4;)
          end
        end
        block  ;; label = @3
          local.get 5
          i32.wrap_i64
          local.tee 1
          i32.eqz
          br_if 0 (;@3;)
          local.get 14
          i32.const 39
          i32.gt_u
          br_if 2 (;@1;)
          local.get 4
          i32.const 336
          i32.add
          local.get 14
          i32.const 2
          i32.shl
          i32.add
          local.get 1
          i32.store
          local.get 14
          i32.const 1
          i32.add
          local.set 14
        end
        local.get 4
        local.get 14
        i32.store offset=496
      end
      local.get 4
      local.get 18
      i32.store offset=832
      local.get 4
      i32.const 672
      i32.add
      local.get 4
      i32.const 504
      i32.add
      i32.const 160
      call $memcpy
      drop
      local.get 4
      i32.const 672
      i32.add
      i32.const 1
      call $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE
      local.set 23
      local.get 4
      local.get 4
      i32.load offset=664
      i32.store offset=1000
      local.get 4
      i32.const 840
      i32.add
      local.get 4
      i32.const 504
      i32.add
      i32.const 160
      call $memcpy
      drop
      local.get 4
      i32.const 840
      i32.add
      i32.const 2
      call $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE
      local.set 24
      local.get 4
      local.get 4
      i32.load offset=664
      i32.store offset=1168
      local.get 4
      i32.const 1008
      i32.add
      local.get 4
      i32.const 504
      i32.add
      i32.const 160
      call $memcpy
      drop
      local.get 4
      i32.const 1008
      i32.add
      i32.const 3
      call $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE
      local.set 25
      block  ;; label = @2
        block  ;; label = @3
          local.get 4
          i32.load offset=160
          local.tee 19
          local.get 4
          i32.load offset=1168
          local.tee 26
          local.get 19
          local.get 26
          i32.gt_u
          select
          local.tee 15
          i32.const 40
          i32.gt_u
          br_if 0 (;@3;)
          local.get 4
          i32.const 672
          i32.add
          i32.const -4
          i32.add
          local.set 16
          local.get 4
          i32.const 840
          i32.add
          i32.const -4
          i32.add
          local.set 14
          local.get 4
          i32.const 1008
          i32.add
          i32.const -4
          i32.add
          local.set 13
          local.get 4
          i32.load offset=664
          local.set 27
          local.get 4
          i32.load offset=832
          local.set 28
          local.get 4
          i32.load offset=1000
          local.set 29
          i32.const 0
          local.set 30
          loop  ;; label = @4
            local.get 30
            local.set 31
            local.get 15
            i32.const 2
            i32.shl
            local.set 1
            block  ;; label = @5
              block  ;; label = @6
                loop  ;; label = @7
                  local.get 1
                  i32.eqz
                  br_if 1 (;@6;)
                  i32.const -1
                  local.get 13
                  local.get 1
                  i32.add
                  i32.load
                  local.tee 10
                  local.get 1
                  i32.const -4
                  i32.add
                  local.tee 1
                  local.get 4
                  i32.add
                  i32.load
                  local.tee 12
                  i32.ne
                  local.get 10
                  local.get 12
                  i32.gt_u
                  select
                  local.tee 10
                  i32.eqz
                  br_if 0 (;@7;)
                  br 2 (;@5;)
                end
              end
              i32.const -1
              i32.const 0
              local.get 1
              select
              local.set 10
            end
            i32.const 0
            local.set 32
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            block  ;; label = @13
                              block  ;; label = @14
                                block  ;; label = @15
                                  block  ;; label = @16
                                    block  ;; label = @17
                                      block  ;; label = @18
                                        block  ;; label = @19
                                          block  ;; label = @20
                                            block  ;; label = @21
                                              block  ;; label = @22
                                                block  ;; label = @23
                                                  block  ;; label = @24
                                                    block  ;; label = @25
                                                      block  ;; label = @26
                                                        block  ;; label = @27
                                                          block  ;; label = @28
                                                            block  ;; label = @29
                                                              block  ;; label = @30
                                                                block  ;; label = @31
                                                                  local.get 10
                                                                  i32.const 1
                                                                  i32.gt_u
                                                                  br_if 0 (;@31;)
                                                                  block  ;; label = @32
                                                                    local.get 15
                                                                    i32.eqz
                                                                    br_if 0 (;@32;)
                                                                    i32.const 1
                                                                    local.set 17
                                                                    local.get 15
                                                                    i32.const 1
                                                                    i32.and
                                                                    local.set 33
                                                                    i32.const 0
                                                                    local.set 18
                                                                    block  ;; label = @33
                                                                      local.get 15
                                                                      i32.const 1
                                                                      i32.eq
                                                                      br_if 0 (;@33;)
                                                                      local.get 15
                                                                      i32.const -2
                                                                      i32.and
                                                                      local.set 19
                                                                      i32.const 0
                                                                      local.set 18
                                                                      i32.const 1
                                                                      local.set 17
                                                                      local.get 4
                                                                      local.set 1
                                                                      local.get 4
                                                                      i32.const 1008
                                                                      i32.add
                                                                      local.set 10
                                                                      loop  ;; label = @34
                                                                        local.get 1
                                                                        local.get 1
                                                                        i32.load
                                                                        local.tee 20
                                                                        local.get 10
                                                                        i32.load
                                                                        i32.const -1
                                                                        i32.xor
                                                                        i32.add
                                                                        local.tee 12
                                                                        local.get 17
                                                                        i32.const 1
                                                                        i32.and
                                                                        i32.add
                                                                        local.tee 21
                                                                        i32.store
                                                                        local.get 1
                                                                        i32.const 4
                                                                        i32.add
                                                                        local.tee 17
                                                                        local.get 17
                                                                        i32.load
                                                                        local.tee 22
                                                                        local.get 10
                                                                        i32.const 4
                                                                        i32.add
                                                                        i32.load
                                                                        i32.const -1
                                                                        i32.xor
                                                                        i32.add
                                                                        local.tee 17
                                                                        local.get 12
                                                                        local.get 20
                                                                        i32.lt_u
                                                                        local.get 21
                                                                        local.get 12
                                                                        i32.lt_u
                                                                        i32.or
                                                                        i32.add
                                                                        local.tee 12
                                                                        i32.store
                                                                        local.get 17
                                                                        local.get 22
                                                                        i32.lt_u
                                                                        local.get 12
                                                                        local.get 17
                                                                        i32.lt_u
                                                                        i32.or
                                                                        local.set 17
                                                                        local.get 10
                                                                        i32.const 8
                                                                        i32.add
                                                                        local.set 10
                                                                        local.get 1
                                                                        i32.const 8
                                                                        i32.add
                                                                        local.set 1
                                                                        local.get 19
                                                                        local.get 18
                                                                        i32.const 2
                                                                        i32.add
                                                                        local.tee 18
                                                                        i32.ne
                                                                        br_if 0 (;@34;)
                                                                      end
                                                                    end
                                                                    block  ;; label = @33
                                                                      local.get 33
                                                                      i32.eqz
                                                                      br_if 0 (;@33;)
                                                                      local.get 4
                                                                      local.get 18
                                                                      i32.const 2
                                                                      i32.shl
                                                                      local.tee 1
                                                                      i32.add
                                                                      local.tee 10
                                                                      local.get 10
                                                                      i32.load
                                                                      local.tee 10
                                                                      local.get 25
                                                                      local.get 1
                                                                      i32.add
                                                                      i32.load
                                                                      i32.const -1
                                                                      i32.xor
                                                                      i32.add
                                                                      local.tee 1
                                                                      local.get 17
                                                                      i32.add
                                                                      local.tee 12
                                                                      i32.store
                                                                      local.get 1
                                                                      local.get 10
                                                                      i32.lt_u
                                                                      local.get 12
                                                                      local.get 1
                                                                      i32.lt_u
                                                                      i32.or
                                                                      local.set 17
                                                                    end
                                                                    local.get 17
                                                                    i32.const 1
                                                                    i32.and
                                                                    i32.eqz
                                                                    br_if 2 (;@30;)
                                                                  end
                                                                  local.get 4
                                                                  local.get 15
                                                                  i32.store offset=160
                                                                  i32.const 8
                                                                  local.set 32
                                                                  local.get 15
                                                                  local.set 19
                                                                end
                                                                local.get 19
                                                                local.get 29
                                                                local.get 19
                                                                local.get 29
                                                                i32.gt_u
                                                                select
                                                                local.tee 15
                                                                i32.const 41
                                                                i32.ge_u
                                                                br_if 1 (;@29;)
                                                                local.get 15
                                                                i32.const 2
                                                                i32.shl
                                                                local.set 1
                                                                block  ;; label = @31
                                                                  block  ;; label = @32
                                                                    loop  ;; label = @33
                                                                      local.get 1
                                                                      i32.eqz
                                                                      br_if 1 (;@32;)
                                                                      i32.const -1
                                                                      local.get 14
                                                                      local.get 1
                                                                      i32.add
                                                                      i32.load
                                                                      local.tee 10
                                                                      local.get 1
                                                                      i32.const -4
                                                                      i32.add
                                                                      local.tee 1
                                                                      local.get 4
                                                                      i32.add
                                                                      i32.load
                                                                      local.tee 12
                                                                      i32.ne
                                                                      local.get 10
                                                                      local.get 12
                                                                      i32.gt_u
                                                                      select
                                                                      local.tee 10
                                                                      i32.eqz
                                                                      br_if 0 (;@33;)
                                                                      br 2 (;@31;)
                                                                    end
                                                                  end
                                                                  i32.const -1
                                                                  i32.const 0
                                                                  local.get 1
                                                                  select
                                                                  local.set 10
                                                                end
                                                                block  ;; label = @31
                                                                  block  ;; label = @32
                                                                    local.get 10
                                                                    i32.const 1
                                                                    i32.le_u
                                                                    br_if 0 (;@32;)
                                                                    local.get 19
                                                                    local.set 15
                                                                    br 1 (;@31;)
                                                                  end
                                                                  block  ;; label = @32
                                                                    local.get 15
                                                                    i32.eqz
                                                                    br_if 0 (;@32;)
                                                                    i32.const 1
                                                                    local.set 17
                                                                    local.get 15
                                                                    i32.const 1
                                                                    i32.and
                                                                    local.set 33
                                                                    i32.const 0
                                                                    local.set 18
                                                                    block  ;; label = @33
                                                                      local.get 15
                                                                      i32.const 1
                                                                      i32.eq
                                                                      br_if 0 (;@33;)
                                                                      local.get 15
                                                                      i32.const -2
                                                                      i32.and
                                                                      local.set 19
                                                                      i32.const 0
                                                                      local.set 18
                                                                      i32.const 1
                                                                      local.set 17
                                                                      local.get 4
                                                                      local.set 1
                                                                      local.get 4
                                                                      i32.const 840
                                                                      i32.add
                                                                      local.set 10
                                                                      loop  ;; label = @34
                                                                        local.get 1
                                                                        local.get 1
                                                                        i32.load
                                                                        local.tee 20
                                                                        local.get 10
                                                                        i32.load
                                                                        i32.const -1
                                                                        i32.xor
                                                                        i32.add
                                                                        local.tee 12
                                                                        local.get 17
                                                                        i32.const 1
                                                                        i32.and
                                                                        i32.add
                                                                        local.tee 21
                                                                        i32.store
                                                                        local.get 1
                                                                        i32.const 4
                                                                        i32.add
                                                                        local.tee 17
                                                                        local.get 17
                                                                        i32.load
                                                                        local.tee 22
                                                                        local.get 10
                                                                        i32.const 4
                                                                        i32.add
                                                                        i32.load
                                                                        i32.const -1
                                                                        i32.xor
                                                                        i32.add
                                                                        local.tee 17
                                                                        local.get 12
                                                                        local.get 20
                                                                        i32.lt_u
                                                                        local.get 21
                                                                        local.get 12
                                                                        i32.lt_u
                                                                        i32.or
                                                                        i32.add
                                                                        local.tee 12
                                                                        i32.store
                                                                        local.get 17
                                                                        local.get 22
                                                                        i32.lt_u
                                                                        local.get 12
                                                                        local.get 17
                                                                        i32.lt_u
                                                                        i32.or
                                                                        local.set 17
                                                                        local.get 10
                                                                        i32.const 8
                                                                        i32.add
                                                                        local.set 10
                                                                        local.get 1
                                                                        i32.const 8
                                                                        i32.add
                                                                        local.set 1
                                                                        local.get 19
                                                                        local.get 18
                                                                        i32.const 2
                                                                        i32.add
                                                                        local.tee 18
                                                                        i32.ne
                                                                        br_if 0 (;@34;)
                                                                      end
                                                                    end
                                                                    block  ;; label = @33
                                                                      local.get 33
                                                                      i32.eqz
                                                                      br_if 0 (;@33;)
                                                                      local.get 4
                                                                      local.get 18
                                                                      i32.const 2
                                                                      i32.shl
                                                                      local.tee 1
                                                                      i32.add
                                                                      local.tee 10
                                                                      local.get 10
                                                                      i32.load
                                                                      local.tee 10
                                                                      local.get 24
                                                                      local.get 1
                                                                      i32.add
                                                                      i32.load
                                                                      i32.const -1
                                                                      i32.xor
                                                                      i32.add
                                                                      local.tee 1
                                                                      local.get 17
                                                                      i32.add
                                                                      local.tee 12
                                                                      i32.store
                                                                      local.get 1
                                                                      local.get 10
                                                                      i32.lt_u
                                                                      local.get 12
                                                                      local.get 1
                                                                      i32.lt_u
                                                                      i32.or
                                                                      local.set 17
                                                                    end
                                                                    local.get 17
                                                                    i32.const 1
                                                                    i32.and
                                                                    i32.eqz
                                                                    br_if 4 (;@28;)
                                                                  end
                                                                  local.get 4
                                                                  local.get 15
                                                                  i32.store offset=160
                                                                  local.get 32
                                                                  i32.const 4
                                                                  i32.or
                                                                  local.set 32
                                                                end
                                                                local.get 15
                                                                local.get 28
                                                                local.get 15
                                                                local.get 28
                                                                i32.gt_u
                                                                select
                                                                local.tee 33
                                                                i32.const 41
                                                                i32.ge_u
                                                                br_if 3 (;@27;)
                                                                local.get 33
                                                                i32.const 2
                                                                i32.shl
                                                                local.set 1
                                                                block  ;; label = @31
                                                                  block  ;; label = @32
                                                                    loop  ;; label = @33
                                                                      local.get 1
                                                                      i32.eqz
                                                                      br_if 1 (;@32;)
                                                                      i32.const -1
                                                                      local.get 16
                                                                      local.get 1
                                                                      i32.add
                                                                      i32.load
                                                                      local.tee 10
                                                                      local.get 1
                                                                      i32.const -4
                                                                      i32.add
                                                                      local.tee 1
                                                                      local.get 4
                                                                      i32.add
                                                                      i32.load
                                                                      local.tee 12
                                                                      i32.ne
                                                                      local.get 10
                                                                      local.get 12
                                                                      i32.gt_u
                                                                      select
                                                                      local.tee 10
                                                                      i32.eqz
                                                                      br_if 0 (;@33;)
                                                                      br 2 (;@31;)
                                                                    end
                                                                  end
                                                                  i32.const -1
                                                                  i32.const 0
                                                                  local.get 1
                                                                  select
                                                                  local.set 10
                                                                end
                                                                block  ;; label = @31
                                                                  block  ;; label = @32
                                                                    local.get 10
                                                                    i32.const 1
                                                                    i32.le_u
                                                                    br_if 0 (;@32;)
                                                                    local.get 15
                                                                    local.set 33
                                                                    br 1 (;@31;)
                                                                  end
                                                                  block  ;; label = @32
                                                                    local.get 33
                                                                    i32.eqz
                                                                    br_if 0 (;@32;)
                                                                    i32.const 1
                                                                    local.set 17
                                                                    local.get 33
                                                                    i32.const 1
                                                                    i32.and
                                                                    local.set 15
                                                                    i32.const 0
                                                                    local.set 18
                                                                    block  ;; label = @33
                                                                      local.get 33
                                                                      i32.const 1
                                                                      i32.eq
                                                                      br_if 0 (;@33;)
                                                                      local.get 33
                                                                      i32.const -2
                                                                      i32.and
                                                                      local.set 19
                                                                      i32.const 0
                                                                      local.set 18
                                                                      i32.const 1
                                                                      local.set 17
                                                                      local.get 4
                                                                      local.set 1
                                                                      local.get 4
                                                                      i32.const 672
                                                                      i32.add
                                                                      local.set 10
                                                                      loop  ;; label = @34
                                                                        local.get 1
                                                                        local.get 1
                                                                        i32.load
                                                                        local.tee 20
                                                                        local.get 10
                                                                        i32.load
                                                                        i32.const -1
                                                                        i32.xor
                                                                        i32.add
                                                                        local.tee 12
                                                                        local.get 17
                                                                        i32.const 1
                                                                        i32.and
                                                                        i32.add
                                                                        local.tee 21
                                                                        i32.store
                                                                        local.get 1
                                                                        i32.const 4
                                                                        i32.add
                                                                        local.tee 17
                                                                        local.get 17
                                                                        i32.load
                                                                        local.tee 22
                                                                        local.get 10
                                                                        i32.const 4
                                                                        i32.add
                                                                        i32.load
                                                                        i32.const -1
                                                                        i32.xor
                                                                        i32.add
                                                                        local.tee 17
                                                                        local.get 12
                                                                        local.get 20
                                                                        i32.lt_u
                                                                        local.get 21
                                                                        local.get 12
                                                                        i32.lt_u
                                                                        i32.or
                                                                        i32.add
                                                                        local.tee 12
                                                                        i32.store
                                                                        local.get 17
                                                                        local.get 22
                                                                        i32.lt_u
                                                                        local.get 12
                                                                        local.get 17
                                                                        i32.lt_u
                                                                        i32.or
                                                                        local.set 17
                                                                        local.get 10
                                                                        i32.const 8
                                                                        i32.add
                                                                        local.set 10
                                                                        local.get 1
                                                                        i32.const 8
                                                                        i32.add
                                                                        local.set 1
                                                                        local.get 19
                                                                        local.get 18
                                                                        i32.const 2
                                                                        i32.add
                                                                        local.tee 18
                                                                        i32.ne
                                                                        br_if 0 (;@34;)
                                                                      end
                                                                    end
                                                                    block  ;; label = @33
                                                                      local.get 15
                                                                      i32.eqz
                                                                      br_if 0 (;@33;)
                                                                      local.get 4
                                                                      local.get 18
                                                                      i32.const 2
                                                                      i32.shl
                                                                      local.tee 1
                                                                      i32.add
                                                                      local.tee 10
                                                                      local.get 10
                                                                      i32.load
                                                                      local.tee 10
                                                                      local.get 23
                                                                      local.get 1
                                                                      i32.add
                                                                      i32.load
                                                                      i32.const -1
                                                                      i32.xor
                                                                      i32.add
                                                                      local.tee 1
                                                                      local.get 17
                                                                      i32.add
                                                                      local.tee 12
                                                                      i32.store
                                                                      local.get 1
                                                                      local.get 10
                                                                      i32.lt_u
                                                                      local.get 12
                                                                      local.get 1
                                                                      i32.lt_u
                                                                      i32.or
                                                                      local.set 17
                                                                    end
                                                                    local.get 17
                                                                    i32.const 1
                                                                    i32.and
                                                                    i32.eqz
                                                                    br_if 6 (;@26;)
                                                                  end
                                                                  local.get 4
                                                                  local.get 33
                                                                  i32.store offset=160
                                                                  local.get 32
                                                                  i32.const 2
                                                                  i32.add
                                                                  local.set 32
                                                                end
                                                                local.get 33
                                                                local.get 27
                                                                local.get 33
                                                                local.get 27
                                                                i32.gt_u
                                                                select
                                                                local.tee 19
                                                                i32.const 41
                                                                i32.ge_u
                                                                br_if 5 (;@25;)
                                                                local.get 19
                                                                i32.const 2
                                                                i32.shl
                                                                local.set 1
                                                                block  ;; label = @31
                                                                  block  ;; label = @32
                                                                    loop  ;; label = @33
                                                                      local.get 1
                                                                      i32.eqz
                                                                      br_if 1 (;@32;)
                                                                      i32.const -1
                                                                      local.get 1
                                                                      i32.const -4
                                                                      i32.add
                                                                      local.tee 1
                                                                      local.get 4
                                                                      i32.const 504
                                                                      i32.add
                                                                      i32.add
                                                                      i32.load
                                                                      local.tee 10
                                                                      local.get 1
                                                                      local.get 4
                                                                      i32.add
                                                                      i32.load
                                                                      local.tee 12
                                                                      i32.ne
                                                                      local.get 10
                                                                      local.get 12
                                                                      i32.gt_u
                                                                      select
                                                                      local.tee 10
                                                                      i32.eqz
                                                                      br_if 0 (;@33;)
                                                                      br 2 (;@31;)
                                                                    end
                                                                  end
                                                                  i32.const -1
                                                                  i32.const 0
                                                                  local.get 1
                                                                  select
                                                                  local.set 10
                                                                end
                                                                block  ;; label = @31
                                                                  block  ;; label = @32
                                                                    local.get 10
                                                                    i32.const 1
                                                                    i32.le_u
                                                                    br_if 0 (;@32;)
                                                                    local.get 33
                                                                    local.set 19
                                                                    br 1 (;@31;)
                                                                  end
                                                                  block  ;; label = @32
                                                                    local.get 19
                                                                    i32.eqz
                                                                    br_if 0 (;@32;)
                                                                    i32.const 1
                                                                    local.set 17
                                                                    local.get 19
                                                                    i32.const 1
                                                                    i32.and
                                                                    local.set 33
                                                                    i32.const 0
                                                                    local.set 18
                                                                    block  ;; label = @33
                                                                      local.get 19
                                                                      i32.const 1
                                                                      i32.eq
                                                                      br_if 0 (;@33;)
                                                                      local.get 19
                                                                      i32.const -2
                                                                      i32.and
                                                                      local.set 15
                                                                      i32.const 0
                                                                      local.set 18
                                                                      i32.const 1
                                                                      local.set 17
                                                                      local.get 4
                                                                      local.set 1
                                                                      local.get 4
                                                                      i32.const 504
                                                                      i32.add
                                                                      local.set 10
                                                                      loop  ;; label = @34
                                                                        local.get 1
                                                                        local.get 1
                                                                        i32.load
                                                                        local.tee 20
                                                                        local.get 10
                                                                        i32.load
                                                                        i32.const -1
                                                                        i32.xor
                                                                        i32.add
                                                                        local.tee 12
                                                                        local.get 17
                                                                        i32.const 1
                                                                        i32.and
                                                                        i32.add
                                                                        local.tee 21
                                                                        i32.store
                                                                        local.get 1
                                                                        i32.const 4
                                                                        i32.add
                                                                        local.tee 17
                                                                        local.get 17
                                                                        i32.load
                                                                        local.tee 22
                                                                        local.get 10
                                                                        i32.const 4
                                                                        i32.add
                                                                        i32.load
                                                                        i32.const -1
                                                                        i32.xor
                                                                        i32.add
                                                                        local.tee 17
                                                                        local.get 12
                                                                        local.get 20
                                                                        i32.lt_u
                                                                        local.get 21
                                                                        local.get 12
                                                                        i32.lt_u
                                                                        i32.or
                                                                        i32.add
                                                                        local.tee 12
                                                                        i32.store
                                                                        local.get 17
                                                                        local.get 22
                                                                        i32.lt_u
                                                                        local.get 12
                                                                        local.get 17
                                                                        i32.lt_u
                                                                        i32.or
                                                                        local.set 17
                                                                        local.get 10
                                                                        i32.const 8
                                                                        i32.add
                                                                        local.set 10
                                                                        local.get 1
                                                                        i32.const 8
                                                                        i32.add
                                                                        local.set 1
                                                                        local.get 15
                                                                        local.get 18
                                                                        i32.const 2
                                                                        i32.add
                                                                        local.tee 18
                                                                        i32.ne
                                                                        br_if 0 (;@34;)
                                                                      end
                                                                    end
                                                                    block  ;; label = @33
                                                                      local.get 33
                                                                      i32.eqz
                                                                      br_if 0 (;@33;)
                                                                      local.get 4
                                                                      local.get 18
                                                                      i32.const 2
                                                                      i32.shl
                                                                      local.tee 1
                                                                      i32.add
                                                                      local.tee 10
                                                                      local.get 10
                                                                      i32.load
                                                                      local.tee 10
                                                                      local.get 4
                                                                      i32.const 504
                                                                      i32.add
                                                                      local.get 1
                                                                      i32.add
                                                                      i32.load
                                                                      i32.const -1
                                                                      i32.xor
                                                                      i32.add
                                                                      local.tee 1
                                                                      local.get 17
                                                                      i32.add
                                                                      local.tee 12
                                                                      i32.store
                                                                      local.get 1
                                                                      local.get 10
                                                                      i32.lt_u
                                                                      local.get 12
                                                                      local.get 1
                                                                      i32.lt_u
                                                                      i32.or
                                                                      local.set 17
                                                                    end
                                                                    local.get 17
                                                                    i32.const 1
                                                                    i32.and
                                                                    i32.eqz
                                                                    br_if 8 (;@24;)
                                                                  end
                                                                  local.get 4
                                                                  local.get 19
                                                                  i32.store offset=160
                                                                  local.get 32
                                                                  i32.const 1
                                                                  i32.add
                                                                  local.set 32
                                                                end
                                                                local.get 31
                                                                local.get 3
                                                                i32.eq
                                                                br_if 10 (;@20;)
                                                                local.get 2
                                                                local.get 31
                                                                i32.add
                                                                local.get 32
                                                                i32.const 48
                                                                i32.add
                                                                i32.store8
                                                                local.get 19
                                                                local.get 4
                                                                i32.load offset=328
                                                                local.tee 34
                                                                local.get 19
                                                                local.get 34
                                                                i32.gt_u
                                                                select
                                                                local.tee 1
                                                                i32.const 41
                                                                i32.ge_u
                                                                br_if 7 (;@23;)
                                                                local.get 31
                                                                i32.const 1
                                                                i32.add
                                                                local.set 30
                                                                local.get 1
                                                                i32.const 2
                                                                i32.shl
                                                                local.set 1
                                                                block  ;; label = @31
                                                                  block  ;; label = @32
                                                                    loop  ;; label = @33
                                                                      local.get 1
                                                                      i32.eqz
                                                                      br_if 1 (;@32;)
                                                                      i32.const -1
                                                                      local.get 1
                                                                      i32.const -4
                                                                      i32.add
                                                                      local.tee 1
                                                                      local.get 4
                                                                      i32.const 168
                                                                      i32.add
                                                                      i32.add
                                                                      i32.load
                                                                      local.tee 10
                                                                      local.get 1
                                                                      local.get 4
                                                                      i32.add
                                                                      i32.load
                                                                      local.tee 12
                                                                      i32.ne
                                                                      local.get 10
                                                                      local.get 12
                                                                      i32.gt_u
                                                                      select
                                                                      local.tee 15
                                                                      i32.eqz
                                                                      br_if 0 (;@33;)
                                                                      br 2 (;@31;)
                                                                    end
                                                                  end
                                                                  i32.const -1
                                                                  i32.const 0
                                                                  local.get 1
                                                                  select
                                                                  local.set 15
                                                                end
                                                                local.get 4
                                                                local.get 19
                                                                i32.store offset=1336
                                                                local.get 4
                                                                i32.const 1176
                                                                i32.add
                                                                local.get 4
                                                                i32.const 160
                                                                call $memcpy
                                                                drop
                                                                local.get 19
                                                                local.get 4
                                                                i32.load offset=496
                                                                local.tee 35
                                                                local.get 19
                                                                local.get 35
                                                                i32.gt_u
                                                                select
                                                                local.tee 32
                                                                i32.const 40
                                                                i32.gt_u
                                                                br_if 14 (;@16;)
                                                                block  ;; label = @31
                                                                  block  ;; label = @32
                                                                    local.get 32
                                                                    br_if 0 (;@32;)
                                                                    i32.const 0
                                                                    local.set 32
                                                                    br 1 (;@31;)
                                                                  end
                                                                  local.get 32
                                                                  i32.const 1
                                                                  i32.and
                                                                  local.set 36
                                                                  i32.const 0
                                                                  local.set 17
                                                                  i32.const 0
                                                                  local.set 18
                                                                  block  ;; label = @32
                                                                    local.get 32
                                                                    i32.const 1
                                                                    i32.eq
                                                                    br_if 0 (;@32;)
                                                                    local.get 32
                                                                    i32.const -2
                                                                    i32.and
                                                                    local.set 33
                                                                    i32.const 0
                                                                    local.set 17
                                                                    local.get 4
                                                                    i32.const 1176
                                                                    i32.add
                                                                    local.set 1
                                                                    local.get 4
                                                                    i32.const 336
                                                                    i32.add
                                                                    local.set 10
                                                                    i32.const 0
                                                                    local.set 18
                                                                    loop  ;; label = @33
                                                                      local.get 1
                                                                      local.get 1
                                                                      i32.load
                                                                      local.tee 20
                                                                      local.get 10
                                                                      i32.load
                                                                      i32.add
                                                                      local.tee 12
                                                                      local.get 17
                                                                      i32.const 1
                                                                      i32.and
                                                                      i32.add
                                                                      local.tee 21
                                                                      i32.store
                                                                      local.get 1
                                                                      i32.const 4
                                                                      i32.add
                                                                      local.tee 17
                                                                      local.get 17
                                                                      i32.load
                                                                      local.tee 22
                                                                      local.get 10
                                                                      i32.const 4
                                                                      i32.add
                                                                      i32.load
                                                                      i32.add
                                                                      local.tee 17
                                                                      local.get 12
                                                                      local.get 20
                                                                      i32.lt_u
                                                                      local.get 21
                                                                      local.get 12
                                                                      i32.lt_u
                                                                      i32.or
                                                                      i32.add
                                                                      local.tee 12
                                                                      i32.store
                                                                      local.get 17
                                                                      local.get 22
                                                                      i32.lt_u
                                                                      local.get 12
                                                                      local.get 17
                                                                      i32.lt_u
                                                                      i32.or
                                                                      local.set 17
                                                                      local.get 10
                                                                      i32.const 8
                                                                      i32.add
                                                                      local.set 10
                                                                      local.get 1
                                                                      i32.const 8
                                                                      i32.add
                                                                      local.set 1
                                                                      local.get 33
                                                                      local.get 18
                                                                      i32.const 2
                                                                      i32.add
                                                                      local.tee 18
                                                                      i32.ne
                                                                      br_if 0 (;@33;)
                                                                    end
                                                                  end
                                                                  block  ;; label = @32
                                                                    local.get 36
                                                                    i32.eqz
                                                                    br_if 0 (;@32;)
                                                                    local.get 4
                                                                    i32.const 1176
                                                                    i32.add
                                                                    local.get 18
                                                                    i32.const 2
                                                                    i32.shl
                                                                    local.tee 1
                                                                    i32.add
                                                                    local.tee 10
                                                                    local.get 10
                                                                    i32.load
                                                                    local.tee 10
                                                                    local.get 4
                                                                    i32.const 336
                                                                    i32.add
                                                                    local.get 1
                                                                    i32.add
                                                                    i32.load
                                                                    i32.add
                                                                    local.tee 1
                                                                    local.get 17
                                                                    i32.add
                                                                    local.tee 12
                                                                    i32.store
                                                                    local.get 1
                                                                    local.get 10
                                                                    i32.lt_u
                                                                    local.get 12
                                                                    local.get 1
                                                                    i32.lt_u
                                                                    i32.or
                                                                    local.set 17
                                                                  end
                                                                  local.get 17
                                                                  i32.const 1
                                                                  i32.and
                                                                  i32.eqz
                                                                  br_if 0 (;@31;)
                                                                  local.get 32
                                                                  i32.const 39
                                                                  i32.gt_u
                                                                  br_if 9 (;@22;)
                                                                  local.get 4
                                                                  i32.const 1176
                                                                  i32.add
                                                                  local.get 32
                                                                  i32.const 2
                                                                  i32.shl
                                                                  i32.add
                                                                  i32.const 1
                                                                  i32.store
                                                                  local.get 32
                                                                  i32.const 1
                                                                  i32.add
                                                                  local.set 32
                                                                end
                                                                local.get 4
                                                                local.get 32
                                                                i32.store offset=1336
                                                                local.get 27
                                                                local.get 32
                                                                local.get 27
                                                                local.get 32
                                                                i32.gt_u
                                                                select
                                                                local.tee 1
                                                                i32.const 41
                                                                i32.ge_u
                                                                br_if 9 (;@21;)
                                                                local.get 1
                                                                i32.const 2
                                                                i32.shl
                                                                local.set 1
                                                                block  ;; label = @31
                                                                  block  ;; label = @32
                                                                    loop  ;; label = @33
                                                                      local.get 1
                                                                      i32.eqz
                                                                      br_if 1 (;@32;)
                                                                      i32.const -1
                                                                      local.get 1
                                                                      i32.const -4
                                                                      i32.add
                                                                      local.tee 1
                                                                      local.get 4
                                                                      i32.const 1176
                                                                      i32.add
                                                                      i32.add
                                                                      i32.load
                                                                      local.tee 10
                                                                      local.get 1
                                                                      local.get 4
                                                                      i32.const 504
                                                                      i32.add
                                                                      i32.add
                                                                      i32.load
                                                                      local.tee 12
                                                                      i32.ne
                                                                      local.get 10
                                                                      local.get 12
                                                                      i32.gt_u
                                                                      select
                                                                      local.tee 10
                                                                      i32.eqz
                                                                      br_if 0 (;@33;)
                                                                      br 2 (;@31;)
                                                                    end
                                                                  end
                                                                  i32.const -1
                                                                  i32.const 0
                                                                  local.get 1
                                                                  select
                                                                  local.set 10
                                                                end
                                                                block  ;; label = @31
                                                                  local.get 15
                                                                  local.get 9
                                                                  i32.lt_s
                                                                  br_if 0 (;@31;)
                                                                  local.get 10
                                                                  local.get 9
                                                                  i32.lt_s
                                                                  br_if 0 (;@31;)
                                                                  local.get 19
                                                                  i32.const 41
                                                                  i32.ge_u
                                                                  br_if 12 (;@19;)
                                                                  block  ;; label = @32
                                                                    local.get 19
                                                                    br_if 0 (;@32;)
                                                                    i32.const 0
                                                                    local.set 19
                                                                    br 19 (;@13;)
                                                                  end
                                                                  local.get 19
                                                                  i32.const -1
                                                                  i32.add
                                                                  i32.const 1073741823
                                                                  i32.and
                                                                  local.tee 1
                                                                  i32.const 1
                                                                  i32.add
                                                                  local.tee 12
                                                                  i32.const 3
                                                                  i32.and
                                                                  local.set 10
                                                                  block  ;; label = @32
                                                                    local.get 1
                                                                    i32.const 3
                                                                    i32.ge_u
                                                                    br_if 0 (;@32;)
                                                                    local.get 4
                                                                    local.set 1
                                                                    i64.const 0
                                                                    local.set 5
                                                                    br 18 (;@14;)
                                                                  end
                                                                  local.get 12
                                                                  i32.const 2147483644
                                                                  i32.and
                                                                  local.set 12
                                                                  local.get 4
                                                                  local.set 1
                                                                  i64.const 0
                                                                  local.set 5
                                                                  loop  ;; label = @32
                                                                    local.get 1
                                                                    local.get 1
                                                                    i64.load32_u
                                                                    i64.const 10
                                                                    i64.mul
                                                                    local.get 5
                                                                    i64.add
                                                                    local.tee 5
                                                                    i64.store32
                                                                    local.get 1
                                                                    i32.const 4
                                                                    i32.add
                                                                    local.tee 17
                                                                    local.get 17
                                                                    i64.load32_u
                                                                    i64.const 10
                                                                    i64.mul
                                                                    local.get 5
                                                                    i64.const 32
                                                                    i64.shr_u
                                                                    i64.add
                                                                    local.tee 5
                                                                    i64.store32
                                                                    local.get 1
                                                                    i32.const 8
                                                                    i32.add
                                                                    local.tee 17
                                                                    local.get 17
                                                                    i64.load32_u
                                                                    i64.const 10
                                                                    i64.mul
                                                                    local.get 5
                                                                    i64.const 32
                                                                    i64.shr_u
                                                                    i64.add
                                                                    local.tee 5
                                                                    i64.store32
                                                                    local.get 1
                                                                    i32.const 12
                                                                    i32.add
                                                                    local.tee 17
                                                                    local.get 17
                                                                    i64.load32_u
                                                                    i64.const 10
                                                                    i64.mul
                                                                    local.get 5
                                                                    i64.const 32
                                                                    i64.shr_u
                                                                    i64.add
                                                                    local.tee 5
                                                                    i64.store32
                                                                    local.get 5
                                                                    i64.const 32
                                                                    i64.shr_u
                                                                    local.set 5
                                                                    local.get 1
                                                                    i32.const 16
                                                                    i32.add
                                                                    local.set 1
                                                                    local.get 12
                                                                    i32.const -4
                                                                    i32.add
                                                                    local.tee 12
                                                                    br_if 0 (;@32;)
                                                                    br 18 (;@14;)
                                                                  end
                                                                end
                                                                local.get 10
                                                                local.get 9
                                                                i32.ge_s
                                                                br_if 15 (;@15;)
                                                                block  ;; label = @31
                                                                  local.get 15
                                                                  local.get 9
                                                                  i32.ge_s
                                                                  br_if 0 (;@31;)
                                                                  local.get 4
                                                                  i32.const 1
                                                                  call $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE
                                                                  drop
                                                                  local.get 4
                                                                  i32.load offset=160
                                                                  local.tee 1
                                                                  local.get 4
                                                                  i32.load offset=664
                                                                  local.tee 10
                                                                  local.get 1
                                                                  local.get 10
                                                                  i32.gt_u
                                                                  select
                                                                  local.tee 1
                                                                  i32.const 41
                                                                  i32.ge_u
                                                                  br_if 13 (;@18;)
                                                                  local.get 1
                                                                  i32.const 2
                                                                  i32.shl
                                                                  local.set 1
                                                                  local.get 4
                                                                  i32.const -4
                                                                  i32.add
                                                                  local.set 17
                                                                  local.get 4
                                                                  i32.const 504
                                                                  i32.add
                                                                  i32.const -4
                                                                  i32.add
                                                                  local.set 18
                                                                  block  ;; label = @32
                                                                    block  ;; label = @33
                                                                      loop  ;; label = @34
                                                                        local.get 1
                                                                        i32.eqz
                                                                        br_if 1 (;@33;)
                                                                        local.get 17
                                                                        local.get 1
                                                                        i32.add
                                                                        local.set 10
                                                                        local.get 18
                                                                        local.get 1
                                                                        i32.add
                                                                        local.set 12
                                                                        local.get 1
                                                                        i32.const -4
                                                                        i32.add
                                                                        local.set 1
                                                                        i32.const -1
                                                                        local.get 12
                                                                        i32.load
                                                                        local.tee 12
                                                                        local.get 10
                                                                        i32.load
                                                                        local.tee 10
                                                                        i32.ne
                                                                        local.get 12
                                                                        local.get 10
                                                                        i32.gt_u
                                                                        select
                                                                        local.tee 10
                                                                        i32.eqz
                                                                        br_if 0 (;@34;)
                                                                        br 2 (;@32;)
                                                                      end
                                                                    end
                                                                    i32.const -1
                                                                    i32.const 0
                                                                    local.get 1
                                                                    select
                                                                    local.set 10
                                                                  end
                                                                  local.get 10
                                                                  i32.const 2
                                                                  i32.ge_u
                                                                  br_if 16 (;@15;)
                                                                end
                                                                local.get 31
                                                                local.get 3
                                                                i32.ge_u
                                                                br_if 13 (;@17;)
                                                                local.get 2
                                                                local.get 30
                                                                i32.add
                                                                local.set 18
                                                                local.get 31
                                                                local.set 1
                                                                i32.const -1
                                                                local.set 10
                                                                block  ;; label = @31
                                                                  loop  ;; label = @32
                                                                    local.get 1
                                                                    i32.const -1
                                                                    i32.eq
                                                                    br_if 1 (;@31;)
                                                                    local.get 10
                                                                    i32.const 1
                                                                    i32.add
                                                                    local.set 10
                                                                    local.get 2
                                                                    local.get 1
                                                                    i32.add
                                                                    local.set 12
                                                                    local.get 1
                                                                    i32.const -1
                                                                    i32.add
                                                                    local.tee 17
                                                                    local.set 1
                                                                    local.get 12
                                                                    i32.load8_u
                                                                    i32.const 57
                                                                    i32.eq
                                                                    br_if 0 (;@32;)
                                                                  end
                                                                  local.get 2
                                                                  local.get 17
                                                                  i32.add
                                                                  local.tee 12
                                                                  i32.const 1
                                                                  i32.add
                                                                  local.tee 1
                                                                  local.get 1
                                                                  i32.load8_u
                                                                  i32.const 1
                                                                  i32.add
                                                                  i32.store8
                                                                  local.get 17
                                                                  i32.const 2
                                                                  i32.add
                                                                  local.get 31
                                                                  i32.gt_u
                                                                  br_if 16 (;@15;)
                                                                  local.get 12
                                                                  i32.const 2
                                                                  i32.add
                                                                  i32.const 48
                                                                  local.get 10
                                                                  call $memset
                                                                  drop
                                                                  br 16 (;@15;)
                                                                end
                                                                local.get 2
                                                                i32.const 49
                                                                i32.store8
                                                                block  ;; label = @31
                                                                  local.get 31
                                                                  i32.eqz
                                                                  br_if 0 (;@31;)
                                                                  local.get 2
                                                                  i32.const 1
                                                                  i32.add
                                                                  i32.const 48
                                                                  local.get 31
                                                                  call $memset
                                                                  drop
                                                                end
                                                                block  ;; label = @31
                                                                  local.get 30
                                                                  local.get 3
                                                                  i32.ge_u
                                                                  br_if 0 (;@31;)
                                                                  local.get 18
                                                                  i32.const 48
                                                                  i32.store8
                                                                  local.get 11
                                                                  i32.const 1
                                                                  i32.add
                                                                  local.set 11
                                                                  local.get 31
                                                                  i32.const 2
                                                                  i32.add
                                                                  local.set 30
                                                                  br 16 (;@15;)
                                                                end
                                                                local.get 30
                                                                local.get 3
                                                                i32.const 1049952
                                                                call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                                                                unreachable
                                                              end
                                                              i32.const 1054864
                                                              i32.const 26
                                                              i32.const 1054848
                                                              call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                                              unreachable
                                                            end
                                                            local.get 15
                                                            i32.const 40
                                                            i32.const 1054848
                                                            call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                                            unreachable
                                                          end
                                                          i32.const 1054864
                                                          i32.const 26
                                                          i32.const 1054848
                                                          call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                                          unreachable
                                                        end
                                                        local.get 33
                                                        i32.const 40
                                                        i32.const 1054848
                                                        call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                                        unreachable
                                                      end
                                                      i32.const 1054864
                                                      i32.const 26
                                                      i32.const 1054848
                                                      call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                                      unreachable
                                                    end
                                                    local.get 19
                                                    i32.const 40
                                                    i32.const 1054848
                                                    call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                                    unreachable
                                                  end
                                                  i32.const 1054864
                                                  i32.const 26
                                                  i32.const 1054848
                                                  call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                                  unreachable
                                                end
                                                local.get 1
                                                i32.const 40
                                                i32.const 1054848
                                                call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                                unreachable
                                              end
                                              local.get 32
                                              i32.const 40
                                              i32.const 1054848
                                              call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                                              unreachable
                                            end
                                            local.get 1
                                            i32.const 40
                                            i32.const 1054848
                                            call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                            unreachable
                                          end
                                          local.get 3
                                          local.get 3
                                          i32.const 1049920
                                          call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                                          unreachable
                                        end
                                        local.get 19
                                        i32.const 40
                                        i32.const 1054848
                                        call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                        unreachable
                                      end
                                      local.get 1
                                      i32.const 40
                                      i32.const 1054848
                                      call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                      unreachable
                                    end
                                    local.get 30
                                    local.get 3
                                    i32.const 1049936
                                    call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                    unreachable
                                  end
                                  local.get 32
                                  i32.const 40
                                  i32.const 1054848
                                  call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                  unreachable
                                end
                                block  ;; label = @15
                                  local.get 30
                                  local.get 3
                                  i32.gt_u
                                  br_if 0 (;@15;)
                                  local.get 0
                                  local.get 11
                                  i32.store16 offset=8
                                  local.get 0
                                  local.get 30
                                  i32.store offset=4
                                  local.get 0
                                  local.get 2
                                  i32.store
                                  local.get 4
                                  i32.const 1344
                                  i32.add
                                  global.set $__stack_pointer
                                  return
                                end
                                local.get 30
                                local.get 3
                                i32.const 1049968
                                call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                unreachable
                              end
                              block  ;; label = @14
                                local.get 10
                                i32.eqz
                                br_if 0 (;@14;)
                                loop  ;; label = @15
                                  local.get 1
                                  local.get 1
                                  i64.load32_u
                                  i64.const 10
                                  i64.mul
                                  local.get 5
                                  i64.add
                                  local.tee 5
                                  i64.store32
                                  local.get 1
                                  i32.const 4
                                  i32.add
                                  local.set 1
                                  local.get 5
                                  i64.const 32
                                  i64.shr_u
                                  local.set 5
                                  local.get 10
                                  i32.const -1
                                  i32.add
                                  local.tee 10
                                  br_if 0 (;@15;)
                                end
                              end
                              local.get 5
                              i32.wrap_i64
                              local.tee 1
                              i32.eqz
                              br_if 0 (;@13;)
                              local.get 19
                              i32.const 39
                              i32.gt_u
                              br_if 1 (;@12;)
                              local.get 4
                              local.get 19
                              i32.const 2
                              i32.shl
                              i32.add
                              local.get 1
                              i32.store
                              local.get 19
                              i32.const 1
                              i32.add
                              local.set 19
                            end
                            local.get 4
                            local.get 19
                            i32.store offset=160
                            local.get 34
                            i32.const 41
                            i32.ge_u
                            br_if 1 (;@11;)
                            block  ;; label = @13
                              local.get 34
                              br_if 0 (;@13;)
                              i32.const 0
                              local.set 34
                              br 4 (;@9;)
                            end
                            local.get 34
                            i32.const -1
                            i32.add
                            i32.const 1073741823
                            i32.and
                            local.tee 1
                            i32.const 1
                            i32.add
                            local.tee 12
                            i32.const 3
                            i32.and
                            local.set 10
                            block  ;; label = @13
                              local.get 1
                              i32.const 3
                              i32.ge_u
                              br_if 0 (;@13;)
                              local.get 4
                              i32.const 168
                              i32.add
                              local.set 1
                              i64.const 0
                              local.set 5
                              br 3 (;@10;)
                            end
                            local.get 12
                            i32.const 2147483644
                            i32.and
                            local.set 12
                            local.get 4
                            i32.const 168
                            i32.add
                            local.set 1
                            i64.const 0
                            local.set 5
                            loop  ;; label = @13
                              local.get 1
                              local.get 1
                              i64.load32_u
                              i64.const 10
                              i64.mul
                              local.get 5
                              i64.add
                              local.tee 5
                              i64.store32
                              local.get 1
                              i32.const 4
                              i32.add
                              local.tee 17
                              local.get 17
                              i64.load32_u
                              i64.const 10
                              i64.mul
                              local.get 5
                              i64.const 32
                              i64.shr_u
                              i64.add
                              local.tee 5
                              i64.store32
                              local.get 1
                              i32.const 8
                              i32.add
                              local.tee 17
                              local.get 17
                              i64.load32_u
                              i64.const 10
                              i64.mul
                              local.get 5
                              i64.const 32
                              i64.shr_u
                              i64.add
                              local.tee 5
                              i64.store32
                              local.get 1
                              i32.const 12
                              i32.add
                              local.tee 17
                              local.get 17
                              i64.load32_u
                              i64.const 10
                              i64.mul
                              local.get 5
                              i64.const 32
                              i64.shr_u
                              i64.add
                              local.tee 5
                              i64.store32
                              local.get 5
                              i64.const 32
                              i64.shr_u
                              local.set 5
                              local.get 1
                              i32.const 16
                              i32.add
                              local.set 1
                              local.get 12
                              i32.const -4
                              i32.add
                              local.tee 12
                              br_if 0 (;@13;)
                              br 3 (;@10;)
                            end
                          end
                          local.get 19
                          i32.const 40
                          i32.const 1054848
                          call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                          unreachable
                        end
                        local.get 34
                        i32.const 40
                        i32.const 1054848
                        call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                        unreachable
                      end
                      block  ;; label = @10
                        local.get 10
                        i32.eqz
                        br_if 0 (;@10;)
                        loop  ;; label = @11
                          local.get 1
                          local.get 1
                          i64.load32_u
                          i64.const 10
                          i64.mul
                          local.get 5
                          i64.add
                          local.tee 5
                          i64.store32
                          local.get 1
                          i32.const 4
                          i32.add
                          local.set 1
                          local.get 5
                          i64.const 32
                          i64.shr_u
                          local.set 5
                          local.get 10
                          i32.const -1
                          i32.add
                          local.tee 10
                          br_if 0 (;@11;)
                        end
                      end
                      local.get 5
                      i32.wrap_i64
                      local.tee 1
                      i32.eqz
                      br_if 0 (;@9;)
                      local.get 34
                      i32.const 39
                      i32.gt_u
                      br_if 1 (;@8;)
                      local.get 4
                      i32.const 168
                      i32.add
                      local.get 34
                      i32.const 2
                      i32.shl
                      i32.add
                      local.get 1
                      i32.store
                      local.get 34
                      i32.const 1
                      i32.add
                      local.set 34
                    end
                    local.get 4
                    local.get 34
                    i32.store offset=328
                    local.get 35
                    i32.const 41
                    i32.ge_u
                    br_if 1 (;@7;)
                    block  ;; label = @9
                      local.get 35
                      br_if 0 (;@9;)
                      i32.const 0
                      local.set 35
                      br 4 (;@5;)
                    end
                    local.get 35
                    i32.const -1
                    i32.add
                    i32.const 1073741823
                    i32.and
                    local.tee 1
                    i32.const 1
                    i32.add
                    local.tee 12
                    i32.const 3
                    i32.and
                    local.set 10
                    block  ;; label = @9
                      local.get 1
                      i32.const 3
                      i32.ge_u
                      br_if 0 (;@9;)
                      local.get 4
                      i32.const 336
                      i32.add
                      local.set 1
                      i64.const 0
                      local.set 5
                      br 3 (;@6;)
                    end
                    local.get 12
                    i32.const 2147483644
                    i32.and
                    local.set 12
                    local.get 4
                    i32.const 336
                    i32.add
                    local.set 1
                    i64.const 0
                    local.set 5
                    loop  ;; label = @9
                      local.get 1
                      local.get 1
                      i64.load32_u
                      i64.const 10
                      i64.mul
                      local.get 5
                      i64.add
                      local.tee 5
                      i64.store32
                      local.get 1
                      i32.const 4
                      i32.add
                      local.tee 17
                      local.get 17
                      i64.load32_u
                      i64.const 10
                      i64.mul
                      local.get 5
                      i64.const 32
                      i64.shr_u
                      i64.add
                      local.tee 5
                      i64.store32
                      local.get 1
                      i32.const 8
                      i32.add
                      local.tee 17
                      local.get 17
                      i64.load32_u
                      i64.const 10
                      i64.mul
                      local.get 5
                      i64.const 32
                      i64.shr_u
                      i64.add
                      local.tee 5
                      i64.store32
                      local.get 1
                      i32.const 12
                      i32.add
                      local.tee 17
                      local.get 17
                      i64.load32_u
                      i64.const 10
                      i64.mul
                      local.get 5
                      i64.const 32
                      i64.shr_u
                      i64.add
                      local.tee 5
                      i64.store32
                      local.get 5
                      i64.const 32
                      i64.shr_u
                      local.set 5
                      local.get 1
                      i32.const 16
                      i32.add
                      local.set 1
                      local.get 12
                      i32.const -4
                      i32.add
                      local.tee 12
                      br_if 0 (;@9;)
                      br 3 (;@6;)
                    end
                  end
                  local.get 34
                  i32.const 40
                  i32.const 1054848
                  call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                  unreachable
                end
                local.get 35
                i32.const 40
                i32.const 1054848
                call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                unreachable
              end
              block  ;; label = @6
                local.get 10
                i32.eqz
                br_if 0 (;@6;)
                loop  ;; label = @7
                  local.get 1
                  local.get 1
                  i64.load32_u
                  i64.const 10
                  i64.mul
                  local.get 5
                  i64.add
                  local.tee 5
                  i64.store32
                  local.get 1
                  i32.const 4
                  i32.add
                  local.set 1
                  local.get 5
                  i64.const 32
                  i64.shr_u
                  local.set 5
                  local.get 10
                  i32.const -1
                  i32.add
                  local.tee 10
                  br_if 0 (;@7;)
                end
              end
              local.get 5
              i32.wrap_i64
              local.tee 1
              i32.eqz
              br_if 0 (;@5;)
              local.get 35
              i32.const 39
              i32.gt_u
              br_if 3 (;@2;)
              local.get 4
              i32.const 336
              i32.add
              local.get 35
              i32.const 2
              i32.shl
              i32.add
              local.get 1
              i32.store
              local.get 35
              i32.const 1
              i32.add
              local.set 35
            end
            local.get 4
            local.get 35
            i32.store offset=496
            local.get 19
            local.get 26
            local.get 19
            local.get 26
            i32.gt_u
            select
            local.tee 15
            i32.const 40
            i32.le_u
            br_if 0 (;@4;)
          end
        end
        local.get 15
        i32.const 40
        i32.const 1054848
        call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
        unreachable
      end
      local.get 35
      i32.const 40
      i32.const 1054848
      call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
      unreachable
    end
    local.get 14
    i32.const 40
    i32.const 1054848
    call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
    unreachable)
  (func $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE (type 2) (param i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.const 1279
        i32.gt_u
        br_if 0 (;@2;)
        local.get 1
        i32.const 5
        i32.shr_u
        local.set 2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 0
                i32.load offset=160
                local.tee 3
                i32.eqz
                br_if 0 (;@6;)
                local.get 3
                i32.const -1
                i32.add
                local.set 4
                local.get 3
                i32.const 2
                i32.shl
                local.get 0
                i32.add
                i32.const -4
                i32.add
                local.set 5
                local.get 3
                local.get 2
                i32.add
                i32.const 2
                i32.shl
                local.get 0
                i32.add
                i32.const -4
                i32.add
                local.set 6
                local.get 3
                i32.const 40
                i32.gt_u
                local.set 3
                loop  ;; label = @7
                  local.get 3
                  br_if 4 (;@3;)
                  local.get 2
                  local.get 4
                  i32.add
                  local.tee 7
                  i32.const 40
                  i32.ge_u
                  br_if 2 (;@5;)
                  local.get 6
                  local.get 5
                  i32.load
                  i32.store
                  local.get 6
                  i32.const -4
                  i32.add
                  local.set 6
                  local.get 5
                  i32.const -4
                  i32.add
                  local.set 5
                  local.get 4
                  i32.const -1
                  i32.add
                  local.tee 4
                  i32.const -1
                  i32.ne
                  br_if 0 (;@7;)
                end
              end
              local.get 1
              i32.const 32
              i32.lt_u
              br_if 4 (;@1;)
              local.get 0
              i32.const 0
              i32.store
              local.get 1
              i32.const 64
              i32.ge_u
              br_if 1 (;@4;)
              br 4 (;@1;)
            end
            local.get 7
            i32.const 40
            i32.const 1054848
            call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
            unreachable
          end
          local.get 0
          i32.const 0
          i32.store offset=4
          local.get 2
          i32.const 1
          local.get 2
          i32.const 1
          i32.gt_u
          select
          local.tee 4
          i32.const 2
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=8
          local.get 4
          i32.const 3
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=12
          local.get 4
          i32.const 4
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=16
          local.get 4
          i32.const 5
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=20
          local.get 4
          i32.const 6
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=24
          local.get 4
          i32.const 7
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=28
          local.get 4
          i32.const 8
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=32
          local.get 4
          i32.const 9
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=36
          local.get 4
          i32.const 10
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=40
          local.get 4
          i32.const 11
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=44
          local.get 4
          i32.const 12
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=48
          local.get 4
          i32.const 13
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=52
          local.get 4
          i32.const 14
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=56
          local.get 4
          i32.const 15
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=60
          local.get 4
          i32.const 16
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=64
          local.get 4
          i32.const 17
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=68
          local.get 4
          i32.const 18
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=72
          local.get 4
          i32.const 19
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=76
          local.get 4
          i32.const 20
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=80
          local.get 4
          i32.const 21
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=84
          local.get 4
          i32.const 22
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=88
          local.get 4
          i32.const 23
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=92
          local.get 4
          i32.const 24
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=96
          local.get 4
          i32.const 25
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=100
          local.get 4
          i32.const 26
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=104
          local.get 4
          i32.const 27
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=108
          local.get 4
          i32.const 28
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=112
          local.get 4
          i32.const 29
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=116
          local.get 4
          i32.const 30
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=120
          local.get 4
          i32.const 31
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=124
          local.get 4
          i32.const 32
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=128
          local.get 4
          i32.const 33
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=132
          local.get 4
          i32.const 34
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=136
          local.get 4
          i32.const 35
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=140
          local.get 4
          i32.const 36
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=144
          local.get 4
          i32.const 37
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=148
          local.get 4
          i32.const 38
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=152
          local.get 4
          i32.const 39
          i32.eq
          br_if 2 (;@1;)
          local.get 0
          i32.const 0
          i32.store offset=156
          local.get 4
          i32.const 40
          i32.eq
          br_if 2 (;@1;)
          i32.const 40
          i32.const 40
          i32.const 1054848
          call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
          unreachable
        end
        local.get 4
        i32.const 40
        i32.const 1054848
        call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
        unreachable
      end
      i32.const 1054890
      i32.const 29
      i32.const 1054848
      call $_ZN4core9panicking5panic17he2fadc50abd7e245E
      unreachable
    end
    local.get 0
    i32.load offset=160
    local.get 2
    i32.add
    local.set 5
    block  ;; label = @1
      local.get 1
      i32.const 31
      i32.and
      local.tee 3
      br_if 0 (;@1;)
      local.get 0
      local.get 5
      i32.store offset=160
      local.get 0
      return
    end
    block  ;; label = @1
      block  ;; label = @2
        local.get 5
        i32.const -1
        i32.add
        local.tee 4
        i32.const 39
        i32.gt_u
        br_if 0 (;@2;)
        local.get 5
        local.set 8
        local.get 0
        local.get 4
        i32.const 2
        i32.shl
        i32.add
        i32.load
        local.tee 6
        i32.const 0
        local.get 1
        i32.sub
        local.tee 1
        i32.shr_u
        local.tee 4
        i32.eqz
        br_if 1 (;@1;)
        block  ;; label = @3
          local.get 5
          i32.const 39
          i32.gt_u
          br_if 0 (;@3;)
          local.get 0
          local.get 5
          i32.const 2
          i32.shl
          i32.add
          local.get 4
          i32.store
          local.get 5
          i32.const 1
          i32.add
          local.set 8
          br 2 (;@1;)
        end
        local.get 5
        i32.const 40
        i32.const 1054848
        call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
        unreachable
      end
      local.get 4
      i32.const 40
      i32.const 1054848
      call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
      unreachable
    end
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.const 1
        i32.add
        local.tee 7
        local.get 5
        i32.ge_u
        br_if 0 (;@2;)
        local.get 1
        i32.const 31
        i32.and
        local.set 1
        local.get 5
        i32.const 2
        i32.shl
        local.get 0
        i32.add
        i32.const -8
        i32.add
        local.set 4
        loop  ;; label = @3
          local.get 5
          i32.const -2
          i32.add
          i32.const 40
          i32.ge_u
          br_if 2 (;@1;)
          local.get 4
          i32.const 4
          i32.add
          local.get 6
          local.get 3
          i32.shl
          local.get 4
          i32.load
          local.tee 6
          local.get 1
          i32.shr_u
          i32.or
          i32.store
          local.get 4
          i32.const -4
          i32.add
          local.set 4
          local.get 7
          local.get 5
          i32.const -1
          i32.add
          local.tee 5
          i32.lt_u
          br_if 0 (;@3;)
        end
      end
      local.get 0
      local.get 2
      i32.const 2
      i32.shl
      i32.add
      local.tee 4
      local.get 4
      i32.load
      local.get 3
      i32.shl
      i32.store
      local.get 0
      local.get 8
      i32.store offset=160
      local.get 0
      return
    end
    i32.const -1
    i32.const 40
    i32.const 1054848
    call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
    unreachable)
  (func $_ZN4core3num7flt2dec8strategy6dragon12format_exact17h6d7faa24131201fdE (type 15) (param i32 i32 i32 i32 i32)
    (local i32 i64 i64 i64 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 848
    i32.sub
    local.tee 5
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            block  ;; label = @13
                              block  ;; label = @14
                                block  ;; label = @15
                                  block  ;; label = @16
                                    block  ;; label = @17
                                      block  ;; label = @18
                                        local.get 1
                                        i64.load
                                        local.tee 6
                                        i64.eqz
                                        br_if 0 (;@18;)
                                        local.get 1
                                        i64.load offset=8
                                        local.tee 7
                                        i64.eqz
                                        br_if 1 (;@17;)
                                        local.get 1
                                        i64.load offset=16
                                        local.tee 8
                                        i64.eqz
                                        br_if 2 (;@16;)
                                        local.get 6
                                        local.get 8
                                        i64.add
                                        local.get 6
                                        i64.lt_u
                                        br_if 3 (;@15;)
                                        local.get 6
                                        local.get 7
                                        i64.lt_u
                                        br_if 4 (;@14;)
                                        local.get 1
                                        i32.load16_u offset=24
                                        local.set 1
                                        local.get 5
                                        local.get 6
                                        i64.store32 offset=8
                                        local.get 5
                                        i32.const 1
                                        i32.const 2
                                        local.get 6
                                        i64.const 4294967296
                                        i64.lt_u
                                        local.tee 9
                                        select
                                        i32.store offset=168
                                        local.get 5
                                        i32.const 0
                                        local.get 6
                                        i64.const 32
                                        i64.shr_u
                                        i32.wrap_i64
                                        local.get 9
                                        select
                                        i32.store offset=12
                                        local.get 5
                                        i32.const 16
                                        i32.add
                                        i32.const 0
                                        i32.const 152
                                        call $memset
                                        drop
                                        local.get 5
                                        i32.const 176
                                        i32.add
                                        i32.const 4
                                        i32.or
                                        i32.const 0
                                        i32.const 156
                                        call $memset
                                        drop
                                        local.get 5
                                        i32.const 1
                                        i32.store offset=176
                                        local.get 5
                                        i32.const 1
                                        i32.store offset=336
                                        local.get 1
                                        i64.extend_i32_u
                                        i64.extend16_s
                                        local.get 6
                                        i64.const -1
                                        i64.add
                                        i64.clz
                                        i64.sub
                                        i64.const 1292913986
                                        i64.mul
                                        i64.const 82746495104
                                        i64.add
                                        i64.const 32
                                        i64.shr_u
                                        i32.wrap_i64
                                        local.tee 9
                                        i32.extend16_s
                                        local.set 10
                                        block  ;; label = @19
                                          block  ;; label = @20
                                            local.get 1
                                            i32.extend16_s
                                            local.tee 11
                                            i32.const 0
                                            i32.lt_s
                                            br_if 0 (;@20;)
                                            local.get 5
                                            i32.const 8
                                            i32.add
                                            local.get 1
                                            call $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE
                                            drop
                                            br 1 (;@19;)
                                          end
                                          local.get 5
                                          i32.const 176
                                          i32.add
                                          i32.const 0
                                          local.get 11
                                          i32.sub
                                          i32.extend16_s
                                          call $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE
                                          drop
                                        end
                                        block  ;; label = @19
                                          block  ;; label = @20
                                            local.get 10
                                            i32.const -1
                                            i32.gt_s
                                            br_if 0 (;@20;)
                                            local.get 5
                                            i32.const 8
                                            i32.add
                                            i32.const 0
                                            local.get 10
                                            i32.sub
                                            i32.extend16_s
                                            call $_ZN4core3num7flt2dec8strategy6dragon9mul_pow1017h3a8b434637fe22e2E
                                            drop
                                            br 1 (;@19;)
                                          end
                                          local.get 5
                                          i32.const 176
                                          i32.add
                                          local.get 9
                                          i32.const 65535
                                          i32.and
                                          call $_ZN4core3num7flt2dec8strategy6dragon9mul_pow1017h3a8b434637fe22e2E
                                          drop
                                        end
                                        local.get 5
                                        local.get 5
                                        i32.load offset=336
                                        local.tee 12
                                        i32.store offset=840
                                        local.get 5
                                        i32.const 680
                                        i32.add
                                        local.get 5
                                        i32.const 176
                                        i32.add
                                        i32.const 160
                                        call $memcpy
                                        drop
                                        local.get 3
                                        local.set 13
                                        block  ;; label = @19
                                          local.get 3
                                          i32.const 10
                                          i32.lt_u
                                          br_if 0 (;@19;)
                                          block  ;; label = @20
                                            block  ;; label = @21
                                              local.get 12
                                              i32.const 40
                                              i32.le_u
                                              br_if 0 (;@21;)
                                              local.get 12
                                              local.set 1
                                              br 1 (;@20;)
                                            end
                                            local.get 5
                                            i32.const 680
                                            i32.add
                                            i32.const -8
                                            i32.add
                                            local.set 14
                                            local.get 3
                                            local.set 13
                                            local.get 12
                                            local.set 1
                                            loop  ;; label = @21
                                              block  ;; label = @22
                                                local.get 1
                                                i32.eqz
                                                br_if 0 (;@22;)
                                                local.get 1
                                                i32.const -1
                                                i32.add
                                                i32.const 1073741823
                                                i32.and
                                                local.tee 9
                                                i32.const 1
                                                i32.add
                                                local.tee 11
                                                i32.const 1
                                                i32.and
                                                local.set 15
                                                local.get 1
                                                i32.const 2
                                                i32.shl
                                                local.set 1
                                                block  ;; label = @23
                                                  block  ;; label = @24
                                                    local.get 9
                                                    br_if 0 (;@24;)
                                                    local.get 5
                                                    i32.const 680
                                                    i32.add
                                                    local.get 1
                                                    i32.add
                                                    local.set 1
                                                    i64.const 0
                                                    local.set 6
                                                    br 1 (;@23;)
                                                  end
                                                  local.get 11
                                                  i32.const 2147483646
                                                  i32.and
                                                  local.set 9
                                                  local.get 14
                                                  local.get 1
                                                  i32.add
                                                  local.set 1
                                                  i64.const 0
                                                  local.set 6
                                                  loop  ;; label = @24
                                                    local.get 1
                                                    i32.const 4
                                                    i32.add
                                                    local.tee 11
                                                    local.get 6
                                                    i64.const 32
                                                    i64.shl
                                                    local.get 11
                                                    i64.load32_u
                                                    i64.or
                                                    local.tee 6
                                                    i64.const 1000000000
                                                    i64.div_u
                                                    local.tee 8
                                                    i64.store32
                                                    local.get 1
                                                    local.get 6
                                                    local.get 8
                                                    i64.const 1000000000
                                                    i64.mul
                                                    i64.sub
                                                    i64.const 32
                                                    i64.shl
                                                    local.get 1
                                                    i64.load32_u
                                                    i64.or
                                                    local.tee 6
                                                    i64.const 1000000000
                                                    i64.div_u
                                                    local.tee 8
                                                    i64.store32
                                                    local.get 6
                                                    local.get 8
                                                    i64.const 1000000000
                                                    i64.mul
                                                    i64.sub
                                                    local.set 6
                                                    local.get 1
                                                    i32.const -8
                                                    i32.add
                                                    local.set 1
                                                    local.get 9
                                                    i32.const -2
                                                    i32.add
                                                    local.tee 9
                                                    br_if 0 (;@24;)
                                                  end
                                                  local.get 1
                                                  i32.const 8
                                                  i32.add
                                                  local.set 1
                                                end
                                                local.get 15
                                                i32.eqz
                                                br_if 0 (;@22;)
                                                local.get 1
                                                i32.const -4
                                                i32.add
                                                local.tee 1
                                                local.get 6
                                                i64.const 32
                                                i64.shl
                                                local.get 1
                                                i64.load32_u
                                                i64.or
                                                i64.const 1000000000
                                                i64.div_u
                                                i64.store32
                                              end
                                              local.get 13
                                              i32.const -9
                                              i32.add
                                              local.tee 13
                                              i32.const 9
                                              i32.le_u
                                              br_if 2 (;@19;)
                                              local.get 5
                                              i32.load offset=840
                                              local.tee 1
                                              i32.const 41
                                              i32.lt_u
                                              br_if 0 (;@21;)
                                            end
                                          end
                                          local.get 1
                                          i32.const 40
                                          i32.const 1054848
                                          call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                          unreachable
                                        end
                                        block  ;; label = @19
                                          block  ;; label = @20
                                            block  ;; label = @21
                                              block  ;; label = @22
                                                local.get 13
                                                i32.const 2
                                                i32.shl
                                                i32.const 1049272
                                                i32.add
                                                i32.load
                                                local.tee 9
                                                i32.eqz
                                                br_if 0 (;@22;)
                                                local.get 5
                                                i32.load offset=840
                                                local.tee 1
                                                i32.const 41
                                                i32.ge_u
                                                br_if 9 (;@13;)
                                                block  ;; label = @23
                                                  local.get 1
                                                  br_if 0 (;@23;)
                                                  i32.const 0
                                                  local.set 1
                                                  br 4 (;@19;)
                                                end
                                                local.get 1
                                                i32.const -1
                                                i32.add
                                                i32.const 1073741823
                                                i32.and
                                                local.tee 11
                                                i32.const 1
                                                i32.add
                                                local.tee 15
                                                i32.const 1
                                                i32.and
                                                local.set 13
                                                local.get 1
                                                i32.const 2
                                                i32.shl
                                                local.set 1
                                                local.get 9
                                                i64.extend_i32_u
                                                local.set 6
                                                local.get 11
                                                br_if 1 (;@21;)
                                                local.get 5
                                                i32.const 680
                                                i32.add
                                                local.get 1
                                                i32.add
                                                local.set 1
                                                i64.const 0
                                                local.set 8
                                                br 2 (;@20;)
                                              end
                                              i32.const 1054919
                                              i32.const 27
                                              i32.const 1054848
                                              call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                              unreachable
                                            end
                                            local.get 15
                                            i32.const 2147483646
                                            i32.and
                                            local.set 9
                                            local.get 1
                                            local.get 5
                                            i32.const 680
                                            i32.add
                                            i32.add
                                            i32.const -8
                                            i32.add
                                            local.set 1
                                            i64.const 0
                                            local.set 8
                                            loop  ;; label = @21
                                              local.get 1
                                              i32.const 4
                                              i32.add
                                              local.tee 11
                                              local.get 8
                                              i64.const 32
                                              i64.shl
                                              local.get 11
                                              i64.load32_u
                                              i64.or
                                              local.tee 8
                                              local.get 6
                                              i64.div_u
                                              local.tee 7
                                              i64.store32
                                              local.get 1
                                              local.get 8
                                              local.get 7
                                              local.get 6
                                              i64.mul
                                              i64.sub
                                              i64.const 32
                                              i64.shl
                                              local.get 1
                                              i64.load32_u
                                              i64.or
                                              local.tee 8
                                              local.get 6
                                              i64.div_u
                                              local.tee 7
                                              i64.store32
                                              local.get 8
                                              local.get 7
                                              local.get 6
                                              i64.mul
                                              i64.sub
                                              local.set 8
                                              local.get 1
                                              i32.const -8
                                              i32.add
                                              local.set 1
                                              local.get 9
                                              i32.const -2
                                              i32.add
                                              local.tee 9
                                              br_if 0 (;@21;)
                                            end
                                            local.get 1
                                            i32.const 8
                                            i32.add
                                            local.set 1
                                          end
                                          block  ;; label = @20
                                            local.get 13
                                            i32.eqz
                                            br_if 0 (;@20;)
                                            local.get 1
                                            i32.const -4
                                            i32.add
                                            local.tee 1
                                            local.get 8
                                            i64.const 32
                                            i64.shl
                                            local.get 1
                                            i64.load32_u
                                            i64.or
                                            local.get 6
                                            i64.div_u
                                            i64.store32
                                          end
                                          local.get 5
                                          i32.load offset=840
                                          local.set 1
                                        end
                                        local.get 1
                                        local.get 5
                                        i32.load offset=168
                                        local.tee 16
                                        local.get 1
                                        local.get 16
                                        i32.gt_u
                                        select
                                        local.tee 17
                                        i32.const 40
                                        i32.gt_u
                                        br_if 6 (;@12;)
                                        block  ;; label = @19
                                          local.get 17
                                          br_if 0 (;@19;)
                                          i32.const 0
                                          local.set 17
                                          br 9 (;@10;)
                                        end
                                        local.get 17
                                        i32.const 1
                                        i32.and
                                        local.set 18
                                        block  ;; label = @19
                                          local.get 17
                                          i32.const 1
                                          i32.ne
                                          br_if 0 (;@19;)
                                          i32.const 0
                                          local.set 13
                                          i32.const 0
                                          local.set 15
                                          br 8 (;@11;)
                                        end
                                        local.get 17
                                        i32.const -2
                                        i32.and
                                        local.set 19
                                        i32.const 0
                                        local.set 13
                                        local.get 5
                                        i32.const 680
                                        i32.add
                                        local.set 1
                                        local.get 5
                                        i32.const 8
                                        i32.add
                                        local.set 9
                                        i32.const 0
                                        local.set 15
                                        loop  ;; label = @19
                                          local.get 1
                                          local.get 1
                                          i32.load
                                          local.tee 14
                                          local.get 9
                                          i32.load
                                          i32.add
                                          local.tee 11
                                          local.get 13
                                          i32.const 1
                                          i32.and
                                          i32.add
                                          local.tee 20
                                          i32.store
                                          local.get 1
                                          i32.const 4
                                          i32.add
                                          local.tee 13
                                          local.get 13
                                          i32.load
                                          local.tee 21
                                          local.get 9
                                          i32.const 4
                                          i32.add
                                          i32.load
                                          i32.add
                                          local.tee 13
                                          local.get 11
                                          local.get 14
                                          i32.lt_u
                                          local.get 20
                                          local.get 11
                                          i32.lt_u
                                          i32.or
                                          i32.add
                                          local.tee 11
                                          i32.store
                                          local.get 13
                                          local.get 21
                                          i32.lt_u
                                          local.get 11
                                          local.get 13
                                          i32.lt_u
                                          i32.or
                                          local.set 13
                                          local.get 9
                                          i32.const 8
                                          i32.add
                                          local.set 9
                                          local.get 1
                                          i32.const 8
                                          i32.add
                                          local.set 1
                                          local.get 19
                                          local.get 15
                                          i32.const 2
                                          i32.add
                                          local.tee 15
                                          i32.ne
                                          br_if 0 (;@19;)
                                          br 8 (;@11;)
                                        end
                                      end
                                      i32.const 1049575
                                      i32.const 28
                                      i32.const 1049984
                                      call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                      unreachable
                                    end
                                    i32.const 1049620
                                    i32.const 29
                                    i32.const 1050000
                                    call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                    unreachable
                                  end
                                  i32.const 1049668
                                  i32.const 28
                                  i32.const 1050016
                                  call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                  unreachable
                                end
                                i32.const 1049712
                                i32.const 54
                                i32.const 1050032
                                call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                unreachable
                              end
                              i32.const 1049784
                              i32.const 55
                              i32.const 1050048
                              call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                              unreachable
                            end
                            local.get 1
                            i32.const 40
                            i32.const 1054848
                            call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                            unreachable
                          end
                          local.get 17
                          i32.const 40
                          i32.const 1054848
                          call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                          unreachable
                        end
                        block  ;; label = @11
                          local.get 18
                          i32.eqz
                          br_if 0 (;@11;)
                          local.get 5
                          i32.const 680
                          i32.add
                          local.get 15
                          i32.const 2
                          i32.shl
                          local.tee 1
                          i32.add
                          local.tee 9
                          local.get 9
                          i32.load
                          local.tee 9
                          local.get 5
                          i32.const 8
                          i32.add
                          local.get 1
                          i32.add
                          i32.load
                          i32.add
                          local.tee 1
                          local.get 13
                          i32.add
                          local.tee 11
                          i32.store
                          local.get 1
                          local.get 9
                          i32.lt_u
                          local.get 11
                          local.get 1
                          i32.lt_u
                          i32.or
                          local.set 13
                        end
                        local.get 13
                        i32.const 1
                        i32.and
                        i32.eqz
                        br_if 0 (;@10;)
                        local.get 17
                        i32.const 39
                        i32.gt_u
                        br_if 1 (;@9;)
                        local.get 5
                        i32.const 680
                        i32.add
                        local.get 17
                        i32.const 2
                        i32.shl
                        i32.add
                        i32.const 1
                        i32.store
                        local.get 17
                        i32.const 1
                        i32.add
                        local.set 17
                      end
                      local.get 5
                      local.get 17
                      i32.store offset=840
                      local.get 17
                      local.get 12
                      local.get 17
                      local.get 12
                      i32.gt_u
                      select
                      local.tee 1
                      i32.const 41
                      i32.ge_u
                      br_if 1 (;@8;)
                      local.get 1
                      i32.const 2
                      i32.shl
                      local.set 1
                      block  ;; label = @10
                        block  ;; label = @11
                          loop  ;; label = @12
                            local.get 1
                            i32.eqz
                            br_if 1 (;@11;)
                            i32.const -1
                            local.get 1
                            i32.const -4
                            i32.add
                            local.tee 1
                            local.get 5
                            i32.const 176
                            i32.add
                            i32.add
                            i32.load
                            local.tee 9
                            local.get 1
                            local.get 5
                            i32.const 680
                            i32.add
                            i32.add
                            i32.load
                            local.tee 11
                            i32.ne
                            local.get 9
                            local.get 11
                            i32.gt_u
                            select
                            local.tee 9
                            i32.eqz
                            br_if 0 (;@12;)
                            br 2 (;@10;)
                          end
                        end
                        i32.const -1
                        i32.const 0
                        local.get 1
                        select
                        local.set 9
                      end
                      block  ;; label = @10
                        local.get 9
                        i32.const 1
                        i32.gt_u
                        br_if 0 (;@10;)
                        local.get 10
                        i32.const 1
                        i32.add
                        local.set 10
                        br 6 (;@4;)
                      end
                      local.get 16
                      i32.const 41
                      i32.ge_u
                      br_if 2 (;@7;)
                      block  ;; label = @10
                        local.get 16
                        br_if 0 (;@10;)
                        i32.const 0
                        local.set 16
                        br 5 (;@5;)
                      end
                      local.get 16
                      i32.const -1
                      i32.add
                      i32.const 1073741823
                      i32.and
                      local.tee 1
                      i32.const 1
                      i32.add
                      local.tee 11
                      i32.const 3
                      i32.and
                      local.set 9
                      block  ;; label = @10
                        local.get 1
                        i32.const 3
                        i32.ge_u
                        br_if 0 (;@10;)
                        local.get 5
                        i32.const 8
                        i32.add
                        local.set 1
                        i64.const 0
                        local.set 6
                        br 4 (;@6;)
                      end
                      local.get 11
                      i32.const 2147483644
                      i32.and
                      local.set 11
                      local.get 5
                      i32.const 8
                      i32.add
                      local.set 1
                      i64.const 0
                      local.set 6
                      loop  ;; label = @10
                        local.get 1
                        local.get 1
                        i64.load32_u
                        i64.const 10
                        i64.mul
                        local.get 6
                        i64.add
                        local.tee 6
                        i64.store32
                        local.get 1
                        i32.const 4
                        i32.add
                        local.tee 13
                        local.get 13
                        i64.load32_u
                        i64.const 10
                        i64.mul
                        local.get 6
                        i64.const 32
                        i64.shr_u
                        i64.add
                        local.tee 6
                        i64.store32
                        local.get 1
                        i32.const 8
                        i32.add
                        local.tee 13
                        local.get 13
                        i64.load32_u
                        i64.const 10
                        i64.mul
                        local.get 6
                        i64.const 32
                        i64.shr_u
                        i64.add
                        local.tee 6
                        i64.store32
                        local.get 1
                        i32.const 12
                        i32.add
                        local.tee 13
                        local.get 13
                        i64.load32_u
                        i64.const 10
                        i64.mul
                        local.get 6
                        i64.const 32
                        i64.shr_u
                        i64.add
                        local.tee 6
                        i64.store32
                        local.get 6
                        i64.const 32
                        i64.shr_u
                        local.set 6
                        local.get 1
                        i32.const 16
                        i32.add
                        local.set 1
                        local.get 11
                        i32.const -4
                        i32.add
                        local.tee 11
                        br_if 0 (;@10;)
                        br 4 (;@6;)
                      end
                    end
                    local.get 17
                    i32.const 40
                    i32.const 1054848
                    call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                    unreachable
                  end
                  local.get 1
                  i32.const 40
                  i32.const 1054848
                  call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                  unreachable
                end
                local.get 16
                i32.const 40
                i32.const 1054848
                call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                unreachable
              end
              block  ;; label = @6
                local.get 9
                i32.eqz
                br_if 0 (;@6;)
                loop  ;; label = @7
                  local.get 1
                  local.get 1
                  i64.load32_u
                  i64.const 10
                  i64.mul
                  local.get 6
                  i64.add
                  local.tee 6
                  i64.store32
                  local.get 1
                  i32.const 4
                  i32.add
                  local.set 1
                  local.get 6
                  i64.const 32
                  i64.shr_u
                  local.set 6
                  local.get 9
                  i32.const -1
                  i32.add
                  local.tee 9
                  br_if 0 (;@7;)
                end
              end
              local.get 6
              i32.wrap_i64
              local.tee 1
              i32.eqz
              br_if 0 (;@5;)
              local.get 16
              i32.const 39
              i32.gt_u
              br_if 2 (;@3;)
              local.get 5
              i32.const 8
              i32.add
              local.get 16
              i32.const 2
              i32.shl
              i32.add
              local.get 1
              i32.store
              local.get 16
              i32.const 1
              i32.add
              local.set 16
            end
            local.get 5
            local.get 16
            i32.store offset=168
          end
          i32.const 0
          local.set 14
          block  ;; label = @4
            block  ;; label = @5
              local.get 10
              i32.extend16_s
              local.tee 1
              local.get 4
              i32.extend16_s
              local.tee 22
              i32.lt_s
              br_if 0 (;@5;)
              local.get 10
              local.get 4
              i32.sub
              i32.extend16_s
              local.get 3
              local.get 1
              local.get 22
              i32.sub
              local.get 3
              i32.lt_u
              select
              local.tee 13
              br_if 1 (;@4;)
              i32.const 0
              local.set 14
            end
            i32.const 0
            local.set 13
            br 2 (;@2;)
          end
          local.get 5
          local.get 12
          i32.store offset=504
          local.get 5
          i32.const 344
          i32.add
          local.get 5
          i32.const 176
          i32.add
          i32.const 160
          call $memcpy
          drop
          local.get 5
          i32.const 344
          i32.add
          i32.const 1
          call $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE
          local.set 23
          local.get 5
          local.get 5
          i32.load offset=336
          i32.store offset=672
          local.get 5
          i32.const 512
          i32.add
          local.get 5
          i32.const 176
          i32.add
          i32.const 160
          call $memcpy
          drop
          local.get 5
          i32.const 512
          i32.add
          i32.const 2
          call $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE
          local.set 24
          local.get 5
          local.get 5
          i32.load offset=336
          i32.store offset=840
          local.get 5
          i32.const 680
          i32.add
          local.get 5
          i32.const 176
          i32.add
          i32.const 160
          call $memcpy
          drop
          local.get 5
          i32.const 176
          i32.add
          i32.const -4
          i32.add
          local.set 17
          local.get 5
          i32.const 344
          i32.add
          i32.const -4
          i32.add
          local.set 19
          local.get 5
          i32.const 512
          i32.add
          i32.const -4
          i32.add
          local.set 21
          local.get 5
          i32.const 680
          i32.add
          i32.const -4
          i32.add
          local.set 20
          local.get 5
          i32.const 680
          i32.add
          i32.const 3
          call $_ZN4core3num6bignum8Big32x408mul_pow217h4ac708fc9b1da84dE
          local.set 25
          local.get 5
          i32.load offset=168
          local.set 14
          local.get 5
          i32.load offset=336
          local.set 12
          local.get 5
          i32.load offset=504
          local.set 26
          local.get 5
          i32.load offset=672
          local.set 27
          local.get 5
          i32.load offset=840
          local.set 28
          i32.const 0
          local.set 29
          block  ;; label = @4
            loop  ;; label = @5
              local.get 29
              local.set 30
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          block  ;; label = @12
                            block  ;; label = @13
                              block  ;; label = @14
                                block  ;; label = @15
                                  block  ;; label = @16
                                    block  ;; label = @17
                                      block  ;; label = @18
                                        block  ;; label = @19
                                          local.get 14
                                          i32.const 41
                                          i32.ge_u
                                          br_if 0 (;@19;)
                                          local.get 30
                                          i32.const 1
                                          i32.add
                                          local.set 29
                                          local.get 14
                                          i32.const 2
                                          i32.shl
                                          local.set 11
                                          i32.const 0
                                          local.set 1
                                          block  ;; label = @20
                                            block  ;; label = @21
                                              block  ;; label = @22
                                                loop  ;; label = @23
                                                  local.get 11
                                                  local.get 1
                                                  i32.eq
                                                  br_if 1 (;@22;)
                                                  local.get 5
                                                  i32.const 8
                                                  i32.add
                                                  local.get 1
                                                  i32.add
                                                  local.set 9
                                                  local.get 1
                                                  i32.const 4
                                                  i32.add
                                                  local.set 1
                                                  local.get 9
                                                  i32.load
                                                  i32.eqz
                                                  br_if 0 (;@23;)
                                                end
                                                local.get 14
                                                local.get 28
                                                local.get 14
                                                local.get 28
                                                i32.gt_u
                                                select
                                                local.tee 31
                                                i32.const 41
                                                i32.ge_u
                                                br_if 4 (;@18;)
                                                local.get 31
                                                i32.const 2
                                                i32.shl
                                                local.set 1
                                                block  ;; label = @23
                                                  block  ;; label = @24
                                                    loop  ;; label = @25
                                                      local.get 1
                                                      i32.eqz
                                                      br_if 1 (;@24;)
                                                      i32.const -1
                                                      local.get 20
                                                      local.get 1
                                                      i32.add
                                                      i32.load
                                                      local.tee 9
                                                      local.get 1
                                                      i32.const -4
                                                      i32.add
                                                      local.tee 1
                                                      local.get 5
                                                      i32.const 8
                                                      i32.add
                                                      i32.add
                                                      i32.load
                                                      local.tee 11
                                                      i32.ne
                                                      local.get 9
                                                      local.get 11
                                                      i32.gt_u
                                                      select
                                                      local.tee 9
                                                      i32.eqz
                                                      br_if 0 (;@25;)
                                                      br 2 (;@23;)
                                                    end
                                                  end
                                                  i32.const -1
                                                  i32.const 0
                                                  local.get 1
                                                  select
                                                  local.set 9
                                                end
                                                i32.const 0
                                                local.set 32
                                                block  ;; label = @23
                                                  local.get 9
                                                  i32.const 2
                                                  i32.ge_u
                                                  br_if 0 (;@23;)
                                                  block  ;; label = @24
                                                    local.get 31
                                                    i32.eqz
                                                    br_if 0 (;@24;)
                                                    i32.const 1
                                                    local.set 15
                                                    local.get 31
                                                    i32.const 1
                                                    i32.and
                                                    local.set 32
                                                    i32.const 0
                                                    local.set 14
                                                    block  ;; label = @25
                                                      local.get 31
                                                      i32.const 1
                                                      i32.eq
                                                      br_if 0 (;@25;)
                                                      local.get 31
                                                      i32.const -2
                                                      i32.and
                                                      local.set 33
                                                      i32.const 0
                                                      local.set 14
                                                      i32.const 1
                                                      local.set 15
                                                      local.get 5
                                                      i32.const 8
                                                      i32.add
                                                      local.set 1
                                                      local.get 5
                                                      i32.const 680
                                                      i32.add
                                                      local.set 9
                                                      loop  ;; label = @26
                                                        local.get 1
                                                        local.get 1
                                                        i32.load
                                                        local.tee 16
                                                        local.get 9
                                                        i32.load
                                                        i32.const -1
                                                        i32.xor
                                                        i32.add
                                                        local.tee 11
                                                        local.get 15
                                                        i32.const 1
                                                        i32.and
                                                        i32.add
                                                        local.tee 4
                                                        i32.store
                                                        local.get 1
                                                        i32.const 4
                                                        i32.add
                                                        local.tee 15
                                                        local.get 15
                                                        i32.load
                                                        local.tee 18
                                                        local.get 9
                                                        i32.const 4
                                                        i32.add
                                                        i32.load
                                                        i32.const -1
                                                        i32.xor
                                                        i32.add
                                                        local.tee 15
                                                        local.get 11
                                                        local.get 16
                                                        i32.lt_u
                                                        local.get 4
                                                        local.get 11
                                                        i32.lt_u
                                                        i32.or
                                                        i32.add
                                                        local.tee 11
                                                        i32.store
                                                        local.get 15
                                                        local.get 18
                                                        i32.lt_u
                                                        local.get 11
                                                        local.get 15
                                                        i32.lt_u
                                                        i32.or
                                                        local.set 15
                                                        local.get 9
                                                        i32.const 8
                                                        i32.add
                                                        local.set 9
                                                        local.get 1
                                                        i32.const 8
                                                        i32.add
                                                        local.set 1
                                                        local.get 33
                                                        local.get 14
                                                        i32.const 2
                                                        i32.add
                                                        local.tee 14
                                                        i32.ne
                                                        br_if 0 (;@26;)
                                                      end
                                                    end
                                                    block  ;; label = @25
                                                      local.get 32
                                                      i32.eqz
                                                      br_if 0 (;@25;)
                                                      local.get 5
                                                      i32.const 8
                                                      i32.add
                                                      local.get 14
                                                      i32.const 2
                                                      i32.shl
                                                      local.tee 1
                                                      i32.add
                                                      local.tee 9
                                                      local.get 9
                                                      i32.load
                                                      local.tee 9
                                                      local.get 25
                                                      local.get 1
                                                      i32.add
                                                      i32.load
                                                      i32.const -1
                                                      i32.xor
                                                      i32.add
                                                      local.tee 1
                                                      local.get 15
                                                      i32.add
                                                      local.tee 11
                                                      i32.store
                                                      local.get 1
                                                      local.get 9
                                                      i32.lt_u
                                                      local.get 11
                                                      local.get 1
                                                      i32.lt_u
                                                      i32.or
                                                      local.set 15
                                                    end
                                                    local.get 15
                                                    i32.const 1
                                                    i32.and
                                                    i32.eqz
                                                    br_if 9 (;@15;)
                                                  end
                                                  local.get 5
                                                  local.get 31
                                                  i32.store offset=168
                                                  i32.const 8
                                                  local.set 32
                                                  local.get 31
                                                  local.set 14
                                                end
                                                local.get 14
                                                local.get 27
                                                local.get 14
                                                local.get 27
                                                i32.gt_u
                                                select
                                                local.tee 33
                                                i32.const 41
                                                i32.ge_u
                                                br_if 8 (;@14;)
                                                local.get 33
                                                i32.const 2
                                                i32.shl
                                                local.set 1
                                                loop  ;; label = @23
                                                  local.get 1
                                                  i32.eqz
                                                  br_if 2 (;@21;)
                                                  i32.const -1
                                                  local.get 21
                                                  local.get 1
                                                  i32.add
                                                  i32.load
                                                  local.tee 9
                                                  local.get 1
                                                  i32.const -4
                                                  i32.add
                                                  local.tee 1
                                                  local.get 5
                                                  i32.const 8
                                                  i32.add
                                                  i32.add
                                                  i32.load
                                                  local.tee 11
                                                  i32.ne
                                                  local.get 9
                                                  local.get 11
                                                  i32.gt_u
                                                  select
                                                  local.tee 9
                                                  i32.eqz
                                                  br_if 0 (;@23;)
                                                  br 3 (;@20;)
                                                end
                                              end
                                              local.get 13
                                              local.get 30
                                              i32.lt_u
                                              br_if 4 (;@17;)
                                              local.get 13
                                              local.get 3
                                              i32.gt_u
                                              br_if 5 (;@16;)
                                              local.get 13
                                              local.get 30
                                              i32.eq
                                              br_if 20 (;@1;)
                                              local.get 2
                                              local.get 30
                                              i32.add
                                              i32.const 48
                                              local.get 13
                                              local.get 30
                                              i32.sub
                                              call $memset
                                              drop
                                              br 20 (;@1;)
                                            end
                                            i32.const -1
                                            i32.const 0
                                            local.get 1
                                            select
                                            local.set 9
                                          end
                                          block  ;; label = @20
                                            block  ;; label = @21
                                              local.get 9
                                              i32.const 1
                                              i32.le_u
                                              br_if 0 (;@21;)
                                              local.get 14
                                              local.set 33
                                              br 1 (;@20;)
                                            end
                                            block  ;; label = @21
                                              local.get 33
                                              i32.eqz
                                              br_if 0 (;@21;)
                                              i32.const 1
                                              local.set 15
                                              local.get 33
                                              i32.const 1
                                              i32.and
                                              local.set 34
                                              i32.const 0
                                              local.set 14
                                              block  ;; label = @22
                                                local.get 33
                                                i32.const 1
                                                i32.eq
                                                br_if 0 (;@22;)
                                                local.get 33
                                                i32.const -2
                                                i32.and
                                                local.set 31
                                                i32.const 0
                                                local.set 14
                                                i32.const 1
                                                local.set 15
                                                local.get 5
                                                i32.const 8
                                                i32.add
                                                local.set 1
                                                local.get 5
                                                i32.const 512
                                                i32.add
                                                local.set 9
                                                loop  ;; label = @23
                                                  local.get 1
                                                  local.get 1
                                                  i32.load
                                                  local.tee 16
                                                  local.get 9
                                                  i32.load
                                                  i32.const -1
                                                  i32.xor
                                                  i32.add
                                                  local.tee 11
                                                  local.get 15
                                                  i32.const 1
                                                  i32.and
                                                  i32.add
                                                  local.tee 4
                                                  i32.store
                                                  local.get 1
                                                  i32.const 4
                                                  i32.add
                                                  local.tee 15
                                                  local.get 15
                                                  i32.load
                                                  local.tee 18
                                                  local.get 9
                                                  i32.const 4
                                                  i32.add
                                                  i32.load
                                                  i32.const -1
                                                  i32.xor
                                                  i32.add
                                                  local.tee 15
                                                  local.get 11
                                                  local.get 16
                                                  i32.lt_u
                                                  local.get 4
                                                  local.get 11
                                                  i32.lt_u
                                                  i32.or
                                                  i32.add
                                                  local.tee 11
                                                  i32.store
                                                  local.get 15
                                                  local.get 18
                                                  i32.lt_u
                                                  local.get 11
                                                  local.get 15
                                                  i32.lt_u
                                                  i32.or
                                                  local.set 15
                                                  local.get 9
                                                  i32.const 8
                                                  i32.add
                                                  local.set 9
                                                  local.get 1
                                                  i32.const 8
                                                  i32.add
                                                  local.set 1
                                                  local.get 31
                                                  local.get 14
                                                  i32.const 2
                                                  i32.add
                                                  local.tee 14
                                                  i32.ne
                                                  br_if 0 (;@23;)
                                                end
                                              end
                                              block  ;; label = @22
                                                local.get 34
                                                i32.eqz
                                                br_if 0 (;@22;)
                                                local.get 5
                                                i32.const 8
                                                i32.add
                                                local.get 14
                                                i32.const 2
                                                i32.shl
                                                local.tee 1
                                                i32.add
                                                local.tee 9
                                                local.get 9
                                                i32.load
                                                local.tee 9
                                                local.get 24
                                                local.get 1
                                                i32.add
                                                i32.load
                                                i32.const -1
                                                i32.xor
                                                i32.add
                                                local.tee 1
                                                local.get 15
                                                i32.add
                                                local.tee 11
                                                i32.store
                                                local.get 1
                                                local.get 9
                                                i32.lt_u
                                                local.get 11
                                                local.get 1
                                                i32.lt_u
                                                i32.or
                                                local.set 15
                                              end
                                              local.get 15
                                              i32.const 1
                                              i32.and
                                              i32.eqz
                                              br_if 8 (;@13;)
                                            end
                                            local.get 5
                                            local.get 33
                                            i32.store offset=168
                                            local.get 32
                                            i32.const 4
                                            i32.or
                                            local.set 32
                                          end
                                          local.get 33
                                          local.get 26
                                          local.get 33
                                          local.get 26
                                          i32.gt_u
                                          select
                                          local.tee 31
                                          i32.const 41
                                          i32.ge_u
                                          br_if 7 (;@12;)
                                          local.get 31
                                          i32.const 2
                                          i32.shl
                                          local.set 1
                                          block  ;; label = @20
                                            block  ;; label = @21
                                              loop  ;; label = @22
                                                local.get 1
                                                i32.eqz
                                                br_if 1 (;@21;)
                                                i32.const -1
                                                local.get 19
                                                local.get 1
                                                i32.add
                                                i32.load
                                                local.tee 9
                                                local.get 1
                                                i32.const -4
                                                i32.add
                                                local.tee 1
                                                local.get 5
                                                i32.const 8
                                                i32.add
                                                i32.add
                                                i32.load
                                                local.tee 11
                                                i32.ne
                                                local.get 9
                                                local.get 11
                                                i32.gt_u
                                                select
                                                local.tee 9
                                                i32.eqz
                                                br_if 0 (;@22;)
                                                br 2 (;@20;)
                                              end
                                            end
                                            i32.const -1
                                            i32.const 0
                                            local.get 1
                                            select
                                            local.set 9
                                          end
                                          block  ;; label = @20
                                            block  ;; label = @21
                                              local.get 9
                                              i32.const 1
                                              i32.le_u
                                              br_if 0 (;@21;)
                                              local.get 33
                                              local.set 31
                                              br 1 (;@20;)
                                            end
                                            block  ;; label = @21
                                              local.get 31
                                              i32.eqz
                                              br_if 0 (;@21;)
                                              i32.const 1
                                              local.set 15
                                              local.get 31
                                              i32.const 1
                                              i32.and
                                              local.set 34
                                              i32.const 0
                                              local.set 14
                                              block  ;; label = @22
                                                local.get 31
                                                i32.const 1
                                                i32.eq
                                                br_if 0 (;@22;)
                                                local.get 31
                                                i32.const -2
                                                i32.and
                                                local.set 33
                                                i32.const 0
                                                local.set 14
                                                i32.const 1
                                                local.set 15
                                                local.get 5
                                                i32.const 8
                                                i32.add
                                                local.set 1
                                                local.get 5
                                                i32.const 344
                                                i32.add
                                                local.set 9
                                                loop  ;; label = @23
                                                  local.get 1
                                                  local.get 1
                                                  i32.load
                                                  local.tee 16
                                                  local.get 9
                                                  i32.load
                                                  i32.const -1
                                                  i32.xor
                                                  i32.add
                                                  local.tee 11
                                                  local.get 15
                                                  i32.const 1
                                                  i32.and
                                                  i32.add
                                                  local.tee 4
                                                  i32.store
                                                  local.get 1
                                                  i32.const 4
                                                  i32.add
                                                  local.tee 15
                                                  local.get 15
                                                  i32.load
                                                  local.tee 18
                                                  local.get 9
                                                  i32.const 4
                                                  i32.add
                                                  i32.load
                                                  i32.const -1
                                                  i32.xor
                                                  i32.add
                                                  local.tee 15
                                                  local.get 11
                                                  local.get 16
                                                  i32.lt_u
                                                  local.get 4
                                                  local.get 11
                                                  i32.lt_u
                                                  i32.or
                                                  i32.add
                                                  local.tee 11
                                                  i32.store
                                                  local.get 15
                                                  local.get 18
                                                  i32.lt_u
                                                  local.get 11
                                                  local.get 15
                                                  i32.lt_u
                                                  i32.or
                                                  local.set 15
                                                  local.get 9
                                                  i32.const 8
                                                  i32.add
                                                  local.set 9
                                                  local.get 1
                                                  i32.const 8
                                                  i32.add
                                                  local.set 1
                                                  local.get 33
                                                  local.get 14
                                                  i32.const 2
                                                  i32.add
                                                  local.tee 14
                                                  i32.ne
                                                  br_if 0 (;@23;)
                                                end
                                              end
                                              block  ;; label = @22
                                                local.get 34
                                                i32.eqz
                                                br_if 0 (;@22;)
                                                local.get 5
                                                i32.const 8
                                                i32.add
                                                local.get 14
                                                i32.const 2
                                                i32.shl
                                                local.tee 1
                                                i32.add
                                                local.tee 9
                                                local.get 9
                                                i32.load
                                                local.tee 9
                                                local.get 23
                                                local.get 1
                                                i32.add
                                                i32.load
                                                i32.const -1
                                                i32.xor
                                                i32.add
                                                local.tee 1
                                                local.get 15
                                                i32.add
                                                local.tee 11
                                                i32.store
                                                local.get 1
                                                local.get 9
                                                i32.lt_u
                                                local.get 11
                                                local.get 1
                                                i32.lt_u
                                                i32.or
                                                local.set 15
                                              end
                                              local.get 15
                                              i32.const 1
                                              i32.and
                                              i32.eqz
                                              br_if 10 (;@11;)
                                            end
                                            local.get 5
                                            local.get 31
                                            i32.store offset=168
                                            local.get 32
                                            i32.const 2
                                            i32.add
                                            local.set 32
                                          end
                                          local.get 31
                                          local.get 12
                                          local.get 31
                                          local.get 12
                                          i32.gt_u
                                          select
                                          local.tee 14
                                          i32.const 41
                                          i32.ge_u
                                          br_if 9 (;@10;)
                                          local.get 14
                                          i32.const 2
                                          i32.shl
                                          local.set 1
                                          block  ;; label = @20
                                            block  ;; label = @21
                                              loop  ;; label = @22
                                                local.get 1
                                                i32.eqz
                                                br_if 1 (;@21;)
                                                i32.const -1
                                                local.get 17
                                                local.get 1
                                                i32.add
                                                i32.load
                                                local.tee 9
                                                local.get 1
                                                i32.const -4
                                                i32.add
                                                local.tee 1
                                                local.get 5
                                                i32.const 8
                                                i32.add
                                                i32.add
                                                i32.load
                                                local.tee 11
                                                i32.ne
                                                local.get 9
                                                local.get 11
                                                i32.gt_u
                                                select
                                                local.tee 9
                                                i32.eqz
                                                br_if 0 (;@22;)
                                                br 2 (;@20;)
                                              end
                                            end
                                            i32.const -1
                                            i32.const 0
                                            local.get 1
                                            select
                                            local.set 9
                                          end
                                          block  ;; label = @20
                                            block  ;; label = @21
                                              local.get 9
                                              i32.const 1
                                              i32.le_u
                                              br_if 0 (;@21;)
                                              local.get 31
                                              local.set 14
                                              br 1 (;@20;)
                                            end
                                            block  ;; label = @21
                                              local.get 14
                                              i32.eqz
                                              br_if 0 (;@21;)
                                              i32.const 1
                                              local.set 15
                                              local.get 14
                                              i32.const 1
                                              i32.and
                                              local.set 34
                                              i32.const 0
                                              local.set 16
                                              block  ;; label = @22
                                                local.get 14
                                                i32.const 1
                                                i32.eq
                                                br_if 0 (;@22;)
                                                local.get 14
                                                i32.const -2
                                                i32.and
                                                local.set 31
                                                i32.const 0
                                                local.set 16
                                                i32.const 1
                                                local.set 15
                                                local.get 5
                                                i32.const 8
                                                i32.add
                                                local.set 1
                                                local.get 5
                                                i32.const 176
                                                i32.add
                                                local.set 9
                                                loop  ;; label = @23
                                                  local.get 1
                                                  local.get 1
                                                  i32.load
                                                  local.tee 4
                                                  local.get 9
                                                  i32.load
                                                  i32.const -1
                                                  i32.xor
                                                  i32.add
                                                  local.tee 11
                                                  local.get 15
                                                  i32.const 1
                                                  i32.and
                                                  i32.add
                                                  local.tee 18
                                                  i32.store
                                                  local.get 1
                                                  i32.const 4
                                                  i32.add
                                                  local.tee 15
                                                  local.get 15
                                                  i32.load
                                                  local.tee 33
                                                  local.get 9
                                                  i32.const 4
                                                  i32.add
                                                  i32.load
                                                  i32.const -1
                                                  i32.xor
                                                  i32.add
                                                  local.tee 15
                                                  local.get 11
                                                  local.get 4
                                                  i32.lt_u
                                                  local.get 18
                                                  local.get 11
                                                  i32.lt_u
                                                  i32.or
                                                  i32.add
                                                  local.tee 11
                                                  i32.store
                                                  local.get 15
                                                  local.get 33
                                                  i32.lt_u
                                                  local.get 11
                                                  local.get 15
                                                  i32.lt_u
                                                  i32.or
                                                  local.set 15
                                                  local.get 9
                                                  i32.const 8
                                                  i32.add
                                                  local.set 9
                                                  local.get 1
                                                  i32.const 8
                                                  i32.add
                                                  local.set 1
                                                  local.get 31
                                                  local.get 16
                                                  i32.const 2
                                                  i32.add
                                                  local.tee 16
                                                  i32.ne
                                                  br_if 0 (;@23;)
                                                end
                                              end
                                              block  ;; label = @22
                                                local.get 34
                                                i32.eqz
                                                br_if 0 (;@22;)
                                                local.get 5
                                                i32.const 8
                                                i32.add
                                                local.get 16
                                                i32.const 2
                                                i32.shl
                                                local.tee 1
                                                i32.add
                                                local.tee 9
                                                local.get 9
                                                i32.load
                                                local.tee 9
                                                local.get 5
                                                i32.const 176
                                                i32.add
                                                local.get 1
                                                i32.add
                                                i32.load
                                                i32.const -1
                                                i32.xor
                                                i32.add
                                                local.tee 1
                                                local.get 15
                                                i32.add
                                                local.tee 11
                                                i32.store
                                                local.get 1
                                                local.get 9
                                                i32.lt_u
                                                local.get 11
                                                local.get 1
                                                i32.lt_u
                                                i32.or
                                                local.set 15
                                              end
                                              local.get 15
                                              i32.const 1
                                              i32.and
                                              i32.eqz
                                              br_if 12 (;@9;)
                                            end
                                            local.get 5
                                            local.get 14
                                            i32.store offset=168
                                            local.get 32
                                            i32.const 1
                                            i32.add
                                            local.set 32
                                          end
                                          block  ;; label = @20
                                            local.get 30
                                            local.get 3
                                            i32.eq
                                            br_if 0 (;@20;)
                                            local.get 2
                                            local.get 30
                                            i32.add
                                            local.get 32
                                            i32.const 48
                                            i32.add
                                            i32.store8
                                            local.get 14
                                            i32.const 41
                                            i32.ge_u
                                            br_if 12 (;@8;)
                                            block  ;; label = @21
                                              local.get 14
                                              br_if 0 (;@21;)
                                              i32.const 0
                                              local.set 14
                                              br 15 (;@6;)
                                            end
                                            local.get 14
                                            i32.const -1
                                            i32.add
                                            i32.const 1073741823
                                            i32.and
                                            local.tee 1
                                            i32.const 1
                                            i32.add
                                            local.tee 11
                                            i32.const 3
                                            i32.and
                                            local.set 9
                                            block  ;; label = @21
                                              local.get 1
                                              i32.const 3
                                              i32.ge_u
                                              br_if 0 (;@21;)
                                              local.get 5
                                              i32.const 8
                                              i32.add
                                              local.set 1
                                              i64.const 0
                                              local.set 6
                                              br 14 (;@7;)
                                            end
                                            local.get 11
                                            i32.const 2147483644
                                            i32.and
                                            local.set 11
                                            local.get 5
                                            i32.const 8
                                            i32.add
                                            local.set 1
                                            i64.const 0
                                            local.set 6
                                            loop  ;; label = @21
                                              local.get 1
                                              local.get 1
                                              i64.load32_u
                                              i64.const 10
                                              i64.mul
                                              local.get 6
                                              i64.add
                                              local.tee 6
                                              i64.store32
                                              local.get 1
                                              i32.const 4
                                              i32.add
                                              local.tee 15
                                              local.get 15
                                              i64.load32_u
                                              i64.const 10
                                              i64.mul
                                              local.get 6
                                              i64.const 32
                                              i64.shr_u
                                              i64.add
                                              local.tee 6
                                              i64.store32
                                              local.get 1
                                              i32.const 8
                                              i32.add
                                              local.tee 15
                                              local.get 15
                                              i64.load32_u
                                              i64.const 10
                                              i64.mul
                                              local.get 6
                                              i64.const 32
                                              i64.shr_u
                                              i64.add
                                              local.tee 6
                                              i64.store32
                                              local.get 1
                                              i32.const 12
                                              i32.add
                                              local.tee 15
                                              local.get 15
                                              i64.load32_u
                                              i64.const 10
                                              i64.mul
                                              local.get 6
                                              i64.const 32
                                              i64.shr_u
                                              i64.add
                                              local.tee 6
                                              i64.store32
                                              local.get 6
                                              i64.const 32
                                              i64.shr_u
                                              local.set 6
                                              local.get 1
                                              i32.const 16
                                              i32.add
                                              local.set 1
                                              local.get 11
                                              i32.const -4
                                              i32.add
                                              local.tee 11
                                              br_if 0 (;@21;)
                                              br 14 (;@7;)
                                            end
                                          end
                                          local.get 3
                                          local.get 3
                                          i32.const 1050080
                                          call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                                          unreachable
                                        end
                                        local.get 14
                                        i32.const 40
                                        i32.const 1054848
                                        call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                        unreachable
                                      end
                                      local.get 31
                                      i32.const 40
                                      i32.const 1054848
                                      call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                      unreachable
                                    end
                                    local.get 30
                                    local.get 13
                                    i32.const 1050064
                                    call $_ZN4core5slice5index22slice_index_order_fail17hf30e3ed6d5fa1a7bE
                                    unreachable
                                  end
                                  local.get 13
                                  local.get 3
                                  i32.const 1050064
                                  call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                                  unreachable
                                end
                                i32.const 1054864
                                i32.const 26
                                i32.const 1054848
                                call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                unreachable
                              end
                              local.get 33
                              i32.const 40
                              i32.const 1054848
                              call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                              unreachable
                            end
                            i32.const 1054864
                            i32.const 26
                            i32.const 1054848
                            call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                            unreachable
                          end
                          local.get 31
                          i32.const 40
                          i32.const 1054848
                          call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                          unreachable
                        end
                        i32.const 1054864
                        i32.const 26
                        i32.const 1054848
                        call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                        unreachable
                      end
                      local.get 14
                      i32.const 40
                      i32.const 1054848
                      call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                      unreachable
                    end
                    i32.const 1054864
                    i32.const 26
                    i32.const 1054848
                    call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                    unreachable
                  end
                  local.get 14
                  i32.const 40
                  i32.const 1054848
                  call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                  unreachable
                end
                block  ;; label = @7
                  local.get 9
                  i32.eqz
                  br_if 0 (;@7;)
                  loop  ;; label = @8
                    local.get 1
                    local.get 1
                    i64.load32_u
                    i64.const 10
                    i64.mul
                    local.get 6
                    i64.add
                    local.tee 6
                    i64.store32
                    local.get 1
                    i32.const 4
                    i32.add
                    local.set 1
                    local.get 6
                    i64.const 32
                    i64.shr_u
                    local.set 6
                    local.get 9
                    i32.const -1
                    i32.add
                    local.tee 9
                    br_if 0 (;@8;)
                  end
                end
                local.get 6
                i32.wrap_i64
                local.tee 1
                i32.eqz
                br_if 0 (;@6;)
                local.get 14
                i32.const 39
                i32.gt_u
                br_if 2 (;@4;)
                local.get 5
                i32.const 8
                i32.add
                local.get 14
                i32.const 2
                i32.shl
                i32.add
                local.get 1
                i32.store
                local.get 14
                i32.const 1
                i32.add
                local.set 14
              end
              local.get 5
              local.get 14
              i32.store offset=168
              local.get 29
              local.get 13
              i32.ne
              br_if 0 (;@5;)
            end
            i32.const 1
            local.set 14
            br 2 (;@2;)
          end
          local.get 14
          i32.const 40
          i32.const 1054848
          call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
          unreachable
        end
        local.get 16
        i32.const 40
        i32.const 1054848
        call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
        unreachable
      end
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 12
                      i32.const 41
                      i32.ge_u
                      br_if 0 (;@9;)
                      block  ;; label = @10
                        local.get 12
                        br_if 0 (;@10;)
                        i32.const 0
                        local.set 12
                        br 3 (;@7;)
                      end
                      local.get 12
                      i32.const -1
                      i32.add
                      i32.const 1073741823
                      i32.and
                      local.tee 1
                      i32.const 1
                      i32.add
                      local.tee 11
                      i32.const 3
                      i32.and
                      local.set 9
                      block  ;; label = @10
                        local.get 1
                        i32.const 3
                        i32.ge_u
                        br_if 0 (;@10;)
                        local.get 5
                        i32.const 176
                        i32.add
                        local.set 1
                        i64.const 0
                        local.set 6
                        br 2 (;@8;)
                      end
                      local.get 11
                      i32.const 2147483644
                      i32.and
                      local.set 11
                      local.get 5
                      i32.const 176
                      i32.add
                      local.set 1
                      i64.const 0
                      local.set 6
                      loop  ;; label = @10
                        local.get 1
                        local.get 1
                        i64.load32_u
                        i64.const 5
                        i64.mul
                        local.get 6
                        i64.add
                        local.tee 6
                        i64.store32
                        local.get 1
                        i32.const 4
                        i32.add
                        local.tee 15
                        local.get 15
                        i64.load32_u
                        i64.const 5
                        i64.mul
                        local.get 6
                        i64.const 32
                        i64.shr_u
                        i64.add
                        local.tee 6
                        i64.store32
                        local.get 1
                        i32.const 8
                        i32.add
                        local.tee 15
                        local.get 15
                        i64.load32_u
                        i64.const 5
                        i64.mul
                        local.get 6
                        i64.const 32
                        i64.shr_u
                        i64.add
                        local.tee 6
                        i64.store32
                        local.get 1
                        i32.const 12
                        i32.add
                        local.tee 15
                        local.get 15
                        i64.load32_u
                        i64.const 5
                        i64.mul
                        local.get 6
                        i64.const 32
                        i64.shr_u
                        i64.add
                        local.tee 6
                        i64.store32
                        local.get 6
                        i64.const 32
                        i64.shr_u
                        local.set 6
                        local.get 1
                        i32.const 16
                        i32.add
                        local.set 1
                        local.get 11
                        i32.const -4
                        i32.add
                        local.tee 11
                        br_if 0 (;@10;)
                        br 2 (;@8;)
                      end
                    end
                    local.get 12
                    i32.const 40
                    i32.const 1054848
                    call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
                    unreachable
                  end
                  block  ;; label = @8
                    local.get 9
                    i32.eqz
                    br_if 0 (;@8;)
                    loop  ;; label = @9
                      local.get 1
                      local.get 1
                      i64.load32_u
                      i64.const 5
                      i64.mul
                      local.get 6
                      i64.add
                      local.tee 6
                      i64.store32
                      local.get 1
                      i32.const 4
                      i32.add
                      local.set 1
                      local.get 6
                      i64.const 32
                      i64.shr_u
                      local.set 6
                      local.get 9
                      i32.const -1
                      i32.add
                      local.tee 9
                      br_if 0 (;@9;)
                    end
                  end
                  local.get 6
                  i32.wrap_i64
                  local.tee 1
                  i32.eqz
                  br_if 0 (;@7;)
                  local.get 12
                  i32.const 39
                  i32.gt_u
                  br_if 1 (;@6;)
                  local.get 5
                  i32.const 176
                  i32.add
                  local.get 12
                  i32.const 2
                  i32.shl
                  i32.add
                  local.get 1
                  i32.store
                  local.get 12
                  i32.const 1
                  i32.add
                  local.set 12
                end
                local.get 5
                local.get 12
                i32.store offset=336
                local.get 5
                i32.load offset=168
                local.tee 1
                local.get 12
                local.get 1
                local.get 12
                i32.gt_u
                select
                local.tee 1
                i32.const 41
                i32.ge_u
                br_if 1 (;@5;)
                local.get 1
                i32.const 2
                i32.shl
                local.set 1
                block  ;; label = @7
                  block  ;; label = @8
                    loop  ;; label = @9
                      local.get 1
                      i32.eqz
                      br_if 1 (;@8;)
                      i32.const -1
                      local.get 1
                      i32.const -4
                      i32.add
                      local.tee 1
                      local.get 5
                      i32.const 176
                      i32.add
                      i32.add
                      i32.load
                      local.tee 9
                      local.get 1
                      local.get 5
                      i32.const 8
                      i32.add
                      i32.add
                      i32.load
                      local.tee 11
                      i32.ne
                      local.get 9
                      local.get 11
                      i32.gt_u
                      select
                      local.tee 9
                      i32.eqz
                      br_if 0 (;@9;)
                      br 2 (;@7;)
                    end
                  end
                  i32.const -1
                  i32.const 0
                  local.get 1
                  select
                  local.set 9
                end
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 9
                    i32.const 255
                    i32.and
                    br_table 0 (;@8;) 1 (;@7;) 6 (;@2;)
                  end
                  local.get 14
                  i32.eqz
                  br_if 5 (;@2;)
                  local.get 13
                  i32.const -1
                  i32.add
                  local.tee 1
                  local.get 3
                  i32.ge_u
                  br_if 3 (;@4;)
                  local.get 2
                  local.get 1
                  i32.add
                  i32.load8_u
                  i32.const 1
                  i32.and
                  i32.eqz
                  br_if 5 (;@2;)
                end
                local.get 13
                local.get 3
                i32.gt_u
                br_if 3 (;@3;)
                local.get 2
                local.get 13
                i32.add
                local.set 15
                i32.const 0
                local.set 1
                local.get 2
                local.set 9
                block  ;; label = @7
                  loop  ;; label = @8
                    local.get 13
                    local.get 1
                    i32.eq
                    br_if 1 (;@7;)
                    local.get 1
                    i32.const 1
                    i32.add
                    local.set 1
                    local.get 9
                    i32.const -1
                    i32.add
                    local.tee 9
                    local.get 13
                    i32.add
                    local.tee 11
                    i32.load8_u
                    i32.const 57
                    i32.eq
                    br_if 0 (;@8;)
                  end
                  local.get 11
                  local.get 11
                  i32.load8_u
                  i32.const 1
                  i32.add
                  i32.store8
                  local.get 13
                  local.get 1
                  i32.sub
                  i32.const 1
                  i32.add
                  local.get 13
                  i32.ge_u
                  br_if 5 (;@2;)
                  local.get 11
                  i32.const 1
                  i32.add
                  i32.const 48
                  local.get 1
                  i32.const -1
                  i32.add
                  call $memset
                  drop
                  br 5 (;@2;)
                end
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 13
                    br_if 0 (;@8;)
                    i32.const 49
                    local.set 1
                    br 1 (;@7;)
                  end
                  local.get 2
                  i32.const 49
                  i32.store8
                  i32.const 48
                  local.set 1
                  local.get 13
                  i32.const 1
                  i32.eq
                  br_if 0 (;@7;)
                  i32.const 48
                  local.set 1
                  local.get 2
                  i32.const 1
                  i32.add
                  i32.const 48
                  local.get 13
                  i32.const -1
                  i32.add
                  call $memset
                  drop
                end
                local.get 10
                i32.const 1
                i32.add
                i32.extend16_s
                local.tee 10
                local.get 22
                i32.le_s
                br_if 4 (;@2;)
                local.get 13
                local.get 3
                i32.ge_u
                br_if 4 (;@2;)
                local.get 15
                local.get 1
                i32.store8
                local.get 13
                i32.const 1
                i32.add
                local.set 13
                br 4 (;@2;)
              end
              local.get 12
              i32.const 40
              i32.const 1054848
              call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
              unreachable
            end
            local.get 1
            i32.const 40
            i32.const 1054848
            call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
            unreachable
          end
          local.get 1
          local.get 3
          i32.const 1050096
          call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
          unreachable
        end
        local.get 13
        local.get 3
        i32.const 1050112
        call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
        unreachable
      end
      local.get 13
      local.get 3
      i32.le_u
      br_if 0 (;@1;)
      local.get 13
      local.get 3
      i32.const 1050128
      call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
      unreachable
    end
    local.get 0
    local.get 10
    i32.store16 offset=8
    local.get 0
    local.get 13
    i32.store offset=4
    local.get 0
    local.get 2
    i32.store
    local.get 5
    i32.const 848
    i32.add
    global.set $__stack_pointer)
  (func $_ZN4core5slice5index22slice_index_order_fail17hf30e3ed6d5fa1a7bE (type 3) (param i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    call $_ZN4core5slice5index25slice_index_order_fail_rt17hdbe4433d56ef37a3E
    unreachable)
  (func $_ZN4core3num7flt2dec8strategy5grisu19format_shortest_opt17ha5578b8fd1a67f89E (type 12) (param i32 i32 i32 i32)
    (local i32 i64 i64 i64 i64 i32 i32 i64 i32 i64 i64 i64 i64 i64 i64 i64 i64 i64 i64 i64 i64 i64 i64 i64 i64 i64 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 4
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 1
                i64.load
                local.tee 5
                i64.eqz
                br_if 0 (;@6;)
                block  ;; label = @7
                  local.get 1
                  i64.load offset=8
                  local.tee 6
                  i64.eqz
                  br_if 0 (;@7;)
                  block  ;; label = @8
                    local.get 1
                    i64.load offset=16
                    local.tee 7
                    i64.eqz
                    br_if 0 (;@8;)
                    block  ;; label = @9
                      local.get 5
                      local.get 7
                      i64.add
                      local.tee 7
                      local.get 5
                      i64.lt_u
                      br_if 0 (;@9;)
                      block  ;; label = @10
                        local.get 5
                        local.get 6
                        i64.lt_u
                        br_if 0 (;@10;)
                        block  ;; label = @11
                          local.get 3
                          i32.const 17
                          i32.lt_u
                          br_if 0 (;@11;)
                          block  ;; label = @12
                            block  ;; label = @13
                              block  ;; label = @14
                                block  ;; label = @15
                                  block  ;; label = @16
                                    local.get 7
                                    i64.const 2305843009213693951
                                    i64.gt_u
                                    br_if 0 (;@16;)
                                    local.get 4
                                    local.get 1
                                    i32.load16_u offset=24
                                    local.tee 1
                                    i32.store16 offset=8
                                    local.get 4
                                    local.get 5
                                    local.get 6
                                    i64.sub
                                    local.tee 8
                                    i64.store
                                    local.get 1
                                    local.get 1
                                    i32.const -32
                                    i32.add
                                    local.get 1
                                    local.get 7
                                    i64.const 4294967296
                                    i64.lt_u
                                    local.tee 9
                                    select
                                    local.tee 10
                                    i32.const -16
                                    i32.add
                                    local.get 10
                                    local.get 7
                                    i64.const 32
                                    i64.shl
                                    local.get 7
                                    local.get 9
                                    select
                                    local.tee 7
                                    i64.const 281474976710656
                                    i64.lt_u
                                    local.tee 9
                                    select
                                    local.tee 10
                                    i32.const -8
                                    i32.add
                                    local.get 10
                                    local.get 7
                                    i64.const 16
                                    i64.shl
                                    local.get 7
                                    local.get 9
                                    select
                                    local.tee 7
                                    i64.const 72057594037927936
                                    i64.lt_u
                                    local.tee 9
                                    select
                                    local.tee 10
                                    i32.const -4
                                    i32.add
                                    local.get 10
                                    local.get 7
                                    i64.const 8
                                    i64.shl
                                    local.get 7
                                    local.get 9
                                    select
                                    local.tee 7
                                    i64.const 1152921504606846976
                                    i64.lt_u
                                    local.tee 9
                                    select
                                    local.tee 10
                                    i32.const -2
                                    i32.add
                                    local.get 10
                                    local.get 7
                                    i64.const 4
                                    i64.shl
                                    local.get 7
                                    local.get 9
                                    select
                                    local.tee 7
                                    i64.const 4611686018427387904
                                    i64.lt_u
                                    local.tee 9
                                    select
                                    local.get 7
                                    i64.const 2
                                    i64.shl
                                    local.get 7
                                    local.get 9
                                    select
                                    local.tee 11
                                    i64.const -1
                                    i64.gt_s
                                    local.tee 12
                                    i32.sub
                                    local.tee 9
                                    i32.sub
                                    i32.extend16_s
                                    local.tee 10
                                    i32.const 0
                                    i32.lt_s
                                    br_if 4 (;@12;)
                                    local.get 4
                                    i64.const -1
                                    local.get 10
                                    i64.extend_i32_u
                                    local.tee 13
                                    i64.shr_u
                                    local.tee 7
                                    local.get 8
                                    i64.and
                                    i64.store offset=16
                                    local.get 8
                                    local.get 7
                                    i64.gt_u
                                    br_if 3 (;@13;)
                                    local.get 4
                                    local.get 1
                                    i32.store16 offset=8
                                    local.get 4
                                    local.get 5
                                    i64.store
                                    local.get 4
                                    local.get 7
                                    local.get 5
                                    i64.and
                                    i64.store offset=16
                                    local.get 5
                                    local.get 7
                                    i64.gt_u
                                    br_if 2 (;@14;)
                                    i32.const -96
                                    local.get 9
                                    i32.sub
                                    i32.extend16_s
                                    i32.const 80
                                    i32.mul
                                    i32.const 86960
                                    i32.add
                                    i32.const 2126
                                    i32.div_s
                                    local.tee 1
                                    i32.const 81
                                    i32.ge_u
                                    br_if 1 (;@15;)
                                    local.get 1
                                    i32.const 4
                                    i32.shl
                                    local.tee 1
                                    i32.const 1050144
                                    i32.add
                                    i64.load
                                    local.tee 7
                                    i64.const 4294967295
                                    i64.and
                                    local.tee 6
                                    local.get 5
                                    local.get 13
                                    i64.const 63
                                    i64.and
                                    local.tee 13
                                    i64.shl
                                    local.tee 5
                                    i64.const 32
                                    i64.shr_u
                                    local.tee 14
                                    i64.mul
                                    local.tee 15
                                    i64.const 32
                                    i64.shr_u
                                    local.tee 16
                                    local.get 7
                                    i64.const 32
                                    i64.shr_u
                                    local.tee 17
                                    local.get 14
                                    i64.mul
                                    i64.add
                                    local.get 17
                                    local.get 5
                                    i64.const 4294967295
                                    i64.and
                                    local.tee 5
                                    i64.mul
                                    local.tee 7
                                    i64.const 32
                                    i64.shr_u
                                    local.tee 18
                                    i64.add
                                    local.set 19
                                    local.get 15
                                    i64.const 4294967295
                                    i64.and
                                    local.get 6
                                    local.get 5
                                    i64.mul
                                    i64.const 32
                                    i64.shr_u
                                    i64.add
                                    local.get 7
                                    i64.const 4294967295
                                    i64.and
                                    i64.add
                                    i64.const 2147483648
                                    i64.add
                                    i64.const 32
                                    i64.shr_u
                                    local.set 20
                                    i64.const 1
                                    i32.const 0
                                    local.get 9
                                    local.get 1
                                    i32.const 1050152
                                    i32.add
                                    i32.load16_u
                                    i32.add
                                    i32.sub
                                    i32.const 63
                                    i32.and
                                    i64.extend_i32_u
                                    local.tee 7
                                    i64.shl
                                    local.tee 21
                                    i64.const -1
                                    i64.add
                                    local.set 15
                                    local.get 6
                                    local.get 8
                                    local.get 13
                                    i64.shl
                                    local.tee 5
                                    i64.const 32
                                    i64.shr_u
                                    local.tee 8
                                    i64.mul
                                    local.tee 13
                                    i64.const 4294967295
                                    i64.and
                                    local.get 6
                                    local.get 5
                                    i64.const 4294967295
                                    i64.and
                                    local.tee 5
                                    i64.mul
                                    i64.const 32
                                    i64.shr_u
                                    i64.add
                                    local.get 17
                                    local.get 5
                                    i64.mul
                                    local.tee 5
                                    i64.const 4294967295
                                    i64.and
                                    i64.add
                                    i64.const 2147483648
                                    i64.add
                                    i64.const 32
                                    i64.shr_u
                                    local.set 22
                                    local.get 17
                                    local.get 8
                                    i64.mul
                                    local.set 8
                                    local.get 5
                                    i64.const 32
                                    i64.shr_u
                                    local.set 23
                                    local.get 13
                                    i64.const 32
                                    i64.shr_u
                                    local.set 13
                                    local.get 1
                                    i32.const 1050154
                                    i32.add
                                    i32.load16_u
                                    local.set 1
                                    block  ;; label = @17
                                      block  ;; label = @18
                                        block  ;; label = @19
                                          block  ;; label = @20
                                            local.get 17
                                            local.get 11
                                            local.get 12
                                            i64.extend_i32_u
                                            i64.shl
                                            local.tee 5
                                            i64.const 32
                                            i64.shr_u
                                            local.tee 24
                                            i64.mul
                                            local.tee 25
                                            local.get 6
                                            local.get 24
                                            i64.mul
                                            local.tee 11
                                            i64.const 32
                                            i64.shr_u
                                            local.tee 26
                                            i64.add
                                            local.get 17
                                            local.get 5
                                            i64.const 4294967295
                                            i64.and
                                            local.tee 5
                                            i64.mul
                                            local.tee 27
                                            i64.const 32
                                            i64.shr_u
                                            local.tee 28
                                            i64.add
                                            local.get 11
                                            i64.const 4294967295
                                            i64.and
                                            local.get 6
                                            local.get 5
                                            i64.mul
                                            i64.const 32
                                            i64.shr_u
                                            i64.add
                                            local.get 27
                                            i64.const 4294967295
                                            i64.and
                                            i64.add
                                            i64.const 2147483648
                                            i64.add
                                            i64.const 32
                                            i64.shr_u
                                            local.tee 29
                                            i64.add
                                            i64.const 1
                                            i64.add
                                            local.tee 27
                                            local.get 7
                                            i64.shr_u
                                            i32.wrap_i64
                                            local.tee 10
                                            i32.const 10000
                                            i32.lt_u
                                            br_if 0 (;@20;)
                                            local.get 10
                                            i32.const 1000000
                                            i32.lt_u
                                            br_if 1 (;@19;)
                                            local.get 10
                                            i32.const 100000000
                                            i32.lt_u
                                            br_if 2 (;@18;)
                                            i32.const 8
                                            i32.const 9
                                            local.get 10
                                            i32.const 1000000000
                                            i32.lt_u
                                            local.tee 9
                                            select
                                            local.set 30
                                            i32.const 100000000
                                            i32.const 1000000000
                                            local.get 9
                                            select
                                            local.set 9
                                            br 3 (;@17;)
                                          end
                                          block  ;; label = @20
                                            local.get 10
                                            i32.const 100
                                            i32.lt_u
                                            br_if 0 (;@20;)
                                            i32.const 2
                                            i32.const 3
                                            local.get 10
                                            i32.const 1000
                                            i32.lt_u
                                            local.tee 9
                                            select
                                            local.set 30
                                            i32.const 100
                                            i32.const 1000
                                            local.get 9
                                            select
                                            local.set 9
                                            br 3 (;@17;)
                                          end
                                          i32.const 10
                                          i32.const 1
                                          local.get 10
                                          i32.const 9
                                          i32.gt_u
                                          local.tee 30
                                          select
                                          local.set 9
                                          br 2 (;@17;)
                                        end
                                        i32.const 4
                                        i32.const 5
                                        local.get 10
                                        i32.const 100000
                                        i32.lt_u
                                        local.tee 9
                                        select
                                        local.set 30
                                        i32.const 10000
                                        i32.const 100000
                                        local.get 9
                                        select
                                        local.set 9
                                        br 1 (;@17;)
                                      end
                                      i32.const 6
                                      i32.const 7
                                      local.get 10
                                      i32.const 10000000
                                      i32.lt_u
                                      local.tee 9
                                      select
                                      local.set 30
                                      i32.const 1000000
                                      i32.const 10000000
                                      local.get 9
                                      select
                                      local.set 9
                                    end
                                    local.get 19
                                    local.get 20
                                    i64.add
                                    local.set 19
                                    local.get 27
                                    local.get 15
                                    i64.and
                                    local.set 5
                                    local.get 30
                                    local.get 1
                                    i32.sub
                                    i32.const 1
                                    i32.add
                                    local.set 31
                                    local.get 27
                                    local.get 8
                                    local.get 13
                                    i64.add
                                    local.get 23
                                    i64.add
                                    local.get 22
                                    i64.add
                                    local.tee 23
                                    i64.sub
                                    i64.const 1
                                    i64.add
                                    local.tee 13
                                    local.get 15
                                    i64.and
                                    local.set 8
                                    i32.const 0
                                    local.set 1
                                    loop  ;; label = @17
                                      local.get 10
                                      local.get 9
                                      i32.div_u
                                      local.set 12
                                      block  ;; label = @18
                                        block  ;; label = @19
                                          block  ;; label = @20
                                            local.get 3
                                            local.get 1
                                            i32.eq
                                            br_if 0 (;@20;)
                                            local.get 2
                                            local.get 1
                                            i32.add
                                            local.tee 32
                                            local.get 12
                                            i32.const 48
                                            i32.add
                                            local.tee 33
                                            i32.store8
                                            local.get 13
                                            local.get 10
                                            local.get 12
                                            local.get 9
                                            i32.mul
                                            i32.sub
                                            local.tee 10
                                            i64.extend_i32_u
                                            local.get 7
                                            i64.shl
                                            local.tee 6
                                            local.get 5
                                            i64.add
                                            local.tee 11
                                            i64.gt_u
                                            br_if 15 (;@5;)
                                            local.get 30
                                            local.get 1
                                            i32.ne
                                            br_if 2 (;@18;)
                                            local.get 1
                                            i32.const 1
                                            i32.add
                                            local.set 1
                                            i64.const 1
                                            local.set 6
                                            loop  ;; label = @21
                                              local.get 6
                                              local.set 11
                                              local.get 8
                                              local.set 13
                                              local.get 1
                                              local.get 3
                                              i32.ge_u
                                              br_if 2 (;@19;)
                                              local.get 2
                                              local.get 1
                                              i32.add
                                              local.get 5
                                              i64.const 10
                                              i64.mul
                                              local.tee 5
                                              local.get 7
                                              i64.shr_u
                                              i32.wrap_i64
                                              i32.const 48
                                              i32.add
                                              local.tee 9
                                              i32.store8
                                              local.get 1
                                              i32.const 1
                                              i32.add
                                              local.set 1
                                              local.get 11
                                              i64.const 10
                                              i64.mul
                                              local.set 6
                                              local.get 13
                                              i64.const 10
                                              i64.mul
                                              local.tee 8
                                              local.get 5
                                              local.get 15
                                              i64.and
                                              local.tee 5
                                              i64.le_u
                                              br_if 0 (;@21;)
                                            end
                                            local.get 8
                                            local.get 5
                                            i64.sub
                                            local.tee 7
                                            local.get 21
                                            i64.ge_u
                                            local.set 10
                                            local.get 6
                                            local.get 27
                                            local.get 19
                                            i64.sub
                                            i64.mul
                                            local.tee 15
                                            local.get 6
                                            i64.add
                                            local.set 20
                                            local.get 7
                                            local.get 21
                                            i64.lt_u
                                            br_if 16 (;@4;)
                                            local.get 15
                                            local.get 6
                                            i64.sub
                                            local.tee 15
                                            local.get 5
                                            i64.le_u
                                            br_if 16 (;@4;)
                                            local.get 2
                                            local.get 1
                                            i32.add
                                            i32.const -1
                                            i32.add
                                            local.set 12
                                            local.get 13
                                            i64.const 10
                                            i64.mul
                                            local.get 21
                                            local.get 5
                                            i64.add
                                            i64.sub
                                            local.set 17
                                            local.get 21
                                            local.get 15
                                            i64.sub
                                            local.set 27
                                            local.get 15
                                            local.get 5
                                            i64.sub
                                            local.set 14
                                            i64.const 0
                                            local.set 7
                                            loop  ;; label = @21
                                              block  ;; label = @22
                                                local.get 5
                                                local.get 21
                                                i64.add
                                                local.tee 6
                                                local.get 15
                                                i64.lt_u
                                                br_if 0 (;@22;)
                                                local.get 14
                                                local.get 7
                                                i64.add
                                                local.get 27
                                                local.get 5
                                                i64.add
                                                i64.ge_u
                                                br_if 0 (;@22;)
                                                i32.const 1
                                                local.set 10
                                                br 18 (;@4;)
                                              end
                                              local.get 12
                                              local.get 9
                                              i32.const -1
                                              i32.add
                                              local.tee 9
                                              i32.store8
                                              local.get 17
                                              local.get 7
                                              i64.add
                                              local.tee 13
                                              local.get 21
                                              i64.ge_u
                                              local.set 10
                                              local.get 6
                                              local.get 15
                                              i64.ge_u
                                              br_if 18 (;@3;)
                                              local.get 7
                                              local.get 21
                                              i64.sub
                                              local.set 7
                                              local.get 6
                                              local.set 5
                                              local.get 13
                                              local.get 21
                                              i64.ge_u
                                              br_if 0 (;@21;)
                                              br 18 (;@3;)
                                            end
                                          end
                                          local.get 3
                                          local.get 3
                                          i32.const 1051708
                                          call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                                          unreachable
                                        end
                                        local.get 1
                                        local.get 3
                                        i32.const 1051740
                                        call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                                        unreachable
                                      end
                                      local.get 1
                                      i32.const 1
                                      i32.add
                                      local.set 1
                                      local.get 9
                                      i32.const 10
                                      i32.lt_u
                                      local.set 12
                                      local.get 9
                                      i32.const 10
                                      i32.div_u
                                      local.set 9
                                      local.get 12
                                      i32.eqz
                                      br_if 0 (;@17;)
                                    end
                                    i32.const 1051680
                                    i32.const 25
                                    i32.const 1051664
                                    call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                    unreachable
                                  end
                                  i32.const 1051600
                                  i32.const 45
                                  i32.const 1051648
                                  call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                                  unreachable
                                end
                                local.get 1
                                i32.const 81
                                i32.const 1051488
                                call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
                                unreachable
                              end
                              local.get 4
                              i32.const 0
                              i32.store offset=32
                              local.get 4
                              i32.const 16
                              i32.add
                              local.get 4
                              local.get 4
                              i32.const 24
                              i32.add
                              call $_ZN4core9panicking13assert_failed17hca637033118b5d2eE
                              unreachable
                            end
                            local.get 4
                            i32.const 0
                            i32.store offset=32
                            local.get 4
                            i32.const 16
                            i32.add
                            local.get 4
                            local.get 4
                            i32.const 24
                            i32.add
                            call $_ZN4core9panicking13assert_failed17hca637033118b5d2eE
                            unreachable
                          end
                          i32.const 1049135
                          i32.const 29
                          i32.const 1049200
                          call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                          unreachable
                        end
                        i32.const 1049856
                        i32.const 45
                        i32.const 1051584
                        call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                        unreachable
                      end
                      i32.const 1049784
                      i32.const 55
                      i32.const 1051568
                      call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                      unreachable
                    end
                    i32.const 1049712
                    i32.const 54
                    i32.const 1051552
                    call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                    unreachable
                  end
                  i32.const 1049668
                  i32.const 28
                  i32.const 1051536
                  call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                  unreachable
                end
                i32.const 1049620
                i32.const 29
                i32.const 1051520
                call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                unreachable
              end
              i32.const 1049575
              i32.const 28
              i32.const 1051504
              call $_ZN4core9panicking5panic17he2fadc50abd7e245E
              unreachable
            end
            local.get 1
            i32.const 1
            i32.add
            local.set 10
            block  ;; label = @5
              block  ;; label = @6
                local.get 1
                local.get 3
                i32.ge_u
                br_if 0 (;@6;)
                local.get 13
                local.get 11
                i64.sub
                local.tee 15
                local.get 9
                i64.extend_i32_u
                local.get 7
                i64.shl
                local.tee 7
                i64.ge_u
                local.set 1
                local.get 27
                local.get 19
                i64.sub
                local.tee 8
                i64.const 1
                i64.add
                local.set 22
                local.get 8
                i64.const -1
                i64.add
                local.tee 21
                local.get 11
                i64.le_u
                br_if 1 (;@5;)
                local.get 15
                local.get 7
                i64.lt_u
                br_if 1 (;@5;)
                local.get 5
                local.get 7
                i64.add
                local.tee 11
                local.get 16
                i64.add
                local.get 18
                i64.add
                local.get 20
                i64.add
                local.get 17
                local.get 14
                local.get 24
                i64.sub
                i64.mul
                i64.add
                local.get 26
                i64.sub
                local.get 28
                i64.sub
                local.get 29
                i64.sub
                local.set 15
                local.get 26
                local.get 28
                i64.add
                local.get 29
                i64.add
                local.get 25
                i64.add
                local.set 8
                i64.const 0
                local.get 19
                local.get 6
                local.get 5
                i64.add
                i64.add
                i64.sub
                local.set 20
                i64.const 2
                local.get 23
                local.get 11
                local.get 6
                i64.add
                i64.add
                i64.sub
                local.set 14
                loop  ;; label = @7
                  block  ;; label = @8
                    local.get 6
                    local.get 11
                    i64.add
                    local.tee 17
                    local.get 21
                    i64.lt_u
                    br_if 0 (;@8;)
                    local.get 20
                    local.get 8
                    i64.add
                    local.get 6
                    local.get 15
                    i64.add
                    i64.ge_u
                    br_if 0 (;@8;)
                    local.get 6
                    local.get 5
                    i64.add
                    local.set 11
                    i32.const 1
                    local.set 1
                    br 3 (;@5;)
                  end
                  local.get 32
                  local.get 33
                  i32.const -1
                  i32.add
                  local.tee 33
                  i32.store8
                  local.get 5
                  local.get 7
                  i64.add
                  local.set 5
                  local.get 14
                  local.get 8
                  i64.add
                  local.set 27
                  block  ;; label = @8
                    local.get 17
                    local.get 21
                    i64.ge_u
                    br_if 0 (;@8;)
                    local.get 11
                    local.get 7
                    i64.add
                    local.set 11
                    local.get 15
                    local.get 7
                    i64.add
                    local.set 15
                    local.get 8
                    local.get 7
                    i64.sub
                    local.set 8
                    local.get 27
                    local.get 7
                    i64.ge_u
                    br_if 1 (;@7;)
                  end
                end
                local.get 27
                local.get 7
                i64.ge_u
                local.set 1
                local.get 6
                local.get 5
                i64.add
                local.set 11
                br 1 (;@5;)
              end
              local.get 10
              local.get 3
              i32.const 1051724
              call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
              unreachable
            end
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 22
                  local.get 11
                  i64.le_u
                  br_if 0 (;@7;)
                  local.get 1
                  i32.eqz
                  br_if 0 (;@7;)
                  local.get 11
                  local.get 7
                  i64.add
                  local.tee 5
                  local.get 22
                  i64.lt_u
                  br_if 1 (;@6;)
                  local.get 22
                  local.get 11
                  i64.sub
                  local.get 5
                  local.get 22
                  i64.sub
                  i64.ge_u
                  br_if 1 (;@6;)
                end
                block  ;; label = @7
                  local.get 11
                  i64.const 2
                  i64.lt_u
                  br_if 0 (;@7;)
                  local.get 11
                  local.get 13
                  i64.const -4
                  i64.add
                  i64.le_u
                  br_if 2 (;@5;)
                end
                local.get 0
                i32.const 0
                i32.store
                br 5 (;@1;)
              end
              local.get 0
              i32.const 0
              i32.store
              br 4 (;@1;)
            end
            local.get 0
            local.get 31
            i32.store16 offset=8
            local.get 0
            local.get 10
            i32.store offset=4
            br 2 (;@2;)
          end
          local.get 5
          local.set 6
        end
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 20
              local.get 6
              i64.le_u
              br_if 0 (;@5;)
              local.get 10
              i32.eqz
              br_if 0 (;@5;)
              local.get 6
              local.get 21
              i64.add
              local.tee 5
              local.get 20
              i64.lt_u
              br_if 1 (;@4;)
              local.get 20
              local.get 6
              i64.sub
              local.get 5
              local.get 20
              i64.sub
              i64.ge_u
              br_if 1 (;@4;)
            end
            block  ;; label = @5
              local.get 11
              i64.const 20
              i64.mul
              local.get 6
              i64.gt_u
              br_if 0 (;@5;)
              local.get 6
              local.get 11
              i64.const -40
              i64.mul
              local.get 8
              i64.add
              i64.le_u
              br_if 2 (;@3;)
            end
            local.get 0
            i32.const 0
            i32.store
            br 3 (;@1;)
          end
          local.get 0
          i32.const 0
          i32.store
          br 2 (;@1;)
        end
        local.get 0
        local.get 31
        i32.store16 offset=8
        local.get 0
        local.get 1
        i32.store offset=4
      end
      local.get 0
      local.get 2
      i32.store
    end
    local.get 4
    i32.const 48
    i32.add
    global.set $__stack_pointer)
  (func $_ZN4core3num7flt2dec8strategy5grisu16format_exact_opt17ha8db8d0ba5fa1a18E (type 15) (param i32 i32 i32 i32 i32)
    (local i64 i32 i32 i64 i64 i64 i64 i32 i32 i64 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 1
                  i64.load
                  local.tee 5
                  i64.eqz
                  br_if 0 (;@7;)
                  local.get 5
                  i64.const 2305843009213693951
                  i64.gt_u
                  br_if 1 (;@6;)
                  local.get 3
                  i32.eqz
                  br_if 3 (;@4;)
                  i32.const -96
                  local.get 1
                  i32.load16_u offset=24
                  local.tee 1
                  i32.const -32
                  i32.add
                  local.get 1
                  local.get 5
                  i64.const 4294967296
                  i64.lt_u
                  local.tee 6
                  select
                  local.tee 1
                  i32.const -16
                  i32.add
                  local.get 1
                  local.get 5
                  i64.const 32
                  i64.shl
                  local.get 5
                  local.get 6
                  select
                  local.tee 5
                  i64.const 281474976710656
                  i64.lt_u
                  local.tee 6
                  select
                  local.tee 1
                  i32.const -8
                  i32.add
                  local.get 1
                  local.get 5
                  i64.const 16
                  i64.shl
                  local.get 5
                  local.get 6
                  select
                  local.tee 5
                  i64.const 72057594037927936
                  i64.lt_u
                  local.tee 6
                  select
                  local.tee 1
                  i32.const -4
                  i32.add
                  local.get 1
                  local.get 5
                  i64.const 8
                  i64.shl
                  local.get 5
                  local.get 6
                  select
                  local.tee 5
                  i64.const 1152921504606846976
                  i64.lt_u
                  local.tee 6
                  select
                  local.tee 1
                  i32.const -2
                  i32.add
                  local.get 1
                  local.get 5
                  i64.const 4
                  i64.shl
                  local.get 5
                  local.get 6
                  select
                  local.tee 5
                  i64.const 4611686018427387904
                  i64.lt_u
                  local.tee 6
                  select
                  local.get 5
                  i64.const 2
                  i64.shl
                  local.get 5
                  local.get 6
                  select
                  local.tee 5
                  i64.const -1
                  i64.gt_s
                  i32.sub
                  local.tee 6
                  i32.sub
                  i32.extend16_s
                  i32.const 80
                  i32.mul
                  i32.const 86960
                  i32.add
                  i32.const 2126
                  i32.div_s
                  local.tee 1
                  i32.const 81
                  i32.ge_u
                  br_if 2 (;@5;)
                  local.get 1
                  i32.const 4
                  i32.shl
                  local.tee 1
                  i32.const 1050154
                  i32.add
                  i32.load16_u
                  local.set 7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        block  ;; label = @11
                          local.get 1
                          i32.const 1050144
                          i32.add
                          i64.load
                          local.tee 8
                          i64.const 4294967295
                          i64.and
                          local.tee 9
                          local.get 5
                          local.get 5
                          i64.const -1
                          i64.xor
                          i64.const 63
                          i64.shr_u
                          i64.shl
                          local.tee 5
                          i64.const 32
                          i64.shr_u
                          local.tee 10
                          i64.mul
                          local.tee 11
                          i64.const 32
                          i64.shr_u
                          local.get 8
                          i64.const 32
                          i64.shr_u
                          local.tee 8
                          local.get 10
                          i64.mul
                          i64.add
                          local.get 8
                          local.get 5
                          i64.const 4294967295
                          i64.and
                          local.tee 5
                          i64.mul
                          local.tee 8
                          i64.const 32
                          i64.shr_u
                          i64.add
                          local.get 11
                          i64.const 4294967295
                          i64.and
                          local.get 9
                          local.get 5
                          i64.mul
                          i64.const 32
                          i64.shr_u
                          i64.add
                          local.get 8
                          i64.const 4294967295
                          i64.and
                          i64.add
                          i64.const 2147483648
                          i64.add
                          i64.const 32
                          i64.shr_u
                          i64.add
                          local.tee 5
                          i32.const -64
                          local.get 6
                          local.get 1
                          i32.const 1050152
                          i32.add
                          i32.load16_u
                          i32.add
                          i32.sub
                          local.tee 1
                          i32.const 63
                          i32.and
                          i64.extend_i32_u
                          local.tee 9
                          i64.shr_u
                          i32.wrap_i64
                          local.tee 12
                          i32.const 10000
                          i32.lt_u
                          br_if 0 (;@11;)
                          local.get 12
                          i32.const 1000000
                          i32.lt_u
                          br_if 1 (;@10;)
                          local.get 12
                          i32.const 100000000
                          i32.lt_u
                          br_if 2 (;@9;)
                          i32.const 8
                          i32.const 9
                          local.get 12
                          i32.const 1000000000
                          i32.lt_u
                          local.tee 6
                          select
                          local.set 13
                          i32.const 100000000
                          i32.const 1000000000
                          local.get 6
                          select
                          local.set 6
                          br 3 (;@8;)
                        end
                        block  ;; label = @11
                          local.get 12
                          i32.const 100
                          i32.lt_u
                          br_if 0 (;@11;)
                          i32.const 2
                          i32.const 3
                          local.get 12
                          i32.const 1000
                          i32.lt_u
                          local.tee 6
                          select
                          local.set 13
                          i32.const 100
                          i32.const 1000
                          local.get 6
                          select
                          local.set 6
                          br 3 (;@8;)
                        end
                        i32.const 10
                        i32.const 1
                        local.get 12
                        i32.const 9
                        i32.gt_u
                        local.tee 13
                        select
                        local.set 6
                        br 2 (;@8;)
                      end
                      i32.const 4
                      i32.const 5
                      local.get 12
                      i32.const 100000
                      i32.lt_u
                      local.tee 6
                      select
                      local.set 13
                      i32.const 10000
                      i32.const 100000
                      local.get 6
                      select
                      local.set 6
                      br 1 (;@8;)
                    end
                    i32.const 6
                    i32.const 7
                    local.get 12
                    i32.const 10000000
                    i32.lt_u
                    local.tee 6
                    select
                    local.set 13
                    i32.const 1000000
                    i32.const 10000000
                    local.get 6
                    select
                    local.set 6
                  end
                  i64.const 1
                  local.get 9
                  i64.shl
                  local.set 14
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 13
                      local.get 7
                      i32.sub
                      i32.const 1
                      i32.add
                      i32.extend16_s
                      local.tee 15
                      local.get 4
                      i32.extend16_s
                      local.tee 7
                      i32.le_s
                      br_if 0 (;@9;)
                      local.get 5
                      local.get 14
                      i64.const -1
                      i64.add
                      local.tee 10
                      i64.and
                      local.set 8
                      local.get 1
                      i32.const 65535
                      i32.and
                      local.set 16
                      local.get 15
                      local.get 4
                      i32.sub
                      i32.extend16_s
                      local.get 3
                      local.get 15
                      local.get 7
                      i32.sub
                      local.get 3
                      i32.lt_u
                      select
                      local.tee 17
                      i32.const -1
                      i32.add
                      local.set 18
                      i32.const 0
                      local.set 1
                      loop  ;; label = @10
                        local.get 12
                        local.get 6
                        i32.div_u
                        local.set 7
                        local.get 3
                        local.get 1
                        i32.eq
                        br_if 7 (;@3;)
                        local.get 12
                        local.get 7
                        local.get 6
                        i32.mul
                        i32.sub
                        local.set 12
                        local.get 2
                        local.get 1
                        i32.add
                        local.get 7
                        i32.const 48
                        i32.add
                        i32.store8
                        local.get 18
                        local.get 1
                        i32.eq
                        br_if 8 (;@2;)
                        local.get 13
                        local.get 1
                        i32.eq
                        br_if 2 (;@8;)
                        local.get 1
                        i32.const 1
                        i32.add
                        local.set 1
                        local.get 6
                        i32.const 10
                        i32.lt_u
                        local.set 7
                        local.get 6
                        i32.const 10
                        i32.div_u
                        local.set 6
                        local.get 7
                        i32.eqz
                        br_if 0 (;@10;)
                      end
                      i32.const 1051680
                      i32.const 25
                      i32.const 1051916
                      call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                      unreachable
                    end
                    local.get 0
                    local.get 2
                    local.get 3
                    i32.const 0
                    local.get 15
                    local.get 4
                    local.get 5
                    i64.const 10
                    i64.div_u
                    local.get 6
                    i64.extend_i32_u
                    local.get 9
                    i64.shl
                    local.get 14
                    call $_ZN4core3num7flt2dec8strategy5grisu16format_exact_opt14possibly_round17h706057cf26fc1f91E
                    return
                  end
                  local.get 1
                  i32.const 1
                  i32.add
                  local.set 1
                  local.get 16
                  i32.const -1
                  i32.add
                  i32.const 63
                  i32.and
                  i64.extend_i32_u
                  local.set 11
                  i64.const 1
                  local.set 5
                  loop  ;; label = @8
                    block  ;; label = @9
                      local.get 5
                      local.get 11
                      i64.shr_u
                      i64.eqz
                      br_if 0 (;@9;)
                      local.get 0
                      i32.const 0
                      i32.store
                      return
                    end
                    local.get 1
                    local.get 3
                    i32.ge_u
                    br_if 7 (;@1;)
                    local.get 2
                    local.get 1
                    i32.add
                    local.get 8
                    i64.const 10
                    i64.mul
                    local.tee 8
                    local.get 9
                    i64.shr_u
                    i32.wrap_i64
                    i32.const 48
                    i32.add
                    i32.store8
                    local.get 5
                    i64.const 10
                    i64.mul
                    local.set 5
                    local.get 8
                    local.get 10
                    i64.and
                    local.set 8
                    local.get 17
                    local.get 1
                    i32.const 1
                    i32.add
                    local.tee 1
                    i32.ne
                    br_if 0 (;@8;)
                  end
                  local.get 0
                  local.get 2
                  local.get 3
                  local.get 17
                  local.get 15
                  local.get 4
                  local.get 8
                  local.get 14
                  local.get 5
                  call $_ZN4core3num7flt2dec8strategy5grisu16format_exact_opt14possibly_round17h706057cf26fc1f91E
                  return
                end
                i32.const 1049575
                i32.const 28
                i32.const 1051832
                call $_ZN4core9panicking5panic17he2fadc50abd7e245E
                unreachable
              end
              i32.const 1051848
              i32.const 36
              i32.const 1051884
              call $_ZN4core9panicking5panic17he2fadc50abd7e245E
              unreachable
            end
            local.get 1
            i32.const 81
            i32.const 1051488
            call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
            unreachable
          end
          i32.const 1051756
          i32.const 33
          i32.const 1051900
          call $_ZN4core9panicking5panic17he2fadc50abd7e245E
          unreachable
        end
        local.get 3
        local.get 3
        i32.const 1051932
        call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
        unreachable
      end
      local.get 0
      local.get 2
      local.get 3
      local.get 17
      local.get 15
      local.get 4
      local.get 12
      i64.extend_i32_u
      local.get 9
      i64.shl
      local.get 8
      i64.add
      local.get 6
      i64.extend_i32_u
      local.get 9
      i64.shl
      local.get 14
      call $_ZN4core3num7flt2dec8strategy5grisu16format_exact_opt14possibly_round17h706057cf26fc1f91E
      return
    end
    local.get 1
    local.get 3
    i32.const 1051948
    call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
    unreachable)
  (func $_ZN4core3num7flt2dec8strategy5grisu16format_exact_opt14possibly_round17h706057cf26fc1f91E (type 18) (param i32 i32 i32 i32 i32 i32 i64 i64 i64)
    (local i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 7
                    local.get 8
                    i64.le_u
                    br_if 0 (;@8;)
                    local.get 7
                    local.get 8
                    i64.sub
                    local.get 8
                    i64.le_u
                    br_if 1 (;@7;)
                    block  ;; label = @9
                      local.get 7
                      local.get 6
                      i64.sub
                      local.get 6
                      i64.le_u
                      br_if 0 (;@9;)
                      local.get 7
                      local.get 6
                      i64.const 1
                      i64.shl
                      i64.sub
                      local.get 8
                      i64.const 1
                      i64.shl
                      i64.ge_u
                      br_if 3 (;@6;)
                    end
                    block  ;; label = @9
                      local.get 6
                      local.get 8
                      i64.le_u
                      br_if 0 (;@9;)
                      local.get 7
                      local.get 6
                      local.get 8
                      i64.sub
                      local.tee 8
                      i64.sub
                      local.get 8
                      i64.le_u
                      br_if 4 (;@5;)
                    end
                    local.get 0
                    i32.const 0
                    i32.store
                    return
                  end
                  local.get 0
                  i32.const 0
                  i32.store
                  return
                end
                local.get 0
                i32.const 0
                i32.store
                return
              end
              local.get 3
              local.get 2
              i32.gt_u
              br_if 1 (;@4;)
              br 4 (;@1;)
            end
            local.get 3
            local.get 2
            i32.gt_u
            br_if 1 (;@3;)
            local.get 1
            local.get 3
            i32.add
            local.set 9
            i32.const 0
            local.set 10
            local.get 1
            local.set 11
            block  ;; label = @5
              loop  ;; label = @6
                local.get 3
                local.get 10
                i32.eq
                br_if 1 (;@5;)
                local.get 10
                i32.const 1
                i32.add
                local.set 10
                local.get 11
                i32.const -1
                i32.add
                local.tee 11
                local.get 3
                i32.add
                local.tee 12
                i32.load8_u
                i32.const 57
                i32.eq
                br_if 0 (;@6;)
              end
              local.get 12
              local.get 12
              i32.load8_u
              i32.const 1
              i32.add
              i32.store8
              local.get 3
              local.get 10
              i32.sub
              i32.const 1
              i32.add
              local.get 3
              i32.ge_u
              br_if 3 (;@2;)
              local.get 12
              i32.const 1
              i32.add
              i32.const 48
              local.get 10
              i32.const -1
              i32.add
              call $memset
              drop
              br 3 (;@2;)
            end
            block  ;; label = @5
              block  ;; label = @6
                local.get 3
                br_if 0 (;@6;)
                i32.const 49
                local.set 10
                br 1 (;@5;)
              end
              local.get 1
              i32.const 49
              i32.store8
              i32.const 48
              local.set 10
              local.get 3
              i32.const 1
              i32.eq
              br_if 0 (;@5;)
              i32.const 48
              local.set 10
              local.get 1
              i32.const 1
              i32.add
              i32.const 48
              local.get 3
              i32.const -1
              i32.add
              call $memset
              drop
            end
            local.get 4
            i32.const 1
            i32.add
            i32.extend16_s
            local.set 4
            local.get 3
            local.get 2
            i32.ge_u
            br_if 2 (;@2;)
            local.get 4
            local.get 5
            i32.extend16_s
            i32.le_s
            br_if 2 (;@2;)
            local.get 9
            local.get 10
            i32.store8
            local.get 3
            i32.const 1
            i32.add
            local.set 3
            br 2 (;@2;)
          end
          local.get 3
          local.get 2
          i32.const 1051964
          call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
          unreachable
        end
        local.get 3
        local.get 2
        i32.const 1051980
        call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
        unreachable
      end
      local.get 3
      local.get 2
      i32.le_u
      br_if 0 (;@1;)
      local.get 3
      local.get 2
      i32.const 1051996
      call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
      unreachable
    end
    local.get 0
    local.get 4
    i32.store16 offset=8
    local.get 0
    local.get 3
    i32.store offset=4
    local.get 0
    local.get 1
    i32.store)
  (func $_ZN4core3num7flt2dec17digits_to_dec_str17h82b39084dda9b16dE (type 19) (param i32 i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 2
              i32.eqz
              br_if 0 (;@5;)
              local.get 1
              i32.load8_u
              i32.const 49
              i32.lt_u
              br_if 1 (;@4;)
              local.get 6
              i32.const 4
              i32.lt_u
              br_if 2 (;@3;)
              local.get 5
              i32.const 2
              i32.store16
              block  ;; label = @6
                block  ;; label = @7
                  local.get 3
                  i32.extend16_s
                  local.tee 6
                  i32.const 1
                  i32.lt_s
                  br_if 0 (;@7;)
                  local.get 5
                  local.get 1
                  i32.store offset=4
                  local.get 3
                  i32.const 65535
                  i32.and
                  local.tee 3
                  local.get 2
                  i32.ge_u
                  br_if 1 (;@6;)
                  local.get 5
                  i32.const 2
                  i32.store16 offset=24
                  local.get 5
                  i32.const 2
                  i32.store16 offset=12
                  local.get 5
                  local.get 3
                  i32.store offset=8
                  local.get 5
                  i32.const 32
                  i32.add
                  local.get 2
                  local.get 3
                  i32.sub
                  local.tee 2
                  i32.store
                  local.get 5
                  i32.const 28
                  i32.add
                  local.get 1
                  local.get 3
                  i32.add
                  i32.store
                  local.get 5
                  i32.const 20
                  i32.add
                  i32.const 1
                  i32.store
                  local.get 5
                  i32.const 16
                  i32.add
                  i32.const 1052170
                  i32.store
                  i32.const 3
                  local.set 1
                  local.get 2
                  local.get 4
                  i32.ge_u
                  br_if 6 (;@1;)
                  local.get 4
                  local.get 2
                  i32.sub
                  local.set 4
                  br 5 (;@2;)
                end
                local.get 5
                i32.const 2
                i32.store16 offset=24
                local.get 5
                i32.const 0
                i32.store16 offset=12
                local.get 5
                i32.const 2
                i32.store offset=8
                local.get 5
                i32.const 1052168
                i32.store offset=4
                local.get 5
                i32.const 32
                i32.add
                local.get 2
                i32.store
                local.get 5
                i32.const 28
                i32.add
                local.get 1
                i32.store
                local.get 5
                i32.const 16
                i32.add
                i32.const 0
                local.get 6
                i32.sub
                local.tee 3
                i32.store
                i32.const 3
                local.set 1
                local.get 4
                local.get 2
                i32.le_u
                br_if 5 (;@1;)
                local.get 4
                local.get 2
                i32.sub
                local.tee 2
                local.get 3
                i32.le_u
                br_if 5 (;@1;)
                local.get 2
                local.get 6
                i32.add
                local.set 4
                br 4 (;@2;)
              end
              local.get 5
              i32.const 0
              i32.store16 offset=12
              local.get 5
              local.get 2
              i32.store offset=8
              local.get 5
              i32.const 16
              i32.add
              local.get 3
              local.get 2
              i32.sub
              i32.store
              block  ;; label = @6
                local.get 4
                br_if 0 (;@6;)
                i32.const 2
                local.set 1
                br 5 (;@1;)
              end
              local.get 5
              i32.const 2
              i32.store16 offset=24
              local.get 5
              i32.const 32
              i32.add
              i32.const 1
              i32.store
              local.get 5
              i32.const 28
              i32.add
              i32.const 1052170
              i32.store
              br 3 (;@2;)
            end
            i32.const 1051756
            i32.const 33
            i32.const 1052048
            call $_ZN4core9panicking5panic17he2fadc50abd7e245E
            unreachable
          end
          i32.const 1052064
          i32.const 33
          i32.const 1052100
          call $_ZN4core9panicking5panic17he2fadc50abd7e245E
          unreachable
        end
        i32.const 1052116
        i32.const 34
        i32.const 1052152
        call $_ZN4core9panicking5panic17he2fadc50abd7e245E
        unreachable
      end
      local.get 5
      i32.const 0
      i32.store16 offset=36
      local.get 5
      i32.const 40
      i32.add
      local.get 4
      i32.store
      i32.const 4
      local.set 1
    end
    local.get 0
    local.get 1
    i32.store offset=4
    local.get 0
    local.get 5
    i32.store)
  (func $_ZN4core3fmt3num3imp52_$LT$impl$u20$core..fmt..Display$u20$for$u20$u32$GT$3fmt17ha297522bf92cf712E (type 2) (param i32 i32) (result i32)
    local.get 0
    i64.load32_u
    i32.const 1
    local.get 1
    call $_ZN4core3fmt3num3imp7fmt_u6417h5f097ceb54602c9bE)
  (func $_ZN4core3fmt3num50_$LT$impl$u20$core..fmt..Debug$u20$for$u20$u32$GT$3fmt17hf5a8f997cd8038acE (type 2) (param i32 i32) (result i32)
    (local i32 i32 i32)
    global.get $__stack_pointer
    i32.const 128
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 1
              i32.load offset=24
              local.tee 3
              i32.const 16
              i32.and
              br_if 0 (;@5;)
              local.get 3
              i32.const 32
              i32.and
              br_if 1 (;@4;)
              local.get 0
              i64.load32_u
              i32.const 1
              local.get 1
              call $_ZN4core3fmt3num3imp7fmt_u6417h5f097ceb54602c9bE
              local.set 0
              br 4 (;@1;)
            end
            local.get 0
            i32.load
            local.set 0
            i32.const 0
            local.set 3
            loop  ;; label = @5
              local.get 2
              local.get 3
              i32.add
              i32.const 127
              i32.add
              i32.const 48
              i32.const 87
              local.get 0
              i32.const 15
              i32.and
              local.tee 4
              i32.const 10
              i32.lt_u
              select
              local.get 4
              i32.add
              i32.store8
              local.get 3
              i32.const -1
              i32.add
              local.set 3
              local.get 0
              i32.const 15
              i32.gt_u
              local.set 4
              local.get 0
              i32.const 4
              i32.shr_u
              local.set 0
              local.get 4
              br_if 0 (;@5;)
            end
            local.get 3
            i32.const 128
            i32.add
            local.tee 0
            i32.const 129
            i32.ge_u
            br_if 1 (;@3;)
            local.get 1
            i32.const 1
            i32.const 1052552
            i32.const 2
            local.get 2
            local.get 3
            i32.add
            i32.const 128
            i32.add
            i32.const 0
            local.get 3
            i32.sub
            call $_ZN4core3fmt9Formatter12pad_integral17hf0a49a9f5edc1e33E
            local.set 0
            br 3 (;@1;)
          end
          local.get 0
          i32.load
          local.set 0
          i32.const 0
          local.set 3
          loop  ;; label = @4
            local.get 2
            local.get 3
            i32.add
            i32.const 127
            i32.add
            i32.const 48
            i32.const 55
            local.get 0
            i32.const 15
            i32.and
            local.tee 4
            i32.const 10
            i32.lt_u
            select
            local.get 4
            i32.add
            i32.store8
            local.get 3
            i32.const -1
            i32.add
            local.set 3
            local.get 0
            i32.const 15
            i32.gt_u
            local.set 4
            local.get 0
            i32.const 4
            i32.shr_u
            local.set 0
            local.get 4
            br_if 0 (;@4;)
          end
          local.get 3
          i32.const 128
          i32.add
          local.tee 0
          i32.const 129
          i32.ge_u
          br_if 1 (;@2;)
          local.get 1
          i32.const 1
          i32.const 1052552
          i32.const 2
          local.get 2
          local.get 3
          i32.add
          i32.const 128
          i32.add
          i32.const 0
          local.get 3
          i32.sub
          call $_ZN4core3fmt9Formatter12pad_integral17hf0a49a9f5edc1e33E
          local.set 0
          br 2 (;@1;)
        end
        local.get 0
        i32.const 128
        i32.const 1052536
        call $_ZN4core5slice5index26slice_start_index_len_fail17h61955cceb128f310E
        unreachable
      end
      local.get 0
      i32.const 128
      i32.const 1052536
      call $_ZN4core5slice5index26slice_start_index_len_fail17h61955cceb128f310E
      unreachable
    end
    local.get 2
    i32.const 128
    i32.add
    global.set $__stack_pointer
    local.get 0)
  (func $_ZN4core3fmt5write17h7d7e1e065b0161edE (type 1) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    local.get 3
    i32.const 3
    i32.store8 offset=40
    local.get 3
    i64.const 137438953472
    i64.store offset=32
    i32.const 0
    local.set 4
    local.get 3
    i32.const 0
    i32.store offset=24
    local.get 3
    i32.const 0
    i32.store offset=16
    local.get 3
    local.get 1
    i32.store offset=12
    local.get 3
    local.get 0
    i32.store offset=8
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 2
            i32.load
            local.tee 5
            br_if 0 (;@4;)
            local.get 2
            i32.const 20
            i32.add
            i32.load
            local.tee 0
            i32.eqz
            br_if 1 (;@3;)
            local.get 2
            i32.load offset=16
            local.set 1
            local.get 0
            i32.const 3
            i32.shl
            local.set 6
            local.get 0
            i32.const -1
            i32.add
            i32.const 536870911
            i32.and
            i32.const 1
            i32.add
            local.set 4
            local.get 2
            i32.load offset=8
            local.set 0
            loop  ;; label = @5
              block  ;; label = @6
                local.get 0
                i32.const 4
                i32.add
                i32.load
                local.tee 7
                i32.eqz
                br_if 0 (;@6;)
                local.get 3
                i32.load offset=8
                local.get 0
                i32.load
                local.get 7
                local.get 3
                i32.load offset=12
                i32.load offset=12
                call_indirect (type 1)
                br_if 4 (;@2;)
              end
              local.get 1
              i32.load
              local.get 3
              i32.const 8
              i32.add
              local.get 1
              i32.const 4
              i32.add
              i32.load
              call_indirect (type 2)
              br_if 3 (;@2;)
              local.get 1
              i32.const 8
              i32.add
              local.set 1
              local.get 0
              i32.const 8
              i32.add
              local.set 0
              local.get 6
              i32.const -8
              i32.add
              local.tee 6
              br_if 0 (;@5;)
              br 2 (;@3;)
            end
          end
          local.get 2
          i32.load offset=4
          local.tee 1
          i32.eqz
          br_if 0 (;@3;)
          local.get 1
          i32.const 5
          i32.shl
          local.set 8
          local.get 1
          i32.const -1
          i32.add
          i32.const 134217727
          i32.and
          i32.const 1
          i32.add
          local.set 4
          local.get 2
          i32.load offset=8
          local.set 0
          i32.const 0
          local.set 6
          loop  ;; label = @4
            block  ;; label = @5
              local.get 0
              i32.const 4
              i32.add
              i32.load
              local.tee 1
              i32.eqz
              br_if 0 (;@5;)
              local.get 3
              i32.load offset=8
              local.get 0
              i32.load
              local.get 1
              local.get 3
              i32.load offset=12
              i32.load offset=12
              call_indirect (type 1)
              br_if 3 (;@2;)
            end
            local.get 3
            local.get 5
            local.get 6
            i32.add
            local.tee 1
            i32.const 28
            i32.add
            i32.load8_u
            i32.store8 offset=40
            local.get 3
            local.get 1
            i32.const 20
            i32.add
            i64.load align=4
            i64.store offset=32
            local.get 1
            i32.const 16
            i32.add
            i32.load
            local.set 9
            local.get 2
            i32.load offset=16
            local.set 10
            i32.const 0
            local.set 11
            i32.const 0
            local.set 7
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 1
                  i32.const 12
                  i32.add
                  i32.load
                  br_table 1 (;@6;) 0 (;@7;) 2 (;@5;) 1 (;@6;)
                end
                local.get 9
                i32.const 3
                i32.shl
                local.set 12
                i32.const 0
                local.set 7
                local.get 10
                local.get 12
                i32.add
                local.tee 12
                i32.const 4
                i32.add
                i32.load
                i32.const 26
                i32.ne
                br_if 1 (;@5;)
                local.get 12
                i32.load
                i32.load
                local.set 9
              end
              i32.const 1
              local.set 7
            end
            local.get 3
            local.get 9
            i32.store offset=20
            local.get 3
            local.get 7
            i32.store offset=16
            local.get 1
            i32.const 8
            i32.add
            i32.load
            local.set 7
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 1
                  i32.const 4
                  i32.add
                  i32.load
                  br_table 1 (;@6;) 0 (;@7;) 2 (;@5;) 1 (;@6;)
                end
                local.get 7
                i32.const 3
                i32.shl
                local.set 9
                local.get 10
                local.get 9
                i32.add
                local.tee 9
                i32.const 4
                i32.add
                i32.load
                i32.const 26
                i32.ne
                br_if 1 (;@5;)
                local.get 9
                i32.load
                i32.load
                local.set 7
              end
              i32.const 1
              local.set 11
            end
            local.get 3
            local.get 7
            i32.store offset=28
            local.get 3
            local.get 11
            i32.store offset=24
            local.get 10
            local.get 1
            i32.load
            i32.const 3
            i32.shl
            i32.add
            local.tee 1
            i32.load
            local.get 3
            i32.const 8
            i32.add
            local.get 1
            i32.load offset=4
            call_indirect (type 2)
            br_if 2 (;@2;)
            local.get 0
            i32.const 8
            i32.add
            local.set 0
            local.get 8
            local.get 6
            i32.const 32
            i32.add
            local.tee 6
            i32.ne
            br_if 0 (;@4;)
          end
        end
        block  ;; label = @3
          local.get 4
          local.get 2
          i32.const 12
          i32.add
          i32.load
          i32.ge_u
          br_if 0 (;@3;)
          local.get 3
          i32.load offset=8
          local.get 2
          i32.load offset=8
          local.get 4
          i32.const 3
          i32.shl
          i32.add
          local.tee 1
          i32.load
          local.get 1
          i32.load offset=4
          local.get 3
          i32.load offset=12
          i32.load offset=12
          call_indirect (type 1)
          br_if 1 (;@2;)
        end
        i32.const 0
        local.set 1
        br 1 (;@1;)
      end
      i32.const 1
      local.set 1
    end
    local.get 3
    i32.const 48
    i32.add
    global.set $__stack_pointer
    local.get 1)
  (func $_ZN4core3fmt9Formatter12pad_integral17hf0a49a9f5edc1e33E (type 20) (param i32 i32 i32 i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.eqz
        br_if 0 (;@2;)
        i32.const 43
        i32.const 1114112
        local.get 0
        i32.load offset=24
        local.tee 6
        i32.const 1
        i32.and
        local.tee 1
        select
        local.set 7
        local.get 1
        local.get 5
        i32.add
        local.set 8
        br 1 (;@1;)
      end
      local.get 5
      i32.const 1
      i32.add
      local.set 8
      local.get 0
      i32.load offset=24
      local.set 6
      i32.const 45
      local.set 7
    end
    block  ;; label = @1
      block  ;; label = @2
        local.get 6
        i32.const 4
        i32.and
        br_if 0 (;@2;)
        i32.const 0
        local.set 2
        br 1 (;@1;)
      end
      block  ;; label = @2
        block  ;; label = @3
          local.get 3
          i32.const 16
          i32.lt_u
          br_if 0 (;@3;)
          local.get 2
          local.get 3
          call $_ZN4core3str5count14do_count_chars17hffbd23dea83a0c54E
          local.set 9
          br 1 (;@2;)
        end
        block  ;; label = @3
          local.get 3
          br_if 0 (;@3;)
          i32.const 0
          local.set 9
          br 1 (;@2;)
        end
        local.get 3
        i32.const 3
        i32.and
        local.set 10
        block  ;; label = @3
          block  ;; label = @4
            local.get 3
            i32.const 4
            i32.ge_u
            br_if 0 (;@4;)
            i32.const 0
            local.set 9
            local.get 2
            local.set 1
            br 1 (;@3;)
          end
          local.get 3
          i32.const -4
          i32.and
          local.set 11
          i32.const 0
          local.set 9
          local.get 2
          local.set 1
          loop  ;; label = @4
            local.get 9
            local.get 1
            i32.load8_s
            i32.const -65
            i32.gt_s
            i32.add
            local.get 1
            i32.load8_s offset=1
            i32.const -65
            i32.gt_s
            i32.add
            local.get 1
            i32.load8_s offset=2
            i32.const -65
            i32.gt_s
            i32.add
            local.get 1
            i32.load8_s offset=3
            i32.const -65
            i32.gt_s
            i32.add
            local.set 9
            local.get 1
            i32.const 4
            i32.add
            local.set 1
            local.get 11
            i32.const -4
            i32.add
            local.tee 11
            br_if 0 (;@4;)
          end
        end
        local.get 10
        i32.eqz
        br_if 0 (;@2;)
        loop  ;; label = @3
          local.get 9
          local.get 1
          i32.load8_s
          i32.const -65
          i32.gt_s
          i32.add
          local.set 9
          local.get 1
          i32.const 1
          i32.add
          local.set 1
          local.get 10
          i32.const -1
          i32.add
          local.tee 10
          br_if 0 (;@3;)
        end
      end
      local.get 9
      local.get 8
      i32.add
      local.set 8
    end
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i32.load offset=8
        br_if 0 (;@2;)
        i32.const 1
        local.set 1
        local.get 0
        i32.load
        local.tee 9
        local.get 0
        i32.const 4
        i32.add
        i32.load
        local.tee 0
        local.get 7
        local.get 2
        local.get 3
        call $_ZN4core3fmt9Formatter12pad_integral12write_prefix17hbf853012c4e1643dE
        br_if 1 (;@1;)
        local.get 9
        local.get 4
        local.get 5
        local.get 0
        i32.load offset=12
        call_indirect (type 1)
        return
      end
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 0
                i32.const 12
                i32.add
                i32.load
                local.tee 11
                local.get 8
                i32.le_u
                br_if 0 (;@6;)
                local.get 6
                i32.const 8
                i32.and
                br_if 4 (;@2;)
                local.get 11
                local.get 8
                i32.sub
                local.tee 9
                local.set 11
                i32.const 1
                local.get 0
                i32.load8_u offset=32
                local.tee 1
                local.get 1
                i32.const 3
                i32.eq
                select
                local.tee 1
                br_table 3 (;@3;) 1 (;@5;) 2 (;@4;) 3 (;@3;)
              end
              i32.const 1
              local.set 1
              local.get 0
              i32.load
              local.tee 9
              local.get 0
              i32.const 4
              i32.add
              i32.load
              local.tee 0
              local.get 7
              local.get 2
              local.get 3
              call $_ZN4core3fmt9Formatter12pad_integral12write_prefix17hbf853012c4e1643dE
              br_if 4 (;@1;)
              local.get 9
              local.get 4
              local.get 5
              local.get 0
              i32.load offset=12
              call_indirect (type 1)
              return
            end
            i32.const 0
            local.set 11
            local.get 9
            local.set 1
            br 1 (;@3;)
          end
          local.get 9
          i32.const 1
          i32.shr_u
          local.set 1
          local.get 9
          i32.const 1
          i32.add
          i32.const 1
          i32.shr_u
          local.set 11
        end
        local.get 1
        i32.const 1
        i32.add
        local.set 1
        local.get 0
        i32.const 4
        i32.add
        i32.load
        local.set 10
        local.get 0
        i32.load offset=28
        local.set 9
        local.get 0
        i32.load
        local.set 0
        block  ;; label = @3
          loop  ;; label = @4
            local.get 1
            i32.const -1
            i32.add
            local.tee 1
            i32.eqz
            br_if 1 (;@3;)
            local.get 0
            local.get 9
            local.get 10
            i32.load offset=16
            call_indirect (type 2)
            i32.eqz
            br_if 0 (;@4;)
          end
          i32.const 1
          return
        end
        i32.const 1
        local.set 1
        local.get 9
        i32.const 1114112
        i32.eq
        br_if 1 (;@1;)
        local.get 0
        local.get 10
        local.get 7
        local.get 2
        local.get 3
        call $_ZN4core3fmt9Formatter12pad_integral12write_prefix17hbf853012c4e1643dE
        br_if 1 (;@1;)
        local.get 0
        local.get 4
        local.get 5
        local.get 10
        i32.load offset=12
        call_indirect (type 1)
        br_if 1 (;@1;)
        i32.const 0
        local.set 1
        block  ;; label = @3
          loop  ;; label = @4
            block  ;; label = @5
              local.get 11
              local.get 1
              i32.ne
              br_if 0 (;@5;)
              local.get 11
              local.set 1
              br 2 (;@3;)
            end
            local.get 1
            i32.const 1
            i32.add
            local.set 1
            local.get 0
            local.get 9
            local.get 10
            i32.load offset=16
            call_indirect (type 2)
            i32.eqz
            br_if 0 (;@4;)
          end
          local.get 1
          i32.const -1
          i32.add
          local.set 1
        end
        local.get 1
        local.get 11
        i32.lt_u
        local.set 1
        br 1 (;@1;)
      end
      local.get 0
      i32.load offset=28
      local.set 6
      local.get 0
      i32.const 48
      i32.store offset=28
      local.get 0
      i32.load8_u offset=32
      local.set 12
      i32.const 1
      local.set 1
      local.get 0
      i32.const 1
      i32.store8 offset=32
      local.get 0
      i32.load
      local.tee 9
      local.get 0
      i32.const 4
      i32.add
      i32.load
      local.tee 10
      local.get 7
      local.get 2
      local.get 3
      call $_ZN4core3fmt9Formatter12pad_integral12write_prefix17hbf853012c4e1643dE
      br_if 0 (;@1;)
      local.get 11
      local.get 8
      i32.sub
      i32.const 1
      i32.add
      local.set 1
      block  ;; label = @2
        loop  ;; label = @3
          local.get 1
          i32.const -1
          i32.add
          local.tee 1
          i32.eqz
          br_if 1 (;@2;)
          local.get 9
          i32.const 48
          local.get 10
          i32.load offset=16
          call_indirect (type 2)
          i32.eqz
          br_if 0 (;@3;)
        end
        i32.const 1
        return
      end
      i32.const 1
      local.set 1
      local.get 9
      local.get 4
      local.get 5
      local.get 10
      i32.load offset=12
      call_indirect (type 1)
      br_if 0 (;@1;)
      local.get 0
      local.get 12
      i32.store8 offset=32
      local.get 0
      local.get 6
      i32.store offset=28
      i32.const 0
      return
    end
    local.get 1)
  (func $_ZN71_$LT$core..ops..range..Range$LT$Idx$GT$$u20$as$u20$core..fmt..Debug$GT$3fmt17h4317476aca2a8a15E (type 2) (param i32 i32) (result i32)
    (local i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 32
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    i32.const 1
    local.set 3
    block  ;; label = @1
      local.get 0
      local.get 1
      call $_ZN4core3fmt3num50_$LT$impl$u20$core..fmt..Debug$u20$for$u20$u32$GT$3fmt17hf5a8f997cd8038acE
      br_if 0 (;@1;)
      local.get 1
      i32.load offset=4
      local.set 4
      local.get 1
      i32.load
      local.set 5
      local.get 2
      i32.const 0
      i32.store offset=28
      local.get 2
      i32.const 1052176
      i32.store offset=24
      i32.const 1
      local.set 3
      local.get 2
      i32.const 1
      i32.store offset=20
      local.get 2
      i32.const 1052240
      i32.store offset=16
      local.get 2
      i32.const 0
      i32.store offset=8
      local.get 5
      local.get 4
      local.get 2
      i32.const 8
      i32.add
      call $_ZN4core3fmt5write17h7d7e1e065b0161edE
      br_if 0 (;@1;)
      local.get 0
      i32.const 4
      i32.add
      local.get 1
      call $_ZN4core3fmt3num50_$LT$impl$u20$core..fmt..Debug$u20$for$u20$u32$GT$3fmt17hf5a8f997cd8038acE
      local.set 3
    end
    local.get 2
    i32.const 32
    i32.add
    global.set $__stack_pointer
    local.get 3)
  (func $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17ha7f33678c3328e3eE (type 10) (param i32) (result i64)
    i64.const 726813373705916399)
  (func $_ZN4core6result13unwrap_failed17h80851617ed2e05abE (type 15) (param i32 i32 i32 i32 i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 64
    i32.sub
    local.tee 5
    global.set $__stack_pointer
    local.get 5
    local.get 1
    i32.store offset=12
    local.get 5
    local.get 0
    i32.store offset=8
    local.get 5
    local.get 3
    i32.store offset=20
    local.get 5
    local.get 2
    i32.store offset=16
    local.get 5
    i32.const 24
    i32.add
    i32.const 12
    i32.add
    i32.const 2
    i32.store
    local.get 5
    i32.const 44
    i32.add
    i32.const 2
    i32.store
    local.get 5
    i32.const 48
    i32.add
    i32.const 12
    i32.add
    i32.const 27
    i32.store
    local.get 5
    i32.const 1052492
    i32.store offset=32
    local.get 5
    i32.const 0
    i32.store offset=24
    local.get 5
    i32.const 28
    i32.store offset=52
    local.get 5
    local.get 5
    i32.const 48
    i32.add
    i32.store offset=40
    local.get 5
    local.get 5
    i32.const 16
    i32.add
    i32.store offset=56
    local.get 5
    local.get 5
    i32.const 8
    i32.add
    i32.store offset=48
    local.get 5
    i32.const 24
    i32.add
    local.get 4
    call $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E
    unreachable)
  (func $_ZN44_$LT$$RF$T$u20$as$u20$core..fmt..Display$GT$3fmt17h1d04e84ea8294252E (type 2) (param i32 i32) (result i32)
    local.get 1
    local.get 0
    i32.load
    local.get 0
    i32.load offset=4
    call $_ZN4core3fmt9Formatter3pad17h7e3d356fb87d6820E)
  (func $_ZN4core5panic10panic_info9PanicInfo7message17hdce92b3e180291bdE (type 13) (param i32) (result i32)
    local.get 0
    i32.load offset=8)
  (func $_ZN4core5panic10panic_info9PanicInfo8location17h669fabfe65fb79d4E (type 13) (param i32) (result i32)
    local.get 0
    i32.load offset=12)
  (func $_ZN4core5panic10panic_info9PanicInfo10can_unwind17hd9c7a52d1ea179a2E (type 13) (param i32) (result i32)
    local.get 0
    i32.load8_u offset=16)
  (func $_ZN4core9panicking19assert_failed_inner17h6bb952400e68e912E (type 19) (param i32 i32 i32 i32 i32 i32 i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 112
    i32.sub
    local.tee 7
    global.set $__stack_pointer
    local.get 7
    local.get 2
    i32.store offset=12
    local.get 7
    local.get 1
    i32.store offset=8
    local.get 7
    local.get 4
    i32.store offset=20
    local.get 7
    local.get 3
    i32.store offset=16
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 0
            i32.const 255
            i32.and
            br_table 0 (;@4;) 1 (;@3;) 2 (;@2;) 0 (;@4;)
          end
          local.get 7
          i32.const 1052357
          i32.store offset=24
          i32.const 2
          local.set 2
          br 2 (;@1;)
        end
        local.get 7
        i32.const 1052355
        i32.store offset=24
        i32.const 2
        local.set 2
        br 1 (;@1;)
      end
      local.get 7
      i32.const 1052348
      i32.store offset=24
      i32.const 7
      local.set 2
    end
    local.get 7
    local.get 2
    i32.store offset=28
    block  ;; label = @1
      local.get 5
      i32.load offset=8
      br_if 0 (;@1;)
      local.get 7
      i32.const 56
      i32.add
      i32.const 20
      i32.add
      i32.const 27
      i32.store
      local.get 7
      i32.const 56
      i32.add
      i32.const 12
      i32.add
      i32.const 27
      i32.store
      local.get 7
      i32.const 88
      i32.add
      i32.const 12
      i32.add
      i32.const 4
      i32.store
      local.get 7
      i32.const 88
      i32.add
      i32.const 20
      i32.add
      i32.const 3
      i32.store
      local.get 7
      i32.const 1052456
      i32.store offset=96
      local.get 7
      i32.const 0
      i32.store offset=88
      local.get 7
      i32.const 28
      i32.store offset=60
      local.get 7
      local.get 7
      i32.const 56
      i32.add
      i32.store offset=104
      local.get 7
      local.get 7
      i32.const 16
      i32.add
      i32.store offset=72
      local.get 7
      local.get 7
      i32.const 8
      i32.add
      i32.store offset=64
      local.get 7
      local.get 7
      i32.const 24
      i32.add
      i32.store offset=56
      local.get 7
      i32.const 88
      i32.add
      local.get 6
      call $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E
      unreachable
    end
    local.get 7
    i32.const 32
    i32.add
    i32.const 16
    i32.add
    local.get 5
    i32.const 16
    i32.add
    i64.load align=4
    i64.store
    local.get 7
    i32.const 32
    i32.add
    i32.const 8
    i32.add
    local.get 5
    i32.const 8
    i32.add
    i64.load align=4
    i64.store
    local.get 7
    local.get 5
    i64.load align=4
    i64.store offset=32
    local.get 7
    i32.const 88
    i32.add
    i32.const 12
    i32.add
    i32.const 4
    i32.store
    local.get 7
    i32.const 88
    i32.add
    i32.const 20
    i32.add
    i32.const 4
    i32.store
    local.get 7
    i32.const 84
    i32.add
    i32.const 29
    i32.store
    local.get 7
    i32.const 56
    i32.add
    i32.const 20
    i32.add
    i32.const 27
    i32.store
    local.get 7
    i32.const 56
    i32.add
    i32.const 12
    i32.add
    i32.const 27
    i32.store
    local.get 7
    i32.const 1052420
    i32.store offset=96
    local.get 7
    i32.const 0
    i32.store offset=88
    local.get 7
    i32.const 28
    i32.store offset=60
    local.get 7
    local.get 7
    i32.const 56
    i32.add
    i32.store offset=104
    local.get 7
    local.get 7
    i32.const 32
    i32.add
    i32.store offset=80
    local.get 7
    local.get 7
    i32.const 16
    i32.add
    i32.store offset=72
    local.get 7
    local.get 7
    i32.const 8
    i32.add
    i32.store offset=64
    local.get 7
    local.get 7
    i32.const 24
    i32.add
    i32.store offset=56
    local.get 7
    i32.const 88
    i32.add
    local.get 6
    call $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E
    unreachable)
  (func $_ZN42_$LT$$RF$T$u20$as$u20$core..fmt..Debug$GT$3fmt17h3fb2133f2241bdd8E (type 2) (param i32 i32) (result i32)
    local.get 0
    i32.load
    local.get 1
    local.get 0
    i32.load offset=4
    i32.load offset=12
    call_indirect (type 2))
  (func $_ZN59_$LT$core..fmt..Arguments$u20$as$u20$core..fmt..Display$GT$3fmt17hf42b4cccce915088E (type 2) (param i32 i32) (result i32)
    local.get 1
    i32.load
    local.get 1
    i32.load offset=4
    local.get 0
    call $_ZN4core3fmt5write17h7d7e1e065b0161edE)
  (func $_ZN4core3fmt5float29float_to_decimal_common_exact17h887dd65cb992c326E (type 21) (param i32 f32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i64 i64 i32 i64)
    global.get $__stack_pointer
    i32.const 1136
    i32.sub
    local.tee 4
    global.set $__stack_pointer
    local.get 1
    i32.reinterpret_f32
    local.set 5
    i32.const 3
    local.set 6
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        f32.const inf (;=inf;)
        f32.ne
        br_if 0 (;@2;)
        br 1 (;@1;)
      end
      local.get 1
      f32.const -inf (;=-inf;)
      f32.eq
      br_if 0 (;@1;)
      block  ;; label = @2
        local.get 1
        local.get 1
        f32.eq
        br_if 0 (;@2;)
        i32.const 2
        local.set 6
        br 1 (;@1;)
      end
      local.get 5
      i32.const 8388607
      i32.and
      i32.const 8388608
      i32.or
      local.get 5
      i32.const 1
      i32.shl
      i32.const 16777214
      i32.and
      local.get 5
      i32.const 23
      i32.shr_u
      i32.const 255
      i32.and
      local.tee 7
      select
      local.tee 8
      i64.extend_i32_u
      local.tee 9
      i64.const 1
      i64.and
      local.set 10
      i32.const 4
      local.set 6
      block  ;; label = @2
        block  ;; label = @3
          i32.const 4
          i32.const 3
          local.get 5
          i32.const 2139095040
          i32.and
          select
          i32.const 2
          local.get 5
          i32.const 2147483647
          i32.and
          select
          i32.const -2
          i32.add
          br_table 2 (;@1;) 0 (;@3;) 1 (;@2;) 2 (;@1;)
        end
        local.get 7
        i32.const -150
        i32.add
        local.set 11
        local.get 10
        i64.eqz
        local.set 6
        i64.const 1
        local.set 12
        br 1 (;@1;)
      end
      i64.const 33554432
      local.get 9
      i64.const 1
      i64.shl
      local.get 8
      i32.const 8388608
      i32.eq
      local.tee 6
      select
      local.set 9
      i64.const 2
      i64.const 1
      local.get 6
      select
      local.set 12
      i32.const -152
      i32.const -151
      local.get 6
      select
      local.get 7
      i32.add
      local.set 11
      local.get 10
      i64.eqz
      local.set 6
    end
    local.get 4
    local.get 11
    i32.store16 offset=1128
    local.get 4
    local.get 12
    i64.store offset=1120
    local.get 4
    i64.const 1
    i64.store offset=1112
    local.get 4
    local.get 9
    i64.store offset=1104
    local.get 4
    local.get 6
    i32.store8 offset=1130
    block  ;; label = @1
      block  ;; label = @2
        local.get 6
        i32.const 2
        i32.ne
        br_if 0 (;@2;)
        i32.const 1052176
        local.set 2
        i32.const 0
        local.set 7
        br 1 (;@1;)
      end
      block  ;; label = @2
        local.get 2
        br_if 0 (;@2;)
        i32.const 1052171
        i32.const 1052176
        local.get 5
        i32.const 0
        i32.lt_s
        select
        local.set 2
        local.get 5
        i32.const 31
        i32.shr_u
        local.set 7
        br 1 (;@1;)
      end
      i32.const 1052171
      i32.const 1052172
      local.get 5
      i32.const 0
      i32.lt_s
      select
      local.set 2
      i32.const 1
      local.set 7
    end
    i32.const 1
    local.set 5
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 6
                i32.const -2
                i32.add
                i32.const 3
                local.get 6
                i32.const 1
                i32.gt_u
                select
                i32.const 255
                i32.and
                br_table 3 (;@3;) 2 (;@4;) 1 (;@5;) 0 (;@6;) 3 (;@3;)
              end
              i32.const -12
              i32.const 5
              local.get 11
              i32.extend16_s
              local.tee 6
              i32.const 0
              i32.lt_s
              select
              local.get 6
              i32.mul
              local.tee 6
              i32.const 16063
              i32.gt_u
              br_if 4 (;@1;)
              local.get 4
              i32.const 1040
              i32.add
              local.get 4
              i32.const 1104
              i32.add
              local.get 4
              i32.const 16
              i32.add
              local.get 6
              i32.const 4
              i32.shr_u
              i32.const 21
              i32.add
              local.tee 5
              i32.const 0
              local.get 3
              i32.sub
              i32.const -32768
              local.get 3
              i32.const 32768
              i32.lt_u
              select
              local.tee 6
              call $_ZN4core3num7flt2dec8strategy5grisu16format_exact_opt17ha8db8d0ba5fa1a18E
              local.get 6
              i32.extend16_s
              local.set 6
              block  ;; label = @6
                block  ;; label = @7
                  local.get 4
                  i32.load offset=1040
                  br_if 0 (;@7;)
                  local.get 4
                  i32.const 1088
                  i32.add
                  local.get 4
                  i32.const 1104
                  i32.add
                  local.get 4
                  i32.const 16
                  i32.add
                  local.get 5
                  local.get 6
                  call $_ZN4core3num7flt2dec8strategy6dragon12format_exact17h6d7faa24131201fdE
                  br 1 (;@6;)
                end
                local.get 4
                i32.const 1088
                i32.add
                i32.const 8
                i32.add
                local.get 4
                i32.const 1040
                i32.add
                i32.const 8
                i32.add
                i32.load
                i32.store
                local.get 4
                local.get 4
                i64.load offset=1040
                i64.store offset=1088
              end
              block  ;; label = @6
                local.get 4
                i32.load16_s offset=1096
                local.tee 5
                local.get 6
                i32.le_s
                br_if 0 (;@6;)
                local.get 4
                i32.const 8
                i32.add
                local.get 4
                i32.load offset=1088
                local.get 4
                i32.load offset=1092
                local.get 5
                local.get 3
                local.get 4
                i32.const 1040
                i32.add
                i32.const 4
                call $_ZN4core3num7flt2dec17digits_to_dec_str17h82b39084dda9b16dE
                local.get 4
                i32.load offset=12
                local.set 5
                local.get 4
                i32.load offset=8
                local.set 6
                br 4 (;@2;)
              end
              i32.const 2
              local.set 5
              local.get 4
              i32.const 2
              i32.store16 offset=1040
              block  ;; label = @6
                local.get 3
                i32.eqz
                br_if 0 (;@6;)
                local.get 4
                i32.const 1056
                i32.add
                local.get 3
                i32.store
                local.get 4
                i32.const 0
                i32.store16 offset=1052
                local.get 4
                i32.const 2
                i32.store offset=1048
                local.get 4
                i32.const 1052168
                i32.store offset=1044
                local.get 4
                i32.const 1040
                i32.add
                local.set 6
                br 4 (;@2;)
              end
              i32.const 1
              local.set 5
              local.get 4
              i32.const 1
              i32.store offset=1048
              local.get 4
              i32.const 1052176
              i32.store offset=1044
              local.get 4
              i32.const 1040
              i32.add
              local.set 6
              br 3 (;@2;)
            end
            i32.const 2
            local.set 5
            local.get 4
            i32.const 2
            i32.store16 offset=1040
            block  ;; label = @5
              local.get 3
              i32.eqz
              br_if 0 (;@5;)
              local.get 4
              i32.const 1056
              i32.add
              local.get 3
              i32.store
              local.get 4
              i32.const 0
              i32.store16 offset=1052
              local.get 4
              i32.const 2
              i32.store offset=1048
              local.get 4
              i32.const 1052168
              i32.store offset=1044
              local.get 4
              i32.const 1040
              i32.add
              local.set 6
              br 3 (;@2;)
            end
            i32.const 1
            local.set 5
            local.get 4
            i32.const 1
            i32.store offset=1048
            local.get 4
            i32.const 1052176
            i32.store offset=1044
            local.get 4
            i32.const 1040
            i32.add
            local.set 6
            br 2 (;@2;)
          end
          local.get 4
          i32.const 3
          i32.store offset=1048
          local.get 4
          i32.const 1049132
          i32.store offset=1044
          local.get 4
          i32.const 2
          i32.store16 offset=1040
          local.get 4
          i32.const 1040
          i32.add
          local.set 6
          br 1 (;@2;)
        end
        local.get 4
        i32.const 3
        i32.store offset=1048
        local.get 4
        i32.const 1052177
        i32.store offset=1044
        local.get 4
        i32.const 2
        i32.store16 offset=1040
        local.get 4
        i32.const 1040
        i32.add
        local.set 6
      end
      local.get 4
      i32.const 1100
      i32.add
      local.get 5
      i32.store
      local.get 4
      local.get 6
      i32.store offset=1096
      local.get 4
      local.get 7
      i32.store offset=1092
      local.get 4
      local.get 2
      i32.store offset=1088
      local.get 0
      local.get 4
      i32.const 1088
      i32.add
      call $_ZN4core3fmt9Formatter19pad_formatted_parts17hd9b4ba6efd75be12E
      local.set 6
      local.get 4
      i32.const 1136
      i32.add
      global.set $__stack_pointer
      local.get 6
      return
    end
    i32.const 1052180
    i32.const 37
    i32.const 1052220
    call $_ZN4core9panicking5panic17he2fadc50abd7e245E
    unreachable)
  (func $_ZN4core3fmt9Formatter19pad_formatted_parts17hd9b4ba6efd75be12E (type 2) (param i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 0
              i32.load offset=8
              i32.eqz
              br_if 0 (;@5;)
              local.get 0
              i32.const 12
              i32.add
              i32.load
              local.set 3
              local.get 2
              i32.const 12
              i32.add
              local.get 1
              i32.const 12
              i32.add
              i32.load
              local.tee 4
              i32.store
              local.get 2
              local.get 1
              i32.load offset=8
              local.tee 5
              i32.store offset=8
              local.get 2
              local.get 1
              i32.load offset=4
              local.tee 6
              i32.store offset=4
              local.get 2
              local.get 1
              i32.load
              local.tee 1
              i32.store
              local.get 0
              i32.load8_u offset=32
              local.set 7
              local.get 0
              i32.load offset=28
              local.set 8
              local.get 0
              i32.load8_u offset=24
              i32.const 8
              i32.and
              br_if 1 (;@4;)
              local.get 8
              local.set 9
              local.get 6
              local.set 1
              local.get 7
              local.set 10
              br 2 (;@3;)
            end
            local.get 0
            i32.load
            local.get 0
            i32.const 4
            i32.add
            i32.load
            local.get 1
            call $_ZN4core3fmt9Formatter21write_formatted_parts17h154a6b873f25313dE
            local.set 5
            br 3 (;@1;)
          end
          local.get 0
          i32.load
          local.get 1
          local.get 6
          local.get 0
          i32.load offset=4
          i32.load offset=12
          call_indirect (type 1)
          br_if 1 (;@2;)
          i32.const 1
          local.set 10
          local.get 0
          i32.const 1
          i32.store8 offset=32
          i32.const 48
          local.set 9
          local.get 0
          i32.const 48
          i32.store offset=28
          i32.const 0
          local.set 1
          local.get 2
          i32.const 0
          i32.store offset=4
          local.get 2
          i32.const 1052176
          i32.store
          i32.const 0
          local.get 3
          local.get 6
          i32.sub
          local.tee 6
          local.get 6
          local.get 3
          i32.gt_u
          select
          local.set 3
        end
        block  ;; label = @3
          local.get 4
          i32.eqz
          br_if 0 (;@3;)
          local.get 4
          i32.const 12
          i32.mul
          local.set 4
          loop  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 5
                    i32.load16_u
                    br_table 0 (;@8;) 2 (;@6;) 1 (;@7;) 0 (;@8;)
                  end
                  local.get 5
                  i32.const 4
                  i32.add
                  i32.load
                  local.set 6
                  br 2 (;@5;)
                end
                local.get 5
                i32.const 8
                i32.add
                i32.load
                local.set 6
                br 1 (;@5;)
              end
              block  ;; label = @6
                local.get 5
                i32.const 2
                i32.add
                i32.load16_u
                local.tee 11
                i32.const 1000
                i32.lt_u
                br_if 0 (;@6;)
                i32.const 4
                i32.const 5
                local.get 11
                i32.const 10000
                i32.lt_u
                select
                local.set 6
                br 1 (;@5;)
              end
              i32.const 1
              local.set 6
              local.get 11
              i32.const 10
              i32.lt_u
              br_if 0 (;@5;)
              i32.const 2
              i32.const 3
              local.get 11
              i32.const 100
              i32.lt_u
              select
              local.set 6
            end
            local.get 5
            i32.const 12
            i32.add
            local.set 5
            local.get 6
            local.get 1
            i32.add
            local.set 1
            local.get 4
            i32.const -12
            i32.add
            local.tee 4
            br_if 0 (;@4;)
          end
        end
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 3
              local.get 1
              i32.le_u
              br_if 0 (;@5;)
              local.get 3
              local.get 1
              i32.sub
              local.tee 1
              local.set 4
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 10
                    i32.const 255
                    i32.and
                    local.tee 5
                    br_table 2 (;@6;) 0 (;@8;) 1 (;@7;) 0 (;@8;) 2 (;@6;)
                  end
                  i32.const 0
                  local.set 4
                  local.get 1
                  local.set 5
                  br 1 (;@6;)
                end
                local.get 1
                i32.const 1
                i32.shr_u
                local.set 5
                local.get 1
                i32.const 1
                i32.add
                i32.const 1
                i32.shr_u
                local.set 4
              end
              local.get 5
              i32.const 1
              i32.add
              local.set 5
              local.get 0
              i32.const 4
              i32.add
              i32.load
              local.set 1
              local.get 0
              i32.load
              local.set 6
              loop  ;; label = @6
                local.get 5
                i32.const -1
                i32.add
                local.tee 5
                i32.eqz
                br_if 2 (;@4;)
                local.get 6
                local.get 9
                local.get 1
                i32.load offset=16
                call_indirect (type 2)
                i32.eqz
                br_if 0 (;@6;)
                br 4 (;@2;)
              end
            end
            local.get 0
            i32.load
            local.get 0
            i32.const 4
            i32.add
            i32.load
            local.get 2
            call $_ZN4core3fmt9Formatter21write_formatted_parts17h154a6b873f25313dE
            local.set 5
            br 1 (;@3;)
          end
          local.get 6
          local.get 1
          local.get 2
          call $_ZN4core3fmt9Formatter21write_formatted_parts17h154a6b873f25313dE
          br_if 1 (;@2;)
          i32.const 0
          local.set 5
          loop  ;; label = @4
            block  ;; label = @5
              local.get 4
              local.get 5
              i32.ne
              br_if 0 (;@5;)
              local.get 4
              local.get 4
              i32.lt_u
              local.set 5
              br 2 (;@3;)
            end
            local.get 5
            i32.const 1
            i32.add
            local.set 5
            local.get 6
            local.get 9
            local.get 1
            i32.load offset=16
            call_indirect (type 2)
            i32.eqz
            br_if 0 (;@4;)
          end
          local.get 5
          i32.const -1
          i32.add
          local.get 4
          i32.lt_u
          local.set 5
        end
        local.get 0
        local.get 7
        i32.store8 offset=32
        local.get 0
        local.get 8
        i32.store offset=28
        br 1 (;@1;)
      end
      i32.const 1
      local.set 5
    end
    local.get 2
    i32.const 16
    i32.add
    global.set $__stack_pointer
    local.get 5)
  (func $_ZN4core3fmt5float29float_to_decimal_common_exact17hdedd1adf7adb591bE (type 22) (param i32 f64 i32 i32) (result i32)
    (local i32 i64 i32 i64 i32 i64 i64 i64 i32 i32 i64)
    global.get $__stack_pointer
    i32.const 1136
    i32.sub
    local.tee 4
    global.set $__stack_pointer
    local.get 1
    i64.reinterpret_f64
    local.set 5
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        local.get 1
        f64.eq
        br_if 0 (;@2;)
        i32.const 2
        local.set 6
        br 1 (;@1;)
      end
      local.get 5
      i64.const 4503599627370495
      i64.and
      local.tee 7
      i64.const 4503599627370496
      i64.or
      local.get 5
      i64.const 1
      i64.shl
      i64.const 9007199254740990
      i64.and
      local.get 5
      i64.const 52
      i64.shr_u
      i32.wrap_i64
      i32.const 2047
      i32.and
      local.tee 8
      select
      local.tee 9
      i64.const 1
      i64.and
      local.set 10
      i32.const 3
      local.set 6
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            i32.const 1
            i32.const 2
            i32.const 4
            local.get 5
            i64.const 9218868437227405312
            i64.and
            local.tee 11
            i64.eqz
            local.tee 12
            select
            local.get 11
            i64.const 9218868437227405312
            i64.eq
            select
            i32.const 3
            i32.const 4
            local.get 12
            select
            local.get 7
            i64.eqz
            select
            i32.const -1
            i32.add
            br_table 3 (;@1;) 0 (;@4;) 1 (;@3;) 2 (;@2;) 3 (;@1;)
          end
          i32.const 4
          local.set 6
          br 2 (;@1;)
        end
        local.get 8
        i32.const -1075
        i32.add
        local.set 13
        local.get 10
        i64.eqz
        local.set 6
        i64.const 1
        local.set 14
        br 1 (;@1;)
      end
      i64.const 18014398509481984
      local.get 9
      i64.const 1
      i64.shl
      local.get 9
      i64.const 4503599627370496
      i64.eq
      local.tee 6
      select
      local.set 9
      i64.const 2
      i64.const 1
      local.get 6
      select
      local.set 14
      i32.const -1077
      i32.const -1076
      local.get 6
      select
      local.get 8
      i32.add
      local.set 13
      local.get 10
      i64.eqz
      local.set 6
    end
    local.get 4
    local.get 13
    i32.store16 offset=1128
    local.get 4
    local.get 14
    i64.store offset=1120
    local.get 4
    i64.const 1
    i64.store offset=1112
    local.get 4
    local.get 9
    i64.store offset=1104
    local.get 4
    local.get 6
    i32.store8 offset=1130
    block  ;; label = @1
      block  ;; label = @2
        local.get 6
        i32.const 2
        i32.ne
        br_if 0 (;@2;)
        i32.const 1052176
        local.set 2
        i32.const 0
        local.set 8
        br 1 (;@1;)
      end
      block  ;; label = @2
        local.get 2
        br_if 0 (;@2;)
        i32.const 1052171
        i32.const 1052176
        local.get 5
        i64.const 0
        i64.lt_s
        select
        local.set 2
        local.get 5
        i64.const 63
        i64.shr_u
        i32.wrap_i64
        local.set 8
        br 1 (;@1;)
      end
      i32.const 1052171
      i32.const 1052172
      local.get 5
      i64.const 0
      i64.lt_s
      select
      local.set 2
      i32.const 1
      local.set 8
    end
    i32.const 1
    local.set 12
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 6
                i32.const -2
                i32.add
                i32.const 3
                local.get 6
                i32.const 1
                i32.gt_u
                select
                i32.const 255
                i32.and
                br_table 3 (;@3;) 2 (;@4;) 1 (;@5;) 0 (;@6;) 3 (;@3;)
              end
              i32.const -12
              i32.const 5
              local.get 13
              i32.extend16_s
              local.tee 6
              i32.const 0
              i32.lt_s
              select
              local.get 6
              i32.mul
              local.tee 6
              i32.const 16063
              i32.gt_u
              br_if 4 (;@1;)
              local.get 4
              i32.const 1040
              i32.add
              local.get 4
              i32.const 1104
              i32.add
              local.get 4
              i32.const 16
              i32.add
              local.get 6
              i32.const 4
              i32.shr_u
              i32.const 21
              i32.add
              local.tee 13
              i32.const 0
              local.get 3
              i32.sub
              i32.const -32768
              local.get 3
              i32.const 32768
              i32.lt_u
              select
              local.tee 6
              call $_ZN4core3num7flt2dec8strategy5grisu16format_exact_opt17ha8db8d0ba5fa1a18E
              local.get 6
              i32.extend16_s
              local.set 6
              block  ;; label = @6
                block  ;; label = @7
                  local.get 4
                  i32.load offset=1040
                  br_if 0 (;@7;)
                  local.get 4
                  i32.const 1088
                  i32.add
                  local.get 4
                  i32.const 1104
                  i32.add
                  local.get 4
                  i32.const 16
                  i32.add
                  local.get 13
                  local.get 6
                  call $_ZN4core3num7flt2dec8strategy6dragon12format_exact17h6d7faa24131201fdE
                  br 1 (;@6;)
                end
                local.get 4
                i32.const 1088
                i32.add
                i32.const 8
                i32.add
                local.get 4
                i32.const 1040
                i32.add
                i32.const 8
                i32.add
                i32.load
                i32.store
                local.get 4
                local.get 4
                i64.load offset=1040
                i64.store offset=1088
              end
              block  ;; label = @6
                local.get 4
                i32.load16_s offset=1096
                local.tee 13
                local.get 6
                i32.le_s
                br_if 0 (;@6;)
                local.get 4
                i32.const 8
                i32.add
                local.get 4
                i32.load offset=1088
                local.get 4
                i32.load offset=1092
                local.get 13
                local.get 3
                local.get 4
                i32.const 1040
                i32.add
                i32.const 4
                call $_ZN4core3num7flt2dec17digits_to_dec_str17h82b39084dda9b16dE
                local.get 4
                i32.load offset=12
                local.set 12
                local.get 4
                i32.load offset=8
                local.set 6
                br 4 (;@2;)
              end
              i32.const 2
              local.set 12
              local.get 4
              i32.const 2
              i32.store16 offset=1040
              block  ;; label = @6
                local.get 3
                i32.eqz
                br_if 0 (;@6;)
                local.get 4
                i32.const 1056
                i32.add
                local.get 3
                i32.store
                local.get 4
                i32.const 0
                i32.store16 offset=1052
                local.get 4
                i32.const 2
                i32.store offset=1048
                local.get 4
                i32.const 1052168
                i32.store offset=1044
                local.get 4
                i32.const 1040
                i32.add
                local.set 6
                br 4 (;@2;)
              end
              i32.const 1
              local.set 12
              local.get 4
              i32.const 1
              i32.store offset=1048
              local.get 4
              i32.const 1052176
              i32.store offset=1044
              local.get 4
              i32.const 1040
              i32.add
              local.set 6
              br 3 (;@2;)
            end
            i32.const 2
            local.set 12
            local.get 4
            i32.const 2
            i32.store16 offset=1040
            block  ;; label = @5
              local.get 3
              i32.eqz
              br_if 0 (;@5;)
              local.get 4
              i32.const 1056
              i32.add
              local.get 3
              i32.store
              local.get 4
              i32.const 0
              i32.store16 offset=1052
              local.get 4
              i32.const 2
              i32.store offset=1048
              local.get 4
              i32.const 1052168
              i32.store offset=1044
              local.get 4
              i32.const 1040
              i32.add
              local.set 6
              br 3 (;@2;)
            end
            i32.const 1
            local.set 12
            local.get 4
            i32.const 1
            i32.store offset=1048
            local.get 4
            i32.const 1052176
            i32.store offset=1044
            local.get 4
            i32.const 1040
            i32.add
            local.set 6
            br 2 (;@2;)
          end
          local.get 4
          i32.const 3
          i32.store offset=1048
          local.get 4
          i32.const 1049132
          i32.store offset=1044
          local.get 4
          i32.const 2
          i32.store16 offset=1040
          local.get 4
          i32.const 1040
          i32.add
          local.set 6
          br 1 (;@2;)
        end
        local.get 4
        i32.const 3
        i32.store offset=1048
        local.get 4
        i32.const 1052177
        i32.store offset=1044
        local.get 4
        i32.const 2
        i32.store16 offset=1040
        local.get 4
        i32.const 1040
        i32.add
        local.set 6
      end
      local.get 4
      i32.const 1100
      i32.add
      local.get 12
      i32.store
      local.get 4
      local.get 6
      i32.store offset=1096
      local.get 4
      local.get 8
      i32.store offset=1092
      local.get 4
      local.get 2
      i32.store offset=1088
      local.get 0
      local.get 4
      i32.const 1088
      i32.add
      call $_ZN4core3fmt9Formatter19pad_formatted_parts17hd9b4ba6efd75be12E
      local.set 6
      local.get 4
      i32.const 1136
      i32.add
      global.set $__stack_pointer
      local.get 6
      return
    end
    i32.const 1052180
    i32.const 37
    i32.const 1052220
    call $_ZN4core9panicking5panic17he2fadc50abd7e245E
    unreachable)
  (func $_ZN4core3fmt5float32float_to_decimal_common_shortest17h0203f2e3c1aa9f40E (type 22) (param i32 f64 i32 i32) (result i32)
    (local i32 i64 i32 i64 i32 i64 i64 i64 i32 i32 i64)
    global.get $__stack_pointer
    i32.const 128
    i32.sub
    local.tee 4
    global.set $__stack_pointer
    local.get 1
    i64.reinterpret_f64
    local.set 5
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        local.get 1
        f64.eq
        br_if 0 (;@2;)
        i32.const 2
        local.set 6
        br 1 (;@1;)
      end
      local.get 5
      i64.const 4503599627370495
      i64.and
      local.tee 7
      i64.const 4503599627370496
      i64.or
      local.get 5
      i64.const 1
      i64.shl
      i64.const 9007199254740990
      i64.and
      local.get 5
      i64.const 52
      i64.shr_u
      i32.wrap_i64
      i32.const 2047
      i32.and
      local.tee 8
      select
      local.tee 9
      i64.const 1
      i64.and
      local.set 10
      i32.const 3
      local.set 6
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            i32.const 1
            i32.const 2
            i32.const 4
            local.get 5
            i64.const 9218868437227405312
            i64.and
            local.tee 11
            i64.eqz
            local.tee 12
            select
            local.get 11
            i64.const 9218868437227405312
            i64.eq
            select
            i32.const 3
            i32.const 4
            local.get 12
            select
            local.get 7
            i64.eqz
            select
            i32.const -1
            i32.add
            br_table 3 (;@1;) 0 (;@4;) 1 (;@3;) 2 (;@2;) 3 (;@1;)
          end
          i32.const 4
          local.set 6
          br 2 (;@1;)
        end
        local.get 8
        i32.const -1075
        i32.add
        local.set 13
        local.get 10
        i64.eqz
        local.set 6
        i64.const 1
        local.set 14
        br 1 (;@1;)
      end
      i64.const 18014398509481984
      local.get 9
      i64.const 1
      i64.shl
      local.get 9
      i64.const 4503599627370496
      i64.eq
      local.tee 6
      select
      local.set 9
      i64.const 2
      i64.const 1
      local.get 6
      select
      local.set 14
      i32.const -1077
      i32.const -1076
      local.get 6
      select
      local.get 8
      i32.add
      local.set 13
      local.get 10
      i64.eqz
      local.set 6
    end
    local.get 4
    local.get 13
    i32.store16 offset=120
    local.get 4
    local.get 14
    i64.store offset=112
    local.get 4
    i64.const 1
    i64.store offset=104
    local.get 4
    local.get 9
    i64.store offset=96
    local.get 4
    local.get 6
    i32.store8 offset=122
    block  ;; label = @1
      block  ;; label = @2
        local.get 6
        i32.const 2
        i32.ne
        br_if 0 (;@2;)
        i32.const 1052176
        local.set 12
        i32.const 0
        local.set 2
        br 1 (;@1;)
      end
      block  ;; label = @2
        local.get 2
        br_if 0 (;@2;)
        i32.const 1052171
        i32.const 1052176
        local.get 5
        i64.const 0
        i64.lt_s
        select
        local.set 12
        local.get 5
        i64.const 63
        i64.shr_u
        i32.wrap_i64
        local.set 2
        br 1 (;@1;)
      end
      i32.const 1052171
      i32.const 1052172
      local.get 5
      i64.const 0
      i64.lt_s
      select
      local.set 12
      i32.const 1
      local.set 2
    end
    i32.const 1
    local.set 13
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 6
              i32.const -2
              i32.add
              i32.const 3
              local.get 6
              i32.const 1
              i32.gt_u
              select
              i32.const 255
              i32.and
              br_table 3 (;@2;) 2 (;@3;) 1 (;@4;) 0 (;@5;) 3 (;@2;)
            end
            local.get 4
            i32.const 32
            i32.add
            local.get 4
            i32.const 96
            i32.add
            local.get 4
            i32.const 15
            i32.add
            i32.const 17
            call $_ZN4core3num7flt2dec8strategy5grisu19format_shortest_opt17ha5578b8fd1a67f89E
            block  ;; label = @5
              block  ;; label = @6
                local.get 4
                i32.load offset=32
                br_if 0 (;@6;)
                local.get 4
                i32.const 80
                i32.add
                local.get 4
                i32.const 96
                i32.add
                local.get 4
                i32.const 15
                i32.add
                i32.const 17
                call $_ZN4core3num7flt2dec8strategy6dragon15format_shortest17hb3d88f8cf15eff79E
                br 1 (;@5;)
              end
              local.get 4
              i32.const 80
              i32.add
              i32.const 8
              i32.add
              local.get 4
              i32.const 32
              i32.add
              i32.const 8
              i32.add
              i32.load
              i32.store
              local.get 4
              local.get 4
              i64.load offset=32
              i64.store offset=80
            end
            local.get 4
            local.get 4
            i32.load offset=80
            local.get 4
            i32.load offset=84
            local.get 4
            i32.load16_u offset=88
            local.get 3
            local.get 4
            i32.const 32
            i32.add
            i32.const 4
            call $_ZN4core3num7flt2dec17digits_to_dec_str17h82b39084dda9b16dE
            local.get 4
            i32.load offset=4
            local.set 13
            local.get 4
            i32.load
            local.set 6
            br 3 (;@1;)
          end
          i32.const 2
          local.set 13
          local.get 4
          i32.const 2
          i32.store16 offset=32
          block  ;; label = @4
            local.get 3
            i32.eqz
            br_if 0 (;@4;)
            local.get 4
            i32.const 48
            i32.add
            i32.const 1
            i32.store
            local.get 4
            i32.const 0
            i32.store16 offset=44
            local.get 4
            i32.const 2
            i32.store offset=40
            local.get 4
            i32.const 1052168
            i32.store offset=36
            local.get 4
            i32.const 32
            i32.add
            local.set 6
            br 3 (;@1;)
          end
          i32.const 1
          local.set 13
          local.get 4
          i32.const 1
          i32.store offset=40
          local.get 4
          i32.const 1052176
          i32.store offset=36
          local.get 4
          i32.const 32
          i32.add
          local.set 6
          br 2 (;@1;)
        end
        local.get 4
        i32.const 3
        i32.store offset=40
        local.get 4
        i32.const 1049132
        i32.store offset=36
        local.get 4
        i32.const 2
        i32.store16 offset=32
        local.get 4
        i32.const 32
        i32.add
        local.set 6
        br 1 (;@1;)
      end
      local.get 4
      i32.const 3
      i32.store offset=40
      local.get 4
      i32.const 1052177
      i32.store offset=36
      local.get 4
      i32.const 2
      i32.store16 offset=32
      local.get 4
      i32.const 32
      i32.add
      local.set 6
    end
    local.get 4
    i32.const 92
    i32.add
    local.get 13
    i32.store
    local.get 4
    local.get 6
    i32.store offset=88
    local.get 4
    local.get 2
    i32.store offset=84
    local.get 4
    local.get 12
    i32.store offset=80
    local.get 0
    local.get 4
    i32.const 80
    i32.add
    call $_ZN4core3fmt9Formatter19pad_formatted_parts17hd9b4ba6efd75be12E
    local.set 6
    local.get 4
    i32.const 128
    i32.add
    global.set $__stack_pointer
    local.get 6)
  (func $_ZN4core3fmt5float32float_to_decimal_common_shortest17hab759cd6e3d171f2E (type 21) (param i32 f32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i64 i64 i32 i64)
    global.get $__stack_pointer
    i32.const 128
    i32.sub
    local.tee 4
    global.set $__stack_pointer
    local.get 1
    i32.reinterpret_f32
    local.set 5
    i32.const 3
    local.set 6
    block  ;; label = @1
      block  ;; label = @2
        local.get 1
        f32.const inf (;=inf;)
        f32.ne
        br_if 0 (;@2;)
        br 1 (;@1;)
      end
      local.get 1
      f32.const -inf (;=-inf;)
      f32.eq
      br_if 0 (;@1;)
      block  ;; label = @2
        local.get 1
        local.get 1
        f32.eq
        br_if 0 (;@2;)
        i32.const 2
        local.set 6
        br 1 (;@1;)
      end
      local.get 5
      i32.const 8388607
      i32.and
      i32.const 8388608
      i32.or
      local.get 5
      i32.const 1
      i32.shl
      i32.const 16777214
      i32.and
      local.get 5
      i32.const 23
      i32.shr_u
      i32.const 255
      i32.and
      local.tee 7
      select
      local.tee 8
      i64.extend_i32_u
      local.tee 9
      i64.const 1
      i64.and
      local.set 10
      i32.const 4
      local.set 6
      block  ;; label = @2
        block  ;; label = @3
          i32.const 4
          i32.const 3
          local.get 5
          i32.const 2139095040
          i32.and
          select
          i32.const 2
          local.get 5
          i32.const 2147483647
          i32.and
          select
          i32.const -2
          i32.add
          br_table 2 (;@1;) 0 (;@3;) 1 (;@2;) 2 (;@1;)
        end
        local.get 7
        i32.const -150
        i32.add
        local.set 11
        local.get 10
        i64.eqz
        local.set 6
        i64.const 1
        local.set 12
        br 1 (;@1;)
      end
      i64.const 33554432
      local.get 9
      i64.const 1
      i64.shl
      local.get 8
      i32.const 8388608
      i32.eq
      local.tee 6
      select
      local.set 9
      i64.const 2
      i64.const 1
      local.get 6
      select
      local.set 12
      i32.const -152
      i32.const -151
      local.get 6
      select
      local.get 7
      i32.add
      local.set 11
      local.get 10
      i64.eqz
      local.set 6
    end
    local.get 4
    local.get 11
    i32.store16 offset=120
    local.get 4
    local.get 12
    i64.store offset=112
    local.get 4
    i64.const 1
    i64.store offset=104
    local.get 4
    local.get 9
    i64.store offset=96
    local.get 4
    local.get 6
    i32.store8 offset=122
    block  ;; label = @1
      block  ;; label = @2
        local.get 6
        i32.const 2
        i32.ne
        br_if 0 (;@2;)
        i32.const 1052176
        local.set 11
        i32.const 0
        local.set 2
        br 1 (;@1;)
      end
      block  ;; label = @2
        local.get 2
        br_if 0 (;@2;)
        i32.const 1052171
        i32.const 1052176
        local.get 5
        i32.const 0
        i32.lt_s
        select
        local.set 11
        local.get 5
        i32.const 31
        i32.shr_u
        local.set 2
        br 1 (;@1;)
      end
      i32.const 1052171
      i32.const 1052172
      local.get 5
      i32.const 0
      i32.lt_s
      select
      local.set 11
      i32.const 1
      local.set 2
    end
    i32.const 1
    local.set 5
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 6
              i32.const -2
              i32.add
              i32.const 3
              local.get 6
              i32.const 1
              i32.gt_u
              select
              i32.const 255
              i32.and
              br_table 3 (;@2;) 2 (;@3;) 1 (;@4;) 0 (;@5;) 3 (;@2;)
            end
            local.get 4
            i32.const 32
            i32.add
            local.get 4
            i32.const 96
            i32.add
            local.get 4
            i32.const 15
            i32.add
            i32.const 17
            call $_ZN4core3num7flt2dec8strategy5grisu19format_shortest_opt17ha5578b8fd1a67f89E
            block  ;; label = @5
              block  ;; label = @6
                local.get 4
                i32.load offset=32
                br_if 0 (;@6;)
                local.get 4
                i32.const 80
                i32.add
                local.get 4
                i32.const 96
                i32.add
                local.get 4
                i32.const 15
                i32.add
                i32.const 17
                call $_ZN4core3num7flt2dec8strategy6dragon15format_shortest17hb3d88f8cf15eff79E
                br 1 (;@5;)
              end
              local.get 4
              i32.const 80
              i32.add
              i32.const 8
              i32.add
              local.get 4
              i32.const 32
              i32.add
              i32.const 8
              i32.add
              i32.load
              i32.store
              local.get 4
              local.get 4
              i64.load offset=32
              i64.store offset=80
            end
            local.get 4
            local.get 4
            i32.load offset=80
            local.get 4
            i32.load offset=84
            local.get 4
            i32.load16_u offset=88
            local.get 3
            local.get 4
            i32.const 32
            i32.add
            i32.const 4
            call $_ZN4core3num7flt2dec17digits_to_dec_str17h82b39084dda9b16dE
            local.get 4
            i32.load offset=4
            local.set 5
            local.get 4
            i32.load
            local.set 6
            br 3 (;@1;)
          end
          i32.const 2
          local.set 5
          local.get 4
          i32.const 2
          i32.store16 offset=32
          block  ;; label = @4
            local.get 3
            i32.eqz
            br_if 0 (;@4;)
            local.get 4
            i32.const 48
            i32.add
            i32.const 1
            i32.store
            local.get 4
            i32.const 0
            i32.store16 offset=44
            local.get 4
            i32.const 2
            i32.store offset=40
            local.get 4
            i32.const 1052168
            i32.store offset=36
            local.get 4
            i32.const 32
            i32.add
            local.set 6
            br 3 (;@1;)
          end
          i32.const 1
          local.set 5
          local.get 4
          i32.const 1
          i32.store offset=40
          local.get 4
          i32.const 1052176
          i32.store offset=36
          local.get 4
          i32.const 32
          i32.add
          local.set 6
          br 2 (;@1;)
        end
        local.get 4
        i32.const 3
        i32.store offset=40
        local.get 4
        i32.const 1049132
        i32.store offset=36
        local.get 4
        i32.const 2
        i32.store16 offset=32
        local.get 4
        i32.const 32
        i32.add
        local.set 6
        br 1 (;@1;)
      end
      local.get 4
      i32.const 3
      i32.store offset=40
      local.get 4
      i32.const 1052177
      i32.store offset=36
      local.get 4
      i32.const 2
      i32.store16 offset=32
      local.get 4
      i32.const 32
      i32.add
      local.set 6
    end
    local.get 4
    i32.const 92
    i32.add
    local.get 5
    i32.store
    local.get 4
    local.get 6
    i32.store offset=88
    local.get 4
    local.get 2
    i32.store offset=84
    local.get 4
    local.get 11
    i32.store offset=80
    local.get 0
    local.get 4
    i32.const 80
    i32.add
    call $_ZN4core3fmt9Formatter19pad_formatted_parts17hd9b4ba6efd75be12E
    local.set 6
    local.get 4
    i32.const 128
    i32.add
    global.set $__stack_pointer
    local.get 6)
  (func $_ZN4core3str5count14do_count_chars17hffbd23dea83a0c54E (type 2) (param i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i32.const 3
        i32.add
        i32.const -4
        i32.and
        local.tee 2
        local.get 0
        i32.sub
        local.tee 3
        local.get 1
        i32.gt_u
        br_if 0 (;@2;)
        local.get 3
        i32.const 4
        i32.gt_u
        br_if 0 (;@2;)
        local.get 1
        local.get 3
        i32.sub
        local.tee 4
        i32.const 4
        i32.lt_u
        br_if 0 (;@2;)
        local.get 4
        i32.const 3
        i32.and
        local.set 5
        i32.const 0
        local.set 6
        i32.const 0
        local.set 1
        block  ;; label = @3
          local.get 2
          local.get 0
          i32.eq
          br_if 0 (;@3;)
          local.get 3
          i32.const 3
          i32.and
          local.set 7
          block  ;; label = @4
            block  ;; label = @5
              local.get 2
              local.get 0
              i32.const -1
              i32.xor
              i32.add
              i32.const 3
              i32.ge_u
              br_if 0 (;@5;)
              i32.const 0
              local.set 1
              local.get 0
              local.set 2
              br 1 (;@4;)
            end
            local.get 3
            i32.const -4
            i32.and
            local.set 8
            i32.const 0
            local.set 1
            local.get 0
            local.set 2
            loop  ;; label = @5
              local.get 1
              local.get 2
              i32.load8_s
              i32.const -65
              i32.gt_s
              i32.add
              local.get 2
              i32.load8_s offset=1
              i32.const -65
              i32.gt_s
              i32.add
              local.get 2
              i32.load8_s offset=2
              i32.const -65
              i32.gt_s
              i32.add
              local.get 2
              i32.load8_s offset=3
              i32.const -65
              i32.gt_s
              i32.add
              local.set 1
              local.get 2
              i32.const 4
              i32.add
              local.set 2
              local.get 8
              i32.const -4
              i32.add
              local.tee 8
              br_if 0 (;@5;)
            end
          end
          local.get 7
          i32.eqz
          br_if 0 (;@3;)
          loop  ;; label = @4
            local.get 1
            local.get 2
            i32.load8_s
            i32.const -65
            i32.gt_s
            i32.add
            local.set 1
            local.get 2
            i32.const 1
            i32.add
            local.set 2
            local.get 7
            i32.const -1
            i32.add
            local.tee 7
            br_if 0 (;@4;)
          end
        end
        local.get 0
        local.get 3
        i32.add
        local.set 0
        block  ;; label = @3
          local.get 5
          i32.eqz
          br_if 0 (;@3;)
          local.get 0
          local.get 4
          i32.const -4
          i32.and
          i32.add
          local.tee 2
          i32.load8_s
          i32.const -65
          i32.gt_s
          local.set 6
          local.get 5
          i32.const 1
          i32.eq
          br_if 0 (;@3;)
          local.get 6
          local.get 2
          i32.load8_s offset=1
          i32.const -65
          i32.gt_s
          i32.add
          local.set 6
          local.get 5
          i32.const 2
          i32.eq
          br_if 0 (;@3;)
          local.get 6
          local.get 2
          i32.load8_s offset=2
          i32.const -65
          i32.gt_s
          i32.add
          local.set 6
        end
        local.get 4
        i32.const 2
        i32.shr_u
        local.set 3
        local.get 6
        local.get 1
        i32.add
        local.set 7
        loop  ;; label = @3
          local.get 0
          local.set 6
          local.get 3
          i32.eqz
          br_if 2 (;@1;)
          local.get 3
          i32.const 192
          local.get 3
          i32.const 192
          i32.lt_u
          select
          local.tee 4
          i32.const 3
          i32.and
          local.set 5
          local.get 4
          i32.const 2
          i32.shl
          local.set 9
          block  ;; label = @4
            block  ;; label = @5
              local.get 4
              i32.const 252
              i32.and
              local.tee 10
              br_if 0 (;@5;)
              i32.const 0
              local.set 2
              br 1 (;@4;)
            end
            local.get 6
            local.get 10
            i32.const 2
            i32.shl
            i32.add
            local.set 8
            i32.const 0
            local.set 2
            local.get 6
            local.set 0
            loop  ;; label = @5
              local.get 0
              i32.eqz
              br_if 1 (;@4;)
              local.get 0
              i32.const 12
              i32.add
              i32.load
              local.tee 1
              i32.const -1
              i32.xor
              i32.const 7
              i32.shr_u
              local.get 1
              i32.const 6
              i32.shr_u
              i32.or
              i32.const 16843009
              i32.and
              local.get 0
              i32.const 8
              i32.add
              i32.load
              local.tee 1
              i32.const -1
              i32.xor
              i32.const 7
              i32.shr_u
              local.get 1
              i32.const 6
              i32.shr_u
              i32.or
              i32.const 16843009
              i32.and
              local.get 0
              i32.const 4
              i32.add
              i32.load
              local.tee 1
              i32.const -1
              i32.xor
              i32.const 7
              i32.shr_u
              local.get 1
              i32.const 6
              i32.shr_u
              i32.or
              i32.const 16843009
              i32.and
              local.get 0
              i32.load
              local.tee 1
              i32.const -1
              i32.xor
              i32.const 7
              i32.shr_u
              local.get 1
              i32.const 6
              i32.shr_u
              i32.or
              i32.const 16843009
              i32.and
              local.get 2
              i32.add
              i32.add
              i32.add
              i32.add
              local.set 2
              local.get 0
              i32.const 16
              i32.add
              local.tee 0
              local.get 8
              i32.ne
              br_if 0 (;@5;)
            end
          end
          local.get 3
          local.get 4
          i32.sub
          local.set 3
          local.get 6
          local.get 9
          i32.add
          local.set 0
          local.get 2
          i32.const 8
          i32.shr_u
          i32.const 16711935
          i32.and
          local.get 2
          i32.const 16711935
          i32.and
          i32.add
          i32.const 65537
          i32.mul
          i32.const 16
          i32.shr_u
          local.get 7
          i32.add
          local.set 7
          local.get 5
          i32.eqz
          br_if 0 (;@3;)
        end
        block  ;; label = @3
          block  ;; label = @4
            local.get 6
            br_if 0 (;@4;)
            i32.const 0
            local.set 0
            br 1 (;@3;)
          end
          local.get 6
          local.get 10
          i32.const 2
          i32.shl
          i32.add
          local.tee 2
          i32.load
          local.tee 0
          i32.const -1
          i32.xor
          i32.const 7
          i32.shr_u
          local.get 0
          i32.const 6
          i32.shr_u
          i32.or
          i32.const 16843009
          i32.and
          local.set 0
          local.get 5
          i32.const 1
          i32.eq
          br_if 0 (;@3;)
          local.get 2
          i32.load offset=4
          local.tee 1
          i32.const -1
          i32.xor
          i32.const 7
          i32.shr_u
          local.get 1
          i32.const 6
          i32.shr_u
          i32.or
          i32.const 16843009
          i32.and
          local.get 0
          i32.add
          local.set 0
          local.get 5
          i32.const 2
          i32.eq
          br_if 0 (;@3;)
          local.get 2
          i32.load offset=8
          local.tee 2
          i32.const -1
          i32.xor
          i32.const 7
          i32.shr_u
          local.get 2
          i32.const 6
          i32.shr_u
          i32.or
          i32.const 16843009
          i32.and
          local.get 0
          i32.add
          local.set 0
        end
        local.get 0
        i32.const 8
        i32.shr_u
        i32.const 459007
        i32.and
        local.get 0
        i32.const 16711935
        i32.and
        i32.add
        i32.const 65537
        i32.mul
        i32.const 16
        i32.shr_u
        local.get 7
        i32.add
        return
      end
      block  ;; label = @2
        local.get 1
        br_if 0 (;@2;)
        i32.const 0
        return
      end
      local.get 1
      i32.const 3
      i32.and
      local.set 2
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.const 4
          i32.ge_u
          br_if 0 (;@3;)
          i32.const 0
          local.set 7
          br 1 (;@2;)
        end
        local.get 1
        i32.const -4
        i32.and
        local.set 1
        i32.const 0
        local.set 7
        loop  ;; label = @3
          local.get 7
          local.get 0
          i32.load8_s
          i32.const -65
          i32.gt_s
          i32.add
          local.get 0
          i32.load8_s offset=1
          i32.const -65
          i32.gt_s
          i32.add
          local.get 0
          i32.load8_s offset=2
          i32.const -65
          i32.gt_s
          i32.add
          local.get 0
          i32.load8_s offset=3
          i32.const -65
          i32.gt_s
          i32.add
          local.set 7
          local.get 0
          i32.const 4
          i32.add
          local.set 0
          local.get 1
          i32.const -4
          i32.add
          local.tee 1
          br_if 0 (;@3;)
        end
      end
      local.get 2
      i32.eqz
      br_if 0 (;@1;)
      loop  ;; label = @2
        local.get 7
        local.get 0
        i32.load8_s
        i32.const -65
        i32.gt_s
        i32.add
        local.set 7
        local.get 0
        i32.const 1
        i32.add
        local.set 0
        local.get 2
        i32.const -1
        i32.add
        local.tee 2
        br_if 0 (;@2;)
      end
    end
    local.get 7)
  (func $_ZN4core3fmt9Formatter12pad_integral12write_prefix17hbf853012c4e1643dE (type 16) (param i32 i32 i32 i32 i32) (result i32)
    (local i32)
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 2
          i32.const 1114112
          i32.eq
          br_if 0 (;@3;)
          i32.const 1
          local.set 5
          local.get 0
          local.get 2
          local.get 1
          i32.load offset=16
          call_indirect (type 2)
          br_if 1 (;@2;)
        end
        local.get 3
        br_if 1 (;@1;)
        i32.const 0
        local.set 5
      end
      local.get 5
      return
    end
    local.get 0
    local.get 3
    local.get 4
    local.get 1
    i32.load offset=12
    call_indirect (type 1))
  (func $_ZN4core3fmt9Formatter21write_formatted_parts17h154a6b873f25313dE (type 1) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 16
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.load offset=4
        local.tee 4
        i32.eqz
        br_if 0 (;@2;)
        i32.const 1
        local.set 5
        local.get 0
        local.get 2
        i32.load
        local.get 4
        local.get 1
        i32.load offset=12
        call_indirect (type 1)
        br_if 1 (;@1;)
      end
      block  ;; label = @2
        local.get 2
        i32.const 12
        i32.add
        i32.load
        local.tee 5
        br_if 0 (;@2;)
        i32.const 0
        local.set 5
        br 1 (;@1;)
      end
      local.get 2
      i32.load offset=8
      local.tee 6
      local.get 5
      i32.const 12
      i32.mul
      i32.add
      local.set 7
      local.get 3
      i32.const 8
      i32.add
      i32.const 4
      i32.add
      local.set 8
      loop  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 6
                i32.load16_u
                br_table 0 (;@6;) 2 (;@4;) 1 (;@5;) 0 (;@6;)
              end
              block  ;; label = @6
                block  ;; label = @7
                  local.get 6
                  i32.load offset=4
                  local.tee 2
                  i32.const 65
                  i32.lt_u
                  br_if 0 (;@7;)
                  local.get 1
                  i32.const 12
                  i32.add
                  i32.load
                  local.set 5
                  loop  ;; label = @8
                    block  ;; label = @9
                      local.get 0
                      i32.const 1052800
                      i32.const 64
                      local.get 5
                      call_indirect (type 1)
                      i32.eqz
                      br_if 0 (;@9;)
                      i32.const 1
                      local.set 5
                      br 8 (;@1;)
                    end
                    local.get 2
                    i32.const -64
                    i32.add
                    local.tee 2
                    i32.const 64
                    i32.gt_u
                    br_if 0 (;@8;)
                    br 2 (;@6;)
                  end
                end
                local.get 2
                i32.eqz
                br_if 3 (;@3;)
              end
              block  ;; label = @6
                block  ;; label = @7
                  local.get 2
                  i32.const 63
                  i32.gt_u
                  br_if 0 (;@7;)
                  local.get 2
                  i32.const 1052800
                  i32.add
                  i32.load8_s
                  i32.const -65
                  i32.le_s
                  br_if 1 (;@6;)
                end
                local.get 0
                i32.const 1052800
                local.get 2
                local.get 1
                i32.const 12
                i32.add
                i32.load
                call_indirect (type 1)
                i32.eqz
                br_if 3 (;@3;)
                i32.const 1
                local.set 5
                br 5 (;@1;)
              end
              i32.const 1052800
              i32.const 64
              i32.const 0
              local.get 2
              i32.const 1052864
              call $_ZN4core3str16slice_error_fail17h48e751b16631e034E
              unreachable
            end
            local.get 0
            local.get 6
            i32.load offset=4
            local.get 6
            i32.const 8
            i32.add
            i32.load
            local.get 1
            i32.const 12
            i32.add
            i32.load
            call_indirect (type 1)
            i32.eqz
            br_if 1 (;@3;)
            i32.const 1
            local.set 5
            br 3 (;@1;)
          end
          local.get 6
          i32.load16_u offset=2
          local.set 2
          local.get 8
          i32.const 0
          i32.store8
          local.get 3
          i32.const 0
          i32.store offset=8
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 6
                      i32.load16_u
                      br_table 2 (;@7;) 1 (;@8;) 0 (;@9;) 2 (;@7;)
                    end
                    local.get 6
                    i32.const 8
                    i32.add
                    local.set 5
                    br 2 (;@6;)
                  end
                  block  ;; label = @8
                    local.get 6
                    i32.load16_u offset=2
                    local.tee 5
                    i32.const 1000
                    i32.lt_u
                    br_if 0 (;@8;)
                    i32.const 4
                    i32.const 5
                    local.get 5
                    i32.const 10000
                    i32.lt_u
                    select
                    local.set 9
                    br 3 (;@5;)
                  end
                  i32.const 1
                  local.set 9
                  local.get 5
                  i32.const 10
                  i32.lt_u
                  br_if 2 (;@5;)
                  i32.const 2
                  i32.const 3
                  local.get 5
                  i32.const 100
                  i32.lt_u
                  select
                  local.set 9
                  br 2 (;@5;)
                end
                local.get 6
                i32.const 4
                i32.add
                local.set 5
              end
              block  ;; label = @6
                local.get 5
                i32.load
                local.tee 9
                i32.const 6
                i32.ge_u
                br_if 0 (;@6;)
                local.get 9
                br_if 1 (;@5;)
                i32.const 0
                local.set 9
                br 2 (;@4;)
              end
              local.get 9
              i32.const 5
              i32.const 1052784
              call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
              unreachable
            end
            local.get 3
            i32.const 8
            i32.add
            local.get 9
            i32.add
            local.set 4
            block  ;; label = @5
              block  ;; label = @6
                local.get 9
                i32.const 1
                i32.and
                br_if 0 (;@6;)
                local.get 2
                local.set 5
                br 1 (;@5;)
              end
              local.get 4
              i32.const -1
              i32.add
              local.tee 4
              local.get 2
              local.get 2
              i32.const 65535
              i32.and
              i32.const 10
              i32.div_u
              local.tee 5
              i32.const 10
              i32.mul
              i32.sub
              i32.const 48
              i32.or
              i32.store8
            end
            local.get 9
            i32.const 1
            i32.eq
            br_if 0 (;@4;)
            local.get 4
            i32.const -2
            i32.add
            local.set 2
            loop  ;; label = @5
              local.get 2
              local.get 5
              i32.const 65535
              i32.and
              local.tee 4
              i32.const 10
              i32.div_u
              local.tee 10
              i32.const 10
              i32.rem_u
              i32.const 48
              i32.or
              i32.store8
              local.get 2
              i32.const 1
              i32.add
              local.get 5
              local.get 10
              i32.const 10
              i32.mul
              i32.sub
              i32.const 48
              i32.or
              i32.store8
              local.get 4
              i32.const 100
              i32.div_u
              local.set 5
              local.get 2
              local.get 3
              i32.const 8
              i32.add
              i32.eq
              local.set 4
              local.get 2
              i32.const -2
              i32.add
              local.set 2
              local.get 4
              i32.eqz
              br_if 0 (;@5;)
            end
          end
          local.get 0
          local.get 3
          i32.const 8
          i32.add
          local.get 9
          local.get 1
          i32.const 12
          i32.add
          i32.load
          call_indirect (type 1)
          i32.eqz
          br_if 0 (;@3;)
          i32.const 1
          local.set 5
          br 2 (;@1;)
        end
        local.get 6
        i32.const 12
        i32.add
        local.tee 6
        local.get 7
        i32.ne
        br_if 0 (;@2;)
      end
      i32.const 0
      local.set 5
    end
    local.get 3
    i32.const 16
    i32.add
    global.set $__stack_pointer
    local.get 5)
  (func $_ZN4core3str16slice_error_fail17h48e751b16631e034E (type 15) (param i32 i32 i32 i32 i32)
    local.get 0
    local.get 1
    local.get 2
    local.get 3
    local.get 4
    call $_ZN4core3str19slice_error_fail_rt17h8712bc75e2bb7542E
    unreachable)
  (func $_ZN4core7unicode12unicode_data15grapheme_extend6lookup17h24f97697985c6c02E (type 13) (param i32) (result i32)
    (local i32 i32 i32 i32 i32)
    local.get 0
    i32.const 11
    i32.shl
    local.set 1
    i32.const 0
    local.set 2
    i32.const 33
    local.set 3
    i32.const 33
    local.set 4
    block  ;; label = @1
      block  ;; label = @2
        loop  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              i32.const -1
              local.get 3
              i32.const 1
              i32.shr_u
              local.get 2
              i32.add
              local.tee 5
              i32.const 2
              i32.shl
              i32.const 1054984
              i32.add
              i32.load
              i32.const 11
              i32.shl
              local.tee 3
              local.get 1
              i32.ne
              local.get 3
              local.get 1
              i32.lt_u
              select
              local.tee 3
              i32.const 1
              i32.ne
              br_if 0 (;@5;)
              local.get 5
              local.set 4
              br 1 (;@4;)
            end
            local.get 3
            i32.const 255
            i32.and
            i32.const 255
            i32.ne
            br_if 2 (;@2;)
            local.get 5
            i32.const 1
            i32.add
            local.set 2
          end
          local.get 4
          local.get 2
          i32.sub
          local.set 3
          local.get 4
          local.get 2
          i32.gt_u
          br_if 0 (;@3;)
          br 2 (;@1;)
        end
      end
      local.get 5
      i32.const 1
      i32.add
      local.set 2
    end
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.const 32
        i32.gt_u
        br_if 0 (;@2;)
        local.get 2
        i32.const 2
        i32.shl
        local.tee 1
        i32.const 1054984
        i32.add
        i32.load
        i32.const 21
        i32.shr_u
        local.set 4
        i32.const 727
        local.set 5
        block  ;; label = @3
          block  ;; label = @4
            local.get 2
            i32.const 32
            i32.eq
            br_if 0 (;@4;)
            local.get 1
            i32.const 1054988
            i32.add
            i32.load
            i32.const 21
            i32.shr_u
            local.set 5
            local.get 2
            br_if 0 (;@4;)
            i32.const 0
            local.set 1
            br 1 (;@3;)
          end
          local.get 1
          i32.const 1054980
          i32.add
          i32.load
          i32.const 2097151
          i32.and
          local.set 3
          i32.const 1
          local.set 1
        end
        local.get 5
        local.get 4
        i32.const -1
        i32.xor
        i32.add
        i32.eqz
        br_if 1 (;@1;)
        i32.const 0
        local.set 2
        local.get 0
        local.get 3
        i32.const 0
        local.get 1
        select
        i32.sub
        local.set 3
        local.get 4
        i32.const 727
        local.get 4
        i32.const 727
        i32.gt_u
        select
        local.set 1
        local.get 5
        i32.const -1
        i32.add
        local.set 5
        loop  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 1
              local.get 4
              i32.eq
              br_if 0 (;@5;)
              local.get 2
              local.get 4
              i32.const 1055116
              i32.add
              i32.load8_u
              i32.add
              local.tee 2
              local.get 3
              i32.le_u
              br_if 1 (;@4;)
              br 4 (;@1;)
            end
            local.get 1
            i32.const 727
            i32.const 1054968
            call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
            unreachable
          end
          local.get 5
          local.get 4
          i32.const 1
          i32.add
          local.tee 4
          i32.ne
          br_if 0 (;@3;)
        end
        local.get 5
        local.set 4
        br 1 (;@1;)
      end
      local.get 2
      i32.const 33
      i32.const 1054952
      call $_ZN4core9panicking18panic_bounds_check17h6f239e91c234b035E
      unreachable
    end
    local.get 4
    i32.const 1
    i32.and)
  (func $_ZN4core7unicode9printable12is_printable17h237b4d5f567ccf43E (type 13) (param i32) (result i32)
    (local i32)
    block  ;; label = @1
      local.get 0
      i32.const 32
      i32.ge_u
      br_if 0 (;@1;)
      i32.const 0
      return
    end
    i32.const 1
    local.set 1
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i32.const 127
        i32.lt_u
        br_if 0 (;@2;)
        local.get 0
        i32.const 65536
        i32.lt_u
        br_if 1 (;@1;)
        block  ;; label = @3
          block  ;; label = @4
            local.get 0
            i32.const 131072
            i32.lt_u
            br_if 0 (;@4;)
            block  ;; label = @5
              local.get 0
              i32.const -205744
              i32.add
              i32.const 712016
              i32.ge_u
              br_if 0 (;@5;)
              i32.const 0
              return
            end
            block  ;; label = @5
              local.get 0
              i32.const -201547
              i32.add
              i32.const 5
              i32.ge_u
              br_if 0 (;@5;)
              i32.const 0
              return
            end
            block  ;; label = @5
              local.get 0
              i32.const -195102
              i32.add
              i32.const 1506
              i32.ge_u
              br_if 0 (;@5;)
              i32.const 0
              return
            end
            block  ;; label = @5
              local.get 0
              i32.const -191457
              i32.add
              i32.const 3103
              i32.ge_u
              br_if 0 (;@5;)
              i32.const 0
              return
            end
            block  ;; label = @5
              local.get 0
              i32.const -183970
              i32.add
              i32.const 14
              i32.ge_u
              br_if 0 (;@5;)
              i32.const 0
              return
            end
            block  ;; label = @5
              local.get 0
              i32.const -2
              i32.and
              i32.const 178206
              i32.ne
              br_if 0 (;@5;)
              i32.const 0
              return
            end
            local.get 0
            i32.const -32
            i32.and
            i32.const 173792
            i32.ne
            br_if 1 (;@3;)
            i32.const 0
            return
          end
          local.get 0
          i32.const 1054042
          i32.const 44
          i32.const 1054130
          i32.const 196
          i32.const 1054326
          i32.const 450
          call $_ZN4core7unicode9printable5check17hbed7f4d4530f0fe8E
          return
        end
        i32.const 0
        local.set 1
        local.get 0
        i32.const -177978
        i32.add
        i32.const 6
        i32.lt_u
        br_if 0 (;@2;)
        local.get 0
        i32.const -1114112
        i32.add
        i32.const -196112
        i32.lt_u
        local.set 1
      end
      local.get 1
      return
    end
    local.get 0
    i32.const 1053372
    i32.const 40
    i32.const 1053452
    i32.const 287
    i32.const 1053739
    i32.const 303
    call $_ZN4core7unicode9printable5check17hbed7f4d4530f0fe8E)
  (func $_ZN42_$LT$str$u20$as$u20$core..fmt..Display$GT$3fmt17h100f1ae78e96c78dE (type 1) (param i32 i32 i32) (result i32)
    local.get 2
    local.get 0
    local.get 1
    call $_ZN4core3fmt9Formatter3pad17h7e3d356fb87d6820E)
  (func $_ZN41_$LT$char$u20$as$u20$core..fmt..Debug$GT$3fmt17h19265808cb3f0ce9E (type 2) (param i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32)
    i32.const 1
    local.set 2
    block  ;; label = @1
      local.get 1
      i32.load
      local.tee 3
      i32.const 39
      local.get 1
      i32.load offset=4
      i32.load offset=16
      local.tee 4
      call_indirect (type 2)
      br_if 0 (;@1;)
      i32.const 1114114
      local.set 2
      i32.const 48
      local.set 5
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      block  ;; label = @10
                        local.get 0
                        i32.load
                        local.tee 1
                        br_table 8 (;@2;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 2 (;@8;) 4 (;@6;) 1 (;@9;) 1 (;@9;) 3 (;@7;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 1 (;@9;) 5 (;@5;) 0 (;@10;)
                      end
                      local.get 1
                      i32.const 92
                      i32.eq
                      br_if 4 (;@5;)
                    end
                    local.get 1
                    call $_ZN4core7unicode12unicode_data15grapheme_extend6lookup17h24f97697985c6c02E
                    i32.eqz
                    br_if 4 (;@4;)
                    local.get 1
                    i32.const 1
                    i32.or
                    i32.clz
                    i32.const 2
                    i32.shr_u
                    i32.const 7
                    i32.xor
                    local.set 5
                    br 5 (;@3;)
                  end
                  i32.const 116
                  local.set 5
                  br 5 (;@2;)
                end
                i32.const 114
                local.set 5
                br 4 (;@2;)
              end
              i32.const 110
              local.set 5
              br 3 (;@2;)
            end
            local.get 1
            local.set 5
            br 2 (;@2;)
          end
          i32.const 1114113
          local.set 2
          block  ;; label = @4
            local.get 1
            call $_ZN4core7unicode9printable12is_printable17h237b4d5f567ccf43E
            i32.eqz
            br_if 0 (;@4;)
            local.get 1
            local.set 5
            br 2 (;@2;)
          end
          local.get 1
          i32.const 1
          i32.or
          i32.clz
          i32.const 2
          i32.shr_u
          i32.const 7
          i32.xor
          local.set 5
        end
        local.get 1
        local.set 2
      end
      i32.const 5
      local.set 6
      loop  ;; label = @2
        local.get 6
        local.set 7
        local.get 2
        local.set 1
        i32.const 1114113
        local.set 2
        i32.const 92
        local.set 0
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  local.get 1
                  i32.const -1114112
                  i32.add
                  i32.const 3
                  local.get 1
                  i32.const 1114111
                  i32.gt_u
                  select
                  br_table 2 (;@5;) 1 (;@6;) 4 (;@3;) 0 (;@7;) 2 (;@5;)
                end
                i32.const 0
                local.set 6
                i32.const 125
                local.set 0
                local.get 1
                local.set 2
                block  ;; label = @7
                  block  ;; label = @8
                    block  ;; label = @9
                      local.get 7
                      i32.const 255
                      i32.and
                      br_table 4 (;@5;) 6 (;@3;) 5 (;@4;) 0 (;@9;) 1 (;@8;) 2 (;@7;) 4 (;@5;)
                    end
                    i32.const 2
                    local.set 6
                    i32.const 123
                    local.set 0
                    local.get 1
                    local.set 2
                    br 5 (;@3;)
                  end
                  i32.const 3
                  local.set 6
                  i32.const 117
                  local.set 0
                  local.get 1
                  local.set 2
                  br 4 (;@3;)
                end
                i32.const 4
                local.set 6
                i32.const 92
                local.set 0
                local.get 1
                local.set 2
                br 3 (;@3;)
              end
              i32.const 1114112
              local.set 2
              local.get 5
              local.set 0
              local.get 7
              local.set 6
              local.get 5
              i32.const 1114112
              i32.ne
              br_if 2 (;@3;)
            end
            local.get 3
            i32.const 39
            local.get 4
            call_indirect (type 2)
            local.set 2
            br 3 (;@1;)
          end
          local.get 7
          i32.const 1
          local.get 5
          select
          local.set 6
          i32.const 48
          i32.const 87
          local.get 1
          local.get 5
          i32.const 2
          i32.shl
          i32.shr_u
          i32.const 15
          i32.and
          local.tee 2
          i32.const 10
          i32.lt_u
          select
          local.get 2
          i32.add
          local.set 0
          local.get 1
          local.set 2
          i32.const 0
          local.get 5
          i32.const -1
          i32.add
          local.tee 1
          local.get 1
          local.get 5
          i32.gt_u
          select
          local.set 5
        end
        local.get 3
        local.get 0
        local.get 4
        call_indirect (type 2)
        i32.eqz
        br_if 0 (;@2;)
      end
      i32.const 1
      return
    end
    local.get 2)
  (func $_ZN4core5slice5index29slice_start_index_len_fail_rt17he819c5908e9fdb1fE (type 3) (param i32 i32 i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    local.get 3
    local.get 1
    i32.store offset=4
    local.get 3
    local.get 0
    i32.store
    local.get 3
    i32.const 8
    i32.add
    i32.const 12
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 28
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 32
    i32.add
    i32.const 12
    i32.add
    i32.const 4
    i32.store
    local.get 3
    i32.const 1052932
    i32.store offset=16
    local.get 3
    i32.const 0
    i32.store offset=8
    local.get 3
    i32.const 4
    i32.store offset=36
    local.get 3
    local.get 3
    i32.const 32
    i32.add
    i32.store offset=24
    local.get 3
    local.get 3
    i32.const 4
    i32.add
    i32.store offset=40
    local.get 3
    local.get 3
    i32.store offset=32
    local.get 3
    i32.const 8
    i32.add
    local.get 2
    call $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E
    unreachable)
  (func $_ZN4core5slice5index27slice_end_index_len_fail_rt17h480eb8e5cf49cf70E (type 3) (param i32 i32 i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    local.get 3
    local.get 1
    i32.store offset=4
    local.get 3
    local.get 0
    i32.store
    local.get 3
    i32.const 8
    i32.add
    i32.const 12
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 28
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 32
    i32.add
    i32.const 12
    i32.add
    i32.const 4
    i32.store
    local.get 3
    i32.const 1052964
    i32.store offset=16
    local.get 3
    i32.const 0
    i32.store offset=8
    local.get 3
    i32.const 4
    i32.store offset=36
    local.get 3
    local.get 3
    i32.const 32
    i32.add
    i32.store offset=24
    local.get 3
    local.get 3
    i32.const 4
    i32.add
    i32.store offset=40
    local.get 3
    local.get 3
    i32.store offset=32
    local.get 3
    i32.const 8
    i32.add
    local.get 2
    call $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E
    unreachable)
  (func $_ZN4core5slice5index25slice_index_order_fail_rt17hdbe4433d56ef37a3E (type 3) (param i32 i32 i32)
    (local i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    local.get 3
    local.get 1
    i32.store offset=4
    local.get 3
    local.get 0
    i32.store
    local.get 3
    i32.const 8
    i32.add
    i32.const 12
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 28
    i32.add
    i32.const 2
    i32.store
    local.get 3
    i32.const 32
    i32.add
    i32.const 12
    i32.add
    i32.const 4
    i32.store
    local.get 3
    i32.const 1053016
    i32.store offset=16
    local.get 3
    i32.const 0
    i32.store offset=8
    local.get 3
    i32.const 4
    i32.store offset=36
    local.get 3
    local.get 3
    i32.const 32
    i32.add
    i32.store offset=24
    local.get 3
    local.get 3
    i32.const 4
    i32.add
    i32.store offset=40
    local.get 3
    local.get 3
    i32.store offset=32
    local.get 3
    i32.const 8
    i32.add
    local.get 2
    call $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E
    unreachable)
  (func $_ZN4core3str19slice_error_fail_rt17h8712bc75e2bb7542E (type 15) (param i32 i32 i32 i32 i32)
    (local i32 i32 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 112
    i32.sub
    local.tee 5
    global.set $__stack_pointer
    local.get 5
    local.get 3
    i32.store offset=12
    local.get 5
    local.get 2
    i32.store offset=8
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          i32.const 257
          i32.lt_u
          br_if 0 (;@3;)
          i32.const 256
          local.set 6
          block  ;; label = @4
            local.get 0
            i32.load8_s offset=256
            i32.const -65
            i32.gt_s
            br_if 0 (;@4;)
            i32.const 255
            local.set 6
            local.get 0
            i32.load8_s offset=255
            i32.const -65
            i32.gt_s
            br_if 0 (;@4;)
            i32.const 254
            local.set 6
            local.get 0
            i32.load8_s offset=254
            i32.const -65
            i32.gt_s
            br_if 0 (;@4;)
            i32.const 253
            local.set 6
          end
          local.get 0
          local.get 6
          i32.add
          i32.load8_s
          i32.const -65
          i32.le_s
          br_if 1 (;@2;)
          local.get 5
          local.get 6
          i32.store offset=20
          local.get 5
          local.get 0
          i32.store offset=16
          i32.const 5
          local.set 6
          i32.const 1053032
          local.set 7
          br 2 (;@1;)
        end
        local.get 5
        local.get 1
        i32.store offset=20
        local.get 5
        local.get 0
        i32.store offset=16
        i32.const 0
        local.set 6
        i32.const 1052176
        local.set 7
        br 1 (;@1;)
      end
      local.get 0
      local.get 1
      i32.const 0
      local.get 6
      local.get 4
      call $_ZN4core3str16slice_error_fail17h48e751b16631e034E
      unreachable
    end
    local.get 5
    local.get 6
    i32.store offset=28
    local.get 5
    local.get 7
    i32.store offset=24
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 2
            local.get 1
            i32.gt_u
            local.tee 6
            br_if 0 (;@4;)
            local.get 3
            local.get 1
            i32.gt_u
            br_if 0 (;@4;)
            block  ;; label = @5
              block  ;; label = @6
                block  ;; label = @7
                  block  ;; label = @8
                    local.get 2
                    local.get 3
                    i32.gt_u
                    br_if 0 (;@8;)
                    block  ;; label = @9
                      block  ;; label = @10
                        local.get 2
                        i32.eqz
                        br_if 0 (;@10;)
                        block  ;; label = @11
                          local.get 2
                          local.get 1
                          i32.lt_u
                          br_if 0 (;@11;)
                          local.get 2
                          local.get 1
                          i32.eq
                          br_if 1 (;@10;)
                          br 2 (;@9;)
                        end
                        local.get 0
                        local.get 2
                        i32.add
                        i32.load8_s
                        i32.const -64
                        i32.lt_s
                        br_if 1 (;@9;)
                      end
                      local.get 3
                      local.set 2
                    end
                    local.get 5
                    local.get 2
                    i32.store offset=32
                    local.get 1
                    local.set 3
                    block  ;; label = @9
                      local.get 2
                      local.get 1
                      i32.ge_u
                      br_if 0 (;@9;)
                      local.get 2
                      i32.const 1
                      i32.add
                      local.tee 6
                      i32.const 0
                      local.get 2
                      i32.const -3
                      i32.add
                      local.tee 3
                      local.get 3
                      local.get 2
                      i32.gt_u
                      select
                      local.tee 3
                      i32.lt_u
                      br_if 6 (;@3;)
                      block  ;; label = @10
                        local.get 3
                        local.get 6
                        i32.eq
                        br_if 0 (;@10;)
                        local.get 0
                        local.get 6
                        i32.add
                        local.get 0
                        local.get 3
                        i32.add
                        local.tee 8
                        i32.sub
                        local.set 6
                        block  ;; label = @11
                          local.get 0
                          local.get 2
                          i32.add
                          local.tee 9
                          i32.load8_s
                          i32.const -65
                          i32.le_s
                          br_if 0 (;@11;)
                          local.get 6
                          i32.const -1
                          i32.add
                          local.set 7
                          br 1 (;@10;)
                        end
                        local.get 3
                        local.get 2
                        i32.eq
                        br_if 0 (;@10;)
                        block  ;; label = @11
                          local.get 9
                          i32.const -1
                          i32.add
                          local.tee 2
                          i32.load8_s
                          i32.const -65
                          i32.le_s
                          br_if 0 (;@11;)
                          local.get 6
                          i32.const -2
                          i32.add
                          local.set 7
                          br 1 (;@10;)
                        end
                        local.get 8
                        local.get 2
                        i32.eq
                        br_if 0 (;@10;)
                        block  ;; label = @11
                          local.get 2
                          i32.const -1
                          i32.add
                          local.tee 2
                          i32.load8_s
                          i32.const -65
                          i32.le_s
                          br_if 0 (;@11;)
                          local.get 6
                          i32.const -3
                          i32.add
                          local.set 7
                          br 1 (;@10;)
                        end
                        local.get 8
                        local.get 2
                        i32.eq
                        br_if 0 (;@10;)
                        block  ;; label = @11
                          local.get 2
                          i32.const -1
                          i32.add
                          local.tee 2
                          i32.load8_s
                          i32.const -65
                          i32.le_s
                          br_if 0 (;@11;)
                          local.get 6
                          i32.const -4
                          i32.add
                          local.set 7
                          br 1 (;@10;)
                        end
                        local.get 8
                        local.get 2
                        i32.eq
                        br_if 0 (;@10;)
                        local.get 6
                        i32.const -5
                        i32.add
                        local.set 7
                      end
                      local.get 7
                      local.get 3
                      i32.add
                      local.set 3
                    end
                    block  ;; label = @9
                      local.get 3
                      i32.eqz
                      br_if 0 (;@9;)
                      block  ;; label = @10
                        local.get 3
                        local.get 1
                        i32.lt_u
                        br_if 0 (;@10;)
                        local.get 3
                        local.get 1
                        i32.eq
                        br_if 1 (;@9;)
                        br 9 (;@1;)
                      end
                      local.get 0
                      local.get 3
                      i32.add
                      i32.load8_s
                      i32.const -65
                      i32.le_s
                      br_if 8 (;@1;)
                    end
                    local.get 3
                    local.get 1
                    i32.eq
                    br_if 6 (;@2;)
                    block  ;; label = @9
                      block  ;; label = @10
                        local.get 0
                        local.get 3
                        i32.add
                        local.tee 2
                        i32.load8_s
                        local.tee 1
                        i32.const -1
                        i32.gt_s
                        br_if 0 (;@10;)
                        local.get 2
                        i32.load8_u offset=1
                        i32.const 63
                        i32.and
                        local.set 0
                        local.get 1
                        i32.const 31
                        i32.and
                        local.set 6
                        local.get 1
                        i32.const -33
                        i32.gt_u
                        br_if 1 (;@9;)
                        local.get 6
                        i32.const 6
                        i32.shl
                        local.get 0
                        i32.or
                        local.set 2
                        br 4 (;@6;)
                      end
                      local.get 5
                      local.get 1
                      i32.const 255
                      i32.and
                      i32.store offset=36
                      i32.const 1
                      local.set 1
                      br 4 (;@5;)
                    end
                    local.get 0
                    i32.const 6
                    i32.shl
                    local.get 2
                    i32.load8_u offset=2
                    i32.const 63
                    i32.and
                    i32.or
                    local.set 0
                    local.get 1
                    i32.const -16
                    i32.ge_u
                    br_if 1 (;@7;)
                    local.get 0
                    local.get 6
                    i32.const 12
                    i32.shl
                    i32.or
                    local.set 2
                    br 2 (;@6;)
                  end
                  local.get 5
                  i32.const 100
                  i32.add
                  i32.const 28
                  i32.store
                  local.get 5
                  i32.const 72
                  i32.add
                  i32.const 20
                  i32.add
                  i32.const 28
                  i32.store
                  local.get 5
                  i32.const 72
                  i32.add
                  i32.const 12
                  i32.add
                  i32.const 4
                  i32.store
                  local.get 5
                  i32.const 48
                  i32.add
                  i32.const 12
                  i32.add
                  i32.const 4
                  i32.store
                  local.get 5
                  i32.const 48
                  i32.add
                  i32.const 20
                  i32.add
                  i32.const 4
                  i32.store
                  local.get 5
                  i32.const 1053132
                  i32.store offset=56
                  local.get 5
                  i32.const 0
                  i32.store offset=48
                  local.get 5
                  i32.const 4
                  i32.store offset=76
                  local.get 5
                  local.get 5
                  i32.const 72
                  i32.add
                  i32.store offset=64
                  local.get 5
                  local.get 5
                  i32.const 24
                  i32.add
                  i32.store offset=96
                  local.get 5
                  local.get 5
                  i32.const 16
                  i32.add
                  i32.store offset=88
                  local.get 5
                  local.get 5
                  i32.const 12
                  i32.add
                  i32.store offset=80
                  local.get 5
                  local.get 5
                  i32.const 8
                  i32.add
                  i32.store offset=72
                  local.get 5
                  i32.const 48
                  i32.add
                  local.get 4
                  call $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E
                  unreachable
                end
                local.get 0
                i32.const 6
                i32.shl
                local.get 2
                i32.load8_u offset=3
                i32.const 63
                i32.and
                i32.or
                local.get 6
                i32.const 18
                i32.shl
                i32.const 1835008
                i32.and
                i32.or
                local.tee 2
                i32.const 1114112
                i32.eq
                br_if 4 (;@2;)
              end
              local.get 5
              local.get 2
              i32.store offset=36
              i32.const 1
              local.set 1
              local.get 2
              i32.const 128
              i32.lt_u
              br_if 0 (;@5;)
              i32.const 2
              local.set 1
              local.get 2
              i32.const 2047
              i32.le_u
              br_if 0 (;@5;)
              i32.const 3
              i32.const 4
              local.get 2
              i32.const 65536
              i32.lt_u
              select
              local.set 1
            end
            local.get 5
            local.get 3
            i32.store offset=40
            local.get 5
            local.get 1
            local.get 3
            i32.add
            i32.store offset=44
            local.get 5
            i32.const 48
            i32.add
            i32.const 12
            i32.add
            i32.const 5
            i32.store
            local.get 5
            i32.const 48
            i32.add
            i32.const 20
            i32.add
            i32.const 5
            i32.store
            local.get 5
            i32.const 108
            i32.add
            i32.const 28
            i32.store
            local.get 5
            i32.const 100
            i32.add
            i32.const 28
            i32.store
            local.get 5
            i32.const 72
            i32.add
            i32.const 20
            i32.add
            i32.const 30
            i32.store
            local.get 5
            i32.const 72
            i32.add
            i32.const 12
            i32.add
            i32.const 31
            i32.store
            local.get 5
            i32.const 1053216
            i32.store offset=56
            local.get 5
            i32.const 0
            i32.store offset=48
            local.get 5
            i32.const 4
            i32.store offset=76
            local.get 5
            local.get 5
            i32.const 72
            i32.add
            i32.store offset=64
            local.get 5
            local.get 5
            i32.const 24
            i32.add
            i32.store offset=104
            local.get 5
            local.get 5
            i32.const 16
            i32.add
            i32.store offset=96
            local.get 5
            local.get 5
            i32.const 40
            i32.add
            i32.store offset=88
            local.get 5
            local.get 5
            i32.const 36
            i32.add
            i32.store offset=80
            local.get 5
            local.get 5
            i32.const 32
            i32.add
            i32.store offset=72
            local.get 5
            i32.const 48
            i32.add
            local.get 4
            call $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E
            unreachable
          end
          local.get 5
          local.get 2
          local.get 3
          local.get 6
          select
          i32.store offset=40
          local.get 5
          i32.const 48
          i32.add
          i32.const 12
          i32.add
          i32.const 3
          i32.store
          local.get 5
          i32.const 48
          i32.add
          i32.const 20
          i32.add
          i32.const 3
          i32.store
          local.get 5
          i32.const 72
          i32.add
          i32.const 20
          i32.add
          i32.const 28
          i32.store
          local.get 5
          i32.const 72
          i32.add
          i32.const 12
          i32.add
          i32.const 28
          i32.store
          local.get 5
          i32.const 1053072
          i32.store offset=56
          local.get 5
          i32.const 0
          i32.store offset=48
          local.get 5
          i32.const 4
          i32.store offset=76
          local.get 5
          local.get 5
          i32.const 72
          i32.add
          i32.store offset=64
          local.get 5
          local.get 5
          i32.const 24
          i32.add
          i32.store offset=88
          local.get 5
          local.get 5
          i32.const 16
          i32.add
          i32.store offset=80
          local.get 5
          local.get 5
          i32.const 40
          i32.add
          i32.store offset=72
          local.get 5
          i32.const 48
          i32.add
          local.get 4
          call $_ZN4core9panicking9panic_fmt17he51e8f27bc0b24b2E
          unreachable
        end
        local.get 3
        local.get 6
        i32.const 1053284
        call $_ZN4core5slice5index22slice_index_order_fail17hf30e3ed6d5fa1a7bE
        unreachable
      end
      i32.const 1051789
      i32.const 43
      local.get 4
      call $_ZN4core9panicking5panic17he2fadc50abd7e245E
      unreachable
    end
    local.get 0
    local.get 1
    local.get 3
    local.get 1
    local.get 4
    call $_ZN4core3str16slice_error_fail17h48e751b16631e034E
    unreachable)
  (func $_ZN4core7unicode9printable5check17hbed7f4d4530f0fe8E (type 23) (param i32 i32 i32 i32 i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32)
    i32.const 1
    local.set 7
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.eqz
        br_if 0 (;@2;)
        local.get 1
        local.get 2
        i32.const 1
        i32.shl
        i32.add
        local.set 8
        local.get 0
        i32.const 65280
        i32.and
        i32.const 8
        i32.shr_u
        local.set 9
        i32.const 0
        local.set 10
        local.get 0
        i32.const 255
        i32.and
        local.set 11
        loop  ;; label = @3
          local.get 1
          i32.const 2
          i32.add
          local.set 12
          local.get 10
          local.get 1
          i32.load8_u offset=1
          local.tee 2
          i32.add
          local.set 13
          block  ;; label = @4
            local.get 1
            i32.load8_u
            local.tee 1
            local.get 9
            i32.eq
            br_if 0 (;@4;)
            local.get 1
            local.get 9
            i32.gt_u
            br_if 2 (;@2;)
            local.get 13
            local.set 10
            local.get 12
            local.set 1
            local.get 12
            local.get 8
            i32.eq
            br_if 2 (;@2;)
            br 1 (;@3;)
          end
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 13
                local.get 10
                i32.lt_u
                br_if 0 (;@6;)
                local.get 13
                local.get 4
                i32.gt_u
                br_if 1 (;@5;)
                local.get 3
                local.get 10
                i32.add
                local.set 1
                loop  ;; label = @7
                  local.get 2
                  i32.eqz
                  br_if 3 (;@4;)
                  local.get 2
                  i32.const -1
                  i32.add
                  local.set 2
                  local.get 1
                  i32.load8_u
                  local.set 10
                  local.get 1
                  i32.const 1
                  i32.add
                  local.set 1
                  local.get 10
                  local.get 11
                  i32.ne
                  br_if 0 (;@7;)
                end
                i32.const 0
                local.set 7
                br 5 (;@1;)
              end
              local.get 10
              local.get 13
              i32.const 1053340
              call $_ZN4core5slice5index22slice_index_order_fail17hf30e3ed6d5fa1a7bE
              unreachable
            end
            local.get 13
            local.get 4
            i32.const 1053340
            call $_ZN4core5slice5index24slice_end_index_len_fail17ha4ac89ea73386400E
            unreachable
          end
          local.get 13
          local.set 10
          local.get 12
          local.set 1
          local.get 12
          local.get 8
          i32.ne
          br_if 0 (;@3;)
        end
      end
      local.get 6
      i32.eqz
      br_if 0 (;@1;)
      local.get 5
      local.get 6
      i32.add
      local.set 11
      local.get 0
      i32.const 65535
      i32.and
      local.set 1
      i32.const 1
      local.set 7
      block  ;; label = @2
        loop  ;; label = @3
          local.get 5
          i32.const 1
          i32.add
          local.set 10
          block  ;; label = @4
            block  ;; label = @5
              local.get 5
              i32.load8_u
              local.tee 2
              i32.extend8_s
              local.tee 13
              i32.const 0
              i32.lt_s
              br_if 0 (;@5;)
              local.get 10
              local.set 5
              br 1 (;@4;)
            end
            local.get 10
            local.get 11
            i32.eq
            br_if 2 (;@2;)
            local.get 13
            i32.const 127
            i32.and
            i32.const 8
            i32.shl
            local.get 5
            i32.load8_u offset=1
            i32.or
            local.set 2
            local.get 5
            i32.const 2
            i32.add
            local.set 5
          end
          local.get 1
          local.get 2
          i32.sub
          local.tee 1
          i32.const 0
          i32.lt_s
          br_if 2 (;@1;)
          local.get 7
          i32.const 1
          i32.xor
          local.set 7
          local.get 5
          local.get 11
          i32.ne
          br_if 0 (;@3;)
          br 2 (;@1;)
        end
      end
      i32.const 1051789
      i32.const 43
      i32.const 1053356
      call $_ZN4core9panicking5panic17he2fadc50abd7e245E
      unreachable
    end
    local.get 7
    i32.const 1
    i32.and)
  (func $_ZN4core3fmt3num50_$LT$impl$u20$core..fmt..Debug$u20$for$u20$u64$GT$3fmt17h4e1f5c934b492f6aE (type 2) (param i32 i32) (result i32)
    (local i32 i32 i64)
    global.get $__stack_pointer
    i32.const 128
    i32.sub
    local.tee 2
    global.set $__stack_pointer
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              local.get 1
              i32.load offset=24
              local.tee 3
              i32.const 16
              i32.and
              br_if 0 (;@5;)
              local.get 3
              i32.const 32
              i32.and
              br_if 1 (;@4;)
              local.get 0
              i64.load
              i32.const 1
              local.get 1
              call $_ZN4core3fmt3num3imp7fmt_u6417h5f097ceb54602c9bE
              local.set 0
              br 4 (;@1;)
            end
            local.get 0
            i64.load
            local.set 4
            i32.const 0
            local.set 0
            loop  ;; label = @5
              local.get 2
              local.get 0
              i32.add
              i32.const 127
              i32.add
              i32.const 48
              i32.const 87
              local.get 4
              i32.wrap_i64
              i32.const 15
              i32.and
              local.tee 3
              i32.const 10
              i32.lt_u
              select
              local.get 3
              i32.add
              i32.store8
              local.get 0
              i32.const -1
              i32.add
              local.set 0
              local.get 4
              i64.const 15
              i64.gt_u
              local.set 3
              local.get 4
              i64.const 4
              i64.shr_u
              local.set 4
              local.get 3
              br_if 0 (;@5;)
            end
            local.get 0
            i32.const 128
            i32.add
            local.tee 3
            i32.const 129
            i32.ge_u
            br_if 1 (;@3;)
            local.get 1
            i32.const 1
            i32.const 1052552
            i32.const 2
            local.get 2
            local.get 0
            i32.add
            i32.const 128
            i32.add
            i32.const 0
            local.get 0
            i32.sub
            call $_ZN4core3fmt9Formatter12pad_integral17hf0a49a9f5edc1e33E
            local.set 0
            br 3 (;@1;)
          end
          local.get 0
          i64.load
          local.set 4
          i32.const 0
          local.set 0
          loop  ;; label = @4
            local.get 2
            local.get 0
            i32.add
            i32.const 127
            i32.add
            i32.const 48
            i32.const 55
            local.get 4
            i32.wrap_i64
            i32.const 15
            i32.and
            local.tee 3
            i32.const 10
            i32.lt_u
            select
            local.get 3
            i32.add
            i32.store8
            local.get 0
            i32.const -1
            i32.add
            local.set 0
            local.get 4
            i64.const 15
            i64.gt_u
            local.set 3
            local.get 4
            i64.const 4
            i64.shr_u
            local.set 4
            local.get 3
            br_if 0 (;@4;)
          end
          local.get 0
          i32.const 128
          i32.add
          local.tee 3
          i32.const 129
          i32.ge_u
          br_if 1 (;@2;)
          local.get 1
          i32.const 1
          i32.const 1052552
          i32.const 2
          local.get 2
          local.get 0
          i32.add
          i32.const 128
          i32.add
          i32.const 0
          local.get 0
          i32.sub
          call $_ZN4core3fmt9Formatter12pad_integral17hf0a49a9f5edc1e33E
          local.set 0
          br 2 (;@1;)
        end
        local.get 3
        i32.const 128
        i32.const 1052536
        call $_ZN4core5slice5index26slice_start_index_len_fail17h61955cceb128f310E
        unreachable
      end
      local.get 3
      i32.const 128
      i32.const 1052536
      call $_ZN4core5slice5index26slice_start_index_len_fail17h61955cceb128f310E
      unreachable
    end
    local.get 2
    i32.const 128
    i32.add
    global.set $__stack_pointer
    local.get 0)
  (func $_ZN4core3fmt3num3imp7fmt_u6417h5f097ceb54602c9bE (type 24) (param i64 i32 i32) (result i32)
    (local i32 i32 i64 i32 i32 i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 3
    global.set $__stack_pointer
    i32.const 39
    local.set 4
    block  ;; label = @1
      block  ;; label = @2
        local.get 0
        i64.const 10000
        i64.ge_u
        br_if 0 (;@2;)
        local.get 0
        local.set 5
        br 1 (;@1;)
      end
      i32.const 39
      local.set 4
      loop  ;; label = @2
        local.get 3
        i32.const 9
        i32.add
        local.get 4
        i32.add
        local.tee 6
        i32.const -4
        i32.add
        local.get 0
        local.get 0
        i64.const 10000
        i64.div_u
        local.tee 5
        i64.const 10000
        i64.mul
        i64.sub
        i32.wrap_i64
        local.tee 7
        i32.const 65535
        i32.and
        i32.const 100
        i32.div_u
        local.tee 8
        i32.const 1
        i32.shl
        i32.const 1052554
        i32.add
        i32.load16_u align=1
        i32.store16 align=1
        local.get 6
        i32.const -2
        i32.add
        local.get 7
        local.get 8
        i32.const 100
        i32.mul
        i32.sub
        i32.const 65535
        i32.and
        i32.const 1
        i32.shl
        i32.const 1052554
        i32.add
        i32.load16_u align=1
        i32.store16 align=1
        local.get 4
        i32.const -4
        i32.add
        local.set 4
        local.get 0
        i64.const 99999999
        i64.gt_u
        local.set 6
        local.get 5
        local.set 0
        local.get 6
        br_if 0 (;@2;)
      end
    end
    block  ;; label = @1
      local.get 5
      i32.wrap_i64
      local.tee 6
      i32.const 99
      i32.le_u
      br_if 0 (;@1;)
      local.get 3
      i32.const 9
      i32.add
      local.get 4
      i32.const -2
      i32.add
      local.tee 4
      i32.add
      local.get 5
      i32.wrap_i64
      local.tee 6
      local.get 6
      i32.const 65535
      i32.and
      i32.const 100
      i32.div_u
      local.tee 6
      i32.const 100
      i32.mul
      i32.sub
      i32.const 65535
      i32.and
      i32.const 1
      i32.shl
      i32.const 1052554
      i32.add
      i32.load16_u align=1
      i32.store16 align=1
    end
    block  ;; label = @1
      block  ;; label = @2
        local.get 6
        i32.const 10
        i32.lt_u
        br_if 0 (;@2;)
        local.get 3
        i32.const 9
        i32.add
        local.get 4
        i32.const -2
        i32.add
        local.tee 4
        i32.add
        local.get 6
        i32.const 1
        i32.shl
        i32.const 1052554
        i32.add
        i32.load16_u align=1
        i32.store16 align=1
        br 1 (;@1;)
      end
      local.get 3
      i32.const 9
      i32.add
      local.get 4
      i32.const -1
      i32.add
      local.tee 4
      i32.add
      local.get 6
      i32.const 48
      i32.add
      i32.store8
    end
    local.get 2
    local.get 1
    i32.const 1052176
    i32.const 0
    local.get 3
    i32.const 9
    i32.add
    local.get 4
    i32.add
    i32.const 39
    local.get 4
    i32.sub
    call $_ZN4core3fmt9Formatter12pad_integral17hf0a49a9f5edc1e33E
    local.set 4
    local.get 3
    i32.const 48
    i32.add
    global.set $__stack_pointer
    local.get 4)
  (func $_ZN4core3fmt5float52_$LT$impl$u20$core..fmt..Display$u20$for$u20$f32$GT$3fmt17h9d95cc5371839148E (type 2) (param i32 i32) (result i32)
    (local i32 f32)
    local.get 1
    i32.load offset=24
    i32.const 1
    i32.and
    local.set 2
    local.get 0
    f32.load
    local.set 3
    block  ;; label = @1
      local.get 1
      i32.load offset=16
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      local.get 3
      local.get 2
      local.get 1
      i32.const 20
      i32.add
      i32.load
      call $_ZN4core3fmt5float29float_to_decimal_common_exact17h887dd65cb992c326E
      return
    end
    local.get 1
    local.get 3
    local.get 2
    i32.const 0
    call $_ZN4core3fmt5float32float_to_decimal_common_shortest17hab759cd6e3d171f2E)
  (func $_ZN4core3fmt5float52_$LT$impl$u20$core..fmt..Display$u20$for$u20$f64$GT$3fmt17h33a38e7afb79522fE (type 2) (param i32 i32) (result i32)
    (local i32 f64)
    local.get 1
    i32.load offset=24
    i32.const 1
    i32.and
    local.set 2
    local.get 0
    f64.load
    local.set 3
    block  ;; label = @1
      local.get 1
      i32.load offset=16
      i32.eqz
      br_if 0 (;@1;)
      local.get 1
      local.get 3
      local.get 2
      local.get 1
      i32.const 20
      i32.add
      i32.load
      call $_ZN4core3fmt5float29float_to_decimal_common_exact17hdedd1adf7adb591bE
      return
    end
    local.get 1
    local.get 3
    local.get 2
    i32.const 0
    call $_ZN4core3fmt5float32float_to_decimal_common_shortest17h0203f2e3c1aa9f40E)
  (func $_ZN4core3fmt3num3imp52_$LT$impl$u20$core..fmt..Display$u20$for$u20$i32$GT$3fmt17h2a64a79a6a2a885aE (type 2) (param i32 i32) (result i32)
    local.get 0
    i32.load
    local.tee 0
    i64.extend_i32_u
    i64.const 0
    local.get 0
    i64.extend_i32_s
    i64.sub
    local.get 0
    i32.const -1
    i32.gt_s
    local.tee 0
    select
    local.get 0
    local.get 1
    call $_ZN4core3fmt3num3imp7fmt_u6417h5f097ceb54602c9bE)
  (func $_ZN4core3fmt3num3imp52_$LT$impl$u20$core..fmt..Display$u20$for$u20$i64$GT$3fmt17h25d8f33ead29d25bE (type 2) (param i32 i32) (result i32)
    (local i64 i64)
    local.get 0
    i64.load
    local.tee 2
    local.get 2
    i64.const 63
    i64.shr_s
    local.tee 3
    i64.xor
    local.get 3
    i64.sub
    local.get 2
    i64.const -1
    i64.gt_s
    local.get 1
    call $_ZN4core3fmt3num3imp7fmt_u6417h5f097ceb54602c9bE)
  (func $_ZN53_$LT$core..fmt..Error$u20$as$u20$core..fmt..Debug$GT$3fmt17hca756aa95c0138eeE (type 2) (param i32 i32) (result i32)
    local.get 1
    i32.load
    i32.const 1054946
    i32.const 5
    local.get 1
    i32.load offset=4
    i32.load offset=12
    call_indirect (type 1))
  (func $_ZN42_$LT$$RF$T$u20$as$u20$core..fmt..Debug$GT$3fmt17h1b991df98dc9a0fdE (type 2) (param i32 i32) (result i32)
    local.get 0
    i32.load
    local.get 1
    call $_ZN4core3fmt3num50_$LT$impl$u20$core..fmt..Debug$u20$for$u20$u64$GT$3fmt17h4e1f5c934b492f6aE)
  (func $memset (type 1) (param i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    call $_ZN17compiler_builtins3mem6memset17h2d50fc5881661f2fE)
  (func $_ZN17compiler_builtins3mem6memcpy17hcd54104a1e4c5b0dE (type 1) (param i32 i32 i32) (result i32)
    (local i32 i32 i32 i32 i32 i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.const 15
        i32.gt_u
        br_if 0 (;@2;)
        local.get 0
        local.set 3
        br 1 (;@1;)
      end
      local.get 0
      i32.const 0
      local.get 0
      i32.sub
      i32.const 3
      i32.and
      local.tee 4
      i32.add
      local.set 5
      block  ;; label = @2
        local.get 4
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        local.set 3
        local.get 1
        local.set 6
        loop  ;; label = @3
          local.get 3
          local.get 6
          i32.load8_u
          i32.store8
          local.get 6
          i32.const 1
          i32.add
          local.set 6
          local.get 3
          i32.const 1
          i32.add
          local.tee 3
          local.get 5
          i32.lt_u
          br_if 0 (;@3;)
        end
      end
      local.get 5
      local.get 2
      local.get 4
      i32.sub
      local.tee 7
      i32.const -4
      i32.and
      local.tee 8
      i32.add
      local.set 3
      block  ;; label = @2
        block  ;; label = @3
          local.get 1
          local.get 4
          i32.add
          local.tee 9
          i32.const 3
          i32.and
          local.tee 6
          i32.eqz
          br_if 0 (;@3;)
          local.get 8
          i32.const 1
          i32.lt_s
          br_if 1 (;@2;)
          local.get 9
          i32.const -4
          i32.and
          local.tee 10
          i32.const 4
          i32.add
          local.set 1
          i32.const 0
          local.get 6
          i32.const 3
          i32.shl
          local.tee 2
          i32.sub
          i32.const 24
          i32.and
          local.set 4
          local.get 10
          i32.load
          local.set 6
          loop  ;; label = @4
            local.get 5
            local.get 6
            local.get 2
            i32.shr_u
            local.get 1
            i32.load
            local.tee 6
            local.get 4
            i32.shl
            i32.or
            i32.store
            local.get 1
            i32.const 4
            i32.add
            local.set 1
            local.get 5
            i32.const 4
            i32.add
            local.tee 5
            local.get 3
            i32.lt_u
            br_if 0 (;@4;)
            br 2 (;@2;)
          end
        end
        local.get 8
        i32.const 1
        i32.lt_s
        br_if 0 (;@2;)
        local.get 9
        local.set 1
        loop  ;; label = @3
          local.get 5
          local.get 1
          i32.load
          i32.store
          local.get 1
          i32.const 4
          i32.add
          local.set 1
          local.get 5
          i32.const 4
          i32.add
          local.tee 5
          local.get 3
          i32.lt_u
          br_if 0 (;@3;)
        end
      end
      local.get 7
      i32.const 3
      i32.and
      local.set 2
      local.get 9
      local.get 8
      i32.add
      local.set 1
    end
    block  ;; label = @1
      local.get 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 3
      local.get 2
      i32.add
      local.set 5
      loop  ;; label = @2
        local.get 3
        local.get 1
        i32.load8_u
        i32.store8
        local.get 1
        i32.const 1
        i32.add
        local.set 1
        local.get 3
        i32.const 1
        i32.add
        local.tee 3
        local.get 5
        i32.lt_u
        br_if 0 (;@2;)
      end
    end
    local.get 0)
  (func $_ZN17compiler_builtins3mem6memset17h2d50fc5881661f2fE (type 1) (param i32 i32 i32) (result i32)
    (local i32 i32 i32)
    block  ;; label = @1
      block  ;; label = @2
        local.get 2
        i32.const 15
        i32.gt_u
        br_if 0 (;@2;)
        local.get 0
        local.set 3
        br 1 (;@1;)
      end
      local.get 0
      i32.const 0
      local.get 0
      i32.sub
      i32.const 3
      i32.and
      local.tee 4
      i32.add
      local.set 5
      block  ;; label = @2
        local.get 4
        i32.eqz
        br_if 0 (;@2;)
        local.get 0
        local.set 3
        loop  ;; label = @3
          local.get 3
          local.get 1
          i32.store8
          local.get 3
          i32.const 1
          i32.add
          local.tee 3
          local.get 5
          i32.lt_u
          br_if 0 (;@3;)
        end
      end
      local.get 5
      local.get 2
      local.get 4
      i32.sub
      local.tee 4
      i32.const -4
      i32.and
      local.tee 2
      i32.add
      local.set 3
      block  ;; label = @2
        local.get 2
        i32.const 1
        i32.lt_s
        br_if 0 (;@2;)
        local.get 1
        i32.const 255
        i32.and
        i32.const 16843009
        i32.mul
        local.set 2
        loop  ;; label = @3
          local.get 5
          local.get 2
          i32.store
          local.get 5
          i32.const 4
          i32.add
          local.tee 5
          local.get 3
          i32.lt_u
          br_if 0 (;@3;)
        end
      end
      local.get 4
      i32.const 3
      i32.and
      local.set 2
    end
    block  ;; label = @1
      local.get 2
      i32.eqz
      br_if 0 (;@1;)
      local.get 3
      local.get 2
      i32.add
      local.set 5
      loop  ;; label = @2
        local.get 3
        local.get 1
        i32.store8
        local.get 3
        i32.const 1
        i32.add
        local.tee 3
        local.get 5
        i32.lt_u
        br_if 0 (;@2;)
      end
    end
    local.get 0)
  (func $memcpy (type 1) (param i32 i32 i32) (result i32)
    local.get 0
    local.get 1
    local.get 2
    call $_ZN17compiler_builtins3mem6memcpy17hcd54104a1e4c5b0dE)
  (table (;0;) 35 35 funcref)
  (memory (;0;) 17)
  (global $__stack_pointer (mut i32) (i32.const 1048576))
  (global (;1;) i32 (i32.const 1056328))
  (global (;2;) i32 (i32.const 1056336))
  (export "memory" (memory 0))
  (export "add" (func $add))
  (export "test_i32_string" (func $test_i32_string))
  (export "test_string_i32" (func $test_string_i32))
  (export "test_string_i32_ret_string" (func $test_string_i32_ret_string))
  (export "test_i64_string_i32" (func $test_i64_string_i32))
  (export "test_string_i64_i32" (func $test_string_i64_i32))
  (export "test_bytes_i64_i32_ret_string" (func $test_bytes_i64_i32_ret_string))
  (export "test_string_string_ret_string" (func $test_string_string_ret_string))
  (export "test_f32_f64_ret_f64" (func $test_f32_f64_ret_f64))
  (export "test_f32_string_f64_ret_string" (func $test_f32_string_f64_ret_string))
  (export "process" (func $process))
  (export "call_winner" (func $call_winner))
  (export "test_string_i64_i32_ret_string" (func $test_bytes_i64_i32_ret_string))
  (export "__data_end" (global 1))
  (export "__heap_base" (global 2))
  (elem (;0;) (i32.const 1) func $_ZN4core3fmt3num3imp52_$LT$impl$u20$core..fmt..Display$u20$for$u20$i32$GT$3fmt17h2a64a79a6a2a885aE $_ZN60_$LT$alloc..string..String$u20$as$u20$core..fmt..Display$GT$3fmt17h14fb461a74ffd756E $_ZN4core3fmt3num3imp52_$LT$impl$u20$core..fmt..Display$u20$for$u20$i64$GT$3fmt17h25d8f33ead29d25bE $_ZN4core3fmt3num3imp52_$LT$impl$u20$core..fmt..Display$u20$for$u20$u32$GT$3fmt17ha297522bf92cf712E $_ZN4core3fmt5float52_$LT$impl$u20$core..fmt..Display$u20$for$u20$f64$GT$3fmt17h33a38e7afb79522fE $_ZN4core3fmt5float52_$LT$impl$u20$core..fmt..Display$u20$for$u20$f32$GT$3fmt17h9d95cc5371839148E $_ZN3std5alloc24default_alloc_error_hook17h8183fc7022a4a0e7E $_ZN4core3ptr100drop_in_place$LT$$RF$mut$u20$std..io..Write..write_fmt..Adapter$LT$alloc..vec..Vec$LT$u8$GT$$GT$$GT$17h30f4c8c66d1c1575E $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_str17h7147b2416db1cf04E $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$10write_char17hd5818ebd4677eb1aE $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_fmt17h2452ec875f334425E $_ZN4core3ptr42drop_in_place$LT$alloc..string..String$GT$17hc02ce641e9e8d8a8E $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17hb75db078a9b00dbdE $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17ha86bcc140bde4c2fE $_ZN4core3ptr70drop_in_place$LT$core..option..Option$LT$alloc..string..String$GT$$GT$17h1c38fec998b374e7E $_ZN90_$LT$std..panicking..begin_panic_handler..PanicPayload$u20$as$u20$core..panic..BoxMeUp$GT$8take_box17h94ddafce7f810e6dE $_ZN90_$LT$std..panicking..begin_panic_handler..PanicPayload$u20$as$u20$core..panic..BoxMeUp$GT$3get17h570cda3222383d8dE $_ZN93_$LT$std..panicking..begin_panic_handler..StrPanicPayload$u20$as$u20$core..panic..BoxMeUp$GT$8take_box17hccf548b0011d7a03E $_ZN93_$LT$std..panicking..begin_panic_handler..StrPanicPayload$u20$as$u20$core..panic..BoxMeUp$GT$3get17h13f86f9299c9908bE $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17h38941f43dd32bc22E $_ZN4core3ptr26drop_in_place$LT$usize$GT$17h052ff139fb0757f4E $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_str17hd86cd7fd76c297c4E $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$10write_char17h0f937bd2ea63dfa4E $_ZN50_$LT$$RF$mut$u20$W$u20$as$u20$core..fmt..Write$GT$9write_fmt17h2389dd1df7b5adb2E $_ZN53_$LT$core..fmt..Error$u20$as$u20$core..fmt..Debug$GT$3fmt17hca756aa95c0138eeE $_ZN4core3ops8function6FnOnce9call_once17h936625e43bb00ae7E $_ZN42_$LT$$RF$T$u20$as$u20$core..fmt..Debug$GT$3fmt17h3fb2133f2241bdd8E $_ZN44_$LT$$RF$T$u20$as$u20$core..fmt..Display$GT$3fmt17h1d04e84ea8294252E $_ZN59_$LT$core..fmt..Arguments$u20$as$u20$core..fmt..Display$GT$3fmt17hf42b4cccce915088E $_ZN71_$LT$core..ops..range..Range$LT$Idx$GT$$u20$as$u20$core..fmt..Debug$GT$3fmt17h4317476aca2a8a15E $_ZN41_$LT$char$u20$as$u20$core..fmt..Debug$GT$3fmt17h19265808cb3f0ce9E $_ZN4core3ptr102drop_in_place$LT$$RF$core..iter..adapters..copied..Copied$LT$core..slice..iter..Iter$LT$u8$GT$$GT$$GT$17h4bf2435251908228E $_ZN36_$LT$T$u20$as$u20$core..any..Any$GT$7type_id17ha7f33678c3328e3eE $_ZN42_$LT$$RF$T$u20$as$u20$core..fmt..Debug$GT$3fmt17h1b991df98dc9a0fdE)
  (data $.rodata (i32.const 1048576) "-\00\00\00\00\00\10\00\00\00\00\00\00\00\10\00\01\00\00\00\00\00\10\00\00\00\00\00\00\00\10\00\01\00\00\00\00\00\10\00\01\00\00\00\08\00\00\00\04\00\00\00\04\00\00\00\09\00\00\00\0a\00\00\00\0b\00\00\00called `Option::unwrap()` on a `None` valuememory allocation of  bytes failed\00\00\00o\00\10\00\15\00\00\00\84\00\10\00\0d\00\00\00library/std/src/alloc.rs\a4\00\10\00\18\00\00\00U\01\00\00\09\00\00\00library/std/src/panicking.rs\cc\00\10\00\1c\00\00\00@\02\00\00\1e\00\00\00\cc\00\10\00\1c\00\00\00?\02\00\00\1f\00\00\00\0c\00\00\00\0c\00\00\00\04\00\00\00\0d\00\00\00\08\00\00\00\08\00\00\00\04\00\00\00\0e\00\00\00\0f\00\00\00\10\00\00\00\04\00\00\00\10\00\00\00\11\00\00\00\08\00\00\00\08\00\00\00\04\00\00\00\12\00\00\00\13\00\00\00\08\00\00\00\00\00\00\00\01\00\00\00\14\00\00\00\15\00\00\00\04\00\00\00\04\00\00\00\16\00\00\00\17\00\00\00\18\00\00\00library/alloc/src/raw_vec.rscapacity overflow\00\00\00\94\01\10\00\11\00\00\00x\01\10\00\1c\00\00\00\0d\02\00\00\05\00\00\00a formatting trait implementation returned an error\00\15\00\00\00\00\00\00\00\01\00\00\00\19\00\00\00library/alloc/src/fmt.rs\04\02\10\00\18\00\00\00d\02\00\00 \00\00\00infassertion failed: edelta >= 0library/core/src/num/diy_float.rs\00\00\00L\02\10\00!\00\00\00L\00\00\00\09\00\00\00L\02\10\00!\00\00\00N\00\00\00\09\00\00\00\01\00\00\00\0a\00\00\00d\00\00\00\e8\03\00\00\10'\00\00\a0\86\01\00@B\0f\00\80\96\98\00\00\e1\f5\05\00\ca\9a;\02\00\00\00\14\00\00\00\c8\00\00\00\d0\07\00\00 N\00\00@\0d\03\00\80\84\1e\00\00-1\01\00\c2\eb\0b\00\945w\00\00\c1o\f2\86#\00\00\00\00\00\81\ef\ac\85[Am-\ee\04\00\00\00\00\00\00\00\00\00\00\01\1fj\bfd\ed8n\ed\97\a7\da\f4\f9?\e9\03O\18\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\01>\95.\09\99\df\03\fd8\15\0f/\e4t#\ec\f5\cf\d3\08\dc\04\c4\da\b0\cd\bc\19\7f3\a6\03&\1f\e9N\02\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\01|.\98[\87\d3\ber\9f\d9\d8\87/\15\12\c6P\dekpnJ\cf\0f\d8\95\d5nq\b2&\b0f\c6\ad$6\15\1dZ\d3B<\0eT\ffc\c0sU\cc\17\ef\f9e\f2(\bcU\f7\c7\dc\80\dc\edn\f4\ce\ef\dc_\f7S\05\00library/core/src/num/flt2dec/strategy/dragon.rsassertion failed: d.mant > 0\00\b8\03\10\00/\00\00\00u\00\00\00\05\00\00\00assertion failed: d.minus > 0\00\00\00\b8\03\10\00/\00\00\00v\00\00\00\05\00\00\00assertion failed: d.plus > 0\b8\03\10\00/\00\00\00w\00\00\00\05\00\00\00assertion failed: d.mant.checked_add(d.plus).is_some()\00\00\b8\03\10\00/\00\00\00x\00\00\00\05\00\00\00assertion failed: d.mant.checked_sub(d.minus).is_some()\00\b8\03\10\00/\00\00\00y\00\00\00\05\00\00\00assertion failed: buf.len() >= MAX_SIG_DIGITS\00\00\00\b8\03\10\00/\00\00\00z\00\00\00\05\00\00\00\b8\03\10\00/\00\00\00\c1\00\00\00\09\00\00\00\b8\03\10\00/\00\00\00\f9\00\00\00T\00\00\00\b8\03\10\00/\00\00\00\fa\00\00\00\0d\00\00\00\b8\03\10\00/\00\00\00\01\01\00\003\00\00\00\b8\03\10\00/\00\00\00\0a\01\00\00\05\00\00\00\b8\03\10\00/\00\00\00\0b\01\00\00\05\00\00\00\b8\03\10\00/\00\00\00\0c\01\00\00\05\00\00\00\b8\03\10\00/\00\00\00\0d\01\00\00\05\00\00\00\b8\03\10\00/\00\00\00\0e\01\00\00\05\00\00\00\b8\03\10\00/\00\00\00K\01\00\00\1f\00\00\00\b8\03\10\00/\00\00\00e\01\00\00\0d\00\00\00\b8\03\10\00/\00\00\00q\01\00\00$\00\00\00\b8\03\10\00/\00\00\00v\01\00\00T\00\00\00\b8\03\10\00/\00\00\00\83\01\00\003\00\00\00\dfE\1a=\03\cf\1a\e6\c1\fb\cc\fe\00\00\00\00\ca\c6\9a\c7\17\fep\ab\dc\fb\d4\fe\00\00\00\00O\dc\bc\be\fc\b1w\ff\f6\fb\dc\fe\00\00\00\00\0c\d6kA\ef\91V\be\11\fc\e4\fe\00\00\00\00<\fc\7f\90\ad\1f\d0\8d,\fc\ec\fe\00\00\00\00\83\9aU1(\5cQ\d3F\fc\f4\fe\00\00\00\00\b5\c9\a6\ad\8f\acq\9da\fc\fc\fe\00\00\00\00\cb\8b\ee#w\22\9c\ea{\fc\04\ff\00\00\00\00mSx@\91I\cc\ae\96\fc\0c\ff\00\00\00\00W\ce\b6]y\12<\82\b1\fc\14\ff\00\00\00\007V\fbM6\94\10\c2\cb\fc\1c\ff\00\00\00\00O\98H8o\ea\96\90\e6\fc$\ff\00\00\00\00\c7:\82%\cb\85t\d7\00\fd,\ff\00\00\00\00\f4\97\bf\97\cd\cf\86\a0\1b\fd4\ff\00\00\00\00\e5\ac*\17\98\0a4\ef5\fd<\ff\00\00\00\00\8e\b25*\fbg8\b2P\fdD\ff\00\00\00\00;?\c6\d2\df\d4\c8\84k\fdL\ff\00\00\00\00\ba\cd\d3\1a'D\dd\c5\85\fdT\ff\00\00\00\00\96\c9%\bb\ce\9fk\93\a0\fd\5c\ff\00\00\00\00\84\a5b}$l\ac\db\ba\fdd\ff\00\00\00\00\f6\da_\0dXf\ab\a3\d5\fdl\ff\00\00\00\00&\f1\c3\de\93\f8\e2\f3\ef\fdt\ff\00\00\00\00\b8\80\ff\aa\a8\ad\b5\b5\0a\fe|\ff\00\00\00\00\8bJ|l\05_b\87%\fe\84\ff\00\00\00\00S0\c14`\ff\bc\c9?\fe\8c\ff\00\00\00\00U&\ba\91\8c\85N\96Z\fe\94\ff\00\00\00\00\bd~)p$w\f9\dft\fe\9c\ff\00\00\00\00\8f\b8\e5\b8\9f\bd\df\a6\8f\fe\a4\ff\00\00\00\00\94}t\88\cf_\a9\f8\a9\fe\ac\ff\00\00\00\00\cf\9b\a8\8f\93pD\b9\c4\fe\b4\ff\00\00\00\00k\15\0f\bf\f8\f0\08\8a\df\fe\bc\ff\00\00\00\00\b611eU%\b0\cd\f9\fe\c4\ff\00\00\00\00\ac\7f{\d0\c6\e2?\99\14\ff\cc\ff\00\00\00\00\06;+*\c4\10\5c\e4.\ff\d4\ff\00\00\00\00\d3\92si\99$$\aaI\ff\dc\ff\00\00\00\00\0e\ca\00\83\f2\b5\87\fdc\ff\e4\ff\00\00\00\00\eb\1a\11\92d\08\e5\bc~\ff\ec\ff\00\00\00\00\cc\88Po\09\cc\bc\8c\99\ff\f4\ff\00\00\00\00,e\19\e2X\17\b7\d1\b3\ff\fc\ff\00\00\00\00\00\00\00\00\00\00@\9c\ce\ff\04\00\00\00\00\00\00\00\00\00\10\a5\d4\e8\e8\ff\0c\00\00\00\00\00\00\00b\ac\c5\ebx\ad\03\00\14\00\00\00\00\00\84\09\94\f8x9?\81\1e\00\1c\00\00\00\00\00\b3\15\07\c9{\ce\97\c08\00$\00\00\00\00\00p\5c\ea{\ce2~\8fS\00,\00\00\00\00\00h\80\e9\ab\a48\d2\d5m\004\00\00\00\00\00E\22\9a\17&'O\9f\88\00<\00\00\00\00\00'\fb\c4\d41\a2c\ed\a2\00D\00\00\00\00\00\a8\ad\c8\8c8e\de\b0\bd\00L\00\00\00\00\00\dbe\ab\1a\8e\08\c7\83\d8\00T\00\00\00\00\00\9a\1dqB\f9\1d]\c4\f2\00\5c\00\00\00\00\00X\e7\1b\a6,iM\92\0d\01d\00\00\00\00\00\ea\8dp\1ad\ee\01\da'\01l\00\00\00\00\00Jw\ef\9a\99\a3m\a2B\01t\00\00\00\00\00\85k}\b4{x\09\f2\5c\01|\00\00\00\00\00w\18\ddy\a1\e4T\b4w\01\84\00\00\00\00\00\c2\c5\9b[\92\86[\86\92\01\8c\00\00\00\00\00=]\96\c8\c5S5\c8\ac\01\94\00\00\00\00\00\b3\a0\97\fa\5c\b4*\95\c7\01\9c\00\00\00\00\00\e3_\a0\99\bd\9fF\de\e1\01\a4\00\00\00\00\00%\8c9\db4\c2\9b\a5\fc\01\ac\00\00\00\00\00\5c\9f\98\a3r\9a\c6\f6\16\02\b4\00\00\00\00\00\ce\be\e9TS\bf\dc\b71\02\bc\00\00\00\00\00\e2A\22\f2\17\f3\fc\88L\02\c4\00\00\00\00\00\a5x\5c\d3\9b\ce \ccf\02\cc\00\00\00\00\00\dfS!{\f3Z\16\98\81\02\d4\00\00\00\00\00:0\1f\97\dc\b5\a0\e2\9b\02\dc\00\00\00\00\00\96\b3\e3\5cS\d1\d9\a8\b6\02\e4\00\00\00\00\00<D\a7\a4\d9|\9b\fb\d0\02\ec\00\00\00\00\00\10D\a4\a7LLv\bb\eb\02\f4\00\00\00\00\00\1a\9c@\b6\ef\8e\ab\8b\06\03\fc\00\00\00\00\00,\84W\a6\10\ef\1f\d0 \03\04\01\00\00\00\00)1\91\e9\e5\a4\10\9b;\03\0c\01\00\00\00\00\9d\0c\9c\a1\fb\9b\10\e7U\03\14\01\00\00\00\00)\f4;b\d9 (\acp\03\1c\01\00\00\00\00\85\cf\a7z^KD\80\8b\03$\01\00\00\00\00-\dd\ac\03@\e4!\bf\a5\03,\01\00\00\00\00\8f\ffD^/\9cg\8e\c0\034\01\00\00\00\00A\b8\8c\9c\9d\173\d4\da\03<\01\00\00\00\00\a9\1b\e3\b4\92\db\19\9e\f5\03D\01\00\00\00\00\d9w\df\ban\bf\96\eb\0f\04L\01\00\00\00\00library/core/src/num/flt2dec/strategy/grisu.rs\00\000\0b\10\00.\00\00\00}\00\00\00\15\00\00\000\0b\10\00.\00\00\00\a9\00\00\00\05\00\00\000\0b\10\00.\00\00\00\aa\00\00\00\05\00\00\000\0b\10\00.\00\00\00\ab\00\00\00\05\00\00\000\0b\10\00.\00\00\00\ac\00\00\00\05\00\00\000\0b\10\00.\00\00\00\ad\00\00\00\05\00\00\000\0b\10\00.\00\00\00\ae\00\00\00\05\00\00\00assertion failed: d.mant + d.plus < (1 << 61)\00\00\000\0b\10\00.\00\00\00\af\00\00\00\05\00\00\000\0b\10\00.\00\00\00\0a\01\00\00\11\00\00\00attempt to divide by zero\00\00\000\0b\10\00.\00\00\00\0d\01\00\00\09\00\00\000\0b\10\00.\00\00\00\16\01\00\00B\00\00\000\0b\10\00.\00\00\00@\01\00\00\09\00\00\00assertion failed: !buf.is_empty()called `Option::unwrap()` on a `None` value0\0b\10\00.\00\00\00\dc\01\00\00\05\00\00\00assertion failed: d.mant < (1 << 61)0\0b\10\00.\00\00\00\dd\01\00\00\05\00\00\000\0b\10\00.\00\00\00\de\01\00\00\05\00\00\000\0b\10\00.\00\00\00#\02\00\00\11\00\00\000\0b\10\00.\00\00\00&\02\00\00\09\00\00\000\0b\10\00.\00\00\00\5c\02\00\00\09\00\00\000\0b\10\00.\00\00\00\bc\02\00\00G\00\00\000\0b\10\00.\00\00\00\d3\02\00\00K\00\00\000\0b\10\00.\00\00\00\df\02\00\00G\00\00\00library/core/src/num/flt2dec/mod.rs\00l\0d\10\00#\00\00\00\bc\00\00\00\05\00\00\00assertion failed: buf[0] > b\5c'0\5c'\00\00\00l\0d\10\00#\00\00\00\bd\00\00\00\05\00\00\00assertion failed: parts.len() >= 4\00\00l\0d\10\00#\00\00\00\be\00\00\00\05\00\00\000..-+\00\00\000NaNassertion failed: buf.len() >= maxlen\00\00\00l\0d\10\00#\00\00\00\7f\02\00\00\0d\00\00\00..\00\00L\0e\10\00\02\00\00\00 \00\00\00\00\00\00\00\01\00\00\00!\00\00\00index out of bounds: the len is  but the index is \00\00h\0e\10\00 \00\00\00\88\0e\10\00\12\00\00\00 \00\00\00\04\00\00\00\04\00\00\00\22\00\00\00matches!===assertion failed: `(left  right)`\0a  left: ``,\0a right: ``: \00\00\00\c7\0e\10\00\19\00\00\00\e0\0e\10\00\12\00\00\00\f2\0e\10\00\0c\00\00\00\fe\0e\10\00\03\00\00\00`\00\00\00\c7\0e\10\00\19\00\00\00\e0\0e\10\00\12\00\00\00\f2\0e\10\00\0c\00\00\00$\0f\10\00\01\00\00\00: \00\00\10\0e\10\00\00\00\00\00H\0f\10\00\02\00\00\00library/core/src/fmt/num.rs\00\5c\0f\10\00\1b\00\00\00e\00\00\00\14\00\00\000x00010203040506070809101112131415161718192021222324252627282930313233343536373839404142434445464748495051525354555657585960616263646566676869707172737475767778798081828384858687888990919293949596979899library/core/src/fmt/mod.rs\00\00\00R\10\10\00\1b\00\00\00p\06\00\00\1e\00\00\000000000000000000000000000000000000000000000000000000000000000000R\10\10\00\1b\00\00\00j\06\00\00-\00\00\00range start index  out of range for slice of length \d0\10\10\00\12\00\00\00\e2\10\10\00\22\00\00\00range end index \14\11\10\00\10\00\00\00\e2\10\10\00\22\00\00\00slice index starts at  but ends at \004\11\10\00\16\00\00\00J\11\10\00\0d\00\00\00[...]byte index  is out of bounds of `\00\00m\11\10\00\0b\00\00\00x\11\10\00\16\00\00\00$\0f\10\00\01\00\00\00begin <= end ( <= ) when slicing `\00\00\a8\11\10\00\0e\00\00\00\b6\11\10\00\04\00\00\00\ba\11\10\00\10\00\00\00$\0f\10\00\01\00\00\00 is not a char boundary; it is inside  (bytes ) of `m\11\10\00\0b\00\00\00\ec\11\10\00&\00\00\00\12\12\10\00\08\00\00\00\1a\12\10\00\06\00\00\00$\0f\10\00\01\00\00\00library/core/src/str/mod.rs\00H\12\10\00\1b\00\00\00\07\01\00\00\1d\00\00\00library/core/src/unicode/printable.rs\00\00\00t\12\10\00%\00\00\00\0a\00\00\00\1c\00\00\00t\12\10\00%\00\00\00\1a\00\00\006\00\00\00\00\01\03\05\05\06\06\02\07\06\08\07\09\11\0a\1c\0b\19\0c\1a\0d\10\0e\0c\0f\04\10\03\12\12\13\09\16\01\17\04\18\01\19\03\1a\07\1b\01\1c\02\1f\16 \03+\03-\0b.\010\031\022\01\a7\02\a9\02\aa\04\ab\08\fa\02\fb\05\fd\02\fe\03\ff\09\adxy\8b\8d\a20WX\8b\8c\90\1c\dd\0e\0fKL\fb\fc./?\5c]_\e2\84\8d\8e\91\92\a9\b1\ba\bb\c5\c6\c9\ca\de\e4\e5\ff\00\04\11\12)147:;=IJ]\84\8e\92\a9\b1\b4\ba\bb\c6\ca\ce\cf\e4\e5\00\04\0d\0e\11\12)14:;EFIJ^de\84\91\9b\9d\c9\ce\cf\0d\11):;EIW[\5c^_de\8d\91\a9\b4\ba\bb\c5\c9\df\e4\e5\f0\0d\11EIde\80\84\b2\bc\be\bf\d5\d7\f0\f1\83\85\8b\a4\a6\be\bf\c5\c7\cf\da\dbH\98\bd\cd\c6\ce\cfINOWY^_\89\8e\8f\b1\b6\b7\bf\c1\c6\c7\d7\11\16\17[\5c\f6\f7\fe\ff\80mq\de\df\0e\1fno\1c\1d_}~\ae\af\7f\bb\bc\16\17\1e\1fFGNOXZ\5c^~\7f\b5\c5\d4\d5\dc\f0\f1\f5rs\8ftu\96&./\a7\af\b7\bf\c7\cf\d7\df\9a@\97\980\8f\1f\d2\d4\ce\ffNOZ[\07\08\0f\10'/\ee\efno7=?BE\90\91Sgu\c8\c9\d0\d1\d8\d9\e7\fe\ff\00 _\22\82\df\04\82D\08\1b\04\06\11\81\ac\0e\80\ab\05\1f\09\81\1b\03\19\08\01\04/\044\04\07\03\01\07\06\07\11\0aP\0f\12\07U\07\03\04\1c\0a\09\03\08\03\07\03\02\03\03\03\0c\04\05\03\0b\06\01\0e\15\05N\07\1b\07W\07\02\06\17\0cP\04C\03-\03\01\04\11\06\0f\0c:\04\1d%_ m\04j%\80\c8\05\82\b0\03\1a\06\82\fd\03Y\07\16\09\18\09\14\0c\14\0cj\06\0a\06\1a\06Y\07+\05F\0a,\04\0c\04\01\031\0b,\04\1a\06\0b\03\80\ac\06\0a\06/1M\03\80\a4\08<\03\0f\03<\078\08+\05\82\ff\11\18\08/\11-\03!\0f!\0f\80\8c\04\82\97\19\0b\15\88\94\05/\05;\07\02\0e\18\09\80\be\22t\0c\80\d6\1a\0c\05\80\ff\05\80\df\0c\f2\9d\037\09\81\5c\14\80\b8\08\80\cb\05\0a\18;\03\0a\068\08F\08\0c\06t\0b\1e\03Z\04Y\09\80\83\18\1c\0a\16\09L\04\80\8a\06\ab\a4\0c\17\041\a1\04\81\da&\07\0c\05\05\80\a6\10\81\f5\07\01 *\06L\04\80\8d\04\80\be\03\1b\03\0f\0d\00\06\01\01\03\01\04\02\05\07\07\02\08\08\09\02\0a\05\0b\02\0e\04\10\01\11\02\12\05\13\11\14\01\15\02\17\02\19\0d\1c\05\1d\08\1f\01$\01j\04k\02\af\03\b1\02\bc\02\cf\02\d1\02\d4\0c\d5\09\d6\02\d7\02\da\01\e0\05\e1\02\e7\04\e8\02\ee \f0\04\f8\02\fa\03\fb\01\0c';>NO\8f\9e\9e\9f{\8b\93\96\a2\b2\ba\86\b1\06\07\096=>V\f3\d0\d1\04\14\1867VW\7f\aa\ae\af\bd5\e0\12\87\89\8e\9e\04\0d\0e\11\12)14:EFIJNOde\5c\b6\b7\1b\1c\07\08\0a\0b\14\1769:\a8\a9\d8\d9\097\90\91\a8\07\0a;>fi\8f\92\11o_\bf\ee\efZb\f4\fc\ffST\9a\9b./'(U\9d\a0\a1\a3\a4\a7\a8\ad\ba\bc\c4\06\0b\0c\15\1d:?EQ\a6\a7\cc\cd\a0\07\19\1a\22%>?\e7\ec\ef\ff\c5\c6\04 #%&(38:HJLPSUVXZ\5c^`cefksx}\7f\8a\a4\aa\af\b0\c0\d0\ae\afno\be\93^\22{\05\03\04-\03f\03\01/.\80\82\1d\031\0f\1c\04$\09\1e\05+\05D\04\0e*\80\aa\06$\04$\04(\084\0bNC\817\09\16\0a\08\18;E9\03c\08\090\16\05!\03\1b\05\01@8\04K\05/\04\0a\07\09\07@ '\04\0c\096\03:\05\1a\07\04\0c\07PI73\0d3\07.\08\0a\81&RK+\08*\16\1a&\1c\14\17\09N\04$\09D\0d\19\07\0a\06H\08'\09u\0bB>*\06;\05\0a\06Q\06\01\05\10\03\05\80\8bb\1eH\08\0a\80\a6^\22E\0b\0a\06\0d\13:\06\0a6,\04\17\80\b9<dS\0cH\09\0aFE\1bH\08S\0dI\07\0a\80\f6F\0a\1d\03GI7\03\0e\08\0a\069\07\0a\816\19\07;\03\1cV\01\0f2\0d\83\9bfu\0b\80\c4\8aLc\0d\840\10\16\8f\aa\82G\a1\b9\829\07*\04\5c\06&\0aF\0a(\05\13\82\b0[eK\049\07\11@\05\0b\02\0e\97\f8\08\84\d6*\09\a2\e7\813\0f\01\1d\06\0e\04\08\81\8c\89\04k\05\0d\03\09\07\10\92`G\09t<\80\f6\0as\08p\15Fz\14\0c\14\0cW\09\19\80\87\81G\03\85B\0f\15\84P\1f\06\06\80\d5+\05>!\01p-\03\1a\04\02\81@\1f\11:\05\01\81\d0*\82\e6\80\f7)L\04\0a\04\02\83\11DL=\80\c2<\06\01\04U\05\1b4\02\81\0e,\04d\0cV\0a\80\ae8\1d\0d,\04\09\07\02\0e\06\80\9a\83\d8\04\11\03\0d\03w\04_\06\0c\04\01\0f\0c\048\08\0a\06(\08\22N\81T\0c\1d\03\09\076\08\0e\04\09\07\09\07\80\cb%\0a\84\06library/core/src/unicode/unicode_data.rslibrary/core/src/num/bignum.rs\00\00`\18\10\00\1e\00\00\00\ac\01\00\00\01\00\00\00assertion failed: noborrowassertion failed: digits < 40assertion failed: other > 0Error\008\18\10\00(\00\00\00P\00\00\00(\00\00\008\18\10\00(\00\00\00\5c\00\00\00\16\00\00\00\00\03\00\00\83\04 \00\91\05`\00]\13\a0\00\12\17 \1f\0c `\1f\ef,\a0+*0 ,o\a6\e0,\02\a8`-\1e\fb`.\00\fe 6\9e\ff`6\fd\01\e16\01\0a!7$\0d\e17\ab\0ea9/\18\a190\1caH\f3\1e\a1L@4aP\f0j\a1QOo!R\9d\bc\a1R\00\cfaSe\d1\a1S\00\da!T\00\e0\e1U\ae\e2aW\ec\e4!Y\d0\e8\a1Y \00\eeY\f0\01\7fZ\00p\00\07\00-\01\01\01\02\01\02\01\01H\0b0\15\10\01e\07\02\06\02\02\01\04#\01\1e\1b[\0b:\09\09\01\18\04\01\09\01\03\01\05+\03<\08*\18\01 7\01\01\01\04\08\04\01\03\07\0a\02\1d\01:\01\01\01\02\04\08\01\09\01\0a\02\1a\01\02\029\01\04\02\04\02\02\03\03\01\1e\02\03\01\0b\029\01\04\05\01\02\04\01\14\02\16\06\01\01:\01\01\02\01\04\08\01\07\03\0a\02\1e\01;\01\01\01\0c\01\09\01(\01\03\017\01\01\03\05\03\01\04\07\02\0b\02\1d\01:\01\02\01\02\01\03\01\05\02\07\02\0b\02\1c\029\02\01\01\02\04\08\01\09\01\0a\02\1d\01H\01\04\01\02\03\01\01\08\01Q\01\02\07\0c\08b\01\02\09\0b\07I\02\1b\01\01\01\01\017\0e\01\05\01\02\05\0b\01$\09\01f\04\01\06\01\02\02\02\19\02\04\03\10\04\0d\01\02\02\06\01\0f\01\00\03\00\03\1d\02\1e\02\1e\02@\02\01\07\08\01\02\0b\09\01-\03\01\01u\02\22\01v\03\04\02\09\01\06\03\db\02\02\01:\01\01\07\01\01\01\01\02\08\06\0a\02\010\1f1\040\07\01\01\05\01(\09\0c\02 \04\02\02\01\038\01\01\02\03\01\01\03:\08\02\02\98\03\01\0d\01\07\04\01\06\01\03\02\c6@\00\01\c3!\00\03\8d\01` \00\06i\02\00\04\01\0a \02P\02\00\01\03\01\04\01\19\02\05\01\97\02\1a\12\0d\01&\08\19\0b.\030\01\02\04\02\02'\01C\06\02\02\02\02\0c\01\08\01/\013\01\01\03\02\02\05\02\01\01*\02\08\01\ee\01\02\01\04\01\00\01\00\10\10\10\00\02\00\01\e2\01\95\05\00\03\01\02\05\04(\03\04\01\a5\02\00\04\00\02P\03F\0b1\04{\016\0f)\01\02\02\0a\031\04\02\02\07\01=\03$\05\01\08>\01\0c\024\09\0a\04\02\01_\03\02\01\01\02\06\01\02\01\9d\01\03\08\15\029\02\01\01\01\01\16\01\0e\07\03\05\c3\08\02\03\01\01\17\01Q\01\02\06\01\01\02\01\01\02\01\02\eb\01\02\04\06\02\01\02\1b\02U\08\02\01\01\02j\01\01\01\02\06\01\01e\03\02\04\01\05\00\09\01\02\f5\01\0a\02\01\01\04\01\90\04\02\02\04\01 \0a(\06\02\04\08\01\09\06\02\03.\0d\01\02\00\07\01\06\01\01R\16\02\07\01\02\01\02z\06\03\01\01\02\01\07\01\01H\02\03\01\01\01\00\02\0b\024\05\05\01\01\01\00\01\06\0f\00\05;\07\00\01?\04Q\01\00\02\00.\02\17\00\01\01\03\04\05\08\08\02\07\1e\04\94\03\007\042\08\01\0e\01\16\05\01\0f\00\07\01\11\02\07\01\02\01\05d\01\a0\07\00\01=\04\00\04\00\07m\07\00`\80\f0\00"))
