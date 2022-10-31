package main

import (
	"fmt"
	"log"
	"os"

	"aspenmesh/kickstarter/types"
	"aspenmesh/kickstarter/util"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:                 "kickstarter",
		Usage:                "Kickstarts an Aspen Mesh application",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:  "api",
				Usage: "Sets the Kickstarter API URL",
				Action: func(c *cli.Context) error {
					if !c.Args().Present() {
						fmt.Println("No API URL passed.  Please rerun with the URL of the API.")
						return nil
					}
					util.SetAPIURL(c.Args().First())
					_, err := util.GetAPIURL()
					if err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
			{
				Name:  "deployment",
				Usage: "Creates a new deployment",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "directory",
						Aliases: []string{"d"},
						Usage:   "The directory to save the yaml files to (defaults to your home directory)",
					},
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "Name of the deployment",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "namespace",
						Aliases: []string{"ns"},
						Usage:   "Namespace for the deployment",
					},
					&cli.StringFlag{
						Name:     "container-image",
						Aliases:  []string{"i"},
						Usage:    "The container image for the deployment pods)",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "container-name",
						Aliases:  []string{"c"},
						Usage:    "The name of the container for the pod",
						Required: true,
					},
					&cli.Int64Flag{
						Name:     "container-port",
						Aliases:  []string{"p"},
						Usage:    "The port that the container should listen on",
						Required: true,
					},
					&cli.StringFlag{
						Name:        "image-pull-policy",
						Usage:       "The pull policy for the image (one of Always | Never | IfNotPresent)",
						DefaultText: "Always",
					},
					&cli.IntFlag{
						Name:        "replicas",
						Usage:       "Number of replicas",
						DefaultText: "1",
					},
					&cli.StringFlag{
						Name:        "deployment-strategy",
						Usage:       "The deployment strategy (one of RollingUpdate | Recreate)",
						DefaultText: "RollingUpdate",
					},
					&cli.StringFlag{
						Name:        "service-account",
						Usage:       "The name of the service account to use",
						DefaultText: "default",
					},
				},
				Action: func(c *cli.Context) error {

					d := types.Deployment{
						Directory:          c.String("directory"),
						Name:               c.String("name"),
						Namespace:          c.String("namespace"),
						ContainerImage:     c.String("container-image"),
						ContainerName:      c.String("container-name"),
						ImagePullPolicy:    c.String("image-pull-policy"),
						DeploymentStrategy: c.String("deployment-strategy"),
						ServiceAccount:     c.String("service-account"),
						Replicas:           c.Int64("replicas"),
						ContainerPort:      c.Int64("container-port"),
					}
					types.GetDeployment(d)

					s := types.Service{
						Directory: c.String("directory"),
						Name:      c.String("name"),
						Namespace: c.String("namespace"),
						Port:      c.Int64("container-port"),
					}
					types.GetService(s)

					return nil
				},
			},
			{
				Name:  "virtual-service",
				Usage: "Creates a new virtual service",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "directory",
						Aliases: []string{"d"},
						Usage:   "The directory to save the yaml files to (defaults to your home directory)",
					},
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "Name of the deployment",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "namespace",
						Aliases: []string{"ns"},
						Usage:   "Namespace for the deployment",
					},
					&cli.StringFlag{
						Name:     "host",
						Usage:    "The host name for this virtual service (can be *)",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "gateway",
						Usage: "The name of the gateway that this virtual service is mapped to",
					},
					&cli.StringFlag{
						Name:  "destination-host",
						Usage: "The host name to route to",
					},
					&cli.Int64Flag{
						Name:  "destination-port",
						Usage: "The port to route to",
					},
				},
				Action: func(c *cli.Context) error {

					vs := types.VirtualService{
						Directory:             c.String("directory"),
						Name:                  c.String("name"),
						Namespace:             c.String("namespace"),
						Host:                  c.String("host"),
						GatewayName:           c.String("gateway"),
						DestinationHost:       c.String("destination-host"),
						DestinationPortNumber: c.Int64("destination-port"),
					}
					types.GetVirtualService(vs)
					return nil
				},
			},
			{
				Name:  "gateway",
				Usage: "Creates a new gateway",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "directory",
						Aliases: []string{"d"},
						Usage:   "The directory to save the yaml files to (defaults to your home directory)",
					},
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "Name of the deployment",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "namespace",
						Aliases: []string{"ns"},
						Usage:   "Namespace for the deployment",
					},
					&cli.StringFlag{
						Name:     "host",
						Usage:    "The host that the gateway is bound to (can be *)",
						Required: true,
					},
					&cli.Int64Flag{
						Name:     "port",
						Usage:    "The port of the gateway ",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "port-name",
						Usage:    "The name of the port for the gateway",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "protocol",
						Usage:    "The protocol being used.  One of HTTP | HTTPS | HTTP/2 | GRPC | TCP | TLS.",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "tls-mode",
						Usage: "The TLS mode to use (if protocol is HTTPS or TLS).  One of PASSTHROUGH | SIMPLE | MUTUAL | AUTO_PASSTHROUGH | ISTIO_MUTUAL",
					},
					&cli.StringFlag{
						Name:  "credential-name",
						Usage: "The name of the TLS credential secret (if protocol is HTTPS or TLS)",
					},
				},
				Action: func(c *cli.Context) error {

					gw := types.Gateway{
						Directory:      c.String("directory"),
						Name:           c.String("name"),
						Namespace:      c.String("namespace"),
						Port:           c.Int64("port"),
						PortName:       c.String("port-name"),
						Protocol:       c.String("protocol"),
						Host:           c.String("host"),
						TLSMode:        c.String("tls-mode"),
						CredentialName: c.String("credential-name"),
					}
					types.GetGateway(gw)
					return nil

				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
