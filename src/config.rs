use std::collections::HashMap;
use std::fs::File;
use std::io::BufRead;
use std::io::BufReader;
use std::io::prelude::*;
use std::io::Result as SIResult;
use std::path::Path;

pub struct Config {
    pub configfile: String,
    pub paths: HashMap<String, String>,
}

impl Config {
    pub fn new(configuration_file: &String) -> Result<Config, &'static str> {
        let kv = Config::load_configuration(&configuration_file)?;

        //    return Err("not enough arguments");
        Ok(Config {
            configfile: configuration_file.clone(),
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

    pub fn save_configuration<'a>(&'a self) -> SIResult<()> {
        let path = Path::new(&self.configfile);

        // Open a file in write-only mode, returns `io::Result<File>`
        let mut file = File::create(&path)?;
        // Write the paths to `file`
        for (key, value) in self.paths.iter() {
            let line = format!("{}| {}\n", key, value);
            file.write_all(line.as_bytes())?;
        };
        Ok(())
    }
}
