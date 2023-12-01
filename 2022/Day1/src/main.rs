


//Day 1 pt 1 solution
fn calculate_best(input: &str) -> i32{
    let split = input.split("\n");
    let mut curr_total = 0;
    let mut best_total = 0;
    for s in split{
        if s.is_empty(){
            if curr_total > best_total {
                best_total = curr_total;
            }
            curr_total = 0;
            continue;
        }
        let value:i32 = s.parse().expect("Encountered not a number");
        curr_total = curr_total + value;
    }
    if curr_total > best_total {
        best_total = curr_total;
    }
    best_total
}

//Day 1 pt 2 solution
fn calculate_best_three(input: &str) -> i32{
    let split = input.split("\n");
    let mut curr_total = 0;
    let mut best_total = 0;
    let mut second_total = 0;
    let mut third_total = 0;
    for s in split{
        if s.is_empty(){
            if curr_total > best_total {
                third_total = second_total;
                second_total = best_total;
                best_total = curr_total;
            } else if curr_total > second_total {
                third_total = second_total;
                second_total = curr_total;
            } else if curr_total > third_total {
                third_total = curr_total;
            }

            curr_total = 0;
            continue;
        }
        let value:i32 = s.parse().expect("Encountered not a number");
        curr_total = curr_total + value;
    }
    if curr_total > best_total {
        best_total = curr_total;
    }
    best_total+second_total+third_total
}

fn main() {
    println!("Hello, world!");
    println!("{}", INPUT);
    let best_score = calculate_best(INPUT);
    println!("best score is {}", best_score);
    let best_three_score = calculate_best_three(INPUT);
    println!("best three score is {}", best_three_score);
}

const INPUT:&str = "1000
2000
3000

4000

5000
6000

7000
8000
9000

1000";
