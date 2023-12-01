use std::collections::HashMap;

#[derive(Clone, Debug, Default)]
struct Robot {
	ore:u64,
	clay:u64,
	obsidian:u64,
	geode:u64,
	ore_robot:u64,
	clay_robot:u64,
	obsidian_robot:u64,
	geode_robot:u64,
}

impl Robot {
	fn get_hash(&self) -> String{
		format!("{}_{}_{}_{}", self.ore_robot, self.clay_robot, self.obsidian_robot, self.geode_robot).to_string()
	}
	fn collect(&mut self, max_geodes: &mut u64){
		self.ore += self.ore_robot;
		self.clay += self.clay_robot;
		self.obsidian += self.obsidian_robot;
		self.geode += self.geode_robot;
		if self.geode > *max_geodes{
			*max_geodes = self.geode;
			//println!("We have exceeded the max geodes with a robot {:?}", self);
		}

	}

}

#[derive(Clone, Debug)]
struct Blueprint {
	id: u64,
	ore_robot_ore: u64,
	clay_robot_ore: u64,
	obsidian_robot_ore: u64,
	obsidian_robot_clay: u64,
	geode_robot_ore: u64,
	geode_robot_obsidian: u64,
}

fn get_int_between(s:&str, start:&str, end:&str) -> u64{
	if start == ""{
		if end == ""{
			return s.parse().unwrap()
		}
		return s.split(end).collect::<Vec<&str>>()[0].parse().unwrap()
	}

	let after_start = s.split(start).collect::<Vec<&str>>()[1];
	if end == ""{
		return after_start.parse().unwrap()
	}
	after_start.split(end).collect::<Vec<&str>>()[0].parse().unwrap()
}

fn read_blueprints(input:&str) -> Vec<Blueprint> {
	let mut blueprints = vec![];

	for (id, line) in input.split("\n").enumerate(){
		let terms = line.split(" ").collect::<Vec<&str>>();
		let id = get_int_between(terms[1],"",":");
		let ore_robot_ore = get_int_between(terms[6],"","");
		let clay_robot_ore = get_int_between(terms[12],"","");
		let obsidian_robot_ore = get_int_between(terms[18],"","");
		let obsidian_robot_clay = get_int_between(terms[21],"","");
		let geode_robot_ore = get_int_between(terms[27],"","");
		let geode_robot_obsidian = get_int_between(terms[30],"","");


		blueprints.push(Blueprint {
			id: id,
			ore_robot_ore: ore_robot_ore,
			clay_robot_ore: clay_robot_ore,
			obsidian_robot_ore: obsidian_robot_ore,
			obsidian_robot_clay: obsidian_robot_clay,
			geode_robot_ore: geode_robot_ore,
			geode_robot_obsidian: geode_robot_obsidian,
		});
	}
	blueprints
}

fn ignore_or_replace_hash(map:&mut HashMap<String, Vec<Robot>>, rob:Robot) {
	let hash = rob.get_hash();
	if !map.contains_key(&hash){
		map.insert(hash, vec![rob]);
		return
	}

	let mut robot_set = map.get_mut(&hash).unwrap();
	for mut alt_rob in robot_set.iter_mut(){
		if rob.ore <= alt_rob.ore && rob.clay <= alt_rob.clay && rob.obsidian <= alt_rob.obsidian && rob.geode <= alt_rob.geode {
			return
		}
		if rob.ore >= alt_rob.ore && rob.clay >= alt_rob.clay && rob.obsidian >= alt_rob.obsidian && rob.geode >= alt_rob.geode {
			alt_rob.ore = rob.ore;
			alt_rob.clay = rob.clay;
			alt_rob.obsidian = rob.obsidian;
			alt_rob.geode = rob.geode;
			return
		}
	}
	robot_set.push(rob);
}


fn find_best(blue:&Blueprint,num_min:usize) -> u64{
	let mut max_geodes = 0u64;

	//Initialize map with first robot
	let mut prev_map:HashMap<String, Vec<Robot>> = HashMap::new();
	let mut starter_rob:Robot =  Default::default();
	starter_rob.ore_robot += 1;
	prev_map.insert(starter_rob.get_hash(), vec![starter_rob]);

	for i in 0..num_min{
		println!("Minute {} is passing for blueprint {} and we have {} hashes and max_geodes {}", i, blue.id, prev_map.len(), max_geodes);
		let mut curr_map = HashMap::new();
		for robot_set in prev_map.values_mut(){
			while let Some(mut robot) = robot_set.pop(){

				//skip pathetic robots that cannot catch up even if they buy a geode robot evey round.
				let tmp:u64 = (num_min-i).try_into().unwrap();
				if robot.geode + tmp*robot.geode_robot+tmp*(tmp+1)/2 < max_geodes{
					continue;
				}

				//try all the actions
				if robot.ore >= blue.ore_robot_ore{
					let mut new_bot = robot.clone();
					new_bot.ore -= blue.ore_robot_ore;
					new_bot.collect(&mut max_geodes);
					new_bot.ore_robot +=1;
					ignore_or_replace_hash(&mut curr_map, new_bot);
				}
				if robot.ore >= blue.clay_robot_ore{
					let mut new_bot = robot.clone();
					new_bot.ore -= blue.clay_robot_ore;
					new_bot.collect(&mut max_geodes);
					new_bot.clay_robot +=1;
					ignore_or_replace_hash(&mut curr_map, new_bot);
				}
				if robot.ore >= blue.obsidian_robot_ore && robot.clay >= blue.obsidian_robot_clay {
					let mut new_bot = robot.clone();
					new_bot.ore -= blue.obsidian_robot_ore;
					new_bot.clay -= blue.obsidian_robot_clay;
					new_bot.collect(&mut max_geodes);
					new_bot.obsidian_robot +=1;
					ignore_or_replace_hash(&mut curr_map, new_bot);
				}
				if robot.ore >= blue.geode_robot_ore && robot.obsidian >= blue.geode_robot_obsidian {
					let mut new_bot = robot.clone();
					new_bot.ore -= blue.geode_robot_ore;
					new_bot.obsidian -= blue.geode_robot_obsidian;
					new_bot.collect(&mut max_geodes);
					new_bot.geode_robot +=1;
					ignore_or_replace_hash(&mut curr_map, new_bot);
				}
				robot.collect(&mut max_geodes);
				ignore_or_replace_hash(&mut curr_map, robot);	
			}
		}
		prev_map = curr_map;
	}
	println!("Finished Blueprint! Maximum number of geodes encountered is {}", max_geodes);
	max_geodes
}


fn main() {
    println!("Hello, world!");
	let blueprints = read_blueprints(INPUT);
	println!("{:?}", blueprints);
	let mut score:u64 = 0;
	for (i,blueprint) in blueprints.into_iter().enumerate(){
		let time:u64 = (1+i).try_into().unwrap();
		score += time*find_best(&blueprint, 24usize);
	}
	println!("Final score is {}", score);
}

/*const INPUT:&str="Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.";
*/

const INPUT:&str="Blueprint 1: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 18 clay. Each geode robot costs 3 ore and 8 obsidian.
Blueprint 2: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 18 clay. Each geode robot costs 4 ore and 20 obsidian.
Blueprint 3: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 18 clay. Each geode robot costs 4 ore and 12 obsidian.
Blueprint 4: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 20 clay. Each geode robot costs 2 ore and 8 obsidian.
Blueprint 5: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 2 ore and 14 clay. Each geode robot costs 3 ore and 8 obsidian.
Blueprint 6: Each ore robot costs 3 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 17 clay. Each geode robot costs 4 ore and 8 obsidian.
Blueprint 7: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 4 ore and 6 clay. Each geode robot costs 3 ore and 11 obsidian.
Blueprint 8: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 20 clay. Each geode robot costs 4 ore and 18 obsidian.
Blueprint 9: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 16 clay. Each geode robot costs 3 ore and 15 obsidian.
Blueprint 10: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 7 clay. Each geode robot costs 4 ore and 11 obsidian.
Blueprint 11: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 9 clay. Each geode robot costs 3 ore and 19 obsidian.
Blueprint 12: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 2 ore and 17 clay. Each geode robot costs 3 ore and 16 obsidian.
Blueprint 13: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 5 clay. Each geode robot costs 3 ore and 12 obsidian.
Blueprint 14: Each ore robot costs 3 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 2 ore and 12 obsidian.
Blueprint 15: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 7 clay. Each geode robot costs 3 ore and 20 obsidian.
Blueprint 16: Each ore robot costs 3 ore. Each clay robot costs 3 ore. Each obsidian robot costs 2 ore and 11 clay. Each geode robot costs 2 ore and 19 obsidian.
Blueprint 17: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 10 clay. Each geode robot costs 3 ore and 14 obsidian.
Blueprint 18: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 12 clay. Each geode robot costs 3 ore and 17 obsidian.
Blueprint 19: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 7 clay. Each geode robot costs 2 ore and 19 obsidian.
Blueprint 20: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 2 ore and 17 clay. Each geode robot costs 3 ore and 19 obsidian.
Blueprint 21: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 2 ore and 14 clay. Each geode robot costs 4 ore and 11 obsidian.
Blueprint 22: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 4 ore and 15 clay. Each geode robot costs 4 ore and 9 obsidian.
Blueprint 23: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 9 clay. Each geode robot costs 3 ore and 7 obsidian.
Blueprint 24: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 15 clay. Each geode robot costs 4 ore and 17 obsidian.
Blueprint 25: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 9 clay. Each geode robot costs 2 ore and 20 obsidian.
Blueprint 26: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 20 clay. Each geode robot costs 3 ore and 14 obsidian.
Blueprint 27: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 15 clay. Each geode robot costs 3 ore and 16 obsidian.
Blueprint 28: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 20 obsidian.
Blueprint 29: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 14 clay. Each geode robot costs 3 ore and 14 obsidian.
Blueprint 30: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 4 ore and 18 clay. Each geode robot costs 3 ore and 13 obsidian.";
