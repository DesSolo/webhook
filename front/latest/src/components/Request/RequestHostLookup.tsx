import React from 'react'

interface RequestHostLookupProps {
    ip: string
}

class RequestHostLookup extends React.Component<RequestHostLookupProps> {
    services = [
        ["Whois", "https://whois.domaintools.com/%s"],
        ["Shodan", "https://www.shodan.io/host/%s"],
        ["Netify", "https://www.netify.ai/resources/ips/%s"],
        ["Censys", "https://search.censys.io/hosts/%s"],
        ["VirusTotal", "https://www.virustotal.com/gui/ip-address/%s/relations"],
        ["IPInfo", "https://ipinfo.io/%s"],

    ]
    render() {
        let ip = this.props.ip

        if (ip.includes(":")) {
            ip = ip.split(":")[0]
        }

        const lookups = []

        for (let i = 0; i < this.services.length; i++) {
            const [name, url] = this.services[i]
            lookups.push(<a key={i} style={{marginRight: '10px'}} href={url.replace('%s', ip)} target="_blank">{name}</a>)
        }

        return (
            <div>
                {lookups}
            </div>
        )
    }
}

export default RequestHostLookup