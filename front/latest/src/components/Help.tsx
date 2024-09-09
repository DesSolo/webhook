import { Space, Typography } from 'antd';


const Help = (props: { channel: string }) => {
    const channelURL = () => {
        return import.meta.env.VITE_WEBHOOK_URL + props.channel
    }

    const items = [
        {
            label: "Your uniq URL",
            value: channelURL(),
        },
        {
            label: "Example cURL",
            value: "curl --request POST --header 'Content-Type: application/json' --url " + channelURL() + " --data '{\"key\": \"value\"}'",
        }
    ]
    return (
        <Space direction="vertical">
            {items.map((item, index) => (
                <Space key={index} direction="vertical">
                    <Typography.Text strong>{item.label}</Typography.Text>
                    <Typography.Text code copyable>{item.value}</Typography.Text>
                </Space>
            ))}
        </Space>
    )
}


export default Help

// export {WEBHOOK_URL}
