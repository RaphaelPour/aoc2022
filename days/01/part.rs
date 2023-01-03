use std::fs::File;
use std::io::{self, prelude::*, BufReader};

// https://stackoverflow.com/a/45882510
fn main() -> io::Result<()> {
    // https://doc.rust-lang.org/std/vec/struct.Vec.html
    let mut calories = Vec::new();
    let mut sum: i32 = 0;

    // https://doc.rust-lang.org/book/ch09-02-recoverable-errors-with-result.html#propagating-errors
    let file = match File::open("input"){
        Err(e) => return Err(e),
        Ok(file) => file,
    };
    for line in  BufReader::new(file).lines() {
        let data = match line {
            Err(e) => return Err(e),
            Ok(data) => data,
        };

        if data == "" {
            calories.push(sum);
            sum = 0;
            continue
        }
        // https://doc.rust-lang.org/std/primitive.str.html#method.parse
        sum += data.parse::<i32>().unwrap();
    }
    
    calories.push(sum);
    calories.sort();
    
    println!("part1: {}", calories.iter().rev().take(1).sum::<i32>());
    println!("part2: {}", calories.iter().rev().take(3).sum::<i32>());
    
    Ok(())
}
