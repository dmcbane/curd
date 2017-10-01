use std::process;

mod args;
mod config;
mod execute;

fn main() {
    let args = args::Args::new().unwrap_or_else(|err| {
        println!("Problem parsing arguments: {}", err);
        process::exit(1);
    });
    let config = config::Config::new(&args.configfile).unwrap_or_else(|err| {
        println!("Problem reading configuration file: {}", err);
        process::exit(2);
    });
    execute::execute_command(&args, &config);
}
