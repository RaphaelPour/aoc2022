use std::fs::File;
use std::io::{self, prelude::*, BufReader};
use std::cmp; 

// https://stackoverflow.com/a/45882510
fn main() -> io::Result<()> {
    let mut max: i32 = 0;
    let mut sum: i32 = 0;

    let file = match File::open("input"){
        Err(e) => return Err(e),
        Ok(file) => file,
    };
    let reader = BufReader::new(file);
    for line in reader.lines() {
        if let Ok(data) = line {
            if data == "" {
                max = cmp::max(max, sum);
                sum = 0;
                continue
            }
            // https://doc.rust-lang.org/std/primitive.str.html#method.parse
            sum += data.parse::<i32>().unwrap();
        }
    }
    
    max = cmp::max(max, sum);
    println!("{}", max);

    Ok(())
}
