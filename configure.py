from mininet.net import Mininet
from mininet.cli import CLI
from mininet.link import Link, TCLink,Intf
from subprocess import Popen, PIPE
from mininet.log import setLogLevel


if '__main__' == __name__:
  setLogLevel('info')
  net = Mininet(link=TCLink)
  key = "net.mptcp.mptcp_enabled"
  value = 1
  p = Popen("sysctl -w %s=%s" % (key, value), shell=True, stdout=PIPE, stderr=PIPE)
  stdout, stderr = p.communicate()
  print ("stdout=",stdout,"stderr=", stderr)
  h1 = net.addHost('h1')
  h2 = net.addHost('h2')
  r1 = net.addHost('r1')
  linkopt={'bw':10}
  net.addLink(r1,h1,cls=TCLink, **linkopt)
  net.addLink(r1,h1,cls=TCLink, **linkopt)
  net.addLink(r1,h2,cls=TCLink, **linkopt)
  net.addLink(r1,h2,cls=TCLink, **linkopt)
  net.build()
  r1.cmd("ifconfig r1-eth0 0")
  r1.cmd("ifconfig r1-eth1 0")
  r1.cmd("ifconfig r1-eth2 0")
  r1.cmd("ifconfig r1-eth3 0")
  h1.cmd("ifconfig h1-eth0 0")
  h1.cmd("ifconfig h1-eth1 0")
  h2.cmd("ifconfig h2-eth0 0")
  h2.cmd("ifconfig h2-eth1 0")
  r1.cmd("echo 1 > /proc/sys/net/ipv4/ip_forward")
  r1.cmd("ifconfig r1-eth0 10.0.0.1 netmask 255.255.255.0")
  r1.cmd("ifconfig r1-eth1 10.0.1.1 netmask 255.255.255.0")
  r1.cmd("ifconfig r1-eth2 10.0.2.1 netmask 255.255.255.0")
  r1.cmd("ifconfig r1-eth3 10.0.3.1 netmask 255.255.255.0")
  h1.cmd("ifconfig h1-eth0 10.0.0.2 netmask 255.255.255.0")
  h1.cmd("ifconfig h1-eth1 10.0.1.2 netmask 255.255.255.0")
  h2.cmd("ifconfig h2-eth0 10.0.2.2 netmask 255.255.255.0")
  h2.cmd("ifconfig h2-eth1 10.0.3.2 netmask 255.255.255.0")
  h1.cmd("ip rule add from 10.0.0.2 table 1")
  h1.cmd("ip rule add from 10.0.1.2 table 2")
  h1.cmd("ip route add 10.0.0.0/24 dev h1-eth0 scope link table 1")
  h1.cmd("ip route add default via 10.0.0.1 dev h1-eth0 table 1")
  h1.cmd("ip route add 10.0.1.0/24 dev h1-eth1 scope link table 2")
  h1.cmd("ip route add default via 10.0.1.1 dev h1-eth1 table 2")
  h1.cmd("ip route add default scope global nexthop via 10.0.0.1 dev h1-eth0")
  h2.cmd("ip rule add from 10.0.2.2 table 1")
  h2.cmd("ip rule add from 10.0.3.2 table 2")
  h2.cmd("ip route add 10.0.2.0/24 dev h2-eth0 scope link table 1")
  h2.cmd("ip route add default via 10.0.2.1 dev h2-eth0 table 1")
  h2.cmd("ip route add 10.0.3.0/24 dev h2-eth1 scope link table 2")
  h2.cmd("ip route add default via 10.0.3.1 dev h2-eth1 table 2")
  h2.cmd("ip route add default scope global nexthop via 10.0.2.1 dev h2-eth0")
  # h1.cmd("python3 ~/Documents/np/code.py receiver localhost")
  # h2.cmd("python3 ~/Documents/np/code.py sender localhost")
  CLI(net)
  net.stop()
