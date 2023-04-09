/* tslint:disable */
/* eslint-disable */
/**
* @param {number} a
* @param {number} b
* @returns {number}
*/
export function add(a: number, b: number): number;
/**
* @param {number} a
* @param {string} b
* @returns {number}
*/
export function test_i32_string(a: number, b: string): number;
/**
* @param {string} a
* @param {number} b
* @returns {number}
*/
export function test_string_i32(a: string, b: number): number;
/**
* @param {string} a
* @param {number} b
* @returns {string}
*/
export function test_string_i32_ret_string(a: string, b: number): string;
/**
* @param {bigint} a
* @param {string} b
* @param {number} c
* @returns {bigint}
*/
export function test_i64_string_i32(a: bigint, b: string, c: number): bigint;
/**
* @param {string} a
* @param {bigint} b
* @param {number} c
* @returns {number}
*/
export function test_string_i64_i32(a: string, b: bigint, c: number): number;
/**
* @param {string} a
* @param {bigint} b
* @param {number} c
* @returns {string}
*/
export function test_string_i64_i32_ret_string(a: string, b: bigint, c: number): string;
/**
* @param {number} a
* @param {Uint8Array} b
* @param {number} c
* @returns {string}
*/
export function test_bytes_i64_i32_ret_string(a: number, b: Uint8Array, c: number): string;
/**
* @param {string} a
* @param {string} b
* @returns {string}
*/
export function test_string_string_ret_string(a: string, b: string): string;
/**
* @param {number} a
* @param {number} b
* @returns {number}
*/
export function test_f32_f64_ret_f64(a: number, b: number): number;
/**
* @param {number} a
* @param {string} b
* @param {number} c
* @returns {string}
*/
export function test_f32_string_f64_ret_string(a: number, b: string, c: number): string;
/**
* @param {string} s
* @returns {string}
*/
export function process(s: string): string;
/**
* @param {string} input
* @returns {number}
*/
export function call_winner(input: string): number;

export type InitInput = RequestInfo | URL | Response | BufferSource | WebAssembly.Module;

export interface InitOutput {
  readonly memory: WebAssembly.Memory;
  readonly add: (a: number, b: number) => number;
  readonly test_i32_string: (a: number, b: number, c: number) => number;
  readonly test_string_i32: (a: number, b: number, c: number) => number;
  readonly test_string_i32_ret_string: (a: number, b: number, c: number, d: number) => void;
  readonly test_i64_string_i32: (a: number, b: number, c: number, d: number) => number;
  readonly test_string_i64_i32: (a: number, b: number, c: number, d: number) => number;
  readonly test_string_i64_i32_ret_string: (a: number, b: number, c: number, d: number, e: number) => void;
  readonly test_bytes_i64_i32_ret_string: (a: number, b: number, c: number, d: number, e: number) => void;
  readonly test_string_string_ret_string: (a: number, b: number, c: number, d: number, e: number) => void;
  readonly test_f32_f64_ret_f64: (a: number, b: number) => number;
  readonly test_f32_string_f64_ret_string: (a: number, b: number, c: number, d: number, e: number) => void;
  readonly process: (a: number, b: number, c: number) => void;
  readonly call_winner: (a: number, b: number) => number;
  readonly __wbindgen_malloc: (a: number) => number;
  readonly __wbindgen_realloc: (a: number, b: number, c: number) => number;
  readonly __wbindgen_add_to_stack_pointer: (a: number) => number;
  readonly __wbindgen_free: (a: number, b: number) => void;
}

export type SyncInitInput = BufferSource | WebAssembly.Module;
/**
* Instantiates the given `module`, which can either be bytes or
* a precompiled `WebAssembly.Module`.
*
* @param {SyncInitInput} module
*
* @returns {InitOutput}
*/
export function initSync(module: SyncInitInput): InitOutput;

/**
* If `module_or_path` is {RequestInfo} or {URL}, makes a request and
* for everything else, calls `WebAssembly.instantiate` directly.
*
* @param {InitInput | Promise<InitInput>} module_or_path
*
* @returns {Promise<InitOutput>}
*/
export default function init (module_or_path?: InitInput | Promise<InitInput>): Promise<InitOutput>;
