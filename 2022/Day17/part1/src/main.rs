use array2d::Array2D;

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
    while true {
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
            let mut return_height = bloc+piece.h;
            while row_has_element(&grid, return_height){
                return_height += 1;
            }
            return return_height
        } else {
            bloc -=1;
        }
    }
    0
}

//y=0 indicates the floor.
fn main() {
    println!("Hello, worldo!");
    let max_cnt = 2022;
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

    for i in 0..max_cnt {
        height = drop_piece(&wind, &mut wind_loc, &p[i%num_p], &mut grid, height);
    }
    print_grid(&grid, 0, 25);

    println!("Done. Max height is {}", height);
}

fn row_has_element(grid:&Array2D<bool>, height:usize) -> bool{
    for j in 0..grid.num_columns(){
        if grid[(height,j)]{
            return true
        }
    }
    false
}



//const INPUT:&str = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>";

const INPUT:&str = ">>>><>><<>>>><<<>><<<<>><<<>>>><<<>>><<<<>><<>>>><<>>>><<<<>>>><><<<<>>><<>>>><<>>>><<><<<<>>>><<>>>><><<<>>>><<>><<<>><<<<>>>><<<<><>>>><<<><>>><<<<>>>><<>><<>><>><<<><<<<>>>><<<><><<>>><<>>><>>><>><>>>><<<<>>><>><<><>><<<>>>><<<<><><<>><<<>>><><<<<>>>><<<>><<<>>><<<>>><<<><<<>>>><<<<>>><<<<>><<><<<>>>><<<>>>><<<<>><><>><<>><<><<>>><>>><>>><<<><<<<>><<<><<<>><<>><><<<>><<<<><<<>>><<>>><>>><<<><>>>><<>>>><<>>><<<>>><<<<>><><<<>><>>>><<><><<>>>><<<<>>><<><<><<><<<>>>><<<>><<<><<>><>><><<<<>><<<>><<>>>><><<<<><<<<>>><<>>>><<<<><<<>>><<>>>><<>><<<>>><<>>><<>><<<>><<<>>><<>>><<>>>><><<>><<<>><<<><><<<>>>><<<>>>><<<>><<<<><<>>><<<>>>><<>>><>>>><<<>>>><<><>><<<>><<<<>>>><>>><<<><<>>><<<>>>><<<<>>>><<>>>><<><<><<<<><><<<>>><<>>><<<<>><<>>><>>>><<<<>><<>><<<<>>><<>>>><<<<>>>><<<<>><<<><<<<>><<<<>><<>><<>><<<<>>><<<>>><<<<>><>><<<><<>>>><<<<>>>><<<<>>><<<<>><<>>>><><><>>><<>>><<<>>>><>><<<<><<<<><><<>>><<<>>>><<<<>>><>><<<<>>>><>><<<>>><>>>><<<>>>><<>>><<<><<>>><<><<<<>>><<<<>>>><<>><>>><<>>><>>><<<>>><<<<>>><<<><>>>><<<<>><>>><<<>><<<>>><><<<><<><>><<<<>>>><<<>>>><<<><<<>><<<<>>>><<<>>>><<>><<<>><<<>>><<><<<<>>><<<<>><>>>><<>>>><<<>><>>>><<<>>><<>>><><>>>><<<>>>><>>>><<<>>><<<>>>><<>>>><<><<<>><<><>>>><<>>><<>>><>>><<<>>><<><<<><<>><<>>><<<<><<<><<<<>>><<<<><>><<>><<<>>><>><<><><<<<>>>><<>>>><<<<>><<<<>><>>><<<><<<<>><<<<>>><<<><<<>>>><>>><>>>><<<<>><<<<><><>>>><<<>>><<<><<<<><>><<<>>>><<<>>><<<<>>><<<<>><<<>><<<>><<<<>>><>>>><<<><<<>>><>><<>>>><<<>>><>>><<<<>>><<<>>>><>>><<><<>>><<>>>><<<><<<>><<>><<>>>><<<>><>>>><<<>>>><<>><>><>>>><<<>>><<>>><>>><<<>>>><<<>><>>><<<>>>><<<<>>>><<<><<><<>><<<>>>><<<<>>><>>><<>>><<<>><<>>><>>>><<<<>><<<<>>>><<>>>><><<<<>><>><<><><>><<<>><<<<>>><<<>>><<<<><<<<>>>><<<<>>>><<<<>>>><>><<>>>><<<<>><>>><<<<><<>>>><<<<>>><><<<<>>>><><<>>><<<<>>>><<<>><<><<><<>>><><>>><<<<><<<>>><<<<><<<<>>>><<>><<>>>><<<<><<<>>>><>><<<<>>><>><<<><>>><<<><>>><<><<>>><<<<><<<<><<>>>><>>><>>><>><>>>><<<<>>><<<<>>><<>>>><<<>>><<<<>>>><><>>><<<<>><<<><<<>>>><<<<>>><<<<>>><<<>>><<<<>><>><<>><<>>><<<>><>>><<>><<<<>>>><<>><<<>><<>><<<>>><<<>>>><<<>>><<<>>>><<>><<><>>><<<>>>><<<<>>>><<<>>>><>>>><<<>>><<<<>>>><<<>>><>>><>>><><>>>><<<>>>><><>>><<><>><<<<>>><<<<><<<<>>>><<<<>>><>><<<>>><<<>><<<<>>><<<<><<>><<><<><<<<>>>><<>><>><>>><>><<<>>>><<<>>>><<>><<><<<<>>><<><<<>><<<>>><<><<<<>>><<>><>><>>>><<><<<>><<<><<<<>><>><<<><>><<<>><<<>>>><<<<>>><<<<>>><<<><<>>>><<<>><<<<>><<<>>><<<<>>><<<<>><<>><<<>>>><<><<<>><<>>><<<<>>>><<<>><<<<>>>><<>>><<<<><<<>>><<<>>>><>>><<<<><<<>><<<><<<>><><<>>>><>>>><<<<>>>><><<<<><<>>>><><<<><<<>>><<>>>><<<<><><<<<>>><<<<>>><<>>><<>>><<>>><<>>><<<<>><>>><<>><<<>>><><>><<<>>><><>><<>><<<>>>><<<<>>>><>><<<<>>><><<<<>><<>>>><>><<>>>><>>><<<<>><<<<>>><>>>><<<>><>>><<<<>><<<<>><<>>><>>>><<<>>><<<<>>><<<>><<>>><<<>><<<<>>><<>>><>>>><<<>>>><<<>>><<<<><<>><<<<>>>><><<<<>>><><<>>><<<<><<>>>><<<<>>><<<<>>><<<<>><<<<>><<>><<<<>>>><<<>><<<><<<>>>><<>><<>>><<<<><<<<>>><<<><<<<>>>><<>><<>><<<>>><<<<>><<<<>><<<>>><>><<<><<<>><>>>><<<>><<<<>>><<>><<<>><>><>><<<>>><<<>>><<<><>><><>>>><<<>>><>>><<>>>><<<>>>><<><<><>>><<<>>>><<<>><<<>>>><>>>><<<><<>><><>><<<>>>><<<<><<<>>>><<>>>><<<<>>>><<>><>>><<<<>><<<<><>>>><<<>><<>>>><<<<><<<>>><<<>>><<<>>>><<<<>><>><>>><<>>><>>><<<<>><<<<><>>><<<>>><<<><<>><>><>>><<<>>><<<<>>>><<<<>>>><<>><<<>>>><>>>><<<>><>>>><<<<>><<<<>>><>><<<<>><<><<<>>>><<<>>><>><<>>>><>>>><<><><<>>><>><<<<>>><<<<><>>><<<>>>><<>>><<<><<<<>>><<<<>>><><<>>><<>><><<>>><>>><><<<>>><>>><<<<>>><<<>><<<><<<<><<>>>><<<<>>><<<>>>><>>><<>><>>><<<>>><><<<>>>><<<>>>><<<>>><<>>><<<>>>><<<>>><<>><><>>><<>>><<<<>>><<<><<<>>><<<<>><<<>>><<><<><>>>><<>><>>><<>>>><<<<>><<<><<<>>>><>>>><>><<<<>>>><<>>>><<<<>>>><<<>>>><>>>><>>><<<>>>><<<>>>><<>>>><<<<>>>><<>>>><<<><<<>>><<<>><>><<<>><<>><>>><<<>>>><<<>>>><<><<<>>>><<<><><<<<>><>>><<><<<<><<>><<<<>>>><<>><<<>>>><<><>><>>><<<>><>><<<<><>>><>><>><<<<>><>>>><<>>>><<>>><>>><<<<>><>>>><<<<>>>><<>>>><<>><<<<>><>>>><>><>><>>>><><<>><<<>>><<<>>>><<<><<<><<<<>><<<<>>>><<<>><<<>>><<<>>><>>>><>>><<<<>>><<<<>>><<<<>>><<<>><<<>><<<<>>>><<<>>>><<<>>>><<<>>><<>>><<<<>>>><<><<<>>><<><>><><<>>>><<>>><<>>><><<<>><<><<<>>>><<>><<<<>>><>>>><>><<><<<<>>>><>><<>><<><<<>>><<>>><<<<><>><<<>><><<<<><<>><<<<>>><<><<<>>>><<<<><>><<<><<<<>><<<<>>>><>>><<<<>>>><<>>><<<>>>><<<<>><<<>>>><<<>>>><<<<>>><<<<><<<><<>><<<<>>>><<<<>>>><<>>><<<>>>><<<>>>><<>>><<>>><<>>><<>>>><<>>><<>>><<><<<<>>>><>>>><<<><<>>><<>><<<<>><>>>><><<>>>><<<<>>><<>><<<>><<<>>>><<><><<>><<<<><>><<<<>>><><<>>><>><<<>>><<<<><>><>><<>>>><><<<>>>><<<<><<<>><<>>><<>><<<<>><<>>><>>>><<<<>><>>><>>><<>>>><<<>>>><<>><<>>>><>>>><<<<>>>><>>><<<<>>>><<<>>><<><<<>><<<>><<><><<<>>>><<<<>><<<><<<<>>><<<<><<<><<>><<>><<>>><>>>><<<<>><<<>>><>>>><<>><<<<>><>><>>><<<<><<<><<<<><<<<>>>><<<<>>>><>>>><<<>><<>>><><><>><<<>>><<>><<<<>>><<<>><<<>><<><<<<>><>>>><<<>>><<<><>>><>>><<>>><><<<<>><<<>><<>><<>>>><<<>>><><<<>>>><>>>><<<<>>>><<<<><<>>><<>>>><>><<<>>>><<>><<>><>>><>>><<>>><<<<>><>>>><>>><>><<<<><<>>><<<>>><<<>>>><<<>><<><<>><<>>>><<>>><>><<>>><>>>><<<<>>><<>><>><<<><>>>><<>>>><<<>>>><<>><<>>>><<>><<<>><<<<>>>><><<>>>><>><<>><<<>>><<<<>>>><<>>>><<><<>>>><<>><<>>>><>>><<<<>>>><>>><<<>>>><<<>><<<<>>>><<<>>>><>>><<<>>><>><>><<<>><<<><<>><<<>><>>>><>>><>>><<>>><<<>>>><<<<><<<>><<<<>><<<<>><>>>><<<>><<><><><<<<><<<<>>>><<>>><>>>><><<<<>>><>><>><>>><<<>>>><>>>><<><<<<>>><<<<>>><<>><><><><<<>>><<<<>><<>><<<>>>><>>><>>><<>>>><<<<>><<>>><<>><<<><>><<<<>><<>>><<<><<<>>>><<<><<><>>>><<>>>><<<<>><<>><<<><<<>>><>><<<>><<<<><<<>>>><>>>><<<>>>><><>>>><<>><<<<>>><<>><<>><>>><<<<>>><<>><<<>>><<<<><<<<>>><<<><<>>><<<<><<>>>><>><<<<>><><<<<>>><<<>>>><<<<>><><<>><<<>><>>>><<<>>>><<<>>><>><<><<<>><>><<>>><><<<<>>>><<<<>>>><<>>>><>>><<>>><<<><>>><<>>>><<<<>>>><<<>>>><>><<<<>>><<>><<<<>><<<<>><<><<<><>>>><<><<<>>>><>><<>>><>>><<>>><<<>>>><<<<>>>><<>><<<<><<<<><><<<>>><<<<>><<>>><<<>>><<<<>>><<<>><<><<<<>><<><<<>><<<<>><<<<>>>><>>><<<<>><<>><>>>><<>>>><<<<>>>><<<>>>><<<<><<<<>><>><>><<<<>><>><<<<><<<>>>><<<<><<>>>><>><<<<>>><<<>>><><>>><<<>><>><<>><>>>><>><<<>>><<<>><<<<>>><<>>>><<<<><<<>>>><<<<><<>><><<<>><>><<<<>>>><<<>>>><<<>>>><><<<<><<<<>><<<><<>><<>>><<<>>><>>>><<<>>>><<<>><><<<<><<<><><>>>><<<>><<<>>>><<<>>><<<>><<><<><>><<><<<<>>><<<<><><><<>>>><>><<<<>>>><<<<>>><<><<>><<>>><>>>><<<>><<<<><<<><<<<>><<<>>><<>>><<<<>><>>>><<<<>><>><><<<<>>><<<>>><<>>>><<<<>>><<<><<>><<<<><<<><<<<>><<><<><<>>><<<>>>><><<<>>>><<<<>>>><<<>><<<<>><<<<>>><>><<<<>>><>>><<<>><>>><<>>>><>>>><<>><>>>><>><>>><<<>>><>>><<>><<<<>>>><<<<>>><<>>>><<<<>>>><<>>><<>>>><<<>>><<<>>><>>><<<>>><<<>><><<<><<<<>><<>>><<>>><>>><<>><<>>>><<>>><<<>>>><<>>><<<<>><<>>>><>>><<<>>>><<>><>><>>>><<>>><>><><<<<>><<>>><<<<><>>>><<<>><><>>>><<<<>>>><<<>>>><<<>>>><>>><<<><><<<<><>>>><<>>>><<<<>>>><<<<><<<<>><<<<>>>><<<>>>><<><<>><<<>>>><<><<<<>><>>>><<>>>><>>><<<<>>>><<>>>><<<><>>>><<<<>>>><<<<>>><<<<>><><<>><<<><<><<<>><<>><>>>><<><>>><<<>><<<>>><<><<><<<<><>><<><<<>><<<<>><>>><>>>><<<>>>><>>>><<<>><>>>><><>>><<<<>>><<<<>>>><<><><<<><<<>>>><<<>>><<<>><<>>>><<><<<>><<<>><<>>>><<>><<<>>><<<<>>><<<<>><<><>>>><<<>>><<><>><<>><>>>><>>>><>>><<<<>><<><>>><<<<><<<>>>><><<>>>><<<>>>><<>>><<<<>>>><<<<>><>>>><>>>><<<<>>>><<<<>>><<<<><<>><<>>><<<>><<<<><<<>><<>><<<>><<<>><<<>>><<<>>><<>>>><>>>><<<<>>><<<<><<>>>><<>><<<<>><>><<>><<<><<<>>>><>><<<<><<<<>><>>><<>><<<>>>><<<>>>><<>>><<<>>><<><>>>><<<>>><><<<<>><<<><<<<><>>>><>>><<>><<><><><<><>><<<>>>><>>>><<>><<<>><<<<>>>><<>><<<><<>>>><<<<>>><<>>>><<<<>>>><<<>><<><><<<<>>><<>>>><<>>>><<<<>>>><<<<><<<>>>><<<<>><<<<>>><<<<><<<<><<><<<<>>><<>>>><>>><><<<<>><<>>><<<<>>>><<<>>><<<<><><<<<><<>>>><<>><<<<>><><<<<>>><<<<>>>><<>>>><<<<>><<<<><>>><<>>>><<<>>><<><>>><<<>>>><<<><<<<><<><<<>><<<<><<<>><><<<>><><>>><<<<>>><<<>>>><><<<>><<<<><<<<>><<<>><<>><><>>><<<>><<<<>>><<<><>>><<<><>><<>><<>><<<>>>><<<>>>><><<><<<<>>><<>>>><<>>><<>><<<>><<<<>>>><<<><<<><<><<<>>><><<<>>>><<><<>><<<>><>><<>>>><<<>>><<><<>><>><<<<>>>><<>>><<><<<<>>>><<><<>>><>>><<>>>><<<<>><<>><<<<>>>><<<>>><>><><<<>><<<<>><>><<<>><<>><<>><<<<>><<<<><><<<>>>><<<<><<<<>><<<<>>>><<>>><<<><<>>><<<<><<<<>>>><<<>>>><<<<><<<>>><<<>>>><<<<>><<<>>><<<<>>><<<>>>><>>>><<<>>><<<>>>><<<<>><<<<>>><<>>>><>>>><<>>>><<<>><<<<>>><<<>>><>>><<<>>>><<>>>><>>><<><>><<>>>><<><<<>>><<<<>>><<>>><<<>><<<>>><<>>><<<>>>><<<>>>><>>><<<<><>>>><<<>>><><<<><<<>><<<>>><<<<>><<>><<<<>>><<<>>><><<>>><<<>>>><>><<>>>><<<>><<<<>><><>>><>>><<>>>><<>>><><<<<>>>><<>>>><<<<>>>><<<<>><<<<>>>><<>>>><<>>><>><>>><<<>>>><<<><<>>><<<<><<<><<>>>><>><<<<>><<>>><><<<><<<><<<>>>><>><><<<<>>>><<<<>>><>>><<<>>>><>><<<>><<><<<<>><>>><>>>><<<<>><<>><>><<<>>>><><>>>><<>>>><>>><<<><<<>><<><<>>><<<<>><<>>><<>>><<<<>>><<<>><<<>><<>><<<<>>>><><<><<<<><<>><<<<><<>>><<>>><<><<<>>><<><<<<>>>><<<>>><<<<>>><<<>>><>>><<>><<<><<<>>>><>><<<<>><<<>>>><<<><<>><<<<>><>><<<<>><<<<>>><<<<>>><>><<<>>><>>><>>>><>>>><<<>>><<<><<<>><<>>>><<<>>><><<>>><>>><>>><<>>>><>>>><<<>><<<><>><<<<>>>><<>>><<<<>>><<<>><<<>>><<<>><>>>><<<<><<<<>>><>>><<<<>>>><<<>><<<<><<>>><<>>>><<<><<<><>>>><<>>><<<<><<<<><<<<><<<<>>><<<<>>><<>><<<><<<>>><<>><>>>><>>><<<<>><<>>><<>>><<<<><<<>>>><<><>>><<<<><<<><<<<>>>><>>>><<<<>>>><<>>>><<>>><<<>>>><<>>><<>><<<<>><<>>>><<><<<<>>>><<<<>>><<><>><<<>><>><><>>>><<<>>>><<<<>>>><<>><<>>><<<<><<<>>><<<>>>><>>>><<>>><<>>><<<>>>><<>>><>>>><<>>>><<<<>>><>><<<>>><<<<>>><>>>><>>>><<<<>>><>>><>><<>>><>><<<<>>><>>>><>>><<<><<><<<<><<<<>>>><<>>><<<<><<<>>>><<<<>>>><<>><<<>>>><<<<>><<>>>><<<>>><>><<<<>>>><<<<>>>><<<<>><><<<<>>><<<<>><<<<>>><<<>><<>><<>>>><<<<>>>><<><>>><<<>>>><<<<>>>><<<>>><>>>><<>>><><<<<>><<<><<>>><<<<><<>>><<<><<<><<<<>>>><<>><>>><<<<>>>><>><><<<<>>><<>>><<<>>><>>>><>>>><><>>><<<<>>><<<>>><<><<<>><<<>>><<<>>><<<<>><<<>>>><<<>>>><<<>><<<<><<<<>><><>>><<>>><<<<><<><<>><<<<>>><<<<>>><<<<>><<><<<>><><<>><<<>>>><<<><<<><<<<>>>><<<><<<<>><<<>>>><>>><<<<>>><<<<>>>><<>><><<<>>>><<>><<<<>>><<<>><<<<><<<>><>>><<<><<<<><<<>>><<<<>>><<<<><<>>><>><<<<>><<><<>>><<<<><<>>><<><<>>>><<<<><<<><>>><<>>>><<>>>><><<>><<<>>><>><<><<<<>>>><<>>>><<<>>>><>>><<<>><<<>>><><>><<>><>><<>><<<><<<<>>>><<<>>><<>><<<>>><<><>>><<>><<<>>><<<><<>>>><>><>>><><<<>>>><<<<><<<>><<<><<<><>>><<<>>><<>><>>>><>>><<";


