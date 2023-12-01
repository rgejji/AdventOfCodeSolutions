use priority_queue::PriorityQueue;

const NUM_ROUNDS:i32 = 20;

enum Op {
	Plus,
	Times
}

enum SameOrU32 {
    U(u32),
    Same
}

struct Monkey {
    items: Vec<u32>,
    operator: Op,
	operator_val: SameOrU32,
    test_div: u32,
	true_monkey: usize,
	false_monkey: usize
}

fn observe_monkeys(input:&str)-> Vec<Monkey> {
	let mut monkeys = Vec::<Monkey>::new();
	let mut line_itr = input.split("\n");

	while let Some(line) = line_itr.next(){
		//get monkey id
        let no_colon = &line[1..line.len()-1];
		let monkey_id  = get_num_from_line(&no_colon, 1, "Error, could not parse monkey id");
		if monkey_id != monkeys.len().try_into().unwrap() {
			println!("Error, monkeys are out of order");
			std::process::exit(1);
		}
		//get items
		let line = line_itr.next().unwrap();
        let item_split:Vec<&str> = line.trim().split(": ").collect();
        //println!("{:?}", item_split);
        
        let items:Vec<u32> = match item_split.get(1) {
			None => {
				Vec::new()
			},
			Some(i) => {
                i.split(", ").map(|s| s.parse().unwrap()).collect::<Vec<u32>>()
            }
		};
	    //get operation
		let line = line_itr.next().unwrap();
        let operator_val = get_op_val(&line, 5, "Error, could not parse operator");

	    let op_terms:Vec<&str> = line.trim().split(" ").collect();
        let operator = match op_terms.get(4) {
			None => {
				println!("Error, could parse operator terms");
			    std::process::exit(1);
			},
			Some(i) => {
				if String::from(*i).eq("*") {Op::Times} else {Op::Plus}
			}
		};
	    //get test
		let line = line_itr.next().unwrap();
        let test_div = get_num_from_line(&line, 3, "Error, could not parse divisor");
        //get true
		let line = line_itr.next().unwrap();
        let true_monkey = usize::try_from(get_num_from_line(&line, 5, "Error, could not parse true monkey")).unwrap();
        //get false
		let line = line_itr.next().unwrap();
        let false_monkey = usize::try_from(get_num_from_line(&line, 5, "Error, could not parse false monkey")).unwrap();
        //make monkey
        let m = Monkey {
            items: items,
            operator: operator,
            operator_val: operator_val,
            test_div: test_div,
            true_monkey: true_monkey,
            false_monkey: false_monkey
        };
        //println!("MONKEY: {:?} {} {} {}", m.items, m.test_div, m.true_monkey, m.false_monkey);
        monkeys.push(m);
        //skip empty line
		line_itr.next();
	}
    monkeys
}

fn monkey_move(monkeys: &mut Vec<Monkey>, counts: &mut Vec<u32> ) {

    for i in 0..monkeys.len() {
        let monkey = &monkeys[i];
        let true_index = monkey.true_monkey;
        let false_index = monkey.false_monkey;
	    let mut true_items = Vec::<u32>::new();
    	let mut false_items = Vec::<u32>::new();

        for item in monkey.items.iter(){
            let term:u32 = match monkey.operator_val {
                SameOrU32::U(i) => i,
                SameOrU32::Same => *item
            };
            let worry = match monkey.operator {
               Op::Plus => *item+term,
               Op::Times => *item*term,
            };
            let new_worry = worry/3;
            if new_worry % monkey.test_div == 0 {
                true_items.push(new_worry);
            } else {
                false_items.push(new_worry);
            }
            //println!("Monkey {} is inspecting  {}", i, item);
            counts[i] += 1;
        }
        monkeys[i].items = Vec::<u32>::new();

        for item in &true_items{
            //println!("Monkey {} is sending {} to {} ", i, *item, true_index);
            monkeys[true_index].items.push(*item);
        }
        for item in &false_items{
            //println!("Monkey {} is sending {} to {}", i, *item, false_index);
            monkeys[false_index].items.push(*item);
        }
    }
}

fn get_num_from_line(line:&str, index:usize, err:&str) -> u32 {
    let line_split: Vec<&str> = line.trim().split(" ").collect();
    match line_split.get(index) {
		None => {
			println!("ERROR: {}", err);
			0
		},
		Some(i) => {
            //println!("{:?}", line_split);
	    	i.parse().expect("Unable to parse value")
		}
	}
}

fn get_op_val(line:&str, index:usize, _err:&str) -> SameOrU32 {
    let line_split: Vec<&str> = line.trim().split(" ").collect();
    match line_split.get(index) {
		None => {
			//println!("{}", err);
			SameOrU32::U(0)
		},
		Some(i) => {
            let old = String::from("old");
            if i.eq(&old) {
                return SameOrU32::Same;
            }
            //println!("{:?}", line_split);
	    	SameOrU32::U(i.parse().expect("Unable to parse value"))
		}
	}
}

fn main() {
    println!("Hello, world!");
    let mut monkeys = observe_monkeys(INPUT);
    let mut counts:Vec<u32> = vec![0; monkeys.len()];
    for _i in 0..NUM_ROUNDS {
        monkey_move(&mut monkeys, &mut counts);
    }
    let mut pq = PriorityQueue::new();
    for (i, val) in counts.iter().enumerate() {
        pq.push(i, val);
    }
    println!("{:?}", counts);
    let a = pq.pop().unwrap();
    let b = pq.pop().unwrap();
    println!("Top two are {} and {} with score {}", a.1, b.1, a.1*b.1);

}
/*
const INPUT:&str = "Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1";
*/
const INPUT:&str ="Monkey 0:
  Starting items: 96, 60, 68, 91, 83, 57, 85
  Operation: new = old * 2
  Test: divisible by 17
    If true: throw to monkey 2
    If false: throw to monkey 5

Monkey 1:
  Starting items: 75, 78, 68, 81, 73, 99
  Operation: new = old + 3
  Test: divisible by 13
    If true: throw to monkey 7
    If false: throw to monkey 4

Monkey 2:
  Starting items: 69, 86, 67, 55, 96, 69, 94, 85
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 6
    If false: throw to monkey 5

Monkey 3:
  Starting items: 88, 75, 74, 98, 80
  Operation: new = old + 5
  Test: divisible by 7
    If true: throw to monkey 7
    If false: throw to monkey 1

Monkey 4:
  Starting items: 82
  Operation: new = old + 8
  Test: divisible by 11
    If true: throw to monkey 0
    If false: throw to monkey 2

Monkey 5:
  Starting items: 72, 92, 92
  Operation: new = old * 5
  Test: divisible by 3
    If true: throw to monkey 6
    If false: throw to monkey 3

Monkey 6:
  Starting items: 74, 61
  Operation: new = old * old
  Test: divisible by 2
    If true: throw to monkey 3
    If false: throw to monkey 1

Monkey 7:
  Starting items: 76, 86, 83, 55
  Operation: new = old + 4
  Test: divisible by 5
    If true: throw to monkey 4
    If false: throw to monkey 0";
