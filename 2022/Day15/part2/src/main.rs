use priority_queue::PriorityQueue;
use core::cmp::Reverse;
use core::cmp::max;

#[derive(Clone, Debug, Eq, Hash, Ord, PartialEq, PartialOrd)]
struct Pos(i64, i64);

const max_val:usize = 4000000;
//const max_val:usize = 20;

fn get_distance(p1:&Pos, p2:&Pos) -> i64 {
    (p1.0-p2.0).abs() + (p1.1-p2.1).abs()
}

fn get_vector(input:&str) -> Vec<(Pos,Pos, i64)> {
    let mut v = vec![];
    for line in input.split("\n"){
        let parts = line.split(" ").collect::<Vec<&str>>();
        let ax = get_int_between(parts[2], "=", ",");
        let ay = get_int_between(parts[3], "=", ":");

        let bx = get_int_between(parts[8], "=", ",");
        let by = get_int_between(parts[9], "=", "");

        let sensor = Pos(ax,ay);
        let beacon = Pos(bx,by);
        let d = get_distance(&sensor, &beacon);
        v.push((sensor, beacon, d));
    }
    v
}

fn found_beacon_v2(v:&Vec<(Pos,Pos, i64)>, query_y:i64) -> bool {
    let mut segments = PriorityQueue::new();
    for (sensor,beacon, max_dist) in v{
        let d = (query_y-sensor.1).abs();
        let sx = sensor.0;
        if d <= *max_dist {
            let start = sx-(max_dist-d);
            let end = sx+(max_dist-d);
            segments.push(Pos(start, end), Reverse(start));
        }
    }
    let mut max_end:i64 = 0;
    while segments.len() > 0{
        let (Pos(start,end), _) = segments.pop().unwrap();
        if start > max_end+1 {
            println!("FOUND MISSING VALUE ({}, {})", max_end+1, query_y);
            println!("SCORE {}", 4000000*(max_end+1)+query_y);
            return true
        }
        max_end=max(max_end, end);
    }
    if max_end < max_val.try_into().unwrap() {
        println!("FOUND MISSING VALUE AT END ({}, {})", max_end+1, query_y);
        println!("SCORE {}", 4000000*(max_end+1)+query_y);
        return true
    }
    return false
}


fn get_int_between(s:&str, start:&str, end:&str) -> i64{
    let after_start = s.split(start).collect::<Vec<&str>>()[1];
    if end == ""{
        return after_start.parse().unwrap()
    }
    after_start.split(end).collect::<Vec<&str>>()[0].parse().unwrap()
}

fn main() {
    println!("Hello, world!");
    let v:Vec<(Pos,Pos, i64)> = get_vector(INPUT);
    println!("READ IN {} VECS", v.len());
    for query in 0..max_val{
        if found_beacon_v2(&v, query.try_into().unwrap()){
            break;
        }
        if query %10000 == 0{
            println!("Query: {}", query);
        }
    }
    println!("DONE");

}

/*
const INPUT:&str ="Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3";
*/

const INPUT:&str = "Sensor at x=2389280, y=2368338: closest beacon is at x=2127703, y=2732666
Sensor at x=1882900, y=3151610: closest beacon is at x=2127703, y=2732666
Sensor at x=2480353, y=3555879: closest beacon is at x=2092670, y=3609041
Sensor at x=93539, y=965767: closest beacon is at x=501559, y=361502
Sensor at x=357769, y=2291291: closest beacon is at x=262473, y=2000000
Sensor at x=2237908, y=1893142: closest beacon is at x=2127703, y=2732666
Sensor at x=2331355, y=3906306: closest beacon is at x=2092670, y=3609041
Sensor at x=3919787, y=2021847: closest beacon is at x=2795763, y=2589706
Sensor at x=3501238, y=3327244: closest beacon is at x=3562181, y=3408594
Sensor at x=1695968, y=2581703: closest beacon is at x=2127703, y=2732666
Sensor at x=3545913, y=3356504: closest beacon is at x=3562181, y=3408594
Sensor at x=1182450, y=1405295: closest beacon is at x=262473, y=2000000
Sensor at x=3067566, y=3753120: closest beacon is at x=3562181, y=3408594
Sensor at x=1835569, y=3983183: closest beacon is at x=2092670, y=3609041
Sensor at x=127716, y=2464105: closest beacon is at x=262473, y=2000000
Sensor at x=3065608, y=3010074: closest beacon is at x=2795763, y=2589706
Sensor at x=2690430, y=2693094: closest beacon is at x=2795763, y=2589706
Sensor at x=2051508, y=3785175: closest beacon is at x=2092670, y=3609041
Sensor at x=2377394, y=3043562: closest beacon is at x=2127703, y=2732666
Sensor at x=1377653, y=37024: closest beacon is at x=501559, y=361502
Sensor at x=2758174, y=2627042: closest beacon is at x=2795763, y=2589706
Sensor at x=1968468, y=2665146: closest beacon is at x=2127703, y=2732666
Sensor at x=3993311, y=3779031: closest beacon is at x=3562181, y=3408594
Sensor at x=159792, y=1923149: closest beacon is at x=262473, y=2000000
Sensor at x=724679, y=3489022: closest beacon is at x=2092670, y=3609041
Sensor at x=720259, y=121267: closest beacon is at x=501559, y=361502
Sensor at x=6, y=46894: closest beacon is at x=501559, y=361502
Sensor at x=21501, y=2098549: closest beacon is at x=262473, y=2000000
Sensor at x=2974083, y=551886: closest beacon is at x=4271266, y=-98555";
