package utils

import "net"

func GetIPAddressList() []string {
	addrs, err := net.InterfaceAddrs()
	var ips []string
	if err != nil {
		return ips
	}

	for _, address := range addrs {
		//fmt.Println(address.(*net.IPNet))
		ips = append(ips, address.(*net.IPNet).IP.String())

		// 检查ip地址判断是否回环地址
		//if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
		//	if ipnet.IP.To4() != nil {
		//		ips = append(ips, ipnet.IP.String())
		//	}
		//}
	}

	return ips
}
