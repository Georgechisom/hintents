use clap::Parser;

#[derive(Parser, Debug)]
#[command(author, version, about, long_about = None)]
struct Args {
    /// The base64 encoded transaction envelope
    #[arg(short, long)]
    envelope: String,

    /// The base64 encoded ledger footprint (JSON components)
    #[arg(short, long)]
    footprint: String,
}

fn main() {
    let args = Args::parse();

    println!("Erst Simulator Starting...");
    println!("Received Envelope: {}...", &args.envelope.chars().take(10).collect::<String>());
    
    // TODO: Initialize Host, Load Contract, Invoke
}
