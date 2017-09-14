extern crate clap;

use std::env;
use std::path::Path;
use self::clap::{App, Arg, ArgGroup};

pub struct Args {
    pub configfile: String,
    pub keyword: String,
    pub list: bool,
    pub clean: bool,
    pub remove: bool,
    pub save: bool,
    pub read: bool,
}

impl Args {
    pub fn new(args: &[String]) -> Result<Args, &'static str> {
        let matches = App::new("CURD - Change to a User's Recurring Directory")
                          .version("1.0.0")
                          .author("H. Dale McBane<h.dale.mcbane@gmail.com>")
                          .about("Allows a user to save and return to paths that they visit often.")
                          .arg(Arg::with_name("keyword")
                               .help("Display the path associated with keyword."))
                          .arg(Arg::with_name("config")
                               .short("c")
                               .long("config")
                               .value_name("FILE")
                               .help("Select a configuration file to use instead of the default (~/.curdrc).")
                               .takes_value(true))
                          .arg(Arg::with_name("list")
                               .short("l")
                               .long("list")
                               .help("List all of the paths saved in the configuration file."))
                          .arg(Arg::with_name("clean")
                               .short("n")
                               .long("clean")
                               .help("Cleanup entries for paths that don't exist."))
                          .arg(Arg::with_name("remove")
                               .short("r")
                               .long("remove")
                               .value_name("KEYWORD")
                               .help("Remove the path specified by the keyword or the default path from the configuration file.")
                               .takes_value(true))
                          .arg(Arg::with_name("save")
                               .short("s")
                               .long("save")
                               .value_name("KEYWORD")
                               .help("Save the current directory to the specified keyword or the default.")
                               .takes_value(true))
                          .group(ArgGroup::with_name("")
                                 .args(&["keyword", "clean", "list", "remove", "save"]))
                          .after_help("The 'keyword' argument, the 'clean' or 'list' flags, nor the 'remove' or 'save' options can be used together.  The 'config' option can be combined with all other flags, options, and args.")
                          .get_matches_from(args);

        // Gets a value for config if supplied by user, or defaults to "~/.curdrc"
        let default_config = Args::get_default_config_filename();
        let config = (matches.value_of("config").unwrap_or(&default_config)).to_string();
        println!("Value for config: {}", config);

        let keyword = if matches.is_present("keyword") {
            matches.value_of("keyword").unwrap().to_string()
        } else if matches.is_present("remove") {
            matches.value_of("remove").unwrap().to_string()
        } else if matches.is_present("save") {
            matches.value_of("save").unwrap().to_string()
        } else {
            "".to_string()
        };
        println!("Value for keyword: {}", keyword);

        let read: bool = matches.is_present("keyword");
        println!("Value for read: {}", read);
        let clean = matches.is_present("clean");
        println!("Value for clean: {}", clean);
        let list = matches.is_present("list");
        println!("Value for list: {}", list);
        let remove = matches.is_present("remove");
        println!("Value for remove: {}", remove);
        let save = matches.is_present("save");
        println!("Value for save: {}", save);

        return Ok(Args {
            configfile: config,
            keyword: keyword,
            list: list,
            clean: clean,
            remove: remove,
            save: save,
            read: read,
        });
    }

    fn get_default_config_filename() -> String {
        let home_dir = if cfg!(windows) {
            env::var("USERPROFILE").unwrap_or(String::from("."))
        } else {
            env::var("HOME").unwrap_or(String::from("."))
        };

        let path = Path::new(&home_dir)
            .join(".curdrc")
            .into_os_string()
            .into_string()
            .unwrap();
        path
    }
}
