import { useState } from 'react'
import { Button, List, Space, Typography } from 'antd'
import RequestMethod from './RequestMethod'
import moment from 'moment'
import { CloseOutlined } from '@ant-design/icons'

const PAGE_SIZE = 9


const RequestList = (props: {
    messages: any[]
    onDeleteRequest: (uuid: string) => void
    onSelectRequest: (uuid: string) => void
}) => {

    const [selected, setSelected] = useState("")
    const [page, setPage] = useState(1)

    const handleSelectRequest = (e: React.MouseEvent<HTMLElement>, uuid: string) => {
        e.stopPropagation()
        setSelected(uuid)
        props.onSelectRequest(uuid)
    }

    const handleDeleteRequest = (e: React.MouseEvent<HTMLElement>, uuid: string) => {
        e.stopPropagation()
        props.onDeleteRequest(uuid)
    }

    const dataSource = () => {
        return props.messages.slice((page - 1) * PAGE_SIZE, page * PAGE_SIZE)
    }

    const renderFooter = () => {
        const totalMessages = props.messages.length
        const lastPage = Math.ceil(totalMessages / PAGE_SIZE)

        const hasMessages = totalMessages > 0
        const isFistPage = page === 1
        const isLastPage = page === lastPage

        return (
            <>
                <Button type="link" onClick={() => setPage(1)} disabled={!hasMessages || isFistPage}>First</Button>
                <Button type="link" onClick={() => setPage(page - 1)} disabled={isFistPage}>← Prev</Button>
                <Button type="link" onClick={() => setPage(page + 1)} disabled={isLastPage || !hasMessages}>Next →</Button>
                <Button type="link" onClick={() => setPage(lastPage)} disabled={!hasMessages || isLastPage}>Last</Button>
            </>
        )
    }

    const renderItem = (item: any) => {
        const color = item.uuid === selected ? "white" : "black"
        return <List.Item
                    key={item.uuid} 
                    onClick={(e) => handleSelectRequest(e, item.uuid)}
                    style={{
                        backgroundColor: item.uuid === selected ? "#5ca0ff" : "transparent",
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
                    <Button type='primary'danger size='small' onClick={(e) => handleDeleteRequest(e, item.uuid)}>
                        <CloseOutlined style={{ fontSize: 10 }} />
                    </Button>
            </List.Item>
    }

    return (
        <List
            header={<div className='request-card-title'>Requests: {props.messages.length}</div>}
            bordered
            style={{
                minWidth: "340px",
                minHeight: "100vh",
                marginLeft: "10px",
            }}
            dataSource={dataSource()}
            renderItem={renderItem}
            footer={renderFooter()}
        />
    )
}

export default RequestList