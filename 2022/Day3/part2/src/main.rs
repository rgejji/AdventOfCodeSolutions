fn sum_priorities(input: &str) -> usize{
    let split: Vec<&str> = input.split("\n").collect();
    let mut total:usize = 0;
    let mut index = 0;

    while index+3 <= split.len(){
        total = total + score_line_priorities(split[index], split[index+1], split[index+2]);
        index += 3;
    }
    total
}

fn score_line_priorities(s1:&str, s2:&str, s3:&str) -> usize {
    const ASCII_A:usize = 65; 
    const ASCII_a:usize = 97; 
    let mut inventory1 = [0; 52];
    let mut inventory2 = [0; 52];
    let bytes1 = s1.as_bytes();
    let bytes2 = s2.as_bytes();
    let bytes3 = s3.as_bytes();
    
    for (_i,b) in bytes1.iter().enumerate() {
        let val = byte_to_int(&b);
        let val = if val >= ASCII_a { val - ASCII_a } else { val + 26 - ASCII_A };
        inventory1[val] = inventory1[val] + 1;
    }
    
    for (_i,b) in bytes2.iter().enumerate() {
        let val = byte_to_int(&b);
        let val = if val >= ASCII_a { val - ASCII_a } else { val + 26 - ASCII_A };
        if inventory1[val] > 0 {
            inventory2[val] = inventory2[val] + inventory1[val];
        }
    }

    for (_i,b) in bytes3.iter().enumerate() {
        let val = byte_to_int(&b);
        let val = if val >= ASCII_a { val - ASCII_a } else { val + 26 - ASCII_A };
        if inventory2[val] > 0 {
            println!("Found priority {}", val+1);
            return val+1;
        }
    }
    0
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

