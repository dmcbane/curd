use std::collections::HashMap;
use std::fs::File;
use std::path::Path;
use std::io::BufRead;
use std::io::BufReader;

pub struct Config {
    pub configfile: String,
    pub paths: HashMap<String, String>,
}

impl Config {
    pub fn new(args: &::args::Args) -> Result<Config, &'static str> {
        let configuration_file = args.configfile.clone();
        let kv = Config::load_configuration(&configuration_file)?;

        //    return Err("not enough arguments");
        Ok(Config {
            configfile: String::from(configuration_file),
            paths: kv,
        })
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
                let value: String = l.drain(2..).collect(); // retrieve the value
                result.insert(key, value);
            }
        }
        Ok(result)
    }
}
