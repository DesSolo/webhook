import React from 'react'
import { List, Space, Typography } from 'antd'

import RequestMethod from './RequestMethod'
import RequestHostLookup from './RequestHostLookup'
import RequestContent from './RequestContent'
import moment from 'moment'

const { Text, Link } = Typography


const url = (request: any) => {
    return request.schema + "://" + request.host + request.uri
}


interface RequestProps {
    request: any
}

class Request extends React.Component<RequestProps> {
    copyAsCurl() {
        const nl = '\\\n  '

        let result = ["curl"]

        result.push('--request', this.props.request.method, nl)
        result.push('--url', url(this.props.request), nl)

        if (this.props.request.headers) {
            Object.entries(this.props.request.headers).map((item: any) => {
                const [key, value] = item
                result.push(`--header '${key}: ${value}'`, nl)
            })
        }

        if (this.props.request.data) {
            result.push(`--data '${atob(this.props.request.data)}'`)
        } else {
            result.splice(-1)
        }

        navigator.clipboard.writeText(result.join(" "))
    }

    requestDetailsItems() {
        const date = moment(this.props.request.date)

        const requestURL = url(this.props.request)
        
        // TODO: note, full uri
        return [
            {
                key: <RequestMethod method={this.props.request.method} />,
                value: <Link href={requestURL} target="_blank">{requestURL}</Link>,
            },
            {
                key: "Host",
                value: <Space>
                    <Text copyable>{this.props.request.ip}</Text> <RequestHostLookup ip={this.props.request.ip} />
                </Space>,
            },
            {
                key: "Date",
                value: <div>{date.format("YYYY-MM-DD HH:mm:ss")} ({date.fromNow()})</div>,
            },
            {
                key: "ID",
                value: this.props.request.uuid,
            }
        ]
    }

    queryStringsItems() {
        let items: { key: string; value: string }[] = []

        this.props.request.query.split("&").forEach((item: string) => {
            const [key, value] = item.split("=")
            
            if (!key) {
                return
            }

            items.push({
                key: key,
                value: value,
            })
        })

        return items
    }

    headersItems() {
        let items: { key: string; value: string }[] = []

        Object.entries(this.props.request.headers).map((item: any) => {
            const [key, value] = item
            items.push({
                key: key,
                value: value,
            })
        })

        return items
    }


    formValuesItems() {
        let items: { key: string; value: string }[] = []

        if (!this.props.request.headers["Content-Type"]?.includes("application/x-www-form-urlencoded")) {
            return items
        }

        atob(this.props.request.data).split("&").forEach((item: string) => {
            const [key, value] = item.split("=")
            items.push({
                key: key,
                value: value,
            })
        })


        return items
    }

    render() {
        function listItem(item: { key: any; value: any }, className: string) {           
            return <List.Item>
                <div style={{ display: "flex" }}>
                    <div style={{minWidth: "150px"}}>{item.key}</div>
                    <div className={className}>{item.value}</div>
                </div>
            </List.Item>
        }

        return (
            <Space direction='vertical'>
                <Space direction="horizontal">
                    <List
                        className='request-card'
                        size="small"
                        header={<div className='request-card-title'>Request details <div className='request-card-action' onClick={() => this.copyAsCurl()}>Copy as cURL</div></div>}
                        bordered
                        dataSource={this.requestDetailsItems()}
                        renderItem={(item) => listItem(item, "")}
                    />
                    <List
                        className='request-card'
                        size="small"
                        header={<div className='request-card-title'>Headers</div>}
                        bordered
                        dataSource={this.headersItems()}
                        renderItem={(item) => listItem(item, "code")}
                    />
                </Space>
                <Space direction="horizontal">
                    <List
                        className='request-card'
                        size="small"
                        header={<div className='request-card-title'>Query strings</div>}
                        bordered
                        dataSource={this.queryStringsItems()}
                        renderItem={(item) => listItem(item, "code")}
                    />
                    <List
                        className='request-card'
                        size="small"
                        header={<div className='request-card-title'>Form values</div>}
                        bordered
                        dataSource={this.formValuesItems()}
                        renderItem={(item) => listItem(item, "code")}
                    />
                </Space>
                <RequestContent content={this.props.request.data} />
            </Space>
        )
    }
}

export default Request
export type { RequestProps }