const RequestMethod = (props: { method: string }) => {

    let style = {
        fontWeight: "bold",
        borderRadius: "4px",
        padding: "2px 6px 2px 6px",
        color: "#fff",
        background: "red"
    }

    switch (props.method) {
        case "GET":
            style.background = "#5cb85c"
            break
        case "POST":
            style.background = "#5bc0de"
            break
        case "DELETE":
            style.background = "#d9534f"
            break
        case "PUT":
            style.background = "#777"
            break
        case "PATCH":
            style.background = "#f0ad4e"
            break
    }

    return (
        <span style={style}>{props.method}</span>
    )
}

export default RequestMethod