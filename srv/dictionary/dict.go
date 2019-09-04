package dictionary

import (
	"context"
	"log"
	"net"
	"os"

	dict "github.com/Flyewzz/balda-backend/pkg/dictionary"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type DictionaryManager struct{}

func (dm *DictionaryManager) CheckWord(ctx context.Context, word *dict.Word) (*dict.Availability, error) {
	return &dict.Availability{}, nil
}

func (dm *DictionaryManager) AddWord(ctx context.Context, word *dict.Word) (*dict.Nothing, error) {
	return &dict.Nothing{}, nil
}

func (dm *DictionaryManager) RemoveWord(ctx context.Context, word *dict.Word) (*dict.Nothing, error) {
	return &dict.Nothing{}, nil
}

func main() {
	viper.SetConfigFile(os.Args[1])
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Cannot read config", err)
		return
	}
	port := viper.GetString("port")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	dict.RegisterDictionaryServer(server, &DictionaryManager{})

}
