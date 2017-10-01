extern crate yaml_rust;

use std::collections::HashMap;
use std::fs::File;
use std::fs::OpenOptions;
use std::io::BufWriter;
use std::io::prelude::*;
use std::io::Result as SIResult;
use std::path::Path;
use self::yaml_rust::{YamlLoader, YamlEmitter};

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
            let mut f = File::open(filename).unwrap();
            let mut contents = String::new();
            f.read_to_string(&mut contents).unwrap();
            let docs = YamlLoader::load_from_str(&contents).unwrap();

            // Multi document support, doc is a yaml::Yaml
            let doc = &docs[0];
            let hsh = doc.as_hash().unwrap();

            for (key, value) in hsh.iter() {
                result.insert(
                    key.as_str().unwrap().to_string(),
                    value.as_str().unwrap().to_string(),
                );
            }
        }
        Ok(result)
    }

    pub fn save_configuration<'a>(&'a self) -> SIResult<()> {
        let file = OpenOptions::new().write(true).truncate(true).open(
            &self.configfile,
        )?;
        let mut writer = BufWriter::new(&file);
        let mut yaml_str = "".to_string();

        // Write the paths to `file`
        for (key, value) in self.paths.iter() {
            let line = format!("{}: {}\n", key, value);
            yaml_str.push_str(&line);
        }
        let docs = YamlLoader::load_from_str(&yaml_str).unwrap();
        let doc = &docs[0];
        let mut out_str = String::new();
        {
            let mut emitter = YamlEmitter::new(&mut out_str);
            emitter.dump(doc).unwrap(); // dump the YAML object to a String
        }
        writer.write_all(out_str.as_bytes())?;
        writer.flush()?;
        Ok(())
    }
}
