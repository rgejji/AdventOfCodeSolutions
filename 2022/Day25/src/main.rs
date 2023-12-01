use std::fs;

fn parse_input(input:&str){
    let mut total_sum = 0;
    for line in input.split("\n"){
        if line.eq(""){
            continue;
        }

        let mut sum = 0;
        for c in line.chars().into_iter(){
            sum = sum*5;
            sum = sum + match c{
                '0' => 0,
                '1' => 1,
                '2' => 2,
                '-' => -1,
                '=' => -2,
                _ => {
                    println!("ERROR, UNEXPECTED CHARACTER IN LINE {}", line);
                    0
                },
            };

        }
        println!("SNAFU: {}", sum);
        total_sum += sum;
    }

    println!("Total sum base 10 is {}", total_sum);

    print_base_snafu(total_sum);
}

fn print_base_snafu(total_sum:i64){
    let mut curr = total_sum;
    let mut output = vec![];
    while curr>0 {
        match curr%5 {
            0 => {
                output.push('0');
            },
            1 => {
                curr-=1;
                output.push('1');
            },
            2 => {
                curr-=2;
                output.push('2');
            },
            3 => {
                curr+=2;
                output.push('=');
            },
            4 => {
                curr+=1;
                output.push('-');
            }, 
            _ => {},
        }

        curr = curr/5;
    }
    let result:String = output.into_iter().rev().collect();
    println!("Result is {}", result);
}


fn main() {
    println!("Hello, world!");
    let input = fs::read_to_string("input.txt").unwrap();
    parse_input(&input);
}
