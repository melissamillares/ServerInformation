package main

import (
	"time"
)

func getServers(r string) []Server {
	outputURL := isURL(r)
	var server Server
	var servers []Server

	if outputURL {
		host := hostName(r)
		IPs := getIP(host) // array with the server IPs
		owner := getInfoWhoIs("OrgName: ", IPs) // server owner		
		country := getInfoWhoIs("Country: ", IPs) // server country							
		ssl := getSSLGrade(host, len(IPs)) // SSL grade		

		for i, value := range ssl {
			server = Server {
				//ID: ,
				Address: IPs[i].String(),
				SSL_grade: value,
				Country: country,
				Owner: owner,
				//DomainID: host,
				Domain: host,
				Created: time.Now(),
			}
			servers = append(servers, server)
		}
	}

	return servers
}

func getDomain(r string, servers []Server) *Domain {	
	outputURL := isURL(r)
	domain := Domain{}

	if outputURL {
		host := hostName(r)
		IPs := getIP(host) // array with the server IPs			
		title := getTitle(r, "title") // domain html title		
		ssl := getSSLGrade(host, len(IPs)) // SSL grade		
		lowerSSL := getLowerGrade(ssl) // lower SSL grade
		serverDown := isServerDown(r)		

		domain = Domain{
			//ID: ,
			URL: host,
			Servers: servers,
			Servers_Changed: false,
			SSL: lowerSSL,
			Previous_SSL: lowerSSL,
			Logo: lowerSSL,
			Title: title,
			Is_Down: serverDown,
			Created: time.Now(),
		}
	}	

	return &domain
}
