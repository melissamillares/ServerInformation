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
				Updated: time.Time{}, // null time
			}
			servers = append(servers, server)
		}
	}

	return servers
}

func getUpdatedServers(r string) []Server {
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
				Address: IPs[i].String(),
				SSL_grade: value,
				Country: country,
				Owner: owner,	
				Domain: host,							
				Updated: time.Now(),			
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
		title, serverDown := getTitle(r) // domain html title	
		logo := getLogo(r) // html logo
		ssl := getSSLGrade(host, len(IPs)) // SSL grade		
		lowerSSL := getLowerGrade(ssl) // lower SSL grade
		//serverDown := isServerDown(r)		

		domain = Domain{
			//ID: ,
			URL: host,
			Servers: servers,
			Servers_Changed: false,
			SSL: lowerSSL,
			Previous_SSL: "",
			Logo: logo,
			Title: title,
			Is_Down: serverDown,
			Created: time.Now(),
			Updated: time.Time{}, // null time
		}
	}	

	return &domain
}

func getUpdatedDomain(r string, servers []Server) *Domain {
	outputURL := isURL(r)
	domain := Domain{}

	if outputURL {
		host := hostName(r)
		IPs := getIP(host) // array with the server IPs							
		ssl := getSSLGrade(host, len(IPs)) // SSL grade	
		title, serverDown := getTitle(r) // domain html title	
		logo := getLogo(r) // html logo	
		lowerSSL := getLowerGrade(ssl) // lower SSL grade
		//serverDown := isServerDown(r)			

		domain = Domain{
			URL: host,
			Servers: servers,
			Servers_Changed: false,
			SSL: lowerSSL,
			//Previous_SSL: lowerSSL,
			Logo: logo,
			Title: title,
			Is_Down: serverDown,
			Updated: time.Now(),
		}
	}	

	return &domain
}
