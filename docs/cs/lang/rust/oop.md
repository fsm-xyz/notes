# 在 Rust 中实现 Java 风格的面向对象特性

Rust 虽然不是传统的面向对象语言，但它提供了许多特性来支持面向对象编程模式。下面是一个示例，展示如何在 Rust 中实现类似 Java 的面向对象特性。

```rust
use std::fmt::Display;

// 定义一个 trait (类似 Java 接口)
trait Animal: Display {
    fn make_sound(&self) -> String;
    fn move_(&self) -> String {
        String::from("moving")
    }
}

// 基类/父类 (使用 struct 和 impl)
struct Mammal {
    name: String,
    age: u8,
}

impl Mammal {
    fn new(name: String, age: u8) -> Self {
        Mammal { name, age }
    }
    
    // 公共方法
    fn get_name(&self) -> &str {
        &self.name
    }
    
    fn get_age(&self) -> u8 {
        self.age
    }
    
    // 受保护的方法 (通过模块系统模拟)
    pub(crate) fn protected_method(&self) {
        println!("This is a protected method");
    }
    
    // 私有方法 (默认就是私有的)
    fn private_method(&self) {
        println!("This is private to Mammal");
    }
}

// 派生类/子类
struct Dog {
    mammal: Mammal,  // 使用组合而不是继承
    breed: String,
}

impl Dog {
    fn new(name: String, age: u8, breed: String) -> Self {
        Dog {
            mammal: Mammal::new(name, age),
            breed,
        }
    }
    
    // 方法重载
    fn fetch(&self, item: &str) -> String {
        format!("{} fetched the {}", self.mammal.get_name(), item)
    }
    
    // 访问"父类"方法
    fn get_name(&self) -> &str {
        self.mammal.get_name()
    }
    
    fn get_age(&self) -> u8 {
        self.mammal.get_age()
    }
}

// 实现 Animal trait (类似实现接口)
impl Animal for Dog {
    fn make_sound(&self) -> String {
        String::from("Woof!")
    }
    
    // 重写默认实现
    fn move_(&self) -> String {
        String::from("running")
    }
}

// 实现 Display trait (类似 Java 的 toString())
impl Display for Dog {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "Dog: {} ({} years old, breed: {})", 
               self.mammal.get_name(), self.mammal.get_age(), self.breed)
    }
}

// 多态示例函数
fn animal_info(animal: &impl Animal) {
    println!("{}", animal);
    println!("Sound: {}", animal.make_sound());
    println!("Movement: {}", animal.move_());
}

// 另一个实现 Animal 的结构体
struct Cat {
    mammal: Mammal,
    color: String,
}

impl Cat {
    fn new(name: String, age: u8, color: String) -> Self {
        Cat {
            mammal: Mammal::new(name, age),
            color,
        }
    }
}

impl Animal for Cat {
    fn make_sound(&self) -> String {
        String::from("Meow!")
    }
}

impl Display for Cat {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "Cat: {} ({} years old, color: {})", 
               self.mammal.get_name(), self.mammal.get_age(), self.color)
    }
}

fn main() {
    println!("=== Java-like OOP in Rust ===");
    
    // 创建对象实例
    let dog = Dog::new(String::from("Buddy"), 5, String::from("Golden Retriever"));
    let cat = Cat::new(String::from("Whiskers"), 3, String::from("Tabby"));
    
    // 调用方法
    println!("Dog name: {}", dog.get_name());
    println!("Dog age: {}", dog.get_age());
    println!("Dog sound: {}", dog.make_sound());
    println!("Dog movement: {}", dog.move_());
    println!("Dog fetch: {}", dog.fetch("ball"));
    
    println!();
    
    // 使用多态
    println!("--- Polymorphism Example ---");
    animal_info(&dog);
    println!();
    animal_info(&cat);
    
    println!();
    
    // 直接使用 Display trait
    println!("--- Display Implementation ---");
    println!("{}", dog);
    println!("{}", cat);
}
```

## 运行结果

```
=== Java-like OOP in Rust ===
Dog name: Buddy
Dog age: 5
Dog sound: Woof!
Dog movement: running
Dog fetch: Buddy fetched the ball

--- Polymorphism Example ---
Dog: Buddy (5 years old, breed: Golden Retriever)
Sound: Woof!
Movement: running

Cat: Whiskers (3 years old, color: Tabby)
Sound: Meow!
Movement: moving

--- Display Implementation ---
Dog: Buddy (5 years old, breed: Golden Retriever)
Cat: Whiskers (3 years old, color: Tabby)
```

## Rust 中实现面向对象的关键点

1. **封装**: 使用 `pub` 关键字控制可见性，默认所有项都是私有的
2. **继承的替代方案**:
   - 使用 trait 实现多态和行为共享
   - 使用组合而不是继承来重用代码
3. **多态**: 通过 trait 对象实现运行时多态
4. **构造函数**: 惯例使用 `new` 关联函数创建实例
5. **方法**: 在 `impl` 块中定义方法，第一个参数通常是 `&self` 或 `&mut self`

虽然 Rust 不支持传统的类继承，但通过 trait 和组合，它可以实现更灵活和安全的面向对象设计模式。