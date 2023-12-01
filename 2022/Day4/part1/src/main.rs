struct Segment {
    start: u32,
    end: u32,
}

fn count_fully_contains(input: &str) -> u32 {
    let split: Vec<&str> = input.split("\n").collect();
    let mut total:u32 = 0;
    let mut index = 0;

    while index+1 <= split.len(){
        total = total + score_line(split[index]);
        index += 1;
    }
    total
}

fn get_segments_from_str(s:&str) -> (Segment, Segment) {
    let pairs: Vec<&str> = s.split(",").collect();
    let a:&str = pairs[0];
    let b:&str = pairs[1];
    let seg_a = get_segment(a);
    let seg_b = get_segment(b);

    (seg_a, seg_b)
}

fn get_segment(s:&str) -> Segment {
    let bds: Vec<&str> = s.split("-").collect();
    let start:u32 = bds[0].parse().unwrap();
    let end:u32 = bds[1].parse().unwrap();
    Segment {
        start: start,
        end: end,
    }
}

fn score_line(s1:&str) -> u32 {
    let (seg_a, seg_b) = get_segments_from_str(s1);
    
    if seg_a.start <= seg_b.start && seg_b.end <= seg_a.end {
        println!("B in A: {}", s1);
        return 1;
    }
    if seg_b.start <= seg_a.start && seg_a.end <= seg_b.end {
        println!("A in B: {}", s1);
        return 1;
    }
    0
}

fn main() {
    let score = count_fully_contains(INPUT);
    println!("Score: {}", score);
}

const INPUT:&str = "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";
