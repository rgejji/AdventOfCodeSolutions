fn score_input(input: &str) -> usize {
    const dist_char:usize = 14;

    let mut total:u32 = 0;
    let bytes = input.as_bytes();
    for (i,b) in bytes.iter().enumerate() {
        //add new one-encoding
        let new_val = byte_to_one_encoding(b);
        total |= new_val;
        //remove old one-encoding if we are past the 4th character and there is no repeat in
        //characters in the last dist_char
        if i>= dist_char{
            let mut overlap = false;
            for j in 0..dist_char{
                if bytes[i-j] == bytes[i-dist_char]{
                    overlap = true;
                }
            }
            if !overlap{
                let old_val = byte_to_one_encoding(&bytes[i-dist_char]);
                //println!("removing: {:b}", old_val);
                total ^= old_val;
            }
        }
        //println!("{:b} {:b}", new_val, total);
        if total.count_ones() == u32::try_from(dist_char).unwrap() {
            return i+1;
        }
    }
    0
}


fn byte_to_one_encoding(c:&u8) -> u32 {
    const ASCII_a:u32 = 97; 
    let result:u32 = (*c).into();
    let shift = u32::try_from(result).unwrap()-ASCII_a;
    1 << shift
}



fn main() {
    println!("Hello, world!");
    let score = score_input(INPUT);
    println!("Score: {}", score);

}

const INPUT:&str = "bvwbjplbgvbhsrlpgdmjqwftvncz";
//const INPUT:&str = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw";                   
