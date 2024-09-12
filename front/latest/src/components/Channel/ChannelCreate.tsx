import { Form, FormProps, Input, InputNumber, Select } from "antd"

type ChannelCreateFields = {
    kind: string
    simple: {
        status_code: number
        content_type: string
        content: string
        timeout: number
    }
}

const onFinishFailed: FormProps<ChannelCreateFields>['onFinishFailed'] = (errorInfo) => {
    console.log('Failed:', errorInfo);
};

const ChannelCreate = (props: {onSuccess: any, onError: any}) => {
    const onFinish = (values: ChannelCreateFields) => {
        fetch('/api/v1/channel', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(values)
        })
        .then( async(response) => {
            if (!response.ok) {
                try {
                    let data = await response.json()
                    props.onError(data.error || data)
                    return
                } catch (error) {
                    props.onError(response.statusText)
                    return
                }
            }

            let data = await response.json()

            props.onSuccess(data.token)
        })
        .catch((error) => {
            console.log("HHERE")
            props.onError(error);
        })
    }
    
    const renderSimple = () => {
        return (
            <>
                <Form.Item<ChannelCreateFields> label="Status code" name={["simple", "status_code"]}>
                    <InputNumber />
                </Form.Item>
                <Form.Item<ChannelCreateFields> label="Content type" name={["simple", "content_type"]}>
                    <Input />
                </Form.Item>
                <Form.Item<ChannelCreateFields> label="Content" name={["simple", "content"]}>
                    <Input.TextArea />
                </Form.Item>
                <Form.Item<ChannelCreateFields> label="Timeout" name={["simple", "timeout"]} help="time in seconds">
                    <InputNumber />
                </Form.Item>
            </>
        )
    }

    return (
        <Form
            name="createRequest"
            labelCol={{ span: 6 }}
            wrapperCol={{ span: 16 }}
            onFinish={onFinish}
            onFinishFailed={onFinishFailed}
            initialValues={{
                kind: "simple",
                simple: {
                    status_code: 200,
                    content_type: "text/plain",
                    content: "",
                    timeout: 0
                }
            }}
        >
            <Form.Item<ChannelCreateFields> label="Kind" name="kind">
                <Select>
                    <Select.Option value="simple">Simple</Select.Option>
                </Select>
            </Form.Item>
            {renderSimple()}
        </Form>
    )
}

export default ChannelCreate