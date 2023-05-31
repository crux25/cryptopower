package lightning

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/decred/dcrlnd/lncfg"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Service struct {
	daemon *Daemon
	client *Client
	config *ServiceConfig
}

// ServiceConfig represents the configuration of the lightning service.
type ServiceConfig struct {
	WorkingDir string
	Network    string
}

func NewService(config *ServiceConfig) (*Service, error) {
	dm := NewDaemon(config)
	return &Service{
		daemon: dm,
		config: config,
	}, nil
}

func (s *Service) Start() error {
	go func() {
		go s.daemon.Start()
		ticker := time.NewTicker(time.Second * 10)

		// Wait for the daemon to start
		// TODO: write a listener outside that doesn't block and doesn't use lightning client.
		time.Sleep(120 * time.Second)

		// TODO: This will fail because the daemon don't have a lightning wallet at this point, so admin macaroon
		// is not created. Wallet should be created before calling this function.
		cl, err := NewClient(buildClienConfig(s.config))
		if err != nil {
			log.Infof("Error creating client %+v \n", err)
		}
		s.client = cl
		// poll and fetch node information to ascertain connectivity.
		for {
			select {
			case <-ticker.C:
				if s.client != nil {
					state, err := s.client.Client.GetInfo(context.Background())
					if err != nil {
						log.Infof("Error getting state %s: \n", err)
						continue
					}
					fmt.Printf("State: %+v\n", state)
				}
			default:
				continue
			}
		}
	}()

	return nil
}

// Initialize a new lightning wallet
func (s *Service) InitWallet(privatePassphrase string) error {
	conn, err := s.getBasicClientCon()
	if err != nil {
		return err
	}

	walletClient := lnrpc.NewWalletUnlockerClient(conn)
	// TODO: Save the wallet seed util the user is able to verify the wallet seed.
	seedR, err := walletClient.GenSeed(context.Background(), &lnrpc.GenSeedRequest{})
	if err != nil {
		return err
	}

	req := &lnrpc.InitWalletRequest{
		WalletPassword:     []byte(privatePassphrase),
		RecoveryWindow:     200,
		CipherSeedMnemonic: seedR.CipherSeedMnemonic,
	}
	_, err = walletClient.InitWallet(context.Background(), req)
	if err != nil {
		return err
	}

	return err
}

func (s *Service) getBasicClientCon() (*grpc.ClientConn, error) {
	creds, err := credentials.NewClientTLSFromFile(path.Join(s.config.WorkingDir, defaultTLSCertFilename), "")
	if err != nil {
		return nil, err
	}
	// Create a dial options array.
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		// We need to use a custom dialer so we can also connect to
		// unix sockets and not just TCP addresses.
		grpc.WithContextDialer(lncfg.ClientAddressDialer(defaultRPCPort)),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(lnrpc.MaxGrpcMsgSize)),
	}

	conn, err := grpc.Dial(defaultRPCHostPort, opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to RPC server: %v",
			err)
	}

	return conn, nil
}

func (s *Service) GetClient() *Client {
	return s.client
}
