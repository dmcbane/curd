extern crate clap;

use std::env;
use std::path::Path;
use self::clap::{App, Arg, ArgGroup};

pub struct Args {
    pub configfile: String,
    pub keyword: String,
    pub clean: bool,
    pub list: bool,
    pub read: bool,
    pub remove: bool,
    pub save: bool,
    pub verbose: bool,
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
                               .help("Remove the path specified by the keyword or the default path from the configuration file."))
                          .arg(Arg::with_name("save")
                               .short("s")
                               .long("save")
                               .help("Save the current directory to the specified keyword or the default."))
                          .arg(Arg::with_name("verbose")
                               .short("v")
                               .long("verbose")
                               .help("Display extra information."))
                          .group(ArgGroup::with_name("One Action")
                                 .args(&["clean", "list", "remove", "save"]))
                          .group(ArgGroup::with_name("No Keyword")
                                 .args(&["clean", "list", "keyword"]))
                          .after_help("The 'clean' and 'list' flags cannot be combined with the keyword arg. The 'remove' or 'save' flags can be used alone or with the 'keyword' arg.  The 'config' option can be combined with all flags,and args.")
                          .get_matches_from(args);

        // display debug information?
        let verbose = matches.is_present("verbose");

        // Gets a value for config if supplied by user, or defaults to "~/.curdrc"
        let default_config = Args::get_default_config_filename();
        let config = (matches.value_of("config").unwrap_or(&default_config)).to_string();

        let keyword = if matches.is_present("keyword") {
            matches.value_of("keyword").unwrap().to_string()
        // } else if matches.is_present("remove") {
        //     matches.value_of("remove").unwrap().to_string()
        // } else if matches.is_present("save") {
        //     matches.value_of("save").unwrap().to_string()
        } else {
            "<default>".to_string()
        };

        let clean = matches.is_present("clean");
        let list = matches.is_present("list");
        let remove = matches.is_present("remove");
        let save = matches.is_present("save");

        // reading if only keyword is provided or reading default if nothing is provided
        let read: bool = !clean && !list && !remove && !save;

        if verbose {
            println!("verbose: {}", verbose);
            println!("default configuration file: {}", default_config);
            println!("configuration file: {}", config);
            println!("keyword: {}", keyword);
            println!("clean: {}", clean);
            println!("list: {}", list);
            println!("remove: {}", remove);
            println!("save: {}", save);
            println!("read: {}", read);
        };

        return Ok(Args {
            configfile: config,
            keyword: keyword,
            clean: clean,
            list: list,
            read: read,
            remove: remove,
            save: save,
            verbose: verbose,
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
