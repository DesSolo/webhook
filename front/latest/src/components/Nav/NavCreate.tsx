import { useState } from "react"
import { Button, message, Modal, Typography } from "antd"
import ChannelCreate from "../Channel/ChannelCreate"
import { PlusOutlined } from "@ant-design/icons"

const NavCreate = (props: { onUpdateChannel: any }) => {
    const [visible, setVisible] = useState(false)
    const [messageApi, contextHolder] = message.useMessage();

    const onSuccess = (channel: string) => {
        props.onUpdateChannel(channel)
        setVisible(false)
    }

    const onError = (err: any) => {
        messageApi.error({
            content: err,
            duration: 2,
        })
    }

    return (
        <>
            {contextHolder}
            <Button
                type="text"
                style={{ color: "white" }}
                onClick={() => setVisible(true)}
                icon={<PlusOutlined />}
            >New
            </Button>
            <Modal
                title="Create new channel"
                open={visible}
                onCancel={() => setVisible(false)}
                footer={[
                    <Button form="createRequest" type="primary" key="submit" htmlType="submit">Create</Button>
                ]}
            >
                <div style={{ marginBottom: 20 }}>
                    <Typography.Text>Customize your response</Typography.Text>
                </div>
                <ChannelCreate onSuccess={onSuccess} onError={onError} />
            </Modal>
        </>
    )
}

export default NavCreate