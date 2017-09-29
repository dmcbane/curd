use std::process;

mod args;
mod config;
mod curd;

fn main() {
    let args = args::Args::new().unwrap_or_else(|err| {
                                                    println!("Problem parsing arguments: {}", err);
                                                    process::exit(1);
                                                });
    let config = config::Config::new(&args.configfile).unwrap_or_else(|err| {
        println!("Problem reading configuration file: {}", err);
        process::exit(2);
    });
    curd::execute_command(&args, &config);
}
