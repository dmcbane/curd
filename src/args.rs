extern crate docopt;

use self::docopt::Docopt;
use std::env;
use std::path::Path;

const USAGE: &'static str = "
CURD - Change to a User's Recurring Directory 1.0.0
H. Dale McBane<h.dale.mcbane@gmail.com>
Save and return to paths you visit often.

Usage:
    curd clean [--config <file>] [--verbose]
    curd (ls | list) [--config <file>] [--verbose]
    curd remove [KEYWORD] [--config <file>] [--verbose]
    curd save [KEYWORD] [--dir <directory>] [--config <file>] [--verbose]
    curd [KEYWORD] [--config <file>] [--verbose]
    curd (-h | --help)
    curd (-V | --version)

Options:
    --config=<file>  Specify configuration filename [default: <<replaceme>>].
    --dir=<directory>  Specify configuration filename [default: <current directory>].
    -h, --help     Show this screen.
    -V, --version  Show version.
    -v, --verbose  Display extra information.

Examples:
    List all of the paths defined in the default configuration file.
        curd ls

    List all of the paths in a specified configuration file.
        curd list --config some_configuration_file

    Clean paths from the default configuration that do not exist in the
    filesystem.
        curd clean

    Read the default path from the default configuration file.
        curd

    Save the current directory as the default path in the default configuration
    file.
        curd save

    Save the specified directory as the specified path in the default
    configuration file.
        curd save curd --dir=~/go/src/github.com/dmcbane/curd

    Remove the specified path from the default configuration file.
        curd remove essay

";

pub struct Args {
    pub configfile: String,
    pub keyword: String,
    pub clean: bool,
    pub list: bool,
    pub read: bool,
    pub remove: bool,
    pub save: bool,
    pub directory: String,
    pub verbose: bool,
}

impl Args {
    pub fn new() -> Result<Args, &'static str> {
        let default_config = Args::get_default_config_filename();
        let usage = USAGE.replace("<<replaceme>>", &default_config);
        let args = Docopt::new(usage)
            .and_then(|dopt| dopt.parse())
            .unwrap_or_else(|e| e.exit());

        // display debug information?
        let verbose = args.get_bool("--verbose");

        //// // Gets a value for config if supplied by user, or defaults to "~/.curdrc"
        let config = args.get_str("--config");
        let directory = args.get_str("--dir");
        let mut keyword = args.get_str("KEYWORD");
        if keyword == "" {
            keyword = "<default>"
        };
        let clean = args.get_bool("clean");
        let list = args.get_bool("ls") || args.get_bool("list");
        let remove = args.get_bool("remove");
        let save = args.get_bool("save");

        // reading if only keyword is provided or reading default if nothing is provided
        let read: bool = !clean && !list && !remove && !save;

        if verbose {
            println!("verbose: {}", verbose);
            println!("configuration file: {}", config);
            println!("keyword: {}", keyword);
            println!("clean: {}", clean);
            println!("list: {}", list);
            println!("remove: {}", remove);
            println!("save: {}", save);
            println!("directory: {}", directory);
            println!("read: {}", read);
        };

        return Ok(Args {
                      configfile: config.to_string(),
                      keyword: keyword.to_string(),
                      clean: clean,
                      list: list,
                      read: read,
                      remove: remove,
                      save: save,
                      directory: directory.to_string(),
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
