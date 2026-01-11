# FFI

生成的是静态依赖

.o -> .a

.so

## C

## C++

+ 手动FFI 
+ bindgen
+ cxx

## 实践

CARGO_MANIFEST_DIR和OUT_DIR

### 编译

#### 直接利用原来的项目进编译

```rust
fn main() {
    let manifest_dir = PathBuf::from(env::var("CARGO_MANIFEST_DIR").unwrap());
    let make_dir = manifest_dir.join("c-project");

    let status = Command::new("make")
        .current_dir(&make_dir)
        .arg("all")
        .status()
        .expect("Failed to execute make");

    if !status.success() {
        panic!(
            "Make failed with exit status: {}",
            status.code().unwrap_or(-1)
        );
    }
    println!("cargo:rerun-if-changed=vanitygen-plusplus/Makefile");
}
```

#### rust内置

```rust
fn main() {
    let manifest_dir = PathBuf::from(env::var("CARGO_MANIFEST_DIR").unwrap());
    let out_dir = manifest_dir.join(env::var("OUT_DIR").unwrap());
    let mut build = cc::Build::new();

    build.files(&[
        "demo.c",
    ]);

    build
        .flag("-ggdb")
        .flag("-Wall")
        .flag("-Wno-deprecated")
        .compile("demo");

    
    // println!("cargo:rustc-link-search=native={}", out_dir.display());
    // println!("cargo:rustc-link-lib=pcre");

    println!("cargo:rerun-if-changed=demo.c");
}
```

### bingen

```rust
// 调用的文件中引入
// #![allow(non_upper_case_globals)]
// #![allow(non_camel_case_types)]
// #![allow(non_snake_case)]

// include!(concat!(env!("OUT_DIR"), "/bindings.rs"));
fn main() {
    let manifest_dir = PathBuf::from(env::var("CARGO_MANIFEST_DIR").unwrap());
    let out_dir = manifest_dir.join(env::var("OUT_DIR").unwrap());
    let bindings = bindgen::Builder::default()
        .header("demo.h")
        .parse_callbacks(Box::new(bindgen::CargoCallbacks::new()))
        .generate()
        .expect("Unable to generate bindings");

    bindings
        .write_to_file(out_dir.join("bindings.rs"))
        .expect("Couldn't write bindings!");
}
```
