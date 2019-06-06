package errors

const UnreadableFile string = "Unable to read file %v"
const UnparsableConfigFile string = "The config.toml file could not be parsed"
const UnexistentConfigField string = "The requested configuration field %v does not exist on loaded config.toml"
