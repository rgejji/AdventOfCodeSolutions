use std::collections::HashMap;
use std::fs;

#[derive(Clone, Debug)]
enum Action {
    Move(usize),
    Turn,
}

#[derive(Clone, Debug)]
struct Node {
    nbrs:Vec<usize>,
    valve_index:Option<usize>,
    name:String,
}


#[derive(Clone, Debug)]
struct State{
    loc: usize,
    e_loc: usize,
    history: usize,
    total: usize,
    visited: bool
}

impl State {
    fn get_state_index(&self, num_nodes:usize) -> usize {
        self.history*num_nodes*num_nodes+self.e_loc*num_nodes+self.loc
    }

    fn get_delta_pressure(&self, pressure:&Vec<usize>) -> usize {
        let mut sum = 0;
        let mut mask = 1;

        for p in pressure {
            if self.history & mask > 0{
                sum+=p;
            }
            mask = mask<<1;
        }

        sum
    }
}

fn construct_states(input:&str) -> (Vec<State>, Vec<Node>, Vec<usize>, HashMap<String, usize>) {
    let mut name_to_index = HashMap::new(); //names to index
    let mut node_names = vec![];  //index to names
    let mut state_list = vec![];
    let mut pressure = vec![];
    let mut node_cnt = 0;
    let lines = input.split("\n").collect::<Vec<&str>>();
    let mut node_list = vec![Node{nbrs: vec![], valve_index:Some(0), name:"".to_string()}; lines.len()];
    for line in lines.iter() {
        if line.is_empty(){
            continue;
        }

        let mut valve_index = None;
        let mut nbrs = vec![];
        let mut node_index = node_cnt;
        let name:String = String::from(line.split(" ").collect::<Vec<&str>>()[1]);
        if !name_to_index.contains_key(&name){
            name_to_index.insert(name.clone(), node_cnt);
            node_names.push(name.clone());
            node_cnt+=1;
        } else {
            node_index = *name_to_index.get(&name).unwrap();
        }

        let rate:usize = line.split("=").collect::<Vec<&str>>()[1].split(";").collect::<Vec<&str>>()[0].parse().unwrap();
        if rate >0 {
            valve_index = Some(pressure.len());
            pressure.push(rate);
        }
        let edge_strs = 
            match line.split("valves ").collect::<Vec<&str>>().get(1).clone(){
                Some(val) => val,
                None => {
                    let tmp_b = line.split("valve ").collect::<Vec<&str>>();
                    tmp_b.get(1).unwrap().clone()
                }
            };    
        for edge_end in edge_strs.split(", "){
            match name_to_index.get(edge_end) {
                Some(v) => {
                    nbrs.push(*v);
                },
                None => {
                    name_to_index.insert(edge_end.to_string(), node_cnt);
                    node_names.push(edge_end.to_string());
                    nbrs.push(node_cnt);
                    node_cnt+=1;
                }
            }
        }
        //println!("Index {} connects to {:?}", node_index, nbrs);
        node_list[node_index] = Node{
            nbrs: nbrs,
            name: name,
            valve_index: valve_index,
        };
        
    }
    //println!("Name to Index {:?}", name_to_index);
    //println!("Index to edge {:?}", node_names);
    //println!("We found {} nodes", node_list.len());
    for no in &node_list{
        println!("Nodes: {:?}", no);
    }
    println!("Pressures are {:?}", pressure);
    println!("Creating states for {} pressures, which is history of size {}", pressure.len(), 1<<pressure.len());
    for i in 0..(1 << pressure.len()){
        for j in 0..node_list.len(){
            for k in 0..node_list.len(){
                state_list.push(State {
                    loc: k,
                    e_loc: j,
                    history: i,
                    total:0,
                    visited: false,
                });
            }
        }
    }
    //Set AA as visited in the sense that we are there
    let aaloc = name_to_index.get("AA").unwrap();
    let start_loc = node_list.len()*aaloc +aaloc;
    state_list[start_loc].visited = true;
    println!("AA loc was found to be at {}. With {} nodes we are starting at {}", aaloc, node_list.len(), start_loc);
    println!("{:?}", state_list[start_loc]);
    (state_list, node_list, pressure, name_to_index)
}


fn turn_on_valve(history:usize, valve_index:Option<usize>) -> usize {
    match valve_index{
        None => return history,
        Some(v) => {
            let mask = 1<<v;
            return mask | history
        },
    }
}


fn iterate_states(curr_states:&mut Vec<State>, nodes:&Vec<Node>, pressures:&Vec<usize>, name_to_ind:&HashMap<String,usize>){
    let end_time = 26;
    //let end_time = 6;
    let mut alt_states = curr_states.clone();

    /*
    let debug_state = State{
        loc: *name_to_ind.get("BB").unwrap(),
        e_loc: *name_to_ind.get("HH").unwrap(),
        history: 1<<2 | 1<<5 | 1<<4 | 1<<0,
        //history: 1<<2 | 1<<5,
        total:0,
        visited: false
        };
    let debug_ind = debug_state.get_state_index(nodes.len());
    */

    for _i in 0..end_time/2 {
        iterate(curr_states, &mut alt_states, nodes, pressures);
        print_highest_total(&alt_states);
        //println!("Test state is at {:?}", curr_states[debug_ind]);

        iterate(&mut alt_states, curr_states, nodes, pressures);
        print_highest_total(curr_states);
        //println!("Test state is at {:?}", curr_states[debug_ind]);
    }
    if end_time%2==1{
        iterate(curr_states, &mut alt_states, nodes, pressures);
        *curr_states = alt_states;
    }
}

fn iterate(curr_states:&mut Vec<State>, next_states:&mut Vec<State>, nodes:&Vec<Node>, pressure:&Vec<usize>){
    let mut num_explored = 0;
    for s in curr_states.iter_mut() {
        if !s.visited{
            continue;
        }
        s.total +=s.get_delta_pressure(pressure);
    }
    for (i,s) in curr_states.into_iter().enumerate() {
        if !s.visited {
            continue;
        }

        //Note: We don't need to copy things over here, since we allow for no-moves in get_nbrs
        //explore using yourself or elephant 
        let nbrs = get_nbrs(s, nodes);
        for n in nbrs{
            if s.total >= next_states[n].total{
                next_states[n].total = s.total;
                next_states[n].visited = true;
            }
        }
        num_explored+=1;
    }
    println!("Explored {} states this round", num_explored);
}

fn print_num_visited(curr_states:&mut Vec<State>){
    let mut cnt = 0;
    for s in curr_states{
        if s.visited{
            cnt+=1;
        }
    }
    println!("{} are now marked as explored", cnt);
}

//returns the state index of the neighbors doing one step from either your loc, or the elephants..
fn get_nbrs(curr:&State, nodes:&Vec<Node>)-> Vec<usize>{
    let mut my_actions = vec![];
    let mut e_actions = vec![];
    let mut nbrs = vec![];

    for n in &nodes[curr.loc].nbrs{
        my_actions.push(Action::Move(*n));
    }
    my_actions.push(Action::Move(curr.loc));
    my_actions.push(Action::Turn);

    for n in &nodes[curr.e_loc].nbrs{
        e_actions.push(Action::Move(*n));
    }
    e_actions.push(Action::Move(curr.e_loc));
    e_actions.push(Action::Turn);

    for a in &my_actions{
        for b in &e_actions{
            let mut tmpState = curr.clone();
            match a {
                Action::Move(v) => {tmpState.loc = *v},
                Action::Turn => {
                    tmpState.history = turn_on_valve(tmpState.history,nodes[curr.loc].valve_index);
                    if tmpState.history == curr.history {
                        continue;
                    }
                }
            }
            match b {
                Action::Move(v) => {tmpState.e_loc = *v},
                Action::Turn => {
                    tmpState.history = turn_on_valve(tmpState.history,nodes[curr.e_loc].valve_index);
                    if tmpState.history == curr.history {
                        continue;
                    }
                }
            }
           
            /*
            match (a,b) {
                (Action::Turn, Action::Turn) => {
                    if tmpState.loc == 3 && tmpState.e_loc == 8{
                        println!("Before the double turn at {:?}", curr);
                        println!("Found the double turn at {:?}", tmpState);
                    }
                },
                _ => {}
            }*/


            /*
            let history_check = 1<<2 | 1<<5 | 1<<4 | 1<<0;
            if tmpState.loc == 3 && tmpState.e_loc == 8{
                //println!("Want history {} Adding State {:?}", history_check, tmpState);
            }*/

            nbrs.push(tmpState.get_state_index(nodes.len()));
        }
    }
    return nbrs
}

fn print_highest_total(states:&Vec<State>){
    let mut best = 0;
    for s in states{
        if s.total > best{
            best = s.total;
        }
    }
    println!("Current best is: {}", best);
    /*
    for s in states {
        if s.total == best  && best> 0{
            println!("{:?}", s);
        }
    }*/
}


fn main() {
    println!("Hello, world!");
    let input = fs::read_to_string("input.txt").unwrap();
    let (mut curr_best_states, nodes, pressures, name_to_index) = construct_states(&input);
    iterate_states(&mut curr_best_states, &nodes, &pressures, &name_to_index);
    print!("Final result:");
    print_highest_total(&curr_best_states);
}
