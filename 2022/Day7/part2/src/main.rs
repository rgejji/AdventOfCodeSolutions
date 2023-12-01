use std::fs;

//As instructed from the blog post
//https://dev.to/deciduously/no-more-tears-no-more-knots-arena-allocated-trees-in-rust-44k6#:~:text=One%20such%20category%20is%20tree,Rust%20hates%20that
#[derive(Debug, Default)]
struct ArenaTree<T> 
where
    T: PartialEq
{
    arena: Vec<Node<T>>,
}
impl<T> ArenaTree<T>
where
    T: PartialEq
{
    fn new_node(&mut self, val: T) -> usize {
        //check if node exists, if so return the idx
        //for node in &self.arena {
        //    if node.val == val {
        //        return node.idx;
        //    }
        //}
        //Otherwise, add new node
        let idx = self.arena.len();
        self.arena.push(Node::new(idx,val));
        idx
    }
}

#[derive(Debug)]
struct Node<T>
where
    T: PartialEq
{
    idx: usize,
    val: T,
    parent: Option<usize>,
    children: Vec<usize>
}

impl<T> Node<T>
where
    T: PartialEq
{
    fn new(idx: usize, val:T) -> Self {
        Self{
            idx,
            val,
            parent: None,
            children: vec![],
        }
    }
    fn child(&self, arena:&ArenaTree<FileObject>, name:String) -> usize{
        for c in &self.children{
            if arena.arena[*c].val.name.eq(&name){
                return *c;
            }
        }
        println!("Error: Could not find child with name {}", name);
        //should really use an option here
        return 0;
    }
}


#[derive(Debug, PartialEq)]
struct FileObject{
    file_type: BaseObject,
    size: i64,
    name: String,
}


#[derive(Debug, Default, PartialEq)]
enum BaseObject {
    Dir,
    #[default]
    File,
}

fn insert(arena:&mut ArenaTree<FileObject>, parent_idx:usize, val:FileObject) {
    let idx = arena.new_node(val);
    arena.arena[idx].parent = Some(parent_idx);
    arena.arena[parent_idx].children.push(idx);
}


fn read_input(input:&str) ->  ArenaTree<FileObject> {
    let mut arena = ArenaTree {
        arena: vec![],
    };
    let mut curr_idx = arena.new_node( FileObject {
        file_type: BaseObject::Dir,
        size: -1,
        name: "/".to_string(),
    });
    arena.arena[curr_idx].parent = Some(curr_idx);
        
    for (i,line) in input.split("\n").enumerate(){
        //skip first two lines since they we start in root object which is a directory
        if i<2  || line.eq("") {
            continue;
        }
        let parts:Vec<&str> = line.split(" ").collect();
        match parts[0] {
            "dir" => {
                let size:i64 = -1;
                let name = parts[1].clone();
                insert(&mut arena, curr_idx, FileObject {
                    file_type: BaseObject::Dir,
                    size: size,
                    name: name.to_string(),
                });
            },
            "$" => {
                match parts[1]{
                    "cd" => {
                        match parts[2] {
                            ".." => {
                                curr_idx = arena.arena[curr_idx].parent.unwrap();
                            },
                            _ => {curr_idx = arena.arena[curr_idx].child(&arena, parts[2].to_string())},
                        }
                    },
                    _ => {},
                }
            }, 
            _ => {
                let size:i64 = parts[0].parse().unwrap();
                let name = parts[1].clone();
                insert(&mut arena, curr_idx, FileObject {
                    file_type: BaseObject::File,
                    size: size,
                    name: name.to_string(),
                });
            }
        }

    }
    arena
}

fn calculate_sizes(arena: &mut ArenaTree<FileObject>, idx: usize) -> i64 {
    match arena.arena[idx].val.file_type{
        BaseObject::File => { return arena.arena[idx].val.size},
        BaseObject::Dir => { 
            let mut sum = 0;
            let tmp = arena.arena[idx].children.clone();
            for child in tmp {
                sum += calculate_sizes(arena, child);
            }
            println!("Calcualted size {}", sum);
            arena.arena[idx].val.size = sum;
            return sum
        },
    }
}

/*
fn sum_of_sizes(arena: &ArenaTree<FileObject>, idx: usize, filter_threshold:i64) -> i64{
    match arena.arena[idx].val.file_type{
        BaseObject::File => { return 0 },
        BaseObject::Dir => {
            let mut sum = 0;

            let tmp = arena.arena[idx].children.clone();
            for child in tmp{
                sum += sum_of_sizes(arena, child, filter_threshold);
            }
            let obj = &arena.arena[idx];
            if obj.val.size <0 {
                println!("ERROR: Encountered a negative dir size");
            }

            if obj.val.size < filter_threshold{
                sum += obj.val.size;
            }
            return sum
        },
    }
}*/


fn get_smallest_sufficient_dir_size(arena: &ArenaTree<FileObject>, idx: usize, filter_threshold:i64) -> i64{
    match arena.arena[idx].val.file_type{
        BaseObject::File => { return 99999999999 },
        BaseObject::Dir => {
            let mut min_suff = 99999999999;
            let tmp = arena.arena[idx].children.clone();
            for child in tmp{
                let val = get_smallest_sufficient_dir_size(arena, child, filter_threshold);
                if val < min_suff {
                    min_suff = val;
                }
            }
            let obj = &arena.arena[idx];
            if obj.val.size >= filter_threshold && obj.val.size < min_suff{
                min_suff = obj.val.size
            }
            return min_suff
        },
    }
}


fn main() {
    let input = fs::read_to_string("input.txt").unwrap();
    let mut arena = read_input(&input);
    calculate_sizes(&mut arena,0);
    let used_size = arena.arena[0].val.size;
    println!("Used size is {}", used_size);
    let unused_size = 70000000-used_size;
    println!("Unused size is {}", unused_size);
    let score = get_smallest_sufficient_dir_size(&mut arena,0, 30000000-unused_size);
        
    println!("Hello, world!");
    println!("Final score is {}", score);
}
