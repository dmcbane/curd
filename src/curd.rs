use std::collections::HashMap;
use std::env;
use std::path::Path;

pub fn execute_command(args: &::args::Args, config: &::config::Config) {
    if args.read {
        // read the value from paths
        if config.paths.contains_key(&args.keyword) {
            println!("{}", config.paths.get(&args.keyword).unwrap());
        }
    } else if args.save {
        // save the value to the paths
        let mut kv: HashMap<String, String> = HashMap::new();
        for (key, value) in config.paths.iter() {
            kv.insert(key.clone(), value.clone());
        };
        kv.insert(args.keyword.clone(), get_current_dir());
        let _ = ::config::Config {
            configfile: config.configfile.clone(),
            paths: kv,
        }.save_configuration();
    } else if args.remove {
        // remove the value from paths
        let mut kv: HashMap<String, String> = HashMap::new();
        for (key, value) in config.paths.iter() {
            if *key != args.keyword {
                kv.insert(key.clone(), value.clone());
            }
        };
        let _ = ::config::Config {
            configfile: config.configfile.clone(),
            paths: kv,
        }.save_configuration();
    } else if args.list {
        // list defined paths
        let mut keys: Vec<&String> = config.paths.keys().collect();
        keys.sort();
        for key in keys.iter_mut() {
            println!("{} - {}", key, config.paths.get(&key.clone()).unwrap());
        }
    } else if args.clean {
        // remove paths that don't exist
        let mut kv: HashMap<String, String> = HashMap::new();
        for (key, value) in config.paths.iter() {
            if path_exists(value) {
                kv.insert(key.clone(), value.clone());
            }
        };
        let _ = ::config::Config {
            configfile: config.configfile.clone(),
            paths: kv,
        }.save_configuration();
    }
}

fn get_current_dir() -> String {
    env::current_dir().unwrap().display().to_string()
}

fn path_exists(path: &String) -> bool {
    Path::new(path).exists()
}
