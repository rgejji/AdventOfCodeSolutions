use array2d::Array2D;
use std::collections::HashMap;
use std::fs;

const order: [char; 4] = ['N', 'S', 'W', 'E'];


#[derive(Debug,Hash,Clone, Eq, PartialEq, Copy)]
struct Loc(usize,usize);

fn read_input(input:&str, x_buff:usize, y_buff:usize) -> (Array2D<char>,Vec<Loc>) {

    let lines:Vec<&str> = input.split("\n").collect();
    let mut arr = Array2D::filled_with('.', lines.len()+2*y_buff,  lines[0].len()+2*x_buff);
    let mut curr_locs = vec![];

    for (i,line) in lines.into_iter().enumerate(){
        for (j, c) in line.chars().into_iter().enumerate(){
            let y=i+y_buff;
            let x=j+x_buff;
            arr[(y,x)] = c;
            if c == '#'{
                curr_locs.push(Loc(y,x));
            }
        }
    }
    (arr, curr_locs)
}

fn no_neighbors(grid:&Array2D<char>, loc:&Loc) -> bool {
    if grid[(loc.0-1,loc.1)] != '.'{
        return false
    }
    if grid[(loc.0-1,loc.1+1)] != '.'{
        return false
    }
    if grid[(loc.0,loc.1+1)] != '.'{
        return false
    }
    if grid[(loc.0+1,loc.1+1)] != '.'{
        return false
    }
    if grid[(loc.0+1,loc.1)] != '.'{
        return false
    }
    if grid[(loc.0+1,loc.1-1)] != '.'{
        return false
    }
    if grid[(loc.0,loc.1-1)] != '.'{
        return false
    }
    if grid[(loc.0-1,loc.1-1)] != '.'{
        return false
    }
    return true

}


//We unsafely assume the grid is so big, elves won't go off the edge. If we have a segfault, resize
//the x_buff and y_buff to be bigger.
fn calculate_moves(grid:&Array2D<char>, locs:&Vec<Loc>, round:usize) -> HashMap<Loc,Loc>{
    let mut planned_moves = HashMap::new();
    let mut valid_moves = HashMap::new();
   
    for l in locs {
        let mut proposal = Loc(l.0,l.1);
        for i in 0..4{
            if no_neighbors(grid, l){
                break;
            }

            match order[(i+round)%4]{
                'N' => {
                    if grid[(l.0-1,l.1-1)] == '.' && grid[(l.0-1,l.1)] == '.' && grid[(l.0-1, l.1+1)] == '.'{
                        proposal = Loc(l.0-1, l.1);
                        break;
                    }
                },
                'S' => {
                    if grid[(l.0+1,l.1-1)] == '.' && grid[(l.0+1,l.1)] == '.' && grid[(l.0+1, l.1+1)] == '.'{
                        proposal = Loc(l.0+1, l.1);
                        break;
                    }
                },
                'W' => {
                    if grid[(l.0-1,l.1-1)] == '.' && grid[(l.0,l.1-1)] == '.' && grid[(l.0+1, l.1-1)] == '.'{
                        proposal = Loc(l.0, l.1-1);
                        break;
                    }
                },
                'E' => {
                    if grid[(l.0-1,l.1+1)] == '.' && grid[(l.0,l.1+1)] == '.' && grid[(l.0+1, l.1+1)] == '.'{
                        proposal = Loc(l.0, l.1+1);
                        break;
                    }
                },
                _ => println!("Error: Invalid order char!!"),
            }
        }
        //println!("Have proposal {:?}", proposal);
        //store proposal
        match planned_moves.get(&proposal){
            None => planned_moves.insert(Loc(proposal.0, proposal.1), (Loc(l.0,l.1), 1)),
            Some(v) => planned_moves.insert(Loc(proposal.0, proposal.1), (Loc(l.0,l.1), v.1+1)),
        };
    }

    //only propose moves that have only 1 elf going to them and ignore non-moves
    for (proposed, terms) in &planned_moves {
        if (terms.1 == 1) && (proposed.0 != terms.0.0 || proposed.1 != terms.0.1){
            valid_moves.insert(terms.0, Loc(proposed.0, proposed.1));
        }
    }
    valid_moves
}

fn execute_moves(grid:&mut Array2D<char>, locs:&mut Vec<Loc>, moves:&HashMap<Loc, Loc>) {
    for old_loc in locs.iter_mut() {
        match moves.get(old_loc){
            Some(new_loc) => {
                if new_loc != old_loc{
                    grid[(new_loc.0,new_loc.1)]='#';
                    grid[(old_loc.0,old_loc.1)]='.';
                    old_loc.0 = new_loc.0;
                    old_loc.1 = new_loc.1;
                }
            }
            None => {},
        }
    }
}

fn print_arr(arr: &Array2D<char>){
    for i in 0..arr.num_rows(){
        for j in 0..arr.num_columns(){
            print!("{}", arr[(i,j)]);
        }
        println!("");
    }
}

fn calculate_score(locs:&Vec<Loc>) -> usize{
    let num_elves = locs.len();
    let mut min_x = 999999;
    let mut min_y = 999999;
    let mut max_x = 0;
    let mut max_y = 0;

    for l in locs{
        if min_x > l.1{ min_x = l.1;}
        if min_y > l.0{ min_y = l.0;}
        if max_x < l.1{ max_x = l.1;}
        if max_y < l.0{ max_y = l.0;}
    }

    let area = (max_y-min_y+1)*(max_x-min_x+1);
    area-num_elves
}

fn main() {
    let num_moves =100000;
    println!("Hello, world!");
    let buff = 100;
    let input = fs::read_to_string("input.txt").unwrap();
    //let input = fs::read_to_string("test_input.txt").unwrap();
    let (mut arr, mut locs) = read_input(&input, buff, buff);

    println!("Starting with extended array of size ({},{})", arr.num_rows(), arr.num_columns());
    println!("Starting moves...");
    for i in 0..num_moves{
        let moves = calculate_moves(&arr, &locs, i);
        if moves.len() == 0 {
            println!("We encountered no moves on round {}", i+1);
            break;
        }
        execute_moves(&mut arr, &mut locs, &moves);
    }
    println!("Calculating score...");
    println!("Final score is {}", calculate_score(&locs));

}


