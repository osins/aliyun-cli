package service

type Cli interface {
	Run(args []*string) (_err error)
}
