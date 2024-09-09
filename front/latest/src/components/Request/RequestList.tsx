import { useState } from 'react'
import { Button, List, Space, Typography } from 'antd'
import RequestMethod from './RequestMethod'
import moment from 'moment'
import { CloseOutlined } from '@ant-design/icons'


const RequestList = (props: any) => {

    const [selected, setSelected] = useState(-1)

    const handleSelectRequest = (e: React.MouseEvent<HTMLElement>, index: number) => {
        e.stopPropagation()
        setSelected(index)
        props.onSelectRequest(index)
    }

    const handleDeleteRequest = (e: React.MouseEvent<HTMLElement>, index: number) => {
        e.stopPropagation()
        props.onDeleteRequest(index)
    }

    const renderItem = (item: any, index: number) => {
        const color = index === selected ? "white" : "black"
        return <List.Item
                    key={item.uuid} 
                    onClick={(e) => handleSelectRequest(e, index)}
                    style={{
                        backgroundColor: index === selected ? "#5ca0ff" : "transparent",
                        color: color,
                        cursor: "pointer",
                    }}
                >
                    <Space direction="vertical">
                        <Space direction="horizontal">
                            <RequestMethod method={item.method} /> 
                            {'#' + item.uuid.split("-")[0]}
                        </Space>
                        <Typography.Text style={{ fontSize: 12, color: color }}>{moment(item.date).format("YYYY-MM-DD HH:mm:ss")}</Typography.Text>
                    </Space>
                    <Button type='primary'danger size='small' onClick={(e) => handleDeleteRequest(e, index)}>
                        <CloseOutlined />
                    </Button>
            </List.Item>
    }

    return (
        <List
            size='large'
            header={<div className='request-card-title'>Requests</div>}
            bordered
            style={{
                minWidth: "350px",
                minHeight: "100vh"
            }}
            dataSource={props.messages}
            renderItem={renderItem}
        />
    )
}

export default RequestList