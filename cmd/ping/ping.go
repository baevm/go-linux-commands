package ping

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

var (
	// packets count
	packetsCount int
	// packet size
	packetSize int
	// ttl
	ttl int
)

func init() {
	Cmd.Flags().IntVarP(&packetsCount, "count", "c", 4, "Number of packets to send")
	Cmd.Flags().IntVarP(&packetSize, "size", "s", 64, "Packet size")
	Cmd.Flags().IntVarP(&ttl, "ttl", "t", 59, "Time To Live")
}

var Cmd = &cobra.Command{
	Use:   "ping [OPTIONS] [URL]",
	Short: "Send packets to host.",
	Long:  `Send ICMP ECHO_REQUEST to network hosts. Works only in sudo mode.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   pingHost,
}

func pingHost(cmd *cobra.Command, args []string) {
	url := args[0]

	var successCount int
	var failedCount int
	packet := make([]byte, packetSize)

	ip, err := net.ResolveIPAddr("ip4:icmp", url)

	if err != nil {
		log.Fatal(err)
	}

	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	conn.IPv4PacketConn().SetTTL(ttl)

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: packet,
		},
	}

	mp, err := msg.Marshal(nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PING %s (%s) %d bytes of data.\n", url, ip.String(), len(mp))

	for i := 0; i < packetsCount; i++ {
		// sleep for 1 second to avoid overflooding
		time.Sleep(time.Second)

		_, err := conn.WriteTo(mp, &net.IPAddr{IP: net.ParseIP(ip.String())})

		if err != nil {
			log.Fatal("Error writing to host %v", err)
		}

		reply, err := parseMessage(conn)

		if err != nil {
			log.Fatal("Error parsing message %v", err)
		}

		switch reply.Code {
		case 0:
			successCount += 1
			fmt.Printf("%d bytes from %s (%s): ttl=%d\n",
				len(reply.Body.(*icmp.Echo).Data),
				url,
				ip.String(),
				ttl,
			)
		case 3:
			failedCount += 1
			fmt.Printf("%s is unreachable\n", ip.String())
		case 11:
			failedCount += 1
			fmt.Printf("%s is slow\n", ip.String()) // Time Exceeded
		default:
			failedCount += 1
			fmt.Printf("%s is unreachable\n", ip.String())
		}

	}

	fmt.Printf("--- %s ping statistics ---\n", url)
	fmt.Printf("%d packets transmitted, %d packets received, %d packets lost \n", packetsCount, successCount, failedCount)
}

func parseMessage(conn *icmp.PacketConn) (*icmp.Message, error) {
	replyBuff := make([]byte, 1500)
	n, _, err := conn.ReadFrom(replyBuff)

	if err != nil {
		return nil, fmt.Errorf("read response error %v", err)
	}
	parsed_reply, err := icmp.ParseMessage(1, replyBuff[:n])

	if err != nil {
		return nil, fmt.Errorf("parse response error %v", err)
	}

	return parsed_reply, nil
}
