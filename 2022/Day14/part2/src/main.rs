use array2d::Array2D;
use core::cmp::min;
use core::cmp::max;

const empty_space:char = '.';


//drop sand will drop the sand, return a bool if the drop was a success and the loc it dropped
//(0,0) on a fail
fn drop_sand(grid:&mut Array2D<char>, r:usize, c:usize) -> ((usize,usize), bool) {
    if r==0 && grid[(r,c)] != empty_space{
        println!("We have successfully filled!!");
        return ((0,0),false);
    }

    if r+1 >= grid.num_rows() {
        println!("Error: Exceeded row");
        return ((r+1,c), false);
    }
    if grid[(r+1,c)] == empty_space {
        return drop_sand(grid, r+1, c)
    }
    if c == 0 {
        println!("Error: Exceeded left bdry");
        return ((r+1,0), false);
    }
    if grid[(r+1,c-1)] == empty_space {
        return drop_sand(grid, r+1, c-1)
    }
    if c+1 >= grid.num_columns() {
        println!("Error: Exceeded right bdry");
        return ((r+1,c+1), false);
    }
    if grid[(r+1,c+1)] == empty_space {
        return drop_sand(grid, r+1, c+1)
    }
    grid[(r,c)] = 'o';
    return ((r,c), true)
}

fn get_pair_int(s: &str, index:usize) -> usize{
    s.split(",").collect::<Vec<&str>>()[index].parse().unwrap()
}


fn read_paths(input: &str) -> (Array2D<char>, usize, usize){
    let mut min_r:usize = 9999999;
    let mut min_c:usize = 9999999;
    let mut max_r:usize = 0;
    let mut max_c:usize = 0;
    //read min and max
    for line in input.split("\n"){
        for pair in line.split(" -> "){
            let valA = get_pair_int(pair,1);
            let valB = get_pair_int(pair,0);
            if valA < min_r {
                min_r = valA;
            }
            if valA > max_r {
                max_r = valA;
            }
            if valB < min_c {
                min_c = valB;
            }
            if valB > max_c {
                max_c = valB;
            }
        }
    }

    println!("Mins are ({},{}) and maxes are ({}, {})", min_r, min_c, max_r, max_c);
    //we want to fill to the 0th row, so we do not shift
    let shift_r = 0;
    //we have two plus the bottom
    let num_rows = 3+max_r;
    //pyramid is at most num_rows^2 elements full and 1+2*num_rows wide
    let num_columns = 1+2*num_rows;
    let shift_c = 500-(num_rows-1);
    println!("vals are {} {} {}", num_rows, num_columns, shift_c);

    let mut grid = Array2D::filled_with(empty_space, num_rows, num_columns);

    //fill bottom
    for c in 0..num_columns{
        grid[(num_rows-1,c)] = '#';
    }



    for line in input.split("\n"){
        let mut prev_r = 9999;
        let mut prev_c = 9999;

        for pair in line.split(" -> "){
            let val_r = get_pair_int(pair,1);
            let val_c = get_pair_int(pair,0);
            let r = val_r-shift_r;
            let c = val_c-shift_c;
            
            if prev_r != 9999 || prev_c != 9999 {
                //check horizontal movement
                if r == prev_r{
                    for curr_c in min(prev_c,c)..(max(prev_c,c)+1) {
                        grid[(r,curr_c)] = '#';
                    }
                } else {
                    for curr_r in min(prev_r,r)..(max(prev_r,r)+1) {
                        grid[(curr_r,c)] = '#';
                    }
                }
            } 
            prev_r = r;
            prev_c = c
        }

    }
    (grid, shift_r, shift_c)
}

fn print_grid(grid: &Array2D<char>){
    for i in 0..grid.num_rows(){
        for j in 0..grid.num_columns(){
            print!("{}",grid[(i,j)]);
        }
        println!("");
    }
    println!("");
}

fn main() {
    println!("Hello, world!");
    let (mut grid, _, shift_c) = read_paths(INPUT);
    let mut not_done = true;
    let mut cnt = 0;
    let mut val = (0,0);
    //print_grid(&grid);
    while not_done {
        (val, not_done) = drop_sand(&mut grid, 0, 500-shift_c);
        //print_grid(&grid);
        cnt+=1;
    }
    println!("Could not place on cnt {}. Answer is {}", cnt, cnt-1);
    println!("Exited with val {:?}", val);
}

/*
const INPUT: &str = "498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9";
*/
const INPUT: &str = "497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
525,75 -> 525,77 -> 517,77 -> 517,83 -> 532,83 -> 532,77 -> 531,77 -> 531,75
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
513,169 -> 518,169
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
510,34 -> 514,34
530,86 -> 530,87 -> 543,87 -> 543,86
517,150 -> 521,150
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65
501,177 -> 506,177
505,115 -> 505,119 -> 502,119 -> 502,126 -> 514,126 -> 514,119 -> 510,119 -> 510,115
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
511,52 -> 516,52
505,115 -> 505,119 -> 502,119 -> 502,126 -> 514,126 -> 514,119 -> 510,119 -> 510,115
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
504,30 -> 508,30
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
523,150 -> 527,150
525,175 -> 530,175
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65
525,75 -> 525,77 -> 517,77 -> 517,83 -> 532,83 -> 532,77 -> 531,77 -> 531,75
532,52 -> 537,52
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
525,75 -> 525,77 -> 517,77 -> 517,83 -> 532,83 -> 532,77 -> 531,77 -> 531,75
505,115 -> 505,119 -> 502,119 -> 502,126 -> 514,126 -> 514,119 -> 510,119 -> 510,115
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
502,103 -> 502,105 -> 498,105 -> 498,112 -> 507,112 -> 507,105 -> 504,105 -> 504,103
510,30 -> 514,30
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
511,150 -> 515,150
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
514,147 -> 518,147
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
511,175 -> 516,175
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
507,32 -> 511,32
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
505,115 -> 505,119 -> 502,119 -> 502,126 -> 514,126 -> 514,119 -> 510,119 -> 510,115
508,177 -> 513,177
502,103 -> 502,105 -> 498,105 -> 498,112 -> 507,112 -> 507,105 -> 504,105 -> 504,103
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
515,177 -> 520,177
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
521,49 -> 526,49
520,153 -> 524,153
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
507,49 -> 512,49
518,175 -> 523,175
505,115 -> 505,119 -> 502,119 -> 502,126 -> 514,126 -> 514,119 -> 510,119 -> 510,115
514,49 -> 519,49
521,173 -> 526,173
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
495,32 -> 499,32
504,34 -> 508,34
510,171 -> 515,171
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
514,129 -> 514,133 -> 507,133 -> 507,141 -> 518,141 -> 518,133 -> 517,133 -> 517,129
517,71 -> 517,72 -> 527,72 -> 527,71
514,129 -> 514,133 -> 507,133 -> 507,141 -> 518,141 -> 518,133 -> 517,133 -> 517,129
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65
526,153 -> 530,153
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
525,75 -> 525,77 -> 517,77 -> 517,83 -> 532,83 -> 532,77 -> 531,77 -> 531,75
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
507,68 -> 507,69 -> 523,69
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
513,32 -> 517,32
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
502,103 -> 502,105 -> 498,105 -> 498,112 -> 507,112 -> 507,105 -> 504,105 -> 504,103
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
516,40 -> 521,40
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
520,36 -> 520,37 -> 528,37 -> 528,36
498,34 -> 502,34
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
514,129 -> 514,133 -> 507,133 -> 507,141 -> 518,141 -> 518,133 -> 517,133 -> 517,129
498,30 -> 502,30
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65
504,175 -> 509,175
502,103 -> 502,105 -> 498,105 -> 498,112 -> 507,112 -> 507,105 -> 504,105 -> 504,103
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
510,46 -> 515,46
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
520,36 -> 520,37 -> 528,37 -> 528,36
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
514,129 -> 514,133 -> 507,133 -> 507,141 -> 518,141 -> 518,133 -> 517,133 -> 517,129
516,34 -> 520,34
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65
525,75 -> 525,77 -> 517,77 -> 517,83 -> 532,83 -> 532,77 -> 531,77 -> 531,75
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
514,173 -> 519,173
492,34 -> 496,34
525,52 -> 530,52
502,103 -> 502,105 -> 498,105 -> 498,112 -> 507,112 -> 507,105 -> 504,105 -> 504,103
507,68 -> 507,69 -> 523,69
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65
520,36 -> 520,37 -> 528,37 -> 528,36
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
505,115 -> 505,119 -> 502,119 -> 502,126 -> 514,126 -> 514,119 -> 510,119 -> 510,115
507,173 -> 512,173
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
504,52 -> 509,52
502,103 -> 502,105 -> 498,105 -> 498,112 -> 507,112 -> 507,105 -> 504,105 -> 504,103
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
517,71 -> 517,72 -> 527,72 -> 527,71
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
524,46 -> 529,46
508,153 -> 512,153
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
530,86 -> 530,87 -> 543,87 -> 543,86
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65
504,26 -> 508,26
514,129 -> 514,133 -> 507,133 -> 507,141 -> 518,141 -> 518,133 -> 517,133 -> 517,129
522,177 -> 527,177
525,75 -> 525,77 -> 517,77 -> 517,83 -> 532,83 -> 532,77 -> 531,77 -> 531,75
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
514,129 -> 514,133 -> 507,133 -> 507,141 -> 518,141 -> 518,133 -> 517,133 -> 517,129
517,46 -> 522,46
507,28 -> 511,28
514,153 -> 518,153
517,144 -> 521,144
517,71 -> 517,72 -> 527,72 -> 527,71
525,75 -> 525,77 -> 517,77 -> 517,83 -> 532,83 -> 532,77 -> 531,77 -> 531,75
514,129 -> 514,133 -> 507,133 -> 507,141 -> 518,141 -> 518,133 -> 517,133 -> 517,129
520,147 -> 524,147
530,86 -> 530,87 -> 543,87 -> 543,86
518,52 -> 523,52
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
517,171 -> 522,171
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
505,115 -> 505,119 -> 502,119 -> 502,126 -> 514,126 -> 514,119 -> 510,119 -> 510,115
497,166 -> 497,165 -> 497,166 -> 499,166 -> 499,160 -> 499,166 -> 501,166 -> 501,157 -> 501,166 -> 503,166 -> 503,163 -> 503,166 -> 505,166 -> 505,164 -> 505,166 -> 507,166 -> 507,162 -> 507,166 -> 509,166 -> 509,157 -> 509,166 -> 511,166 -> 511,162 -> 511,166 -> 513,166 -> 513,162 -> 513,166 -> 515,166 -> 515,163 -> 515,166
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
528,49 -> 533,49
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65
529,177 -> 534,177
501,32 -> 505,32
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
501,28 -> 505,28
493,23 -> 493,15 -> 493,23 -> 495,23 -> 495,15 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,17 -> 499,23 -> 501,23 -> 501,17 -> 501,23 -> 503,23 -> 503,14 -> 503,23 -> 505,23 -> 505,16 -> 505,23
502,103 -> 502,105 -> 498,105 -> 498,112 -> 507,112 -> 507,105 -> 504,105 -> 504,103
535,100 -> 535,93 -> 535,100 -> 537,100 -> 537,92 -> 537,100 -> 539,100 -> 539,99 -> 539,100 -> 541,100 -> 541,96 -> 541,100 -> 543,100 -> 543,92 -> 543,100 -> 545,100 -> 545,97 -> 545,100 -> 547,100 -> 547,95 -> 547,100 -> 549,100 -> 549,93 -> 549,100 -> 551,100 -> 551,90 -> 551,100
513,43 -> 518,43
520,43 -> 525,43
505,65 -> 505,58 -> 505,65 -> 507,65 -> 507,55 -> 507,65 -> 509,65 -> 509,60 -> 509,65 -> 511,65 -> 511,56 -> 511,65 -> 513,65 -> 513,62 -> 513,65";
