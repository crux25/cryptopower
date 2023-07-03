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
	basicRPCConn *grpc.ClientConn
	daemon       *Daemon
	client       *Client
	config       *ServiceConfig
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

		// Poll every 2 secs and start the client once the daemon is ready.
		time.Sleep(2 * time.Second)

		for {
			select {
			case <-ticker.C:
				resp, err := s.getWalletState()
				if err != nil {
					fmt.Print(err)
				}

				if resp.State == lnrpc.WalletState_RPC_ACTIVE {
					cl, err := NewClient(buildClienConfig(s.config))
					if err == nil {
						s.client = cl
						// Return from this function if the client started succefully.
						return
					}
					log.Infof("Error creating client %+v \n", err)
				}

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

func (s *Service) UnlockWallet(privatePassphrase string) error {
	conn, err := s.getBasicClientCon()
	if err != nil {
		return err
	}

	req := &lnrpc.UnlockWalletRequest{
		WalletPassword: []byte(privatePassphrase),
		StatelessInit:  false,
	}

	client := lnrpc.NewWalletUnlockerClient(conn)
	_, err = client.UnlockWallet(context.Background(), req)
	if err != nil {
		return err
	}
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
	if s.basicRPCConn != nil {
		return s.basicRPCConn, nil
	}
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

	s.basicRPCConn, err = grpc.Dial(defaultRPCHostPort, opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to RPC server: %v",
			err)
	}

	return s.basicRPCConn, nil
}

func (s *Service) getWalletState() (*lnrpc.SubscribeStateResponse, error) {
	conn, err := s.getBasicClientCon()
	if err != nil {
		return nil, err
	}
	stateConn := lnrpc.NewStateClient(conn)
	req := &lnrpc.SubscribeStateRequest{}
	stream, err := stateConn.SubscribeState(context.Background(), req)
	if err != nil {
		return nil, err
	}

	// Recieve state message
	return stream.Recv()
}

func (s *Service) GetClient() *Client {
	return s.client
}

// IsWalletCreated returns true if the wallet has been created.
func (s *Service) IsWalletCreated() bool {
	resp, _ := s.getWalletState()
	if resp != nil {
		if resp.State != lnrpc.WalletState_NON_EXISTING {
			return true
		}
	}
	return false
}
