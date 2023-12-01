use array2d::Array2D;
use std::cmp::max;
use std::collections::HashMap;

struct Piece {
    t:usize,
    w:usize,
    h:usize,
    fill:Vec<(usize,usize)> 
}

fn read_wind(input:&str) -> Vec<i64> {
    let mut v = vec![];
    for c in input.chars(){
        if c == '>'{
            v.push(1);
        }
        if c == '<' {
            v.push(-1);
        }
    }
    v
}


fn no_collision(piece:&Piece, grid: &Array2D<bool>, height:usize, lloc:usize) -> bool {
    for (a,b) in &piece.fill {
        if grid[(height+a,lloc+b)] {
            return false
        }
    }
    return true
}

fn apply_collision(piece:&Piece,grid:&mut Array2D<bool>, height:usize, lloc:usize) {
    for (a,b) in &piece.fill {
        grid[(height+a,lloc+b)] = true;
    }
}

fn print_grid(grid:& Array2D<bool>, start_row:usize, num_rows:usize){
    let end_row = start_row+num_rows;
    for i in 0..num_rows {
        for j in 0..grid.num_columns(){
            if grid[(end_row-1-i,j)]{
                print!("#");
            } else {
                print!(".");
            }
        }
        println!("");
    }
    println!("~~~~~~~");
    println!("");
}

fn drop_piece(wind:&Vec<i64>, wind_loc:&mut usize, piece:&Piece, grid:&mut Array2D<bool>, height:usize) -> usize {
    let num_wind:usize = wind.len();
    let mut lloc:usize = 2;
    let mut bloc = (height + 3).try_into().unwrap();

    //println!("Height is {}, starting piece at x={},y={}", height, lloc, bloc);
    loop {
        //println!("Wind loc {}", *wind_loc);
        //println!("Piece {} at lloc {} and height {}", piece.t, lloc, bloc);
        //wind tries to move rock
        let dir = wind[*wind_loc];
        *wind_loc = (*wind_loc+1)%num_wind;
        match dir {
            -1  =>  {
                if lloc >= 1 && no_collision(&piece, &grid, bloc, lloc-1) {
                    lloc -=1;
                }
                //println!("moves left x={}, y={}", lloc, bloc);
            }
            _ => {
                if lloc+piece.w < 7 && no_collision(&piece, &grid, bloc, lloc+1) {
                    lloc+=1;
                }
                //println!("moves right x={}, y={}", lloc, bloc);
            }

        }

        //rock falls
        if bloc == 0 || !no_collision(&piece, &grid, bloc-1, lloc) {
            //println!("Stopping x={}, y={}", lloc, bloc);
            apply_collision(&piece, grid, bloc, lloc);
            return max(height, bloc+piece.h)
        } else {
            bloc -=1;
        }
    }
}

fn calculate_hash(index:usize, piece:usize, wind:usize, height:usize, grid:&Array2D<bool>, hashes:&mut HashMap<String, (usize, usize)>) -> (usize, usize) {
    let mut val:usize = 0;

    for i in height-10..height{
        for j in 0..grid.num_columns(){
            val = val << 1;
            if grid[(i,j)]{
                val +=1;
            }
        }
    }
    let hash = format!("{}_{}_{}", piece, wind, val).to_string();
    //println!("Hash: {}", hash);
    if hashes.contains_key(&hash){
        return *hashes.get(&hash).unwrap()
    }
    hashes.insert(hash, (index, height));
    return (0,0)

}

//y=0 indicates the floor.
fn main() {
    println!("Hello, worldo!");
    let max_cnt = 1000000;
    
    //let max_cnt =1000;
    //let max_cnt =10;
    
    //let max_cnt = 12;
    //let max_cnt = 10;

    let p = vec![Piece { t:0, w:4, h:1, fill:vec![(0,0),(0,1),(0,2),(0,3)]},
        Piece {t:1, w:3, h:3, fill:vec![(1,0),(0,1),(1,1),(2,1),(1,2)]},
        Piece {t:2, w:3, h:3, fill:vec![(0,0),(0,1),(0,2), (1,2), (2,2)]},
        Piece {t:3, w:1, h:4, fill:vec![(0,0), (1,0), (2,0), (3,0)]},
        Piece {t:4, w:2, h:2, fill:vec![(0,0), (1,0), (0,1), (1,1)]}];
    let num_p = p.len();

    let wind = read_wind(INPUT);
    let mut wind_loc = 0;
    let mut height:usize = 0;

    let mut grid = Array2D::filled_with(false, 3*max_cnt, 7);
    let mut height_history = vec![];
    let mut height_a:usize =0;
    let mut height_b:usize =0;
    let mut i_a:usize = 0;
    let mut i_b:usize = 0;
    let mut hashes = HashMap::new();
    for i in 0..max_cnt {
        let piece = i%num_p;
        height = drop_piece(&wind, &mut wind_loc, &p[piece], &mut grid, height);
        height_history.push(height);
        //println!("HEIGHT IS {}", height);
        if height > 10 {
            //calculate hash and return collision
            (i_a, height_a) = calculate_hash(i, piece, wind_loc%wind.len(), height, &grid, &mut hashes);
            //handle collision
            if height_a != 0 {
                i_b = i;
                height_b = height.try_into().unwrap();
                println!("We found duplicate hashes {} and {}", height_a, height_b);
                break;
            }
        }
    }

    if height_b == 0{
        println!("error: we did not find a duplicate. Exit now");
    }

    const trillion:usize = 1000000000000;
    let freq = i_b-i_a;
    let mut num_skips = trillion/freq;
    while num_skips*freq+i_a > trillion{
        num_skips-=1;
    }

    println!("Num skips: {}", num_skips);
    println!("Estimate: {}", num_skips*(height_b-height_a));

    println!("Num skip steps: {}", (num_skips-1)*freq);
    let num_pieces_after_skip = trillion - ((num_skips-1)*freq+i_b);
    println!("Performing remaining {} steps", num_pieces_after_skip);
    height = (height_history[num_pieces_after_skip+i_a]-height_history[i_a]) + (height_b-height_a)*num_skips + height_a;

    //Account for height being one off
    println!("Done. Max height is {}", height-1);
}


const INPUT:&str = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>";



