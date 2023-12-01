use pathfinding::prelude::astar;
use pathfinding::prelude::bfs;
use pathfinding::prelude::dijkstra;
use array2d::Array2D;

#[derive(Clone, Debug, Eq, Hash, Ord, PartialEq, PartialOrd)]
struct State(i32, i32, i32);

#[derive(Clone, Debug, Eq, Hash, Ord, PartialEq, PartialOrd)]
struct Pos(i32, i32);

impl State {
  fn successors(&self, board:&Vec<Array2D<Vec<char>>>) -> Vec<State> {
    let &State(r, c, n) = self;
    let nr = board[0].num_rows();
    let nc = board[0].num_columns();
    let mut valid = Vec::<State>::new();
    let ordinal = vec![State(r+1,c,n+1), State(r-1,c,n+1), State(r,c+1,n+1), State(r,c-1,n+1), State(r,c,n+1)];
    for ord_pos in ordinal{
        if ord_pos.0 >=0 && ord_pos.0 < nr.try_into().unwrap() && ord_pos.1 >=0 && ord_pos.1 < nc.try_into().unwrap() {
            let next_row:usize = ord_pos.0.try_into().unwrap();
            let next_col:usize = ord_pos.1.try_into().unwrap();
            let next_board_num:usize = ord_pos.2.try_into().unwrap();
            let next_board = &board[next_board_num];
            if next_board[(next_row,next_col)].len() == 0 {
                //println!("Exploring ({},{}) at n={}", ord_pos.0, ord_pos.1, ord_pos.2);
                valid.push(State(ord_pos.0,ord_pos.1,ord_pos.2));
            }
        }
    }
    valid
  }

  fn weighted_successors(&self, board:&Vec<Array2D<Vec<char>>>) -> Vec<(State,i32)> {
    let succ:Vec<State> = self.successors(board);
    succ.into_iter().map(|p| (p,1)).collect()
  }

  fn distance(&self, other: &State) -> i32 {
    return (self.0.abs_diff(other.0) + self.1.abs_diff(other.1)) as i32
  }
}


fn create_array_steps(input:&str, num_steps:i32) -> (Vec<Array2D<Vec<char>>>, State, State) { 
    let splits:Vec<&str> = input.split("\n").collect();
    let num_rows = splits.len();
    let num_cols = splits.get(0).unwrap().len();
   
    let mut array = Array2D::filled_with(Vec::<char>::new(), num_rows, num_cols);
    for (i,line) in splits.iter().enumerate() {
        for (j,b) in line.chars().into_iter().enumerate() {
            if b != '.'{
                array[(i,j)].push(b);
            }
        }
    }

    let start = find_empty(&array,0);
    let end = find_empty(&array, (num_rows-1).try_into().unwrap());
    

    //make master array
    let mut master_array = vec![];
    for _i in 0..num_steps {
        let mut new_array = Array2D::filled_with(Vec::<char>::new(), num_rows, num_cols);
        simulate_move(&array, &mut new_array);
        master_array.push(array);
        array = new_array;
    }

    /*for arr in &master_array{
        println!("{:?}", arr);
    }*/

    (master_array, start, end)
}

fn find_empty(array:&Array2D<Vec<char>>, row: i32)-> State{
    for c in 0..array.num_columns(){
        let r:usize = row.try_into().unwrap();
        if array[(r,c)].len() == 0{
            return State(row,c.try_into().unwrap(),0)
        }
    }
    println!("Error, could not find empty space in row {}", row);
    State(0,0,0)
}


fn simulate_move(A:&Array2D<Vec<char>>, B:&mut Array2D<Vec<char>>){
    let nr = A.num_rows();
    let nc = A.num_columns();
    for i in 0..nr{
        for j in 0..nc{
            for c in &A[(i,j)]{
                match c{
                    '#' => B[(i,j)].push('#'),
                    '>' => {
                        if j+1<nc-1{
                            B[(i,j+1)].push('>')
                        } else {
                            B[(i,1)].push('>')
                        }

                    },
                    '<' => {
                        if j-1>0{
                            B[(i,j-1)].push('<')
                        } else {
                            B[(i,nc-2)].push('<')
                        }
                    },
                    'v' => {
                        if i+1<nr-1{
                            B[(i+1,j)].push('v')
                        } else {
                            B[(1,j)].push('v')
                        }

                    },
                    '^' => {
                        if i-1>0{
                            B[(i-1,j)].push('^')
                        } else {
                            B[(nr-2,j)].push('^')
                        }
                    },
                    _ => println!("ERROR: ENCOUNTERED UNEXPECTED CHARACTER {}", c),

                }
            }
        }
    }
}

fn main() {
    let n = 1000;
    println!("Hello, world!");
    let (master_array, start, goal) = create_array_steps(INPUT, n);
    println!("starting loc is {:?} AND goal IS {:?}", start, goal);
    let begin:State = State(start.0, start.1, 0);
    let result = astar(&begin, |p| p.weighted_successors(&master_array), |p| p.distance(&goal), |p| p.0 == goal.0 && p.1 == goal.1);
    println!("Path took {} steps", result.expect("no path found").1);
    //let result = bfs(&begin, |p| p.successors(&master_array), |p| p.0 == goal.0 && p.1 == goal.1);
    //println!("Path took {} steps", result.expect("no path found").len()-1);

}
/*
const INPUT:&str = "#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#";
*/
const INPUT:&str = "#.########################################################################################################################
#<^vvv<v^v>^vvv.v<vv<>^^<^><.vvv^>v<^<<<<^>^vv^>^<^>>v>v^v^>^v>vv>^.^..v<^<^^^>>>><<>v<^^^v<>>><>.vv^vv>v.vv<<^<.v...^v.<#
#<^^><><^^^<^>>v^^>v<^vvvv><vv<v.<vv^v<>v^.v^>v^v<>^<>^<vv.v^.^<<<^^^..<.><<v<vv^<v^^>>.><v<vv.>v>><<v.v^<^<v^^^^vv^v>^v>#
#<^.>>v<<vv^^<<<<<.<v<>^>v^v<v<<<^<v^<<vvvv>v<^v<^v<>v.<>v^^^.^vv.<vv>v<^^v^>>.><^.><<<>v^>><^.<><>v^vv>>^v>>.><vvv^<<>v<#
#>v^v<>^<.^.^>><>>..<vv<>^>>vv<.<.<^v<^<>>><<^v.v^<v<vv><v.<v.>v^v^<.v>^v^<^.v>.v.>^<v^>vv><<v.>^>v><<^vv<v^>v>v>.>v>>><.#
#><<<^<v.><>>^>v<^><^<^v.^v<vv<<^^vvv><^^>>>>v^.^^v^<v^^.vv^v^<.v<v^^v<^..>>v<^v^.v^^^^.>v^<vv^<<>><<>vv><^<<^^><..vv.v^.#
#<^>^.^v.><^>.v<v<>^<^<<<<<^^><vv.>^v><^<.>>.><^<v>>>^^^vvv<>v^^^<.^v^<^v<v>^.^.<>v^.v>v<v^<^^vv.^<><>v^>>.^>>><>><>v..^<#
#>.^>^>^.v<vv.<>.>>v<>><^vv>>^<<v^.><^^>^^^.<><><.<v<<^.^<><<^>v<v<v<v>^^<^<v<<...<<>v<v<>.^v>>^^^^<<<><v<^>.^v^.>v<v>v.>#
#<.>.^.<>^^<^v<<^^^^v..<><>^vvv<^>..v<>^vvvv<<<>v^<v.v^<<>v.<^v^>^>>vv<.<<.<^^^v^v<vv<.>^<^.v.<<<v.v^v>v^^<vv^^^>v>v^>><>#
#<v>v>^<vvv^v<^>^v>.<<^v^^>v>^v>^.<<v<vv<<v^<v<^vvv.<>.<>^vv.>^vvv<v<><v^.<^v<..><<><^v..v>^>v<>>.<<>^<^v>v.^><<^v<<.v^>>#
#<<<^<<^v^vv><><<>vv<>...<^<<<vv^^^...<><<<^vv^<vv<^<v>^<^^^^>v^>>.<^<^<>>v^><><v^v>>><v>^^<^.>><><^<v<^.>>v^>^v^v<v^^^<<#
#<>.><^v.vv><.v<v>^.>^><<^<><vv>^v<^>v<^.v^.v<>^.>^>vv>^^>>v<^><^v.>v^>^vvv^^<v<>..v^^<>..v^>v^^^<<>vv^vv>v><<v<.>><<vvv>#
#><><.^^<<<<.<<><v^<v>^>^>vv..v>^v<<v^<^^.<<>^>>.<>>^<<v.^>^^vv>^^v.^<<.>>><v^>.>>^.>>^<^v>^<>v<v.v<.<^.v>^<>v^v>.^.><^v>#
#><>>>>^<>>>v^v^v^vv.<^^.^.^>^v>>v<<^..^>><<.^<v><>v>>><<>vv.<>v.^<^^<v^vv^>v<>v><^^vv>vvv.v><v.<<<<>>>vv>v.>v>>^^v>.^<><#
#<<><><^>^v<>v^.^<.^<v^<v.>.>v^v.<^><<vv<vvv^<>^>.v>>v^><><<<v<.^v>><vv<^v<>^v>v.^<^.><.^<.v^v<^^><>^.^>v>.>^<..<<^>^vv..#
#><v>v.v<>>v<<><>>^<>vv^<<>^vvv^<>^>>v><^<v.<^<<<^><<<<<^.^v><><<><><v.vv^^v.<<<<>.<<>v^.<^.v<v>>v^<..^><v<<^^<^>^^v^>.<>#
#.<<>>>v.<>^<<^v.^<>v<><vv>.^v.^>.<v.>v<.^vvv>v^<<v<><^<v^.>vv^.<<^v^v>>vvv>.>>^>^<<<v.>>.vv>>>^v^vv^v<v<<^>>><.<>>^>v^<<#
#<<>><><>>^v>^>.<v^v..v.v^v^>.><^.><^<>^.>>v.^>vv^v><^..<>^vv<>.<<.<>v^^..>^<>v><<^>^v><<v>..>vvv<.>>v<v<vv<^^v>^v<v^v<v.#
#>vvv.><<<^<>>><<<v><>>>>^v<<.<><vvv.<<^^v.v<>^<>^v.^^>.<v>v>v^v^^<^.v^>^><v>^>>^<>>vvvv<.vv^v<^v^^^.^>^><><^vvv.>.v^<<v>#
#>vv.><>>^.^v<.>vv>>.^v^vv>^v^>>><v<<vv.><^><>.><>>><^<>v^v<>^<<<^>^>^<.><.v>><v<>v>>><vv^.>^v.>v^<>v>v^vv^<v<<^v^v<<^>^>#
#<v^vvv<><<<vv.<^v<<v^^>v^<vvv^^vvv^^>v<.<>v^>>>>^>.^><^>^^<^<^.v>>>v>^<<<v>>.v^v<v><^>>v.<^>^v^vv>>^^vv<<v<v<<vv^<^<.vv>#
#<^>v^^^^vv><>>^>v>^>v<vv<><v.v<^vv><v<^^^^^.<<><v<><^.>^^.<^v>^<v>^<<.v>>^v^^.^<.<v>.^><.>^v^^v<><v<.v<.>>><^>^vv<.^.>><#
#<>^^vvv<<v^vv<<>v>.^<>v^>.vv<vv^>.<^>v>v^>.^v<.<>^<vv^v><^v.^v>^^>v^>>><>vv<>vv>.^v><>.<^^v<<<^>^^^><>^^>^.^>vv<.<><<^>>#
#<<^^..>v>vv<v<>>v<><^<v<v^>.<>v>^^>^v<.v^.<>>.<<>>v^><<^^<>^.<v^vvv^v^<<<^v.^vv^^^<v><<<v^<<>^.<><vvv^<>v<^>.>>><>^v<>v<#
#<>>^.>^^^vv<>^>^<v.^^v><.v<^<vv>^^v><<>v<>>>>>.<^^vv^^^>.<<>v^>>.^.<v^>.>.v^vv>><>.><<^^>>v^>>>^>v<>.^^v>^>vvv^^^^.^^^.<#
#>v<vvv<^^>v^>>^^<vvvv.v.v>^<v^>>v><><<v^.<><v<^>v.<^v<>>^v<.^><^v>^^>v^<v^v>>^vv><v.^<><<v<<<v^v>vv><<^<>v^<<^...<<.<>^<#
########################################################################################################################.#";
