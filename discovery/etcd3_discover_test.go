package discovery

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/akkagao/gms/plugin"
	"github.com/akkagao/gms/plugin/registe"
)

func TestNewEtcd3Discover(t *testing.T) {
	plugins := registe.NewEtcd3RegistePlugin("gmsDemo", []string{"localhost:2379"})
	plugins.Start()

	registePlugin := plugins.(plugin.IRegistePlugin)
	registePlugin.Registe("127.0.0.1", 1024)

	discover, err := NewEtcd3Discover("gmsDemo", []string{"localhost:2379"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(discover.GetServer())

	timeout := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timeout.C:
			fmt.Println("---")
			rand.Seed(time.Now().UnixNano())
			registePlugin.Registe("127.0.0.1", rand.Intn(1000))
			assert.NoError(t, err)

			fmt.Println(discover.GetServer())
		}
	}

	select {}
}

func TestNewEtcd3Discover2(t *testing.T) {
	plugins := registe.NewEtcd3RegistePlugin("gmsDemo", []string{"localhost:2379"})
	plugins.Start()

	registePlugin := plugins.(plugin.IRegistePlugin)
	registePlugin.Registe("127.0.0.1", 1024)

	discover, err := NewEtcd3Discover("gmsDemo", []string{"localhost:2379"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(discover.GetServer())

	timeout := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timeout.C:
			fmt.Println(discover.GetServer())
		}
	}

	select {}
}
