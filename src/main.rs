use std::env;
use std::process;

mod args;
mod config;

fn main() {
    let arg_vec: Vec<String> = env::args().collect();
    let args = args::Args::new(&arg_vec).unwrap_or_else(|err| {
        println!("Problem parsing arguments: {}", err);
        process::exit(1);
    });
    let config = config::Config::new(&args).unwrap_or_else(|err| {
        println!("Problem reading configuration file: {}", err);
        process::exit(2);
    });
    println!("config.configfile: {}", config.configfile);
    println!("config.directories: {:?}", config.directories);
}
