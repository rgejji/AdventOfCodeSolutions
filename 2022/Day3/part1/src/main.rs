fn sum_priorities(input: &str) -> usize{
    let split = input.split("\n");
    let mut total:usize = 0;
    for s in split {
        total = total + score_line_priorities(s);
    }
    total
}

fn score_line_priorities(s: &str) -> usize {
    const ASCII_A:usize = 65; 
    const ASCII_a:usize = 97; 
    let mut total = 0;
    let mut inventory = [0; 52];
    let bytes = s.as_bytes();
    let n = bytes.len();
    //println!("Length is {}", n);
    for (i,b) in bytes.iter().enumerate() {
        let val = byte_to_int(&b);
        let val = if val >= ASCII_a { val - ASCII_a } else { val + 26 - ASCII_A };

        if i < n/2 {
            //println!("inventorying {}", i);
            inventory[val] = inventory[val] + 1;
        } else {
            //println!("checking {}", i);
            if inventory[val] > 0 {
                println!("Have priority {}", val+1);
                total += val +1;
                break;
            }
        }
    }
    //println!("Done");
    total
    
}

fn byte_to_int(c:&u8) -> usize {
    let result:u32 = (*c).into();
    usize::try_from(result).unwrap()
}

fn main() {
    let score = sum_priorities(INPUT);
    println!("Score: {}", score);
}

const INPUT:&str = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";
