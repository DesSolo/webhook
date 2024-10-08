import React from 'react'
import { Card, Checkbox, Typography } from 'antd'
import {decode as b64Decode} from 'js-base64'


interface RequestContentProps {
    content: string
}

interface RequestContentState {
    formatJSON: boolean
}

class RequestContent extends React.Component<RequestContentProps, RequestContentState> {
    constructor(props: RequestContentProps) {
        super(props)
        this.state = {
            formatJSON: true
        }
    }
    renderActions() {
        return <div className='request-card-action'>
            <Checkbox checked={this.state.formatJSON} onChange={() => this.setState({ formatJSON: !this.state.formatJSON })}>
                Format JSON
            </Checkbox>
            <Typography.Text copyable={{ text: b64Decode(this.props.content) }}/>
        </div>
    }

    render() {
        let data = this.props.content

        if (!data) {
            return <div style={{color: "#ccc"}}>No content</div>
        }

        try {
            data = b64Decode(data)
        } catch {}
        
        if (this.state.formatJSON) {
            try {
                data = JSON.stringify(JSON.parse(data), null, 2)
            } catch {}
        }

        return <Card size="small" title={<div>Request content {this.renderActions()}</div>}>
            <pre className="code" style={{
                border: '1px solid #ccc',
                padding: '10px',
            }}>{data}</pre>
        </Card>
    }
}

export default RequestContent