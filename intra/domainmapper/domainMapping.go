package domainmappackage

import (
	"github.com/eycorsican/go-tun2socks/common/log"
)

var Global_DomainList_Controller []DomainMap

func init() {
	Global_DomainList_Controller = []DomainMap{}
}

func FetchDomainFromController(GlobalDomain []DomainMap, ipaddress string) string {
	if GlobalDomain != nil {

		log.Warnf("@CyberMine: fetching for domain controller")
		for i := 0; i < len(GlobalDomain); i++ {
			for j := 0; j < len(GlobalDomain[i].Ipaddresses); j++ {
				if GlobalDomain[i].Ipaddresses[j] == ipaddress {
					return GlobalDomain[i].Domain
				}
			}
		}
	} else {
		log.Warnf("@CyberMine: fetching for domain controller is null")
	}
	return "empty"
}

func IsDomainExist(GlobalDomain []DomainMap, domainToSearch string) bool {
	if GlobalDomain != nil {

		if len(GlobalDomain) <= 0 {
			log.Warnf("@CyberMine: Domain list is null")
			return false
		}
		for i := 0; i < len(GlobalDomain); i++ {
			if GlobalDomain[i].Domain == domainToSearch {
				log.Warnf("@CyberMine: Domain Already Exists: Found" + domainToSearch)
				return true
			}
		}
		log.Warnf("@CyberMine: Domain does'nt exist: Not Found" + domainToSearch)
		return false
		//log.Warnf("@CyberMine: checking for domain exist")
		//sort.Slice(GlobalDomain, func(i, j int) bool {
		//	log.Warnf("CyberMine: domain exist scan completed1")
		//	return GlobalDomain[i].Domain <= GlobalDomain[j].Domain
		//})
		//idx := sort.Search(len(GlobalDomain), func(i int) bool {
		//	log.Warnf("CyberMine: domain exist scan completed2")
		//	return string(GlobalDomain[i].Domain) >= domainToSearch
		//})
		//if GlobalDomain[idx].Domain == domainToSearch {
		//	log.Warnf("CyberMine: domain exist scan completed3")
		//	return true
		//}

	} else {
		log.Warnf("@CyberMine: checking for domain exist is null")
	}
	return false
}

type DomainMap struct {
	Domain      string
	Ipaddresses []string
}
