use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::cmp; 

// based on https://doc.rust-lang.org/stable/rust-by-example/std_misc/file/read_lines.html
fn main() {
    let mut max: i32 = 0;
    let mut sum: i32 = 0;
    if let Ok(lines) = read_lines("input") {
        for line in lines {
            if let Ok(data) = line {
                if data == "" {
                    max = cmp::max(max, sum);
                    sum = 0;
                    continue
                }
                sum += data.parse::<i32>().unwrap();
            }
        }
    }
    
    max = cmp::max(max, sum);
    println!("{}", max);
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
