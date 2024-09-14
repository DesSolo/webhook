import { CopyOutlined, DownOutlined } from "@ant-design/icons"
import { Dropdown, Space } from "antd"

const NavCopy = (props: { uuid: string }) => {
    const renderLabel = (name : string, text: string) => {
        return <div onClick={() => navigator.clipboard.writeText(text)}>
            <div>{name}</div>
            <div style={{ fontSize: 12, color: "#777" }}>{text}</div>
        </div>
    }

    const items = [
        {
            key: "url",
            label: renderLabel("URL", import.meta.env.VITE_WEBHOOK_URL + props.uuid),
        },
        {
            key: "id",
            label: renderLabel("ID", props.uuid),
        }
    ]

    return (
        <Dropdown menu={{ items }}>
            <a onClick={(e) => e.preventDefault()}>
                <Space style={{color: "white"}}>
                    <CopyOutlined />Copy<DownOutlined />
                </Space>
            </a>
        </Dropdown>
    )
}

export default NavCopy