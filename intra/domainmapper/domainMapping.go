package domainmapper

var Global_DomainList_Controller []DomainMap

func init() {
	Global_DomainList_Controller = []DomainMap{}
}

func FetchDomainFromController(GlobalDomain []DomainMap, ipaddress string) string {
	if GlobalDomain != nil {

		//log.Warnf("@CyberMine: fetching for domain controller")
		for i := 0; i < len(GlobalDomain); i++ {
			for j := 0; j < len(GlobalDomain[i].Ipaddresses); j++ {
				if GlobalDomain[i].Ipaddresses[j] == ipaddress {
					return GlobalDomain[i].Domain
				}
			}
		}
	} else {
		//log.Warnf("@CyberMine: fetching for domain controller is null")
	}
	return "empty"
}

func IsDomainExist(GlobalDomain []DomainMap, domainToSearch string) bool {
	if GlobalDomain != nil {

		if len(GlobalDomain) <= 0 {
			//log.Warnf("@CyberMine: Domain list is null")
			return false
		}
		for i := 0; i < len(GlobalDomain); i++ {
			if GlobalDomain[i].Domain == domainToSearch {
				//log.Warnf("@CyberMine: Domain Already Exists: Found" + domainToSearch)
				return true
			}
		}
		//log.Warnf("@CyberMine: Domain does'nt exist: Not Found" + domainToSearch)
		return false

	} else {
		//log.Warnf("@CyberMine: checking for domain exist is null")
	}
	return false
}

type DomainMap struct {
	Domain      string
	Ipaddresses []string
}
