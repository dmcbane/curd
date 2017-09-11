pub struct Args {
    pub configfile: String,
}

impl Args {
    pub fn new(args: &[String]) -> Result<Args, &'static str> {
        if args.len() < 2 {
            // return Err("not enough arguments");
            return Ok(Args { configfile: String::new() });
        } else {
            return Ok(Args { configfile: args[1].clone() });
        }
    }
}
