use std::collections::HashMap;
use std::env;
use std::fs::File;
use std::io::BufRead;
use std::io::BufReader;
use std::path::Path;
use std::process;

fn main() {
    let arg_vec: Vec<String> = env::args().collect();
    let args = Args::new(&arg_vec)
        .unwrap_or_else(|err| {
            println!("Problem parsing arguments: {}", err);
            process::exit(1);
        });
    let config = Config::new(&args)
        .unwrap_or_else(|err| {
            println!("Problem reading configuration file: {}", err);
            process::exit(2);
        });
}

struct Args {
    configfile: String,
}

impl Args {
    fn new(args: &[String]) -> Result<Args, &'static str> {
        if args.len() < 2 {
            // return Err("not enough arguments");
            return Ok(Args { configfile: String::new() })
        } else {
            return Ok(Args { configfile: args[1].clone() });
        }
    }
}

struct Config {
    configfile: String,
    directories: HashMap<String, String>,
}

impl Config {
    fn new(args: &Args) -> Result<Config, &'static str> {
        let mut configuration_file = get_default_config_filename();
        if args.configfile.len() > 0 {
            configuration_file = args.configfile.clone();
        }
        let dirs = load_configuration(&configuration_file)?;

        //    return Err("not enough arguments");
        Ok(Config { configfile: String::from(configuration_file), directories: dirs })
    }
}

fn get_default_config_filename () -> String {
   let home_dir = if cfg!(windows) {
        env::var("USERPROFILE").unwrap_or(String::from("."))
   } else {
       env::var("HOME").unwrap_or(String::from("."))
   };

   let path = Path::new(&home_dir).join(".curdrc").into_os_string().into_string().unwrap();
   path
}

fn load_configuration(filename: &str) -> Result<HashMap<String, String>, &'static str> {
    let mut result: HashMap<String, String> = HashMap::new();
    if Path::new(filename).exists() {
        let f = File::open(filename).unwrap();
        let file = BufReader::new(&f);

        for line in file.lines() {
            let mut l = line.unwrap();
            let pipe = l.find('|').unwrap_or(l.len());
            let key: String = l.drain(..pipe).collect(); // retrieve the key
            let value: String = l.drain(2..).collect();  // retrieve the value
            result.insert(key, value);
        }
    }
    Ok(result)
}
