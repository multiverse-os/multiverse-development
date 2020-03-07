extern crate clap;

extern crate portal_util;
extern crate jailer;

fn main() {
    if let Err(error) = jailer::run(
        jailer::clap_app().get_matches(),
        portal_util::get_time(portal_util::ClockType::Monotonic) / 1000,
        portal_util::get_time(portal_util::ClockType::ProcessCpu) / 1000,
    ) {
        panic!("Jailer error: {}", error);
    }
}
