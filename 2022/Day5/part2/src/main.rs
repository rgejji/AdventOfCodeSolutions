//const N:usize = 3;
const N:usize = 9;

fn create_stacks(stacks:&mut[Vec<char>; N]) {
    let mut inits: Vec<&str> = Default::default(); 
    /*let a = "ZN";
    let b = "MCD";
    let c = "P";*/

    let a = "RNFVLJSM";
    let b = "PNDZFJWH";
    let c = "WRCDG";
    let d = "NBS";
    let e = "MZWPCBFN";
    let f = "PRMW";
    let g = "RTNGLSW";
    let h = "QTHFNBV";
    let i = "LMHZNF";
    inits.push(a);
    inits.push(b);
    inits.push(c);
    inits.push(d);
    inits.push(e);
    inits.push(f);
    inits.push(g);
    inits.push(h);
    inits.push(i);

    for (cnt, stack) in stacks.iter_mut().enumerate(){ 
        //let stack = &mut stacks[cnt];
        for curr in inits[cnt].chars(){
            stack.push(curr);
        }
    }
}

fn perform_moves(input:&str, stacks:&mut[Vec<char>;N]) { 
    for line in input.split("\n"){
        let words: Vec<&str> = line.split(" ").collect();
        let num:usize = words[1].parse().unwrap();
        let from:usize= words[3].parse().unwrap();
        let to:usize = words[5].parse().unwrap();

        //println!("Move {} from {} to {}", num, from, to);
        move_cmd(num,from-1,to-1, stacks);
    }
}

//We change the move command to copy to a tmp vector
//Doing this reverses the order
fn move_cmd(num:usize, start:usize, end:usize, stacks:&mut[Vec<char>]){
    let mut tmp:Vec<char> = Default::default();
    for _i in 0..num {
        let val = stacks[start].pop();

        if val.is_none() {
            println!("Error, cannot pop from an empty stack!");
            return;
        }
        tmp.push(val.unwrap());
    }
    for _i in 0..num {
        let val = tmp.pop().unwrap();
        stacks[end].push(val);
    }
}

fn evaluate(stacks:&mut[Vec<char>]){
    for v in stacks{
        let val = v.pop();
        print!("{}", val.unwrap());
    }
    println!("");
}

fn main() {
    //let stacks: vec![Vec<char>; N] = Default::default();
    let mut stacks: [Vec<char>; N] = Default::default();
    create_stacks(&mut stacks);
    perform_moves(INPUT,&mut stacks);
    evaluate(&mut stacks);
    //println!("stacks: {:?}",stacks);
}

const INPUT:&str = "move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";
