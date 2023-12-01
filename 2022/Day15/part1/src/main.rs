use std::collections::HashSet;

#[derive(Clone, Debug, Eq, Hash, Ord, PartialEq, PartialOrd)]
struct Pos(i32, i32);

fn get_distance(p1:&Pos, p2:&Pos) -> i32 {
    (p1.0-p2.0).abs() + (p1.1-p2.1).abs()
}

fn get_no_beacons(input:&str, query_y:i32) -> i32 {
    let mut no_sensors = HashSet::new();
    for line in input.split("\n"){
        let parts = line.split(" ").collect::<Vec<&str>>();
        let ax = get_int_between(parts[2], "=", ",");
        let ay = get_int_between(parts[3], "=", ":");

        let bx = get_int_between(parts[8], "=", ",");
        let by = get_int_between(parts[9], "=", "");

        let sensor = Pos(ax,ay);
        let beacon = Pos(bx,by);
        
        if sensor.1 == query_y {
            println!("Warning: Found a sensor in the beacon counting line. Occlusion rules have not yet been stated.Ignoring since sensor will have a min beacon distance");
        }

        let max_dist = get_distance(&sensor, &beacon);
        let d = (query_y-sensor.1).abs();
        if d > max_dist {
            continue
        }

        //println!("EXAMINING: {:?} {} {}", sensor, max_dist, d);
        for i in 0..(max_dist+1-d){
            if beacon.0 != sensor.0+i || beacon.1 != query_y {
                //println!("({},{}),({},{}) is adding {},{}", sensor.0, sensor.1, beacon.0, beacon.1, sensor.0+i, query_y);
                no_sensors.insert(Pos(sensor.0+i, query_y));
            } 
            if beacon.0 != sensor.0-i || beacon.1 != query_y {
                //println!("Adding {},{}", sensor.0-i, query_y);
                no_sensors.insert(Pos(sensor.0-i, query_y));
            }

        }
        
    }
    no_sensors.len().try_into().unwrap()
}

fn get_int_between(s:&str, start:&str, end:&str) -> i32{
    let after_start = s.split(start).collect::<Vec<&str>>()[1];
    if end == ""{
        return after_start.parse().unwrap()
    }
    after_start.split(end).collect::<Vec<&str>>()[0].parse().unwrap()
}

fn main() {
    //let query = 10;
    let query = 2000000;
    println!("Hello, world!");
    let val = get_no_beacons(INPUT, query);
    println!("{}",val);

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
