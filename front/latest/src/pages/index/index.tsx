import React from "react"
import { v4 as uuid } from 'uuid';
import { Space, Layout, ConfigProvider, Spin } from "antd"

import Help from "../../components/Help";
import Request from "../../components/Request/Request"
import RequestList from "../../components/Request/RequestList"
import NavCopy from "../../components/Nav/NavCopy";
import { LoadingOutlined } from "@ant-design/icons";

import "./index.css"
import NavCreate from "../../components/Nav/NavCreate";

const { Header, Content } = Layout;


interface IndexState {
    messages: any[]
    selectedRequest: any
    channel: string
    connected: boolean
    // ws: any
}



class Index extends React.Component<any, IndexState> {

    private wsRef: React.RefObject<WebSocket>;

    constructor(props: any) {
        super(props)
        this.state = {
            messages: [],
            selectedRequest: null,
            channel: "",
            connected: false,
        }

        this.wsRef = React.createRef()

        this.handleUpdateChannel = this.handleUpdateChannel.bind(this)
        this.handleSelectRequest = this.handleSelectRequest.bind(this)
        this.handleDeleteRequest = this.handleDeleteRequest.bind(this)
    }

    connectToChannel(channel: string) {
        console.log("connectToWebSocket", channel)

        this.setState({connected: false})

        if (this.wsRef.current) {
            console.log("Closing previous websocket")
            this.wsRef.current.close()
        }

        const ws = new WebSocket("/ws?channel=" + channel)

        // @ts-ignore
        this.wsRef.current = ws

        ws.onerror = (event) => {
            console.log("Websocket error", event)
        }

        ws.onopen = () => {
            this.setState({
                connected: true,
                channel: channel
            })
        }

        ws.onmessage = (event) => {
            const request = JSON.parse(event.data)

            this.setState({ messages: [...this.state.messages, request] })

            if (!this.state.selectedRequest) {
                this.setState({ selectedRequest: request })
            }
        }

    }

    componentDidMount(): void {
        const uid = uuid()
        this.connectToChannel(uid)
    }

    handleUpdateChannel(channel: string) {
        this.connectToChannel(channel)
    }

    handleSelectRequest(index: number) {
        this.setState({ selectedRequest: this.state.messages[index] })
    }

    handleDeleteRequest(index: number) {
        this.state.messages.splice(index, 1)
        this.setState({ messages: this.state.messages })
    }

    render() {
        return (
            <ConfigProvider renderEmpty={() => <div>(empty)</div>}>
                <Layout>
                    <Header style={{ alignItems: 'center' }}>
                        <div>
                            <img src="/logo.svg" width={"50px"} style={{ marginTop: "10px" }} />
                            <div style={{ float: "right" }}>
                                <NavCopy uuid={this.state.channel} />
                                <NavCreate onUpdateChannel={this.handleUpdateChannel} />
                            </div>
                        </div>
                    </Header>

                    <Spin indicator={<LoadingOutlined style={{ fontSize: 24 }} spin />} spinning={!this.state.connected} tip="Connecting...">
                        <Content>
                            <div
                                style={{
                                    padding: '24px 0',
                                    background: '#fff',
                                }}
                            >
                                <Space direction="horizontal" align="start">
                                    <RequestList messages={this.state.messages} onDeleteRequest={this.handleDeleteRequest} onSelectRequest={this.handleSelectRequest} />
                                    {!this.state.selectedRequest && <Help channel={this.state.channel} />}
                                    {this.state.selectedRequest && <Request request={this.state.selectedRequest} />}
                                </Space>
                            </div>
                        </Content>
                    </Spin>
                </Layout>
            </ConfigProvider>
        )
    }
}

export default Index