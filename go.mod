module github.com/rancher/multi-cluster-app

require (
	github.com/rancher/goautoneg v0.0.0-20120707110453-a547fc61f48d // indirect
	github.com/rancher/norman v0.0.0-20181015231214-04cb04ac0697
	github.com/rancher/types v0.0.0-20181018184021-9cf979f1663b
	github.com/sirupsen/logrus v1.0.4-0.20170822132746-89742aefa4b2
	github.com/urfave/cli v1.18.0 // indirect
	golang.org/x/sys v0.0.0-20181011152604-fa43e7bc11ba // indirect
	k8s.io/client-go v2.0.0-alpha.0.0.20180621152933-b0722d92a7c1+incompatible

)

replace github.com/rancher/types v0.0.0-20181018184021-9cf979f1663b => github.com/prachidamle/types v0.0.0-20181020064654-2c3b435f0686
