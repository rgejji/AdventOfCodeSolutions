fn part1_calculate_score(input: &str)-> u32{
    let split = input.split("\n");
    let mut total:u32 = 0;
    for s in split {
        total = total + score_round(s);
    }
    total
}

fn score_round(s: &str) -> u32 {
    let mut score:u32 = 0;
    let uppercase_diff = 23;
    const ASCII_A:u32 = 65; 

    let b = s.as_bytes();
    let their_ascii:u32 = b[0].into();
    let my_ascii:u32 = b[2].into();
    let their_strat:u32 = (their_ascii-ASCII_A)%3;
    let my_strat:u32 = (my_ascii-ASCII_A-uppercase_diff)%3;

    if my_strat == their_strat {
        score = score + 3;
    }
    else if my_strat == (their_strat +1)%3{
        //println!("won");
        score = score + 6;
    }

    if my_strat == 0 {
        //println!("rock");
        score = score + 1;
    }
    else if my_strat == 1 {
        //println!("paper");
        score = score + 2;
    }
    else {
        //println!("scissors");
        score = score + 3;
    }
    //println!("subtotal: {}", score);
    score

}

fn part2_calculate_score(input: &str)-> u32{
    let split = input.split("\n");
    let mut total:u32 = 0;
    for s in split {
        total = total + part2_score_round(s);
    }
    total
}


fn part2_score_round(s: &str) -> u32 {
    let mut score:u32 = 0;
    let uppercase_diff = 23;
    const ASCII_A:u32 = 65; 

    let b = s.as_bytes();
    let their_ascii:u32 = b[0].into();
    let my_ascii:u32 = b[2].into();
    let their_strat:u32 = (their_ascii-ASCII_A)%3;
    let my_strat:u32 = (my_ascii-ASCII_A-uppercase_diff)%3;
    let my_state = (their_strat+(3+my_strat-1))%3;

    //0 for loss, 3 for tie, 6 for win
    score = score + my_strat*3;
    //score my choice. 0-Rock->1 pt, 1-Paper -> 2 pts, ...
    score = score + my_state + 1;
    score

}


fn main() {
    let score = part1_calculate_score(INPUT);
    println!("Score: {}", score);
    let score = part2_calculate_score(INPUT);
    println!("Score: {}", score);
}

const INPUT:&str = "A Y
B X
C Z";

