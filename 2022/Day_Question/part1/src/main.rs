use pathfinding::prelude::dijkstra;
use graph::prelude::*;
use std::collections::HashMap;
use std::collections::HashSet;
//use core::char::from_digit;
use std::hash::{Hash, Hasher};

const penalty:i32 = 500;
const ASCII_A:u32 = 65;

struct Flows{
    graph: DirectedCsrGraph<u32>,
    rates: HashMap<u32, i32>
}

#[derive(Clone, Debug, Eq, PartialEq)]
struct State{
    pos: u32,
    valve_state: HashSet<u32>,
    time: u32,
    pressure_delta: i32,
    pressure_total: i32
}

impl Hash for State {
    fn hash<H: Hasher>(&self, s: &mut H) {
        self.pos.hash(s);
        self.time.hash(s);
        self.pressure_delta.hash(s);
        self.pressure_total.hash(s);
    }
}

impl State { 
  fn successors(&self, flows:&Flows) -> Vec<(State, i32)> {
    let loc = self.pos.clone();
    let neighbors = flows.graph.out_neighbors(self.pos).as_slice();
    let mut valid = vec![];
    //Check if opening valve is beneficial
    if !self.valve_state.contains(&loc) && flows.rates[&loc] > 0 { 
        let mut new_valve_state = self.valve_state.clone();
        new_valve_state.insert(loc);
        let open_valve = State {
            pos: loc.clone(),
            valve_state: new_valve_state,
            time: self.time+1,
            pressure_delta: self.pressure_delta+flows.rates[&loc],
            pressure_total: self.pressure_total+self.pressure_delta
        };
        let edge_weight:i32 = penalty - open_valve.pressure_delta;
        //println!("Exploring opening valve at state:\nTime: {}\nPos: {}\nRel Pressure {}", open_valve.time, itos(open_valve.pos), open_valve.pressure_delta); 
        valid.push((open_valve, edge_weight));
    }

    //Otherwise move somewhere
    for neighbor in neighbors {
        let new_valve_state = self.valve_state.clone();
        let move_valve = State {
            pos: neighbor.clone(),
            valve_state: new_valve_state,
            time: self.time+1,
            pressure_delta: self.pressure_delta,
            pressure_total: self.pressure_total+self.pressure_delta
        };
        let edge_weight:i32 = penalty - move_valve.pressure_delta;

        valid.push((move_valve, edge_weight));
    }
    valid
  }
}

fn stoi(s:String) -> u32{
    let tmp:Vec<&u8> = s.as_bytes().iter().collect();
    u32::try_from(*tmp[0]).unwrap()-ASCII_A+26*(u32::try_from(*tmp[1]).unwrap()-ASCII_A)
}


fn itos(i:u32) -> String {
    let first = i%26;
    let second = (i-first)/26;
    //format!("{}", i)
    format!("{}{}", (first+ASCII_A) as u8 as char, (second+ASCII_A) as u8 as char).to_string() 
}

fn construct_flows(input:&str) -> Flows {
    let mut edge_list = vec![];
    let mut rates = HashMap::new();
    for line in input.split("\n") {
        let name:String = String::from(line.split(" ").collect::<Vec<&str>>()[1]);
        let id = stoi(name);
        let rate:i32 = line.split("=").collect::<Vec<&str>>()[1].split(";").collect::<Vec<&str>>()[0].parse().unwrap();
        let edge_strs = 
            match line.split("valves ").collect::<Vec<&str>>().get(1).clone(){
                Some(val) => val,
                None => {
                    let tmp_b = line.split("valve ").collect::<Vec<&str>>();
                    tmp_b.get(1).unwrap().clone()
                }
            };    
        for edge_end in edge_strs.split(", "){
            edge_list.push((id, stoi(edge_end.to_string())));
        }
        rates.insert(id, rate);
    }

    let graph: DirectedCsrGraph<u32> = GraphBuilder::new()
        .csr_layout(CsrLayout::Sorted)
        .edges(edge_list)
        .build();
    Flows {
        graph: graph,
        rates: rates
    }
}

fn main() {
    println!("Hello, world!");
    let f = construct_flows(INPUT);

    let begin:State = State {
        pos: stoi("AA".to_string()),
        valve_state: HashSet::new(),
        time: 0,
        pressure_delta: 0,
        pressure_total: 0,
    };
    let result = dijkstra(&begin, |p| p.successors(&f), |p| p.time == 30).expect("Error, no short path found");
    let last = result.0.len()-1;
    println!("Have total pressure released {}", result.0[last].pressure_total)
}

/*
const INPUT: &str = "Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II";
*/

const INPUT: &str = "Valve WT has flow rate=0; tunnels lead to valves BD, FQ
Valve UG has flow rate=0; tunnels lead to valves FQ, YB
Valve FN has flow rate=0; tunnels lead to valves TV, GA
Valve RU has flow rate=11; tunnels lead to valves YZ, QS, BL, BT, WJ
Valve RH has flow rate=0; tunnels lead to valves AS, II
Valve FL has flow rate=0; tunnels lead to valves HR, PQ
Valve KQ has flow rate=18; tunnels lead to valves FR, BN
Valve PM has flow rate=25; tunnels lead to valves YZ, FR
Valve RQ has flow rate=0; tunnels lead to valves FQ, MW
Valve BL has flow rate=0; tunnels lead to valves RU, IR
Valve FF has flow rate=0; tunnels lead to valves QS, ED
Valve KP has flow rate=0; tunnels lead to valves QM, MA
Valve YB has flow rate=0; tunnels lead to valves UG, HR
Valve TV has flow rate=17; tunnels lead to valves BD, MT, FN
Valve HY has flow rate=0; tunnels lead to valves DW, IU
Valve KF has flow rate=0; tunnels lead to valves AA, HR
Valve YC has flow rate=0; tunnels lead to valves II, MA
Valve EE has flow rate=0; tunnels lead to valves AA, CD
Valve ED has flow rate=9; tunnels lead to valves HG, FF
Valve SA has flow rate=0; tunnels lead to valves MW, LS
Valve II has flow rate=20; tunnels lead to valves YC, CY, QP, RH
Valve BN has flow rate=0; tunnels lead to valves BT, KQ
Valve MO has flow rate=0; tunnels lead to valves XO, VI
Valve YZ has flow rate=0; tunnels lead to valves RU, PM
Valve WJ has flow rate=0; tunnels lead to valves RU, QP
Valve AW has flow rate=0; tunnels lead to valves HR, DW
Valve MJ has flow rate=0; tunnels lead to valves BP, AA
Valve DW has flow rate=4; tunnels lead to valves AU, CB, HY, GL, AW
Valve QM has flow rate=0; tunnels lead to valves KP, FQ
Valve LF has flow rate=5; tunnels lead to valves LS, QN, AU, BP, ZY
Valve QS has flow rate=0; tunnels lead to valves FF, RU
Valve BT has flow rate=0; tunnels lead to valves BN, RU
Valve VI has flow rate=22; tunnel leads to valve MO
Valve LS has flow rate=0; tunnels lead to valves LF, SA
Valve QD has flow rate=0; tunnels lead to valves HR, ZY
Valve HG has flow rate=0; tunnels lead to valves AS, ED
Valve BD has flow rate=0; tunnels lead to valves WT, TV
Valve CD has flow rate=0; tunnels lead to valves EE, MW
Valve QP has flow rate=0; tunnels lead to valves II, WJ
Valve MW has flow rate=7; tunnels lead to valves PQ, SA, CB, CD, RQ
Valve AU has flow rate=0; tunnels lead to valves DW, LF
Valve RR has flow rate=0; tunnels lead to valves AS, MA
Valve GA has flow rate=0; tunnels lead to valves FN, MA
Valve MT has flow rate=0; tunnels lead to valves CY, TV
Valve HR has flow rate=14; tunnels lead to valves KF, YB, QD, AW, FL
Valve AS has flow rate=16; tunnels lead to valves RR, RH, HG, IR
Valve CY has flow rate=0; tunnels lead to valves MT, II
Valve AA has flow rate=0; tunnels lead to valves OX, KF, GL, MJ, EE
Valve IU has flow rate=0; tunnels lead to valves XO, HY
Valve XO has flow rate=23; tunnels lead to valves IU, MO
Valve FR has flow rate=0; tunnels lead to valves KQ, PM
Valve CB has flow rate=0; tunnels lead to valves MW, DW
Valve ZY has flow rate=0; tunnels lead to valves QD, LF
Valve BP has flow rate=0; tunnels lead to valves LF, MJ
Valve QN has flow rate=0; tunnels lead to valves LF, FQ
Valve IR has flow rate=0; tunnels lead to valves AS, BL
Valve PQ has flow rate=0; tunnels lead to valves FL, MW
Valve GL has flow rate=0; tunnels lead to valves AA, DW
Valve OX has flow rate=0; tunnels lead to valves MA, AA
Valve MA has flow rate=10; tunnels lead to valves RR, YC, GA, OX, KP
Valve FQ has flow rate=12; tunnels lead to valves QN, WT, UG, RQ, QM";


