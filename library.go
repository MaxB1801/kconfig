package main

// define structs for input yaml
type Cluster struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	IP       string `yaml:"ip"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type KConfig struct {
	Savedir struct {
		Filepaths []string `yaml:"filepaths"`
	} `yaml:"savedir"`
	Clusters []Cluster `yaml:"clusters"`
}

// define struct for the ssh config
type Config struct {
	Clusters       []Clusters  `yaml:"clusters"`
	Contexts       []Context   `yaml:"contexts"`
	CurrentContext string      `yaml:"current-context"`
	Kind           string      `yaml:"kind"`
	Preferences    Preferences `yaml:"preferences"`
	Users          []User      `yaml:"users"`
}

type Clusters struct {
	Name          string `yaml:"name"`
	ServerContext struct {
		CertificateAuthorityData string `yaml:"certificate-authority-data"`
		Server                   string `yaml:"server"`
	} `yaml:"cluster"`
}

type Context struct {
	Name    string `yaml:"name"`
	Context struct {
		Cluster string `yaml:"cluster"`
		User    string `yaml:"user"`
	} `yaml:"context"`
}

type Preferences struct{}

type User struct {
	Name string `yaml:"name"`
	User struct {
		ClientCertificateData string `yaml:"client-certificate-data"`
		ClientKeyData         string `yaml:"client-key-data"`
	} `yaml:"user"`
}
